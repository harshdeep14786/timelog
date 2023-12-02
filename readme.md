run the file using "localhost:9090/current_time"





#Setup Mysql mysql -u root -p

CREATE USER 'kaurharshdeep'@'localhost' IDENTIFIED BY 'kaurharsh123';

GRANT ALL PRIVILEGES ON . TO 'kaurharshdeep'@'localhost' WITH GRANT OPTION;

#create database CREATE DATABASE goAPI; USE DATABASE goAPI;

#create table CREATE TABLE time_log ( id INT AUTO_INCREMENT PRIMARY KEY, timestamp DATETIME NOT NULL );

install mysql driver
go get -u github.com/go-sql-driver/mysql

Previously we have done the Toronto time assignment. Based on that assignment we have added one
functionality to add the time log in the database. We have created the database name goAPI. The
username for the database is harshdeep and password is Kaurharshdeep@123. We have inserted the
time stamp in the database for the history log. After then we have created the docker repository with
the name timelog .
