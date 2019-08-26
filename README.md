<p align="center">
  <img src="https://user-images.githubusercontent.com/4137581/63674200-a3cba480-c803-11e9-9f47-b90669bef337.png" height="130px"/>
</p>
<h1 align="center">Sampark</h1>

## Introduction

Sampark is a contact book REST API written in Golang.

### API Specification

- POST `/contacts` Creates a contact
- GET `/contacts` Lists contacts
- GET `/contacts/{id}` Fetches contact by ID
- GET `/contacts/search` Searches a contact
- PUT `/contacts/id` Updates a contact
- DELETE `/contacts/id` Deletes contact by ID

## Development

> Pre-requisites: Install latest stable versions of Docker and Docker Compose.

1. Clone the repository locally.
2. Add .env file in the repository root by modifying the .env.template file as needed.
3. From repository root, run:
```bash
docker-compose up
```
4. API will be up and running at http://localhost:5000.

### Dependency Management

`Sampark` uses Go modules with semantic versioning and is tested with Go 1.12+.

* Update all direct and indirect dependencies using `go get -u`.
* Remove unused dependencies using `go mod tidy`.
* Add a new dependency using `go get <path-to-dependency>`.

#### Dependencies

* [echo](https://echo.labstack.com/) - Web framework
* [realize](https://github.com/oxequa/realize) - Live reloading
* [mgo.v2](https://gopkg.in/mgo.v2) - MongoDB driver
<!-- * [testify](https://github.com/stretchr/testify) - Assertions library -->

## References

* [Go Modules](https://github.com/golang/go/wiki/Modules)
<!-- * [Using MongoDB Go Driver](https://vkt.sh/go-mongodb-driver-cookbook/) -->

## Acknowledgements

> Logo credit goes to [Freepik](https://www.flaticon.com/authors/freepik) from [Flaticon](https://www.flaticon.com) and is licensed under [Creative Commons BY 3.0](http://creativecommons.org/licenses/by/3.0).
