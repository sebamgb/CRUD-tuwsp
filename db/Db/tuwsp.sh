#!/bin/bash
/opt/mssql-tools/bin/sqlcmd -S `echo "$SERVER_DB,$PORT_DB"` -U `echo "$USER_DB"` -P `echo "$PASSWORD_DB"` -d `echo "$DATABASE"`