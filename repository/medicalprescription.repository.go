package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"patient/ent"
	"patient/ent/medication"
	grpcutil "patient/grpc/util"
	patient2 "patient/proto/patient"
	"time"
)
import db "database/sql"

type medicalprescriptionRepository struct {
	client *ent.Client
	db     *db.DB
}

func (p medicalprescriptionRepository) UpsertMedicalPrescription(ctx context.Context, tx *ent.Tx, request *patient2.UpsertMedicalPrescriptionRequest) error {
	employeeId, err := grpcutil.GetEmployeeId(ctx)
	if err != nil {
		return err
	}

	medicationIds := make([]uuid.UUID, 0)
	fee := 0.0
	for _, med := range request.Medications {
		medicationIds = append(medicationIds, uuid.MustParse(med.MedicationId))
	}

	medications, err := p.client.Medication.Query().Where(medication.IDIn(medicationIds...)).All(ctx)

	medicationById := make(map[uuid.UUID]*ent.Medication)
	for _, med := range medications {
		medicationById[med.ID] = med
	}

	for _, med := range request.Medications {
		medQ, ok := medicationById[uuid.MustParse(med.MedicationId)]
		if !ok {
			return fmt.Errorf("medication with id %s not found", medQ.ID)
		}

		fee += medQ.Price * float64(med.Quantity)
	}

	if err != nil {
		return err
	}

	x, err := tx.MedicalPrescription.Create().SetMedicalHistoryID(uuid.MustParse(request.MedicalHistoryId)).
		SetCreatedBy(uuid.MustParse(employeeId)).SetFee(fee).SetPrescriptionDate(time.Now()).Save(ctx)

	if err != nil {
		return err
	}

	bulkCre := make([]*ent.PrescriptionMedicationCreate, 0)

	for _, med := range request.Medications {
		bulkCre = append(bulkCre, tx.PrescriptionMedication.Create().SetPrescriptionID(x.ID).SetMedicationID(uuid.MustParse(med.MedicationId)).SetQuantity(med.Quantity))
	}

	if len(bulkCre) > 0 {
		if err := tx.PrescriptionMedication.CreateBulk(bulkCre...).Exec(ctx); err != nil {
			return err
		}
	}

	// Update medication quantity
	for _, med := range request.Medications {
		medQ, ok := medicationById[uuid.MustParse(med.MedicationId)]
		if !ok {
			return fmt.Errorf("medication with id %s not found", medQ.ID)
		}

		if err := p.client.Medication.UpdateOneID(medQ.ID).SetQuantity(medQ.Quantity - med.Quantity).Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}

func NewMedicalPrescriptionRepository(client *ent.Client, db *db.DB) MedicalPrescriptionRepository {
	return &medicalprescriptionRepository{client: client, db: db}
}
