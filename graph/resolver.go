//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"github.com/JonMallozzi/Go_Fiber_GraphQL_DEMO/graph/postgres"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UsersRepo postgres.UserRepo
}
