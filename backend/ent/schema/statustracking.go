package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// StatusTracking holds the schema definition for the StatusTracking entity.
type StatusTracking struct {
	ent.Schema
}

// Fields of the StatusTracking.
func (StatusTracking) Fields() []ent.Field {
	return []ent.Field{
		field.String("status").NotEmpty().Unique(),
	}
}

// Edges of the StatusTracking.
func (StatusTracking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("status_id",Treatment.Type),
	}
}
