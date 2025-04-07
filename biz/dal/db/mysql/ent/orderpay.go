// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kcers/biz/dal/db/mysql/ent/order"
	"kcers/biz/dal/db/mysql/ent/orderpay"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// OrderPay is the model entity for the OrderPay schema.
type OrderPay struct {
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
	// 订单id
	OrderID int64 `json:"order_id,omitempty"`
	// 减免
	Remission float64 `json:"remission,omitempty"`
	// 实际付款
	Pay float64 `json:"pay,omitempty"`
	// 备注
	Note string `json:"note,omitempty"`
	// 支付方式
	PayWay string `json:"pay_way,omitempty"`
	// 操作人id
	CreateID int64 `json:"create_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderPayQuery when eager-loading is set.
	Edges        OrderPayEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrderPayEdges holds the relations/edges for other nodes in the graph.
type OrderPayEdges struct {
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderPayEdges) OrderOrErr() (*Order, error) {
	if e.Order != nil {
		return e.Order, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: order.Label}
	}
	return nil, &NotLoadedError{edge: "order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderPay) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderpay.FieldRemission, orderpay.FieldPay:
			values[i] = new(sql.NullFloat64)
		case orderpay.FieldID, orderpay.FieldDelete, orderpay.FieldCreatedID, orderpay.FieldOrderID, orderpay.FieldCreateID:
			values[i] = new(sql.NullInt64)
		case orderpay.FieldNote, orderpay.FieldPayWay:
			values[i] = new(sql.NullString)
		case orderpay.FieldCreatedAt, orderpay.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderPay fields.
func (op *OrderPay) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderpay.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			op.ID = int64(value.Int64)
		case orderpay.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				op.CreatedAt = value.Time
			}
		case orderpay.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				op.UpdatedAt = value.Time
			}
		case orderpay.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				op.Delete = value.Int64
			}
		case orderpay.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				op.CreatedID = value.Int64
			}
		case orderpay.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				op.OrderID = value.Int64
			}
		case orderpay.FieldRemission:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field remission", values[i])
			} else if value.Valid {
				op.Remission = value.Float64
			}
		case orderpay.FieldPay:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field pay", values[i])
			} else if value.Valid {
				op.Pay = value.Float64
			}
		case orderpay.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				op.Note = value.String
			}
		case orderpay.FieldPayWay:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pay_way", values[i])
			} else if value.Valid {
				op.PayWay = value.String
			}
		case orderpay.FieldCreateID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_id", values[i])
			} else if value.Valid {
				op.CreateID = value.Int64
			}
		default:
			op.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderPay.
// This includes values selected through modifiers, order, etc.
func (op *OrderPay) Value(name string) (ent.Value, error) {
	return op.selectValues.Get(name)
}

// QueryOrder queries the "order" edge of the OrderPay entity.
func (op *OrderPay) QueryOrder() *OrderQuery {
	return NewOrderPayClient(op.config).QueryOrder(op)
}

// Update returns a builder for updating this OrderPay.
// Note that you need to call OrderPay.Unwrap() before calling this method if this OrderPay
// was returned from a transaction, and the transaction was committed or rolled back.
func (op *OrderPay) Update() *OrderPayUpdateOne {
	return NewOrderPayClient(op.config).UpdateOne(op)
}

// Unwrap unwraps the OrderPay entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (op *OrderPay) Unwrap() *OrderPay {
	_tx, ok := op.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderPay is not a transactional entity")
	}
	op.config.driver = _tx.drv
	return op
}

// String implements the fmt.Stringer.
func (op *OrderPay) String() string {
	var builder strings.Builder
	builder.WriteString("OrderPay(")
	builder.WriteString(fmt.Sprintf("id=%v, ", op.ID))
	builder.WriteString("created_at=")
	builder.WriteString(op.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(op.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", op.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", op.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", op.OrderID))
	builder.WriteString(", ")
	builder.WriteString("remission=")
	builder.WriteString(fmt.Sprintf("%v", op.Remission))
	builder.WriteString(", ")
	builder.WriteString("pay=")
	builder.WriteString(fmt.Sprintf("%v", op.Pay))
	builder.WriteString(", ")
	builder.WriteString("note=")
	builder.WriteString(op.Note)
	builder.WriteString(", ")
	builder.WriteString("pay_way=")
	builder.WriteString(op.PayWay)
	builder.WriteString(", ")
	builder.WriteString("create_id=")
	builder.WriteString(fmt.Sprintf("%v", op.CreateID))
	builder.WriteByte(')')
	return builder.String()
}

// OrderPays is a parsable slice of OrderPay.
type OrderPays []*OrderPay
