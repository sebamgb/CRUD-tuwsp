USE master;
GO
IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = N'tuwsp')
BEGIN
    CREATE DATABASE tuwsp
END
GO
USE tuwsp;
GO
IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE table_name = 'protocols') 
BEGIN
    CREATE TABLE protocols 
    (
        id NVARCHAR(32),
        protocol NVARCHAR(15) NOT NULL,
        CONSTRAINT tuwsp_PK_protocols PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE table_name = 'auths') 
BEGIN
    CREATE TABLE auths 
    (
        id NVARCHAR(32),
        email NVARCHAR(255) NOT NULL UNIQUE,
        created_at DATETIME2 NOT NULL,
        password NVARCHAR(255) NOT NULL,
        CONSTRAINT tuwsp_PK_auths PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE table_name = 'urls') 
BEGIN
    CREATE TABLE urls 
    (
        id NVARCHAR(32),
        domain NVARCHAR(255) NOT NULL,
        protocol_id NVARCHAR(32) NOT NULL,
        FOREIGN KEY(protocol_id) REFERENCES protocols(id),
        CONSTRAINT tuwsp_PK_urls PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE table_name = 'users') 
BEGIN
    CREATE TABLE users 
    (
        id NVARCHAR(32),
        name NVARCHAR(70) NULL,
        nick_name NVARCHAR(50) NOT NULL UNIQUE,
        url_id NVARCHAR(32) NULL,
        FOREIGN KEY(url_id) REFERENCES urls(id),
        CONSTRAINT tuwsp_PK_users PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE table_name = 'info_users') 
BEGIN
    CREATE TABLE info_users 
    (
        id NVARCHAR(32),
        phone INT NOT NULL,
        country NVARCHAR(20) NOT NULL,
        cod_country CHAR(2) NOT NULL,
        birthday DATE,
        user_id NVARCHAR(32),
        FOREIGN KEY(user_id) REFERENCES users(id),
        CONSTRAINT tuwsp_PK_info_users PRIMARY KEY (id)
    )
END
GO
IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE table_name = 'query_values') 
BEGIN
    CREATE TABLE query_values 
    (
        number INT IDENTITY(1,1),
        id TINYINT NOT NULL,
        value_param NTEXT NOT NULL,
        url_id NVARCHAR(32) NOT NULL,
        FOREIGN KEY(url_id) REFERENCES urls(id),
        CONSTRAINT tuwsp_PK_query_values PRIMARY KEY (number)
    )
END
GO
IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE table_name = 'query_keys') 
BEGIN
    CREATE TABLE query_keys 
    (
        number INT IDENTITY(1,1),
        id TINYINT NOT NULL,
        key_param NVARCHAR(255) NOT NULL,
        url_id NVARCHAR(32) NOT NULL,
        FOREIGN KEY(url_id) REFERENCES urls(id),
        CONSTRAINT tuwsp_PK_query_keys PRIMARY KEY (number)
    )
END
GO
IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE table_name = 'endponts') 
BEGIN
    CREATE TABLE endponts 
    (
        number INT IDENTITY(1,1),
        id TINYINT NOT NULL,
        endpoint NVARCHAR(255) NOT NULL,
        url_id NVARCHAR(32) NOT NULL,
        FOREIGN KEY(url_id) REFERENCES urls(id),
        CONSTRAINT tuwsp_PK_endponts PRIMARY KEY (number)
    )
END
GO