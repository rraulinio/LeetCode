# Write your MySQL query statement below
SELECT DISTINCT a.Email FROM Person a WHERE (SELECT COUNT(b.Email) FROM Person b WHERE b.Email = a.Email) > 1
