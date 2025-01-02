ALTER TABLE `cars`
    DROP COLUMN `location`,       
    ADD COLUMN `latitude` VARCHAR(255) NOT NULL,  
    ADD COLUMN `longitude` VARCHAR(255) NOT NULL;
