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
	config                        config.Config
	client                        *ent.Client
	logger                        *zap.SugaredLogger
	redis                         *redis.Client
	medicalHistoriesRepository    repository.MedicalHistoriesRepository
	medicalHistoryGrpcMapper      mapper.MedicalHistoryGrpcMapper
	medicalTreatmentGrpcMapper    mapper.MedicalTreatmentGrpcMapper
	medicalSurgeryGrpcMapper      mapper.MedicalSurgeryGrpcMapper
	medicalPrescriptionGrpcMapper mapper.MedicalPrescriptionGrpcMapper
	patientRepository             repository.PatientRepository
	patientGrpcMapper             mapper.PatientGrpcMapper
	patientGrpcValidator          validator.PatientValidator
	medicalTreatmentRepository    repository.MedicalTreatmentRepository
	medicalSurgeryRepository      repository.MedicalSurgeryRepository
	medicalPrescriptionRepository repository.MedicalPrescriptionRepository
	medicalPrescriptionValidator  validator.MedicalPrescriptionValidator
	medicalHistoryValidator       validator.MedicalRecordValidator
	medicalTreatmentValidator     validator.MedicalTreatmentValidator
	medicalSurgeryValidator       validator.MedicalSurgeryValidator
}

func NewPatientGrpcService(
	redis *redis.Client,
	client *ent.Client,
	logger *zap.SugaredLogger,
	config config.Config,
	medicalHistoriesRepository repository.MedicalHistoriesRepository,
	medicalHistoryGrpcMapper mapper.MedicalHistoryGrpcMapper,
	medicalTreatmentGrpcMapper mapper.MedicalTreatmentGrpcMapper,
	medicalSurgeryGrpcMapper mapper.MedicalSurgeryGrpcMapper,
	medicalPrescriptionGrpcMapper mapper.MedicalPrescriptionGrpcMapper,
	patientRepository repository.PatientRepository,
	patientGrpcMapper mapper.PatientGrpcMapper,
	patientGrpcValidator validator.PatientValidator,
	medicalTreatmentRepository repository.MedicalTreatmentRepository,
	medicalSurgeryRepository repository.MedicalSurgeryRepository,
	medicalPrescriptionRepository repository.MedicalPrescriptionRepository,
) PatientGrpcService {
	return &patientGrpcService{
		redis:                         redis,
		client:                        client,
		logger:                        logger,
		config:                        config,
		medicalHistoriesRepository:    medicalHistoriesRepository,
		medicalHistoryGrpcMapper:      medicalHistoryGrpcMapper,
		medicalTreatmentGrpcMapper:    medicalTreatmentGrpcMapper,
		medicalSurgeryGrpcMapper:      medicalSurgeryGrpcMapper,
		medicalPrescriptionGrpcMapper: medicalPrescriptionGrpcMapper,
		patientRepository:             patientRepository,
		patientGrpcMapper:             patientGrpcMapper,
		patientGrpcValidator:          patientGrpcValidator,
		medicalTreatmentRepository:    medicalTreatmentRepository,
		medicalSurgeryRepository:      medicalSurgeryRepository,
		medicalPrescriptionRepository: medicalPrescriptionRepository,
	}
}

func (s *patientGrpcService) GetMedicalHistoryDetail(ctx context.Context, request *gcommon.IdRequest) (*patient.GetMedicalHistoryDetailResponse, error) {
	medicalHistory, err := s.medicalHistoriesRepository.FindMedicalHistoryById(ctx, uuid.MustParse(request.Id))
	if err != nil {
		return nil, fmt.Errorf("error while finding medical history: %v", err)
	}

	return &patient.GetMedicalHistoryDetailResponse{
		MedicalHistory:       s.medicalHistoryGrpcMapper.ConvertMedicalHistory(medicalHistory),
		MedicalTreatments:    s.medicalTreatmentGrpcMapper.ConvertMedicalTreatmentSlice(medicalHistory.Edges.MedicalTreatment),
		MedicalSurgeries:     s.medicalSurgeryGrpcMapper.ConvertMedicalSurgerySlice(medicalHistory.Edges.MedicalSurgery),
		MedicalPrescriptions: s.medicalPrescriptionGrpcMapper.ConvertMedicalPrescriptionSlice(medicalHistory.Edges.MedicalPrescription),
	}, nil
}

func (s *patientGrpcService) GetPatient(ctx context.Context, request *gcommon.IdRequest) (*patient.GetPatientResponse, error) {
	patientById, err := s.patientRepository.FindById(ctx, uuid.MustParse(request.Id))
	if err != nil {
		return nil, fmt.Errorf("error while finding employee: %v", err)
	}

	return &patient.GetPatientResponse{
		Patient: s.patientGrpcMapper.ConvertPatient(patientById),
	}, nil
}

func (s *patientGrpcService) UpsertMedicalRecord(ctx context.Context, request *patient.UpsertMedicalRecordRequest) (*gcommon.EmptyResponse, error) {
	err := s.medicalHistoryValidator.ValidateUpsertMedicalRecordRequest(ctx, request)

	if err != nil {
		return nil, err
	}

	err = withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.medicalHistoriesRepository.UpsertMedicalRecord(ctx, tx, request)
	})

	if err != nil {
		return nil, err
	}

	return &gcommon.EmptyResponse{}, nil
}

func (s *patientGrpcService) UpsertMedicalTreatment(ctx context.Context, request *patient.UpsertMedicalTreatmentRequest) (*gcommon.EmptyResponse, error) {
	err := s.medicalTreatmentValidator.ValidateUpsertMedicalTreatmentRequest(ctx, request)

	if err != nil {
		return nil, err
	}

	err = withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.medicalTreatmentRepository.UpsertMedicalTreatment(ctx, tx, request)
	})

	if err != nil {
		return nil, err
	}

	return &gcommon.EmptyResponse{}, nil
}

func (s *patientGrpcService) UpsertMedicalSurgery(ctx context.Context, request *patient.UpsertMedicalSurgeryRequest) (*gcommon.EmptyResponse, error) {
	err := s.medicalSurgeryValidator.ValidateUpsertMedicalSurgeryRequest(ctx, request)

	if err != nil {
		return nil, err
	}

	err = withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.medicalSurgeryRepository.UpsertMedicalSurgery(ctx, tx, request)
	})

	if err != nil {
		return nil, err
	}

	return &gcommon.EmptyResponse{}, nil
}

func (s *patientGrpcService) UpsertMedicalPrescription(ctx context.Context, request *patient.UpsertMedicalPrescriptionRequest) (*gcommon.EmptyResponse, error) {
	err := s.medicalPrescriptionValidator.ValidateUpsertMedicalPrescriptionRequest(ctx, request)

	if err != nil {
		return nil, err
	}

	err = withTransaction(ctx, s.client, func(tx *ent.Tx) error {
		return s.medicalPrescriptionRepository.UpsertMedicalPrescription(ctx, tx, request)
	})

	if err != nil {
		return nil, err
	}

	return &gcommon.EmptyResponse{}, nil
}

func (s *patientGrpcService) GetMedicalHistory(ctx context.Context, request *patient.GetMedicalHistoryRequest) (*patient.GetMedicalHistoryResponse, error) {
	medicalHistories, total, err := s.medicalHistoriesRepository.FindMedicalHistories(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error while finding medical history: %v", err)
	}

	return &patient.GetMedicalHistoryResponse{
		PageMetadata:     grpcutil.AsPageMetadata(request.Pageable, total),
		MedicalHistories: s.medicalHistoryGrpcMapper.ConvertMedicalHistorySlice(medicalHistories),
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
		PageMetadata: grpcutil.AsPageMetadata(request.Pageable, total),
		Patients:     s.patientGrpcMapper.ConvertPatientSlice(patients),
	}, nil
}
