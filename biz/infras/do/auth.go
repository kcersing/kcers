package do

type Auth interface {
	UpdateApiAuth(roleIDStr string, apis []int64) error
	ApiAuth(roleIDStr string) (infos []*ApiAuthInfo, err error)
	UpdateMenuAuth(roleID int64, menuIDs []int64) error
	MenuAuth(roleID int64) (menuIDs []int64, err error)
}
