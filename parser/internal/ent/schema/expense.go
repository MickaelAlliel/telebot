package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type Expense struct {
	ent.Schema
}

func (Expense) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount"),
		field.String("category"),
		field.String("method"),
		field.String("ownerName"),
		field.Time("timestamp"),
		field.Time("createdAt").Default(time.Now()).Immutable(),
	}
}

func (Expense) Edges() []ent.Edge {
	return nil
}

func (Expense) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("amount"),
		index.Fields("category"),
		index.Fields("method"),
		index.Fields("ownerName"),
		index.Fields("timestamp"),
	}
}
