package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("subtitle").NotEmpty(),
		field.String("time").Default(time.Now().Format("2006-01-02")),
		field.String("cover"),
		field.Int("hits").Positive(),
		field.String("content_url").NotEmpty(),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("remarks", Remark.Type),
	}
}
