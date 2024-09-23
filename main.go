package main

import (
	"context"
	"fmt"
	"go-oapi-test/api"
	dbCon "go-oapi-test/db/sqlc"
	"go-oapi-test/tools"
	"go-oapi-test/util"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	middleWare "github.com/oapi-codegen/nethttp-middleware"

	amqp "github.com/rabbitmq/amqp091-go"
)

var db *dbCon.Queries

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not loadconfig: %v", err)
	}

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, uri *url.URL) ([]byte, error) {
		return os.ReadFile(uri.Path)
	}

	doc, err := loader.LoadFromFile("api.yaml")
	if err != nil {
		panic(err)
	}

	if err = doc.Validate(loader.Context); err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(context.Background(), config.DbSource)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			fmt.Println("Error closing connection...")
		}
	}(conn, context.Background())

	db = dbCon.New(conn)

	fmt.Println("PostgreSql connected successfully...")

	rconn, err := amqp.Dial(config.RabbitMq)
	if err != nil {
		log.Fatalf("Could not connect to rabbitmq: %v", err)
	}
	defer rconn.Close()

	rabbitChannel, err := rconn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer rabbitChannel.Close()

	err = rabbitChannel.ExchangeDeclare("branch-max-users-changed", "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Could not create exchange: %v", err)
	}

	_, err = rabbitChannel.QueueDeclare("branch-max-users-changed", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Could not create query: %v", err)
	}

	err = rabbitChannel.ExchangeDeclare("branch-group-branch-changed", "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Could not create exchange: %v", err)
	}

	_, err = rabbitChannel.QueueDeclare("branch-group-branch-changed", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Could not create query: %v", err)
	}

	r := chi.NewRouter()
	validatorOptions := &middleWare.Options{}
	a, err := tools.NewJwsAuthenticator(config)
	if err != nil {
		log.Fatalln("error creating authenticator:", err)
	}
	validatorOptions.Options.AuthenticationFunc = tools.NewAuthenticator(a)
	r.Use(middleWare.OapiRequestValidatorWithOptions(doc, validatorOptions))
	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		fileBytes, _ := os.ReadFile("api.yaml")
		w.Header().Set("Content-Type", "text")
		_, err := w.Write(fileBytes)
		if err != nil {
			return
		}
	})

	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer(db, *a, rabbitChannel)

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(&server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
