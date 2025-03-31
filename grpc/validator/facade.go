package grpcvalidator

import (
	"context"
	"patient/proto/patient"
)

type (
	PatientValidator interface {
		ValidateUpsertPatientRequest(ctx context.Context, req *patient.UpsertPatientRequest) error
	}
)
