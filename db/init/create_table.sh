#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table test if not exists (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    name varchar(50) NOT NULL,
    );"
$CMD_MYSQL -e  "insert into article values (1, 'マキマ');"
$CMD_MYSQL -e  "insert into article values (1, 'デンジ');"
