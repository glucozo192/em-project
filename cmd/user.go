package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

// generalCmd represents the general command
var generalCmd = &cobra.Command{
	Use:   "general",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		loadGeneral(ctx)

		errChan := make(chan error)
		start(ctx, errChan)
		err := <-errChan
		log.Printf("Server have an error: %v, server stop now\n", err.Error())
		if err := stop(ctx); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func loadGeneral(ctx context.Context) {
	srv.loadPostgres(ctx)

	if err := srv.loadGrpcClients(ctx); err != nil {
		panic(err)
	}

	srv.loadDefault(ctx)
	// srv.userServer = grpc_server.NewGrpcServer(srv.cfg.UserlServiceEndpoint)
	// srv.processors = append(srv.processors, srv.userServer)
	srv.factories = append(srv.factories, srv.postgresClient)
}