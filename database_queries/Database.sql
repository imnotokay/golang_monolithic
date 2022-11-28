CREATE DATABASE customers_db;

USE customers_db;

CREATE TABLE customers (
	id int auto_increment, 
	name varchar(200), 
	typeDoc int, 
	document varchar(200), 
	creationDate datetime, 
	modificationDate datetime,
    constraint pk_customers primary key (id)
);

CREATE TABLE products (
	id int auto_increment, 
    name varchar(200), 
    unitPrice double, 
    tax double, 
    creationDate datetime, 
    modificationDate datetime,
    constraint pk_products primary key (id)
);