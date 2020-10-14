// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/Haiserest/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DoctorEmail applies equality check predicate on the "doctor_email" field. It's identical to DoctorEmailEQ.
func DoctorEmail(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorEmail), v))
	})
}

// DoctorName applies equality check predicate on the "doctor_name" field. It's identical to DoctorNameEQ.
func DoctorName(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorName), v))
	})
}

// DoctorEmailEQ applies the EQ predicate on the "doctor_email" field.
func DoctorEmailEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailNEQ applies the NEQ predicate on the "doctor_email" field.
func DoctorEmailNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailIn applies the In predicate on the "doctor_email" field.
func DoctorEmailIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDoctorEmail), v...))
	})
}

// DoctorEmailNotIn applies the NotIn predicate on the "doctor_email" field.
func DoctorEmailNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDoctorEmail), v...))
	})
}

// DoctorEmailGT applies the GT predicate on the "doctor_email" field.
func DoctorEmailGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailGTE applies the GTE predicate on the "doctor_email" field.
func DoctorEmailGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailLT applies the LT predicate on the "doctor_email" field.
func DoctorEmailLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailLTE applies the LTE predicate on the "doctor_email" field.
func DoctorEmailLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailContains applies the Contains predicate on the "doctor_email" field.
func DoctorEmailContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailHasPrefix applies the HasPrefix predicate on the "doctor_email" field.
func DoctorEmailHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailHasSuffix applies the HasSuffix predicate on the "doctor_email" field.
func DoctorEmailHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailEqualFold applies the EqualFold predicate on the "doctor_email" field.
func DoctorEmailEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDoctorEmail), v))
	})
}

// DoctorEmailContainsFold applies the ContainsFold predicate on the "doctor_email" field.
func DoctorEmailContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDoctorEmail), v))
	})
}

// DoctorNameEQ applies the EQ predicate on the "doctor_name" field.
func DoctorNameEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDoctorName), v))
	})
}

// DoctorNameNEQ applies the NEQ predicate on the "doctor_name" field.
func DoctorNameNEQ(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDoctorName), v))
	})
}

// DoctorNameIn applies the In predicate on the "doctor_name" field.
func DoctorNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDoctorName), v...))
	})
}

// DoctorNameNotIn applies the NotIn predicate on the "doctor_name" field.
func DoctorNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDoctorName), v...))
	})
}

// DoctorNameGT applies the GT predicate on the "doctor_name" field.
func DoctorNameGT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDoctorName), v))
	})
}

// DoctorNameGTE applies the GTE predicate on the "doctor_name" field.
func DoctorNameGTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDoctorName), v))
	})
}

// DoctorNameLT applies the LT predicate on the "doctor_name" field.
func DoctorNameLT(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDoctorName), v))
	})
}

// DoctorNameLTE applies the LTE predicate on the "doctor_name" field.
func DoctorNameLTE(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDoctorName), v))
	})
}

// DoctorNameContains applies the Contains predicate on the "doctor_name" field.
func DoctorNameContains(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDoctorName), v))
	})
}

// DoctorNameHasPrefix applies the HasPrefix predicate on the "doctor_name" field.
func DoctorNameHasPrefix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDoctorName), v))
	})
}

// DoctorNameHasSuffix applies the HasSuffix predicate on the "doctor_name" field.
func DoctorNameHasSuffix(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDoctorName), v))
	})
}

// DoctorNameEqualFold applies the EqualFold predicate on the "doctor_name" field.
func DoctorNameEqualFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDoctorName), v))
	})
}

// DoctorNameContainsFold applies the ContainsFold predicate on the "doctor_name" field.
func DoctorNameContainsFold(v string) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDoctorName), v))
	})
}

// HasDoctorID applies the HasEdge predicate on the "doctor_id" edge.
func HasDoctorID() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorIDTable, DoctorIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDoctorIDWith applies the HasEdge predicate on the "doctor_id" edge with a given conditions (other predicates).
func HasDoctorIDWith(preds ...predicate.Treatment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DoctorIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DoctorIDTable, DoctorIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
