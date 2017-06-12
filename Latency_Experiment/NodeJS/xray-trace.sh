for (( a=0; a<=1000; a++ ))
do
    aws lambda invoke \
    --function-name xray_base_NodeJS \
    --region us-west-2 \
    --profile default \
    output.txt

    EPOCH=$(date +%s)

    TRACEIDS=$(aws xray get-trace-summaries --start-time $(($EPOCH-5)) --end-time $(($EPOCH)) --query 'TraceSummaries[*].Id' --output text)
    aws xray batch-get-traces --trace-ids $TRACEIDS --query 'Traces[*].Duration' --output text >> xray.dat
    #echo "," >> xray.json
done
