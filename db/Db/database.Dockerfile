FROM mcr.microsoft.com/mssql/server:2019-CU18-ubuntu-20.04

WORKDIR /usr/src

COPY ["./entrypoint.sh", "/usr/local/bin"]

COPY [ ".", "/usr/src"]

EXPOSE 1433

CMD [ "entrypoint.sh" ]