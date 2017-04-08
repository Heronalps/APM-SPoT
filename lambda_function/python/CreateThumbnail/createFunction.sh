aws lambda create-function \
--region us-west-2 \
--function-name CreateThumbnail \
--zip-file fileb://~/CreateThumbnail.zip \
--role arn:aws:iam::396803809648:role/AWSLambdaRole \
--handler CreateThumbnail.handler \
--runtime python2.7 \
--profile default \
--timeout 10 \
--memory-size 1024
