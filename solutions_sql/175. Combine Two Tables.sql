# Write your MySQL query statement below
#
# Table: Person
#
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | personId    | int     |
# | lastName    | varchar |
# | firstName   | varchar |
# +-------------+---------+
#
# Table: Address
#
# +-------------+---------+
# | Column Name | Type    |
# +-------------+---------+
# | addressId   | int     |
# | personId    | int     |
# | city        | varchar |
# | state       | varchar |
# +-------------+---------+
#
# in  {"headers":{"Person":["personId","lastName","firstName"],"Address":["addressId","personId","city","state"]},"rows":{"Person":[[1,"Wang","Allen"],[2,"Alice","Bob"]],"Address":[[1,2,"New York City","New York"],[2,3,"Leetcode","California"]]}}
# out {"headers": ["firstName", "lastName", "city", "state"], "values": [["Allen", "Wang", null, null], ["Bob", "Alice", "New York City", "New York"]]}
#
# Write an SQL query to report the first name, last name, city, and state of each person in the Person table. If the address of a personId is not present in the Address table, report null instead.

select firstName, lastName, city, state from
    Person as p
        left join Address as a
            on a.personId = p.personId;