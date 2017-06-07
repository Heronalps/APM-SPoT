for (( a=0; a<=1000; a++ ))
do
    aws lambda invoke \
    --function-name Latency_Base \
    --region us-west-2 \
    --profile default \
    output.txt

    #cat output.txt >> dataset.csv
    #echo "," >> dataset.csv
    EPOCH=$(<output.txt)

    TRACEIDS=$(aws xray get-trace-summaries --start-time $(($EPOCH-10)) --end-time $(($EPOCH)) --query 'TraceSummaries[*].Id' --output text)
    aws xray batch-get-traces --trace-ids $TRACEIDS --query 'Traces[*]' >> xray.json
    echo "," >> xray.json
done
