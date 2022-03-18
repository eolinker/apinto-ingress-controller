package ingress

import (
	"fmt"
	"github.com/eolinker/apinto-ingress-controller/pkg/config"
	"github.com/eolinker/apinto-ingress-controller/pkg/ingress"
	"github.com/eolinker/apinto-ingress-controller/pkg/ingress/signals"
	"github.com/eolinker/eosc/log"

	"github.com/spf13/cobra"
	"os"
	"strings"
	"sync"
)

func dief(template string, args ...interface{}) {
	if !strings.HasSuffix(template, "\n") {
		template += "\n"
	}
	fmt.Fprintf(os.Stderr, template, args...)
	os.Exit(1)
}

// NewIngressCommand creates the ingress sub command for apinto-ingress-controller.
func NewIngressCommand() *cobra.Command {
	var configPath string
	cfg := config.NewDefaultConfig()

	cmd := &cobra.Command{
		Use: "ingress [flags]",
		Long: `launch the ingress controller
			You can run apinto-ingress-controller from configuration file or command line options,
			if you run it from configuration file, other command line options will be ignored.
			
			Run from configuration file:
			
				apinto-ingress-controller ingress --config-path /path/to/config.json
			
			Both json and yaml are supported as the configuration file format.
			
			Run from command line options:
			
				apinto-ingress-controller ingress --default-apinto-cluster-base-url http://apinto-service:9400/api/
			
			Before you run apinto-ingress-controller, be sure all related resources, like CRDs (ApisixRoute, ApisixUpstream and etc),
			the apinto cluster and others are created`,
		Run: func(cmd *cobra.Command, args []string) {
			if configPath != "" {
				c, err := config.NewConfigFromFile(configPath)
				if err != nil {
					dief("failed to initialize configuration: %s", err)
				}
				cfg = c
			}
			// 配置格式校验
			if err := cfg.Validate(); err != nil {
				dief("bad configuration: %s", err)
			}
			InitLogger(cfg.Log)
			log.Info("apinto ingress controller started")

			stop := signals.SetupSignalHandler()
			client, err := ingress.NewController(cfg)
			if err != nil {
				dief("failed to create ingress controller: %s", err)
			}
			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := client.Run(stop); err != nil {
					dief("failed to run ingress controller: %s", err)
				}
			}()
			wg.Wait()
			log.Info("apinto ingress controller exited")
			log.Close()
		},
	}

	cmd.PersistentFlags().StringVar(&configPath, "config-path", "/etc/ingress/config.yaml", "configuration file path for apinto-ingress-controller")
	cmd.PersistentFlags().StringVar(&cfg.Log.LogLevel, "log-level", "info", "error log level")
	cmd.PersistentFlags().StringVar(&cfg.Log.LogOutput, "log-output", "stderr", "error log output file")
	cmd.PersistentFlags().StringVar(&cfg.HTTPListen, "http-listen", ":8080", "the HTTP Server listen address")
	cmd.PersistentFlags().StringVar(&cfg.HTTPSListen, "https-listen", ":8443", "the HTTPS Server listen address")
	cmd.PersistentFlags().BoolVar(&cfg.EnableProfiling, "enable-profiling", true, "enable profiling via web interface host:port/debug/pprof")
	cmd.PersistentFlags().StringVar(&cfg.APINTO.DefaultClusterBaseURL, "default-apinto-cluster-base-url", "", "the base URL of admin api / manager api for the default APISIX cluster")
	cmd.PersistentFlags().StringVar(&cfg.APINTO.DefaultClusterAdminKey, "default-apinto-cluster-admin-key", "", "admin key used for the authorization of admin api / manager api for the default APISIX cluster")
	cmd.PersistentFlags().StringVar(&cfg.APINTO.DefaultClusterName, "default-apinto-cluster-name", "default", "name of the default apinto cluster")

	return cmd
}
