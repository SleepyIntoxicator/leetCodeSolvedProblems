# Please write a DELETE statement and DO NOT write a SELECT statement.
# Write your MySQL query statement below
#
# Table: Person
#
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | id          | int     |
# | email       | varchar |
# +-------------+---------+
#
# in  {"headers": {"Person": ["id", "email"]}, "rows": {"Person": [[1, "john@example.com"], [2, "bob@example.com"], [3, "john@example.com"]]}}
# out {"headers": ["id", "email"], "values": [[1, "john@example.com"], [2, "bob@example.com"]]}
#
# Write an SQL query to delete all the duplicate emails, keeping only one unique email with the smallest id. Note that you are supposed to write a DELETE statement and not a SELECT one.

delete p1 from Person p1, Person p2
    where p1.id > p2.id and p1.email = p2.email;

# Alternative solution

delete p1 from Person p1
    left join Person p2 on p1.email = p2.email
    where p1.id > p2.id;
