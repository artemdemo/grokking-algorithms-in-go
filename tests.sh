#!/bin/sh

echo "binary_search/*"
go test ./binary_search -test.v -cover

echo "quicksort/*"
go test ./quicksort -test.v -cover

echo "tree/*"
go test ./tree -test.v -cover
