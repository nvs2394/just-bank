# just-bank

Learn go with real example

- Golang

- Rest API

- Hexagonal Architect

### APIs:

> Get All Customers      
GET  /customers


> Get Customer by ID    
GET  /customers/{customer_id}

> Create new account    
POST /customers/{customer_id}/accounts

> Make a transaction     
POST /customers/{customer_id}/accounts/{account_id}

### Database Schema

![EPR Just Bank](./just-bank-db.png)
### How to run local


- Start MySQL instance

- Run `SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=just_bank_user DB_PASSWORD=password DB_NAME=just_bank_db go run main.go`

![Principle: Separate User-Side, Business Logic and Server-Side](https://blog.octo.com/wp-content/uploads/2020/06/archi_hexa_en_00.png)

![The Hexagone](https://blog.octo.com/wp-content/uploads/2020/06/archi_hexa_en_06.png)

### Code coverage