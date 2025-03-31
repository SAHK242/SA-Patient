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
	}
)
