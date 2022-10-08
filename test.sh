#!/usr/bin/env bash

coverage=$(go test ./modules/* -covermode count -coverprofile coverage | grep -v 'coverage: 100.0%')
if [[ "$coverage" != "" ]]; then
  go tool cover -func=coverage | grep -v "100.0%"
fi
