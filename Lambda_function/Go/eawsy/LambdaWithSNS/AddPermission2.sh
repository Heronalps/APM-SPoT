aws lambda add-permission \
    --function-name lambdaWithSNS_easwy \
    --statement-id sns-x-account \
    --action "lambda:InvokeFunction" \
    --principal sns.amazonaws.com \
    --source-arn arn:aws:sns:us-west-2:444253827891:lambda-x-account
