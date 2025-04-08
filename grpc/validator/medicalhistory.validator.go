package grpcvalidator

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"patient/proto/patient"
	"patient/repository"
	"strings"
)

type medicalHistoryGrpcValidator struct {
	logger            *zap.SugaredLogger
	patientRepository repository.PatientRepository
}

func NewMedicalHistoryValidator(
	logger *zap.SugaredLogger,
	patientRepository repository.PatientRepository,
) MedicalRecordValidator {
	return &medicalHistoryGrpcValidator{
		logger:            logger,
		patientRepository: patientRepository,
	}
}

func (v *medicalHistoryGrpcValidator) ValidateUpsertMedicalRecordRequest(ctx context.Context, req *patient.UpsertMedicalRecordRequest) error {
	exist, err := v.patientRepository.ExistByPatientId(ctx, uuid.MustParse(req.PatientId))

	if err != nil {
		return status.Error(codes.Internal, "error while checking patient existence")
	}

	if !exist {
		return status.Error(codes.NotFound, "patient not found")
	}

	if strings.TrimSpace(req.Reason) == "" {
		return status.Error(codes.InvalidArgument, "reason is required")
	}

	if strings.TrimSpace(req.Diagnosis) == "" {
		return status.Error(codes.InvalidArgument, "diagnosis is required")
	}

	return nil
}
