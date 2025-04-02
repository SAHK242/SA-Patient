package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"patient/ent"
	"patient/ent/patient"
	"patient/proto/gcommon"
	patient2 "patient/proto/patient"
	"time"
)
import db "database/sql"

type patientRepository struct {
	client *ent.Client
	db     *db.DB
}

func (p patientRepository) FindById(ctx context.Context, id uuid.UUID) (*ent.Patient, error) {
	return p.client.Patient.Get(ctx, id)
}

func (p patientRepository) UpsertPatient(ctx context.Context, tx *ent.Tx, request *patient2.UpsertPatientRequest) error {
	var err error
	var patientCre *ent.Patient

	if request.Id != "" {
		patientCre, err = tx.Patient.UpdateOneID(uuid.MustParse(request.Id)).
			SetAddress(request.Address).
			SetFirstName(request.FirstName).
			SetLastName(request.LastName).
			SetPhoneNumber(request.PhoneNumber).
			SetDateOfBirth(time.UnixMilli(request.DateOfBirth)).
			SetCurrentPatientType(int32(request.PatientType)).
			SetGender(int32(request.Gender)).Save(ctx)
	} else {
		patientCre, err = tx.Patient.Create().
			SetAddress(request.Address).
			SetFirstName(request.FirstName).
			SetLastName(request.LastName).
			SetPhoneNumber(request.PhoneNumber).
			SetDateOfBirth(time.UnixMilli(request.DateOfBirth)).
			SetCurrentPatientType(int32(request.PatientType)).
			SetGender(int32(request.Gender)).Save(ctx)
	}

	if err != nil {
		return err
	}

	//Not error, insert inpatient/outpatient
	if request.PatientType == gcommon.PatientType_PATIENT_TYPE_INPATIENT {
		return tx.Inpatient.Create().
			SetPatientID(patientCre.ID).
			SetRegisterDate(time.Now()).
			Exec(ctx)
	}

	if request.PatientType == gcommon.PatientType_PATIENT_TYPE_OUTPATIENT {
		return tx.Outpatient.Create().
			SetPatientID(patientCre.ID).
			SetRegisterDate(time.Now()).
			Exec(ctx)
	}

	return nil
}

func (p patientRepository) FindAllPatient(ctx context.Context, req *patient2.ListPatientRequest) ([]*ent.Patient, int, error) {
	baseQuery := p.client.Patient.Query()

	if req.Search != "" {
		baseQuery.Where(func(selector *sql.Selector) {
			enNamePredicate := sql.ExprP("concat(first_name,' ',last_name) ilike " + "'%" + req.Search + "%'")
			viNamePredicate := sql.ExprP("concat(last_name,' ',first_name) ilike " + "'%" + req.Search + "%'")
			selector.Where(sql.Or(enNamePredicate, viNamePredicate))
		})
	}

	pagination := ToPagination(req.Pageable)

	var records []*ent.Patient
	var total int
	var err error

	if pagination.PagingIgnored {
		records, err = baseQuery.Order(ToOrderSpecifier(pagination.Sort, patient.FieldFirstName, Ascending)).All(ctx)

		if err != nil {
			return []*ent.Patient{}, 0, err
		}

		return records, len(records), err
	}

	total, err = baseQuery.Clone().Count(ctx)

	if err != nil {
		return []*ent.Patient{}, 0, err
	}

	records, err = baseQuery.Clone().
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		Order(ToOrderSpecifier(pagination.Sort, patient.FieldFirstName, Ascending)).
		All(ctx)

	if err != nil {
		return []*ent.Patient{}, 0, err
	}
	return records, total, nil

}

func NewPatientRepository(client *ent.Client, db *db.DB) PatientRepository {
	return &patientRepository{client: client, db: db}
}
