package config

import (
	"time"

	"github.com/spf13/viper"
)

// GRpcApplication represents grpc application config
type GRpcApplication struct {
	GRpcPort int

	MaxConcurrentStreams uint32
	MaxConnectionIdle    time.Duration
	Time                 time.Duration
	Timeout              time.Duration

	MaxRecvMsgSize int
	MaxSendMsgSize int

	DefaultPaginationLimit int
	MaxPaginationLimit     int
}

// AppConfig represents general app configuration
type AppConfig struct {
	Name     string
	Version  string
	Env      string
	LogLevel string
}

var grpc_app GRpcApplication
var appConfig AppConfig

// GRpcApp contains grpc app configurations
func GRpcApp() GRpcApplication {
	return grpc_app
}

// App returns general app configuration
func App() AppConfig {
	return appConfig
}

func loadApp() {
	appConfig = AppConfig{
		Name:     viper.GetString("app.name"),
		Version:  viper.GetString("app.version"),
		Env:      viper.GetString("app.env"),
		LogLevel: viper.GetString("app.log_level"),
	}

	grpc_app = GRpcApplication{
		GRpcPort:             viper.GetInt("grpc_app.grpc_port"),
		MaxConcurrentStreams: viper.GetUint32("grpc_app.max_concurrent_Streams"),

		MaxConnectionIdle: viper.GetDuration("grpc_app.max_connection_idle") * time.Second,
		Time:              viper.GetDuration("grpc_app.time") * time.Second,
		Timeout:           viper.GetDuration("grpc_app.timeout") * time.Second,

		MaxRecvMsgSize: viper.GetInt("grpc_app.max_recv_msg_size"),
		MaxSendMsgSize: viper.GetInt("grpc_app.max_send_msg_size"),

		MaxPaginationLimit:     viper.GetInt("grpc_app.max_pagination_limit"),
		DefaultPaginationLimit: viper.GetInt("grpc_app.default_pagination_limit"),
	}
}
