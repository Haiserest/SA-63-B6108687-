// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Haiserest/app/ent/patient"
	"github.com/Haiserest/app/ent/statustracking"
	"github.com/Haiserest/app/ent/treatment"
	"github.com/Haiserest/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// TreatmentCreate is the builder for creating a Treatment entity.
type TreatmentCreate struct {
	config
	mutation *TreatmentMutation
	hooks    []Hook
}

// SetDateAppoint sets the date_appoint field.
func (tc *TreatmentCreate) SetDateAppoint(t time.Time) *TreatmentCreate {
	tc.mutation.SetDateAppoint(t)
	return tc
}

// SetNillableDateAppoint sets the date_appoint field if the given value is not nil.
func (tc *TreatmentCreate) SetNillableDateAppoint(t *time.Time) *TreatmentCreate {
	if t != nil {
		tc.SetDateAppoint(*t)
	}
	return tc
}

// SetSymptom sets the symptom field.
func (tc *TreatmentCreate) SetSymptom(s string) *TreatmentCreate {
	tc.mutation.SetSymptom(s)
	return tc
}

// SetPatientIDID sets the patient_id edge to Patient by id.
func (tc *TreatmentCreate) SetPatientIDID(id int) *TreatmentCreate {
	tc.mutation.SetPatientIDID(id)
	return tc
}

// SetNillablePatientIDID sets the patient_id edge to Patient by id if the given value is not nil.
func (tc *TreatmentCreate) SetNillablePatientIDID(id *int) *TreatmentCreate {
	if id != nil {
		tc = tc.SetPatientIDID(*id)
	}
	return tc
}

// SetPatientID sets the patient_id edge to Patient.
func (tc *TreatmentCreate) SetPatientID(p *Patient) *TreatmentCreate {
	return tc.SetPatientIDID(p.ID)
}

// SetDoctorIDID sets the doctor_id edge to User by id.
func (tc *TreatmentCreate) SetDoctorIDID(id int) *TreatmentCreate {
	tc.mutation.SetDoctorIDID(id)
	return tc
}

// SetNillableDoctorIDID sets the doctor_id edge to User by id if the given value is not nil.
func (tc *TreatmentCreate) SetNillableDoctorIDID(id *int) *TreatmentCreate {
	if id != nil {
		tc = tc.SetDoctorIDID(*id)
	}
	return tc
}

// SetDoctorID sets the doctor_id edge to User.
func (tc *TreatmentCreate) SetDoctorID(u *User) *TreatmentCreate {
	return tc.SetDoctorIDID(u.ID)
}

// SetStatusIDID sets the status_id edge to StatusTracking by id.
func (tc *TreatmentCreate) SetStatusIDID(id int) *TreatmentCreate {
	tc.mutation.SetStatusIDID(id)
	return tc
}

// SetNillableStatusIDID sets the status_id edge to StatusTracking by id if the given value is not nil.
func (tc *TreatmentCreate) SetNillableStatusIDID(id *int) *TreatmentCreate {
	if id != nil {
		tc = tc.SetStatusIDID(*id)
	}
	return tc
}

// SetStatusID sets the status_id edge to StatusTracking.
func (tc *TreatmentCreate) SetStatusID(s *StatusTracking) *TreatmentCreate {
	return tc.SetStatusIDID(s.ID)
}

// Mutation returns the TreatmentMutation object of the builder.
func (tc *TreatmentCreate) Mutation() *TreatmentMutation {
	return tc.mutation
}

// Save creates the Treatment in the database.
func (tc *TreatmentCreate) Save(ctx context.Context) (*Treatment, error) {
	if _, ok := tc.mutation.DateAppoint(); !ok {
		v := treatment.DefaultDateAppoint()
		tc.mutation.SetDateAppoint(v)
	}
	if _, ok := tc.mutation.Symptom(); !ok {
		return nil, &ValidationError{Name: "symptom", err: errors.New("ent: missing required field \"symptom\"")}
	}
	var (
		err  error
		node *Treatment
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TreatmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TreatmentCreate) SaveX(ctx context.Context) *Treatment {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TreatmentCreate) sqlSave(ctx context.Context) (*Treatment, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TreatmentCreate) createSpec() (*Treatment, *sqlgraph.CreateSpec) {
	var (
		t     = &Treatment{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: treatment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: treatment.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.DateAppoint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: treatment.FieldDateAppoint,
		})
		t.DateAppoint = value
	}
	if value, ok := tc.mutation.Symptom(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: treatment.FieldSymptom,
		})
		t.Symptom = value
	}
	if nodes := tc.mutation.PatientIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.PatientIDTable,
			Columns: []string{treatment.PatientIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.DoctorIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.DoctorIDTable,
			Columns: []string{treatment.DoctorIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.StatusIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   treatment.StatusIDTable,
			Columns: []string{treatment.StatusIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: statustracking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return t, _spec
}
