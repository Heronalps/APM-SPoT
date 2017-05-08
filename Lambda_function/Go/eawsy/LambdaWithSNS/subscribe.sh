aws sns subscribe \
    --topic-arn arn:aws:sns:us-west-2:444253827891:lambda-x-account \
    --protocol lambda \
    --notification-endpoint arn:aws:lambda:us-west-2:396803809648:function:lambdaWithSNS_easwy
