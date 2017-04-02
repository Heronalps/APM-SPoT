aws lambda create-function \
  --region us-west-2 \
  --function-name HelloPython \
  --zip-file fileb://hello_python.zip \
  --role arn:aws:iam::444253827891:role/lambda_basic_execution \
  --handler hello_python.my_handler \
  --runtime python2.7 \
  --timeout 15 \
  --memory-size 512
