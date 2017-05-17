for (( a=1; a<=10000; a++ ))
do
    aws lambda invoke \
    --function-name Latency_Base \
    --region us-west-2 \
    --profile default \
    output.txt

    cat output.txt >> dataset2.csv
    echo "," >> dataset2.csv
done
