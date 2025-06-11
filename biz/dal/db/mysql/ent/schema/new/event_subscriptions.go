package new

import (
	"kcers/biz/dal/db/mysql/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type EventSubscriptions struct {
	ent.Schema
}

func (EventSubscriptions) Fields() []ent.Field {
	return []ent.Field{

		field.String("name").Comment("名称").Optional(),
		field.String("event_type").Comment("订阅的事件类型").Optional(),
		field.String("last_processed_id").Comment("最后处理的事件ID").Optional(),
		field.String("last_processed_version").Comment("最后处理的事件版本").Optional(),
		field.String("last_processed_at").Comment("最后处理时间").Optional(),
		field.Int64("is_active").Comment("是否活跃").Optional(),
		field.String("error_count").Comment("处理错误次数").Optional(),
		field.String("last_error").Comment("最后错误信息").Optional(),
	}
}

func (EventSubscriptions) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (EventSubscriptions) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("amount", OrderAmount.Type),
		edge.To("item", OrderItem.Type),
		edge.To("pay", OrderPay.Type),
		edge.To("order_contents", MemberContract.Type),
		edge.To("sales", OrderSales.Type),

		edge.From("order_venues", Venue.Type).Ref("venue_orders").Field("venue_id").Unique(),
		edge.From("order_members", Member.Type).Ref("member_orders").Field("member_id").Unique(),
		edge.From("order_creates", User.Type).Ref("created_orders").Field("created_id").Unique(),
	}
}

func (EventSubscriptions) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (EventSubscriptions) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "event_subscriptions"},
		entsql.WithComments(true),
	}
}
