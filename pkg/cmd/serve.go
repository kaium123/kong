package cmd

import (
	"context"
	"fmt"
	"github.com/api_gaurd/access_control/repository"
	"github.com/api_gaurd/access_control/server"
	"github.com/api_gaurd/access_control/service"
	"github.com/api_gaurd/pkg/conn"
	"github.com/api_gaurd/pkg/log"
	accesscontrolpb "github.com/api_gaurd/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve run gRPC server on defined port on env",
		Long:  `Serve run gRPC server on defined port on env`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := conn.ConnectDefaultDB(); err != nil {
				log.Fatal(err)
			} else {
				log.Info(context.Background(), "Database connected successfully!")
			}
		},
		Run: serve,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Error(ctx, fmt.Sprintf("failed to listen on :50052: %v", err))
		return // stop here to avoid panic
	}

	db := conn.DefaultDB()
	repo := repository.NewAccessControlPostgreSQL(db)
	svc := service.NewAccessControlService(repo)
	srv := server.NewAccessControlServer(svc)

	grpcServer := grpc.NewServer()
	accesscontrolpb.RegisterAccessControlServiceServer(grpcServer, &srv)

	log.Println("AccessControlService running on :50052")

	if err := grpcServer.Serve(lis); err != nil {
		log.Error(ctx, fmt.Sprintf("failed to serve gRPC: %v", err))
	}
}
