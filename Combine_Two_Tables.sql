# Write your MySQL query statement below
SELECT a.FirstName, a.LastName, b.City, b.State from Person a LEFT JOIN Address b ON a.PersonId = b.PersonId
