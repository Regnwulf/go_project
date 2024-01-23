# Go CRUD Project

This is a simple CRUD (Create, Read, Update, Delete) project in Go using the Gin framework.

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/Regnwulf/go_project.git
    ```

2. Change into the project directory:

    ```bash
    cd go_project
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Run the project:

    ```bash
    go run main.go
    ```

The server should now be running at `http://localhost:8080`.

## API Endpoints

### List Products
```bash
curl http://localhost:8080/products
```

### Get a Specific Product (replace <ID> with the real product ID)
```bash
curl http://localhost:8080/products/<ID>
```

### Create a Product
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "Product 1", "price": 19.99}' http://localhost:8080/products
```

### Update a Product (replace <ID> with the real product ID)
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name": "Updated Product", "price": 29.99}' http://localhost:8080/products/<ID>
```

### Delete a Product (replace <ID> with the real product ID)
```bash
curl -X DELETE http://localhost:8080/products/<ID>
```

# Project Structure

The project follows the following structure:

- **controllers**: Contains the request handlers (controllers).
- **models**: Defines the data structures (models).
- **services**: Implements the business logic (services).
- **main.go**: Entry point of the application.

Feel free to modify the project structure or add more details to this README based on your specific requirements.
