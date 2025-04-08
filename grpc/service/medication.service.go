package grpcservice

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"patient/config"
	"patient/ent"
	mapper "patient/grpc/mapper"
	grpcutil "patient/grpc/util"
	validator "patient/grpc/validator"
	"patient/proto/gcommon"
	"patient/proto/patient"
	"patient/repository"
)

type medicationGrpcService struct {
	config                  config.Config
	client                  *ent.Client
	logger                  *zap.SugaredLogger
	redis                   *redis.Client
	medicationRepository    repository.MedicationRepository
	medicationGrpcMapper    mapper.MedicationGrpcMapper
	medicationGrpcValidator validator.MedicationValidator
}

func NewMedicationGrpcService(
	redis *redis.Client,
	client *ent.Client,
	logger *zap.SugaredLogger,
	config config.Config,
	medicationRepository repository.MedicationRepository,
	medicationGrpcMapper mapper.MedicationGrpcMapper,
	medicationGrpcValidator validator.MedicationValidator,
) MedicationGrpcService {
	return &medicationGrpcService{
		redis:                   redis,
		client:                  client,
		logger:                  logger,
		config:                  config,
		medicationRepository:    medicationRepository,
		medicationGrpcMapper:    medicationGrpcMapper,
		medicationGrpcValidator: medicationGrpcValidator,
	}
}

func (s *medicationGrpcService) ListMedication(ctx context.Context, request *patient.ListMedicationRequest) (*patient.ListMedicationResponse, error) {
	medications, total, err := s.medicationRepository.FindAllMedication(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error while finding medication: %v", err)
	}

	return &patient.ListMedicationResponse{
		PageMetadata: grpcutil.AsPageMetadata(request.Pageable, total),
		Medications:  s.medicationGrpcMapper.ConvertMedicationSlice(medications),
	}, nil
}

func (s *medicationGrpcService) UpsertMedication(ctx context.Context, request *patient.UpsertMedicationRequest) (*gcommon.EmptyResponse, error) {
	err := s.medicationGrpcValidator.ValidateUpsertMedicationRequest(ctx, request)

	if err != nil {
		return nil, err
	}

	err = withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.medicationRepository.UpsertMedication(ctx, tx, request)
	})

	if err != nil {
		return nil, err
	}

	return &gcommon.EmptyResponse{}, nil
}
