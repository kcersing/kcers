// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kcers/biz/dal/db/mysql/ent/dictionary"
	"kcers/biz/dal/db/mysql/ent/dictionarydetail"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// DictionaryDetail is the model entity for the DictionaryDetail schema.
type DictionaryDetail struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// last delete  1:已删除
	Delete int64 `json:"delete,omitempty"`
	// created
	CreatedID int64 `json:"created_id,omitempty"`
	// 状态[0:禁用;1:正常]
	Status int64 `json:"status,omitempty"`
	// the title shown in the ui | 展示名称 （建议配合i18n）
	Title string `json:"title,omitempty"`
	// key | 键
	Key string `json:"key,omitempty"`
	// value | 值
	Value string `json:"value,omitempty"`
	// Dictionary ID | 字典ID
	DictionaryID int64 `json:"dictionary_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DictionaryDetailQuery when eager-loading is set.
	Edges        DictionaryDetailEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DictionaryDetailEdges holds the relations/edges for other nodes in the graph.
type DictionaryDetailEdges struct {
	// Dictionary holds the value of the dictionary edge.
	Dictionary *Dictionary `json:"dictionary,omitempty"`
	// Property holds the value of the property edge.
	Property []*ProductProperty `json:"property,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DictionaryOrErr returns the Dictionary value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DictionaryDetailEdges) DictionaryOrErr() (*Dictionary, error) {
	if e.Dictionary != nil {
		return e.Dictionary, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: dictionary.Label}
	}
	return nil, &NotLoadedError{edge: "dictionary"}
}

// PropertyOrErr returns the Property value or an error if the edge
// was not loaded in eager-loading.
func (e DictionaryDetailEdges) PropertyOrErr() ([]*ProductProperty, error) {
	if e.loadedTypes[1] {
		return e.Property, nil
	}
	return nil, &NotLoadedError{edge: "property"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e DictionaryDetailEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DictionaryDetail) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case dictionarydetail.FieldID, dictionarydetail.FieldDelete, dictionarydetail.FieldCreatedID, dictionarydetail.FieldStatus, dictionarydetail.FieldDictionaryID:
			values[i] = new(sql.NullInt64)
		case dictionarydetail.FieldTitle, dictionarydetail.FieldKey, dictionarydetail.FieldValue:
			values[i] = new(sql.NullString)
		case dictionarydetail.FieldCreatedAt, dictionarydetail.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DictionaryDetail fields.
func (dd *DictionaryDetail) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dictionarydetail.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dd.ID = int64(value.Int64)
		case dictionarydetail.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dd.CreatedAt = value.Time
			}
		case dictionarydetail.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dd.UpdatedAt = value.Time
			}
		case dictionarydetail.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				dd.Delete = value.Int64
			}
		case dictionarydetail.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				dd.CreatedID = value.Int64
			}
		case dictionarydetail.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				dd.Status = value.Int64
			}
		case dictionarydetail.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				dd.Title = value.String
			}
		case dictionarydetail.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				dd.Key = value.String
			}
		case dictionarydetail.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				dd.Value = value.String
			}
		case dictionarydetail.FieldDictionaryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dictionary_id", values[i])
			} else if value.Valid {
				dd.DictionaryID = value.Int64
			}
		default:
			dd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the DictionaryDetail.
// This includes values selected through modifiers, order, etc.
func (dd *DictionaryDetail) GetValue(name string) (ent.Value, error) {
	return dd.selectValues.Get(name)
}

// QueryDictionary queries the "dictionary" edge of the DictionaryDetail entity.
func (dd *DictionaryDetail) QueryDictionary() *DictionaryQuery {
	return NewDictionaryDetailClient(dd.config).QueryDictionary(dd)
}

// QueryProperty queries the "property" edge of the DictionaryDetail entity.
func (dd *DictionaryDetail) QueryProperty() *ProductPropertyQuery {
	return NewDictionaryDetailClient(dd.config).QueryProperty(dd)
}

// QueryUsers queries the "users" edge of the DictionaryDetail entity.
func (dd *DictionaryDetail) QueryUsers() *UserQuery {
	return NewDictionaryDetailClient(dd.config).QueryUsers(dd)
}

// Update returns a builder for updating this DictionaryDetail.
// Note that you need to call DictionaryDetail.Unwrap() before calling this method if this DictionaryDetail
// was returned from a transaction, and the transaction was committed or rolled back.
func (dd *DictionaryDetail) Update() *DictionaryDetailUpdateOne {
	return NewDictionaryDetailClient(dd.config).UpdateOne(dd)
}

// Unwrap unwraps the DictionaryDetail entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dd *DictionaryDetail) Unwrap() *DictionaryDetail {
	_tx, ok := dd.config.driver.(*txDriver)
	if !ok {
		panic("ent: DictionaryDetail is not a transactional entity")
	}
	dd.config.driver = _tx.drv
	return dd
}

// String implements the fmt.Stringer.
func (dd *DictionaryDetail) String() string {
	var builder strings.Builder
	builder.WriteString("DictionaryDetail(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dd.ID))
	builder.WriteString("created_at=")
	builder.WriteString(dd.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(dd.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", dd.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", dd.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", dd.Status))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(dd.Title)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(dd.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(dd.Value)
	builder.WriteString(", ")
	builder.WriteString("dictionary_id=")
	builder.WriteString(fmt.Sprintf("%v", dd.DictionaryID))
	builder.WriteByte(')')
	return builder.String()
}

// DictionaryDetails is a parsable slice of DictionaryDetail.
type DictionaryDetails []*DictionaryDetail
