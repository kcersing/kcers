package member

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"kcers/biz/dal/cache"
	"kcers/biz/dal/config"
	db "kcers/biz/dal/db/mysql"
	"kcers/biz/dal/db/mysql/ent"
	member2 "kcers/biz/dal/db/mysql/ent/member"
	"kcers/biz/dal/db/mysql/ent/memberproduct"
	"kcers/biz/dal/db/mysql/ent/memberproductproperty"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"kcers/biz/dal/enums"
	"kcers/biz/dal/minio"
	"kcers/biz/infras/do"
	"kcers/biz/infras/service"
	"kcers/biz/infras/service/system"
	userService "kcers/biz/infras/service/user"
	"kcers/biz/pkg/encrypt"
	"kcers/idl_gen/model/member"
	"strconv"
	"time"
)

type Member struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (m *Member) Login(req *member.LoginReq) (resp *member.MemberInfo, err error) {
	memberEnt, err := m.db.Member.Query().Where(member2.MobileEQ(req.Mobile)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get member failed")
		return nil, err
	}
	resp = m.entMemberInfo(*memberEnt)

	return resp, nil
}

func (m *Member) Create(req *member.CreateOrUpdateMemberReq) error {

	birthday, _ := time.Parse(time.DateOnly, req.Birthday)
	var gender = enums.ReturnMemberGenderKey(req.Gender)

	tx, err := m.db.Tx(m.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}

	mer := tx.Member.Create().
		SetMobile(req.Mobile).
		SetName(req.Name).
		SetUsername(req.Username).
		SetCondition(0)

	noe, err := mer.Save(m.ctx)

	if err != nil {
		err = service.Rollback(tx, errors.Wrap(err, "create user failed"))
		return err
	}

	_, err = tx.MemberProfile.Create().
		SetMember(noe).
		SetGender(gender).
		SetBirthday(birthday).
		SetSource(req.Source).
		SetIntention(req.Intention).
		SetEmail(req.Email).
		SetWecom(req.Wecom).
		Save(m.ctx)

	if err != nil {
		err = service.Rollback(tx, errors.Wrap(err, "create Member Profile failed"))
		return err
	}
	_, err = tx.MemberDetails.Create().
		SetMember(noe).
		Save(m.ctx)

	if err != nil {
		err = service.Rollback(tx, errors.Wrap(err, "create Member Details failed"))
		return err
	}
	_, err = tx.Face.Create().
		SetMemberFaces(noe).
		Save(m.ctx)

	if err != nil {
		err = service.Rollback(tx, errors.Wrap(err, "create Face failed"))
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (m *Member) UpdateMemberAvatar(id int64, avatar string) error {
	_, err := m.db.Member.Update().Where(member2.IDEQ(id)).SetAvatar(avatar).Save(m.ctx)
	return err
}
func (m *Member) UpdateMemberPassword(id int64, password string) error {
	password, _ = encrypt.Crypt(password)
	_, err := m.db.Member.Update().Where(member2.IDEQ(id)).SetPassword(password).Save(m.ctx)
	return err
}
func (m *Member) ChangePassword(ID int64, oldPassword, newPassword string) error {

	targetMember, err := m.db.Member.Query().Where(member2.IDEQ(ID)).First(m.ctx)
	if err != nil {
		return errors.Wrap(err, "targetUser not found")
	}
	// check old password
	if ok := encrypt.VerifyPassword(oldPassword, targetMember.Password); !ok {
		err = errors.New("wrong old password")
		return err
	}
	// update password
	password, _ := encrypt.Crypt(newPassword)
	_, err = m.db.Member.Update().Where(member2.IDEQ(ID)).SetPassword(password).Save(m.ctx)

	return err
}
func (m *Member) Update(req *member.CreateOrUpdateMemberReq) error {

	tx, err := m.db.Tx(m.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}
	up := tx.Member.Update().
		Where(member2.IDEQ(req.ID)).
		SetName(req.Name).
		SetUsername(req.Username).
		SetMobile(req.Mobile)

	_, err = up.Save(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "update member failed")
		return err
	}

	birthday, _ := time.Parse(time.DateOnly, req.Birthday)
	var gender = enums.ReturnMemberGenderKey(req.Gender)

	_, err = tx.MemberProfile.Update().
		SetGender(gender).
		SetBirthday(birthday).
		SetSource(req.Source).
		SetIntention(req.Intention).
		SetEmail(req.Email).
		SetWecom(req.Wecom).
		Save(m.ctx)

	if err != nil {
		err = errors.Wrap(err, "update member profile failed")
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (m *Member) Delete(id int64) error {
	_, err := m.db.Member.Update().Where(member2.IDEQ(id)).SetDelete(1).Save(m.ctx)
	return err
}

func (m *Member) List(req *member.MemberListReq) (resp []*member.MemberInfo, total int, err error) {
	var predicates []predicate.Member
	if req.Name != "" {
		predicates = append(predicates, member2.NameEQ(req.Name))
	}
	lists, err := m.db.Member.Query().Where(predicates...).
		Order(ent.Desc(member2.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, m.entMemberInfo(*v))
	}

	total, _ = m.db.Member.Query().Where(predicates...).Count(m.ctx)
	return
}
func (m *Member) entMemberProfile(v ent.Member) (profile *member.MemberProfile) {
	p, _ := v.QueryMemberProfile().First(m.ctx)

	var gender = "未知"
	//var gradeName, intentionName, sourceName, createdName string
	var gradeName, intentionName, sourceName string
	var age int64
	if p != nil {
		gender = enums.ReturnMemberGenderValues(p.Gender)

		if !p.Birthday.IsZero() {
			age = int64(time.Now().Sub(p.Birthday).Hours() / 24 / 365)
		}
		if p.Intention > 0 {
			intentionName = system.NewDictionary(m.ctx, m.c).GetDictionaryDetailTitle(p.Intention)
		}
		if p.Source > 0 {
			sourceName = system.NewDictionary(m.ctx, m.c).GetDictionaryDetailTitle(p.Source)
		}

		profile = &member.MemberProfile{
			Intention:     p.Intention,
			Source:        p.Source,
			GradeName:     gradeName,
			IntentionName: intentionName,
			SourceName:    sourceName,
			Email:         p.Email,
			Gender:        gender,
			Age:           age,
			Wecom:         p.Wecom,
			Birthday:      p.Birthday.Format(time.DateOnly),
			RelationUid:   p.RelationUID,
			RelationUname: p.RelationUname,
			RelationMid:   p.RelationMid,
			RelationMname: p.RelationMame,
		}
	}
	return profile
}
func (m *Member) entMemberDetail(v ent.Member) (detail *member.MemberDetail) {
	d, _ := v.QueryMemberDetails().First(m.ctx)
	detail = &member.MemberDetail{
		MoneySum:        d.MoneySum,
		ProductId:       d.ProductID,
		ProductName:     d.ProductName,
		EntrySum:        d.EntrySum,
		EntryLastAt:     d.EntryLastAt.Format(time.DateTime),
		EntryDeadlineAt: d.EntryDeadlineAt.Format(time.DateTime),
		ClassLastAt:     d.ClassLastAt.Format(time.DateTime),
	}
	return detail
}
func (m *Member) entMemberInfo(v ent.Member) *member.MemberInfo {

	var createdName = userService.NewUser(m.ctx, m.c).GetUserName(v.CreatedID)
	memberInfo := &member.MemberInfo{
		ID:            v.ID,
		Username:      v.Username,
		Mobile:        v.Mobile,
		Avatar:        minio.URLconvert(m.ctx, m.c, v.Avatar),
		Condition:     v.Condition,
		Status:        v.Status,
		ConditionName: enums.ReturnMemberConditionValues(v.Condition),
		CreatedAt:     v.CreatedAt.Format(time.DateTime),
		UpdatedAt:     v.UpdatedAt.Format(time.DateTime),
		CreatedId:     v.CreatedID,
		CreatedName:   createdName,
	}

	return memberInfo

}
func (m *Member) Info(id int64) (info *member.MemberInfo, err error) {

	memberEnt, err := m.db.Member.Query().Where(member2.IDEQ(id)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get member failed")
		return info, err
	}

	info = m.entMemberInfo(*memberEnt)
	return
}

func (m *Member) Search(option, value string) (memberInfo *member.MemberInfo, err error) {

	switch option {
	case "1":
		id, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return memberInfo, errors.New("会员账号ID不正确")
		}
		return m.Info(id)
	case "2":
		id, err := m.db.Member.Query().Where(member2.Mobile(value)).FirstID(m.ctx)
		if err != nil {
			return memberInfo, errors.New("未找到该会员")
		}
		return m.Info(id)
	case "3":
		p, err := m.db.MemberProduct.Query().Where(memberproduct.SnEQ(value)).First(m.ctx)
		if err != nil {
			return memberInfo, errors.New("未找到该会员商品")
		}
		return m.Info(p.MemberID)
	case "4":
		p, err := m.db.MemberProductProperty.Query().Where(memberproductproperty.SnEQ(value)).First(m.ctx)
		if err != nil {
			return memberInfo, errors.New("未找到该会员商品属性")
		}
		return m.Info(p.MemberID)
	default:
		return memberInfo, errors.New("搜索类型不规范")
	}

}

func (m *Member) UpdateStatus(ID int64, status int64) error {
	_, err := m.db.Member.Update().Where(member2.IDEQ(ID)).SetStatus(status).Save(m.ctx)
	return err
}

func NewMember(ctx context.Context, c *app.RequestContext) do.Member {
	return &Member{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
