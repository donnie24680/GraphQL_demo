package main

import (
	"GraphQL_demo/graph"
	"GraphQL_demo/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	oo "github.com/donnie24680/liboo"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"net/http"
)

const infuraURL = "https://rpc.payload.de"

func main() {
	defer func() {
		if err := recover(); nil != err {
			oo.LogW("panic err: %v", err)
		}
	}()

	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	resolver := &graph.Resolver{EthClient: client}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/query", srv)
	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))

	log.Println("connect to http://localhost:8080/playground for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))

	oo.WaitExitSignal()
}
