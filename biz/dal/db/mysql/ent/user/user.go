// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
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
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldSideMode holds the string denoting the side_mode field in the database.
	FieldSideMode = "side_mode"
	// FieldBaseColor holds the string denoting the base_color field in the database.
	FieldBaseColor = "base_color"
	// FieldActiveColor holds the string denoting the active_color field in the database.
	FieldActiveColor = "active_color"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldWecom holds the string denoting the wecom field in the database.
	FieldWecom = "wecom"
	// FieldJob holds the string denoting the job field in the database.
	FieldJob = "job"
	// FieldOrganization holds the string denoting the organization field in the database.
	FieldOrganization = "organization"
	// FieldDefaultVenueID holds the string denoting the default_venue_id field in the database.
	FieldDefaultVenueID = "default_venue_id"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"
	// EdgeCreatedOrders holds the string denoting the created_orders edge name in mutations.
	EdgeCreatedOrders = "created_orders"
	// EdgeUserEntry holds the string denoting the user_entry edge name in mutations.
	EdgeUserEntry = "user_entry"
	// EdgeUserFace holds the string denoting the user_face edge name in mutations.
	EdgeUserFace = "user_face"
	// Table holds the table name of the user in the database.
	Table = "sys_users"
	// TokenTable is the table that holds the token relation/edge.
	TokenTable = "sys_tokens"
	// TokenInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokenInverseTable = "sys_tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "user_token"
	// CreatedOrdersTable is the table that holds the created_orders relation/edge.
	CreatedOrdersTable = "order"
	// CreatedOrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	CreatedOrdersInverseTable = "order"
	// CreatedOrdersColumn is the table column denoting the created_orders relation/edge.
	CreatedOrdersColumn = "create_id"
	// UserEntryTable is the table that holds the user_entry relation/edge.
	UserEntryTable = "entry_logs"
	// UserEntryInverseTable is the table name for the EntryLogs entity.
	// It exists in this package in order to avoid circular dependency with the "entrylogs" package.
	UserEntryInverseTable = "entry_logs"
	// UserEntryColumn is the table column denoting the user_entry relation/edge.
	UserEntryColumn = "user_id"
	// UserFaceTable is the table that holds the user_face relation/edge.
	UserFaceTable = "faces"
	// UserFaceInverseTable is the table name for the Face entity.
	// It exists in this package in order to avoid circular dependency with the "face" package.
	UserFaceInverseTable = "faces"
	// UserFaceColumn is the table column denoting the user_face relation/edge.
	UserFaceColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDelete,
	FieldCreatedID,
	FieldStatus,
	FieldUsername,
	FieldPassword,
	FieldNickname,
	FieldSideMode,
	FieldBaseColor,
	FieldActiveColor,
	FieldRoleID,
	FieldMobile,
	FieldEmail,
	FieldWecom,
	FieldJob,
	FieldOrganization,
	FieldDefaultVenueID,
	FieldAvatar,
	FieldGender,
	FieldBirthday,
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
	// DefaultSideMode holds the default value on creation for the "side_mode" field.
	DefaultSideMode string
	// DefaultBaseColor holds the default value on creation for the "base_color" field.
	DefaultBaseColor string
	// DefaultActiveColor holds the default value on creation for the "active_color" field.
	DefaultActiveColor string
	// DefaultRoleID holds the default value on creation for the "role_id" field.
	DefaultRoleID int64
	// DefaultGender holds the default value on creation for the "gender" field.
	DefaultGender int64
)

// OrderOption defines the ordering options for the User queries.
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

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// BySideMode orders the results by the side_mode field.
func BySideMode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSideMode, opts...).ToFunc()
}

// ByBaseColor orders the results by the base_color field.
func ByBaseColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBaseColor, opts...).ToFunc()
}

// ByActiveColor orders the results by the active_color field.
func ByActiveColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActiveColor, opts...).ToFunc()
}

// ByRoleID orders the results by the role_id field.
func ByRoleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleID, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByWecom orders the results by the wecom field.
func ByWecom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWecom, opts...).ToFunc()
}

// ByJob orders the results by the job field.
func ByJob(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJob, opts...).ToFunc()
}

// ByOrganization orders the results by the organization field.
func ByOrganization(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganization, opts...).ToFunc()
}

// ByDefaultVenueID orders the results by the default_venue_id field.
func ByDefaultVenueID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultVenueID, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByBirthday orders the results by the birthday field.
func ByBirthday(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthday, opts...).ToFunc()
}

// ByTokenField orders the results by token field.
func ByTokenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokenStep(), sql.OrderByField(field, opts...))
	}
}

// ByCreatedOrdersCount orders the results by created_orders count.
func ByCreatedOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedOrdersStep(), opts...)
	}
}

// ByCreatedOrders orders the results by created_orders terms.
func ByCreatedOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserEntryCount orders the results by user_entry count.
func ByUserEntryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserEntryStep(), opts...)
	}
}

// ByUserEntry orders the results by user_entry terms.
func ByUserEntry(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserEntryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserFaceCount orders the results by user_face count.
func ByUserFaceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserFaceStep(), opts...)
	}
}

// ByUserFace orders the results by user_face terms.
func ByUserFace(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserFaceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, TokenTable, TokenColumn),
	)
}
func newCreatedOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreatedOrdersTable, CreatedOrdersColumn),
	)
}
func newUserEntryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserEntryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserEntryTable, UserEntryColumn),
	)
}
func newUserFaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserFaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserFaceTable, UserFaceColumn),
	)
}
