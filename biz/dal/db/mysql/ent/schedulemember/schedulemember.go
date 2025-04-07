// Code generated by ent, DO NOT EDIT.

package schedulemember

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the schedulemember type in the database.
	Label = "schedule_member"
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
	// FieldVenueID holds the string denoting the venue_id field in the database.
	FieldVenueID = "venue_id"
	// FieldScheduleID holds the string denoting the schedule_id field in the database.
	FieldScheduleID = "schedule_id"
	// FieldScheduleName holds the string denoting the schedule_name field in the database.
	FieldScheduleName = "schedule_name"
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// FieldMemberProductID holds the string denoting the member_product_id field in the database.
	FieldMemberProductID = "member_product_id"
	// FieldMemberProductPropertyID holds the string denoting the member_product_property_id field in the database.
	FieldMemberProductPropertyID = "member_product_property_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldSignStartAt holds the string denoting the sign_start_at field in the database.
	FieldSignStartAt = "sign_start_at"
	// FieldSignEndAt holds the string denoting the sign_end_at field in the database.
	FieldSignEndAt = "sign_end_at"
	// FieldMemberName holds the string denoting the member_name field in the database.
	FieldMemberName = "member_name"
	// FieldMemberProductName holds the string denoting the member_product_name field in the database.
	FieldMemberProductName = "member_product_name"
	// FieldMemberProductPropertyName holds the string denoting the member_product_property_name field in the database.
	FieldMemberProductPropertyName = "member_product_property_name"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// EdgeSchedule holds the string denoting the schedule edge name in mutations.
	EdgeSchedule = "schedule"
	// Table holds the table name of the schedulemember in the database.
	Table = "schedule_member"
	// ScheduleTable is the table that holds the schedule relation/edge.
	ScheduleTable = "schedule_member"
	// ScheduleInverseTable is the table name for the Schedule entity.
	// It exists in this package in order to avoid circular dependency with the "schedule" package.
	ScheduleInverseTable = "schedule"
	// ScheduleColumn is the table column denoting the schedule relation/edge.
	ScheduleColumn = "schedule_id"
)

// Columns holds all SQL columns for schedulemember fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDelete,
	FieldCreatedID,
	FieldStatus,
	FieldVenueID,
	FieldScheduleID,
	FieldScheduleName,
	FieldMemberID,
	FieldMemberProductID,
	FieldMemberProductPropertyID,
	FieldType,
	FieldStartAt,
	FieldEndAt,
	FieldSignStartAt,
	FieldSignEndAt,
	FieldMemberName,
	FieldMemberProductName,
	FieldMemberProductPropertyName,
	FieldRemark,
}

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
	// DefaultStartAt holds the default value on creation for the "start_at" field.
	DefaultStartAt func() time.Time
	// DefaultEndAt holds the default value on creation for the "end_at" field.
	DefaultEndAt func() time.Time
	// DefaultSignStartAt holds the default value on creation for the "sign_start_at" field.
	DefaultSignStartAt func() time.Time
	// DefaultSignEndAt holds the default value on creation for the "sign_end_at" field.
	DefaultSignEndAt func() time.Time
)

// OrderOption defines the ordering options for the ScheduleMember queries.
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

// ByVenueID orders the results by the venue_id field.
func ByVenueID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueID, opts...).ToFunc()
}

// ByScheduleID orders the results by the schedule_id field.
func ByScheduleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScheduleID, opts...).ToFunc()
}

// ByScheduleName orders the results by the schedule_name field.
func ByScheduleName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScheduleName, opts...).ToFunc()
}

// ByMemberID orders the results by the member_id field.
func ByMemberID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberID, opts...).ToFunc()
}

// ByMemberProductID orders the results by the member_product_id field.
func ByMemberProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductID, opts...).ToFunc()
}

// ByMemberProductPropertyID orders the results by the member_product_property_id field.
func ByMemberProductPropertyID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductPropertyID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByStartAt orders the results by the start_at field.
func ByStartAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartAt, opts...).ToFunc()
}

// ByEndAt orders the results by the end_at field.
func ByEndAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndAt, opts...).ToFunc()
}

// BySignStartAt orders the results by the sign_start_at field.
func BySignStartAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignStartAt, opts...).ToFunc()
}

// BySignEndAt orders the results by the sign_end_at field.
func BySignEndAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignEndAt, opts...).ToFunc()
}

// ByMemberName orders the results by the member_name field.
func ByMemberName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberName, opts...).ToFunc()
}

// ByMemberProductName orders the results by the member_product_name field.
func ByMemberProductName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductName, opts...).ToFunc()
}

// ByMemberProductPropertyName orders the results by the member_product_property_name field.
func ByMemberProductPropertyName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductPropertyName, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByScheduleField orders the results by schedule field.
func ByScheduleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newScheduleStep(), sql.OrderByField(field, opts...))
	}
}
func newScheduleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScheduleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ScheduleTable, ScheduleColumn),
	)
}
