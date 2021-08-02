package main

import (
	"fmt"
	"github.com/hjjddb/gqlgen-todos/graph"
	"github.com/hjjddb/gqlgen-todos/graph/generated"
	"github.com/hjjddb/gqlgen-todos/graph/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "gorm.io/driver/sqlite"
)

const defaultPort = "8088"

var db *gorm.DB

func initDB() {
	var err error
	dburi := "test.sqlite"
	db, err = gorm.Open(sqlite.Open(dburi), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Can't connect to database")
	}

	db.AutoMigrate(&model.User{}, &model.Wallet{}, &model.Giftcard{}, &model.GiftcardTransaction{}, &model.CryptoTransaction{})

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}