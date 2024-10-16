# Redis Pub/Sub Messaging System

## Overview

**Pub/Sub** (Publish/Subscribe) in Redis is a messaging paradigm that allows for message broadcasting to multiple subscribers. In this model, publishers send messages to channels without knowing who (if anyone) will receive them, while subscribers listen for messages on channels they are interested in. This decouples the communication between the sender and the receiver.

## Key Features of Redis Pub/Sub

- **Decoupled Architecture**: Publishers and subscribers do not need to know about each other.
- **Real-time Messaging**: Messages can be delivered to subscribers in real-time.
- **Multiple Subscribers**: A single message can be sent to multiple subscribers across different channels.
- **Lightweight**: Pub/Sub is efficient for high-throughput message broadcasting.

### Operations in Redis Pub/Sub

| Command        | Description                                               | Example Command              |
| -------------- | --------------------------------------------------------- | ---------------------------- |
| `PUBLISH`      | Publish a message to a channel                            | `PUBLISH mychannel "Hello!"` |
| `SUBSCRIBE`    | Subscribe to one or more channels                         | `SUBSCRIBE mychannel`        |
| `UNSUBSCRIBE`  | Unsubscribe from one or more channels                     | `UNSUBSCRIBE mychannel`      |
| `PSUBSCRIBE`   | Subscribe to channels using a pattern (wildcard matching) | `PSUBSCRIBE my*`             |
| `PUNSUBSCRIBE` | Unsubscribe from channels using a pattern                 | `PUNSUBSCRIBE my*`           |


# Redis Transactions

## Overview

Redis Transactions allow you to execute a series of commands as a single atomic operation. This means that either all the commands in the transaction are executed successfully, or none are. Redis implements transactions using the `MULTI`, `EXEC`, `WATCH`, and `DISCARD` commands, ensuring that commands can be queued up for execution while preventing other clients from modifying the data until the transaction is complete.

## Key Features of Redis Transactions

- **Atomicity**: Transactions in Redis ensure that either all commands are executed or none at all.
- **Isolation**: Commands within a transaction are not visible to other clients until the transaction is executed.
- **Multiple Commands**: You can queue multiple commands in a transaction.

### Operations in Redis Transactions

| Command   | Description                                                                 | Example Command |
| --------- | --------------------------------------------------------------------------- | --------------- |
| `MULTI`   | Start a transaction block. Subsequent commands will be queued.              | `MULTI`         |
| `EXEC`    | Execute all commands in the transaction block.                              | `EXEC`          |
| `DISCARD` | Discard all commands in the transaction block.                              | `DISCARD`       |
| `WATCH`   | Watch a key for changes. If the key is modified, the transaction will fail. | `WATCH mykey`   |


### Translaction `failure` scenario

`Client-A`
```bash
1> SET balance 100
2> SET transaction_amount 50
3> GET balance
"100"
4> WATCH balance
5> MULTI
7> DECRBY balance 50
QUEUED
8> EXEC
(error) EXECABORT Transaction discarded because of previous errors.
```

`Client-B`
```bash
6> SET balance 50
```