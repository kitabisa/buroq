package mutation

import "github.com/graphql-go/graphql"

// GetMutationSchema schema for mutation / edit record
func GetMutationSchema() graphql.ObjectConfig {
	mutationFields := graphql.Fields{}

	rootMutation := graphql.ObjectConfig{Name: "MutationQuery", Fields: mutationFields}

	return rootMutation
}
