aws lambda invoke \
--invocation-type RequestResponse \
--function-name ProcessDynamoDBStream \
--region us-west-2 \
--payload file://./input.txt \
--profile default \
outputfile.txt
