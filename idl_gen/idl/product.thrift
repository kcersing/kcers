namespace go product

include "../base/base.thrift"


struct Property {
    1:optional i64 productId=0 (api.raw = "productId")
    /**名称 */
    2:optional string name ="" (api.raw = "name")
    /**定价 */
    3:optional double price=0   (api.raw = "price")
    /**时长 */
    4:optional i64 duration =0  (api.raw = "duration")
    /**单次时长 */
    5:optional i64 length =0 (api.raw = "length")
    /**次数 */
    6:optional i64 count =0  (api.raw = "count")
    /**类型 */
    7:optional string type="" (api.raw = "type")
    /**数据 */
    8:optional string data=""  (api.raw = "data")
    9:optional i64 id =0  (api.raw = "id")

        /**标签*/
        18: optional list<base.List> tags = {}  (api.raw = "tags")
        /**合同*/
        19: optional list<base.List> contracts = {} (api.raw = "contracts")

}

struct Product {
    /**商品名 */
    1:optional string name="" (api.raw = "name")
    /**主图 */
    2:optional string pic (api.raw = "pic")
    /**详情 */
    3:optional string description (api.raw = "description")
    /**属性 */
    4:optional list <Property> propertys (api.raw = "propertys")
    /**价格 */
    5:optional double price (api.raw = "price")
    /**库存 */
    6:optional i64 stock (api.raw = "stock")
    /**状态*/
    7:optional i64 status (api.raw = "status")
    8:optional i64 id (api.raw = "id")
    12: optional i64 isSales=0 (api.raw = "isSales")
    13: optional string signSalesAt = "" (api.raw = "signSalesAt")
    14: optional string endSalesAt = "" (api.raw = "endSalesAt")
    16: optional string createdAt = ""  (api.raw = "createdAt")
    17: optional string updatedAt = "" (api.raw = "updatedAt")
    20: optional i64 createdId = 0 (api.raw = "createdId")
    21: optional string createdName = "" (api.raw = "createdName")
}

struct CreateOrUpdatePropertyReq {
    1: optional i64 id =0 (api.raw = "id")
    /**名称*/
    2: optional string name ="" (api.raw = "name")
    /**定价*/
    3: optional double price =0(api.raw = "price")
    /**时长*/
    4: optional i64 duration =0(api.raw = "duration")
    /**单次时长*/
    5: optional i64 length =0(api.raw = "length")
    /**次数*/
    6: optional i64 count=0 (api.raw = "count")
    /**类型*/
    7: optional string type ="" (api.raw = "type")
    8: optional list<i64> venueId =0(api.raw = "venueId")

    /**标签-数组*/
    16: optional list<i64> tagId=0   (api.raw = "tagId")
    /**合同-数组*/
    17: optional list<i64> contractId =0 (api.raw = "contractId")

}

struct CreateOrUpdateReq {
    1: optional i64 id =0(api.raw = "id")
    /**商品名*/
    2: optional string name (api.raw = "name")
    /**主图*/
    3: optional string pic (api.raw = "pic")
    /**详情*/
    4: optional string description ="" (api.raw = "description")
    /**价格*/
    5: optional double price =0(api.raw = "price")
    /**库存*/
    6: optional i64 stock =0(api.raw = "stock")
    /**属性*/
    7:optional  list<i64> propertys =0(api.raw = "propertys")

    10: optional i64 CreateId=0 (api.raw = "createId")

    /**销售方式 1会员端*/
    12: optional i64 isSales = 0 (api.raw = "isSales")
    /**销售开始时间*/
    13: optional string signSalesAt ="" (api.raw = "signSalesAt")
    /**销售结束时间*/
    14: optional string endSalesAt ="" (api.raw = "endSalesAt")

}

struct ListReq {
    1: i64 page=0 (api.raw = "page")
    2: i64 pageSize =10(api.raw = "pageSize")
    3: optional string name="" (api.raw = "name")
    4: optional list<i64> status =0 (api.raw = "status")
    5: optional list<i64> venueId =0  (api.raw = "venue")
    6: optional list<string> createdTime=0  (api.raw = "createdTime")
    7: optional string type ="" (api.raw = "type")
}
struct PropertyListReq{
    1: i64 page=0 (api.raw = "page")
    2: i64 pageSize=10 (api.raw = "pageSize")
    3: optional string name (api.raw = "name")
    4: optional list<i64> status =0(api.raw = "status")
    5: optional list<i64> venueId =0  (api.raw = "venue")
    6: optional list<string> createdTime=0  (api.raw = "createdTime")
    7: optional string type ="" (api.raw = "type")
}

service ProductService {
    // 添加属性
    base.NilResponse CreateProperty(1: CreateOrUpdatePropertyReq req) (api.post = "/service/property/create")
    // 编辑属性
    base.NilResponse UpdateProperty(1: CreateOrUpdatePropertyReq req) (api.post = "/service/property/update")
    // 删除属性
    base.NilResponse DeleteProperty(1: base.IDReq req) (api.post = "/service/property/del")
    // 商品列表
    base.NilResponse ListProperty(1: PropertyListReq req) (api.post = "/service/property/list")

    // 添加商品
    base.NilResponse Create(1: CreateOrUpdateReq req) (api.post = "/service/product/create")
    // 编辑商品
    base.NilResponse Update(1: CreateOrUpdateReq req) (api.post = "/service/product/update")
    // 删除商品
    base.NilResponse Delete(1: base.IDReq req) (api.post = "/service/product/del")
    // 商品列表
    base.NilResponse List(1: ListReq req) (api.post = "/service/product/list")
    // 上架0/下架1
    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/product/status")

    // 商品详情
    base.NilResponse InfoById(1: base.IDReq req) (api.post = "/service/product/info")

    base.NilResponse ProductListExport(1: ListReq req) (api.post = "/service/product/export")

}



