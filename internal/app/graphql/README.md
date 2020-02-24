# buroq graphql stuff
you have to enable the graphql first if you want to use it.

Change `is_enabled` value to `true` and you can configure the `route` on `params/buroq.toml`.

you can use the graphql by hit the endpoint `localhost:[port]/[route]`.

## mutation

Add all data manipulation in this schema (insert, update, delete)

## query

Add all query operation in this schema (select). I give you two examples `books` to get all books and `book(id: {id})` to get book with specific id.

## resolver

Write all function to resolve graphql request on here. You can access services from resolver by using `resolverInst.svc.{your-service}.{your-function-on-that-service}`.

## types

Define your query-able item for every model or operation

## Uninstall

If you don't want to use graphql, just disable `graphql.is_enabled` on `params/buroq.toml`.

## TODO
Modify error response from graphql resolver.