package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Remark holds the schema definition for the Remark entity.
type Remark struct {
	ent.Schema
}

func (Remark) Fields() []ent.Field {
	return []ent.Field{
		field.Int("blog_id"),
		field.Bool("is_reply").Default(false),
		field.String("target_id"),
		field.String("content").NotEmpty(),
		field.String("create_at").NotEmpty(),
		field.Bool("if_del").Default(false),
		field.String("user_name").NotEmpty(),
		field.String("target_name").NotEmpty(),
	}
}

func (Remark) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("blogs", Blog.Type).Ref("remarks").Unique(),
	}
}
