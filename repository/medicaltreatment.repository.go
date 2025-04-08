package repository

import (
	"context"
	"github.com/google/uuid"
	"patient/ent"
	grpcutil "patient/grpc/util"
	patient2 "patient/proto/patient"
	"time"
)
import db "database/sql"

type treatmentRepository struct {
	client *ent.Client
	db     *db.DB
}

func (p treatmentRepository) UpsertMedicalTreatment(ctx context.Context, tx *ent.Tx, request *patient2.UpsertMedicalTreatmentRequest) error {
	startDate := time.UnixMilli(request.StartDate)
	employeeId, err := grpcutil.GetEmployeeId(ctx)
	if err != nil {
		return err
	}

	if request.Id != "" {
		if request.EndDate != 0 {
			endDate := time.UnixMilli(request.EndDate)
			return tx.MedicalTreatment.UpdateOneID(uuid.MustParse(request.Id)).
				SetMedicalHistoryID(uuid.MustParse(request.MedicalHistoryId)).
				SetName(request.Name).
				SetStartDate(startDate).
				SetEndDate(endDate).
				SetMainDoctorID(uuid.MustParse(request.MainDoctorId)).
				SetResult(request.Result).
				SetDescription(request.Description).
				SetFee(request.Fee).
				SetSupportDoctorIds(request.SupportDoctorIds).
				SetSupportNurseIds(request.SupportNurseIds).SetUpdatedBy(uuid.MustParse(employeeId)).Exec(ctx)
		}
	}

	cre := tx.MedicalTreatment.Create().
		SetMedicalHistoryID(uuid.MustParse(request.MedicalHistoryId)).
		SetName(request.Name).
		SetStartDate(startDate).
		SetMainDoctorID(uuid.MustParse(request.MainDoctorId)).
		SetResult(request.Result).
		SetDescription(request.Description).
		SetFee(request.Fee).
		SetSupportDoctorIds(request.SupportDoctorIds).
		SetSupportNurseIds(request.SupportNurseIds).SetCreatedBy(uuid.MustParse(employeeId)).SetUpdatedBy(uuid.MustParse(employeeId))

	if request.EndDate != 0 {
		endDate := time.UnixMilli(request.EndDate)
		cre.SetEndDate(endDate)
	}

	return cre.Exec(ctx)
}

func NewMedicalTreatmentRepository(client *ent.Client, db *db.DB) MedicalTreatmentRepository {
	return &treatmentRepository{client: client, db: db}
}
