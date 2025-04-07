package do

import "kcers/idl_gen/model/member"

type Member interface {
	Create(req *member.CreateOrUpdateMemberReq) error
	Update(req *member.CreateOrUpdateMemberReq) error
	Delete(id int64) error
	List(req *member.MemberListReq) (resp []*member.MemberInfo, total int, err error)
	Info(id int64) (info *member.MemberInfo, err error)
	ChangePassword(id int64, oldPassword, newPassword string) error
	UpdateStatus(id, status int64) error
	Search(option, value string) (memberInfo *member.MemberInfo, err error)

	ProductSearch(members []int64) (info *member.MemberProductInfo, err error)
	PropertySearch(memberProducts []int64) (info *member.MemberPropertyInfo, err error)

	ProductList(req *member.MemberProductListReq) (resp []*member.MemberProductInfo, total int, err error)
	ProductDetail(id int64) (info *member.MemberProductInfo, err error)
	ProductUpdate(req *member.MemberProductInfo) error
	ProductUpdateStatus(id int64, status int64) error

	PropertyList(req *member.MemberPropertyListReq) (resp []*member.MemberPropertyInfo, total int, err error)
	PropertyDetail(id int64) (info *member.MemberPropertyInfo, err error)
	PropertyUpdate(req *member.MemberPropertyInfo) error
	PropertyUpdateStatus(id int64, status int64) error

	ContractList(req *member.MemberContractListReq) (resp []*member.MemberContractInfo, total int, err error)
}

// MPStatus 会员产品状态
type MPStatus int

const (
	MPStatusUnfinished = iota
	MPStatusNotActivated
	MPStatusActivated
	MPStatusExpire
	MPStatusExhaust
	MPStatusUpgrade
	MPStatusFreeze
)

var MPStatusNames = map[MPStatus]string{
	MPStatusUnfinished:   "未完成",
	MPStatusNotActivated: "未激活",
	MPStatusActivated:    "已激活",
	MPStatusExpire:       "已到期",
	MPStatusExhaust:      "已完结",
	MPStatusUpgrade:      "已升级",
	MPStatusFreeze:       "已冻结",
}

func (s MPStatus) String() string {
	return MPStatusNames[s]
}
