// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kcers/biz/dal/db/mysql/ent/member"
	"kcers/biz/dal/db/mysql/ent/memberprofile"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MemberProfile is the model entity for the MemberProfile schema.
type MemberProfile struct {
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
	// 意向
	Intention int64 `json:"intention,omitempty"`
	// 来源
	Source int64 `json:"source,omitempty"`
	// 会员id
	MemberID int64 `json:"member_id,omitempty"`
	// 性别 | [1:女性;2:男性;3:保密]
	Gender int64 `json:"gender,omitempty"`
	// 出生日期
	Birthday time.Time `json:"birthday,omitempty"`
	// email | 邮箱号
	Email string `json:"email,omitempty"`
	// wecom | 微信号
	Wecom string `json:"wecom,omitempty"`
	// 关联会员
	RelationMid int64 `json:"relation_mid,omitempty"`
	// 关联会员
	RelationMame string `json:"relation_mame,omitempty"`
	// 跟进人员工
	RelationUID int64 `json:"relation_uid,omitempty"`
	// 跟进人员工
	RelationUname string `json:"relation_uname,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberProfileQuery when eager-loading is set.
	Edges        MemberProfileEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MemberProfileEdges holds the relations/edges for other nodes in the graph.
type MemberProfileEdges struct {
	// Member holds the value of the member edge.
	Member *Member `json:"member,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MemberOrErr returns the Member value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberProfileEdges) MemberOrErr() (*Member, error) {
	if e.Member != nil {
		return e.Member, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: member.Label}
	}
	return nil, &NotLoadedError{edge: "member"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MemberProfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case memberprofile.FieldID, memberprofile.FieldDelete, memberprofile.FieldCreatedID, memberprofile.FieldIntention, memberprofile.FieldSource, memberprofile.FieldMemberID, memberprofile.FieldGender, memberprofile.FieldRelationMid, memberprofile.FieldRelationUID:
			values[i] = new(sql.NullInt64)
		case memberprofile.FieldEmail, memberprofile.FieldWecom, memberprofile.FieldRelationMame, memberprofile.FieldRelationUname:
			values[i] = new(sql.NullString)
		case memberprofile.FieldCreatedAt, memberprofile.FieldUpdatedAt, memberprofile.FieldBirthday:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MemberProfile fields.
func (mp *MemberProfile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case memberprofile.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mp.ID = int64(value.Int64)
		case memberprofile.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mp.CreatedAt = value.Time
			}
		case memberprofile.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mp.UpdatedAt = value.Time
			}
		case memberprofile.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				mp.Delete = value.Int64
			}
		case memberprofile.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				mp.CreatedID = value.Int64
			}
		case memberprofile.FieldIntention:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field intention", values[i])
			} else if value.Valid {
				mp.Intention = value.Int64
			}
		case memberprofile.FieldSource:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				mp.Source = value.Int64
			}
		case memberprofile.FieldMemberID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				mp.MemberID = value.Int64
			}
		case memberprofile.FieldGender:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				mp.Gender = value.Int64
			}
		case memberprofile.FieldBirthday:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birthday", values[i])
			} else if value.Valid {
				mp.Birthday = value.Time
			}
		case memberprofile.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				mp.Email = value.String
			}
		case memberprofile.FieldWecom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wecom", values[i])
			} else if value.Valid {
				mp.Wecom = value.String
			}
		case memberprofile.FieldRelationMid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relation_mid", values[i])
			} else if value.Valid {
				mp.RelationMid = value.Int64
			}
		case memberprofile.FieldRelationMame:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field relation_mame", values[i])
			} else if value.Valid {
				mp.RelationMame = value.String
			}
		case memberprofile.FieldRelationUID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relation_uid", values[i])
			} else if value.Valid {
				mp.RelationUID = value.Int64
			}
		case memberprofile.FieldRelationUname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field relation_uname", values[i])
			} else if value.Valid {
				mp.RelationUname = value.String
			}
		default:
			mp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MemberProfile.
// This includes values selected through modifiers, order, etc.
func (mp *MemberProfile) Value(name string) (ent.Value, error) {
	return mp.selectValues.Get(name)
}

// QueryMember queries the "member" edge of the MemberProfile entity.
func (mp *MemberProfile) QueryMember() *MemberQuery {
	return NewMemberProfileClient(mp.config).QueryMember(mp)
}

// Update returns a builder for updating this MemberProfile.
// Note that you need to call MemberProfile.Unwrap() before calling this method if this MemberProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (mp *MemberProfile) Update() *MemberProfileUpdateOne {
	return NewMemberProfileClient(mp.config).UpdateOne(mp)
}

// Unwrap unwraps the MemberProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mp *MemberProfile) Unwrap() *MemberProfile {
	_tx, ok := mp.config.driver.(*txDriver)
	if !ok {
		panic("ent: MemberProfile is not a transactional entity")
	}
	mp.config.driver = _tx.drv
	return mp
}

// String implements the fmt.Stringer.
func (mp *MemberProfile) String() string {
	var builder strings.Builder
	builder.WriteString("MemberProfile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(mp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", mp.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", mp.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("intention=")
	builder.WriteString(fmt.Sprintf("%v", mp.Intention))
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(fmt.Sprintf("%v", mp.Source))
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(fmt.Sprintf("%v", mp.MemberID))
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", mp.Gender))
	builder.WriteString(", ")
	builder.WriteString("birthday=")
	builder.WriteString(mp.Birthday.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(mp.Email)
	builder.WriteString(", ")
	builder.WriteString("wecom=")
	builder.WriteString(mp.Wecom)
	builder.WriteString(", ")
	builder.WriteString("relation_mid=")
	builder.WriteString(fmt.Sprintf("%v", mp.RelationMid))
	builder.WriteString(", ")
	builder.WriteString("relation_mame=")
	builder.WriteString(mp.RelationMame)
	builder.WriteString(", ")
	builder.WriteString("relation_uid=")
	builder.WriteString(fmt.Sprintf("%v", mp.RelationUID))
	builder.WriteString(", ")
	builder.WriteString("relation_uname=")
	builder.WriteString(mp.RelationUname)
	builder.WriteByte(')')
	return builder.String()
}

// MemberProfiles is a parsable slice of MemberProfile.
type MemberProfiles []*MemberProfile
