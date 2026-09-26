---
id: collect-240926-datacamp/datacamp/top-34-questions-d-entretien-mysql-et-reponses-pour-2026-3
title: "top-34-questions-d-entretien-mysql-et-reponses-pour-2026"
domain: datacamp
role: reference
task: reference
actors: []
dates: ["2024-12-24"]
keywords: ["cost", "memory", "parameters"]
source: docs/RAG/clean_en/datacamp/top-34-questions-d-entretien-mysql-et-reponses-pour-2026.md
source_anchor: ""
source_lines: [377, 562]
sha256: e26f6a2d928be80e1dcaa5eaef32a1bcf35adc58f134b0235c3b495f11023099
---

# top-34-questions-d-entretien-mysql-et-reponses-pour-2026

| **history_id** | **order_id** | **action** | **timestamp** | 
| **1** | 1 | inserted | 2024-12-24 10:00:00 | 
| **2** | 2 | inserted | 2024-12-24 11:00:00 | 

### 22. Why does adding an index speed up SQL queries?

Without an index, the database must scan every row to find an entry. An index is like a table of contents, allowing direct access to the relevant rows. Adding an index therefore reduces search time and speeds up queries.

Indexes generally rely on data structures such as B-trees or hash tables, which facilitate searches, lookups, and range scans.

Example of creating an index:

```
-- Without an index:
SELECT * FROM employees WHERE last_name = 'Smith';
-- Adding an index on the last_name column:
CREATE INDEX idx_last_name ON employees(last_name);
-- With the index, the database quickly locates rows with 'Smith' in last_name.
```
Indexes also have drawbacks:

- 
Slower writes: `INSERT`, `UPDATE`, and `DELETE` operations are slower because the index must be updated with every change.
- 
Storage cost: indexes consume additional disk space.

### 23. What data types should be used for a product's weight and price, and why?

For weight, `DECIMAL` is generally the safest choice. `FLOAT` and `REAL` store decimals but use floating-point arithmetic, which is a source of small rounding discrepancies.

For weights requiring precision (shipping, inventory), `DECIMAL(8, 3)` offers precise control with 3 decimal places without unpleasant rounding surprises. `FLOAT` is only acceptable if a small margin of error is tolerated.

### 24. How do you find duplicates in SQL with a window function?

Here is how to spot duplicates with the `ROW_NUMBER()` function:

```
WITH DuplicateCheck AS (
    SELECT product_name, 
           category,
           ROW_NUMBER() OVER(
               PARTITION BY product_name, category 
               ORDER BY id
           ) AS row_num
    FROM sales
)
SELECT *
FROM DuplicateCheck
WHERE row_num > 1;
```
Explanation:

1. `ROW_NUMBER()` numbers each row in the result. 

2. `PARTITION BY` groups by `product_name` and `category`. 

3. Within each group, numbering starts at 1.

4. Any row with `row_num` > 1 is a duplicate. 

For example, for these records:

- Product A, Category X, row_num = 1
- Product A, Category X, row_num = 2 (duplicate)
- Product B, Category Y, row_num = 1

The query will return the second row, because its `row_num` is greater than 1.

### 25. How do you create and use a stored procedure with parameters in MySQL? Give an example.

Stored procedures let you save and reuse complex queries to make operations more efficient and maintainable. Here is an example with parameters.

Suppose you have a student database and need to filter by age. Let's create a procedure that takes an age parameter:

First, a simple procedure with an input parameter:

```
CREATE PROCEDURE get_student_info(IN age INT)
BEGIN
    SELECT * FROM student WHERE student.age = age;
END;
```
To execute it, you `CALL` it with the desired age:

`CALL get_student_info(21);`
You can go further with output parameters. For example, counting students of a given age:

```
CREATE PROCEDURE count_students_by_age(IN age INT, OUT student_count INT)
BEGIN
    SELECT COUNT(*) INTO student_count FROM students WHERE students.age = age;
END;
```
To retrieve the result:

```
SET @count = 0;
CALL count_students_by_age(21, @count);
SELECT @count AS total_students;
```
### 26. Why is referential integrity important in a database?

Referential integrity ensures consistency in relationships between tables. The foreign key guarantees that values in one table match a unique value in the referenced table.

Concrete example: you manage an e-commerce database with a `Customers` table and an `Orders` table. Each order must belong to a real customer. Referential integrity, through foreign keys, ensures that:

- You cannot create an order for a nonexistent customer.
- You cannot delete a customer who has orders (unless there is an explicit rule on what happens to those orders).
- You cannot modify the ID of a customer who has existing orders.

Thus, by creating a foreign key constraint:

```
ALTER TABLE Orders
ADD FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID);
```
The database automatically enforces these rules:

- 
Each `CustomerID` in `Orders` must exist in `Customers`.
- 
Any attempt at a violation (inserting an invalid `CustomerID`) is rejected.

This prevents critical inconsistencies, such as orders that cannot be linked to customers, or incomplete reports.

## MySQL Questions for Database Administrators

If you are specifically applying for a database administrator position, here are some possible questions.

### 27. Why does a large-scale application use sharding? What are the associated challenges?

Sharding distributes a large volume of data across multiple servers. Each shard contains part of the data. By distributing the load, you avoid high-end hardware. Speed and scalability improve, but there are challenges:

- Some queries such as joins become difficult, which complicates data management.
- As it grows, some shards may become saturated, creating hot spots that degrade performance.

### 28. Explain the role of redo logs in MySQL crash recovery.

On every modification, MySQL must write to disk. Writing directly to data files is slow and risky. This is why, before modifying these files, MySQL first writes what it is going to do to the redo log. This is safer than updating data files on the fly.

For example, you update a customer's address:

1. MySQL first writes this change to the redo log.
2. It then confirms your transaction as committed.
3. Finally, it applies the modification to the data files.

Crash recovery comes into play if MySQL goes down after step 1 or 2 but before step 3. On restart, MySQL reads its redo logs and then finishes the work by replaying the changes. This guarantees that a committed transaction is not lost, even in the event of a failure at the wrong moment.

### 29. What storage engines are available in MySQL and how do they differ?

MySQL supports several storage engines, each optimized for different needs. Here is a comparison of the most common ones:

| Storage engine | Key features | Ideal for |
|---|---|---|
| **InnoDB** | Default engine. ACID compliant, row-level locking, transactions, and foreign keys. | E-commerce, financial systems, anything requiring data integrity. |
| **MyISAM** | Fast reads, table-level locking. No transactions or foreign keys. | Read-heavy applications where speed takes priority over integrity. |
| **Memory** | Data in RAM. Extremely fast but lost on restart. | Caching, session management, temporary data. |
| **CSV** | Stores data in CSV files. No indexing. | Data exchange between apps or simple flat storage. |
| **Archive** | High compression. INSERT and SELECT only. No indexes. | Logs or histories that are rarely queried. |
| **NDB (Clustered)** | Distributed storage, high availability, fault tolerance, transactions supported. | Large-scale distributed applications requiring real time. |

### 30. How do you set a default storage engine in MySQL?

Start by checking the current default engine:

`SHOW ENGINES;`
InnoDB is recommended as the default engine because it supports key features:

- ACID-compliant transactions
- Foreign key constraints
- Crash recovery
- Row-level locking

To temporarily change the default engine for your session:

`SET default_storage_engine = 'InnoDB';`
For a permanent change, edit the MySQL configuration file by adding the following line under the `[mysqld]` section:

`default-storage-engine = InnoDB`
### 31. How do you repair corrupted tables in MySQL?

Start by checking all databases with the following command:

`mysqlcheck --check --all-databases -u root -p`
It scans the tables and reports any corruption. You can then run the repair:

