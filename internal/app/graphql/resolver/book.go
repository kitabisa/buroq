package resolver

import (
	"database/sql"
	"errors"

	"github.com/graphql-go/graphql"
)

/*
	This is sample code for simulating get book data
*/

type Book struct {
	ID          uint64         `json:"id"`
	Title       string         `json:"title"`
	Author      string         `json:"author"`
	Description sql.NullString `json:"description"`
}

var AllBooks = []Book{
	Book{
		ID:          1,
		Title:       "The Pragmatic Programmer",
		Author:      "Andrew Hunt and Dave Thomas",
		Description: sql.NullString{},
	},
	Book{
		ID:          2,
		Title:       "Introduction to Algorithms",
		Author:      "Thomas H. Cormen, Charles E. Leiserson, Ronald L. Rivest, and Clifford Stein",
		Description: sql.NullString{String: "Lorem ipsum dolor sit amet", Valid: true},
	},
	Book{
		ID:          3,
		Title:       "Code: The Hidden Language of Computer Hardware and Software",
		Author:      "Charles Petzold",
		Description: sql.NullString{},
	},
}

// GetBooks handle get all books data
func GetBooks(p graphql.ResolveParams) (result interface{}, err error) {
	result = AllBooks
	return
}

// GetBook handle get all books data
func GetBook(p graphql.ResolveParams) (result interface{}, err error) {
	id := p.Args["id"].(int)

	if id > len(AllBooks) {
		err = errors.New("Book not found")

		return
	}

	result = AllBooks[(id - 1)]

	return
}
