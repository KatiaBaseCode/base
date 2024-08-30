# Base Framework

Base is a lightweight, modular framework for building RESTful APIs in Go. It provides a structured approach to developing web applications with a focus on simplicity, flexibility, and scalability.

## Features

- Modular architecture with support for multiple modules
- Built-in support for CRUD operations
- Automatic routing and request handling
- Integrated logging and error handling
- Configurable database connections 
- Support for multiple database types (MySQL, PostgreSQL, SQLite)
- JWT-based authentication and authorization
- Environment-based configuration via `.env` files
- Command-line tool for generating and removing modules via `bin/base` command
- Customizable routes and middleware
- Integrated Swagger documentation

## Installation

To install Base, make sure you have Go installed on your system, then run:

```
git clone github.com/base-api/base
```

Install swag if not already installed:
```
go get -u github.com/swaggo/swag/cmd/swag
```

## Usage

Run the following command to generate Swagger documentation:
```
swag init --parseDependency --parseInternal --parseVendor
```

Or add as alias in .bashrc or .zshrc:
```
alias swagg='swag init --parseDependency --parseInternal --parseVendor'
```

To start the application, run:

```
go run main.go
```

The application will start on port 8080 by default. You can access the Swagger documentation at `http://localhost:8080/swagger/index.html`.

## Configuration

Base uses environment variables for configuration. You can create a `.env` file in the root directory of your application to set the following variables:

```
SERVER_ADDRESS=:8080
JWT_SECRET=my_super_secret_key_123
API_KEY=api
ENV=debug
DB_DRIVER=sqlite
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=XXX
DB_NAME=base
DB_PATH=storage/base.db
LOG_LEVEL=info
CORS_ALLOWED_ORIGINS=http://localhost:3000
```

# bin/base - Module Generator and Destroyer

The `bin/base` command-line tool streamlines the process of creating and removing modules in your Base application.

## Features

- Generate new modules with a single command
- Create standardized structure for models, controllers, and services
- Automatically update the main application to include new modules
- Destroy existing modules and clean up related code

## Usage

### Generating a New Module

To generate a new module, use the `g` command followed by the module name and field definitions:

```
bin/base g [module_name] [field:type...]
```

Example:
```
bin/base g user name:string age:int email:string
```

This command will:
1. Create a new directory `app/users/` (note the plural form)
2. Generate `model.go`, `controller.go`, `service.go`, and `mod.go` files in the new directory
3. Update `app/init.go` to register the new module

### Destroying an Existing Module

To remove an existing module, use the `d` command followed by the module name:

```
bin/base d [module_name]
```

Example:
```
bin/base d user
```

This command will:
1. Remove the `app/users/` directory and all its contents
2. Update `app/init.go` to unregister the module

## Generated File Structure

For each module, the following files are generated:

- `model.go`: Defines the data structure for the module
- `controller.go`: Handles HTTP requests and responses
- `service.go`: Contains business logic for the module
- `mod.go`: Initializes the module and sets up routes

## Customization

You can customize the generated code by modifying the templates in the `templates/` directory.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

If you encounter any issues or have questions, please file an issue on the GitHub repository.