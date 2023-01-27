/*
Copyright © 2023 Anisur Rahman 'sunny.cse7575@gmail.com'
*/
package cmd

import (
	"fmt"
	"github.com/anisurahman75/apiDesign/api"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

// startServerCmd represents the startServer command
var startServerCmd = &cobra.Command{
	Use:   "startServer",
	Short: "CMD: StartServer for Running this apiServer ",
	Long:  `This API server provides endpoints to create,read,update & delete users and Books.`,
	Run: func(cmd *cobra.Command, args []string) {
		var port string = ":3000"

		if len(args) > 0 {
			port = args[0]
		}
		log.Println("start called! start the server at port", port)
		s := api.CreateNewServer()
		s.MountHandlers()
		err := http.ListenAndServe(port, s.Router)
		if err != nil {
			fmt.Printf("error : %s\n", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
	var port string
	startServerCmd.Flags().StringVar(&port, "port", ":3000", "Running this port")
}
