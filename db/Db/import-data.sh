#!/bin/bash
sleep 90s
/opt/mssql-tools/bin/sqlcmd -S `echo "$SERVER_DB,$PORT_DB"` -U sa -P Passw0rd -d master -i up.sql