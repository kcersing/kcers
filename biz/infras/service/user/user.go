package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"kcers/biz/dal/cache"
	"kcers/biz/dal/config"
	db "kcers/biz/dal/db/mysql"
	"kcers/biz/dal/db/mysql/ent"
	"kcers/biz/dal/db/mysql/ent/predicate"
	user2 "kcers/biz/dal/db/mysql/ent/user"
	"kcers/biz/dal/enums"
	"kcers/biz/infras/do"
	"kcers/biz/infras/service"
	"kcers/biz/pkg/encrypt"
	"kcers/idl_gen/model/dictionary"
	"kcers/idl_gen/model/user"
	"time"
)

type User struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (u *User) SetDefaultVenue(id, venueId int64) error {
	_, err := u.db.User.Update().
		Where(user2.IDEQ(id)).
		SetDefaultVenueID(venueId).
		Save(u.ctx)

	if err != nil {
		err = errors.Wrap(err, "update DefaultVenue  ID   failed")
		return err
	}

	return nil
}

func (u *User) SetRole(id int64, roleID []int64) error {
	_, err := u.db.User.Update().
		Where(user2.IDEQ(id)).
		AddRoleIDs(roleID...).
		Save(u.ctx)

	if err != nil {
		err = errors.Wrap(err, "update user role failed")
		return err
	}

	return nil
}

func (u *User) GetUserName(id int64) (name string) {
	if id == 0 {
		return ""
	}
	first, _ := u.db.User.Query().Where(user2.IDEQ(id)).First(u.ctx)
	if first != nil {
		return first.Name
	}
	return ""
}
func (u *User) Info(id int64) (info *user.UserInfo, err error) {
	//info = new(user.UserInfo)
	//userInterface, exist := u.cache.Get("userInfo" + strconv.Itoa(int(id)))
	//if exist {
	//	if u, ok := userInterface.(*user.UserInfo); ok {
	//		return u, nil
	//	}
	//}
	userEnt, err := u.db.User.Query().Where(user2.IDEQ(id)).First(u.ctx)
	if err != nil {
		err = errors.Wrap(err, "get user failed")
		return info, err
	}
	info = u.entUserInfo(*userEnt)
	//if !userEnt.Birthday.IsZero() {
	//	info.Age = int64(time.Now().Sub(userEnt.Birthday).Hours() / 24 / 365)
	//}
	//info.Birthday = userEnt.Birthday.Format(time.DateTime)
	//u.cache.SetWithTTL("userInfo"+strconv.Itoa(int(info.Id)), &info, 1, 1*time.Hour)
	return
}

func (u *User) Create(req *user.CreateOrUpdateUserReq) error {

	ok, _ := u.db.User.Query().Where(user2.Username(req.Username)).Exist(u.ctx)
	if ok {
		return errors.New("账号重复！")
	}
	ok, _ = u.db.User.Query().Where(user2.Mobile(req.Mobile)).Exist(u.ctx)
	if ok {
		return errors.New("手机号重复！")
	}

	var gender = enums.ReturnMemberGenderKey(req.Gender)

	parsedTime, _ := time.Parse(time.DateOnly, req.Birthday)
	password, _ := encrypt.Crypt(req.Password)
	tx, err := u.db.Tx(u.ctx)
	if err != nil {
		return errors.Wrap(err, "starting a transaction:")
	}
	noe, err := tx.User.Create().
		SetAvatar(req.Avatar).
		SetMobile(req.Mobile).
		SetEmail(req.Email).
		SetStatus(req.Status).
		SetUsername(req.Mobile).
		SetName(req.Name).
		SetBirthday(parsedTime).
		SetGender(gender).
		SetWecom(req.Wecom).
		SetJobTime(req.JobTime).
		SetPassword(password).
		SetFunctions(req.Functions).
		AddRoleIDs(req.RoleId...).
		SetDetail(req.Detail).
		AddTagIDs(req.UserTags...).
		AddVenueIDs(req.VenueId...).
		SetDefaultVenueID(req.DefaultVenueId).
		Save(u.ctx)

	if err != nil {
		err = service.Rollback(tx, errors.Wrap(err, "create user failed"))
		return err
	}
	_, err = tx.Face.Create().
		SetUserFaces(noe).
		Save(u.ctx)

	if err != nil {
		err = service.Rollback(tx, errors.Wrap(err, "create Face failed"))
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (u *User) Update(req *user.CreateOrUpdateUserReq) error {

	var gender = enums.ReturnMemberGenderKey(req.Gender)
	parsedTime, _ := time.Parse(time.DateOnly, req.Birthday)

	_, err := u.db.User.Update().
		Where(user2.IDEQ(req.ID)).
		SetAvatar(req.Avatar).
		SetMobile(req.Mobile).
		SetEmail(req.Email).
		SetStatus(req.Status).
		SetUsername(req.Mobile).
		SetName(req.Name).
		SetBirthday(parsedTime).
		SetGender(gender).
		SetStatus(1).
		SetWecom(req.Wecom).
		Save(u.ctx)

	if err != nil {
		err = errors.Wrap(err, "update user failed")
		return err
	}

	return nil
}

func (u *User) ChangePassword(userId int64, newPassword string) error {
	////get user info
	//targetUser, err := u.db.User.Query().Where(user.IDEQ(userID)).First(u.ctx)
	//if err != nil {
	//	return errors.Wrap(err, "targetUser not found")
	//}
	//// check old password
	//if ok := encrypt.VerifyPassword(oldPassword, targetUser.Password); !ok {
	//	err = errors.New("wrong old password")
	//	return err
	//}
	// update password
	password, _ := encrypt.Crypt(newPassword)
	_, err := u.db.User.Update().Where(user2.IDEQ(userId)).SetPassword(password).Save(u.ctx)

	return err
}

func (u *User) List(req *user.UserListReq) (userList []*user.UserInfo, total int, err error) {
	var predicates []predicate.User
	if req.Mobile != "" {
		predicates = append(predicates, user2.MobileEQ(req.Mobile))
	}

	if req.Name != "" {
		predicates = append(predicates, user2.Name(req.Name))
	}

	users, err := u.db.User.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(user2.FieldID)).
		Limit(int(req.PageSize)).All(u.ctx)
	if err != nil {
		err = errors.Wrap(err, "get user list failed")
		return userList, total, err
	}
	// copy to UserInfo struct

	for _, v := range users {
		mr := u.entUserInfo(*v)
		userList = append(userList, mr)
	}
	total, _ = u.db.User.Query().Where(predicates...).Count(u.ctx)
	return userList, total, nil
}

func (u *User) UpdateUserStatus(id, status int64) error {
	_, err := u.db.User.Update().Where(user2.IDEQ(id)).SetStatus(status).Save(u.ctx)
	return err
}

func (u *User) DeleteUser(id int64) error {
	_, err := u.db.User.Delete().Where(user2.IDEQ(id)).Exec(u.ctx)
	return err
}

func (u *User) entUserInfo(userEnt ent.User) (info *user.UserInfo) {
	info = &user.UserInfo{
		ID:             userEnt.ID,
		Status:         userEnt.Status,
		Username:       userEnt.Username,
		Name:           userEnt.Name,
		Mobile:         userEnt.Mobile,
		CreatedAt:      userEnt.CreatedAt.Format(time.DateTime),
		UpdatedAt:      userEnt.UpdatedAt.Format(time.DateTime),
		Detail:         userEnt.Detail,
		JobTime:        &userEnt.JobTime,
		DefaultVenueId: userEnt.DefaultVenueID,
	}

	info.Gender = enums.ReturnMemberGenderValues(userEnt.Gender)

	var venues []*user.Venues
	Venues, _ := userEnt.QueryVenues().All(u.ctx)
	if len(Venues) > 0 {
		for _, ve := range Venues {
			venues = append(venues, &user.Venues{
				ID:   ve.ID,
				Name: ve.Name,
			})

		}
	}
	info.Venues = venues
	Tags, _ := userEnt.QueryTags().All(u.ctx)
	if len(Tags) > 0 {
		for _, d := range Tags {

			info.UserTags = append(info.UserTags, &dictionary.DictionaryDetail{
				ID:     d.ID,
				Title:  d.Title,
				Key:    d.Key,
				Value:  d.Value,
				Status: d.Status,
			})
		}
	}

	roles, _ := userEnt.QueryRoles().All(u.ctx)
	if roles != nil {
		var userRole []*user.UserRole
		var userRoleIds []int64
		for _, v := range roles {
			userRole = append(userRole, &user.UserRole{
				Name:  v.Name,
				Value: v.Value,
				ID:    v.ID,
			})
			userRoleIds = append(userRoleIds, v.ID)
		}
		info.UserRoleIds = userRoleIds
		info.UserRole = userRole
	}

	return info
}

func (l *User) Login(req *user.LoginReq) (res *user.LoginResp, err error) {
	only, err := l.db.User.Query().
		Where(
			user2.UsernameEQ(req.Username),
			user2.Status(1),
		).Only(l.ctx)
	if err != nil {
		return nil, err
	}
	if only == nil {
		err = errors.New("login user not exist")
		return nil, err
	}
	//VerifyPassword
	if ok := encrypt.VerifyPassword(req.Password, only.Password); !ok {
		err = errors.New("wrong password")
		return nil, err
	}

	info, err := NewUser(l.ctx, l.c).Info(only.ID)
	if err != nil {
		return nil, err
	}

	res = new(user.LoginResp)
	res.Username = info.Username
	res.UserId = info.ID
	res.UserRole = info.UserRole
	res.UserRoleIds = info.UserRoleIds
	return res, err
}

func NewUser(ctx context.Context, c *app.RequestContext) do.User {
	return &User{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
