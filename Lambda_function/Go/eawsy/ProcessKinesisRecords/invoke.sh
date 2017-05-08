aws lambda  invoke \
--invocation-type RequestResponse \
--function-name ProcessKinesisStream_easwy \
--region us-west-2 \
--payload file://./input.txt \
--profile default \
outputfile.txt
