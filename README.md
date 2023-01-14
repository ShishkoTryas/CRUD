# # Book API

## How to use this
Please create database postgres, or if u not it will error. For table, it will auto add(auto migrate) automatic.

And after that, download this repo, and copy this text, and run in terminal. and its done.

    go run main.go

## API 

* List of All Books:
    * GET http://localhost:8080/books - Get all books
    * GET http://localhost:8080/book/:id - Get by id
    * POST http://localhost:8080/books - Create new book
    * PUT http://localhost:8080/book/:id - Update book by id
    * DELETE http://localhost:8080/book/:id - Delete book by id
    
## API POST

 POST http://localhost:8080/books
 
 {
    "name": "bookName",
    "price": 1234
 }
 
 ## API PUT

PUT http://localhost:8080/book/4

{
    "name": "UpdateBookName",
    "price": 1234
 }
