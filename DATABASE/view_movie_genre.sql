CREATE VIEW VIEW_MOVIE_GENRE AS 
SELECT
A.idgenre,A.movieid,C.nmgenre,
B.movietitle,B.label,B.movietype,B.views,C.genredisplay
FROM tbl_trx_moviegenre as A 
JOIN tbl_trx_movie as B ON B.movieid = A.movieid
JOIN tbl_mst_movie_genre as C ON C.idgenre = A.idgenre
WHERE B.enabled = 1  