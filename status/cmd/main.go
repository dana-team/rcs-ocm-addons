package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	goflag "flag"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
	"open-cluster-management.io/addon-framework/pkg/version"

	"github.com/dana-team/rcs-ocm-addons/status/statushub"
	"github.com/dana-team/rcs-ocm-addons/status/statusspoke"
)

func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	var logger logr.Logger

	zapLog, err := zap.NewDevelopment()
	if err != nil {
		if _, er := fmt.Fprintf(os.Stderr, "failed to start up logger %v\n", err); er != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}

	logger = zapr.NewLogger(zapLog)

	command := newCommand(logger)
	if err := command.Execute(); err != nil {
		if _, er := fmt.Fprintf(os.Stderr, "%v\n", err); er != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}
}

// newCommand creates a new Cobra command for the status-addon CLI.
// It initializes the root command with the appropriate metadata and subcommands.
func newCommand(logger logr.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status-addon",
		Short: "status-addon",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				if _, er := fmt.Fprintf(os.Stderr, "%v\n", err); er != nil {
					os.Exit(1)
				}
			}
			os.Exit(1)
		},
	}

	if v := version.Get().String(); len(v) == 0 {
		cmd.Version = "<unknown>"
	} else {
		cmd.Version = v
	}

	cmd.AddCommand(statushub.NewManagerCommand("capp-status-addon", logger.WithName("capp-status-manager")))
	cmd.AddCommand(statusspoke.NewAgentCommand("capp-status-addon", logger.WithName("capp-status-agent")))

	return cmd
}
