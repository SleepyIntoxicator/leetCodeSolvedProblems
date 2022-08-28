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
# in  {"headers": {"Employee": ["id", "salary"]}, "argument": 2, "rows": {"Employee": [[1, 100], [2, 200], [3, 300]]}}
# out {"headers": ["getNthHighestSalary(2)"], "values": [[200]]}
#
# Write an SQL query to report the nth highest salary from the Employee table. If there is no nth highest salary, the query should report null.

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    SET @S := CONCAT('getNthHighestSalary(', N,')');
    SET N = N - 1;
    RETURN (
        # Write your MySQL query statement below.
        select (
                   select salary from Employee
                   group by salary
                   order by salary desc
                   limit 1 OFFSET N
               ) as S
    );
END