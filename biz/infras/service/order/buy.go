package order

import (
	"bytes"
	"context"
	"github.com/pkg/errors"
	"kcers/biz/dal/config"
	"kcers/biz/dal/db/mysql/ent"
	member2 "kcers/biz/dal/db/mysql/ent/member"
	order2 "kcers/biz/dal/db/mysql/ent/order"
	product2 "kcers/biz/dal/db/mysql/ent/product"
	productproperty2 "kcers/biz/dal/db/mysql/ent/productproperty"
	user2 "kcers/biz/dal/db/mysql/ent/user"
	venue2 "kcers/biz/dal/db/mysql/ent/venue"
	"kcers/biz/dal/minio"
	"kcers/biz/infras/service"
	"kcers/biz/pkg/utils"
	"kcers/idl_gen/model/order"
	"strconv"
	"sync"
	"time"
)

func (o *Order) Buy(req *order.BuyReq) (string, error) {

	one := &ent.Order{}
	mp := &ent.MemberProduct{}
	var err error

	userId, exist := o.c.Get("user_id")
	if !exist || userId == nil {
		return "", errors.Wrap(err, "Unauthorized")
	}

	uId, ok := userId.(string)
	if !ok {
		return "", errors.Wrap(err, "user id 转换失败")
	}
	uid, _ := strconv.ParseInt(uId, 10, 64)

	create, err := o.db.User.Query().Where(user2.IDEQ(uid)).First(o.ctx)
	if err != nil {
		return "", errors.Wrap(err, "未找到创建人")
	}

	venue, err := o.db.Venue.Query().Where(venue2.IDEQ(req.VenueId)).First(o.ctx)
	if err != nil {
		return "", errors.Wrap(err, "未找到场馆")
	}

	members, err := o.db.Member.Query().Where(member2.IDEQ(req.MemberId)).First(o.ctx)

	if err != nil {
		return "", errors.Wrap(err, "未找到会员")
	}

	products, err := o.db.Product.Query().Where(product2.IDEQ(req.ProductId)).First(o.ctx)
	if err != nil {
		return "", errors.Wrap(err, "未找到产品")
	}

	layout := time.DateOnly
	assignAt, err := time.Parse(layout, req.AssignAt)
	if err != nil {
		return "", errors.Wrap(err, "时间转换失败")
	}

	errChan := make(chan error, 99)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(5)

	tx, err := o.db.Tx(o.ctx)

	if err != nil {
		return "", errors.Wrap(err, "starting a transaction:")
	}

	one, err = tx.Order.Create().
		SetOrderSn(utils.CreateCn()).
		SetOrderVenues(venue).
		SetOrderMembers(members).
		SetOrderCreates(create).
		SetStatus(0).
		SetNature(req.NatureType).
		//SetSource(req.Source).
		//SetDevice(req.Device).
		Save(o.ctx)
	if err != nil {
		return "", service.Rollback(tx, errors.Wrap(err, "创建 Order 失败"))
	}
	mp, err = tx.MemberProduct.Create().
		SetProductID(req.ProductId).
		// SetVenue (req.Venue).
		SetSn(utils.CreateCn()).
		SetMembers(members).
		SetVenueID(venue.ID).
		SetOrderID(one.ID).
		SetName(products.Name).
		SetStatus(0).
		Save(o.ctx)
	if err != nil {
		return "", service.Rollback(tx, errors.Wrap(err, "创建会员产品失败"))
	}
	_, err = tx.Order.Update().
		Where(order2.ID(one.ID)).
		SetMemberProductID(mp.ID).
		SetOrderVenues(venue).
		Save(o.ctx)

	if err != nil {
		return "", service.Rollback(tx, errors.Wrap(err, "创建 Order 失败"))
	}

	go func() {
		_, err := tx.OrderItem.Create().
			SetOrder(one).
			SetProductID(products.ID).
			//SetData(fmt.Sprintf("%+v", req)).
			SetData(*req).
			Save(o.ctx)
		if err != nil {
			err = errors.Wrap(err, "创建 Order Item 失败")
			errChan <- err
		}

		wg.Done()
	}()

	go func() {
		_, err = tx.OrderAmount.Create().
			SetOrder(one).
			SetTotal(req.Total).
			SetResidue(req.Total).
			Save(o.ctx)
		if err != nil {
			err = errors.Wrap(err, "创建Order Amount失败")
			errChan <- err
		}

		wg.Done()
	}()

	go func() {
		for _, v := range req.StaffsId {
			_, err = tx.OrderSales.Create().
				SetOrder(one).
				SetSalesID(v.ID).
				SetRatio(v.Ratio).
				SetMemberID(members.ID).
				Save(o.ctx)
			if err != nil {
				err = errors.Wrap(err, "创建Order Sales失败")
				errChan <- err
			}
		}

		wg.Done()
	}()

	go func() {

		if len(req.Propertys) > 0 {
			for _, v := range req.Propertys {
				if v.PropertyId > 0 {
					p, err := o.db.ProductProperty.Query().
						Where(productproperty2.IDEQ(v.PropertyId)).
						First(o.ctx)
					if err != nil {
						err = errors.Wrap(err, "查询Product Property失败")
						errChan <- err
					} else {
						venues, err := p.QueryVenues().All(o.ctx)
						if err != nil {
							err = errors.Wrap(err, "查询Product Property venues失败")
							errChan <- err
						}

						_, err = tx.MemberProductProperty.Create().
							SetMemberID(members.ID).
							SetOwner(mp).
							SetType(p.Type).
							SetPropertyID(p.ID).
							SetPrice(p.Price).
							SetDuration(p.Duration).
							SetLength(p.Length).
							SetName(p.Name).
							SetCount(v.Quantity).
							SetCountSurplus(v.Quantity).
							AddVenues(venues...).
							SetValidityAt(assignAt).
							SetStatus(0).
							Save(o.ctx)
						//tx.MemberProductProperty.Update().Where(memberproductproperty.ID(pp.ID)).SetSn(pp.ID).Save()
						if err != nil {
							err = errors.Wrap(err, "创建Member Product Property失败")
							errChan <- err
						}
					}
				}
			}
		}
		wg.Done()
	}()

	//go func() {
	//	signImg := ""
	//
	//
	//	for i, v := range req.ContractId {
	//		contracts, err := o.db.Contract.Query().
	//			Where(contract2.IDEQ(v)).
	//			First(o.ctx)
	//		if err != nil {
	//			err = errors.Wrap(err, "合同 Contract 失败")
	//			errChan <- err
	//		}
	//		memberContract, err := tx.MemberContract.Create().
	//			SetMember(members).
	//			SetOrderID(one.ID).
	//			SetContractID(contracts.ID).
	//			SetMemberProduct(mp).
	//			SetName(contracts.Name).
	//			SetSign(signImg).
	//			SetStatus(0).
	//			Save(o.ctx)
	//		if err != nil {
	//			err = errors.Wrap(err, "创建Member Contract "+string(i)+" 失败")
	//			errChan <- err
	//		}
	//
	//		_, err = tx.MemberContractContent.Create().
	//			SetContract(memberContract).
	//			SetContent(contracts.Content).
	//			Save(o.ctx)
	//
	//		if err != nil {
	//			err = errors.Wrap(err, "创建 Member Contract Content 失败")
	//			errChan <- err
	//		}
	//	}
	//
	//	wg.Done()
	//}()

	wg.Wait()
	select {
	case result := <-errChan:
		return "", service.Rollback(tx, result)
	default:
	}

	if err = tx.Commit(); err != nil {
		return "", err
	}
	return one.OrderSn, nil

}

func OrderSignImg(pic string) (img string, err error) {

	filename := minio.NewFileName(0, time.Now().UnixMicro()) + ".png"
	buffer := new(bytes.Buffer)
	// 将字符串写入Buffer
	_, err = buffer.WriteString(pic)
	if err != nil {
		err = errors.Wrap(err, "写入失败")
		return "", err
	}
	uploadInfo, err := minio.PutToBucketByBuf(context.Background(), config.GlobalServerConfig.Minio.ImgBucketName, filename, buffer)

	if err != nil {
		err = errors.Wrap(err, "签名上传失败")
		return "", err
	}
	signImg := config.GlobalServerConfig.Minio.ImgBucketName + "/" + uploadInfo.Key
	return signImg, nil
}
