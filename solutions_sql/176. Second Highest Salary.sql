# Write your MySQL query statement below
#
# Table: Employee
#
# +-------------+------+
# | Column Name | Type |
# +-------------+------+
# | id          | int  |
# | salary      | int  |
# +-------------+------+
#
# in  {"headers":{"Employee":["id","salary"]},"rows":{"Employee":[[1,100],[2,200],[3,300]]}}
# out {"headers": ["SecondHighestSalary"], "values": [[200]]}
#
# Write an SQL query to report the second highest salary from the Employee table. If there is no second highest salary, the query should report null.

select (select salary from Employee
            group by salary
            order by salary desc
            limit 1, 1) as SecondHighestSalary