# SPDX-FileCopyrightText: The RamenDR authors
# SPDX-License-Identifier: Apache-2.0

PROG := kubectl-ramen

cover := cover.out
output := $(cover) $(PROG)

all: $(PROG)

test: reuse quick-tests

quick-tests:
	go test -cover -coverprofile=$(cover) ./...

cover:
	go tool cover -html=$(cover)

reuse:
	reuse lint

$(PROG):
	go build -o $(PROG)

clean:
	rm -f $(output)
