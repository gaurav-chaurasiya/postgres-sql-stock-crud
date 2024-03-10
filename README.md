# postgres-sql-stock-crud


Postgress and Golang API

install packages command


mkdir go-postgres-sql-stock \n
cd go-postgres-sql-stock
go mod go-postgres-sql-stock
go mod init go-postgres-sql-stock
go mod tidy
go build

go get github.com/go-logr/logr@v1.4.1
go get github.com/gorilla/mux@v1.8.1
go get github.com/joho/godotenv@v1.5.1
go get github.com/lib/pq@v1.10.9
go get golang.org/x/mod@v0.8.0
go get golang.org/x/sys@v0.5.0
go get golang.org/x/text@v0.14.0
go get golang.org/x/tools@v0.6.0

API list

All item list
localhost:8000/api/stock 
Note 
Body None,method GET

Get item by id
localhost:8000/api/stock/14
Note 
Body None,method GET

Update item by id
localhost:8000/api/stock/15
Note
Body row,method PUT
{
    "Name":"n1 name 123",
    "Price":2133,
    "Company":"c2 new dwew"
}

Delete item by id
localhost:8000/api/deletestock/14
Note
Body none,method DELETE

Create item by id
localhost:8000/api/newstock
Note
Body raw-json,method POST
{
    "Name":"n1 name",
    "Price":213,
    "Company":"c new dwew"
}

