# goenvgen [![Go](https://github.com/potproject/goenvgen/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/potproject/goenvgen/actions/workflows/go.yml) [![codecov](https://codecov.io/gh/potproject/goenvgen/branch/master/graph/badge.svg?token=FWVQKED8BO)](https://codecov.io/gh/potproject/goenvgen) [![Go Report Card](https://goreportcard.com/badge/github.com/potproject/goenvgen)](https://goreportcard.com/report/github.com/potproject/goenvgen)

Automatically generate Go Code from dotenv files.

This cli automatically determines the type information from the dotenv information and generates a Go package that can retrieve it.
`os.GetEnv` the code becomes complicated because of the lack of type information.
This package eliminates the need to use `os.GetEnv`.

## Usage

### Installation

```
# Go >= 1.17:
$ go install github.com/potproject/goenvgen

# Go < 1.17:
$ go get github.com/potproject/goenvgen
```

### Generate Code

Add your application configuration to your `.env` file in the root of your project.

```
BASE_HOST=localhost
BASE_PORT=8080
ADMIN_IDS=123,234,345,456

USER_JSON={"Alice": {"ID": 100}, "Bob": {"ID": 200}}
```

The following code will be generated automatically Go Package.

```
$ goenvgen -p envgen .env
Generated envgen Packages.

$ tree envgen
envgen
├── USER_JSON.go
└── envgen.go
```

### About generated package

This package itself does not read environment variables from dotenv, so [joho/godotenv](https://github.com/joho/godotenv) is required if you want to use .env file.

This is because it may control whether the dotenv file is loaded in the production and development environments.

#### Example Codes:

```go
package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"[Your package path]/envgen"
)

func main() {
	// Loading .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup envgen package from environment variables
	err = envgen.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Getting Environment Variables
	host := envgen.Get().BASE_HOST()
	port := envgen.Get().BASE_PORT()
	fmt.Printf("Running to http://%s:%d/", host, port)
  
	// Slice Type
	ids := envgen.Get().ADMIN_IDS()
	for _, id := range ids {
		fmt.Printf("ID: %d \r\n", id)
	}

	// JSON Type
	user := envgen.Get().USER_JSON()
	fmt.Printf("Bod ID: %d \r\n", user.Bob.ID)
}

```

#### Results:
```bash
$ go run main.go
Running to http://localhost:8080/
ID: 123 
ID: 234
ID: 345
ID: 456
Bob ID: 200
```

Enjoy Development!

## Details

```
Usage of envgen:
  -p string
        the name of the package for the generated code (default "envgen")
  -r string
        Required Type setting (example: "-r ENV_REQ_BOOL,ENV_REQ_STRING")
  -t string
        Manually type definition setting (example: "-t ENV_BOOL=bool,ENV_S_INT=[]int")
```

### Automatic determination of type

At the time of code generation, the type is automatically assigned from the values in the .env file.

Automatic generation is generated by the following rule.

|  value  |  is Type ...  | Example |
| ---- | ---- | ---- |
| JSON strings  |  Generate struct for [ChimeraCoder/gojson](https://github.com/ChimeraCoder/gojson)  | '{"Alice": {"ID": 100}, "Bob": {"ID": 200}}' |
| `true` or `false`  |  `boolean`  | 'true' |
| integral number  |  `int64`  | '123' |
| floating point number | `float64` | '1.1' |
| `true` or `false` with comma | `[]boolean` | 'true,false' |
| integral number with comma | `[]int64` | '100,200' |
| floating point number with comma | `[]float64` | '0.1,0.2' |
| charactor with comma | `[]string` | 'string1,string2' |
| otherwise | `string` | 'string' |

### Manual determination of type

The type can be set manually by specifying the `-t` option when generating the code.

The types that can be set are `bool`, `int`, `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64`, `float32`, `float64`, `interface`, and `string`. or their slice types.

#### Example:
```
goenvgen -p envgen -t BASE_PORT=uint16,ADMIN_IDS=[]int
```

### Required Environment Variable

You can make the value required by specifying the `-r` option when generating the code.

#### Example:
```
BASE_HOST=
BASE_PORT=8080
```

```
goenvgen -p envgen -r BASE_HOST,BASE_PORT
```

```bash
$ go run .\main.go
BASE_HOST is required
```

## Development & Contributing

Contributions are most welcome!

```bash
$ git clone https://https://github.com/potproject/goenvgen.git

# Code Generating Test
$ cd test
$ vi .env.test # Change .env.test File
$ go run main.go
$ go test ./envgentest/
```

## License

MIT License
