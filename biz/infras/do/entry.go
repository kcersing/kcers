package do

import "kcers/idl_gen/model/entry"

type Entry interface {
	Create(logsReq *entry.EntryInfo) error
	List(req *entry.EntryListReq) (list []*entry.EntryInfo, total int, err error)
}
