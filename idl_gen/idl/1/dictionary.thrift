namespace go admin.dictionary

include "../../base/base.thrift"


// 字典列表请求数据
struct DictionaryPageReq {
    1:  optional string title (api.raw = "title" )
    2:  optional string name (api.raw = "name" )
    3:  i64 page (api.raw = "page" )
    4:  i64 pageSize (api.raw = "pageSize")
}

//字典名获取字典键值请求数据
struct DictionaryDetailReq{
    1:  optional string name (api.raw = "name" )
    2:  optional i64 dictionaryId (api.raw = "dictionaryId" )

}
// 字典信息
struct DictionaryInfo {
    1:  i64 id (api.raw = "id" )
    2:  string title (api.raw = "title" )
    3:  string name (api.raw = "name" )
    5:  i64 status (api.raw = "status" )
    6:  string description (api.raw = "description" )
    7:  string createdAt (api.raw = "createdAt" )
    8:  string updatedAt (api.raw = "updatedAt" )
}
// 字典键值信息
struct DictionaryDetail {
    1:  i64 id (api.raw = "id" )
    2:  string title (api.raw = "title" )
    3:  string key (api.raw = "key" )
    4:  string value (api.raw = "value" )
    5:  i64 status (api.raw = "status" )
    6:  string createdAt (api.raw = "createdAt" )
    7:  string updatedAt (api.raw = "updatedAt" )
    8:  i64 parentId (api.raw = "parentId" )
}
// dictionary service
service dictionaryService {
  // 创建字典信息
  base.NilResponse CreateDictionary(1: DictionaryInfo req) (api.post = "/service/dict/create")

  // 更新字典信息
  base.NilResponse UpdateDictionary(1: DictionaryInfo req) (api.post = "/service/dict/update")

  // 删除字典信息
  base.NilResponse DeleteDictionary(1: base.IdReq req) (api.post =  "/service/dict")

  // 获取字典列表
  base.NilResponse DictionaryList(1: DictionaryPageReq req) (api.get = "/service/dict/list")

  // 创建字典键值信息
  base.NilResponse CreateDictionaryDetail(1: DictionaryDetail req) (api.post = "/service/dict/detail/create")

  // 更新字典键值信息
  base.NilResponse UpdateDictionaryDetail(1: DictionaryDetail req) (api.post = "/service/dict/detail/update")

  // 删除字典键值信息
  base.NilResponse DeleteDictionaryDetail(1: base.IdReq req) (api.get = "/service/dict/detail")

  // 根据字典名获取字典键值列表
  base.NilResponse DetailByDictionaryList(1: DictionaryDetailReq req) (api.post = "/service/dict/detail/list")

}