#!/bin/bash
apt-get update -y
apt-get install -y git ntp python3-pip
pip3 install fnv
git config --global --unset https.proxy
wget https://storage.googleapis.com/golang/go1.8.1.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.8.1.linux-amd64.tar.gz
echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile
echo "export PATH=\$PATH:\$HOME/go/bin" >> /home/ubuntu/.profile
mkdir -p /home/ubuntu/go
export GOPATH=/home/ubuntu/go
source /etc/profile
go get -u github.com/fasthall/gochariots/...

chown -R ubuntu /home/ubuntu/go
chgrp -R ubuntu /home/ubuntu/go

echo "server ntp1.ucsb.edu iburst" > /etc/ntp.conf
echo "server ntp2.ucsb.edu" >> /etc/ntp.conf

echo "169.231.235.109	app1" >> /etc/hosts
echo "169.231.235.32	controller1" >> /etc/hosts
echo "169.231.235.49	batcher1" >> /etc/hosts
echo "169.231.235.50	filter1" >> /etc/hosts
echo "169.231.235.53	queue1" >> /etc/hosts
echo "169.231.235.55	maintainer1" >> /etc/hosts
echo "169.231.235.59	app2" >> /etc/hosts
echo "169.231.235.64	controller2" >> /etc/hosts
echo "169.231.235.71	batcher2" >> /etc/hosts
echo "169.231.235.74	filter2" >> /etc/hosts
echo "169.231.235.75	queue2" >> /etc/hosts
echo "169.231.235.92	maintainer2" >> /etc/hosts
echo "169.231.235.142   indexer1" >> /etc/hosts
echo "169.231.235.161   indexer2" >> /etc/hosts