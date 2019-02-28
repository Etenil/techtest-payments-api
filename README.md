# techtest-payments-api
A demo payments API in Go.

_Note: The Author learned Go in 2 days then produced this._

## Requirements

- Golang 1.11.5
- [Go Dep](https://github.com/golang/dep) (Golang dependency manager)
- [Tavern](https://github.com/taverntesting/tavern) (API testing)

## Starting the API

Run the API like so:

```
dep ensure
go run main.go
```

The API will be running on port 8080.

## Unit tests

To run the unit tests, run the following:

```
go test ./...
```

## API tests

To run API tests, you need python3 intalled on your computer. Then
install tavern and pytest with `pip` like so:

```
pip install pytest tavern
```

Then start the API and run the api tests like so:

```
py.test apitests
```

_Remember to stop and start the API before running the API tests!_

