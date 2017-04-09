aws lambda create-function --profile heronalps \
    --function-name SNS-X-Account \
    --runtime python2.7 \
    --role arn:aws:iam::444253827891:role/AWSLambdaRole \
    --handler lambda_handler.lambda_handler \
    --description "SNS X Account Test Function" \
    --timeout 60 \
    --memory-size 128 \
    --zip-file fileb://./LambdaWithSNS.zip
