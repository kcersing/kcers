namespace go order

include "../base/base.thrift"


struct OrderItem {
        1:  optional i64 id=0 (api.raw = "id")
        2:  optional i64 productId=0 (api.raw = "productId")
        3:  optional i64 relatedUserProductId=0 (api.raw = "relatedUserProductId")
        6:  optional i64 orderId=0 (api.raw = "orderId")
        7:  optional string data="" (api.raw = "data")
        8:  optional string name="" (api.raw = "name")
        16:  optional string createdAt="" (api.raw = "createdAt")
        17:  optional string updatedAt="" (api.raw = "updatedAt")
        9: optional i64 number=1 (api.raw = "number")
}



struct propertyItem{
    1:optional i64 propertyId=0 (api.raw = "property")
    2:optional i64 quantity=0 (api.raw = "quantity")
    3:optional string type="" (api.raw = "type")
}
struct staffItem{
    1:optional i64 id=0 (api.raw = "quantity")
    2:optional i64 ratio=0 (api.raw = "quantity")
}

struct ListOrderReq {
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional string mobile="" (api.raw = "mobile")
    /**销售员工ID*/
    4: optional list<i64> sellId=0 (api.raw = "sellId")
    /**产品ID*/
    5: optional list<i64> productId=0 (api.raw = "productId")
    6: optional list<i64> venueId=0 (api.raw = "venueId")
    7: optional list<i64> status=0 (api.raw = "status")
    /**产品编号*/
    8: optional string orderSn="" (api.raw = "orderSn")
    /**订单完成时间*/
    9: optional string startCompletionAt="" (api.raw = "startCompletionAt")
    10: optional string endCompletionAt="" (api.raw = "endCompletionAt")
    /**产品类型*/
    11: optional string productType="" (api.raw = "productType")
    /**订单业务*/
    12: optional string nature="" (api.raw = "nature")
    /**产品名称*/
    13: optional string name ="" (api.raw = "name ")
    14: optional string memberName="" (api.raw = "memberName")
    15: optional i64 memberId=0 (api.raw = "memberId")
}

struct UpdateOrderReq {
    1: OrderItem order (api.raw = "order")
}

struct BuyReq{
    1: optional i64 venueId=0 (api.raw = "venueId")
    2: optional i64 productId=0 (api.raw = "productId")
    3: optional i64 memberId=0 (api.raw = "memberId")
    /**数量 默认1*/
    4: optional i64 number=1 (api.raw = "number")
    /**金额*/
    5: optional double total=0 (api.raw = "total")
    /**来源 [pc电脑端;wxc微信小程序] 默认wxc */
    6: optional string device="pc" (api.raw = "device")
    7: optional string assignAt ="" (api.raw = "assignAt")
    8: optional list<propertyItem> propertys ={} (api.raw = "propertys")
    9: optional i64 natureType = 0 (api.raw = "natureType")
    10: optional list<staffItem> staffsId ={} (api.raw = "staffsId")

}

struct ContractSignMemberReq{
    1:optional string signImg ="" (api.raw = "signImg")
    2:optional list<i64> contractId = 0 (api.raw = "contractId")
}


struct OrderAllCountReq{
    1: optional i64 venueId=0 (api.raw = "venueId")
    2: optional string payWay="" (api.raw = "payWay")
    3: optional string startAt="" (api.raw = "startAt")
    4: optional string endAt="" (api.raw = "endAt")
    /**状态  1收入 2支出*/
    5:  optional i64 status=0 (api.raw = "status")
    255: optional i64 page=1 (api.raw = "page")
    256: optional i64 pageSize=100 (api.raw = "pageSize")
}
service OrderService {

    base.NilResponse Buy(1: BuyReq req) (api.post = "/service/order/buy")
    base.NilResponse Update(1: UpdateOrderReq req) (api.post = "/service/order/update")
    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/order/status")
    base.NilResponse ListOrder(1: ListOrderReq req )(api.post = "/service/order/list") // 订单列表
    base.NilResponse GetOrderById(1: base.IDReq req) (api.get = "/service/order/info") // 订单详情
    base.NilResponse OrderListExport(1: ListOrderReq req) (api.post = "/service/order/list/export")


    base.NilResponse OrderAllCount(1: OrderAllCountReq req) (api.post = "/service/order/all-count")

    base.NilResponse ContractSignMember(1: ContractSignMemberReq req) (api.post = "/service/order/contract-sign-member")



}
