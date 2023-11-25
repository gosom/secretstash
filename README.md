# SecretStash

SecretStash is a Go library for managing secrets. It works with AWS Secrets Manager and environment variables. You can add more secret providers too.

## Features

- Works with many secret providers.
- Easy to add more providers.
- Simple to use.

## Getting Started

### Prerequisites

- You need Go on your computer.
- If using AWS Secrets Manager, set up AWS credentials.

### Installation

```bash
go get github.com/gosom/secretstash
```

### Usage

Here's how to use it:

```go
package main

import (
    "fmt"
    "os"

    "github.com/gosom/secretstash"
    "github.com/gosom/secretstash/awssecret"
    "github.com/gosom/secretstash/envprovider"
)

func main() {
    // Start SecretStash with environment and AWS providers
    stash, err := secretstash.New(
        envprovider.New(),
        awssecret.New("eu-north-1"),
    )
    if err != nil {
        panic(err)
    }

    // Example: Getting a secret from an environment variable
    os.Setenv("foo", "value")
    fooVal, err := stash.GetSecret("foo")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(fooVal) // prints "value"
    }

    // Example: Getting a secret from AWS Secrets Manager
    // Make sure the secret "bar" is in AWS Secrets Manager
    barVal, err := stash.GetSecret("bar")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(barVal)
    }
}
```

## Contributing

Feel free to help! You can add new features, fix bugs, or suggest ideas.

## License

This project is licensed under the MIT License.

