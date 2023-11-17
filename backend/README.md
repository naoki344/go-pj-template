# 初期設定
- node(.node-versionに記載のバージョン)をインストールする
- go をインストールする
- 以下のコマンドを実行する
```
	make setup
```

# ogenファイルの生成
```
ogen -package ogen -target ogen -clean ./openapi/openapi.yaml
 ```

# DBのmigration(mysqldef)を利用
```
mysqldef -u ${DB_USERNAME} -p ${DB_PASSWORD} -h ${DB_HOST} -P ${DB_PORT} ${DB_NAME}  < ./migrations/mysql/schemas.sql
```

## Deploy Stack


```bash
cdk bootstrap

cdk deploy
```

## Useful commands

- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
- `go test` run unit tests
- ` mysql -u ${DB_USERNAME} -p  -h ${DB_HOST} -P ${DB_PORT}` connect mysql




# 開発ツール

## 一覧
| ツール名 | 目的 | 設定ファイル |
|:--:|:--:|:--:|
| [pre-commit](https://pre-commit.com/) | Git hooks | .pre-commit-config.yaml |
| [direnv](https://direnv.net) | 環境変数の管理 | .envrc |

## pre-commitの設定
[pre-commit](https://pre-commit.com/)利用している。以下のコマンドでpre-commitをインストールし設定を適応する。
```sh
brew install pre-commit

pre-commit install
```


## テストファイルのテンプレート生成
```
gotests -w -all -template testify internal/usecase/customer/customer.go
```
