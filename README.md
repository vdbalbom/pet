# pet
Project in purpose of increase my skills in good practices in programming.

## Using Docker (Recommended)

### To Test

``` make test ```

### To run

``` make run ```

### To lint*
``` make lint ```

## Not using Docker

#### To Test

``` go test ./... ```

#### To install pet

``` go install ./... ```

#### To run (need to install pet first)

``` ${GOPATH}/bin/pet ```

#### To install lint*
``` go get -u github.com/golang/lint/golint ```

#### To lint* (need to install lint first)
``` golint ./... ```

## *Lint:

source: <link> https://github.com/golang/lint </link>
