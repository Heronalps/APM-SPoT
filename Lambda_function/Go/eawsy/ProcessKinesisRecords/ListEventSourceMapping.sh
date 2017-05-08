aws lambda list-event-source-mappings \
--region us-west-2 \
--function-name ProcessKinesisStream_easwy \
--event-source arn:aws:kinesis:us-west-2:396803809648:stream/MyKinesisStream \
--debug
