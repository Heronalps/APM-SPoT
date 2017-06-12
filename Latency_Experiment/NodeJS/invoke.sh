for (( a=1; a<=2000; a++ ))
do
    aws lambda invoke \
    --function-name xray_base_NodeJS \
    --region us-west-2 \
    --profile default \
    output.txt

    cat output.txt | sed -e 's/^.*"\(.*\)",*/\1/' >> dataset.dat
    #echo "" >> dataset.dat
done
