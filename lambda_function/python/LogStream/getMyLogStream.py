import time
def get_my_log_stream(event, context):
    """
    "Insert Infinite loop"
    while 1:
        print(event)
    """

    """
    "Memory Leak"
    "Configure event {"seed" : 2}"
    for key, value in event.items():
        seed = event["seed"]
        event[key * seed] = value
    get_my_log_stream(event, context)
    """

    """
    "Exit with an error"
    #raise IndexError
    #print(a) #reference a undefined variable - Raise NameError
    #for #Syntax error
    #sys.exit("Intentional Exit") #Force Exit
    """


    print("Log stream name:", context.log_stream_name)
    print("Log group name:",  context.log_group_name)
    print("Request ID:",context.aws_request_id)
    print("Mem. limits(MB):", context.memory_limit_in_mb)
    # Code will execute quickly, so we add a 1 second intentional delay so you can see that in time remaining value.
    time.sleep(1)
    print("Time remaining (MS):", context.get_remaining_time_in_millis())
