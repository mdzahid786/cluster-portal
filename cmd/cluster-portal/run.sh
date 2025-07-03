#!/bin/bash

export DB_URL=localhost:3306/clusterdb
export DB_USER=root
export DB_PASSWORD=""

mysql -u$DB_USER -p -e "CREATE DATABASE IF NOT EXISTS clusterdb;"
mysql -u$DB_USER -p clusterdb < ../../schema/init.sql