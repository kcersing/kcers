// Code generated by ent, DO NOT EDIT.

package orderpay

import (
	"kcers/biz/dal/db/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldUpdatedAt, v))
}

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldDelete, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldCreatedID, v))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldOrderID, v))
}

// Remission applies equality check predicate on the "remission" field. It's identical to RemissionEQ.
func Remission(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldRemission, v))
}

// Pay applies equality check predicate on the "pay" field. It's identical to PayEQ.
func Pay(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldPay, v))
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldNote, v))
}

// PayWay applies equality check predicate on the "pay_way" field. It's identical to PayWayEQ.
func PayWay(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldPayWay, v))
}

// CreateID applies equality check predicate on the "create_id" field. It's identical to CreateIDEQ.
func CreateID(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldCreateID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNEQ(FieldDelete, v))
}

// DeleteIn applies the In predicate on the "delete" field.
func DeleteIn(vs ...int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIn(FieldDelete, vs...))
}

// DeleteNotIn applies the NotIn predicate on the "delete" field.
func DeleteNotIn(vs ...int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotIn(FieldDelete, vs...))
}

// DeleteGT applies the GT predicate on the "delete" field.
func DeleteGT(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGT(FieldDelete, v))
}

// DeleteGTE applies the GTE predicate on the "delete" field.
func DeleteGTE(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGTE(FieldDelete, v))
}

// DeleteLT applies the LT predicate on the "delete" field.
func DeleteLT(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLT(FieldDelete, v))
}

// DeleteLTE applies the LTE predicate on the "delete" field.
func DeleteLTE(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLTE(FieldDelete, v))
}

// DeleteIsNil applies the IsNil predicate on the "delete" field.
func DeleteIsNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIsNull(FieldDelete))
}

// DeleteNotNil applies the NotNil predicate on the "delete" field.
func DeleteNotNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotNull(FieldDelete))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLTE(FieldCreatedID, v))
}

// CreatedIDIsNil applies the IsNil predicate on the "created_id" field.
func CreatedIDIsNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIsNull(FieldCreatedID))
}

// CreatedIDNotNil applies the NotNil predicate on the "created_id" field.
func CreatedIDNotNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotNull(FieldCreatedID))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotIn(FieldOrderID, vs...))
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIsNull(FieldOrderID))
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotNull(FieldOrderID))
}

// RemissionEQ applies the EQ predicate on the "remission" field.
func RemissionEQ(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldRemission, v))
}

// RemissionNEQ applies the NEQ predicate on the "remission" field.
func RemissionNEQ(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNEQ(FieldRemission, v))
}

// RemissionIn applies the In predicate on the "remission" field.
func RemissionIn(vs ...float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIn(FieldRemission, vs...))
}

// RemissionNotIn applies the NotIn predicate on the "remission" field.
func RemissionNotIn(vs ...float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotIn(FieldRemission, vs...))
}

// RemissionGT applies the GT predicate on the "remission" field.
func RemissionGT(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGT(FieldRemission, v))
}

// RemissionGTE applies the GTE predicate on the "remission" field.
func RemissionGTE(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGTE(FieldRemission, v))
}

// RemissionLT applies the LT predicate on the "remission" field.
func RemissionLT(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLT(FieldRemission, v))
}

// RemissionLTE applies the LTE predicate on the "remission" field.
func RemissionLTE(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLTE(FieldRemission, v))
}

// RemissionIsNil applies the IsNil predicate on the "remission" field.
func RemissionIsNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIsNull(FieldRemission))
}

// RemissionNotNil applies the NotNil predicate on the "remission" field.
func RemissionNotNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotNull(FieldRemission))
}

// PayEQ applies the EQ predicate on the "pay" field.
func PayEQ(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldPay, v))
}

// PayNEQ applies the NEQ predicate on the "pay" field.
func PayNEQ(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNEQ(FieldPay, v))
}

// PayIn applies the In predicate on the "pay" field.
func PayIn(vs ...float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIn(FieldPay, vs...))
}

// PayNotIn applies the NotIn predicate on the "pay" field.
func PayNotIn(vs ...float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotIn(FieldPay, vs...))
}

// PayGT applies the GT predicate on the "pay" field.
func PayGT(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGT(FieldPay, v))
}

// PayGTE applies the GTE predicate on the "pay" field.
func PayGTE(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGTE(FieldPay, v))
}

// PayLT applies the LT predicate on the "pay" field.
func PayLT(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLT(FieldPay, v))
}

// PayLTE applies the LTE predicate on the "pay" field.
func PayLTE(v float64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLTE(FieldPay, v))
}

// PayIsNil applies the IsNil predicate on the "pay" field.
func PayIsNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIsNull(FieldPay))
}

// PayNotNil applies the NotNil predicate on the "pay" field.
func PayNotNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotNull(FieldPay))
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldHasSuffix(FieldNote, v))
}

// NoteIsNil applies the IsNil predicate on the "note" field.
func NoteIsNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIsNull(FieldNote))
}

// NoteNotNil applies the NotNil predicate on the "note" field.
func NoteNotNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotNull(FieldNote))
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldContainsFold(FieldNote, v))
}

// PayWayEQ applies the EQ predicate on the "pay_way" field.
func PayWayEQ(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldPayWay, v))
}

// PayWayNEQ applies the NEQ predicate on the "pay_way" field.
func PayWayNEQ(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNEQ(FieldPayWay, v))
}

// PayWayIn applies the In predicate on the "pay_way" field.
func PayWayIn(vs ...string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIn(FieldPayWay, vs...))
}

// PayWayNotIn applies the NotIn predicate on the "pay_way" field.
func PayWayNotIn(vs ...string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotIn(FieldPayWay, vs...))
}

// PayWayGT applies the GT predicate on the "pay_way" field.
func PayWayGT(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGT(FieldPayWay, v))
}

// PayWayGTE applies the GTE predicate on the "pay_way" field.
func PayWayGTE(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGTE(FieldPayWay, v))
}

// PayWayLT applies the LT predicate on the "pay_way" field.
func PayWayLT(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLT(FieldPayWay, v))
}

// PayWayLTE applies the LTE predicate on the "pay_way" field.
func PayWayLTE(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLTE(FieldPayWay, v))
}

// PayWayContains applies the Contains predicate on the "pay_way" field.
func PayWayContains(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldContains(FieldPayWay, v))
}

// PayWayHasPrefix applies the HasPrefix predicate on the "pay_way" field.
func PayWayHasPrefix(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldHasPrefix(FieldPayWay, v))
}

// PayWayHasSuffix applies the HasSuffix predicate on the "pay_way" field.
func PayWayHasSuffix(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldHasSuffix(FieldPayWay, v))
}

// PayWayIsNil applies the IsNil predicate on the "pay_way" field.
func PayWayIsNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIsNull(FieldPayWay))
}

// PayWayNotNil applies the NotNil predicate on the "pay_way" field.
func PayWayNotNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotNull(FieldPayWay))
}

// PayWayEqualFold applies the EqualFold predicate on the "pay_way" field.
func PayWayEqualFold(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEqualFold(FieldPayWay, v))
}

// PayWayContainsFold applies the ContainsFold predicate on the "pay_way" field.
func PayWayContainsFold(v string) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldContainsFold(FieldPayWay, v))
}

// CreateIDEQ applies the EQ predicate on the "create_id" field.
func CreateIDEQ(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldEQ(FieldCreateID, v))
}

// CreateIDNEQ applies the NEQ predicate on the "create_id" field.
func CreateIDNEQ(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNEQ(FieldCreateID, v))
}

// CreateIDIn applies the In predicate on the "create_id" field.
func CreateIDIn(vs ...int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIn(FieldCreateID, vs...))
}

// CreateIDNotIn applies the NotIn predicate on the "create_id" field.
func CreateIDNotIn(vs ...int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotIn(FieldCreateID, vs...))
}

// CreateIDGT applies the GT predicate on the "create_id" field.
func CreateIDGT(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGT(FieldCreateID, v))
}

// CreateIDGTE applies the GTE predicate on the "create_id" field.
func CreateIDGTE(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldGTE(FieldCreateID, v))
}

// CreateIDLT applies the LT predicate on the "create_id" field.
func CreateIDLT(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLT(FieldCreateID, v))
}

// CreateIDLTE applies the LTE predicate on the "create_id" field.
func CreateIDLTE(v int64) predicate.OrderPay {
	return predicate.OrderPay(sql.FieldLTE(FieldCreateID, v))
}

// CreateIDIsNil applies the IsNil predicate on the "create_id" field.
func CreateIDIsNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldIsNull(FieldCreateID))
}

// CreateIDNotNil applies the NotNil predicate on the "create_id" field.
func CreateIDNotNil() predicate.OrderPay {
	return predicate.OrderPay(sql.FieldNotNull(FieldCreateID))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.OrderPay {
	return predicate.OrderPay(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.OrderPay {
	return predicate.OrderPay(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderPay) predicate.OrderPay {
	return predicate.OrderPay(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderPay) predicate.OrderPay {
	return predicate.OrderPay(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderPay) predicate.OrderPay {
	return predicate.OrderPay(sql.NotPredicates(p))
}
