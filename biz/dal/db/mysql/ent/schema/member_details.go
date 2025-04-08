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

type MemberDetails struct {
	ent.Schema
}

func (MemberDetails) Fields() []ent.Field {
	return []ent.Field{

		field.Int64("member_id").Comment("会员id").Optional(),
		field.Float("money_sum").Default(3).Comment("消费总金额").Optional(),
		field.Int64("product_id").Default(0).Comment("首次的产品").Optional(),
		field.String("product_name").Comment("首次的产品").Optional(),
		field.Int64("entry_sum").Default(0).Comment("进馆总次数").Optional(),
		field.Time("entry_last_at").Optional().Comment("最后一次进馆时间"),
		field.Time("entry_deadline_at").Optional().Comment("进馆最后期限时间"),
		field.Time("class_last_at").Optional().Comment("最后一次上课时间"),
	}
}

func (MemberDetails) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (MemberDetails) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("member", Member.Type).
			Ref("member_details").
			Field("member_id").
			Unique(),
	}
}

func (MemberDetails) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("member_id"),
	}
}

func (MemberDetails) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member_details", Options: "AUTO_INCREMENT = 100000"},
		entsql.WithComments(true),
	}
}
