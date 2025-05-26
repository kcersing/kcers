package schema

import (
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
	"kcers/biz/dal/db/mysql/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Contract struct {
	ent.Schema
}

func (Contract) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Comment("name | 名称"),
		field.String("content").Optional().Comment("content | 内容"),
	}
}

func (Contract) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Contract) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("property", ProductProperty.Type).Ref("contracts"),
	}
}

func (Contract) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}

func (Contract) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "contracts"},
		entsql.WithComments(true),
	}
}
