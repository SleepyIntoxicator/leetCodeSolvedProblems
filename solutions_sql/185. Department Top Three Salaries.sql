# Write your MySQL query statement below
#
# Employee
#
# +--------------+---------+
# | Column Name  | Type    |
# +--------------+---------+
# | id           | int     | pk
# | name         | varchar |
# | salary       | int     |
# | departmentId | int     | fk
# +--------------+---------+
#
# Department
#
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | id          | int     | pk
# | name        | varchar |
# +-------------+---------+

# in {"headers": {"Employee": ["id", "name", "salary", "departmentId"], "Department": ["id", "name"]}, "rows": {"Employee": [[1, "Joe", 85000, 1], [2, "Henry", 80000, 2], [3, "Sam", 60000, 2], [4, "Max", 90000, 1], [5, "Janet", 69000, 1], [6, "Randy", 85000, 1], [7, "Will", 70000, 1]], "Department": [[1, "IT"], [2, "Sales"]]}}
# out {"headers": ["Department", "Employee", "Salary"], "values": [["IT", "Max", 90000], ["IT", "Joe", 85000], ["IT", "Randy", 85000], ["IT", "Will", 70000], ["Sales", "Henry", 80000], ["Sales", "Sam", 60000]]}

# Write an SQL query to find the employees who are high earners in each of the departments.

select D.name as 'Department',
    ME.name as 'Employee',
    ME.salary as 'Salary'
    from (
         select name, salary, departmentId,
                DENSE_RANK() over (PARTITION BY departmentId
                        ORDER BY salary desc) as rn
         from Employee) ME
         left join Department D on ME.departmentId = D.id
    where ME.rn <= 3;