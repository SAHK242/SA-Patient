package grpcvalidator

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"patient/proto/patient"
	"patient/repository"
	"time"
)

type medicalSurgeryGrpcValidator struct {
	logger                     *zap.SugaredLogger
	medicalHistoriesRepository repository.MedicalHistoriesRepository
}

func NewMedicalSurgeryValidator(
	logger *zap.SugaredLogger,
	medicalHistoriesRepository repository.MedicalHistoriesRepository,
) MedicalSurgeryValidator {
	return &medicalSurgeryGrpcValidator{
		logger:                     logger,
		medicalHistoriesRepository: medicalHistoriesRepository,
	}
}

func (v *medicalSurgeryGrpcValidator) ValidateUpsertMedicalSurgeryRequest(ctx context.Context, req *patient.UpsertMedicalSurgeryRequest) error {
	exist, err := v.medicalHistoriesRepository.ExistMedicalHistoryById(ctx, uuid.MustParse(req.MedicalHistoryId))

	if err != nil {
		return err
	}

	if !exist {
		return status.Error(codes.InvalidArgument, "medical record not found")
	}

	if req.Name == "" {
		return status.Error(codes.InvalidArgument, "name is required")
	}

	if req.MainDoctorId == "" {
		return status.Error(codes.InvalidArgument, "main doctor id is required")
	}

	if req.EndDate != 0 && req.StartDate > req.EndDate {
		return status.Error(codes.InvalidArgument, "start date must be less than end date")
	}

	if req.Fee <= 0 {
		return status.Error(codes.InvalidArgument, "fee must be greater than 0")
	}

	loc := time.Now().Location()

	startDate := time.UnixMilli(req.StartDate).In(loc)
	now := time.Now().In(loc)

	// Compare only the day, month, and year
	if startDate.Year() < now.Year() ||
		(startDate.Year() == now.Year() && startDate.Month() < now.Month()) ||
		(startDate.Year() == now.Year() && startDate.Month() == now.Month() && startDate.Day() < now.Day()) {
		return status.Error(codes.InvalidArgument, "start date must be in the future")
	}

	return nil
}
