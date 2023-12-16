package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Chapter holds the schema definition for the Chapter entity.
type Chapter struct {
	ent.Schema
}

// Fields of the Chapter.
func (Chapter) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("description").NotEmpty(),
	}
}

// Edges of the Chapter.
func (Chapter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type),
		edge.To("pages", Page.Type),
		edge.From("book", Book.Type).Ref("chapters").Unique(),
	}
}

func (Chapter) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
