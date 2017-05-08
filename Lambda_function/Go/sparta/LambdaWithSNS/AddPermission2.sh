aws lambda add-permission \
    --function-name SNS-X-Account-mainlambdahandlerLambdab419f67c0e60a-UM4AJA396X7O \
    --statement-id sns-x-account \
    --action "lambda:InvokeFunction" \
    --principal sns.amazonaws.com \
    --source-arn arn:aws:sns:us-west-2:444253827891:lambda-x-account
