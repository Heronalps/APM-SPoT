aws lambda add-permission --profile heronalps \
    --function-name SNS-X-Account \
    --statement-id sns-x-account \
    --action "lambda:InvokeFunction" \
    --principal sns.amazonaws.com \
    --source-arn arn:aws:sns:us-west-2:396803809648:lambda-x-account
