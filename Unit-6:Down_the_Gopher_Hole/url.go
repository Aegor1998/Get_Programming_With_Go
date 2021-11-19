/*
In the Go standard library, there’s a function to parse web addresses (see golang.org/pkg/net/url/#Parse).
Display the error that occurs when url.Parse is used with an invalid web address, such as one containing a
space: https://a b.com/.

Use the %#v format verb with Printf to learn more about the error. Then perform a *url.Error type assertion
to access and print the fields of the underlying structure.

Note:
A URL, or Uniform Resource Locator, is the address of a page on the World Wide Web.
*/
package main

import (
	"fmt"
	"net/url"
)

func main() {
	badURL()
	goodURL()
}

func badURL() {
	bad, err := url.ParseRequestURI("https://a b.com")
	if err != nil {
		fmt.Printf("\nError Encountered: %#v\n", err)
	} else {
		fmt.Println(bad, " is a correct url")
	}
}

func goodURL() {
	good, err := url.ParseRequestURI("https://www.google.com")
	if err != nil {
		fmt.Printf("\nError Encountered: %#v\n", err)
		fmt.Printf("\nError Encountered: %#v\n")
	} else {
		fmt.Println(good, " is a correct url")
	}
}
