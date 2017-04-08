aws elasticache create-cache-cluster \
    --cache-cluster-id ClusterForLambdaTest \
    --cache-node-type cache.t2.micro \
    --engine memcached \
    --security-group-ids sg-668f4e1d \
    --num-cache-nodes 1
