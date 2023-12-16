package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("description").NotEmpty(),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type),
		edge.To("chapters", Chapter.Type),
		edge.From("user", User.Type).Ref("books").Unique(),
	}
}

func (Book) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
