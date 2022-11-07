# Write your MySQL query statement below
#
# Table: Weather
#
# +---------------+---------+
# | Column Name   | Type    |
# +---------------+---------+
# | id            | int     |
# | recordDate    | date    |
# | temperature   | int     |
# +---------------+---------+
#
# in  {"headers": {"Weather": ["id", "recordDate", "temperature"]}, "rows": {"Weather": [[1, "2015-01-01", 10], [2, "2015-01-02", 25], [3, "2015-01-03", 20], [4, "2015-01-04", 30]]}}
# out {"headers": ["Id"], "values": [[2], [4]]}
#
# Write an SQL query to find all dates' Id with higher temperatures compared to its previous dates (yesterday).

select w1.id as Id from Weather w1
    join Weather w2 on w2.recordDate = SUBDATE(w1.recordDate, 1) and
    w1.temperature > w2.temperature;