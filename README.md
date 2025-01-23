# bookManagementSystemAPI

## Installation

```sh
    go get
    go mod tidy
```

## Running the app

```sh
    go run main.go
```
## 1.Search all books
```sh
curl --location 'http://localhost:8080/books'
```
## 2.Search for a book by ID
```sh
curl --location 'http://localhost:8080/books/search/id/1'
```
## 3.Search for books by Title
```sh
curl --location 'http://localhost:8080/books/search?title=Learning%20Fiber'
```
## 4.Search for books by Author
```sh
curl --location 'http://localhost:8080/books/search?author=Alice%20Smith'
```
## 5.Search for books by Category
```sh
curl --location 'http://localhost:8080/books/search?category=Programming'
```
## 6.Borrow book
```sh
curl --location --request POST 'http://localhost:8080/books/borrow/1'
```
## 7.Return book
```sh
curl --location --request POST 'http://localhost:8080/books/return/1'
```
## 8.Display the most borrowed books
```sh
curl --location 'http://localhost:8080/books/most-borrowed'
```