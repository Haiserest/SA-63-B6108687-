// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Haiserest/app/ent/patient"
	"github.com/Haiserest/app/ent/schema"
	"github.com/Haiserest/app/ent/statustracking"
	"github.com/Haiserest/app/ent/treatment"
	"github.com/Haiserest/app/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescPatientAge is the schema descriptor for patient_age field.
	patientDescPatientAge := patientFields[0].Descriptor()
	// patient.PatientAgeValidator is a validator for the "patient_age" field. It is called by the builders before save.
	patient.PatientAgeValidator = patientDescPatientAge.Validators[0].(func(int) error)
	// patientDescPatientName is the schema descriptor for patient_name field.
	patientDescPatientName := patientFields[1].Descriptor()
	// patient.PatientNameValidator is a validator for the "patient_name" field. It is called by the builders before save.
	patient.PatientNameValidator = patientDescPatientName.Validators[0].(func(string) error)
	statustrackingFields := schema.StatusTracking{}.Fields()
	_ = statustrackingFields
	// statustrackingDescStatus is the schema descriptor for status field.
	statustrackingDescStatus := statustrackingFields[0].Descriptor()
	// statustracking.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	statustracking.StatusValidator = statustrackingDescStatus.Validators[0].(func(string) error)
	treatmentFields := schema.Treatment{}.Fields()
	_ = treatmentFields
	// treatmentDescDateAppoint is the schema descriptor for date_appoint field.
	treatmentDescDateAppoint := treatmentFields[0].Descriptor()
	// treatment.DefaultDateAppoint holds the default value on creation for the date_appoint field.
	treatment.DefaultDateAppoint = treatmentDescDateAppoint.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescDoctorEmail is the schema descriptor for doctor_email field.
	userDescDoctorEmail := userFields[0].Descriptor()
	// user.DoctorEmailValidator is a validator for the "doctor_email" field. It is called by the builders before save.
	user.DoctorEmailValidator = userDescDoctorEmail.Validators[0].(func(string) error)
	// userDescDoctorName is the schema descriptor for doctor_name field.
	userDescDoctorName := userFields[1].Descriptor()
	// user.DoctorNameValidator is a validator for the "doctor_name" field. It is called by the builders before save.
	user.DoctorNameValidator = userDescDoctorName.Validators[0].(func(string) error)
}