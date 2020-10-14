// Code generated by entc, DO NOT EDIT.

package patient

const (
	// Label holds the string label denoting the patient type in the database.
	Label = "patient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPatientAge holds the string denoting the patient_age field in the database.
	FieldPatientAge = "patient_age"
	// FieldPatientName holds the string denoting the patient_name field in the database.
	FieldPatientName = "patient_name"

	// EdgePatientID holds the string denoting the patient_id edge name in mutations.
	EdgePatientID = "patient_id"

	// Table holds the table name of the patient in the database.
	Table = "patients"
	// PatientIDTable is the table the holds the patient_id relation/edge.
	PatientIDTable = "treatments"
	// PatientIDInverseTable is the table name for the Treatment entity.
	// It exists in this package in order to avoid circular dependency with the "treatment" package.
	PatientIDInverseTable = "treatments"
	// PatientIDColumn is the table column denoting the patient_id relation/edge.
	PatientIDColumn = "patient_patient_id"
)

// Columns holds all SQL columns for patient fields.
var Columns = []string{
	FieldID,
	FieldPatientAge,
	FieldPatientName,
}

var (
	// PatientAgeValidator is a validator for the "patient_age" field. It is called by the builders before save.
	PatientAgeValidator func(int) error
	// PatientNameValidator is a validator for the "patient_name" field. It is called by the builders before save.
	PatientNameValidator func(string) error
)