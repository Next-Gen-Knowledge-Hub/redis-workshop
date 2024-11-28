# /bin/bash

echo "Starting redis instances in docker..."

for i in {1..6}; do
    docker run -d \
        --name redis-$i \
        --network host \
        -v redis-data-$i:/data \
        redis \
        --port 700$((i - 1)) \
        --cluster-enabled yes \
        --cluster-config-file nodes.conf \
        --cluster-node-timeout 5000 \
        --appendonly yes
done

# create cluster

redis-cli -p 7000 --cluster create \
    127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 \
    --cluster-replicas 1

echo "Done ;)"
