package types

import "github.com/graphql-go/graphql"

/* BookType
Here's sample type for graphql implementation in buroq
*/

// BookType type for books
var BookType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Book",
	Description: `Book data mapped to the book table.`,
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
	},
})
