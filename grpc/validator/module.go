package grpcvalidator

import "go.uber.org/fx"

var Module = fx.Provide(NewPatientGrpcValidator)
