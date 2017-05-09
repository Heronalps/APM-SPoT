# APM-SPoT

Application Performance Monitor tool for Serverless Platform of Things project.

This is a part of SPoT project for which RACELab collaborates with Huawei Cloud.

Project Progress:

4/7/17
Wrote a set of functions in Python that communicate other AWS services. Monitored functions and construct metrics chart by CloudWatch. Detailed report.

4/14/17
Prototype AWS lambda by docker (both on local machine and play-with-docker) and FaaS framework.

docker image: docker run -p 4000:80 heronalps/first_repo:version1.0

4/21/17
Wrote experimental Golang functions used to test AWS services mechanism.		

4/28/17
Test Golang AWS framework Sparta with the monitoring functions to find the duration and memory limitations of spawned child_process of Go lambda function.

5/5/17
Implemented getLogStream, LambdaWithSNS, ProcessDynamoDBStream, ProcessKinesisStream functions in Go under sparta and eawsy framework.
