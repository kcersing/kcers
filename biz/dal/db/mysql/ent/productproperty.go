// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kcers/biz/dal/db/mysql/ent/productproperty"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ProductProperty is the model entity for the ProductProperty schema.
type ProductProperty struct {
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
	// 类型
	Type string `json:"type,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 总时长
	Duration int64 `json:"duration,omitempty"`
	// 单次时长
	Length int64 `json:"length,omitempty"`
	// 次数
	Count int64 `json:"count,omitempty"`
	// 定价
	Price float64 `json:"price,omitempty"`
	// Data holds the value of the "data" field.
	Data string `json:"data,omitempty"`
	// 创建人id
	CreateID int64 `json:"create_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductPropertyQuery when eager-loading is set.
	Edges        ProductPropertyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProductPropertyEdges holds the relations/edges for other nodes in the graph.
type ProductPropertyEdges struct {
	// Product holds the value of the product edge.
	Product []*Product `json:"product,omitempty"`
	// Venues holds the value of the venues edge.
	Venues []*Venue `json:"venues,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e ProductPropertyEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// VenuesOrErr returns the Venues value or an error if the edge
// was not loaded in eager-loading.
func (e ProductPropertyEdges) VenuesOrErr() ([]*Venue, error) {
	if e.loadedTypes[1] {
		return e.Venues, nil
	}
	return nil, &NotLoadedError{edge: "venues"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductProperty) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case productproperty.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case productproperty.FieldID, productproperty.FieldDelete, productproperty.FieldCreatedID, productproperty.FieldStatus, productproperty.FieldDuration, productproperty.FieldLength, productproperty.FieldCount, productproperty.FieldCreateID:
			values[i] = new(sql.NullInt64)
		case productproperty.FieldType, productproperty.FieldName, productproperty.FieldData:
			values[i] = new(sql.NullString)
		case productproperty.FieldCreatedAt, productproperty.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductProperty fields.
func (pp *ProductProperty) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productproperty.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pp.ID = int64(value.Int64)
		case productproperty.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pp.CreatedAt = value.Time
			}
		case productproperty.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pp.UpdatedAt = value.Time
			}
		case productproperty.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				pp.Delete = value.Int64
			}
		case productproperty.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				pp.CreatedID = value.Int64
			}
		case productproperty.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pp.Status = value.Int64
			}
		case productproperty.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pp.Type = value.String
			}
		case productproperty.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pp.Name = value.String
			}
		case productproperty.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				pp.Duration = value.Int64
			}
		case productproperty.FieldLength:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field length", values[i])
			} else if value.Valid {
				pp.Length = value.Int64
			}
		case productproperty.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				pp.Count = value.Int64
			}
		case productproperty.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pp.Price = value.Float64
			}
		case productproperty.FieldData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value.Valid {
				pp.Data = value.String
			}
		case productproperty.FieldCreateID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_id", values[i])
			} else if value.Valid {
				pp.CreateID = value.Int64
			}
		default:
			pp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductProperty.
// This includes values selected through modifiers, order, etc.
func (pp *ProductProperty) Value(name string) (ent.Value, error) {
	return pp.selectValues.Get(name)
}

// QueryProduct queries the "product" edge of the ProductProperty entity.
func (pp *ProductProperty) QueryProduct() *ProductQuery {
	return NewProductPropertyClient(pp.config).QueryProduct(pp)
}

// QueryVenues queries the "venues" edge of the ProductProperty entity.
func (pp *ProductProperty) QueryVenues() *VenueQuery {
	return NewProductPropertyClient(pp.config).QueryVenues(pp)
}

// Update returns a builder for updating this ProductProperty.
// Note that you need to call ProductProperty.Unwrap() before calling this method if this ProductProperty
// was returned from a transaction, and the transaction was committed or rolled back.
func (pp *ProductProperty) Update() *ProductPropertyUpdateOne {
	return NewProductPropertyClient(pp.config).UpdateOne(pp)
}

// Unwrap unwraps the ProductProperty entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pp *ProductProperty) Unwrap() *ProductProperty {
	_tx, ok := pp.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductProperty is not a transactional entity")
	}
	pp.config.driver = _tx.drv
	return pp
}

// String implements the fmt.Stringer.
func (pp *ProductProperty) String() string {
	var builder strings.Builder
	builder.WriteString("ProductProperty(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", pp.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", pp.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pp.Status))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(pp.Type)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pp.Name)
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", pp.Duration))
	builder.WriteString(", ")
	builder.WriteString("length=")
	builder.WriteString(fmt.Sprintf("%v", pp.Length))
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", pp.Count))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", pp.Price))
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(pp.Data)
	builder.WriteString(", ")
	builder.WriteString("create_id=")
	builder.WriteString(fmt.Sprintf("%v", pp.CreateID))
	builder.WriteByte(')')
	return builder.String()
}

// ProductProperties is a parsable slice of ProductProperty.
type ProductProperties []*ProductProperty
