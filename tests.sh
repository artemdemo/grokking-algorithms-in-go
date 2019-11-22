#!/bin/bash

if [[ $* == *--html* ]]
then
  go test ./binary_search ./quicksort ./tree ./graph -test.v -cover -coverprofile="coverage.out"
  go tool cover -html="coverage.out"
else
  echo "binary_search/*"
  go test ./binary_search -test.v -cover

  echo "quicksort/*"
  go test ./quicksort -test.v -cover

  echo "tree/*"
  go test ./tree -test.v -cover

  echo "graph/*"
  go test ./graph -test.v -cover
fi
