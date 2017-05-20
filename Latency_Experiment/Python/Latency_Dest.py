import time

def lambda_handler(event, context):

    time_diff = time.time() - event['timestamp']
    return time_diff
