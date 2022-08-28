# Write your MySQL query statement below

# +--------------+---------+
# | Column Name  | Type    |
# +--------------+---------+
# | player_id    | int     |
# | device_id    | int     |
# | event_date   | date    |
# | games_played | int     |
# +--------------+---------+

# in  {"headers":{"Activity":["player_id","device_id","event_date","games_played"]},"rows":{"Activity":[[1,2,"2016-03-01",5],[1,2,"2016-05-02",6],[2,3,"2017-06-25",1],[3,1,"2016-03-02",0],[3,4,"2018-07-03",5]]}}
# out {"headers": ["player_id", "first_login"], "values": [[1, "2016-03-01"], [2, "2017-06-25"], [3, "2016-03-02"]]}

# Write an SQL query to report the first login date for each player.

select player_id, min(event_date) as "first_login"
    from Activity
    group by player_id
    order by player_id