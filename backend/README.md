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
