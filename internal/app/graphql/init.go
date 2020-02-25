package graphql

import (
	"github.com/sirupsen/logrus"
	"github.com/graphql-go/graphql"
	"github.com/kitabisa/buroq/internal/app/graphql/query"
	"github.com/kitabisa/buroq/internal/app/graphql/resolver"
	"github.com/kitabisa/buroq/internal/app/service"
)

// InitGraphqlSchema init graphql schema
func InitGraphqlSchema(svc *service.Services) (schema graphql.Schema) {
	// Init graphql: load all schema and connect to services
	resolver.Init(
		resolver.WithServices(svc),
	)
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(query.GetQuerySchema()),
		// Mutation: graphql.NewObject(mutation.GetMutationSchema()),
	})
	if err != nil {
		logrus.Fatalf("Failed to create schema for the graphql: %s", err)
		return
	}

	return schema
}
