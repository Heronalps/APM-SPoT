aws lambda invoke \
--invocation-type RequestResponse \
--function-name Process-DynamoDB-Stream-G-mainprocessDynamoDBStrea-4Z8O14K5B8WG \
--region us-west-2 \
--payload file://./input.txt \
--profile default \
outputfile.txt
