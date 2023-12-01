run the file using "localhost:9090/current_time"





#Setup Mysql mysql -u root -p

CREATE USER 'kaurharshdeep'@'localhost' IDENTIFIED BY 'kaurharsh123';

GRANT ALL PRIVILEGES ON . TO 'kaurharshdeep'@'localhost' WITH GRANT OPTION;

#create database CREATE DATABASE goAPI; USE DATABASE goAPI;

#create table CREATE TABLE time_log ( id INT AUTO_INCREMENT PRIMARY KEY, timestamp DATETIME NOT NULL );

install mysql driver
go get -u github.com/go-sql-driver/mysql
