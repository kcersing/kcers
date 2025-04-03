package do

import "kcers/idl_gen/model/contract"

type Contract interface {
	Create(req *contract.ContractInfo) error
	Update(req *contract.ContractInfo) error
	UpdateStatus(ID, status int64) error
	Info(id int64) (info *contract.ContractInfo, err error)
	List(req *contract.ContractListReq) (list []*contract.ContractInfo, total int, err error)
}
