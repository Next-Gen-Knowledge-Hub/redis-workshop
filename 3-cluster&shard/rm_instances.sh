# /bin/bash

echo "Starting removing redises instances in docker..."

for i in {1..6}; do
    docker stop redis-$i
    docker rm redis-$i

    docker volume rm redis-data-$i
done

echo "Done"
