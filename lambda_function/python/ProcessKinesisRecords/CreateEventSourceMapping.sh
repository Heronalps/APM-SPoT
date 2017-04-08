aws lambda create-event-source-mapping \
--region us-west-2 \
--function-name ProcessKinesisRecords \
--event-source  arn:aws:kinesis:us-west-2:396803809648:stream/MyKinesisStream \
--batch-size 100 \
--starting-position TRIM_HORIZON \
--profile default
