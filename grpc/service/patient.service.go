package grpcservice

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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

type patientGrpcService struct {
	config               config.Config
	client               *ent.Client
	logger               *zap.SugaredLogger
	redis                *redis.Client
	patientRepository    repository.PatientRepository
	patientGrpcMapper    mapper.PatientGrpcMapper
	patientGrpcValidator validator.PatientValidator
}

func (s *patientGrpcService) GetPatient(ctx context.Context, request *gcommon.IdRequest) (*patient.GetPatientResponse, error) {
	patientById, err := s.patientRepository.FindById(ctx, uuid.MustParse(request.Id))
	if err != nil {
		return nil, fmt.Errorf("error while finding employee: %v", err)
	}

	return &patient.GetPatientResponse{
		PatientDetail: s.patientGrpcMapper.ConvertPatient(patientById),
	}, nil
}

func (s *patientGrpcService) UpsertPatient(ctx context.Context, request *patient.UpsertPatientRequest) (*gcommon.EmptyResponse, error) {
	err := s.patientGrpcValidator.ValidateUpsertPatientRequest(ctx, request)

	if err != nil {
		return nil, err
	}

	err = withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.patientRepository.UpsertPatient(ctx, tx, request)
	})

	if err != nil {
		return nil, err
	}

	return &gcommon.EmptyResponse{}, nil
}

func (s *patientGrpcService) ListPatient(ctx context.Context, request *patient.ListPatientRequest) (*patient.ListPatientResponse, error) {
	patients, total, err := s.patientRepository.FindAllPatient(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error while finding patient: %v", err)
	}

	return &patient.ListPatientResponse{
		PageMetadata:   grpcutil.AsPageMetadata(request.Pageable, total),
		PatientDetails: s.patientGrpcMapper.ConvertPatientSlice(patients),
	}, nil
}

func NewPatientGrpcService(
	redis *redis.Client,
	patientRepository repository.PatientRepository,
	client *ent.Client,
	logger *zap.SugaredLogger,
	config config.Config,
	patientGrpcMapper mapper.PatientGrpcMapper,
	patientGrpcValidator validator.PatientValidator,
) PatientGrpcService {
	return &patientGrpcService{
		redis:                redis,
		patientRepository:    patientRepository,
		client:               client,
		logger:               logger,
		config:               config,
		patientGrpcMapper:    patientGrpcMapper,
		patientGrpcValidator: patientGrpcValidator,
	}
}
