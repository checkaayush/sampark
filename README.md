<div align="center">
  <p>
    <img src="https://user-images.githubusercontent.com/4137581/63674200-a3cba480-c803-11e9-9f47-b90669bef337.png" height="130px"/>
  </p>
  <h1>Sampark</h1>
  <a href="https://goreportcard.com/report/github.com/checkaayush/sampark">
    <img src="https://goreportcard.com/badge/github.com/checkaayush/sampark"/>
  </a>
</div>

## Introduction

Sampark is a contact book REST API written in Golang. It uses MongoDB as the primary data store and is built using [SOLID](https://en.wikipedia.org/wiki/SOLID) design priciples and [12-Factor App](https://12factor.net/) methodology in mind.

### API Specification

- GET `/v1/health` Health check to indicate API health
- POST `/v1/contacts` Creates a contact
- GET `/v1/contacts` Lists and searches contacts
- GET `/v1/contacts/{id}` Fetches contact by ID
- PATCH `/v1/contacts/{id}` Updates a contact
- DELETE `/v1/contacts/{id}` Deletes contact by ID

### API Remarks

- API has been hosted on Heroku (https://sampark.herokuapp.com)
- API supports CRUD operations on `Contact` entity
- Each contact has a unique email address, which is ensured by having a unique index on the `email` field
- Allows searching by name and email address. GET `/v1/contacts?name=<NAME>&email=<EMAIL>` lets you search via name and/or email
- Search supports pagination and returns 10 items by default per invocation. Example: GET `/v1/contacts?page=1&limit=5`.
- Added tests for each functionality
- Basic authentication has been added using environment variables
- Some preliminary load tests will ensure that the code can scale-out for millions of contacts

### Error Codes

| Code Range | Description                                                                                                                             |
| ---------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| 2xx        | This range of response code indicates that request was fulfilled successfully and no error was encountered.                             |
| 400        | This return code indicates that there was an error in fulfilling the request because the supplied parameters are invalid or inadequate. |
| 401        | This return code means that we are not able to authenticate your request. Please re-check your username and password.                   |
| 5xx        | This response code indicates that there was an internal server error while processing the request.                                      |

## Development

> Pre-requisites: Install latest stable versions of Docker and Docker Compose.

1. Clone the repository locally.
2. Add .env file in the repository root by modifying the .env.template file as needed.
3. From repository root, run:
```bash
make start
```
4. API will be up and running at http://localhost:5000.

## Testing

From repository root, run:
```bash
make test
```

### Dependency Management

`Sampark` uses Go modules with semantic versioning and is tested with Go 1.12+.

* Update all direct and indirect dependencies using `go get -u`.
* Remove unused dependencies using `go mod tidy`.
* Add a new dependency using `go get <path-to-dependency>`.

#### Dependencies

* [echo](https://echo.labstack.com/) - Web framework
* [realize](https://github.com/oxequa/realize) - Live reloading
* [mgo.v2](https://gopkg.in/mgo.v2) - MongoDB driver
* [testify](https://github.com/stretchr/testify) - Assertions library

## References

* [Go Modules](https://github.com/golang/go/wiki/Modules)
<!-- * [Using MongoDB Go Driver](https://vkt.sh/go-mongodb-driver-cookbook/) -->

## Acknowledgements

> Logo credit goes to [Freepik](https://www.flaticon.com/authors/freepik) from [Flaticon](https://www.flaticon.com) and is licensed under [Creative Commons BY 3.0](http://creativecommons.org/licenses/by/3.0).
