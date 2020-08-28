# bubble-effect

Bakery shopping cart calculator for The CAI Bakery Challenge

### Build project

`go build`

### Run sample

`./bubble-effect`

### Run tests

From the root of the project:

`
go test ./...
`

### Possible future enhancements:

- Represent price in cents (int), since floats are not the best choice for money calculations
- Add shopping cart functionality so it is easier to interact with the cart, e.g., store cart state, add ability to set date, add item, remove item, and clear cart
- Revisit shop data retrieval, as it might make sense to cache or initialize at some other point in time
