# PostgreSQL-DB-connection with GO
* In this project you will get to know how to connect your postgres database with your GO project.
* Before moving ahead please make sure you have installed postgres in your system.
* Reference link for postgres installation -> `https://youtu.be/QGLIZRG_aMI` or follow this `https://www.codingcommanders.com/postgresql/install.php` 
### Run the below commands in your terminal to setup the postgres database user, password and database name, again make sure postgres is installed properly.
- sudo su postgres
- psql
- ALTER USER postgres WITH PASSWORD 'postgres';
- CREATE DATABASE golang_database;
- ALTER DATABASE golang_database OWNER TO postgres;