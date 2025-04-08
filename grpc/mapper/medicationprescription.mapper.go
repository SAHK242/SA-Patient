package grpcmapper

import (
	"patient/ent"
	"patient/proto/patient"
)

type medicalPrescriptionGrpcMapper struct{}

func (p medicalPrescriptionGrpcMapper) ConvertMedicalPrescription(from *ent.MedicalPrescription) *patient.MedicalPrescription {
	if from == nil {
		return nil
	}

	to := &patient.MedicalPrescription{
		Id:               from.ID.String(),
		MedicalHistoryId: from.MedicalHistoryID.String(),
		CreatedAt:        from.CreatedAt.UnixMilli(),
		CreatedBy:        &patient.User{Id: from.CreatedBy.String()},
		Fee:              from.Fee,
		PrescriptionDate: from.PrescriptionDate.UnixMilli(),
	}

	if from.Edges.PrescriptionMedication != nil {
		for _, v := range from.Edges.PrescriptionMedication {
			to.Medications = append(to.Medications, &patient.MedicalMedication{
				Id: v.ID.String(),
				Medication: &patient.Medication{
					Id:          v.MedicationID.String(),
					Name:        v.Edges.Medication.Name,
					Price:       v.Edges.Medication.Price,
					Quantity:    v.Edges.Medication.Quantity,
					ExpiredDate: v.Edges.Medication.ExpiredDate.UnixMilli(),
					Effects:     v.Edges.Medication.Effects,
				},
				Quantity: v.Quantity,
			})
		}
	}

	return to
}

func (p medicalPrescriptionGrpcMapper) ConvertMedicalPrescriptionSlice(from []*ent.MedicalPrescription) []*patient.MedicalPrescription {
	to := make([]*patient.MedicalPrescription, len(from))
	for i, v := range from {
		to[i] = p.ConvertMedicalPrescription(v)
	}
	return to
}

func NewMedicalPrescriptionGrpcMapper() MedicalPrescriptionGrpcMapper {
	return &medicalPrescriptionGrpcMapper{}
}
