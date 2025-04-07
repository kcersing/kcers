// Code generated by ent, DO NOT EDIT.

package membercontractcontent

import (
	"kcers/biz/dal/db/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldUpdatedAt, v))
}

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldDelete, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldCreatedID, v))
}

// MemberContractID applies equality check predicate on the "member_contract_id" field. It's identical to MemberContractIDEQ.
func MemberContractID(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldMemberContractID, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldContent, v))
}

// SignImg applies equality check predicate on the "sign_img" field. It's identical to SignImgEQ.
func SignImg(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldSignImg, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNEQ(FieldDelete, v))
}

// DeleteIn applies the In predicate on the "delete" field.
func DeleteIn(vs ...int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIn(FieldDelete, vs...))
}

// DeleteNotIn applies the NotIn predicate on the "delete" field.
func DeleteNotIn(vs ...int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotIn(FieldDelete, vs...))
}

// DeleteGT applies the GT predicate on the "delete" field.
func DeleteGT(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGT(FieldDelete, v))
}

// DeleteGTE applies the GTE predicate on the "delete" field.
func DeleteGTE(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGTE(FieldDelete, v))
}

// DeleteLT applies the LT predicate on the "delete" field.
func DeleteLT(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLT(FieldDelete, v))
}

// DeleteLTE applies the LTE predicate on the "delete" field.
func DeleteLTE(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLTE(FieldDelete, v))
}

// DeleteIsNil applies the IsNil predicate on the "delete" field.
func DeleteIsNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIsNull(FieldDelete))
}

// DeleteNotNil applies the NotNil predicate on the "delete" field.
func DeleteNotNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotNull(FieldDelete))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLTE(FieldCreatedID, v))
}

// CreatedIDIsNil applies the IsNil predicate on the "created_id" field.
func CreatedIDIsNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIsNull(FieldCreatedID))
}

// CreatedIDNotNil applies the NotNil predicate on the "created_id" field.
func CreatedIDNotNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotNull(FieldCreatedID))
}

// MemberContractIDEQ applies the EQ predicate on the "member_contract_id" field.
func MemberContractIDEQ(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldMemberContractID, v))
}

// MemberContractIDNEQ applies the NEQ predicate on the "member_contract_id" field.
func MemberContractIDNEQ(v int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNEQ(FieldMemberContractID, v))
}

// MemberContractIDIn applies the In predicate on the "member_contract_id" field.
func MemberContractIDIn(vs ...int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIn(FieldMemberContractID, vs...))
}

// MemberContractIDNotIn applies the NotIn predicate on the "member_contract_id" field.
func MemberContractIDNotIn(vs ...int64) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotIn(FieldMemberContractID, vs...))
}

// MemberContractIDIsNil applies the IsNil predicate on the "member_contract_id" field.
func MemberContractIDIsNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIsNull(FieldMemberContractID))
}

// MemberContractIDNotNil applies the NotNil predicate on the "member_contract_id" field.
func MemberContractIDNotNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotNull(FieldMemberContractID))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldHasSuffix(FieldContent, v))
}

// ContentIsNil applies the IsNil predicate on the "content" field.
func ContentIsNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIsNull(FieldContent))
}

// ContentNotNil applies the NotNil predicate on the "content" field.
func ContentNotNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotNull(FieldContent))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldContainsFold(FieldContent, v))
}

// SignImgEQ applies the EQ predicate on the "sign_img" field.
func SignImgEQ(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEQ(FieldSignImg, v))
}

// SignImgNEQ applies the NEQ predicate on the "sign_img" field.
func SignImgNEQ(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNEQ(FieldSignImg, v))
}

// SignImgIn applies the In predicate on the "sign_img" field.
func SignImgIn(vs ...string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIn(FieldSignImg, vs...))
}

// SignImgNotIn applies the NotIn predicate on the "sign_img" field.
func SignImgNotIn(vs ...string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotIn(FieldSignImg, vs...))
}

// SignImgGT applies the GT predicate on the "sign_img" field.
func SignImgGT(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGT(FieldSignImg, v))
}

// SignImgGTE applies the GTE predicate on the "sign_img" field.
func SignImgGTE(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldGTE(FieldSignImg, v))
}

// SignImgLT applies the LT predicate on the "sign_img" field.
func SignImgLT(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLT(FieldSignImg, v))
}

// SignImgLTE applies the LTE predicate on the "sign_img" field.
func SignImgLTE(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldLTE(FieldSignImg, v))
}

// SignImgContains applies the Contains predicate on the "sign_img" field.
func SignImgContains(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldContains(FieldSignImg, v))
}

// SignImgHasPrefix applies the HasPrefix predicate on the "sign_img" field.
func SignImgHasPrefix(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldHasPrefix(FieldSignImg, v))
}

// SignImgHasSuffix applies the HasSuffix predicate on the "sign_img" field.
func SignImgHasSuffix(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldHasSuffix(FieldSignImg, v))
}

// SignImgIsNil applies the IsNil predicate on the "sign_img" field.
func SignImgIsNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldIsNull(FieldSignImg))
}

// SignImgNotNil applies the NotNil predicate on the "sign_img" field.
func SignImgNotNil() predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldNotNull(FieldSignImg))
}

// SignImgEqualFold applies the EqualFold predicate on the "sign_img" field.
func SignImgEqualFold(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldEqualFold(FieldSignImg, v))
}

// SignImgContainsFold applies the ContainsFold predicate on the "sign_img" field.
func SignImgContainsFold(v string) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.FieldContainsFold(FieldSignImg, v))
}

// HasContract applies the HasEdge predicate on the "contract" edge.
func HasContract() predicate.MemberContractContent {
	return predicate.MemberContractContent(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ContractTable, ContractColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContractWith applies the HasEdge predicate on the "contract" edge with a given conditions (other predicates).
func HasContractWith(preds ...predicate.MemberContract) predicate.MemberContractContent {
	return predicate.MemberContractContent(func(s *sql.Selector) {
		step := newContractStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MemberContractContent) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MemberContractContent) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MemberContractContent) predicate.MemberContractContent {
	return predicate.MemberContractContent(sql.NotPredicates(p))
}
