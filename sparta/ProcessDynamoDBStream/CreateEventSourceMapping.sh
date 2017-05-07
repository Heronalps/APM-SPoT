aws lambda create-event-source-mapping \
--region us-west-2 \
--function-name Process-DynamoDB-Stream-G-mainprocessDynamoDBStrea-4Z8O14K5B8WG \
--event-source arn:aws:dynamodb:us-west-2:396803809648:table/GoExample/stream/2017-05-07T00:47:26.013 \
--batch-size 100 \
--starting-position TRIM_HORIZON \
--profile default
