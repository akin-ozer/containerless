package main

import (
	"github.com/akin-ozer/containerless/cmd/check"
	"github.com/akin-ozer/containerless/cmd/delete"
	"github.com/akin-ozer/containerless/cmd/deploy"
	"github.com/akin-ozer/containerless/cmd/get"
	"github.com/akin-ozer/containerless/cmd/install"
	"github.com/akin-ozer/containerless/cmd/purge"
	"github.com/spf13/cobra"
)

func main() {

	var cmdDeploy = &cobra.Command{
		Use:   "deploy [image to deploy] [image name]",
		Short: "Deploy app to cluster",
		Long:  `Deploy docker image to cluster as serverless application, output will be URL of the application.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			deploy.Deploy(args[0], args[1])
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

	var cmdCheck = &cobra.Command{
		Use:   "check [noargs]",
		Short: "checks environment dependencies",
		Long: `Checks environment  with docker running.
Required binaries: kubectl, kind.`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			check.Check()
		},
	}

	var cmdPurge = &cobra.Command{
		Use:   "purge [noargs]",
		Short: "purge environment",
		Long:  `Purge environment with deleting kind cluster.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			purge.Purge()
		},
	}

	var cmdDelete = &cobra.Command{
		Use:   "delete [deployment name]",
		Short: "delete deployment",
		Long:  `Delete previously deployed service.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			delete.Delete(args[0])
		},
	}

	var cmdGet = &cobra.Command{
		Use:   "get [noargs]",
		Short: "get deployments",
		Long:  `Get all deployments.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			get.Get()
		},
	}

	//adding subcommands and flags
	//cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	//cmdInstall.AddCommand(cmdTimes)
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdDeploy, cmdInstall, cmdPurge, cmdCheck, cmdGet, cmdDelete)

	rootCmd.Execute()
}
