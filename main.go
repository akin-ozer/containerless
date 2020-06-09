package main

import (
	"github.com/akin-ozer/containerless/cmd/deploy"
	"github.com/akin-ozer/containerless/cmd/install"
	"github.com/akin-ozer/containerless/cmd/remove"
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

	var cmdRemove = &cobra.Command{
		Use:   "remove [noargs]",
		Short: "remove environment",
		Long:  `Remove environment to machine with deleting kind cluster.`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			remove.Remove()
		},
	}

	//adding subcommands and flags
	//cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	//cmdInstall.AddCommand(cmdTimes)
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdDeploy, cmdInstall, cmdRemove)

	rootCmd.Execute()
}
