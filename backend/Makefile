.PHONY: help test
.DEFAULT_GOAL := help


install:
	# TODO: nodeのバージョンチェック
	# TODO: goのインストールチェック
	npm install
	go install -v github.com/ogen-go/ogen/cmd/ogen@latest
	go install -v github.com/google/wire/cmd/wire@latest
	go install -v github.com/k0kubun/sqldef/cmd/mysqldef@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest
	go install github.com/fe3dback/go-arch-lint@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/cweill/gotests/gotests@latest

migrate:
	mysqldef -u ${DB_USERNAME} -p '${DB_PASSWORD}' -h ${DB_HOST} -P ${DB_PORT} ${DB_NAME}  < ./migrations/mysql/schemas.sql

update-pkg:
	go mod tidy

deploy-setting:
	cdk bootstrap

build-wire:
	echo "build wire file"
	cd ./cmd/api/ && wire

build-ogen:
	echo "build ogen file"
	ogen -package ogen -target api/lib/ogen -clean ./api/openapi/dist/openapi/openapi.yaml

build-openapi:
	./node_modules/.bin/openapi-generator-cli generate -i ./api/openapi/docsroot.yaml -g openapi-yaml -o api/openapi/dist --generate-alias-as-model

setup-all: setup-install migrate build-openapi build-ogen build-wire

build-all: build-openapi build-ogen build-wire

deploy: build-all
	cd ./cdk/ && npx cdk deploy

test:
	go test ./test/...

build-mock:
	./test/scripts/genmock.sh


help:
	cat README.md

goclean:
	go clean --cache
	go clean --modcache

gofmt:
	gofmt -s -w .

gofumpt:
	@if command -v gofumpt &> /dev/null; then \
		go install mvdan.cc/gofumpt@latest; \
	fi
	gofumpt -w .


lint-init:
	@if command -v golangci-lint &> /dev/null; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@if command -v go-arch-lint &> /dev/null; then \
		go install github.com/fe3dback/go-arch-lint@latest; \
	fi

lint-run: lint-init
	golangci-lint run
	go-arch-lint check

lint-fix: lint-init
	golangci-lint run --fix
