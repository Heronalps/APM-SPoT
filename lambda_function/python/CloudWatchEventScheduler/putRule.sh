aws events put-rule \
--name MyCloudWatchRule \
--schedule-expression 'rate(1 minute)'
