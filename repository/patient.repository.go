package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/google/uuid"
	"patient/ent"
	"patient/ent/medicalhistories"
	"patient/ent/medication"
	"patient/ent/patient"
	grpcutil "patient/grpc/util"
	patient2 "patient/proto/patient"
	"time"
)
import db "database/sql"

type patientRepository struct {
	client *ent.Client
	db     *db.DB
}

func (p patientRepository) FindMedicalHistoryById(ctx context.Context, id uuid.UUID) (*ent.MedicalHistories, error) {
	return p.client.MedicalHistories.Query().Where(medicalhistories.ID(id)).WithMedicalSurgery().WithMedicalPrescription(
		func(query *ent.MedicalPrescriptionQuery) {
			query.WithPrescriptionMedication(
				func(query *ent.PrescriptionMedicationQuery) {
					query.WithMedication()
				})
		},
	).WithMedicalTreatment().Only(ctx)
}

func (p patientRepository) FindMedicalHistories(ctx context.Context, request *patient2.GetMedicalHistoryRequest) ([]*ent.MedicalHistories, int, error) {
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

func (p patientRepository) FindAllMedication(ctx context.Context, request *patient2.ListMedicationRequest) ([]*ent.Medication, int, error) {
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

func (p patientRepository) UpsertMedicalSurgery(ctx context.Context, tx *ent.Tx, request *patient2.UpsertMedicalSurgeryRequest) error {
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
				SetFee(float64(request.Fee)).
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
		SetFee(float64(request.Fee)).
		SetSupportDoctorIds(request.SupportDoctorIds).
		SetSupportNurseIds(request.SupportNurseIds).SetCreatedBy(uuid.MustParse(employeeId)).SetUpdatedBy(uuid.MustParse(employeeId))

	if request.EndDate != 0 {
		endDate := time.UnixMilli(request.EndDate)
		cre.SetEndDate(endDate)
	}

	return cre.Exec(ctx)
}

func (p patientRepository) UpsertMedicalPrescription(ctx context.Context, tx *ent.Tx, request *patient2.UpsertMedicalPrescriptionRequest) error {
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

func (p patientRepository) ExistMedicationByIdIn(ctx context.Context, ids []uuid.UUID) (bool, error) {
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

func (p patientRepository) ExistMedicalHistoryById(ctx context.Context, id uuid.UUID) (bool, error) {
	return p.client.MedicalHistories.Query().Where(medicalhistories.ID(id)).Exist(ctx)
}

func (p patientRepository) UpsertMedicalRecord(ctx context.Context, tx *ent.Tx, request *patient2.UpsertMedicalRecordRequest) error {
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

func (p patientRepository) ExistByPatientId(ctx context.Context, id uuid.UUID) (bool, error) {
	exist, err := p.client.Patient.Query().Where(patient.ID(id)).Exist(ctx)

	if err != nil {
		return false, err
	}

	return exist, nil
}

func NewPatientRepository(client *ent.Client, db *db.DB) PatientRepository {
	return &patientRepository{client: client, db: db}
}

func (p patientRepository) UpsertMedication(ctx context.Context, tx *ent.Tx, request *patient2.UpsertMedicationRequest) error {
	employeeId, err := grpcutil.GetEmployeeId(ctx)
	if err != nil {
		return err
	}
	if request.Id != "" {
		return tx.Medication.UpdateOneID(uuid.MustParse(request.Id)).SetName(request.Name).SetPrice(float64(request.Price)).SetEffects(request.Effects).SetExpiredDate(time.UnixMilli(request.ExpiredDate)).SetQuantity(request.Quantity).SetUpdatedBy(uuid.MustParse(employeeId)).Exec(ctx)
	}

	return tx.Medication.Create().SetName(request.Name).SetPrice(float64(request.Price)).SetEffects(request.Effects).SetExpiredDate(time.UnixMilli(request.ExpiredDate)).SetQuantity(request.Quantity).SetUpdatedBy(uuid.MustParse(employeeId)).SetCreatedBy(uuid.MustParse(employeeId)).Exec(ctx)
}

func (p patientRepository) FindById(ctx context.Context, id uuid.UUID) (*ent.Patient, error) {
	return p.client.Patient.Get(ctx, id)
}

func (p patientRepository) UpsertPatient(ctx context.Context, tx *ent.Tx, request *patient2.UpsertPatientRequest) error {
	var err error
	employeeId, e := grpcutil.GetEmployeeId(ctx)

	if e != nil {
		return e
	}

	if request.Id != "" {
		err = tx.Patient.UpdateOneID(uuid.MustParse(request.Id)).
			SetAddress(request.Address).
			SetFirstName(request.FirstName).
			SetLastName(request.LastName).
			SetPhone(request.Phone).
			SetDateOfBirth(time.UnixMilli(request.DateOfBirth)).
			SetGender(int32(request.Gender)).
			SetUpdatedBy(uuid.MustParse(employeeId)).
			Exec(ctx)
	} else {
		err = tx.Patient.Create().
			SetAddress(request.Address).
			SetFirstName(request.FirstName).
			SetLastName(request.LastName).
			SetPhone(request.Phone).
			SetDateOfBirth(time.UnixMilli(request.DateOfBirth)).
			SetCreatedBy(uuid.MustParse(employeeId)).
			SetUpdatedBy(uuid.MustParse(employeeId)).
			SetGender(int32(request.Gender)).Exec(ctx)
	}

	return err
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
