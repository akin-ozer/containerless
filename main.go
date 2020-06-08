package main

import (
	"github.com/akin-ozer/containerless/cmd/deploy"
	"github.com/akin-ozer/containerless/cmd/install"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func main() {
	var echoTimes int

	var cmdDeploy = &cobra.Command{
		Use:   "deploy [image to deploy]",
		Short: "Deploy app to cluster",
		Long:  `Deploy docker image to cluster as serverless application, output will be URL of the application.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			deploy.Deploy(strings.Join(args, " "))
			//fmt.Println("Print: " + strings.Join(args, " ")docker.io/realvega/quarkus-hello)
		},
	}

	var cmdInstall = &cobra.Command{
		Use:   "install [noargs]",
		Short: "install environment",
		Long: `Installs environment to machine with docker running.
Required binaries: kubectl, kind.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			install.Install()
		},
	}

	var cmdTimes = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "Echo anything to the screen more times",
		Long: `echo things multiple times back to the user by providing
a count and a string.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}

	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdDeploy, cmdInstall)
	cmdInstall.AddCommand(cmdTimes)
	rootCmd.Execute()
}
