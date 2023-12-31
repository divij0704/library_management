LIBRARY MANAGEMENT USING GO : 

The project is a Go language-based HTTP API using the GoFr framework, integrating MongoDB for data persistence. The Library Management project is a simple Go (Golang) application designed to manage a library's book inventory through a RESTful API. It utilizes the Gin framework for handling HTTP requests and MongoDB as the database for persistent storage. The project supports basic CRUD (Create, Read, Update, Delete) operations on book entities, allowing users to add, retrieve, update, and delete book records. The code is organized following a standard Go project structure, promoting modularity and maintainability.

Postman Collection including the documentation for the project : https://www.postman.com/supply-architect-28849416/workspace/library-management/collection/31823378-02f09e2a-0314-468a-9a74-c1d1de82ad85

Steps to run the project:
1. Prerequisites : Go, MongoDB
2. Setting up the project: 
   i. Clone the repository on your system 
   git clone https://github.com/divij0704/library_management.git
   cd library_management
   
   ii.Install Dependencies: 
   go mod download
4. Configuring MongoDB: Start your MongoDB server locally using the default string -              mongodb://localhost:27017
5. Running the Project: Open the terminal and enter : go run main.go
6. The API is now running. You can access it using a tool like Postman or your web browser.
   
API Endpoints:

1.Create a Book:
Endpoint: POST http://localhost:8080/books (provide the book details in JSON format)

2.Get All Books:
Endpoint: GET http://localhost:8080/books

3.Get a Book by ID:
Endpoint: GET http://localhost:8080/books/{id}

4.Update a Book by ID:
Endpoint: PUT http://localhost:8080/books/{id} (provide the book details in JSON format)

5.Delete a Book by ID:
Endpoint: DELETE http://localhost:8080/books/{id}
