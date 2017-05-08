aws lambda list-event-source-mappings \
--region us-west-2 \
--function-name ProcessDynamoDBStream_easwy \
--event-source arn:aws:dynamodb:us-west-2:396803809648:table/GoExample/stream/2017-05-07T00:47:26.013 \
--profile default
