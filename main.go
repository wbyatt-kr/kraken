package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"sync"
	
	"kraken/admin"
	"kraken/events"
	"kraken/gateway"
	"kraken/persistence"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
		panic(err)
	}
}

func sqlQueries() (*persistence.Queries, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionString := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	
	if err != nil {
		return nil, err
	}

	queries := persistence.New(db)

	return queries, nil
}

func main() {
	var err error
	ctx := context.Background()


	queries, err := sqlQueries()
	
	apiGateway := gateway.Gateway{ Queries: queries }

	if err != nil {
		panic(err)
	}


	wg := &sync.WaitGroup{}
	wg.Add(1)

	gatewayServer, err := apiGateway.New(ctx, wg, ":8080")
	
	reloadEvent := events.Event{}

	reloadEvent.On(
		func() {
			log.Printf("Reloading configuration")
			log.Printf("Killing old server")

			wg.Add(1)

			if err := gatewayServer.Shutdown(context.TODO()); err != nil {
				panic(err)
			}

			log.Printf("Starting new server")
			gatewayServer, err = apiGateway.New(ctx, wg, ":8080")
		},
	)

	adminInterface := admin.Admin{ Queries: queries, ReloadEvent: reloadEvent }

	go func() {
		err = adminInterface.New(ctx, ":8081")
		failOnError(err, "Failed to create admin interface")
	}()
	

	wg.Wait();

	log.Printf("main: done. exiting")
}