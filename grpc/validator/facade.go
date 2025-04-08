package grpcvalidator

import (
	"context"
	"patient/proto/patient"
)

type (
	PatientValidator interface {
		ValidateUpsertPatientRequest(ctx context.Context, req *patient.UpsertPatientRequest) error
	}

	MedicationValidator interface {
		ValidateUpsertMedicationRequest(ctx context.Context, req *patient.UpsertMedicationRequest) error
	}

	MedicalRecordValidator interface {
		ValidateUpsertMedicalRecordRequest(ctx context.Context, req *patient.UpsertMedicalRecordRequest) error
	}

	MedicalTreatmentValidator interface {
		ValidateUpsertMedicalTreatmentRequest(ctx context.Context, req *patient.UpsertMedicalTreatmentRequest) error
	}

	MedicalSurgeryValidator interface {
		ValidateUpsertMedicalSurgeryRequest(ctx context.Context, req *patient.UpsertMedicalSurgeryRequest) error
	}

	MedicalPrescriptionValidator interface {
		ValidateUpsertMedicalPrescriptionRequest(ctx context.Context, req *patient.UpsertMedicalPrescriptionRequest) error
	}
)
