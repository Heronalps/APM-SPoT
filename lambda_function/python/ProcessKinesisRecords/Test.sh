aws kinesis put-record \
--stream-name MyKinesisStream \
--data "This is a test. final" \
--partition-key shardId-000000000000 \
--region us-west-2 \
--profile default
