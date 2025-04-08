package repository

import (
	"context"
	"github.com/google/uuid"
	"patient/ent"
	"patient/proto/patient"
)

type (
	PatientRepository interface {
		UpsertPatient(ctx context.Context, tx *ent.Tx, request *patient.UpsertPatientRequest) error
		FindAllPatient(ctx context.Context, request *patient.ListPatientRequest) ([]*ent.Patient, int, error)
		FindById(ctx context.Context, id uuid.UUID) (*ent.Patient, error)
		ExistByPatientId(ctx context.Context, id uuid.UUID) (bool, error)
	}

	MedicalTreatmentRepository interface {
		UpsertMedicalTreatment(ctx context.Context, tx *ent.Tx, request *patient.UpsertMedicalTreatmentRequest) error
	}

	MedicalSurgeryRepository interface {
		UpsertMedicalSurgery(ctx context.Context, tx *ent.Tx, request *patient.UpsertMedicalSurgeryRequest) error
	}

	MedicationRepository interface {
		FindAllMedication(ctx context.Context, request *patient.ListMedicationRequest) ([]*ent.Medication, int, error)
		UpsertMedication(ctx context.Context, tx *ent.Tx, request *patient.UpsertMedicationRequest) error
		ExistMedicationByIdIn(ctx context.Context, ids []uuid.UUID) (bool, error)
	}

	MedicalHistoriesRepository interface {
		FindMedicalHistories(ctx context.Context, request *patient.GetMedicalHistoryRequest) ([]*ent.MedicalHistories, int, error)
		FindMedicalHistoryById(ctx context.Context, id uuid.UUID) (*ent.MedicalHistories, error)
		UpsertMedicalRecord(ctx context.Context, tx *ent.Tx, request *patient.UpsertMedicalRecordRequest) error
		ExistMedicalHistoryById(ctx context.Context, id uuid.UUID) (bool, error)
	}

	MedicalPrescriptionRepository interface {
		UpsertMedicalPrescription(ctx context.Context, tx *ent.Tx, request *patient.UpsertMedicalPrescriptionRequest) error
	}
)
