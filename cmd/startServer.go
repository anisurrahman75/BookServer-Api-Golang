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

		log.Println("\n----------------------StartServer Called!")
		fmt.Println("----------------------Port: ", api.Port, "\n----------------------Authentication: ", api.Auth, "\n\n")
		s := api.CreateNewServer()
		s.MountHandlers()
		tem := ":" + api.Port
		err := http.ListenAndServe(tem, s.Router)
		if err != nil {
			fmt.Printf("error : %s\n", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
	startServerCmd.PersistentFlags().StringVarP(&api.Port, "Port", "p", "3030", "This flag sets the Port of the server")
	startServerCmd.PersistentFlags().BoolVarP(&api.Auth, "Auth", "a", true, "This flag will impose/bypass authentication to API server")
}
