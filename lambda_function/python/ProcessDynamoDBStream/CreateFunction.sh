aws lambda create-function \
--region us-west-2 \
--function-name ProcessDynamoDBStream \
--zip-file fileb://./ProcessDynamoDBStream.zip \
--role arn:aws:iam::396803809648:role/AWSLambdaRole \
--handler ProcessDynamoDBStream.lambda_handler \
--runtime python2.7 \
--profile default
