package server

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"patient/config"

	"patient/grpc/controller"
	"patient/proto/patient"
)

var Module = fx.Provide(NewGrpcServer)

type GrpcServerProps struct {
	fx.In
	*controller.PatientGrpcController
	*controller.MedicationGrpcController
	config.Config
	*zap.SugaredLogger
	ChainedInterceptor grpc.ServerOption `name:"chainedGrpcInterceptor"`
}

func NewGrpcServer(lifecycle fx.Lifecycle, props GrpcServerProps) *grpc.Server {
	grpcServer := grpc.NewServer(props.ChainedInterceptor)

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", fmt.Sprintf(":%d", props.Config.MustGetInt("GRPC_PORT")))

			if err != nil {
				panic(err)
			}

			patient.RegisterPatientServiceServer(grpcServer, props.PatientGrpcController)
			patient.RegisterMedicationServiceServer(grpcServer, props.MedicationGrpcController)

			go func() {
				props.SugaredLogger.Infof("gRPC server is running on %s", lis.Addr().String())

				if e := grpcServer.Serve(lis); e != nil {
					panic(e)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			props.SugaredLogger.Info("gRPC server is shutting down...")
			grpcServer.GracefulStop()
			props.SugaredLogger.Info("gRPC server is shutdown gracefully.")
			return nil
		},
	})

	return grpcServer
}
