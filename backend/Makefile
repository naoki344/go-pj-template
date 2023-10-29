.PHONY: help test
.DEFAULT_GOAL := help


install:
	# TODO: nodeのバージョンチェック
	# TODO: goのインストールチェック
	npm install
	go install -v github.com/ogen-go/ogen/cmd/ogen@latest
	go install -v github.com/google/wire/cmd/wire@latest
	go install -v github.com/k0kubun/sqldef/cmd/mysqldef@latest

migrate:
	mysqldef -u ${DB_USERNAME} -p '${DB_PASSWORD}' -h ${DB_HOST} -P ${DB_PORT} ${DB_NAME}  < ./migrations/mysql/schemas.sql

update-pkg:
	go mod tidy

deploy-setting:
	cdk bootstrap

build-wire:
	echo "build wire file"
	wire

build-ogen:
	echo "build ogen file"
	ogen -package ogen -target ogen -clean ./openapi/openapi.yaml

build-openapi:
	swagger-cli bundle -o openapi/openapi.yaml -t yaml openapi/docsroot.yaml -r

setup-all: setup-install migrate build-openapi build-ogen build-wire

build-all: build-openapi build-ogen build-wire

deploy: build-all
	cdk deploy

test:
	go test

help:
	cat README.md
