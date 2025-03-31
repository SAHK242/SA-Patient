package grpc

import (
	mapper "patient/grpc/mapper"
	"patient/grpc/server"
	service "patient/grpc/service"
	"patient/interceptor"

	"go.uber.org/fx"

	validator "patient/grpc/validator"

	"patient/grpc/controller"
)

var Module = fx.Module("grpc",
	interceptor.Module,
	mapper.Module,
	validator.Module,
	service.Module,
	controller.Module,
	server.Module,
)
