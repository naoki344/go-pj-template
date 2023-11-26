## デプロイ手順

### デプロイ用のS3バケットを作成しておく
```
aws s3 mb s3://go-template-dev-cost-management

```


### ビルド&デプロイ
```
sam build

sam package \
    --output-template-file packaged.yaml \
    --s3-bucket go-template-dev-cost-management

sam deploy \
    --template-file packaged.yaml \
    --stack-name NotifyBillingToSlack \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides SlackWebhookUrl=${SLACK_WEBHOOK_URL}

```
