# Write your MySQL query statement below
#
# Table: Customers
#
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | id          | int     |
# | name        | varchar |
# +-------------+---------+
#
# Table: Orders
#
# +-------------+------+
# | Column Name | Type |
# +-------------+------+
# | id          | int  |
# | customerId  | int  |
# +-------------+------+
#
# in  {"headers": {"Customers": ["id", "name"], "Orders": ["id", "customerId"]}, "rows": {"Customers": [[1, "Joe"], [2, "Henry"], [3, "Sam"], [4, "Max"]], "Orders": [[1, 3], [2, 1]]}}
# out {"headers": ["Customers"], "values": [["Henry"], ["Max"]]}
#
# Write an SQL query to report all customers who never order anything.

select name as Customers
    from Customers as c
        left join Orders o
            on c.id = o.customerId
    where o.customerId is NULL;

# Alternative solution

select name as Customers
    from Customers as c
    where c.id not in (select customerId from Orders);