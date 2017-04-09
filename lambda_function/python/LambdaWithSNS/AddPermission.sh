aws sns add-permission --profile default \
    --region us-west-2 \
    --topic-arn arn:aws:sns:us-west-2:396803809648:lambda-x-account \
    --label lambda-access \
    --aws-account-id 444253827891 \
    --action-name Subscribe ListSubscriptionsByTopic Receive
