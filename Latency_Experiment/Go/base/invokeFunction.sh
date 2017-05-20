for (( a=1; a<=5; a++ ))
do
    aws lambda invoke \
    --invocation-type RequestResponse \
    --function-name Latency_base_Go \
    --region us-west-2 \
    --profile default \
    output.txt

    cat output.txt >> dataset.csv
    echo "," >> dataset.csv
done
