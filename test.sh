#!/usr/bin/env bash

pkgs=github.com/akromio/go-expected/...
coverage=$(go test $pkgs -covermode=count -coverprofile coverage | grep -v 'coverage: 100.0%')
if [[ "$coverage" != "" ]]; then
  go tool cover -func=coverage | grep -v "100.0%"
fi