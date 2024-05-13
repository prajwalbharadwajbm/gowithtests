# Hello World Greetings in Go

This repository contains a simple Go application that demonstrates how to generate multilingual greetings. It supports English, Spanish, and French greetings. The application is built using Test-Driven Development (TDD) principles, and the repository includes both the application code and tests.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You need to have Go installed on your machine. To install Go, visit [https://golang.org/dl/](https://golang.org/dl/) and download the appropriate version for your operating system.

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/prajwalbharadwajbm/gowithtests.git
cd gowithtests/01-Hello,World
```

### Running the Tests
Execute the tests to ensure everything is working correctly:

```bash
go test
```

### Running the Application
Run the main application:

```bash
go run hello.go
```

### Documentation
This project uses godoc to generate documentation from comments in the code. To view the documentation, you'll need to run the godoc server locally.

#### Installing godoc
If godoc is not already installed, you can install it using the following command:

```bash
go get golang.org/x/tools/cmd/godoc
go install golang.org/x/tools/cmd/godoc
```

#### Running godoc Locally
Run the following command in the root directory of this project:
```bash
godoc -http=:6060
```
This command starts a local godoc server on port 6060.

#### Viewing the Documentation
Open your web browser and visit http://localhost:6060/pkg/ to view the documentation. Navigate to the specific package for this project to see detailed documentation about the functions and types defined in the project.
