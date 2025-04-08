package grpcvalidator

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"patient/proto/patient"
	"patient/repository"
)

type medicalPrescriptionGrpcValidator struct {
	logger                     *zap.SugaredLogger
	medicalHistoriesRepository repository.MedicalHistoriesRepository
	medicationRepository       repository.MedicationRepository
}

func NewMedicalPrescriptionValidator(
	logger *zap.SugaredLogger,
	medicalHistoriesRepository repository.MedicalHistoriesRepository,
	medicationRepository repository.MedicationRepository,
) MedicalPrescriptionValidator {
	return &medicalPrescriptionGrpcValidator{
		logger:                     logger,
		medicalHistoriesRepository: medicalHistoriesRepository,
		medicationRepository:       medicationRepository,
	}
}

func (v *medicalPrescriptionGrpcValidator) ValidateUpsertMedicalPrescriptionRequest(ctx context.Context, req *patient.UpsertMedicalPrescriptionRequest) error {
	exist, err := v.medicalHistoriesRepository.ExistMedicalHistoryById(ctx, uuid.MustParse(req.MedicalHistoryId))

	if err != nil {
		return err
	}

	if !exist {
		return status.Error(codes.InvalidArgument, "medical record not found")
	}

	medicationIds := make([]uuid.UUID, 0)

	for _, y := range req.Medications {
		if y.Quantity <= 0 {
			return status.Error(codes.InvalidArgument, "quantity must > 0")
		}
		medicationIds = append(medicationIds, uuid.MustParse(y.MedicationId))
	}

	existMedication, err := v.medicationRepository.ExistMedicationByIdIn(ctx, medicationIds)

	if err != nil {
		return err
	}

	if !existMedication {
		return status.Error(codes.InvalidArgument, "medication id not valid")
	}

	return nil
}
