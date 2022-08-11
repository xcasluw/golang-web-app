# Bookings and Reservations

This is the repository for my bookings and reservations project.

- Built in GO version 1.13
- Uses the [chi router](http://github.com/go-chi/chi)
- Uses [alex edwards SCS](https://github.com/alexedwards/scs) session management
- Uses [nosurf](https://github.com/justinas/nosurf)

# How to run tests

### default
`go test`

### more detailed test
`go test -v`

### test with coverage
`go test -cover`

### test with html file to check uncovered lines
`go test -coverprofile=coverage.out && go tool cover -html=coverage.out`