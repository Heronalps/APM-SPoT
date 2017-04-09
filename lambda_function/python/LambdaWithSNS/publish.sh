aws sns publish \
    --topic-arn arn:aws:sns:us-west-2:396803809648:lambda-x-account \
    --message file://./message.txt \
    --subject Test
