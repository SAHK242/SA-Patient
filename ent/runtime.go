// Code generated by ent, DO NOT EDIT.

package ent

import (
	"patient/ent/medicalhistories"
	"patient/ent/medicalprescription"
	"patient/ent/medicalsurgery"
	"patient/ent/medicaltreatment"
	"patient/ent/medication"
	"patient/ent/patient"
	"patient/ent/prescriptionmedication"
	"patient/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	medicalhistoriesFields := schema.MedicalHistories{}.Fields()
	_ = medicalhistoriesFields
	// medicalhistoriesDescReason is the schema descriptor for reason field.
	medicalhistoriesDescReason := medicalhistoriesFields[2].Descriptor()
	// medicalhistories.ReasonValidator is a validator for the "reason" field. It is called by the builders before save.
	medicalhistories.ReasonValidator = medicalhistoriesDescReason.Validators[0].(func(string) error)
	// medicalhistoriesDescDiagnosis is the schema descriptor for diagnosis field.
	medicalhistoriesDescDiagnosis := medicalhistoriesFields[3].Descriptor()
	// medicalhistories.DiagnosisValidator is a validator for the "diagnosis" field. It is called by the builders before save.
	medicalhistories.DiagnosisValidator = medicalhistoriesDescDiagnosis.Validators[0].(func(string) error)
	// medicalhistoriesDescHasTreatment is the schema descriptor for has_treatment field.
	medicalhistoriesDescHasTreatment := medicalhistoriesFields[4].Descriptor()
	// medicalhistories.DefaultHasTreatment holds the default value on creation for the has_treatment field.
	medicalhistories.DefaultHasTreatment = medicalhistoriesDescHasTreatment.Default.(bool)
	// medicalhistoriesDescHasSurgery is the schema descriptor for has_surgery field.
	medicalhistoriesDescHasSurgery := medicalhistoriesFields[5].Descriptor()
	// medicalhistories.DefaultHasSurgery holds the default value on creation for the has_surgery field.
	medicalhistories.DefaultHasSurgery = medicalhistoriesDescHasSurgery.Default.(bool)
	// medicalhistoriesDescHasPrescription is the schema descriptor for has_prescription field.
	medicalhistoriesDescHasPrescription := medicalhistoriesFields[6].Descriptor()
	// medicalhistories.DefaultHasPrescription holds the default value on creation for the has_prescription field.
	medicalhistories.DefaultHasPrescription = medicalhistoriesDescHasPrescription.Default.(bool)
	// medicalhistoriesDescCreatedAt is the schema descriptor for created_at field.
	medicalhistoriesDescCreatedAt := medicalhistoriesFields[9].Descriptor()
	// medicalhistories.DefaultCreatedAt holds the default value on creation for the created_at field.
	medicalhistories.DefaultCreatedAt = medicalhistoriesDescCreatedAt.Default.(func() time.Time)
	// medicalhistoriesDescUpdatedAt is the schema descriptor for updated_at field.
	medicalhistoriesDescUpdatedAt := medicalhistoriesFields[10].Descriptor()
	// medicalhistories.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	medicalhistories.DefaultUpdatedAt = medicalhistoriesDescUpdatedAt.Default.(func() time.Time)
	// medicalhistories.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	medicalhistories.UpdateDefaultUpdatedAt = medicalhistoriesDescUpdatedAt.UpdateDefault.(func() time.Time)
	// medicalhistoriesDescID is the schema descriptor for id field.
	medicalhistoriesDescID := medicalhistoriesFields[0].Descriptor()
	// medicalhistories.DefaultID holds the default value on creation for the id field.
	medicalhistories.DefaultID = medicalhistoriesDescID.Default.(func() uuid.UUID)
	medicalprescriptionFields := schema.MedicalPrescription{}.Fields()
	_ = medicalprescriptionFields
	// medicalprescriptionDescPrescriptionDate is the schema descriptor for prescription_date field.
	medicalprescriptionDescPrescriptionDate := medicalprescriptionFields[2].Descriptor()
	// medicalprescription.DefaultPrescriptionDate holds the default value on creation for the prescription_date field.
	medicalprescription.DefaultPrescriptionDate = medicalprescriptionDescPrescriptionDate.Default.(func() time.Time)
	// medicalprescriptionDescFee is the schema descriptor for fee field.
	medicalprescriptionDescFee := medicalprescriptionFields[3].Descriptor()
	// medicalprescription.DefaultFee holds the default value on creation for the fee field.
	medicalprescription.DefaultFee = medicalprescriptionDescFee.Default.(float64)
	// medicalprescriptionDescCreatedAt is the schema descriptor for created_at field.
	medicalprescriptionDescCreatedAt := medicalprescriptionFields[4].Descriptor()
	// medicalprescription.DefaultCreatedAt holds the default value on creation for the created_at field.
	medicalprescription.DefaultCreatedAt = medicalprescriptionDescCreatedAt.Default.(func() time.Time)
	// medicalprescriptionDescID is the schema descriptor for id field.
	medicalprescriptionDescID := medicalprescriptionFields[0].Descriptor()
	// medicalprescription.DefaultID holds the default value on creation for the id field.
	medicalprescription.DefaultID = medicalprescriptionDescID.Default.(func() uuid.UUID)
	medicalsurgeryFields := schema.MedicalSurgery{}.Fields()
	_ = medicalsurgeryFields
	// medicalsurgeryDescStartDate is the schema descriptor for start_date field.
	medicalsurgeryDescStartDate := medicalsurgeryFields[2].Descriptor()
	// medicalsurgery.DefaultStartDate holds the default value on creation for the start_date field.
	medicalsurgery.DefaultStartDate = medicalsurgeryDescStartDate.Default.(func() time.Time)
	// medicalsurgeryDescName is the schema descriptor for name field.
	medicalsurgeryDescName := medicalsurgeryFields[4].Descriptor()
	// medicalsurgery.NameValidator is a validator for the "name" field. It is called by the builders before save.
	medicalsurgery.NameValidator = medicalsurgeryDescName.Validators[0].(func(string) error)
	// medicalsurgeryDescFee is the schema descriptor for fee field.
	medicalsurgeryDescFee := medicalsurgeryFields[7].Descriptor()
	// medicalsurgery.DefaultFee holds the default value on creation for the fee field.
	medicalsurgery.DefaultFee = medicalsurgeryDescFee.Default.(float64)
	// medicalsurgeryDescCreatedAt is the schema descriptor for created_at field.
	medicalsurgeryDescCreatedAt := medicalsurgeryFields[11].Descriptor()
	// medicalsurgery.DefaultCreatedAt holds the default value on creation for the created_at field.
	medicalsurgery.DefaultCreatedAt = medicalsurgeryDescCreatedAt.Default.(func() time.Time)
	// medicalsurgeryDescUpdatedAt is the schema descriptor for updated_at field.
	medicalsurgeryDescUpdatedAt := medicalsurgeryFields[14].Descriptor()
	// medicalsurgery.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	medicalsurgery.DefaultUpdatedAt = medicalsurgeryDescUpdatedAt.Default.(func() time.Time)
	// medicalsurgery.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	medicalsurgery.UpdateDefaultUpdatedAt = medicalsurgeryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// medicalsurgeryDescID is the schema descriptor for id field.
	medicalsurgeryDescID := medicalsurgeryFields[0].Descriptor()
	// medicalsurgery.DefaultID holds the default value on creation for the id field.
	medicalsurgery.DefaultID = medicalsurgeryDescID.Default.(func() uuid.UUID)
	medicaltreatmentFields := schema.MedicalTreatment{}.Fields()
	_ = medicaltreatmentFields
	// medicaltreatmentDescStartDate is the schema descriptor for start_date field.
	medicaltreatmentDescStartDate := medicaltreatmentFields[2].Descriptor()
	// medicaltreatment.DefaultStartDate holds the default value on creation for the start_date field.
	medicaltreatment.DefaultStartDate = medicaltreatmentDescStartDate.Default.(func() time.Time)
	// medicaltreatmentDescName is the schema descriptor for name field.
	medicaltreatmentDescName := medicaltreatmentFields[4].Descriptor()
	// medicaltreatment.NameValidator is a validator for the "name" field. It is called by the builders before save.
	medicaltreatment.NameValidator = medicaltreatmentDescName.Validators[0].(func(string) error)
	// medicaltreatmentDescFee is the schema descriptor for fee field.
	medicaltreatmentDescFee := medicaltreatmentFields[7].Descriptor()
	// medicaltreatment.DefaultFee holds the default value on creation for the fee field.
	medicaltreatment.DefaultFee = medicaltreatmentDescFee.Default.(float64)
	// medicaltreatmentDescCreatedAt is the schema descriptor for created_at field.
	medicaltreatmentDescCreatedAt := medicaltreatmentFields[11].Descriptor()
	// medicaltreatment.DefaultCreatedAt holds the default value on creation for the created_at field.
	medicaltreatment.DefaultCreatedAt = medicaltreatmentDescCreatedAt.Default.(func() time.Time)
	// medicaltreatmentDescUpdatedAt is the schema descriptor for updated_at field.
	medicaltreatmentDescUpdatedAt := medicaltreatmentFields[14].Descriptor()
	// medicaltreatment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	medicaltreatment.DefaultUpdatedAt = medicaltreatmentDescUpdatedAt.Default.(func() time.Time)
	// medicaltreatment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	medicaltreatment.UpdateDefaultUpdatedAt = medicaltreatmentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// medicaltreatmentDescID is the schema descriptor for id field.
	medicaltreatmentDescID := medicaltreatmentFields[0].Descriptor()
	// medicaltreatment.DefaultID holds the default value on creation for the id field.
	medicaltreatment.DefaultID = medicaltreatmentDescID.Default.(func() uuid.UUID)
	medicationFields := schema.Medication{}.Fields()
	_ = medicationFields
	// medicationDescName is the schema descriptor for name field.
	medicationDescName := medicationFields[1].Descriptor()
	// medication.NameValidator is a validator for the "name" field. It is called by the builders before save.
	medication.NameValidator = medicationDescName.Validators[0].(func(string) error)
	// medicationDescQuantity is the schema descriptor for quantity field.
	medicationDescQuantity := medicationFields[4].Descriptor()
	// medication.DefaultQuantity holds the default value on creation for the quantity field.
	medication.DefaultQuantity = medicationDescQuantity.Default.(int64)
	// medicationDescPrice is the schema descriptor for price field.
	medicationDescPrice := medicationFields[5].Descriptor()
	// medication.DefaultPrice holds the default value on creation for the price field.
	medication.DefaultPrice = medicationDescPrice.Default.(float64)
	// medicationDescCreatedAt is the schema descriptor for created_at field.
	medicationDescCreatedAt := medicationFields[8].Descriptor()
	// medication.DefaultCreatedAt holds the default value on creation for the created_at field.
	medication.DefaultCreatedAt = medicationDescCreatedAt.Default.(func() time.Time)
	// medicationDescUpdatedAt is the schema descriptor for updated_at field.
	medicationDescUpdatedAt := medicationFields[9].Descriptor()
	// medication.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	medication.DefaultUpdatedAt = medicationDescUpdatedAt.Default.(func() time.Time)
	// medication.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	medication.UpdateDefaultUpdatedAt = medicationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// medicationDescID is the schema descriptor for id field.
	medicationDescID := medicationFields[0].Descriptor()
	// medication.DefaultID holds the default value on creation for the id field.
	medication.DefaultID = medicationDescID.Default.(func() uuid.UUID)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescGender is the schema descriptor for gender field.
	patientDescGender := patientFields[4].Descriptor()
	// patient.GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	patient.GenderValidator = patientDescGender.Validators[0].(func(int32) error)
	// patientDescCreatedAt is the schema descriptor for created_at field.
	patientDescCreatedAt := patientFields[7].Descriptor()
	// patient.DefaultCreatedAt holds the default value on creation for the created_at field.
	patient.DefaultCreatedAt = patientDescCreatedAt.Default.(func() time.Time)
	// patientDescUpdatedAt is the schema descriptor for updated_at field.
	patientDescUpdatedAt := patientFields[8].Descriptor()
	// patient.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	patient.DefaultUpdatedAt = patientDescUpdatedAt.Default.(func() time.Time)
	// patient.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	patient.UpdateDefaultUpdatedAt = patientDescUpdatedAt.UpdateDefault.(func() time.Time)
	// patientDescID is the schema descriptor for id field.
	patientDescID := patientFields[0].Descriptor()
	// patient.DefaultID holds the default value on creation for the id field.
	patient.DefaultID = patientDescID.Default.(func() uuid.UUID)
	prescriptionmedicationFields := schema.PrescriptionMedication{}.Fields()
	_ = prescriptionmedicationFields
	// prescriptionmedicationDescQuantity is the schema descriptor for quantity field.
	prescriptionmedicationDescQuantity := prescriptionmedicationFields[3].Descriptor()
	// prescriptionmedication.DefaultQuantity holds the default value on creation for the quantity field.
	prescriptionmedication.DefaultQuantity = prescriptionmedicationDescQuantity.Default.(int64)
	// prescriptionmedicationDescID is the schema descriptor for id field.
	prescriptionmedicationDescID := prescriptionmedicationFields[0].Descriptor()
	// prescriptionmedication.DefaultID holds the default value on creation for the id field.
	prescriptionmedication.DefaultID = prescriptionmedicationDescID.Default.(func() uuid.UUID)
}
