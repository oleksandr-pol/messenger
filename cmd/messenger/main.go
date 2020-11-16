package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/oleksandr-pol/messenger/internal/env"
	"github.com/oleksandr-pol/simple-go-service/pkg/utils"
)

func main() {
	var port int
	var dbPort int
	var dbName string
	var dbHost string
	var dbUserName string
	var dbPass string

	flag.IntVar(&port, "p", utils.GetDefaultIntVal(os.Getenv("PORT"), 8000), "specify port to use.  defaults to 8000")
	flag.IntVar(&dbPort, "dbPort", utils.GetDefaultIntVal(os.Getenv("DB_PORT"), 5432), "specify data base host name. defaults to 5432")
	flag.StringVar(&dbName, "dbName", utils.GetDefaultStringVal(os.Getenv("DB_NAME"), "messenger"), "specify data base name. defaults to messenger")
	flag.StringVar(&dbHost, "dbHost", utils.GetDefaultStringVal(os.Getenv("DB_HOST"), "localhost"), "specify data base host name. defaults to localhost")
	flag.StringVar(&dbUserName, "dbUserName", utils.GetDefaultStringVal(os.Getenv("DB_USER_NAME"), "oleksandr"), "specify data base host name. defaults to oleksandr")
	flag.StringVar(&dbPass, "dbPass", utils.GetDefaultStringVal(os.Getenv("DB_PASS"), "empty"), "no default value")
	flag.Parse()
	dbConf := env.DbConfig{DbHostName: dbHost, DbHostPort: dbPort, DbUserName: dbUserName, DbPassword: dbPass, DbName: dbName}

	db, err := env.NewDbCon(dbConf)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	router, setUpErr := env.SetUpServer(db)
	if setUpErr != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Listing for requests at http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
