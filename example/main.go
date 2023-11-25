package main

import (
	"fmt"
	"os"

	"github.com/gosom/secretstash"
	"github.com/gosom/secretstash/awssecret"
	"github.com/gosom/secretstash/envprovider"
)

// Make sure that you have set the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
// environment variables before running this example or
// the AWS Secrets Manager provider will fail.

func main() {
	// Initialize a new SecretStash with the environment provider and the AWS
	// Secrets Manager provider.
	stash, err := secretstash.New(
		envprovider.New(),
		awssecret.New("eu-north-1"),
	)
	if err != nil {
		panic(err)
	}

	os.Setenv("foo", "value")

	fooVal, err := stash.GetSecret("foo")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fooVal) // prints "value"
	}

	// if you have a secret named "bar" in AWS Secrets Manager, this will printed
	barVal, err := stash.GetSecret("bar")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(barVal)
	}
}
