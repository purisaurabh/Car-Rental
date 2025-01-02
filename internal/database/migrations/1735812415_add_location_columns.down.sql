ALTER TABLE `cars`
    DROP COLUMN `latitude`,              
    DROP COLUMN `longitude`,             
    ADD COLUMN `location` POINT NOT NULL; 
