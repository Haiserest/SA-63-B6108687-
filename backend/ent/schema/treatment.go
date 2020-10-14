package schema

import (
	"time"
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// Treatment holds the schema definition for the Treatment entity.
type Treatment struct {
	ent.Schema
}

// Fields of the Treatment.
func (Treatment) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date_appoint").Default(time.Now),
		field.String("symptom"),
	}
}

// Edges of the Treatment.
func (Treatment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("patient_id",Patient.Type).
			Ref("patient_id").
			Unique(),

		edge.From("doctor_id",User.Type).
			Ref("doctor_id").
			Unique(),

		edge.From("status_id",StatusTracking.Type).
			Ref("status_id").
			Unique(),
		}
}
