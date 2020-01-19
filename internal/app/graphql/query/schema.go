package query

import (
	"github.com/graphql-go/graphql"
	"github.com/kitabisa/buroq/internal/app/graphql/types"
)

// GetQuerySchema schema for query / get record
func GetQuerySchema() graphql.ObjectConfig {
	queryFields := graphql.Fields{
		"books": &graphql.Field{
			Type: types.BookType,
		},
		"book": &graphql.Field{
			Type: graphql.NewList(types.BookType),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields}

	return rootQuery
}
