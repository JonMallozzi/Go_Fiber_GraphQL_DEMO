package main

import (
	"net/http"

	"github.com/JonMallozzi/Go_Fiber_GraphQL_DEMO/graph/postgres"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/JonMallozzi/Go_Fiber_GraphQL_DEMO/graph"
	"github.com/JonMallozzi/Go_Fiber_GraphQL_DEMO/graph/generated"
	"github.com/gofiber/fiber"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	//"github.com/arsmn/gqlgen" //provides GrahpQL with fasthttp
	//graph/model/generated
)

//An example set up for basic fiber
func main() {

	db := postgres.PostgresConnect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Type("html")
		c.Send("<h1> Hello World! </h1>")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) {
		c.Type("html")
		if c.Params("name") != "" {
			c.Send("<h1> Hello " + c.Params("name") + "! </h1>")
		} else {
			c.Send("Please Provide A Name to Be Greeted")
		}
	})

	//configuring data for the resolvers
	config := generated.Config{Resolvers: &graph.Resolver{
		UsersRepo: postgres.UserRepo{DB: db},
	}}

	app.Post("/query", func(ctx *fiber.Ctx) {
		h := handler.NewDefaultServer(generated.NewExecutableSchema(config))
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			h.ServeHTTP(writer, request)
		})(ctx.Fasthttp)
	})

	app.Get("/playground", func(ctx *fiber.Ctx) {
		h := playground.Handler("GraphQL", "/query")
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			h.ServeHTTP(writer, request)
		})(ctx.Fasthttp)
	})

	app.Listen(3000)
}
