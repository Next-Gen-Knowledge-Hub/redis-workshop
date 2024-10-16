# Redis String Data Structure and Operations

In Redis, a **string** is the simplest and most commonly used data structure. A Redis string is a sequence of bytes, up to 512 MB in length. It can store any type of data, including text, numbers, or serialized objects. Strings are binary-safe, meaning they can store both plain text and binary data.

## Key Features of Redis Strings

- **Simplicity**: Strings are simple key-value pairs.
- **Versatility**: Strings can hold any type of data, including JSON, images, or serialized objects.
- **Atomic Operations**: Operations on strings are atomic, meaning they are executed as a single operation.

### Operations on Redis Strings/Numerics
| Command  | Description                                               | Example Command                    |
| -------- | --------------------------------------------------------- | ---------------------------------- |
| `SET`    | Set the value of a key                                    | `SET mykey "Hello, Redis!"`        |
| `GET`    | Get the value of a key                                    | `GET mykey`                        |
| `DEL`    | Delete a key                                              | `DEL mykey`                        |
| `EXISTS` | Check if a key exists                                     | `EXISTS mykey`                     |
| `INCR`   | Increment the integer value of a key by one               | `INCR mycounter`                   |
| `DECR`   | Decrement the integer value of a key by one               | `DECR mycounter`                   |
| `INCRBY` | Increment the integer value of a key by a specific amount | `INCRBY mycounter 5`               |
| `DECRBY` | Decrement the integer value of a key by a specific amount | `DECRBY mycounter 3`               |
| `APPEND` | Append a value to a key                                   | `APPEND mykey " Welcome!"`         |
| `STRLEN` | Get the length of the value stored at a key               | `STRLEN mykey`                     |
| `SETEX`  | Set the value of a key with an expiration time            | `SETEX mykey 300 "Temporary"`      |
| `MSET`   | Set multiple keys to multiple values                      | `MSET key1 "value1" key2 "value2"` |
| `MGET`   | Get the values of multiple keys                           | `MGET key1 key2`                   |


# Redis List Data Structure and Operations

A **list** in Redis is a collection of ordered strings, which can contain duplicate values. Lists are implemented as linked lists, allowing for efficient insertion and deletion of elements from both ends.

## Key Features of Redis Lists

- **Ordered**: Elements are maintained in the order they were added.
- **Flexible Size**: The size of the list can grow or shrink dynamically.
- **Efficient Access**: You can efficiently access and manipulate elements at both ends of the list.

### Redis List Operations

| Command  | Description                                                                                    | Example Command            |
| -------- | ---------------------------------------------------------------------------------------------- | -------------------------- |
| `LPUSH`  | Prepend one or multiple values to a list                                                       | `LPUSH mylist "World"`     |
| `RPUSH`  | Append one or multiple values to a list                                                        | `RPUSH mylist "Redis"`     |
| `LPOP`   | Remove and return the first element of the list                                                | `LPOP mylist`              |
| `RPOP`   | Remove and return the last element of the list                                                 | `RPOP mylist`              |
| `BLPOP`  | Remove and return the first element of one or more lists; blocks until an element is available | `BLPOP list1 list2 0`      |
| `BRPOP`  | Remove and return the last element of one or more lists; blocks until an element is available  | `BRPOP list1 list2 0`      |
| `LRANGE` | Get a range of elements from a list                                                            | `LRANGE mylist 0 -1`       |
| `LREM`   | Remove elements from a list                                                                    | `LREM mylist 0 "Hello"`    |
| `LLEN`   | Get the length of the list                                                                     | `LLEN mylist`              |
| `LSET`   | Set the value of an element in a list by index                                                 | `LSET mylist 0 "NewValue"` |
| `LINDEX` | Get an element from a list by index                                                            | `LINDEX mylist 0`          |
| `LTRIM`  | Trim a list to the specified range                                                             | `LTRIM mylist 0 1`         |

# Redis Set Data Structure

## Overview

A **Set** in Redis is an unordered collection of unique strings. Sets are useful for storing lists of items where duplicates are not allowed. Redis provides a variety of operations for manipulating sets, including adding, removing, and performing set operations like intersections, unions, and differences.

## Key Features of Redis Sets

- **Unordered**: The elements in a set do not have a specific order.
- **Unique**: A set cannot contain duplicate elements.
- **Efficient Operations**: Sets are designed for high performance, especially for operations involving membership tests and set manipulations.

### Operations on Redis Sets

| Command       | Description                                     | Example Command           |
| ------------- | ----------------------------------------------- | ------------------------- |
| `SADD`        | Add one or more members to a set                | `SADD myset "Hello"`      |
| `SREM`        | Remove one or more members from a set           | `SREM myset "World"`      |
| `SMEMBERS`    | Get all the members of a set                    | `SMEMBERS myset`          |
| `SISMEMBER`   | Determine if a given value is a member of a set | `SISMEMBER myset "Hello"` |
| `SCARD`       | Get the number of members in a set              | `SCARD myset`             |
| `SINTER`      | Return the intersection of one or more sets     | `SINTER set1 set2`        |
| `SUNION`      | Return the union of one or more sets            | `SUNION set1 set2`        |
| `SDIFF`       | Return the difference between two sets          | `SDIFF set1 set2`         |
| `SPOP`        | Remove and return a random member from a set    | `SPOP myset`              |
| `SRANDMEMBER` | Get one or multiple random members from a set   | `SRANDMEMBER myset 2`     |
| `SMOVE`       | Move a member from one set to another set       | `SMOVE set1 set2 "Hello"` |


# Redis Hash Data Structure

## Overview

A **Hash** in Redis is a collection of key-value pairs where the keys and values are both strings. Hashes are useful for representing objects or records, allowing you to store multiple fields associated with a single key. Each field in a hash can be accessed individually, making it an efficient way to manage related data.

## Key Features of Redis Hashes

- **Field-Value Structure**: Hashes allow you to store multiple fields for a single key, which can be useful for representing objects.
- **Efficient Storage**: Hashes use a memory-efficient representation, especially when storing a large number of small objects.
- **Atomic Operations**: Operations on hashes are atomic, ensuring consistency when modifying fields.

### Operations on Redis Hashes

| Command   | Description                                                                                     | Example Command                                |
| --------- | ----------------------------------------------------------------------------------------------- | ---------------------------------------------- |
| `HSET`    | Set the value of a field in a hash                                                              | `HSET myhash field1 "value1"`                  |
| `HGET`    | Get the value of a field in a hash                                                              | `HGET myhash field1`                           |
| `HDEL`    | Delete one or more fields from a hash                                                           | `HDEL myhash field1 field2`                    |
| `HGETALL` | Get all fields and values in a hash                                                             | `HGETALL myhash`                               |
| `HKEYS`   | Get all field names in a hash                                                                   | `HKEYS myhash`                                 |
| `HVALS`   | Get all values in a hash                                                                        | `HVALS myhash`                                 |
| `HLEN`    | Get the number of fields in a hash                                                              | `HLEN myhash`                                  |
| `HEXISTS` | Check if a field exists in a hash                                                               | `HEXISTS myhash field1`                        |
| `HINCRBY` | Increment the integer value of a field by a given number                                        | `HINCRBY myhash field1 10`                     |
| `HMSET`   | Set multiple fields to multiple values (deprecated in newer Redis versions; use `HSET` instead) | `HMSET myhash field1 "value1" field2 "value2"` |
| `HMGET`   | Get the values of multiple fields in a hash                                                     | `HMGET myhash field1 field2`                   |


# Redis Sorted Set (ZSet) Data Structure

## Overview

A **Sorted Set** (ZSet) in Redis is a collection of unique elements, where each element is associated with a score that determines its order in the set. Unlike regular sets, sorted sets maintain their elements in a specific order based on their scores, allowing for efficient range queries and ordering operations.

## Key Features of Redis Sorted Sets

- **Unique Elements**: Each element in a sorted set is unique, and duplicate entries are not allowed.
- **Ordered by Score**: Elements are sorted based on their associated scores, allowing for efficient retrieval of elements in order.
- **Range Queries**: Sorted sets support range queries, enabling you to fetch elements within a specified score range.

### Operations on Redis Sorted Sets

| Command            | Description                                             | Example Command               |
| ------------------ | ------------------------------------------------------- | ----------------------------- |
| `ZADD`             | Add one or more members to a sorted set with a score    | `ZADD myzset 1 "one" 2 "two"` |
| `ZREM`             | Remove one or more members from a sorted set            | `ZREM myzset "one"`           |
| `ZRANGE`           | Get a range of members in a sorted set (by index)       | `ZRANGE myzset 0 -1`          |
| `ZRANGEBYSCORE`    | Get members in a sorted set within a score range        | `ZRANGEBYSCORE myzset 1 2`    |
| `ZCARD`            | Get the number of members in a sorted set               | `ZCARD myzset`                |
| `ZSCORE`           | Get the score of a member in a sorted set               | `ZSCORE myzset "two"`         |
| `ZINCRBY`          | Increment the score of a member in a sorted set         | `ZINCRBY myzset 2 "one"`      |
| `ZREVRANGE`        | Get a range of members in a sorted set in reverse order | `ZREVRANGE myzset 0 -1`       |
| `ZRANK`            | Get the rank of a member in a sorted set                | `ZRANK myzset "two"`          |
| `ZREMRangeByScore` | Remove members from a sorted set within a score range   | `ZREMRangeByScore myzset 0 1` |
