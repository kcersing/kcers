// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kcers/biz/dal/db/mysql/ent/contract"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Contract is the model entity for the Contract schema.
type Contract struct {
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
	// name | 名称
	Name string `json:"name,omitempty"`
	// content | 内容
	Content string `json:"content,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContractQuery when eager-loading is set.
	Edges        ContractEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ContractEdges holds the relations/edges for other nodes in the graph.
type ContractEdges struct {
	// Property holds the value of the property edge.
	Property []*ProductProperty `json:"property,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PropertyOrErr returns the Property value or an error if the edge
// was not loaded in eager-loading.
func (e ContractEdges) PropertyOrErr() ([]*ProductProperty, error) {
	if e.loadedTypes[0] {
		return e.Property, nil
	}
	return nil, &NotLoadedError{edge: "property"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Contract) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contract.FieldID, contract.FieldDelete, contract.FieldCreatedID, contract.FieldStatus:
			values[i] = new(sql.NullInt64)
		case contract.FieldName, contract.FieldContent:
			values[i] = new(sql.NullString)
		case contract.FieldCreatedAt, contract.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Contract fields.
func (c *Contract) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contract.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case contract.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case contract.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case contract.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				c.Delete = value.Int64
			}
		case contract.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				c.CreatedID = value.Int64
			}
		case contract.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = value.Int64
			}
		case contract.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case contract.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				c.Content = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Contract.
// This includes values selected through modifiers, order, etc.
func (c *Contract) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryProperty queries the "property" edge of the Contract entity.
func (c *Contract) QueryProperty() *ProductPropertyQuery {
	return NewContractClient(c.config).QueryProperty(c)
}

// Update returns a builder for updating this Contract.
// Note that you need to call Contract.Unwrap() before calling this method if this Contract
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Contract) Update() *ContractUpdateOne {
	return NewContractClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Contract entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Contract) Unwrap() *Contract {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Contract is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Contract) String() string {
	var builder strings.Builder
	builder.WriteString("Contract(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", c.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(c.Content)
	builder.WriteByte(')')
	return builder.String()
}

// Contracts is a parsable slice of Contract.
type Contracts []*Contract
