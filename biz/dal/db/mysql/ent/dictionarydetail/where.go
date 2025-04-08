// Code generated by ent, DO NOT EDIT.

package dictionarydetail

import (
	"kcers/biz/dal/db/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldUpdatedAt, v))
}

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldDelete, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldCreatedID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldStatus, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldTitle, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldKey, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldValue, v))
}

// DictionaryID applies equality check predicate on the "dictionary_id" field. It's identical to DictionaryIDEQ.
func DictionaryID(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldDictionaryID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNEQ(FieldDelete, v))
}

// DeleteIn applies the In predicate on the "delete" field.
func DeleteIn(vs ...int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIn(FieldDelete, vs...))
}

// DeleteNotIn applies the NotIn predicate on the "delete" field.
func DeleteNotIn(vs ...int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotIn(FieldDelete, vs...))
}

// DeleteGT applies the GT predicate on the "delete" field.
func DeleteGT(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGT(FieldDelete, v))
}

// DeleteGTE applies the GTE predicate on the "delete" field.
func DeleteGTE(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGTE(FieldDelete, v))
}

// DeleteLT applies the LT predicate on the "delete" field.
func DeleteLT(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLT(FieldDelete, v))
}

// DeleteLTE applies the LTE predicate on the "delete" field.
func DeleteLTE(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLTE(FieldDelete, v))
}

// DeleteIsNil applies the IsNil predicate on the "delete" field.
func DeleteIsNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIsNull(FieldDelete))
}

// DeleteNotNil applies the NotNil predicate on the "delete" field.
func DeleteNotNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotNull(FieldDelete))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLTE(FieldCreatedID, v))
}

// CreatedIDIsNil applies the IsNil predicate on the "created_id" field.
func CreatedIDIsNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIsNull(FieldCreatedID))
}

// CreatedIDNotNil applies the NotNil predicate on the "created_id" field.
func CreatedIDNotNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotNull(FieldCreatedID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotNull(FieldStatus))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldContainsFold(FieldTitle, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldContainsFold(FieldKey, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldContainsFold(FieldValue, v))
}

// DictionaryIDEQ applies the EQ predicate on the "dictionary_id" field.
func DictionaryIDEQ(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldEQ(FieldDictionaryID, v))
}

// DictionaryIDNEQ applies the NEQ predicate on the "dictionary_id" field.
func DictionaryIDNEQ(v int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNEQ(FieldDictionaryID, v))
}

// DictionaryIDIn applies the In predicate on the "dictionary_id" field.
func DictionaryIDIn(vs ...int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIn(FieldDictionaryID, vs...))
}

// DictionaryIDNotIn applies the NotIn predicate on the "dictionary_id" field.
func DictionaryIDNotIn(vs ...int64) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotIn(FieldDictionaryID, vs...))
}

// DictionaryIDIsNil applies the IsNil predicate on the "dictionary_id" field.
func DictionaryIDIsNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldIsNull(FieldDictionaryID))
}

// DictionaryIDNotNil applies the NotNil predicate on the "dictionary_id" field.
func DictionaryIDNotNil() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.FieldNotNull(FieldDictionaryID))
}

// HasDictionary applies the HasEdge predicate on the "dictionary" edge.
func HasDictionary() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DictionaryTable, DictionaryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDictionaryWith applies the HasEdge predicate on the "dictionary" edge with a given conditions (other predicates).
func HasDictionaryWith(preds ...predicate.Dictionary) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(func(s *sql.Selector) {
		step := newDictionaryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProperty applies the HasEdge predicate on the "property" edge.
func HasProperty() predicate.DictionaryDetail {
	return predicate.DictionaryDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PropertyTable, PropertyPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPropertyWith applies the HasEdge predicate on the "property" edge with a given conditions (other predicates).
func HasPropertyWith(preds ...predicate.ProductProperty) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(func(s *sql.Selector) {
		step := newPropertyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DictionaryDetail) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DictionaryDetail) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DictionaryDetail) predicate.DictionaryDetail {
	return predicate.DictionaryDetail(sql.NotPredicates(p))
}
