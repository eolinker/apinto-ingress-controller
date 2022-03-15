package cmd

import (
	"github.com/eolinker/apinto-ingress-controller/cmd/ingress"
	"github.com/spf13/cobra"
)

func NewIngressControllerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "apinto-ingress-controller [command]",
	}
	cmd.AddCommand(ingress.NewIngressCommand())
	return cmd
}
