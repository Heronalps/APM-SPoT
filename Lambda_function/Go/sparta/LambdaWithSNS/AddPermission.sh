aws sns add-permission --profile heronalps \
    --region us-west-2 \
    --topic-arn arn:aws:sns:us-west-2:444253827891:lambda-x-account \
    --label lambda-access \
    --aws-account-id 396803809648 \
    --action-name Subscribe ListSubscriptionsByTopic Receive
