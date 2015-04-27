NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
ECHO=echo

DEPS=$(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

all: deps format test build

build:
	@mkdir -p bin/
	@$(ECHO) "$(OK_COLOR)==> Building...$(NO_COLOR)"
	@go install

format:
	@$(ECHO) "$(OK_COLOR)==> Formatting...$(NO_COLOR)"
	go fmt

deps:
	@$(ECHO) "$(OK_COLOR)==> Installing dependencies...$(NO_COLOR)"
	@go get -d -v
	@echo $(DEPS) | xargs -n1 go get -d

clean:
	@rm -rf bin/ pkg/ src/

test:
	@$(ECHO) "$(OK_COLOR)==> Running Tests...$(NO_COLOR)"
	go test ./...

updatedeps:
	@$(ECHO) "$(OK_COLOR)==> Updating all dependencies$(NO_COLOR)"
	@go get -d -v -u ./..
	@echo $(DEPS) | xargs -n1 go get -d -u

.PHONY: all clean deps format test updatedeps build
