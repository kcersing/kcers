// Code generated by ent, DO NOT EDIT.

package productproperty

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the productproperty type in the database.
	Label = "product_property"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDelete holds the string denoting the delete field in the database.
	FieldDelete = "delete"
	// FieldCreatedID holds the string denoting the created_id field in the database.
	FieldCreatedID = "created_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
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
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// FieldCreateID holds the string denoting the create_id field in the database.
	FieldCreateID = "create_id"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeContracts holds the string denoting the contracts edge name in mutations.
	EdgeContracts = "contracts"
	// EdgeVenues holds the string denoting the venues edge name in mutations.
	EdgeVenues = "venues"
	// Table holds the table name of the productproperty in the database.
	Table = "product_property"
	// ProductTable is the table that holds the product relation/edge. The primary key declared below.
	ProductTable = "product_propertys"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "product"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "product_property_tags"
	// TagsInverseTable is the table name for the DictionaryDetail entity.
	// It exists in this package in order to avoid circular dependency with the "dictionarydetail" package.
	TagsInverseTable = "sys_dictionary_details"
	// ContractsTable is the table that holds the contracts relation/edge. The primary key declared below.
	ContractsTable = "product_property_contracts"
	// ContractsInverseTable is the table name for the Contract entity.
	// It exists in this package in order to avoid circular dependency with the "contract" package.
	ContractsInverseTable = "contracts"
	// VenuesTable is the table that holds the venues relation/edge. The primary key declared below.
	VenuesTable = "product_property_venues"
	// VenuesInverseTable is the table name for the Venue entity.
	// It exists in this package in order to avoid circular dependency with the "venue" package.
	VenuesInverseTable = "venue"
)

// Columns holds all SQL columns for productproperty fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDelete,
	FieldCreatedID,
	FieldStatus,
	FieldType,
	FieldName,
	FieldDuration,
	FieldLength,
	FieldCount,
	FieldPrice,
	FieldData,
	FieldCreateID,
}

var (
	// ProductPrimaryKey and ProductColumn2 are the table columns denoting the
	// primary key for the product relation (M2M).
	ProductPrimaryKey = []string{"product_id", "product_property_id"}
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"product_property_id", "dictionary_detail_id"}
	// ContractsPrimaryKey and ContractsColumn2 are the table columns denoting the
	// primary key for the contracts relation (M2M).
	ContractsPrimaryKey = []string{"product_property_id", "contract_id"}
	// VenuesPrimaryKey and VenuesColumn2 are the table columns denoting the
	// primary key for the venues relation (M2M).
	VenuesPrimaryKey = []string{"product_property_id", "venue_id"}
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
	// DefaultDelete holds the default value on creation for the "delete" field.
	DefaultDelete int64
	// DefaultCreatedID holds the default value on creation for the "created_id" field.
	DefaultCreatedID int64
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int64
)

// OrderOption defines the ordering options for the ProductProperty queries.
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

// ByDelete orders the results by the delete field.
func ByDelete(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelete, opts...).ToFunc()
}

// ByCreatedID orders the results by the created_id field.
func ByCreatedID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
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

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByData orders the results by the data field.
func ByData(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldData, opts...).ToFunc()
}

// ByCreateID orders the results by the create_id field.
func ByCreateID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateID, opts...).ToFunc()
}

// ByProductCount orders the results by product count.
func ByProductCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductStep(), opts...)
	}
}

// ByProduct orders the results by product terms.
func ByProduct(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByContractsCount orders the results by contracts count.
func ByContractsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newContractsStep(), opts...)
	}
}

// ByContracts orders the results by contracts terms.
func ByContracts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContractsStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProductTable, ProductPrimaryKey...),
	)
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
	)
}
func newContractsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContractsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ContractsTable, ContractsPrimaryKey...),
	)
}
func newVenuesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VenuesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, VenuesTable, VenuesPrimaryKey...),
	)
}
