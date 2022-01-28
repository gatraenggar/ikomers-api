# iKomers-API (in progress)
## What is this ?
GraphQL API for iKomers app (e-commerce-like app). A self-developed project for learning purpose.
This repository is created as a medium to learn how to implement GraphQL using Go Language to solve an e-commerce business use-case.

## How to test this app in your local environment ?
### Prerequisite installations
#### 1. MySQL database
#### 2. Go Language

### Installation
#### 1. `git clone https://github.com/gatraenggar/ikomers-api.git`
#### 2. `cd ikomers-api`
#### 3. `go mod tidy` to cleans up unused dependencies or adds missing dependencies

### Configuration
#### 4. Create a PostgreSQL database for testing
#### 5. Rename `example.env` to `.env`. Then change the    `db_user`, `db_password`, & `db_name` value in that `.env` file based on yours
#### 5.1. Change `access_token_key` and `refresh_token_key` in `.env` file with your generated encryption keys
You can generate your own encryption keys with following steps:
1. Enter this site https://www.allkeysgenerator.com/Random/Security-Encryption-Key-Generator.aspx
2. Choose `Encryption key`
3. Check the `Hex ?` field
4. Click the `512-bit` option
5. Click `Get new results` button, then copy the generated string to the config file

### Run the App
#### 6. `go run main.go migrate_test` to migrate/create the database tables
#### 7. `go test -v ./...` to run all the tests