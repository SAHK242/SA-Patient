package grpcvalidator

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"patient/proto/patient"
)

type medicationGrpcValidator struct {
	logger *zap.SugaredLogger
}

func NewMedicationValidator(
	logger *zap.SugaredLogger,
) MedicationValidator {
	return &medicationGrpcValidator{
		logger: logger,
	}
}

func (v *medicationGrpcValidator) ValidateUpsertMedicationRequest(ctx context.Context, req *patient.UpsertMedicationRequest) error {
	if req.Effects == "" {
		return status.Error(codes.InvalidArgument, "effects is required")
	}

	if req.Name == "" {
		return status.Error(codes.InvalidArgument, "name is required")
	}

	if req.Price <= 0 {
		return status.Error(codes.InvalidArgument, "price must be greater than 0")
	}

	if req.Quantity <= 0 {
		return status.Error(codes.InvalidArgument, "quantity must be greater than 0")
	}

	if req.ExpiredDate == 0 {
		return status.Error(codes.InvalidArgument, "expiration date is required")
	}

	return nil
}
