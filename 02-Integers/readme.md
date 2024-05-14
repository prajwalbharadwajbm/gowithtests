# Integers Package

The `integers` package provides utility function add, which we will be testing and wraping errors using '%d'. 

### Usage
#### Add Function
The Add function takes two integers and returns their sum.

### Tests
TestAdd and ExampleAdd are what we are learning in this package. Any function written starting with Example will be shown as example in godoc and also can be run when Output field is commented inside the example function.
Run godoc locally to understand how it looks.

### Documentation
You can view the documentation for this package by running godoc -http=:6060 and navigating to http://localhost:6060/pkg/github.com/prajwalbharadwajbm/gowithtests/integers/ in your web browser.

### Example Code
The integers package includes testable examples that are compiled and executed as part of the test suite. These examples also appear in the GoDoc documentation.

### ExampleAdd Function
The ExampleAdd function demonstrates how to use the Add function by enabling you to run it through documentation.

### Running Tests
To run the tests for the integers package, use the following command:

```bash
go test -v
```
You should see output similar to the following:

```bash
=== RUN   TestAdder
--- PASS: TestAdder (0.00s)
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
PASS
ok      github.com/prajwalbharadwajbm/gowithtests/integers 0.002s
Project Structure
```
Your project directory might look something like this:
```bash
gowithtests
|-> 01-Hello,World
|   |- go.mod
|   |- hello.go
|   |- hello_test.go
|   |- README.md
|-> 02-Integers
|   |- adder.go
|   |- adder_test.go
|   |- README.md
|- go.mod
```
