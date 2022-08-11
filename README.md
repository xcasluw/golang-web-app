# Bookings and Reservations

This is the repository for my bookings and reservations project.

- Built in GO version 1.13
- Uses the [chi router](http://github.com/go-chi/chi)
- Uses [alex edwards SCS](https://github.com/alexedwards/scs) session management
- Uses [nosurf](https://github.com/justinas/nosurf)

# Easy way to start the application
- if you are using linux or mac, make sure to run to create an executable:
`chmod +x run.sh`

- and after that, just run: `./run.sh`

- if you don't want to use this way, you can run: `go run ./cmd/web`


# How to run tests

### default (current directory)
`go test`

### more detailed test (current directory)
`go test -v`

### test with coverage (current directory)
`go test -cover`

### test with html file to check uncovered lines (current directory)
`go test -coverprofile=coverage.out && go tool cover -html=coverage.out`

### test all directories (root application directory)
`go test -v ./... `