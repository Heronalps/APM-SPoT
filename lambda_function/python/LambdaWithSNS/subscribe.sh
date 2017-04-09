aws sns subscribe --profile heronalps \
    --topic-arn arn:aws:sns:us-west-2:396803809648:lambda-x-account \
    --protocol lambda \
    --notification-endpoint arn:aws:lambda:us-west-2:444253827891:function:SNS-X-Account
