package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"kcers/biz/dal/db/mysql/ent/schema/mixins"
)

type VenueMember struct {
	ent.Schema
}

func (VenueMember) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("member_id").Comment("会员id").Optional(),
		field.Int64("venue_id").Comment("场馆id").Optional(),

		field.Float("money_sum").Default(3).Comment("消费总金额").Optional(),
		field.Int64("product_id").Default(0).Comment("首次的产品").Optional(),
		field.String("product_name").Comment("首次的产品").Optional(),
		field.Int64("entry_sum").Default(0).Comment("进馆总次数").Optional(),
		field.Time("entry_last_at").Optional().Comment("最后一次进馆时间"),
		field.Time("entry_deadline_at").Optional().Comment("进馆最后期限时间"),
		field.Time("class_last_at").Optional().Comment("最后一次上课时间"),

		field.Int64("relation_uid").Default(0).Comment("跟进人员工").Optional(),
		field.String("relation_uname").Comment("跟进人员工").Optional(),
	}
}

func (VenueMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (VenueMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("member", Member.Type).
			Ref("member_details").
			Field("member_id").
			Unique(),
	}
}

func (VenueMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("member_id"),
		index.Fields("venue_id"),
	}
}

func (VenueMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_details"},
		entsql.WithComments(true),
	}
}
