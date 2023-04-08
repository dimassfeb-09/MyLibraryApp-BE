-- Active: 1680419549061@@127.0.0.1@3306@mylibraryapps

SELECT *
FROM book
    JOIN category ON category.id = book.category_id
WHERE
    title LIKE '%Laskar%'
    OR writer LIKE '%Laskar%'
    OR category.name = 'Sastra';
