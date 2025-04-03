package do

type Login interface {
	UserLogin(username, password string) (res *LoginResp, err error)
	MemberLogin(username, password string) (res *LoginResp, err error)
}

type LoginResp struct {
	Id        int64  `json:"Id"`
	Username  string `json:"username"`
	RoleName  string `json:"roleName"`
	RoleValue string `json:"roleValue"`
	RoleID    int64  `json:"roleID"`
}
