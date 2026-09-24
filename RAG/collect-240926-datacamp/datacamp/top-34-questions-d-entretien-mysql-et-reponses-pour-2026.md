---
id: collect-240926-datacamp/datacamp/top-34-questions-d-entretien-mysql-et-reponses-pour-2026
title: "top-34-questions-d-entretien-mysql-et-reponses-pour-2026"
domain: datacamp
role: reference
task: reference
actors: ["Microsoft", "Oracle", "United States"]
dates: ["2024-01-01", "2024-12-24", "2024-12-31", "9999-12-31"]
keywords: ["cost", "memory", "parameters"]
source: docs/RAG/clean_en/datacamp/top-34-questions-d-entretien-mysql-et-reponses-pour-2026.md
source_anchor: ""
source_lines: [1, 687]
sha256: 06982b7e0dfd2036f57ee4da77a82abed74e9f58b2f236ad1e86235661bcc07c
---

# top-34-questions-d-entretien-mysql-et-reponses-pour-2026

<!-- source: https://www.datacamp.com/fr/blog/mysql-interview-questions -->

Course

You've probably noticed that MySQL appears in almost every job posting related to databases. And for good reason: MySQL powers a multitude of services, from the social networks you browse to the applications you use every day.

In this guide, I've gathered the most common MySQL interview questions. We'll go over the essentials — from concepts every junior must master — to the more advanced topics expected for senior positions. I'll also share tips to help you present yourself as a confident candidate in your upcoming data interviews.

## What is MySQL?

MySQL is an open-source RDBMS (relational database management system) based on SQL, which organizes data into structured tables. It is developed by Oracle Corporation.

It ranked as the most popular open-source DBMS in 2024. However, the 2025 Stack Overflow Developer Survey showed that PostgreSQL has become the most used database by professional developers, surpassing MySQL for the first time.

That doesn't mean MySQL has lost its luster: it remains extremely popular — 40.5% usage among developers in 2025 — and still powers countless web applications, CMS, and enterprise tools. And if you work on web applications or the LAMP stack, MySQL remains a top-tier skill.

In 2024, MySQL was the most popular open-source DBMS in the world, with a score of 1061. Source: Statista.

## Basic MySQL Questions

During the initial exchanges, the interviewer may ask fundamental questions to assess your understanding of basic database and MySQL concepts.

### 1. What is a database and how does it differ from a DBMS?

A database is a storage container that holds data which can be accessed, modified, and analyzed. For example, social platforms store in databases who liked our posts.

A DBMS (Database Management System) is the software that allows you to interact with that data and administer it, for example by creating users and managing their permissions. MySQL is one of the most popular DBMSs. Other examples include PostgreSQL, MongoDB, and Microsoft SQL Server.

### 2. How does MySQL differ from other relational database management systems?

MySQL is an open-source RDBMS that uses SQL to manage data. It is known for its simplicity, speed, and compatibility with web applications.

Here's what sets MySQL apart:

- Simplicity and performance: MySQL is often praised for its ease of use and optimized performance, making it a preferred choice for web developers and startups.
- Advanced features: While MySQL excels in simplicity, it may lack some advanced features found in other RDBMSs like PostgreSQL, such as more complete ACID transaction support, advanced indexing, and a wider range of data types.
- Storage engines: MySQL allows you to choose different storage engines (e.g., InnoDB, MyISAM) for tables, offering flexibility depending on the use case.

MySQL is ideal for scenarios requiring speed and scalability. For more complex needs and "enterprise" features, PostgreSQL may be a better choice.

### 3. What are the main data types available in MySQL?

MySQL supports a variety of data types, grouped into:

- 
Numeric: `INT` ,`DECIMAL` ,`FLOAT` ,`DOUBLE` , etc.
- 
**Strings:**`CHAR` ,`VARCHAR` ,`TEXT` ,`BLOB` .
- 
**Date/time:**`DATE` ,`DATETIME` ,`TIMESTAMP` ,`TIME` .
- 
**JSON:** for storing JSON objects.

### 4. What is the difference between the INT and DECIMAL data types?

`INT` stores whole numbers without decimals. It is used when fractions are not needed. Conversely, `DECIMAL` allows you to store financial values and is suitable for precise calculations with decimals. 

### 5. How does DATE differ from DATETIME in MySQL?

The `DATE` type in MySQL stores a date in the format year, month, day: 

`YYYY-MM-DD`

The `DATETIME` type stores the date with the time, in the format: 

`YYYY-MM-DD HH:MM:SS` 

### 6. What is a foreign key and how do you use it in a database?

A foreign key is a field in one table that references the primary key of another table.

For example, in a `customers` table that stores customer information, each customer has a unique `customer_id`. In another table called `transactions` (which records purchases), `customer_id` is used as a foreign key. The `customer_id` in the transactions table links each purchase to a specific customer in the `customers` table.

Here's what it looks like in SQL:

```
CREATE TABLE customers (
    customer_id INT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);
CREATE TABLE transactions (
    transaction_id INT PRIMARY KEY,
    customer_id INT,
    amount DECIMAL(10,2),
    date DATE,
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);
```
### 7. What are the differences between INNER JOIN, LEFT JOIN, RIGHT JOIN and FULL JOIN?

Joins combine rows from two or more tables based on related columns. Their differences:

- 
INNER JOIN: returns only the rows that have a match in both tables.
- 
LEFT JOIN: returns all rows from the left table and the matching rows from the right table. If there is no match, the columns from the right table are `NULL`.
- 
RIGHT JOIN: similar to `LEFT JOIN`, but returns all rows from the right table and the matches from the left.
- 
**FULL JOIN**: combines the results of `LEFT JOIN` and `RIGHT JOIN`, including unmatched rows from both sides. **Note:** MySQL does not natively support the `FULL JOIN` syntax. To get the same result, use a `UNION` of a `LEFT JOIN` and a `RIGHT JOIN`

### 8. What is the difference between DELETE, TRUNCATE, and DROP in MySQL?

The `DELETE`, `TRUNCATE`, and `DROP` commands may look similar, but their behavior differs:

DELETE: deletes rows from a table based on a condition. Can be rolled back if executed within a transaction. Example:

`DELETE FROM employees WHERE department_id = 5;`
**TRUNCATE: deletes all rows from a table, but the structure remains intact. Faster than `DELETE` and cannot be rolled back. Example:**

`TRUNCATE TABLE employees;`
DROP: completely deletes the table (structure and data), as well as its dependencies (indexes, etc.). Example:

`DROP TABLE employees;`
### 9. How do you create and modify a table in MySQL? Give examples.

To create tables, use the `CREATE TABLE` statement, and to modify them, most often `ALTER TABLE`. Examples:

Creating a table:

`CREATE TABLE employees (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50), department VARCHAR(50));`
Adding a column:

`ALTER TABLE employees ADD COLUMN salary DECIMAL(10, 2);`
### 10. What is a temporary table in SQL?

A temporary table exists only during the current session. Once the session is closed, the table is deleted. It is used to store intermediate results, for example to test, filter, or prepare data before inserting it into a permanent table.

Example:

```
CREATE TEMPORARY TABLE temp_employees (
    id INT,
    name VARCHAR(50)
);
INSERT INTO temp_employees VALUES (1, 'John Doe');
SELECT * FROM temp_employees;
```
### 11. What is a subquery in MySQL? Explain with an example.

A subquery (or nested query) is a query included within another one. It allows complex operations to be broken down into simpler steps. For example, to find employees earning above the average:

```
SELECT first_name, last_name, salary
FROM employees
WHERE salary > (
    SELECT AVG(salary)
    FROM employees
);
```
Breakdown:

1. 
The inner query `SELECT AVG(salary) FROM employees` first calculates the average salary.
2. 
The outer query then uses this average to find the employees who exceed it.

### 12. How do you use the INSERT statement to add data to a table? Best practices?

The `INSERT` statement is used to add data. The basic syntax:

```
INSERT INTO table_name (column1, column2, ...) 
VALUES (value1, value2, ...); 
```
A few best practices to follow with `INSERT`:

1. 
Explicitly list the columns. The code is clearer and remains robust if the structure changes.
2. 
For `AUTO_INCREMENT` columns (such as IDs), do not include them in the `INSERT`. MySQL handles them to avoid duplicates.
3. 
Be consistent in the use of quotation marks for strings. Single quotes are often preferred, but both work.
4. 
To insert multiple rows, prefer a single statement for better performance.

### 13. What is the purpose of the AUTO_INCREMENT attribute in MySQL?

The `AUTO_INCREMENT` attribute generates unique and sequential numbers for a column, usually the primary key.

Example of creating a table with an `AUTO_INCREMENT` column:

```
CREATE TABLE employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    department VARCHAR(50)
);
```
And inserting rows:

```
INSERT INTO employees (name, department) VALUES ('John Doe', 'Sales');
INSERT INTO employees (name, department) VALUES ('Jane Smith', 'Marketing');
```
### 14. What is a view in MySQL?

A view is a saved query that behaves like a virtual table. It allows you to encapsulate a complex query under a name and reuse it like a table. This way, you do not have to retype the entire query each time.

For example, to simplify querying employees with their department name, you can create a view:

```
CREATE VIEW employee_details AS
SELECT 
    e.id,
    e.name,
    d.department_name,
    e.salary
FROM 
    employees e
JOIN 
    departments d ON e.department_id = d.department_id;
```
You can then query the `employee_details` view like a table:

`SELECT * FROM employee_details;`
However, insertion and updating through views are limited. Most are read-only and hide direct access to the tables, thereby strengthening security. Views can also slow down certain queries, because the underlying query is executed on each access.

## Improving SQL for beginners

## Intermediate MySQL questions

In this section, we cover intermediate-level topics. These questions are mainly intended to assess your knowledge of data types and MySQL structure.

### 15. What are system-versioned tables and how do they work?

System-versioned tables keep the complete history of changes to a table. Since they retain the previous versions of each row, they can be used for auditing and data recovery.

They work by adding two additional columns — `StartTime` and `EndTime` — which record the validity period of each row. During inserts, updates, or deletions, these timestamps are adjusted:

- 
Insert: a new row is added with a `StartTime` set to the current timestamp and an `EndTime` at `9999-12-31 23:59:59` — MySQL's maximum `DATETIME` value, used as a sentinel to indicate "currently active row".
- 
Update: the `EndTime` of the original row is set to the current timestamp to mark it as invalid, then a new updated row is created with `StartTime` at the current moment and `EndTime` "infinite".
- 
Delete: the `EndTime` of the existing row is set to the current timestamp, indicating that it is no longer valid.

With SQL's `FOR SYSTEM_TIME` clause, you can query the state of the table at a given moment or over a period. For example:

- 
`FOR SYSTEM_TIME AS OF '2024-01-01'`: retrieves the state of the table as it was on January 1, 2024.
- 
`FOR SYSTEM_TIME BETWEEN '2024-01-01' AND '2024-12-31'`: displays all rows valid during that interval.

### 16. What is a MySQL transaction and how do you use it?

A transaction is a set of operations executed as a unit. It guarantees data integrity by requiring that all operations succeed or fail together.

Example of use:

```
START TRANSACTION;
UPDATE accounts SET balance = balance - 500 WHERE account_id = 1;
UPDATE accounts SET balance = balance + 500 WHERE account_id = 2;
COMMIT; -- Saves permanently
-- or
ROLLBACK; -- Cancels the changes
```
### 17. What is a default constraint in MySQL? How do you set a default value for a column?

A default constraint assigns a value to a column when no explicit value is provided during an `INSERT`. It guarantees the validity of the column if the user omits it during input.

Example of creation with a default value:

```
CREATE TABLE employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    status VARCHAR(10) DEFAULT 'active'
);
```
You can then insert a row without specifying `status`:

`INSERT INTO employees (name) VALUES ('John Doe');`
This approach reduces the risk of `NULL` or invalid data on critical columns and simplifies queries by avoiding the need to handle default cases in code.

| **Field** | **Type** | **Null** | **Key** | **Default** | **Extra** | 
| **id** | INT | NO | PRI | NULL | AUTO_INCREMENT | 
| **name** | VARCHAR(50) | YES |  | NULL |  | 
| **status** | VARCHAR(10) | YES |  | active |  | 

This command is useful because:

- It helps developers understand the schema before writing queries.
- It is useful for debugging, especially on unfamiliar databases.
- It allows you to quickly identify constraints, such as primary keys or default values.

### 18. What is the difference between CHAR and VARCHAR in MySQL?

Both store strings, but differ in how they manage space:

- 
`CHAR(n)` always stores exactly `n` characters, padding with spaces if needed. Fixed length, which can be slightly faster for columns of uniform length (country codes, statuses, etc.).
- 
`VARCHAR(n)` stores only the characters entered, up to `n`. It saves space for variable-length data, with a slight overhead for length management.

```
CREATE TABLE example (
    country_code CHAR(2),      -- Always 2 characters, e.g. 'US', 'UK'
    email VARCHAR(255)         -- Variable length up to 255
);
```
Rule of thumb: use `CHAR` for fixed-length values, `VARCHAR` for the rest.

### 19. How do you use string functions in SQL to process text?

Various string functions help manipulate names and other text. For example:

- 
The `LENGTH()` function displays the number of characters in a name.
- 
`UPPER()` and `LOWER()` convert text to uppercase or lowercase.
- 
`CONCAT()` combines the first name and last name into a single column.
- 
`SUBSTRING()` extracts specific parts of a text. For example, isolating the month from a birthdate.

Example query:

```
SELECT 
    UPPER(first_name) AS upper_name,
    CONCAT(first_name, ' ', last_name) AS full_name,
    SUBSTRING(birthdate, 6, 2) AS birth_month,
    TRIM(last_name) AS trimmed_last_name,
    REPLACE(first_name, 'a', '@') AS replaced_name
FROM employees;
```
This query:

- 
Converts first names to uppercase.
- 
Combines first name and last name into a full name.
- 
Extracts the month from the `birthdate` column.
- 
Removes extra spaces from names.
- 
Replaces all occurrences of "a" with "@" in first names.

### 20. How do you update a specific row in SQL?

Use the `UPDATE` statement with a `WHERE` clause to identify the record to modify. 

For example, to update the genre of the movie "Inception" (2010) to "Sci-Fi":

```
UPDATE movies
SET genre = 'Sci-Fi'
WHERE movie_title = 'Inception' AND year = 2010;
```
Here, `UPDATE movies` specifies the target table, and the `WHERE` clause targets the row where the title is "Inception" and the year is "2010". 

## Advanced MySQL Questions

Advanced questions assess your ability to handle complex MySQL scenarios and give insight into your decision-making approach.

### 21. What is a trigger in MySQL? How do you implement it?

A trigger is a set of actions executed when an event occurs in the database. Triggers can run before or after events such as `INSERT`, `UPDATE`, or `DELETE`. 

For example, suppose there is an `orders` table where new orders are added. We can create a trigger that logs each order in an `order_history` table:

```
CREATE TRIGGER after_order_insert
AFTER INSERT ON orders
FOR EACH ROW
BEGIN
    INSERT INTO order_history (order_id, action, timestamp)
    VALUES (NEW.order_id, 'inserted', NOW());
END;
```
After the trigger executes, the `order_history` table is automatically updated:

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

`mysqlcheck --repair database_name table_name -u root -p`
Be careful: in cases of severe corruption, data loss is possible. Remember to back up beforehand.

## Scenario-Based and Troubleshooting MySQL Questions

These questions assess your experience with complex real-world scenarios and your troubleshooting sense.

### 32. Describe a scenario where you used subqueries in MySQL.

Here is a sample answer:

In my previous role, I managed the database of an e-commerce store and had to prepare a product report. The goal was to identify products generating above-average sales, which involved a multi-step analysis using subqueries.

Here is the SQL query I designed:

```
SELECT 
    p.product_id,
    p.product_name,
    s.sales_amount
FROM products p
JOIN sales s ON p.product_id = s.product_id
WHERE s.sales_amount > (
    SELECT AVG(sales_amount)
    FROM sales
)
ORDER BY s.sales_amount DESC;
```
I first established a threshold by calculating the average sale across all products using a subquery in the `WHERE` clause. This dynamic threshold served as a reference for evaluating each product.

The main query then joined the products and sales tables to get the useful information, while the `WHERE` clause excluded products below the average. 

Structured this way, the query made it possible to identify the top products in a single operation.

### 33. Explain a situation where you combined data from multiple tables with SQL joins.

Sample answer:

Recently, I worked on a project with two main tables — one containing product sales and the other product details. My mission: create a report with `sales`, `product name`, `category`, and `price`.

To combine the data, I used an INNER JOIN on the common column `product_id` to link transactions and product details:

```
SELECT 
    s.sales_date,
    p.product_name,
    p.category,
    s.quantity_sold,
    p.price
FROM 
    sales s
INNER JOIN 
    products p
ON 
    s.product_id = p.product_id;
```
The report provided a clear view of sales trends, helping stakeholders identify high-performing categories and those needing improvement.

### 34. Do you have experience with triggers? Describe your use of them.

Here is a sample answer:

Yes, I have solid experience with triggers. In my last role, I set up an `AFTER UPDATE` trigger for auditing price changes. 

Specifically, I created a trigger that automatically records the history whenever a product price changes. Here is the SQL script:

```
CREATE TRIGGER tr_AuditPriceChanges
AFTER UPDATE ON Products
FOR EACH ROW
BEGIN
    -- Log only if the price actually changed
    IF OLD.UnitPrice <> NEW.UnitPrice THEN
        INSERT INTO PriceAudit (
            ProductID,
            OldPrice,
            NewPrice,
            ChangedBy,
            ChangeDate,
            PercentageChange
        )
        VALUES (
            NEW.ProductID,
            OLD.UnitPrice,
            NEW.UnitPrice,
            CURRENT_USER(),
            NOW(),
            ROUND(((NEW.UnitPrice - OLD.UnitPrice) / OLD.UnitPrice * 100), 2)
        );
    END IF;
END;
```
The strengths of this solution:

1. The trigger only fires if the price has actually changed.
2. It captures the user who made the change via `SYSTEM_USER`.
3. It calculates the percentage change for reporting.
4. It includes a condition to ignore updates with no price change.

I also added error handling and logging after identifying edge cases with `NULL` prices.

## Tips for preparing for a MySQL interview

If you're just starting out, here are some tips to succeed in your upcoming interviews:

Master the MySQL fundamentals: Learn the fundamentals of databases such as indexing, transactions, and the query optimizer. Understand how MySQL processes queries and handles storage. You'll write more efficient queries and be able to justify your choices in an interview.

Practice hands-on: Install MySQL on your computer and practice regularly. Create test databases, write different types of queries, and try to optimize them. Real practice is the best way to understand and gain confidence.

To go further, explore DataCamp resources:

- For an intro to SQL: Introduction to SQL Course
- To practice: Applying SQL to real-world problems

Familiarize yourself with MySQL tools and integrations: discover MySQL Workbench and other administration and monitoring tools. You can also explore how MySQL integrates with Python and relevant frameworks, to show your ease in a real development environment.

## Conclusion

There you go! We've gone through 34 MySQL interview questions to help you land your next job. Whether you're aiming for a junior role or an advanced data administrator position, you need to master MySQL fundamentals, query optimization, and administration to stand out.

**To deepen your knowledge of other DBMSs, discover** DataCamp's SQL courses.

## Become SQL certified

I'm a content strategist who loves simplifying complex topics. I've helped companies like Splunk, Hackernoon, and Tiiny Host create engaging and informative content for their audiences.
