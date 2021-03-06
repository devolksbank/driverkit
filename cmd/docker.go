package cmd

import (
	"github.com/falcosecurity/driverkit/pkg/driverbuilder"
	"github.com/sirupsen/logrus"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// NewDockerCmd creates the `driverkit docker` command.
func NewDockerCmd(rootOpts *RootOptions, rootFlags *pflag.FlagSet) *cobra.Command {
	dockerCmd := &cobra.Command{
		Use:   "docker",
		Short: "Build Falco kernel modules and eBPF probes against a docker daemon.",
		Run: func(c *cobra.Command, args []string) {
			f := c.Flags()
			image, err := f.GetString("image")
			if err != nil {
				logger.WithError(err).Fatal("exiting")
			}
			logrus.WithField("processor", c.Name()).Info("driver building, it will take a few seconds")
			if !configOptions.DryRun {
				if err := driverbuilder.NewDockerBuildProcessor(viper.GetInt("timeout"), viper.GetString("proxy"), image).Start(rootOpts.toBuild()); err != nil {
					logger.WithError(err).Fatal("exiting")
				}
			}
		},
	}
	// Add root flags
	dockerCmd.PersistentFlags().AddFlagSet(rootFlags)
	dockerCmd.PersistentFlags().StringVarP(&Image, "image", "i", "", "image to use")
	return dockerCmd
}
