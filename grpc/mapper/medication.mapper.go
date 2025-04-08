package grpcmapper

import (
	"patient/ent"
	"patient/proto/patient"
)

type medicationGrpcMapper struct{}

func (p medicationGrpcMapper) ConvertMedication(from *ent.Medication) *patient.Medication {
	if from == nil {
		return nil
	}
	to := &patient.Medication{
		Id:          from.ID.String(),
		Name:        from.Name,
		Price:       from.Price,
		Quantity:    from.Quantity,
		ExpiredDate: from.ExpiredDate.UnixMilli(),
		Effects:     from.Effects,
	}
	return to
}

func (p medicationGrpcMapper) ConvertMedicationSlice(from []*ent.Medication) []*patient.Medication {
	to := make([]*patient.Medication, len(from))
	for i, v := range from {
		to[i] = p.ConvertMedication(v)
	}
	return to
}

func NewMedicationGrpcMapper() MedicationGrpcMapper {
	return &medicationGrpcMapper{}
}
