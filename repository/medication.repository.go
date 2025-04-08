package repository

import (
	"context"
	"github.com/google/uuid"
	"patient/ent"
	"patient/ent/medication"
	grpcutil "patient/grpc/util"
	patient2 "patient/proto/patient"
	"time"
)
import db "database/sql"

type medicationRepository struct {
	client *ent.Client
	db     *db.DB
}

func (p medicationRepository) FindAllMedication(ctx context.Context, request *patient2.ListMedicationRequest) ([]*ent.Medication, int, error) {
	baseQuery := p.client.Medication.Query()

	if request.Search != "" {
		baseQuery.Where(medication.NameContainsFold(request.Search))
	}

	pagination := ToPagination(request.Pageable)

	var records []*ent.Medication
	var total int
	var err error

	if pagination.PagingIgnored {
		records, err = baseQuery.Order(ToOrderSpecifier(pagination.Sort, medication.FieldName, Ascending)).All(ctx)

		if err != nil {
			return []*ent.Medication{}, 0, err
		}

		return records, len(records), err
	}

	total, err = baseQuery.Clone().Count(ctx)

	if err != nil {
		return []*ent.Medication{}, 0, err
	}

	records, err = baseQuery.Clone().
		Limit(pagination.Limit).
		Offset(pagination.Offset).
		Order(ToOrderSpecifier(pagination.Sort, medication.FieldName, Ascending)).
		All(ctx)

	if err != nil {
		return []*ent.Medication{}, 0, err
	}
	return records, total, nil
}

func (p medicationRepository) ExistMedicationByIdIn(ctx context.Context, ids []uuid.UUID) (bool, error) {
	if len(ids) == 0 {
		return false, nil
	}

	count, err := p.client.Medication.Query().
		Where(medication.IDIn(ids...)).
		Count(ctx)

	if err != nil {
		return false, err
	}

	return count == len(ids), nil
}

func (p medicationRepository) UpsertMedication(ctx context.Context, tx *ent.Tx, request *patient2.UpsertMedicationRequest) error {
	employeeId, err := grpcutil.GetEmployeeId(ctx)
	if err != nil {
		return err
	}
	if request.Id != "" {
		return tx.Medication.UpdateOneID(uuid.MustParse(request.Id)).SetName(request.Name).SetPrice(float64(request.Price)).SetEffects(request.Effects).SetExpiredDate(time.UnixMilli(request.ExpiredDate)).SetQuantity(request.Quantity).SetUpdatedBy(uuid.MustParse(employeeId)).Exec(ctx)
	}

	return tx.Medication.Create().SetName(request.Name).SetPrice(float64(request.Price)).SetEffects(request.Effects).SetExpiredDate(time.UnixMilli(request.ExpiredDate)).SetQuantity(request.Quantity).SetUpdatedBy(uuid.MustParse(employeeId)).SetCreatedBy(uuid.MustParse(employeeId)).Exec(ctx)
}

func NewMedicationRepository(client *ent.Client, db *db.DB) MedicationRepository {
	return &medicationRepository{client: client, db: db}
}
