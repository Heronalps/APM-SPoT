from boto3 import client as boto3_client
from datetime import datetime
import json

lambda_client = boto3_client('lambda')

def lambda_handler(event, context):
    msg = {"key":"new_invocation", "at": datetime.now()}
    invoke_response = lambda_client.invoke(FunctionName="Latency_dest",
                                           InvocationType='RequestResponse,
                                           Payload=json.dumps(msg))
    return invoke_response['Payload'].read()
