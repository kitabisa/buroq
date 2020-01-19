package resolver

import "github.com/graphql-go/graphql"

import "log"

// GetBooks handle get all books data
func GetBooks(p graphql.ResolveParams) (result interface{}, err error) {
	log.Println("invoked")
	return
}
