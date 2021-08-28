# parallel-requests
A script written in Go that performs requests and returns the MD5 hashed response for each of them. The requests can be performed in parallel.

## Usage

To use the tool, first compile the code with the following command, or be sure that you have the binary.

```
go build main.go
```

The program can receive any number of URLs as arguments, and a request will be performed for each of them. The program will print the URL and the hashed response of a GET request made.

- Example 1: `./main http://google.com`
Console output: `http://google.com 1e2668de156fd73b31a5085e85d92c58`

- Example 2: `./main http://google.com http://reddit.com http://facebook.com`
Console output:
```
http://google.com 1e2668de156fd73b31a5085e85d92c58
http://facebook.com b2136a40aaec53a0181b3315d74cfcf2
http://reddit.com edd2761c38efcc82b631c6dd39766df5
```

The tool can be used by passing a flag, to indicate the concurrency level that will be used to perform requests. The flag is called `-parallel` and the following argument should be the desired number for concurrent requests.

- Example: `./main -parallel 3 http://google.com`

If you want the program to be executed sequentially, pass 1 as argument for the `-parallel` flag.

If no URL is provided, the program will return without any outputs.

## Code structure

The implementation of the requests uses a code design based on Clean Architecture. The code structure is:
```bash
.
├── core         # Folder in which the main implementation was done
├──── gateways   # Interface for the request gateway
├──── usecases   # Implementation of the tool use case (hashing the response of the requests and calling the requests gateway)
├── gateways     # Concrete implementation of the requests gateway (where the request is made)
├── mocks        # Contains the interfaces mocks for testing 
├── main.go      # The main package
```

By using an interface for the Gateway and one interface for the Use Case, the code could be easily extensible. This could be useful, for example, if it was needed to start performing requests other than GET requests. Or, if one wanted to output something other than the hash of the responses.

- The implemented Use Case gets the response of a request and returns the MD5 hash of the response.
- The implemented Gateway performs a GET request and returns the request response.

## Tests

For the Use Case tests, an exernal library were used to accomodate for mocks. The test can be found in the `usecases` directory.

## Concurrency

In the main function, the code parses the instruction to get the desired concurrency level. If the desired concurrency is 1, the code will perform the requests sequentially.

If the concurrency is above 1, a Go channel will be created and a function to perform the requests will be called in Go routines, receiving the URL information from the Go channel.

The maximum concurrency allowed is of 50 Go routines.
