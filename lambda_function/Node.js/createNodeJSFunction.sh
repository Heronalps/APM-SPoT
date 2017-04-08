aws lambda create-function \
  --region us-west-2 \
  --function-name ReturnBucketName \
  --zip-file 'fileb://../lambda_function/Node.js/Test_Environment_Variables.zip' \
  --role arn:aws:iam::396803809648:role/AWSLambdaRole \
  --environment Variables={S3_BUCKET=Test} \
  --handler index.handler \
  --runtime nodejs6.10 \
  --publish \
