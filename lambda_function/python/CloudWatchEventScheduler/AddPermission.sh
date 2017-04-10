aws lambda add-permission \
    --statement-id 'statement-id' \
    --action 'lambda:InvokeFunction' \
    --principal events.amazonaws.com \
    --source-arn arn:aws:events:us-west-2:396803809648:rule/CheckWebsiteScheduledEvent \
    --function-name CheckAmazon \
    --region us-west-2
