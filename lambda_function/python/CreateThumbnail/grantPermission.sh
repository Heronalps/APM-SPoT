aws lambda add-permission \
--function-name CreateThumbnail \
--region us-west-2 \
--statement-id 1 \
--action "lambda:InvokeFunction" \
--principal s3.amazonaws.com \
--source-arn arn:aws:s3:::spotbucketexample \
--source-account 396803809648 \
--profile default
