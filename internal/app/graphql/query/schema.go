package query

import "github.com/graphql-go/graphql"

// GetQuerySchema schema for query / get record
func GetQuerySchema() graphql.ObjectConfig {
	queryFields := graphql.Fields{}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields}

	return rootQuery
}
