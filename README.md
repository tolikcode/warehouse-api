# warehouse-api

Backend for [warehouse-client](https://github.com/tolikcode/warehouse-client.git)

## Development information

### Prerequisites
- An existing empty [Postgres](https://www.postgresql.org/) database

### Steps
 1. Update `warehouse-api-src/db/db.go` with the connection string to your database
 2. Update localhost urls in `warehouse-api-src/main.go` if needed
 2. `cd warehouse-api-src`
 3. `go mod download`
 4. `go run .`


 ## TODO
 - Add unit tests
 - To accomplish previous step, it would be helpful to have explicit dependencies and DI. [wire](https://github.com/google/wire) looks like a good option with wiring at build time. Also might extract separate services from controllers.
 - Add authentication
 - Finish infrastructure (add a database cluster). Also I prefer to have infra in the same language as the application itself, so rewrite it in golang.
 - Create a deployment script (with `docker build`, `docker push`, `cdk deploy` etc). Add github action for build and deploy
 - Simplify local development environment bootstrap for new users. Add a config file. Maybe run everything in docker
 - Add caching if a lot of requests are expected
 - If large inventory.json and products.json are expected, then process them asynchronously (with 201 Accepted etc)
 - Don't know yet how GORM (ORM used in the project) handles db connections. Research to make sure connections are pooled and reopened if closed
 - Api versioning. Fix bugs and edge cases :) 
 

