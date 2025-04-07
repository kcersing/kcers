// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kcers/biz/dal/db/mysql/ent/member"
	"kcers/biz/dal/db/mysql/ent/membercontract"
	"kcers/biz/dal/db/mysql/ent/memberproduct"
	"kcers/biz/dal/db/mysql/ent/order"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MemberContract is the model entity for the MemberContract schema.
type MemberContract struct {
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
	// 会员id
	MemberID int64 `json:"member_id,omitempty"`
	// 原始合同id
	ContractID int64 `json:"contract_id,omitempty"`
	// 订单id
	OrderID int64 `json:"order_id,omitempty"`
	// 场馆id
	VenueID int64 `json:"venue_id,omitempty"`
	// 会员产品id
	MemberProductID int64 `json:"member_product_id,omitempty"`
	// name | 名称
	Name string `json:"name,omitempty"`
	// sign | 签字
	Sign string `json:"sign,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberContractQuery when eager-loading is set.
	Edges        MemberContractEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MemberContractEdges holds the relations/edges for other nodes in the graph.
type MemberContractEdges struct {
	// Content holds the value of the content edge.
	Content []*MemberContractContent `json:"content,omitempty"`
	// Member holds the value of the member edge.
	Member *Member `json:"member,omitempty"`
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// MemberProduct holds the value of the member_product edge.
	MemberProduct *MemberProduct `json:"member_product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ContentOrErr returns the Content value or an error if the edge
// was not loaded in eager-loading.
func (e MemberContractEdges) ContentOrErr() ([]*MemberContractContent, error) {
	if e.loadedTypes[0] {
		return e.Content, nil
	}
	return nil, &NotLoadedError{edge: "content"}
}

// MemberOrErr returns the Member value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberContractEdges) MemberOrErr() (*Member, error) {
	if e.Member != nil {
		return e.Member, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: member.Label}
	}
	return nil, &NotLoadedError{edge: "member"}
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberContractEdges) OrderOrErr() (*Order, error) {
	if e.Order != nil {
		return e.Order, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: order.Label}
	}
	return nil, &NotLoadedError{edge: "order"}
}

// MemberProductOrErr returns the MemberProduct value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberContractEdges) MemberProductOrErr() (*MemberProduct, error) {
	if e.MemberProduct != nil {
		return e.MemberProduct, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: memberproduct.Label}
	}
	return nil, &NotLoadedError{edge: "member_product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MemberContract) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case membercontract.FieldID, membercontract.FieldDelete, membercontract.FieldCreatedID, membercontract.FieldStatus, membercontract.FieldMemberID, membercontract.FieldContractID, membercontract.FieldOrderID, membercontract.FieldVenueID, membercontract.FieldMemberProductID:
			values[i] = new(sql.NullInt64)
		case membercontract.FieldName, membercontract.FieldSign:
			values[i] = new(sql.NullString)
		case membercontract.FieldCreatedAt, membercontract.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MemberContract fields.
func (mc *MemberContract) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case membercontract.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mc.ID = int64(value.Int64)
		case membercontract.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mc.CreatedAt = value.Time
			}
		case membercontract.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mc.UpdatedAt = value.Time
			}
		case membercontract.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				mc.Delete = value.Int64
			}
		case membercontract.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				mc.CreatedID = value.Int64
			}
		case membercontract.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				mc.Status = value.Int64
			}
		case membercontract.FieldMemberID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				mc.MemberID = value.Int64
			}
		case membercontract.FieldContractID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field contract_id", values[i])
			} else if value.Valid {
				mc.ContractID = value.Int64
			}
		case membercontract.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				mc.OrderID = value.Int64
			}
		case membercontract.FieldVenueID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field venue_id", values[i])
			} else if value.Valid {
				mc.VenueID = value.Int64
			}
		case membercontract.FieldMemberProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_product_id", values[i])
			} else if value.Valid {
				mc.MemberProductID = value.Int64
			}
		case membercontract.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mc.Name = value.String
			}
		case membercontract.FieldSign:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sign", values[i])
			} else if value.Valid {
				mc.Sign = value.String
			}
		default:
			mc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MemberContract.
// This includes values selected through modifiers, order, etc.
func (mc *MemberContract) Value(name string) (ent.Value, error) {
	return mc.selectValues.Get(name)
}

// QueryContent queries the "content" edge of the MemberContract entity.
func (mc *MemberContract) QueryContent() *MemberContractContentQuery {
	return NewMemberContractClient(mc.config).QueryContent(mc)
}

// QueryMember queries the "member" edge of the MemberContract entity.
func (mc *MemberContract) QueryMember() *MemberQuery {
	return NewMemberContractClient(mc.config).QueryMember(mc)
}

// QueryOrder queries the "order" edge of the MemberContract entity.
func (mc *MemberContract) QueryOrder() *OrderQuery {
	return NewMemberContractClient(mc.config).QueryOrder(mc)
}

// QueryMemberProduct queries the "member_product" edge of the MemberContract entity.
func (mc *MemberContract) QueryMemberProduct() *MemberProductQuery {
	return NewMemberContractClient(mc.config).QueryMemberProduct(mc)
}

// Update returns a builder for updating this MemberContract.
// Note that you need to call MemberContract.Unwrap() before calling this method if this MemberContract
// was returned from a transaction, and the transaction was committed or rolled back.
func (mc *MemberContract) Update() *MemberContractUpdateOne {
	return NewMemberContractClient(mc.config).UpdateOne(mc)
}

// Unwrap unwraps the MemberContract entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mc *MemberContract) Unwrap() *MemberContract {
	_tx, ok := mc.config.driver.(*txDriver)
	if !ok {
		panic("ent: MemberContract is not a transactional entity")
	}
	mc.config.driver = _tx.drv
	return mc
}

// String implements the fmt.Stringer.
func (mc *MemberContract) String() string {
	var builder strings.Builder
	builder.WriteString("MemberContract(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(mc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", mc.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", mc.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", mc.Status))
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(fmt.Sprintf("%v", mc.MemberID))
	builder.WriteString(", ")
	builder.WriteString("contract_id=")
	builder.WriteString(fmt.Sprintf("%v", mc.ContractID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", mc.OrderID))
	builder.WriteString(", ")
	builder.WriteString("venue_id=")
	builder.WriteString(fmt.Sprintf("%v", mc.VenueID))
	builder.WriteString(", ")
	builder.WriteString("member_product_id=")
	builder.WriteString(fmt.Sprintf("%v", mc.MemberProductID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(mc.Name)
	builder.WriteString(", ")
	builder.WriteString("sign=")
	builder.WriteString(mc.Sign)
	builder.WriteByte(')')
	return builder.String()
}

// MemberContracts is a parsable slice of MemberContract.
type MemberContracts []*MemberContract
