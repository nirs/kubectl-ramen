# SPDX-FileCopyrightText: The RamenDR authors
# SPDX-License-Identifier: Apache-2.0

# We can run as kubectl-ramen or oc-ramen.
HOST := kubectl

prog := $(HOST)-ramen
cover := cover.out
output := $(cover) $(prog)

all:
	go build -o $(prog)

test: reuse lint quick-tests

lint:
	golangci-lint run

quick-tests:
	go test -cover -coverprofile=$(cover) ./...

cover:
	go tool cover -html=$(cover)

reuse:
	reuse lint

clean:
	rm -f $(output)
