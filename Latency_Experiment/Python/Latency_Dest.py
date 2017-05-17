import os
import sys
import logging
import rds_config
import pymysql
import time
#rds_settings
rds_host  = "latencyexperiment.cxfpxjlukepw.us-west-2.rds.amazonaws.com:3306"
name = rds_config.db_username
password = rds_config.db_password
db_name = rds_config.db_name

logger = logging.getLogger()
logger.setLevel(logging.INFO)

try:
    conn = pymysql.connect(rds_host, user=name, passwd=password, db=db_name, connect_timeout=5)
except:
    logger.error("ERROR: Unexpected error: Could not connect to MySql instance.")
    sys.exit()

logger.info("SUCCESS: Connection to RDS mysql instance succeeded")

def lambda_handler(event, context):

    time_diff = time.time() - event['timestamp']

    with conn.cursor() as cur:
        insert_stmt = "insert into invoke_time (time_diff) values(%s)"
        cur.execute(insert_stmt, time_diff)
        conn.commit()
