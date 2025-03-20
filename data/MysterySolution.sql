-- SELECT * FROM crime_scene_report
SELECT description FROM crime_scene_report
WHERE city = 'SQL City'
AND
type = 'murder'
AND date >= 20180110 AND date <= 20180120
LIMIT 10    

-- Security footage shows that there were 2 witnesses. The first witness lives at the last house on "Northwestern Dr". The second witness, named Annabel, lives somewhere on "Franklin Ave".


-- SELECT count(*) FROM interview
SELECT transcript FROm interview
WHERE person_id = (SELECT id FROM person
WHERE address_street_name LIKE 'Northwest%'
ORDER BY address_number DESC
LIMIT 1)


-- I heard a gunshot and then saw a man run out. He had a "Get Fit Now Gym" bag. The membership number on the bag started with "48Z". Only gold members have those bags. The man got into a car with a plate that included "H42W".


SELECT transcript FROm interview
WHERE person_id = (
SELECT id FROM person
WHERE address_street_name LIKE 'Franklin%'
AND name LIKE 'Annabel%'
LIMIT 1)

-- I saw the murder happen, and I recognized the killer from my gym when I was working out last week on January the 9th.


SELECT * FROM get_fit_now_member
WHERE id LIKE '48Z%'
AND
membership_status = 'gold'

-- 48Z7A,"28819","Joe Germuska","20160305","gold"
-- 48Z55,"67318","Jeremy Bowers","20160101","gold"

SELECT * FROM get_fit_now_member gf
JOIN
person p
ON gf.person_id = p.id
JOIN drivers_license d
ON p.license_id = d.id
WHERE gf.id LIKE '48Z%'
AND
gf.membership_status = 'gold'

-- 48Z55,"67318","Jeremy Bowers","20160101","gold","67318","Jeremy Bowers","423327","530","Washington Pl, Apt 3A","871539279","423327","30","70","brown","brown","male","0H42W2","Chevrolet","Spark LS"

SELECT * FROM interview
WHERE person_id = 67318

-- I was hired by a woman with a lot of money. I don't know her name but I know she's around 5'5" (65") or 5'7" (67"). She has red hair and she drives a Tesla Model S. I know that she attended the SQL Symphony Concert 3 times in December 2017.

SELECT * FROM person p
JOIN drivers_license d
ON p.license_id = d.id
WHERE d.gender = 'female'
AND d.car_make = 'Tesla'
AND d.car_model = 'Model S'
AND d.hair_color = 'red'

-- 78881,"Red Korb","918773","107","Camerata Dr","961388910","918773","48","65","black","red","female","917UU3","Tesla","Model S"
-- 90700,"Regina George","291182","332","Maple Ave","337169072","291182","65","66","blue","red","female","08CM64","Tesla","Model S"
-- 99716,"Miranda Priestly","202298","1883","Golden Ave","987756388","202298","68","66","green","red","female","500123","Tesla","Model S"





