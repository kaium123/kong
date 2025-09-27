package conn

// examplePb "vcs.technonext.com/carrybee/photon/example"

type GRPCClients struct {
	//Example Service
	// ExampleStoreServiceClient            examplePb.StoreServiceClient
}

var (
// Example Service
// exampleStoreServiceClient            examplePb.StoreServiceClient
)

func ConnectGrpcClient() error {
	/*
		grpcServiceConfig := config.GRPCServices()

		// Example Service
		exampleRpc, err := grpc.NewClient(
			fmt.Sprintf("%s:%d", grpcServiceConfig.ExampleService.Host, grpcServiceConfig.ExampleService.Port),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed to connect Example Micro gRPC server, reason: %v", err))
		} else {
			fmt.Println("Connected to Example Micro gRPC Server")
			exampleStoreServiceClient = examplePb.NewStoreServiceClient(exampleRpc)
		}

		return err
	*/
	return nil
}

func GrpcClients() *GRPCClients {
	return &GRPCClients{
		// ExampleStoreServiceClient:            ExampleStoreServiceClient,
	}
}
