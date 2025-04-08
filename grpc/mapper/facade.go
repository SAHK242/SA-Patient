package grpcmapper

import (
	"patient/ent"
	"patient/proto/patient"
)

type (
	PatientGrpcMapper interface {
		ConvertPatient(from *ent.Patient) *patient.Patient
		ConvertPatientSlice(from []*ent.Patient) []*patient.Patient
	}

	MedicationGrpcMapper interface {
		ConvertMedication(from *ent.Medication) *patient.Medication
		ConvertMedicationSlice(from []*ent.Medication) []*patient.Medication
	}

	MedicalHistoryGrpcMapper interface {
		ConvertMedicalHistory(from *ent.MedicalHistories) *patient.MedicalHistory
		ConvertMedicalHistorySlice(from []*ent.MedicalHistories) []*patient.MedicalHistory
	}

	MedicalTreatmentGrpcMapper interface {
		ConvertMedicalTreatment(from *ent.MedicalTreatment) *patient.MedicalTreatment
		ConvertMedicalTreatmentSlice(from []*ent.MedicalTreatment) []*patient.MedicalTreatment
	}

	MedicalSurgeryGrpcMapper interface {
		ConvertMedicalSurgery(from *ent.MedicalSurgery) *patient.MedicalSurgery
		ConvertMedicalSurgerySlice(from []*ent.MedicalSurgery) []*patient.MedicalSurgery
	}

	MedicalPrescriptionGrpcMapper interface {
		ConvertMedicalPrescription(from *ent.MedicalPrescription) *patient.MedicalPrescription
		ConvertMedicalPrescriptionSlice(from []*ent.MedicalPrescription) []*patient.MedicalPrescription
	}
)
