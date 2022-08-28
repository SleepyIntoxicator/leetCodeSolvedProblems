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
# in  {"headers": {"Person": ["id", "email"]}, "rows": {"Person": [[1, "a@b.com"], [2, "c@d.com"], [3, "a@b.com"]]}}
# out {"headers": ["email"], "values": [["a@b.com"]]}
#
# Write an SQL query to report all the duplicate emails.

select email from Person
    group by email
        having count(*) > 1;