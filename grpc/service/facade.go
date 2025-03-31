package grpcservice

import (
	"context"
	"patient/proto/gcommon"
	"patient/proto/patient"
)

type (
	PatientGrpcService interface {
		GetPatient(ctx context.Context, request *gcommon.IdRequest) (*patient.GetPatientResponse, error)
		ListPatient(ctx context.Context, request *patient.ListPatientRequest) (*patient.ListPatientResponse, error)
		UpsertPatient(ctx context.Context, request *patient.UpsertPatientRequest) (*gcommon.EmptyResponse, error)
	}
)
