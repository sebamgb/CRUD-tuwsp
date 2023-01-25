USE master;
GO
IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = N'tuwsp')
BEGIN
    CREATE DATABASE tuwsp
END
GO
USE tuwsp;
GO
IF NOT EXISTS (SELECT * FROM sys.schemas WHERE name = N'url') 
BEGIN
    EXEC ('CREATE SCHEMA url')
END
GO
IF NOT EXISTS (SELECT * FROM sys.schemas WHERE name = N'info') 
BEGIN
    EXEC ('CREATE SCHEMA info')
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'forms' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'info'))
BEGIN
    CREATE TABLE info.forms
    (
        id NVARCHAR(32),
        title NVARCHAR(50) NOT NULL,
        app NVARCHAR(20),
        key_value NTEXT,
        CONSTRAINT tuwsp_PK_forms PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'protocols' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'url'))
BEGIN
    CREATE TABLE url.protocols
    (
        id NVARCHAR(32),
        protocol NVARCHAR(15) NOT NULL,
        CONSTRAINT tuwsp_PK_protocols PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'auths' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'info'))
BEGIN
    CREATE TABLE info.auths 
    (
        id NVARCHAR(32),
        email NVARCHAR(255) NOT NULL UNIQUE,
        created_at DATETIME2 NOT NULL,
        password NVARCHAR(255) NOT NULL,
        CONSTRAINT tuwsp_PK_auths PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'urls' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'url'))
BEGIN
    CREATE TABLE url.urls 
    (
        id NVARCHAR(32),
        domain NVARCHAR(255) NOT NULL,
        protocol_id NVARCHAR(32) NOT NULL,
        FOREIGN KEY(protocol_id) REFERENCES url.protocols(id),
        CONSTRAINT tuwsp_PK_urls PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'users' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'url'))
BEGIN
    CREATE TABLE url.users 
    (
        id NVARCHAR(32),
        name NVARCHAR(70) NULL,
        nick_name NVARCHAR(50) NOT NULL UNIQUE,
        url_id NVARCHAR(32) NULL,
        FOREIGN KEY(url_id) REFERENCES url.urls(id),
        CONSTRAINT tuwsp_PK_users PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'info_users' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'url'))
BEGIN
    CREATE TABLE url.info_users 
    (
        id NVARCHAR(32),
        phone INT NOT NULL UNIQUE,
        country NVARCHAR(20) NOT NULL,
        cod_country CHAR(2) NOT NULL,
        birthday DATE,
        user_id NVARCHAR(32),
        FOREIGN KEY(user_id) REFERENCES url.users(id),
        CONSTRAINT tuwsp_PK_info_users PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'query_values' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'url'))
BEGIN
    CREATE TABLE url.query_values 
    (
        number INT NOT NULL,
        id TINYINT NOT NULL,
        value_param NTEXT NOT NULL,
        user_id NVARCHAR(32) NOT NULL,
        FOREIGN KEY(user_id) REFERENCES url.users(id),
        CONSTRAINT tuwsp_PK_query_values PRIMARY KEY (number)
    )
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'query_keys' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'url'))
BEGIN
    CREATE TABLE url.query_keys 
    (
        number INT NOT NULL,
        id TINYINT NOT NULL,
        key_param NVARCHAR(255) NOT NULL,
        url_id NVARCHAR(32) NOT NULL,
        FOREIGN KEY(url_id) REFERENCES url.urls(id),
        CONSTRAINT tuwsp_PK_query_keys PRIMARY KEY (number)
    )
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'endpoints' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'url'))
BEGIN
    CREATE TABLE url.endpoints 
    (
        number INT NOT NULL,
        id TINYINT NOT NULL,
        endpoint NVARCHAR(255) NOT NULL,
        url_id NVARCHAR(32) NOT NULL,
        FOREIGN KEY(url_id) REFERENCES url.urls(id),
        CONSTRAINT tuwsp_PK_endponts PRIMARY KEY (number)
    )
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'dashboards' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'info'))
BEGIN
    CREATE TABLE info.dashboards
    (
        id NVARCHAR(32),
        title NVARCHAR(50) NOT NULL,
        menu NTEXT,
        app NVARCHAR(20),
        owner NVARCHAR(32) NOT NULL,
        FOREIGN KEY (owner) REFERENCES info.auths(id),
        CONSTRAINT tuwsp_PK_dashboards PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'logins' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'info'))
BEGIN
    CREATE TABLE info.logins
    (
        id NVARCHAR(32),
        email NVARCHAR(250) NOT NULL UNIQUE,
        password NVARCHAR(32) NOT NULL,
        created_at DATETIME2,
        log_out DATETIME2,
        auth_id NVARCHAR(32) NOT NULL,
        form_id NVARCHAR(32) NOT NULL,
        FOREIGN KEY (auth_id) REFERENCES info.auths(id),
        FOREIGN KEY (form_id) REFERENCES info.forms(id),
        CONSTRAINT tuwsp_PK_logins PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM sys.tables
 WHERE name = 'signups' AND schema_id = (SELECT schema_id FROM sys.schemas
 WHERE name = N'info'))
BEGIN
    CREATE TABLE info.signups
    (
        id NVARCHAR(32),
        name NVARCHAR(70),
        nick_name NVARCHAR(50),
        email NVARCHAR(255) NOT NULL UNIQUE,
        phone INT,
        password NVARCHAR(250) NOT NULL,
        confirm_password NVARCHAR(250) NOT NULL,
        form_id NVARCHAR(32) NOT NULL,
        FOREIGN KEY (form_id) REFERENCES info.forms(id),
        CONSTRAINT tuwsp_PK_signups PRIMARY KEY (id)
    )
END
GO