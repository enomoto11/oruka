package schema

import (
	"oruka/ent/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// TimeRecord holds the schema definition for the TimeRecord entity.
type TimeRecord struct {
	ent.Schema
}

// Fields of the TimeRecord.
func (TimeRecord) Fields() []ent.Field {
	return nil
}

// Edges of the TimeRecord.
func (TimeRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("timeKeeper", User.Type).Ref("timeRecords").Unique(),
	}
}

func (TimeRecord) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}
