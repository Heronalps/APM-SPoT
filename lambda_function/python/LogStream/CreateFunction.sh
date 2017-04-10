aws lambda create-function \
    --function-name GetMyLogStream \
    --runtime python2.7 \
    --role arn:aws:iam::396803809648:role/AWSLambdaRole \
    --handler get_my_log_stream.get_my_log_stream \
    --timeout 60 \
    --memory-size 128 \
    --zip-file fileb://./getMyLogStream.zip
