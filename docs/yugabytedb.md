# YugabyteDB Learning Path

## Universe

group of nodes that collectively function as a resilient and scalable distributed database.

## Deployment configurations

* Single availability zone
* Multiple availability zones in a region
* Multiple regions, with synchronous and asynchronous replication choices

## Organization of data

* Univers: one or more namespaces
* Namespace: one or more user tables

> Yugabyte automatically shards, replicates, and load-balances these tables across the nodes in the universe

### YSQL

A namespace in ysql is referred to as a database and is logically identical to a namespace in other RDBMS

### YCQL

A namespace in ycql is referred to as a keyspace and is logically identical to a keyspace in Apache Cassandra's cql

## Component Services

Universe: two sets of servers.

* Each node would have two different servers running.
* Servers:
  * Yugabyte Tablet Server (YB-TServer)
  * Yugabyte Master Server (YB-Master)
* both kinds of servers form two respective distributed services usin RAFT as a building block.
* High Availability: failure detection, leader election, data replication in the RAFT implementation

## Tablet Server

* Responsible for hosting and serving user data, as well as dealing with all the queries

## Master Server

* Responsible for keeping system metadata, coordinating system-wide operations
  * Creating, altering, and dropping tables, and load balancing

![yugabyte universe node components](https://docs.yugabyte.com/images/architecture/4_node_cluster.png "yugabyte universe node components")

## Universe vs. cluster

* Universe is comprised of one primary cluster and zero or more read replica clusters
  * A primary cluster can perform both writes and reads. Replication between nodes in a primary cluster is performed synchronously
  * Read replica clusters can perform only reads; writes sent to read replica clusters get automatically rerouted to the primary cluster of the universe. These clusters enable reads in regions that are far away from the primary cluster with timeline-consistent data. This ensures low latency reads for geo-distributed applications. Data is brought into the read replica clusters through asynchronous replication from the primary cluster.
  * Nodes in a read replica cluster act as Raft observers that do not participate in the write path involving the Raft leader and Raft followers present in the primary cluster.

## Core functions

## Universe creation

Steps to create a YugabyteDB universe:

1. Start YB-Masters
![yugabyte master creation](https://docs.yugabyte.com/images/architecture/create_universe_masters.png "yugabyte master creation")
    * Bring sufficient number of YB-Masters, with each being made aware of the others.
        * As many as the replication factor
    * YB-Masters initialize themselves with a universally unique identifier(UUID), learn about each other and perform a leader election.
    * At the end of this step, one of the masters establises itself as the leader.
![yugabyte leader election](https://docs.yugabyte.com/images/architecture/create_universe_tserver_heartbeat.png "yugabyte leader election")
2. Start YB-TServers
![yugabyte tablet server creation](https://docs.yugabyte.com/images/architecture/create_universe_tserver_heartbeat.png "yugabyte tablet server creation")
    * You need to start as many tablet servers as there are nodes, with the master addresses being passed to them on start up.
    * They start sending heartbeats to the masters, communicating the fact that they are alive.
