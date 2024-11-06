/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tyagnii/gw-currency-wallet/config"
	"github.com/tyagnii/gw-currency-wallet/internal/db"
	"github.com/tyagnii/gw-currency-wallet/internal/handlers"
	"github.com/tyagnii/gw-currency-wallet/internal/logger"
	"github.com/tyagnii/gw-currency-wallet/internal/token"
	"os"

	"github.com/spf13/cobra"

	_ "github.com/tyagnii/gw-currency-wallet/docs"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		// Init logger
		sLogger, err := logger.NewSugaredLogger()
		if err != nil {
			panic(err)
		}

		// Read config.env
		err = config.ReadConfig("config.env")
		if err != nil {
			sLogger.DPanicf("Cannot read configuration file: %v", err)
			os.Exit(1)
		}

		// Init database schema
		if err := db.InitSchema(); err != nil {
			sLogger.DPanicf("Cannot init schema: %v", err)
			os.Exit(1)
		}

		// LoadEnvironment for token package
		token.LoadEnvironment()

		// Init restapi routers
		r, err := handlers.NewRouter(sLogger)
		if err != nil {
			sLogger.DPanicf("Cannot create router: %v", err)
			os.Exit(1)
		}

		// Run swagger
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// Run server
		sLogger.Fatal(r.Run())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
