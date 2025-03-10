# Consistency 
- Consistency refers to a database query returning the same data each time the same request is made.
## eventual consistency and strong consistency:
- Eventual Consistency is a guarantee that when an update is made in a distributed database, that update will eventually be reflected in all nodes that store the data, resulting in the same response every time the data is queried.
- Strong consistency means the latest data is returned, but, due to internal consistency methods, it may result with higher latency or delay because all nodes need to be data replicated before displaying data. which takes a bit more time.
-  With eventual consistency, results are less consistent early on, but they are provided much faster with low latency. Early results of eventual consistency data queries may not have the most recent updates because it takes time for updates to reach replicas across a database cluster.
-  The records stay available, but once the transaction has completed across a majority of nodes, the transaction is deemed successful. Data replication across all nodes can take a little more time, but the data in all nodes will become consistent eventually.

## ACID vs. BASE: How is Eventual Consistency Different from Strong Consistency?
- Distributed databases with a BASE model give high availability.
- Early results of eventual consistency data queries may not have the most recent updates. This is because it takes time for updates to reach replicas across a database cluster. In strong consistency, data is sent to every replica the moment a query is made. This causes delay because responses to any new requests must wait while the replicas are updated. When the data is consistent, the waiting requests are handled in the order they arrived and the cycle repeats.
- In contrast to **SQL’s ACID** guarantees, **NoSQL** databases provide so-called **BASE** guarantees.
-  A BASE enables availability and relaxes the stringent consistency. The acronym BASE designates:

  - **Basic Availability** – Data is available most of the time, even during a partial system failure.
  - **Soft state** – replicas are not consistent all the time.
  - **Eventual consistency**– data will become consistent at some point in time, with no guarantee when.

- As such, NoSQL databases sacrifice a degree of consistency in order to increase availability. Rather than providing strong consistency, they provide eventual consistency. This means that a datastore that provides BASE guarantees can occasionally fail to return the result of the latest WRITE.

- Eventual consistency in NoSQL supports the BASE (basically available eventually consistent) pattern for speed and scalability.

## Eventual Consistency Examples:
- concept is same but implementation  of eventual consistency can differ from player to player

- **DynamoDB and Cassandra** use quorum-based models for balancing consistency and availability, especially in write-heavy systems.
- **Redis emphasizes** low-latency reads/writes with asynchronous replication, making it great for caching but limited in strong consistency.
- **Elasticsearch** is search-optimized, but its eventual consistency is influenced by both replication lag and index refresh intervals, blending models from others.

#### Eventual Consistency in DynamoDB Architecture:
- DynamoDB is a fully managed NoSQL database by AWS.
- It uses replication across multiple availability zones (AZs) in a region for high availability and fault tolerance.
- It is based on ideas from Amazon Dynamo, which uses ring topology with quorum-based reads/writes.
##### Consistency Model:

**DynamoDB offers two types of read consistency:**
- Eventually Consistent Reads (default): Reads might not reflect the results of a recently completed write immediately.
- Strongly Consistent Reads: Returns the latest value, assuming a majority quorum.
##### How it works:
- Writes are replicated across multiple nodes (AZs).
- Eventually consistent reads may hit a node that hasn't received the latest update yet.
- With strong consistency, DynamoDB ensures that the read happens after confirming write propagation to a majority of replicas.
##### Ring Topology Analogy:
- Though AWS hides the internals, conceptually it's similar to Dynamo ring-based architecture where data is distributed using consistent hashing, and each item is replicated across multiple nodes.
####  Eventual Consistency in Cassandra
- Peer-to-peer distributed database, unlike DynamoDB’s managed service.
- Nodes are organized in a ring topology using consistent hashing.
- No master node — all nodes are equal.
##### Consistency Model:

- Cassandra uses tunable consistency:
  - You choose the number of replicas that must acknowledge a read/write operation (e.g., QUORUM, ALL, ONE).
  - Write may be accepted by a few nodes and propagated to others asynchronously (hinted handoff, anti-entropy repairs).
##### How it works:

- Write goes to a coordinator node, which forwards to replica nodes.
- Read requests can return stale data if they hit a node that hasn’t received the latest write — eventual consistency.
- If the read hits a QUORUM, it increases the chances of returning the latest data.
- Cassandra uses read repair and gossip protocols to eventually bring all replicas in sync.
