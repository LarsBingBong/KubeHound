package main

import (
	"context"

	"github.com/DataDog/KubeHound/pkg/config"
	"github.com/DataDog/KubeHound/pkg/kubehound/core"
	"github.com/spf13/cobra"
)

// TODO config file path

var (
	cfgFile = config.DefaultConfigPath
)

var (
	rootCmd = &cobra.Command{
		Use:   "kubehound-local",
		Short: "A local Kubehound instance",
		Long:  `A local instance of Kubehound - a Kubernetes attack path generator`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return core.Launch(context.Background(), core.WithConfigPath(cfgFile))
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, "application config file")
}
