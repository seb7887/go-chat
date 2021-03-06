BIN			= $(CURDIR)/bin
PKGS		= $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))

GO			= go
Q				= $(if $(filter 1,$V),,@)
M				= $(shell printf "\033[34;1m▶\033[0m")

$(BIN):
				@mkdir -p $@

.PHONY: all
all: fmt $(BIN) ; $(info $(M) building executable) @ ## Build binary
																$Q go build -o ./bin/challenge -race ./cmd/main.go

.PHONY: fmt
fmt: ; $(info $(M) running gofmt...) @ ## Run gofmt on all source files
																$Q $(GO) fmt $(PKGS)

.PHONY: clean
clean:
	@rm -rf $(BIN)

.PHONY: run
run:
	$Q $(GO) run ./cmd/main.go

.PHONY: migrate
migrate: ; $(info $(M) migrating schema...) @ ## Run database migration file
																$Q sqlite3 chat.db '.read ./migrations/init.sql'

.PHONY: test
test:
	$Q $(GO) test -cover ./...
