# s7-agent-api-sdk

[![GoDoc](https://godoc.org/github.com/tmconsulting/s7-golang-sdk?status.svg)](https://godoc.org/github.com/tmconsulting/s7-golang-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/tmconsulting/s7-golang-sdk)](https://goreportcard.com/report/github.com/tmconsulting/s7-golang-sdk)

written on [Golang](https://golang.org/) SDK for [S7 Agent API](https://s7airlines.atlassian.net/wiki/spaces/GAAPI/pages/152122444/S7+Agent+API) (Application programming interface based on the IATA NDC standards with an advanced search for 
embedding in the online travel solutions). 

## Features

With this SDK you can perform:

* searchFlightsFlex operation
* searchFlights operation
* book operation
* reprice operation
* demandTickets operation
* servicePrice operation

## Compatible S7 Agent API versions

This package was tested against S7 Agent API v0.47

## Installation

It is go gettable

    $ go get github.com/tmconsulting/s7-golang-sdk

```go
package main

import (
	s7Sdk "github.com/tmconsulting/s7-golang-sdk"
)
...
```

## Usage examples

There are several usage examples in `./example` folder. Try it out.

## Tests

There are no tests yet. :( Feel free to help us to change this situation!

## Contribution

Contribution, in any kind of way, is highly welcome!
It doesn't matter if you are not able to write code.
Creating issues or holding talks and help other people to use 
[s7-agent-api-sdk](https://github.com/tmconsulting/s7-golang-sdk) is contribution, too!

A few examples:

* Correct typos in the README / documentation
* Reporting bugs
* Implement a new feature or endpoint
* Sharing the love if like to use [s7-agent-api-sdk](https://github.com/tmconsulting/s7-golang-sdk) and help people 
to get use to it

If you are new to pull requests, checkout [Collaborating on projects using issues and pull requests / Creating a pull request](https://help.github.com/articles/creating-a-pull-request/).

## License

SDK is released under the [MIT License](./LICENSE).