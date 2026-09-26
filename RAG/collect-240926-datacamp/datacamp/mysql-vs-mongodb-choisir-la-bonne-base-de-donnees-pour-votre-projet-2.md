---
id: collect-240926-datacamp/datacamp/mysql-vs-mongodb-choisir-la-bonne-base-de-donnees-pour-votre-projet-2
title: "mysql-vs-mongodb-choisir-la-bonne-base-de-donnees-pour-votre-projet"
domain: datacamp
role: reference
task: reference
actors: ["EU"]
dates: []
keywords: ["cost", "embedding", "memory"]
source: docs/RAG/clean_en/datacamp/mysql-vs-mongodb-choisir-la-bonne-base-de-donnees-pour-votre-projet.md
source_anchor: ""
source_lines: [98, 180]
sha256: 03baebb8d2deebc0366b50ebe60d5d1018b9d4180cf9924156fbadd5f8d0f4d9
---

# mysql-vs-mongodb-choisir-la-bonne-base-de-donnees-pour-votre-projet

## Scaling Strategies for MySQL vs MongoDB

There are two opposing strategies for scaling: vertical scaling and horizontal scaling. Vertical scaling involves improving a single machine by adding CPUs, RAM, or faster disks to handle increased load.

This approach is simple in that it does not require clusters or load balancers. However, it is limited by hardware and can be expensive. There is a single point of failure: if the server goes down, everything stops.

In contrast, horizontal scaling involves adding machines to share the workload. It allows you to scale at a lower cost (no need for a huge, expensive server) and ensures operations continue if one machine fails. However, because it requires more coordination, such as load balancing and data synchronization, it can lead to network overhead.

### Limits of Vertical Scaling in MySQL

Vertical scaling of MySQL happens in two steps: first, you upgrade the hardware, then you tune the configuration settings to use those new resources efficiently. As far as hardware is concerned, you can add CPU cores, increase RAM, and use faster storage and network components. Once the machine is powerful enough, use these configuration settings:

- `innodb_buffer_pool_size` . Increase the memory so that it matches your RAM.
- `innodb_log_file_size` ,`innodb_log_buffer_size` : increase them to batch more write operations.
- `max_connections` ,`table_open_cache, thread_cache_size` . Increase them to keep more clients and tables warm in memory.

However, vertical scaling has practical limits. At some point, a single machine will reach its limits in terms of CPU cores or memory. There is a point of diminishing returns for hardware investment; doubling the price you pay for CPU cores, for example, does not double the CPU power.

Write-intensive workloads may encounter issues with locked rows or indexes. A server failure can take the entire database offline.

MySQL can use and does use horizontal techniques, but this requires additional configuration compared to MongoDB's built-in sharding model.

### MongoDB's horizontal scaling via sharding

MongoDB uses a horizontal scaling strategy, called "sharding," which divides a large collection into smaller chunks based on a "shard key" (for example, ranges of document IDs). Each chunk is stored on a different server.

A group of configuration servers maintains a map indicating the zones containing the different chunks. Query routers use this directory to send client requests to the appropriate storage units. An automatic balancer redistributes chunks as data and traffic patterns change, ensuring that no unit is overloaded. Increase capacity by adding storage servers.

This approach is ideal for applications subject to unpredictable or massive loads, such as social media, video streams, and e-commerce during major sales. With storage and query loads distributed across many machines, the system remains responsive and resilient during traffic spikes.

## Security and compliance considerations

For any database deployment, it is essential to ensure that data remains secure and compliant with industry regulations. Let's compare MySQL and MongoDB in terms of authentication, encryption, and support for standards such as HIPAA and GDPR.

### Authentication and encryption

#### MySQL

MySQL secures data by combining user authentication, encrypted connections, and disk-level encryption. Each user must log in with a unique account and password and is granted only the precise privileges they need, either individually or based on their role. All data in transit can be protected by TLS/SSL, enforced as SSL-only, and validated by server certificates. Data at rest is encrypted by the InnoDB engine. Encryption keys can be changed without service interruption.

#### MongoDB

MongoDB secures data through authentication, encryption, and auditing. Each user logs in with a unique username and password (via SCRAM, x.509 certificates, Kerberos, or LDAP) and is assigned permissions through roles, whether built-in or custom. Traffic in transit is protected by TLS/SSL, and data at rest is protected by the WiredTiger storage engine. Detailed audit logs record login attempts and data changes to facilitate security audits and compliance.

### Compliance

HIPAA, the U.S. law that protects medical records, requires strict access controls, encryption in transit and at rest, and comprehensive auditing. MySQL meets these requirements with role-based user accounts, TLS/SSL for data in flight, InnoDB Transparent Data Encryption for data at rest, and audit plugins that record every access.

MongoDB meets the same standards by using SCRAM or x.509 authentication, TLS/SSL, WiredTiger storage encryption, enterprise audit logging, and optional client-side field-level encryption.

The GDPR grants EU residents the right to access, correct, and erase their personal data, imposing strict access controls, encryption, and auditability.

MySQL meets these requirements through role-based privileges, TLS/SSL for secure connections, InnoDB transparent data encryption, audit plugins for tracking changes, and foreign key cascades for enforcing deletions.

MongoDB meets GDPR standards with SCRAM/x.509 authentication, TLS/SSL, WiredTiger at-rest encryption, enterprise audit logging, flexible document-level updates and deletions, and region-specific deployment options for data residency.

## Comparative table of the advantages and disadvantages of MySQL and MongoDB

| Aspect | MySQL (relational) | MongoDB (document-oriented) | 
| Data paradigm | Structured data with a fixed schema. | Unstructured or semi-structured data with a flexible schema. | 
| Storage model | Tables of rows and columns. | JSON-like documents stored in collections. | 
| Schema definition | Explicit. Defined upfront via `CREATE TABLE`; columns have data types and constraints (`NOT NULL`, `UNIQUE`, foreign keys). | Implicit. Documents can contain any field. No prior definition is required. | 
| Relationships | Enforced by foreign key constraints in tables. | This involves embedding related data in a single document or referencing other documents. | 
| Flexibility | Rigid. Adding or modifying columns requires `ALTER TABLE` migrations and careful planning. | Flexible. Fields can be added or removed on the fly without downtime. | 
| Schema migration | Scripts (`ALTER TABLE`, `CREATE TABLE`), potential table locks, testing, and rollback planning. | Document-level update scripts using `$set`, `$rename`, `$unset`, generally idempotent and executable in the background. | 
| Indexing and queries | Optimized for multi-table `JOIN`s and aggregations. | Indexes on document fields and subfields. Aggregation pipelines support grouping, filtering, and `$lookup` joins. | 

## Use Case Analysis: Choosing the Right Tool

Choosing the right database depends on your application's workload, data model, and growth patterns. Let's compare the scenarios in which MySQL and MongoDB provide the most value.

### Ideal Scenarios for MySQL

MySQL excels in applications that require reliable, high-volume transaction processing and well-structured data.

Financial institutions use MySQL to record account transfers and balance updates, e-commerce platforms use it to manage orders, inventory, and customer information. Healthcare systems store patient records and appointment data, and large enterprises use ERP or CRM systems that connect sales, billing, and support.

MySQL's rich feature set (GROUP BY, JOIN, window functions, and indexed lookups) makes it easy to build structured reports, such as monthly sales totals or average response times, and to connect easily to BI tools.

### MongoDB's Strengths in Modern Applications

MongoDB's schema-less design and horizontal scaling make it the ideal solution for applications that need to evolve quickly or handle massive datasets. New features can be added without costly schema migrations.

