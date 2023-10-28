# 初期設定
- node をインストールする
- go をインストールする
- 以下のコマンドを実行する
```
npm install

go install -v github.com/ogen-go/ogen/cmd/ogen@latest
```

# ogenファイルの生成
```
ogen -package ogen -target ogen -clean ./openapi/openapi.yaml
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

# Update Go Package
```
go mod tidy
```

