aws lambda create-function \
--function-name CreateTableAddRecordsAndRead  \
--region us-west-2 \
--zip-file fileb://../lambda_function/python/app.zip \
--role arn:aws:iam::396803809648:role/lambda-vpc-execution-role \
--handler app.handler \
--runtime python2.7  \
--timeout 30 \
--vpc-config SubnetIds="subnet-872c44e0","subnet-4423a00d","subnet-222c377a",SecurityGroupIds=sg-668f4e1d \
--memory-size 1024
