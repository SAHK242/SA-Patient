package repository

import (
	"context"
	"github.com/google/uuid"
	"patient/ent"
	"patient/ent/medicalhistories"
	grpcutil "patient/grpc/util"
	patient2 "patient/proto/patient"
	"time"
)
import db "database/sql"

type medicalhistoriesRepository struct {
	client *ent.Client
	db     *db.DB
}

func (p medicalhistoriesRepository) FindMedicalHistoryById(ctx context.Context, id uuid.UUID) (*ent.MedicalHistories, error) {
	return p.client.MedicalHistories.Query().Where(medicalhistories.ID(id)).WithMedicalSurgery().WithMedicalPrescription(
		func(query *ent.MedicalPrescriptionQuery) {
			query.WithPrescriptionMedication(
				func(query *ent.PrescriptionMedicationQuery) {
					query.WithMedication()
				})
		},
	).WithMedicalTreatment().Only(ctx)
}

func (p medicalhistoriesRepository) FindMedicalHistories(ctx context.Context, request *patient2.GetMedicalHistoryRequest) ([]*ent.MedicalHistories, int, error) {
	baseQuery := p.client.MedicalHistories.Query().Where(medicalhistories.PatientID(uuid.MustParse(request.PatientId)))

	baseQuery.WithMedicalSurgery().WithMedicalPrescription(
		func(query *ent.MedicalPrescriptionQuery) {
			query.WithPrescriptionMedication()
		},
	).WithMedicalTreatment()

	if request.CreatedBy != "" {
		baseQuery.Where(medicalhistories.CreatedBy(uuid.MustParse(request.CreatedBy)))
	}

	if request.DateRange.FromDate != 0 {
		fromDate := time.UnixMilli(request.DateRange.FromDate)
		baseQuery.Where(medicalhistories.CreatedAtGTE(fromDate))
	}

	if request.DateRange.ToDate != 0 {
		toDate := time.UnixMilli(request.DateRange.ToDate)
		baseQuery.Where(medicalhistories.CreatedAtLTE(toDate))
	}

	pagination := ToPagination(request.Pageable)

	var records []*ent.MedicalHistories
	var total int
	var err error

	if pagination.PagingIgnored {
		records, err = baseQuery.Order(ToOrderSpecifier(pagination.Sort, medicalhistories.FieldCreatedAt, Descending)).All(ctx)

		if err != nil {
			return []*ent.MedicalHistories{}, 0, err
		}

		return records, len(records), err
	}

	total, err = baseQuery.Clone().Count(ctx)

	if err != nil {
		return []*ent.MedicalHistories{}, 0, err
	}

	records, err = baseQuery.Clone().
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		Order(ToOrderSpecifier(pagination.Sort, medicalhistories.FieldCreatedAt, Descending)).
		All(ctx)

	if err != nil {
		return []*ent.MedicalHistories{}, 0, err
	}
	return records, total, nil
}

func (p medicalhistoriesRepository) ExistMedicalHistoryById(ctx context.Context, id uuid.UUID) (bool, error) {
	return p.client.MedicalHistories.Query().Where(medicalhistories.ID(id)).Exist(ctx)
}

func (p medicalhistoriesRepository) UpsertMedicalRecord(ctx context.Context, tx *ent.Tx, request *patient2.UpsertMedicalRecordRequest) error {
	var medicalEndDate *time.Time
	if request.MedicalEndDate != 0 {
		t := time.UnixMilli(request.MedicalEndDate)
		medicalEndDate = &t
	} else {
		medicalEndDate = nil
	}

	employeeId, err := grpcutil.GetEmployeeId(ctx)

	if err != nil {
		return err
	}

	if request.Id != "" {
		return tx.MedicalHistories.UpdateOneID(uuid.MustParse(request.Id)).
			SetReason(request.Reason).
			SetDiagnosis(request.Diagnosis).
			SetDoctorNotes(request.DoctorNotes).
			SetNillableMedicalEndDate(medicalEndDate).
			SetUpdatedBy(uuid.MustParse(employeeId)).
			Exec(ctx)
	}

	return tx.MedicalHistories.Create().
		SetPatientID(uuid.MustParse(request.PatientId)).
		SetReason(request.Reason).
		SetDiagnosis(request.Diagnosis).
		SetDoctorNotes(request.DoctorNotes).
		SetNillableMedicalEndDate(medicalEndDate).
		SetCreatedBy(uuid.MustParse(employeeId)).
		SetUpdatedBy(uuid.MustParse(employeeId)).
		Exec(ctx)
}

func NewMedicalHistoriesRepository(client *ent.Client, db *db.DB) MedicalHistoriesRepository {
	return &medicalhistoriesRepository{client: client, db: db}
}
