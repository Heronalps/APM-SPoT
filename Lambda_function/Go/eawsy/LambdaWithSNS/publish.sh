aws sns publish --profile heronalps \
    --topic-arn arn:aws:sns:us-west-2:444253827891:lambda-x-account \
    --message file://./message.txt \
    --subject Test
