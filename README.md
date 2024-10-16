# Redis Workshop
Redis workshop around redis features and how to setup and manage a simple redis cluster.

This repository contains following topics

1. [Setup a simple redis](https://github.com/Next-Gen-Knowledge-Hub/redis-workshop/tree/main/1-setup)
    - 1.1 Install redis docker image
    - 1.2 Connect to redis with `Go` client
    - 1.3 Do some simple stuf with it
2. [Redis internal](https://github.com/Next-Gen-Knowledge-Hub/redis-workshop/tree/main/2-internals)
    - 2.1 [Redis `Data-strucutes` and `Operations`](https://github.com/Next-Gen-Knowledge-Hub/redis-workshop/blob/main/2-internals/README_Datastructures.md)
      - 2.1.1 Strings (SET, GET...)
      - 2.2.2 Numerics (INCR, DECR...)
      - 2.2.3 Lists (RPUSH, RPOP, BLPOP...)
      - 2.2.4 Sets (SADD, SREM...)
      - 2.2.5 Hashs (HSET, HGET...)
      - 2.2.6 Sorted Sets (ZADD, ZREM...)
    - 2.2 [Redis Operations](https://github.com/Next-Gen-Knowledge-Hub/redis-workshop/blob/main/2-internals/README_Operations.md)
      - 2.2.1 Pun/Sub
      - 2.2.2 Basic transactions
      - 2.2.3 Locking
    - 2.3 [Scripting with `Lua`](https://github.com/Next-Gen-Knowledge-Hub/redis-workshop/blob/main/2-internals/README_Lua.md)
      - 2.3.1 Store and load script into redis
      - 2.3.2 Evaluation or redis scrypts
3. [Redis cluster & sharding](https://github.com/Next-Gen-Knowledge-Hub/redis-workshop/blob/main/3-cluster&shard/README_cluster&shard.md)
    - 3.1 Setup a redis cluster inside docker
    - 3.2 Workaround redis cluster with `Go`
    - 3.3 Configurations (`listpack` configs)
    - 3.4 Sharding
      - 3.4.1 Internal sharding structure
      - 3.4.2 Clientside key sharding
4. [Usecase (implementation `Go`)](https://github.com/Next-Gen-Knowledge-Hub/redis-workshop/tree/main/4-imp)
      - 4.1 Metric collector
      - 4.2 Autocomplete
      - 4.3 Lock
      - 4.4 Semaphor
      - 4.5 Push
      - 4.6 Clientside sharding
 
Feel free to use and make any change ;)
