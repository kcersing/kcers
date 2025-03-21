// Code generated by ent, DO NOT EDIT.

package memberproductproperty

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the memberproductproperty type in the database.
	Label = "member_product_property"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// FieldMemberProductID holds the string denoting the member_product_id field in the database.
	FieldMemberProductID = "member_product_id"
	// FieldSn holds the string denoting the sn field in the database.
	FieldSn = "sn"
	// FieldPropertyID holds the string denoting the property_id field in the database.
	FieldPropertyID = "property_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldCountSurplus holds the string denoting the count_surplus field in the database.
	FieldCountSurplus = "count_surplus"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldValidityAt holds the string denoting the validity_at field in the database.
	FieldValidityAt = "validity_at"
	// FieldCancelAt holds the string denoting the cancel_at field in the database.
	FieldCancelAt = "cancel_at"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeVenues holds the string denoting the venues edge name in mutations.
	EdgeVenues = "venues"
	// Table holds the table name of the memberproductproperty in the database.
	Table = "member_product_property"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "member_product_property"
	// OwnerInverseTable is the table name for the MemberProduct entity.
	// It exists in this package in order to avoid circular dependency with the "memberproduct" package.
	OwnerInverseTable = "member_product"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "member_product_id"
	// VenuesTable is the table that holds the venues relation/edge. The primary key declared below.
	VenuesTable = "member_product_property_venues"
	// VenuesInverseTable is the table name for the Venue entity.
	// It exists in this package in order to avoid circular dependency with the "venue" package.
	VenuesInverseTable = "venue"
)

// Columns holds all SQL columns for memberproductproperty fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldMemberID,
	FieldMemberProductID,
	FieldSn,
	FieldPropertyID,
	FieldType,
	FieldName,
	FieldDuration,
	FieldLength,
	FieldCount,
	FieldCountSurplus,
	FieldPrice,
	FieldValidityAt,
	FieldCancelAt,
}

var (
	// VenuesPrimaryKey and VenuesColumn2 are the table columns denoting the
	// primary key for the venues relation (M2M).
	VenuesPrimaryKey = []string{"member_product_property_id", "venue_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int64
	// DefaultCount holds the default value on creation for the "count" field.
	DefaultCount int64
	// DefaultCountSurplus holds the default value on creation for the "count_surplus" field.
	DefaultCountSurplus int64
)

// OrderOption defines the ordering options for the MemberProductProperty queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByMemberID orders the results by the member_id field.
func ByMemberID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberID, opts...).ToFunc()
}

// ByMemberProductID orders the results by the member_product_id field.
func ByMemberProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductID, opts...).ToFunc()
}

// BySn orders the results by the sn field.
func BySn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSn, opts...).ToFunc()
}

// ByPropertyID orders the results by the property_id field.
func ByPropertyID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPropertyID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByLength orders the results by the length field.
func ByLength(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLength, opts...).ToFunc()
}

// ByCount orders the results by the count field.
func ByCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCount, opts...).ToFunc()
}

// ByCountSurplus orders the results by the count_surplus field.
func ByCountSurplus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountSurplus, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByValidityAt orders the results by the validity_at field.
func ByValidityAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValidityAt, opts...).ToFunc()
}

// ByCancelAt orders the results by the cancel_at field.
func ByCancelAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCancelAt, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByVenuesCount orders the results by venues count.
func ByVenuesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVenuesStep(), opts...)
	}
}

// ByVenues orders the results by venues terms.
func ByVenues(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVenuesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newVenuesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VenuesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, VenuesTable, VenuesPrimaryKey...),
	)
}
