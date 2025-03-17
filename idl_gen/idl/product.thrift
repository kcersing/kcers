namespace go admin.product

include "../base/base.thrift"


struct Property {
    1: i64 productId (api.raw = "productId")
    /**名称 */
    2: string name  (api.raw = "name")
    /**定价 */
    3: double price  (api.raw = "price")
    /**时长 */
    4: i64 duration  (api.raw = "duration")
    /**单次时长 */
    5: i64 length (api.raw = "length")
    /**次数 */
    6: i64 count  (api.raw = "count")
    /**类型 */
    7: string type  (api.raw = "type")
    /**数据 */
    8: string data (api.raw = "data")
    9: i64 id  (api.raw = "id")
}

struct Product {
    /**商品名 */
    1: string name (api.raw = "data")
    /**主图 */
    2: string pic (api.raw = "data")
    /**详情 */
    3: string description (api.raw = "data")
    /**属性 */
    4: list <Property> property (api.raw = "data")
    /**价格 */
    5: double price (api.raw = "data")
    /**库存 */
    6: i64 stock (api.raw = "data")
    /**状态*/
    7: i64 status (api.raw = "data")
    8: i64 id (api.raw = "data")
}

struct CreateOrUpdatePropertyReq {
    1: optional i64 id (api.raw = "id")
    /**名称*/
    2: optional string name (api.raw = "name")
    /**定价*/
    3: optional double price (api.raw = "price")
    /**时长*/
    4: optional i64 duration (api.raw = "duration")
    /**单次时长*/
    5: optional i64 length (api.raw = "length")
    /**次数*/
    6: optional i64 count (api.raw = "count")
    /**类型*/
    7: optional string type (api.raw = "type")
    8: optional list<i64> venueId (api.raw = "venueId")

}

struct CreateOrUpdateReq {
    1: optional i64 id (api.raw = "id")
    /**商品名*/
    2: optional string name (api.raw = "name")
    /**主图*/
    3: optional string pic (api.raw = "pic")
    /**详情*/
    4: optional string description  (api.raw = "description")
    /**价格*/
    5: optional double price (api.raw = "price")
    /**库存*/
    6: optional i64 stock (api.raw = "stock")
    /**属性*/
    7:optional  list<i64> propertys (api.raw = "propertys")

    10: optional i64 CreateId (api.raw = "createId")
}

struct ListReq {
    1: i64 page (api.raw = "page")
    2: i64 pageSize (api.raw = "pageSize")
    3: optional string name (api.raw = "name")
    4: optional list<i64> status (api.raw = "status")
    5: optional list<i64> venue (api.raw = "venue")
    6: optional list<string> createdTime (api.raw = "createdTime")
    7: optional string type (api.raw = "type")
}
struct PropertyListReq{
    1: i64 page (api.raw = "page")
    2: i64 pageSize (api.raw = "pageSize")
    3: optional string name (api.raw = "name")
    4: optional list<i64> status (api.raw = "status")
    5: optional list<i64> venue (api.raw = "venue")
    6: optional list<string> createdTime (api.raw = "createdTime")
    7: optional string type (api.raw = "type")
}

service ProductService {
    // 添加属性
    base.NilResponse CreateProperty(1: CreateOrUpdatePropertyReq req) (api.post = "/api/admin/property/create")
    // 编辑属性
    base.NilResponse UpdateProperty(1: CreateOrUpdatePropertyReq req) (api.post = "/api/admin/property/update")
    // 删除属性
    base.NilResponse DeleteProperty(1: base.IdReq req) (api.post = "/api/admin/property/del")
    // 商品列表
    base.NilResponse ListProperty(1: PropertyListReq req) (api.post = "/api/admin/property/list")

    // 添加商品
    base.NilResponse Create(1: CreateOrUpdateReq req) (api.post = "/api/admin/product/create")
    // 编辑商品
    base.NilResponse Update(1: CreateOrUpdateReq req) (api.post = "/api/admin/product/update")
    // 删除商品
    base.NilResponse Delete(1: base.IdReq req) (api.post = "/api/admin/product/del")
    // 商品列表
    base.NilResponse List(1: ListReq req) (api.post = "/api/admin/product/list")
    // 上架0/下架1
    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/product/status")

    // 商品详情
    base.NilResponse InfoById(1: base.IdReq req) (api.post = "/api/admin/product/info")



}



