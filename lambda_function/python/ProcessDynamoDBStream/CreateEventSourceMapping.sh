aws lambda create-event-source-mapping \
--region us-west-2 \
--function-name ProcessDynamoDBStream \
--event-source arn:aws:dynamodb:us-west-2:396803809648:table/MyStreamTable/stream/2017-04-08T23:32:37.022 \
--batch-size 100 \
--starting-position TRIM_HORIZON \
--profile default
