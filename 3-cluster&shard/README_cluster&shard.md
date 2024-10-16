# Redis Cluster Sharding

Redis Cluster is a distributed version of Redis that enables horizontal scaling and high availability by partitioning data across multiple nodes using **sharding**. This setup allows Redis to distribute the data load, ensuring seamless performance even at scale.

## Key Concepts

### 1. **Sharding with Hash Slots**
Redis Cluster divides data into **16,384 hash slots**. When a key is added, it is hashed to a specific slot using the CRC16 algorithm. Each node in the cluster is responsible for a range of these hash slots.

For example:
- The key `user:123` is hashed and assigned to a specific slot, say 1000.
- If Node A is responsible for hash slots 0–5000, the key-value pair will be stored on Node A.

### 2. **Data Distribution**
Redis Cluster automatically distributes keys across the available nodes by assigning different hash slot ranges to each node. This ensures that the load is evenly distributed and provides efficient data storage.

### 3. **Replica Nodes**
Redis Cluster supports **replication**, where each master node has one or more replica nodes. These replicas can take over if the master node fails, ensuring fault tolerance and high availability.

### 4. **Automatic Failover**
In case of node failure, Redis Cluster performs automatic failover by promoting a replica to become the new master. This keeps the cluster operational even during node outages.

## Operations in Redis Cluster

1. **Key Sharding**: Redis automatically determines which node stores a given key using its hash slot mapping.
2. **Node Failover**: If a node fails, its replica automatically takes over the corresponding hash slots.
3. **Scaling**: Adding or removing nodes from the cluster dynamically redistributes hash slots across the nodes.

## Sharding Example

Consider three nodes in the cluster:

- Node A: responsible for hash slots 0-5000
- Node B: responsible for hash slots 5001-10000
- Node C: responsible for hash slots 10001-16383

For a key `item:2000`, it might hash to slot 3500. Since Node A holds that range of slots, `item:2000` will be stored on Node A.

If Node A fails, its replica takes over, ensuring that the data in its slots remains accessible.

## Commands for Redis Cluster Sharding

- **Gossip Protocol**: Redis Cluster uses a gossip-based protocol to maintain node state information and redistribute data across the cluster.
- **Hash Slots Management**: Redis handles the assignment of hash slots to nodes, and users do not need to manually configure this.
- **`CLUSTER INFO`**: Provides details about the cluster state.
- **`CLUSTER NODES`**: Lists the nodes and their roles (master or replica) in the cluster.

## Example Cluster Configuration

```bash
cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
appendonly yes
```

## Use Cases

- **Large-scale** applications: Redis Cluster is ideal for apps needing horizontal scalability.
- **High availability**: Applications requiring fault tolerance benefit from Redis Cluster’s failover and replication mechanisms.
Limitations
- **Multi-key operations**: Redis Cluster only supports multi-key commands if all keys map to the same hash slot.
Manual rebalancing: When adding or removing nodes, hash slots must sometimes be manually rebalanced.

## Conclusion
Redis Cluster provides a powerful, scalable, and highly available solution for managing large-scale Redis deployments through sharding. By distributing keys across multiple nodes and providing built-in replication and failover, Redis Cluster ensures seamless data operations with minimal downtime.
