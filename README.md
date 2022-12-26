# web3-indexer
Purpose of this repo is to create a simple stack to index ethereum blockchain data.
# Getting Started
```bash
# For Go section
cd src
cp .env.example .env # Update .env with your own values
```
# Components (WIP)
- Etherscan API and Endpoint
- Scheduler service - polls Etherscan API for new blocks and saved to DB 
- DB - PostgresQL 
- Frontend - ReactJS