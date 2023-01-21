#!/bin/bash
/opt/mssql-tools/bin/sqlcmd -S `echo "$SERVER_DB,$PORT_DB"` -U sa -P Passw0rd -d `echo "$DATABASE"`