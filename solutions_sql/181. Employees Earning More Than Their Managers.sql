# Write your MySQL query statement below
#
# Table: Employee
#
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | id          | int     |
# | name        | varchar |
# | salary      | int     |
# | managerId   | int     |
# +-------------+---------+
#
# in  {"headers": {"Employee": ["id", "name", "salary", "managerId"]}, "rows": {"Employee": [[1, "Joe", 70000, 3], [2, "Henry", 80000, 4], [3, "Sam", 60000, null], [4, "Max", 90000, null]]}}
# out {"headers": ["Employee"], "values": [["Joe"]]}
#
# Write an SQL query to find the employees who earn more than their managers.

select name as Employee
    from Employee e
    left join
        (select id, salary from Employee) as mgr
        on e.managerId = mgr.id
    where e.salary > mgr.salary;
