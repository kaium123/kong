package cmd

import (
	"context"
	"os"

	"github.com/api_gaurd/pkg/config"
	"github.com/api_gaurd/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// cfgFile store the configuration file name
	cfgFile                 string
	secretFile              string
	verbose, prettyPrintLog bool

	// rootCmd is the root command of backup service
	rootCmd = &cobra.Command{
		Use:   "api_gaurd",
		Short: "API Gaurd is a tool for managing apis.",
		Long:  `API Gaurd is a tool for managing apis`,
	}
)

func init() {
	cobra.OnInitialize(InitConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")
	rootCmd.PersistentFlags().BoolVarP(&prettyPrintLog, "pretty", "p", false, "pretty print verbose/log")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// set the value to viper config
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ctx := context.Background()
		log.Error(ctx, err)
		os.Exit(1)
	}
}

func InitConfig() {
	ctx := context.Background()
	log.Info(ctx, "Loading configurations")
	if err := config.Init(cfgFile, secretFile); err != nil {
		log.Warn("Failed to load configuration")
		//log.Fatal(err)
	}
	log.Info(ctx, "Configurations loaded successfully!")

	// Log as JSON instead of the default ASCII formatter.
	log.SetLogFormatter(&logrus.JSONFormatter{
		PrettyPrint: prettyPrintLog,
	})

	// by defualt only log the warning severity or above.
	log.SetLogLevel(logrus.TraceLevel) // logrus.WarnLevel
	if verbose {                       // if -v flag pass override previous value
		log.SetLogLevel(logrus.TraceLevel)
	}
}
