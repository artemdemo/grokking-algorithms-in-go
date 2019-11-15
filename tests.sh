#!/bin/sh

#echo "binary_search/*"
#go test ./binary_search -test.v -cover
#
#echo "quicksort/*"
#go test ./quicksort -test.v -cover
#
#echo "tree/*"
#go test ./tree -test.v -cover

# In order to run test with coverage,
# you need to use following commands:
go test ./tree ./binary_search ./quicksort -test.v -cover -coverprofile="coverage.out"
go tool cover -html="coverage.out"
