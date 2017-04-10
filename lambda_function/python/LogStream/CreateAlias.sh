aws lambda create-alias \
 --function-name GetMyLogStream \
 --name "Prod" \
 --function-version "\$LATEST"
