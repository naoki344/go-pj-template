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
	mysqldef -u ${DB_USERNAME} -p ${DB_PASSWORD} -h ${DB_HOST} -P ${DB_PORT} ${DB_NAME}  < ./migrations/mysql/schemas.sql

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

deploy: build-wire build-ogen
	cdk deploy

setup-all: setup-install migrate build-wire build-ogen

help:
	cat README.md
