# EnvServer-Rodina

This is a simple webserver with two endpoints.

## __Endpoints:__
- `/env:` returns all the environment variables of the server as a JSON object.
- `/env/<key>:` returns the value of the specified environment variable key within the http request as a plain text.


## __Manual:__

1. Clone the repository:
```sh
$ git clone https://github.com/codescalersinternships/EnvServer-Rodina.git
```
2. Go to the repository directory:
```sh
$ cd EnvServer-Rodina
```
3. Install dependencies:
```sh
$ go get -d ./...
```
4. Build the package:
```sh
$ go build -o "bin/app" cmd/main.go
```
 ### __How to use without docker?__

1. Run the app as follows:
```sh
$ ./bin/app -p portNumber
```
Note: you need to move the binary to any `$PATH` directory first.

2. Get all environment variables as follows:
```sh
$ http://localhost:8080/env
```
3. Get the specific environment value for a key as follows:
```sh
$ http://localhost:8080/env/key
```

### __How to use docker?__

Run the app using docker as follows:
```sh
$ docker-compose up -d
```

### __How to test?__

Run all the tests as follows: 
```sh
go test ./....
```
If all tests pass on, the result should show that the tests were successful as follows:
```sh
PASS
ok      github.com/codescalersinternships/EnvServer-Rodina/internal   0.003s
```
If any test fails, the output will tell you which test failed.
