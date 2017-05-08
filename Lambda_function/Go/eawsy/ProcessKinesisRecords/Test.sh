aws kinesis put-record \
--stream-name MyKinesisStream \
--data "This is a test from RaceLab. final" \
--partition-key shardId-000000000000 \
--region us-west-2 \
--profile default
