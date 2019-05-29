package cmd

import (
	"github.com/uber/aresdb/cmd/aresd/cmd"
	"github.com/uber/aresdb/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uber/aresdb/broker/config"
	"github.com/uber/aresdb/utils"
	"github.com/uber/aresdb/gateway"
	"time"
	"github.com/uber/aresdb/metastore"
)

func Execute(setters ...cmd.Option) {
	loggerFactory := common.NewLoggerFactory()
	options := &cmd.Options{
		ServerLogger: loggerFactory.GetDefaultLogger(),
		QueryLogger:  loggerFactory.GetLogger("query"),
		Metrics:      common.NewNoopMetrics(),
	}

	for _, setter := range setters {
		setter(options)
	}

	cmd := &cobra.Command{
		Use:     "aresbroker",
		Short:   "AresDB broker",
		Long:    `AresDB broker is the gateway to send queries to AresDB`,
		Example: `./ares --config config/ares-broker.yaml --port 9374`,
		Run: func(cmd *cobra.Command, args []string) {

			cfg, err := ReadConfig(options.DefaultCfg, cmd.Flags())
			if err != nil {
				options.ServerLogger.With("err", err.Error()).Fatal("failed to read configs")
			}

			start(
				cfg,
				options.ServerLogger,
				options.QueryLogger,
				options.Metrics,
				options.HTTPWrappers...,
			)
		},
	}
	AddFlags(cmd)
	cmd.Execute()
}

func start(cfg config.BrokerConfig, logger common.Logger, queryLogger common.Logger, metricsCfg common.Metrics, httpWrappers ...utils.HTTPHandlerWrapper) {
	logger.With("config", cfg).Info("Starting aresdb broker service")

	scope, closer, err := metricsCfg.NewRootScope()
	if err != nil {
		logger.Fatal("Failed to create new root scope", err)
	}
	defer closer.Close()

	// Init common components.
	utils.Init(common.AresServerConfig{}, logger, queryLogger, scope)

	scope.Counter("restart").Inc(1)
	serverRestartTimer := scope.Timer("restart").Start()
	defer serverRestartTimer.Stop()

	controllerClientCfg := cfg.ControllerConfig
	if controllerClientCfg == nil {
		logger.Fatal("Missing controller client config", err)
	}

	controllerClient := gateway.NewControllerHTTPClient(controllerClientCfg.Address, time.Duration(controllerClientCfg.TimeoutSec)*time.Second, controllerClientCfg.Headers)
	controllerClient.FetchAllSchemas()
	schemaFetchJob := metastore.NewSchemaFetchJob(5*60, metaStore, metastore.NewTableSchameValidator(), controllerClient, cfg.Cluster.ClusterName, "")
	// immediate initial fetch
	schemaFetchJob.FetchSchema()
	go schemaFetchJob.Run()
}

// AddFlags adds flags to command
func AddFlags(cmd *cobra.Command) {
	cmd.Flags().String("config", "config/ares-broker.yaml", "Ares broker config file")
	cmd.Flags().IntP("port", "p", 0, "Ares broker service port")
}

// ReadConfig populates BrokerConfig
func ReadConfig(defaultConfig map[string]interface{}, flags *pflag.FlagSet) (config.BrokerConfig, error) {

}