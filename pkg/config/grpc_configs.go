package config

import "github.com/spf13/viper"

type ServiceGrpcPath struct {
	Host string
	Port int
}

type GRPC_Services struct {
	EnableGRPCCall bool
	// ExampleService ServiceGrpcPath
}

var grpc_services GRPC_Services

func GRPCServices() GRPC_Services {
	return grpc_services
}

func loadGRPCServices() {
	grpc_services = GRPC_Services{
		EnableGRPCCall: viper.GetBool("grpc_services.enable_grpc_call"),
		/* ExampleService: ServiceGrpcPath{
			Host: viper.GetString("grpc_services.example.host"),
			Port: viper.GetInt("grpc_services.example.port"),
		}, */
	}
}
