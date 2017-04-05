aws lambda update-function-configuration \
  --function-name 'ReturnBucketName' \
  --region us-west-2 \
  --environment Variables={S3_BUCKET=Prod} \
  --profile default \
