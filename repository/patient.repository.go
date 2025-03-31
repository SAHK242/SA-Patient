package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"patient/ent"
	"patient/ent/patient"
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
	if request.Id != "" {
		return tx.Patient.UpdateOneID(uuid.MustParse(request.Id)).
			SetAddress(request.Address).
			SetFirstName(request.FirstName).
			SetLastName(request.LastName).
			SetPhoneNumber(request.PhoneNumber).
			SetDateOfBirth(time.UnixMilli(request.DateOfBirth)).
			SetGender(int32(request.Gender)).Exec(ctx)
	} else {
		return tx.Patient.Create().
			SetAddress(request.Address).
			SetFirstName(request.FirstName).
			SetLastName(request.LastName).
			SetPhoneNumber(request.PhoneNumber).
			SetDateOfBirth(time.UnixMilli(request.DateOfBirth)).
			SetGender(int32(request.Gender)).Exec(ctx)
	}
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
