/*
Copyright © 2023 Anisur Rahman 'sunny.cse7575@gmail.com'
*/
package cmd

import (
	"flag"
	"fmt"
	"github.com/anisurahman75/apiDesign/api/handler"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func GetConnectionString() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3307"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}

	password := os.Getenv("DB_PASS")
	if password == "" {
		password = "@root"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "crudgo"
	}

	str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
	fmt.Println(str)
	return str
}
func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open(GetConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("----------Finally Connected -----------")
	return db
}

var startServerCmd = &cobra.Command{
	Use:   "startServer",
	Short: "CMD: StartServer for Running this apiServer ",
	Long:  `This API server provides endpoints to create,read,update & delete users and Books.`,
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		log.Println("\n----------------------StartServer Called!")
		fmt.Println("----------------------Port: ", handler.Port, "\n----------------------Authentication: ", handler.Auth, "\n\n")
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error getting env, %v", err)
		}
		fmt.Println("We are getting the env values")
		s := handler.CreateNewServer()
		s.MountHandlers()
		tem := ":" + handler.Port
		err := http.ListenAndServe(tem, s.Router)
		if err != nil {
			fmt.Printf("error : %s\n", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
	startServerCmd.PersistentFlags().StringVarP(&handler.Port, "Port", "p", "3030", "This flag sets the Port of the server")
	startServerCmd.PersistentFlags().BoolVarP(&handler.Auth, "Auth", "a", true, "This flag will impose/bypass authentication to API server")
}
