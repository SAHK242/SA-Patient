package grpcvalidator

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"patient/proto/gcommon"
	"patient/proto/patient"
	"patient/repository"
)

type patientGrpcValidator struct {
	logger            *zap.SugaredLogger
	patientRepository repository.PatientRepository
}

func NewPatientGrpcValidator(
	logger *zap.SugaredLogger,
	patientRepository repository.PatientRepository,
) PatientValidator {
	return &patientGrpcValidator{
		logger:            logger,
		patientRepository: patientRepository,
	}
}

func (v *patientGrpcValidator) ValidateUpsertPatientRequest(ctx context.Context, req *patient.UpsertPatientRequest) error {
	if req.FirstName == "" {
		return status.Error(codes.InvalidArgument, "first name is required")
	}

	if req.LastName == "" {
		return status.Error(codes.InvalidArgument, "last name is required")
	}

	if req.Gender != gcommon.Gender_GENDER_FEMALE && req.Gender != gcommon.Gender_GENDER_MALE && req.Gender != gcommon.Gender_GENDER_OTHER {
		return status.Error(codes.InvalidArgument, "gender type is invalid")
	}

	if req.DateOfBirth == 0 {
		return status.Error(codes.InvalidArgument, "date of birth is required")
	}

	if req.Address == "" {
		return status.Error(codes.InvalidArgument, "address is required")
	}

	if req.Phone == "" {
		return status.Error(codes.InvalidArgument, "phone number is required")
	}

	return nil
}
