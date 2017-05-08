aws lambda create-function \
--region us-west-2 \
--function-name ProcessDynamoDBStream_easwy \
--zip-file fileb://./handler.zip \
--role arn:aws:iam::396803809648:role/AWSLambdaRole \
--handler handler.Handle \
--runtime python2.7 \
--profile default \
--timeout 10 \
--memory-size 1024
