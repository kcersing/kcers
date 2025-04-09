// Code generated by ent, DO NOT EDIT.

package logs

import (
	"kcers/biz/dal/db/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldUpdatedAt, v))
}

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldDelete, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldCreatedID, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldType, v))
}

// Method applies equality check predicate on the "method" field. It's identical to MethodEQ.
func Method(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldMethod, v))
}

// API applies equality check predicate on the "api" field. It's identical to APIEQ.
func API(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldAPI, v))
}

// Success applies equality check predicate on the "success" field. It's identical to SuccessEQ.
func Success(v bool) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldSuccess, v))
}

// ReqContent applies equality check predicate on the "req_content" field. It's identical to ReqContentEQ.
func ReqContent(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldReqContent, v))
}

// RespContent applies equality check predicate on the "resp_content" field. It's identical to RespContentEQ.
func RespContent(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldRespContent, v))
}

// IP applies equality check predicate on the "ip" field. It's identical to IPEQ.
func IP(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldIP, v))
}

// UserAgent applies equality check predicate on the "user_agent" field. It's identical to UserAgentEQ.
func UserAgent(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldUserAgent, v))
}

// Operatorsr applies equality check predicate on the "operatorsr" field. It's identical to OperatorsrEQ.
func Operatorsr(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldOperatorsr, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldTime, v))
}

// Identity applies equality check predicate on the "identity" field. It's identical to IdentityEQ.
func Identity(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldIdentity, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Logs {
	return predicate.Logs(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Logs {
	return predicate.Logs(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Logs {
	return predicate.Logs(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Logs {
	return predicate.Logs(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldDelete, v))
}

// DeleteIn applies the In predicate on the "delete" field.
func DeleteIn(vs ...int64) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldDelete, vs...))
}

// DeleteNotIn applies the NotIn predicate on the "delete" field.
func DeleteNotIn(vs ...int64) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldDelete, vs...))
}

// DeleteGT applies the GT predicate on the "delete" field.
func DeleteGT(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldDelete, v))
}

// DeleteGTE applies the GTE predicate on the "delete" field.
func DeleteGTE(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldDelete, v))
}

// DeleteLT applies the LT predicate on the "delete" field.
func DeleteLT(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldDelete, v))
}

// DeleteLTE applies the LTE predicate on the "delete" field.
func DeleteLTE(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldDelete, v))
}

// DeleteIsNil applies the IsNil predicate on the "delete" field.
func DeleteIsNil() predicate.Logs {
	return predicate.Logs(sql.FieldIsNull(FieldDelete))
}

// DeleteNotNil applies the NotNil predicate on the "delete" field.
func DeleteNotNil() predicate.Logs {
	return predicate.Logs(sql.FieldNotNull(FieldDelete))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldCreatedID, v))
}

// CreatedIDIsNil applies the IsNil predicate on the "created_id" field.
func CreatedIDIsNil() predicate.Logs {
	return predicate.Logs(sql.FieldIsNull(FieldCreatedID))
}

// CreatedIDNotNil applies the NotNil predicate on the "created_id" field.
func CreatedIDNotNil() predicate.Logs {
	return predicate.Logs(sql.FieldNotNull(FieldCreatedID))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContainsFold(FieldType, v))
}

// MethodEQ applies the EQ predicate on the "method" field.
func MethodEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldMethod, v))
}

// MethodNEQ applies the NEQ predicate on the "method" field.
func MethodNEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldMethod, v))
}

// MethodIn applies the In predicate on the "method" field.
func MethodIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldMethod, vs...))
}

// MethodNotIn applies the NotIn predicate on the "method" field.
func MethodNotIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldMethod, vs...))
}

// MethodGT applies the GT predicate on the "method" field.
func MethodGT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldMethod, v))
}

// MethodGTE applies the GTE predicate on the "method" field.
func MethodGTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldMethod, v))
}

// MethodLT applies the LT predicate on the "method" field.
func MethodLT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldMethod, v))
}

// MethodLTE applies the LTE predicate on the "method" field.
func MethodLTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldMethod, v))
}

// MethodContains applies the Contains predicate on the "method" field.
func MethodContains(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContains(FieldMethod, v))
}

// MethodHasPrefix applies the HasPrefix predicate on the "method" field.
func MethodHasPrefix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasPrefix(FieldMethod, v))
}

// MethodHasSuffix applies the HasSuffix predicate on the "method" field.
func MethodHasSuffix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasSuffix(FieldMethod, v))
}

// MethodEqualFold applies the EqualFold predicate on the "method" field.
func MethodEqualFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEqualFold(FieldMethod, v))
}

// MethodContainsFold applies the ContainsFold predicate on the "method" field.
func MethodContainsFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContainsFold(FieldMethod, v))
}

// APIEQ applies the EQ predicate on the "api" field.
func APIEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldAPI, v))
}

// APINEQ applies the NEQ predicate on the "api" field.
func APINEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldAPI, v))
}

// APIIn applies the In predicate on the "api" field.
func APIIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldAPI, vs...))
}

// APINotIn applies the NotIn predicate on the "api" field.
func APINotIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldAPI, vs...))
}

// APIGT applies the GT predicate on the "api" field.
func APIGT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldAPI, v))
}

// APIGTE applies the GTE predicate on the "api" field.
func APIGTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldAPI, v))
}

// APILT applies the LT predicate on the "api" field.
func APILT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldAPI, v))
}

// APILTE applies the LTE predicate on the "api" field.
func APILTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldAPI, v))
}

// APIContains applies the Contains predicate on the "api" field.
func APIContains(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContains(FieldAPI, v))
}

// APIHasPrefix applies the HasPrefix predicate on the "api" field.
func APIHasPrefix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasPrefix(FieldAPI, v))
}

// APIHasSuffix applies the HasSuffix predicate on the "api" field.
func APIHasSuffix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasSuffix(FieldAPI, v))
}

// APIEqualFold applies the EqualFold predicate on the "api" field.
func APIEqualFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEqualFold(FieldAPI, v))
}

// APIContainsFold applies the ContainsFold predicate on the "api" field.
func APIContainsFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContainsFold(FieldAPI, v))
}

// SuccessEQ applies the EQ predicate on the "success" field.
func SuccessEQ(v bool) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldSuccess, v))
}

// SuccessNEQ applies the NEQ predicate on the "success" field.
func SuccessNEQ(v bool) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldSuccess, v))
}

// ReqContentEQ applies the EQ predicate on the "req_content" field.
func ReqContentEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldReqContent, v))
}

// ReqContentNEQ applies the NEQ predicate on the "req_content" field.
func ReqContentNEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldReqContent, v))
}

// ReqContentIn applies the In predicate on the "req_content" field.
func ReqContentIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldReqContent, vs...))
}

// ReqContentNotIn applies the NotIn predicate on the "req_content" field.
func ReqContentNotIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldReqContent, vs...))
}

// ReqContentGT applies the GT predicate on the "req_content" field.
func ReqContentGT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldReqContent, v))
}

// ReqContentGTE applies the GTE predicate on the "req_content" field.
func ReqContentGTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldReqContent, v))
}

// ReqContentLT applies the LT predicate on the "req_content" field.
func ReqContentLT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldReqContent, v))
}

// ReqContentLTE applies the LTE predicate on the "req_content" field.
func ReqContentLTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldReqContent, v))
}

// ReqContentContains applies the Contains predicate on the "req_content" field.
func ReqContentContains(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContains(FieldReqContent, v))
}

// ReqContentHasPrefix applies the HasPrefix predicate on the "req_content" field.
func ReqContentHasPrefix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasPrefix(FieldReqContent, v))
}

// ReqContentHasSuffix applies the HasSuffix predicate on the "req_content" field.
func ReqContentHasSuffix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasSuffix(FieldReqContent, v))
}

// ReqContentIsNil applies the IsNil predicate on the "req_content" field.
func ReqContentIsNil() predicate.Logs {
	return predicate.Logs(sql.FieldIsNull(FieldReqContent))
}

// ReqContentNotNil applies the NotNil predicate on the "req_content" field.
func ReqContentNotNil() predicate.Logs {
	return predicate.Logs(sql.FieldNotNull(FieldReqContent))
}

// ReqContentEqualFold applies the EqualFold predicate on the "req_content" field.
func ReqContentEqualFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEqualFold(FieldReqContent, v))
}

// ReqContentContainsFold applies the ContainsFold predicate on the "req_content" field.
func ReqContentContainsFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContainsFold(FieldReqContent, v))
}

// RespContentEQ applies the EQ predicate on the "resp_content" field.
func RespContentEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldRespContent, v))
}

// RespContentNEQ applies the NEQ predicate on the "resp_content" field.
func RespContentNEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldRespContent, v))
}

// RespContentIn applies the In predicate on the "resp_content" field.
func RespContentIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldRespContent, vs...))
}

// RespContentNotIn applies the NotIn predicate on the "resp_content" field.
func RespContentNotIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldRespContent, vs...))
}

// RespContentGT applies the GT predicate on the "resp_content" field.
func RespContentGT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldRespContent, v))
}

// RespContentGTE applies the GTE predicate on the "resp_content" field.
func RespContentGTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldRespContent, v))
}

// RespContentLT applies the LT predicate on the "resp_content" field.
func RespContentLT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldRespContent, v))
}

// RespContentLTE applies the LTE predicate on the "resp_content" field.
func RespContentLTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldRespContent, v))
}

// RespContentContains applies the Contains predicate on the "resp_content" field.
func RespContentContains(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContains(FieldRespContent, v))
}

// RespContentHasPrefix applies the HasPrefix predicate on the "resp_content" field.
func RespContentHasPrefix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasPrefix(FieldRespContent, v))
}

// RespContentHasSuffix applies the HasSuffix predicate on the "resp_content" field.
func RespContentHasSuffix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasSuffix(FieldRespContent, v))
}

// RespContentIsNil applies the IsNil predicate on the "resp_content" field.
func RespContentIsNil() predicate.Logs {
	return predicate.Logs(sql.FieldIsNull(FieldRespContent))
}

// RespContentNotNil applies the NotNil predicate on the "resp_content" field.
func RespContentNotNil() predicate.Logs {
	return predicate.Logs(sql.FieldNotNull(FieldRespContent))
}

// RespContentEqualFold applies the EqualFold predicate on the "resp_content" field.
func RespContentEqualFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEqualFold(FieldRespContent, v))
}

// RespContentContainsFold applies the ContainsFold predicate on the "resp_content" field.
func RespContentContainsFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContainsFold(FieldRespContent, v))
}

// IPEQ applies the EQ predicate on the "ip" field.
func IPEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldIP, v))
}

// IPNEQ applies the NEQ predicate on the "ip" field.
func IPNEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldIP, v))
}

// IPIn applies the In predicate on the "ip" field.
func IPIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldIP, vs...))
}

// IPNotIn applies the NotIn predicate on the "ip" field.
func IPNotIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldIP, vs...))
}

// IPGT applies the GT predicate on the "ip" field.
func IPGT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldIP, v))
}

// IPGTE applies the GTE predicate on the "ip" field.
func IPGTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldIP, v))
}

// IPLT applies the LT predicate on the "ip" field.
func IPLT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldIP, v))
}

// IPLTE applies the LTE predicate on the "ip" field.
func IPLTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldIP, v))
}

// IPContains applies the Contains predicate on the "ip" field.
func IPContains(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContains(FieldIP, v))
}

// IPHasPrefix applies the HasPrefix predicate on the "ip" field.
func IPHasPrefix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasPrefix(FieldIP, v))
}

// IPHasSuffix applies the HasSuffix predicate on the "ip" field.
func IPHasSuffix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasSuffix(FieldIP, v))
}

// IPIsNil applies the IsNil predicate on the "ip" field.
func IPIsNil() predicate.Logs {
	return predicate.Logs(sql.FieldIsNull(FieldIP))
}

// IPNotNil applies the NotNil predicate on the "ip" field.
func IPNotNil() predicate.Logs {
	return predicate.Logs(sql.FieldNotNull(FieldIP))
}

// IPEqualFold applies the EqualFold predicate on the "ip" field.
func IPEqualFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEqualFold(FieldIP, v))
}

// IPContainsFold applies the ContainsFold predicate on the "ip" field.
func IPContainsFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContainsFold(FieldIP, v))
}

// UserAgentEQ applies the EQ predicate on the "user_agent" field.
func UserAgentEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldUserAgent, v))
}

// UserAgentNEQ applies the NEQ predicate on the "user_agent" field.
func UserAgentNEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldUserAgent, v))
}

// UserAgentIn applies the In predicate on the "user_agent" field.
func UserAgentIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldUserAgent, vs...))
}

// UserAgentNotIn applies the NotIn predicate on the "user_agent" field.
func UserAgentNotIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldUserAgent, vs...))
}

// UserAgentGT applies the GT predicate on the "user_agent" field.
func UserAgentGT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldUserAgent, v))
}

// UserAgentGTE applies the GTE predicate on the "user_agent" field.
func UserAgentGTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldUserAgent, v))
}

// UserAgentLT applies the LT predicate on the "user_agent" field.
func UserAgentLT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldUserAgent, v))
}

// UserAgentLTE applies the LTE predicate on the "user_agent" field.
func UserAgentLTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldUserAgent, v))
}

// UserAgentContains applies the Contains predicate on the "user_agent" field.
func UserAgentContains(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContains(FieldUserAgent, v))
}

// UserAgentHasPrefix applies the HasPrefix predicate on the "user_agent" field.
func UserAgentHasPrefix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasPrefix(FieldUserAgent, v))
}

// UserAgentHasSuffix applies the HasSuffix predicate on the "user_agent" field.
func UserAgentHasSuffix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasSuffix(FieldUserAgent, v))
}

// UserAgentIsNil applies the IsNil predicate on the "user_agent" field.
func UserAgentIsNil() predicate.Logs {
	return predicate.Logs(sql.FieldIsNull(FieldUserAgent))
}

// UserAgentNotNil applies the NotNil predicate on the "user_agent" field.
func UserAgentNotNil() predicate.Logs {
	return predicate.Logs(sql.FieldNotNull(FieldUserAgent))
}

// UserAgentEqualFold applies the EqualFold predicate on the "user_agent" field.
func UserAgentEqualFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEqualFold(FieldUserAgent, v))
}

// UserAgentContainsFold applies the ContainsFold predicate on the "user_agent" field.
func UserAgentContainsFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContainsFold(FieldUserAgent, v))
}

// OperatorsrEQ applies the EQ predicate on the "operatorsr" field.
func OperatorsrEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldOperatorsr, v))
}

// OperatorsrNEQ applies the NEQ predicate on the "operatorsr" field.
func OperatorsrNEQ(v string) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldOperatorsr, v))
}

// OperatorsrIn applies the In predicate on the "operatorsr" field.
func OperatorsrIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldOperatorsr, vs...))
}

// OperatorsrNotIn applies the NotIn predicate on the "operatorsr" field.
func OperatorsrNotIn(vs ...string) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldOperatorsr, vs...))
}

// OperatorsrGT applies the GT predicate on the "operatorsr" field.
func OperatorsrGT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldOperatorsr, v))
}

// OperatorsrGTE applies the GTE predicate on the "operatorsr" field.
func OperatorsrGTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldOperatorsr, v))
}

// OperatorsrLT applies the LT predicate on the "operatorsr" field.
func OperatorsrLT(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldOperatorsr, v))
}

// OperatorsrLTE applies the LTE predicate on the "operatorsr" field.
func OperatorsrLTE(v string) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldOperatorsr, v))
}

// OperatorsrContains applies the Contains predicate on the "operatorsr" field.
func OperatorsrContains(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContains(FieldOperatorsr, v))
}

// OperatorsrHasPrefix applies the HasPrefix predicate on the "operatorsr" field.
func OperatorsrHasPrefix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasPrefix(FieldOperatorsr, v))
}

// OperatorsrHasSuffix applies the HasSuffix predicate on the "operatorsr" field.
func OperatorsrHasSuffix(v string) predicate.Logs {
	return predicate.Logs(sql.FieldHasSuffix(FieldOperatorsr, v))
}

// OperatorsrIsNil applies the IsNil predicate on the "operatorsr" field.
func OperatorsrIsNil() predicate.Logs {
	return predicate.Logs(sql.FieldIsNull(FieldOperatorsr))
}

// OperatorsrNotNil applies the NotNil predicate on the "operatorsr" field.
func OperatorsrNotNil() predicate.Logs {
	return predicate.Logs(sql.FieldNotNull(FieldOperatorsr))
}

// OperatorsrEqualFold applies the EqualFold predicate on the "operatorsr" field.
func OperatorsrEqualFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldEqualFold(FieldOperatorsr, v))
}

// OperatorsrContainsFold applies the ContainsFold predicate on the "operatorsr" field.
func OperatorsrContainsFold(v string) predicate.Logs {
	return predicate.Logs(sql.FieldContainsFold(FieldOperatorsr, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...int64) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...int64) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldTime, v))
}

// TimeIsNil applies the IsNil predicate on the "time" field.
func TimeIsNil() predicate.Logs {
	return predicate.Logs(sql.FieldIsNull(FieldTime))
}

// TimeNotNil applies the NotNil predicate on the "time" field.
func TimeNotNil() predicate.Logs {
	return predicate.Logs(sql.FieldNotNull(FieldTime))
}

// IdentityEQ applies the EQ predicate on the "identity" field.
func IdentityEQ(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldEQ(FieldIdentity, v))
}

// IdentityNEQ applies the NEQ predicate on the "identity" field.
func IdentityNEQ(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldNEQ(FieldIdentity, v))
}

// IdentityIn applies the In predicate on the "identity" field.
func IdentityIn(vs ...int64) predicate.Logs {
	return predicate.Logs(sql.FieldIn(FieldIdentity, vs...))
}

// IdentityNotIn applies the NotIn predicate on the "identity" field.
func IdentityNotIn(vs ...int64) predicate.Logs {
	return predicate.Logs(sql.FieldNotIn(FieldIdentity, vs...))
}

// IdentityGT applies the GT predicate on the "identity" field.
func IdentityGT(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldGT(FieldIdentity, v))
}

// IdentityGTE applies the GTE predicate on the "identity" field.
func IdentityGTE(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldGTE(FieldIdentity, v))
}

// IdentityLT applies the LT predicate on the "identity" field.
func IdentityLT(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldLT(FieldIdentity, v))
}

// IdentityLTE applies the LTE predicate on the "identity" field.
func IdentityLTE(v int64) predicate.Logs {
	return predicate.Logs(sql.FieldLTE(FieldIdentity, v))
}

// IdentityIsNil applies the IsNil predicate on the "identity" field.
func IdentityIsNil() predicate.Logs {
	return predicate.Logs(sql.FieldIsNull(FieldIdentity))
}

// IdentityNotNil applies the NotNil predicate on the "identity" field.
func IdentityNotNil() predicate.Logs {
	return predicate.Logs(sql.FieldNotNull(FieldIdentity))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Logs) predicate.Logs {
	return predicate.Logs(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Logs) predicate.Logs {
	return predicate.Logs(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Logs) predicate.Logs {
	return predicate.Logs(sql.NotPredicates(p))
}
