# web3-indexer
Purpose of this repo is to create a simple stack to index ethereum blockchain data.
# Getting Started
```bash
# For Go section
cp app.env.example app.env # Update the env variables
go run cmd/main/main.go # Run the app
```
# Components (WIP)
- Etherscan API and Endpoint
- Scheduler service - polls Etherscan API for new blocks and saved to DB 
- DB - PostgresQL 
- Frontend - ReactJS
# References
- GORM setup - https://codevoweb.com/setup-golang-gorm-restful-api-project-with-postgres/
- GORM datatypes - https://github.com/go-gorm/datatypes
- Logging - https://github.com/sirupsen/logrus