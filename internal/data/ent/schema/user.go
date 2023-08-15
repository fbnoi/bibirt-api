package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Unique(),
		field.Uint8("type"),
		field.String("name"),
		field.String("pwd").Optional(),
		field.String("salt").Optional(),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.Uint8("status"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
