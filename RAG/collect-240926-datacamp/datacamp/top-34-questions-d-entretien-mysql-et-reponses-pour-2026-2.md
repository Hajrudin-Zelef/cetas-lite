---
id: collect-240926-datacamp/datacamp/top-34-questions-d-entretien-mysql-et-reponses-pour-2026-2
title: "top-34-questions-d-entretien-mysql-et-reponses-pour-2026"
domain: datacamp
role: reference
task: reference
actors: ["United States"]
dates: ["2024-01-01", "2024-12-31", "9999-12-31"]
keywords: []
source: docs/RAG/clean_en/datacamp/top-34-questions-d-entretien-mysql-et-reponses-pour-2026.md
source_anchor: ""
source_lines: [169, 376]
sha256: 3f6f49311805bfa3727d06b6046b6deb8a988c9fa34fa8df4bda2ff0b4936a6a
---

# top-34-questions-d-entretien-mysql-et-reponses-pour-2026

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

