aws lambda create-alias \
--region us-west-2 \
--function-name HelloPython \
--description "sample alias" \
--function-version "\$LATEST" \
--name DEV \
--profile default
