# go_http_server

This whole repo is a part of a project-course on [Boot.dev](https://www.boot.dev/lessons/50f37da8-72c0-4860-a7d1-17e4bda5c243)
The idea is to build a HTTP Server in Go without frameworks, to really get into how they work and so forth.

Goals of This Course

  - Understand what web servers are and how they power real-world web applications
  - Build a production-style HTTP server in Go, without the use of a framework
  - Use JSON, headers, and status codes to communicate with clients via a RESTful API
  - Learn what makes Go a great language for building fast web servers
  - Use type safe SQL to store and retrieve data from a Postgres database
  - Implement a secure authentication/authorization system with well-tested cryptography libraries
  - Build and understand webhooks and API keys
  - Document the REST API with markdown

## How to use locally.

To use this repo locally a few things are required:
- [PostgreSQL](https://www.postgresql.org/)
- [Go](https://go.dev/)
- [Goose](https://github.com/pressly/goose)

You will need to have these fields in your .env file, placed in the root of the project:
```
DB_URL = <connection string to Postgres Database>
PLATFORM = dev 
JWT_SECRET = <base64 string for JWT signing>
POLKA_KEY = <a random string, could be 123 if you want> this is not a real API, just used to mock the webhook.
```

After setting that up navigate to the sql/schema folder and run:
```
goose postgres <db_conn_string> up 
```
to migrate all the required tables to the database.

Then run
```
go run .
```
from the root of the project, if successfull the server should now be ready to accept CRUD operations.

To start using the web server, first create a new user by sending a POST request to /api/users. The endpoint expects a JSON body like so:
```
{
  "email": "email@test.com",
  "password": "123456"
}
```

The make a POST request to /api/login, using the same JSON body as above, to authenticate and recive a JWT token.

A logged in user can then post "chirps" by sending a POST request, with their JWT Token in the Authorization header, to /api/chirps with a JSON body constructed thusly:
```
{
  "body": "this was a triumph"
}
```

All "chirps" can be fetched by making a GET request to /api/chirps, this also accepts optional query parameters like /api/chirps?sort=asc/desc for sorting or a userId, /api/chirps/?author_id=123 for getting "chirps" for only that user. 

A specific "chirp" can be fetched by supplying an id parameter like so: /api/chirps/{id}. They can also be deleted by sending a DELETE request to /api/chirps/{chirpID}

The JWT Tokens are set to have a lifetime of one hour. A user can recieve a new token by making a POST request to /api/refresh with their refresh token, found in the user response when logging in. A refresh token can also be revoked by sending a POST request to /api/revoke.

Calling a webhook can be simulate by making a POST requst to /api/polka/webhooks this "upgrades" a user to supposedely recieve some form of additional benefit. The endpoint expects a specific JSON body, as seen below:
```
{
  "event": "user.upgraded",
  "data": {
    "user_id": "<id of the user to be upgraded>"
  }
}
```

The remaning endpoints are mainly for metrics or to check the status of the server. A GET request can be sent to /api/healthz to know wether or not the server is ready to recieve connections while /admin/metrics will return some metric data about the number of hits on the server.

> [!CAUTION]
> However making a POST requst to /admin/reset will RESET the database to empty, so this is used mainly for development and will only work then the PLATFORM .env key is set to dev. 