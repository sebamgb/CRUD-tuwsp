USE tempdb;
GO
DECLARE @SQL nvarchar(1000);
IF EXISTS (SELECT 1 FROM sys.databases WHERE [name] = N'tuwsp')
BEGIN
    SET @SQL = N'USE [tuwsp];

                 ALTER DATABASE [tuwsp] SET SINGLE_USER WITH ROLLBACK IMMEDIATE;
                 USE [tempdb];

                 DROP DATABASE tuwsp;';
    EXEC (@SQL);
END;
GO
CREATE DATABASE [tuwsp];
GO
USE [tuwsp];
GO
IF NOT EXISTS (SELECT * FROM sys.schemas WHERE name = N'[url]') 
EXEC ('CREATE SCHEMA [url]');
GO
IF NOT EXISTS (SELECT * FROM sys.schemas WHERE name = N'[info]') 
EXEC ('CREATE SCHEMA [info]');
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[info].[key_value_labels]') AND type in (N'U'))
DROP TABLE [info].[key_value_labels];
GO
CREATE TABLE [info].[key_value_labels]
(
    [id] NVARCHAR(32),
    [title] NVARCHAR(50) NULL,
    [label_name] NVARCHAR(50) NULL,
    [label_nickname] NVARCHAR(50) NULL,
    [label_email] NVARCHAR(50) NULL,
    [label_phone] NVARCHAR(50) NULL,
    [label_birthday] NVARCHAR(50) NULL,
    [label_country] NVARCHAR(50) NULL,
    [label_password] NVARCHAR(50) NULL,
    [label_confirm_password] NVARCHAR(50) NULL,
    [input_submit] NVARCHAR(50) NULL,
    CONSTRAINT [tuwsp_PK_key_value_labels] PRIMARY KEY ([id])
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[url].[protocols]') AND type in (N'U'))
DROP TABLE [url].[protocols];
GO
CREATE TABLE [url].[protocols]
(
    [id] NVARCHAR(32),
    [protocol] NVARCHAR(15) NOT NULL,
    CONSTRAINT [tuwsp_PK_protocols] PRIMARY KEY ([id])
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[url].[urls]') AND type in (N'U'))
DROP TABLE [url].[urls];
GO
CREATE TABLE [url].[urls]
(
    [id] NVARCHAR(32),
    [domain] NVARCHAR(255) NOT NULL,
    [protocol_id] NVARCHAR(32) NOT NULL,
    FOREIGN KEY([protocol_id]) REFERENCES [url].[protocols]([id]),
    CONSTRAINT [tuwsp_PK_urls] PRIMARY KEY ([id])
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[url].[users]') AND type in (N'U'))
DROP TABLE [url].[users];
GO
CREATE TABLE [url].[users]
(
    [id] NVARCHAR(32),
    [name] NVARCHAR(70) NULL,
    [nick_name] NVARCHAR(50) NOT NULL UNIQUE,
    [url_id] NVARCHAR(32) NULL,
    FOREIGN KEY([url_id]) REFERENCES [url].[urls]([id]),
    CONSTRAINT [tuwsp_PK_users] PRIMARY KEY (id)
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[info].[forms]') AND type in (N'U'))
DROP TABLE [info].[forms];
GO
CREATE TABLE [info].[forms]
(
    [id] NVARCHAR(32),
    [app] NVARCHAR(20),
    [key_value] NVARCHAR(32) NULL,
    FOREIGN KEY ([key_value]) REFERENCES [info].[key_value_labels]([id]),
    CONSTRAINT [tuwsp_PK_forms] PRIMARY KEY ([id])
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[info].[key_value_placeholders]') AND type in (N'U'))
DROP TABLE [info].[key_value_placeholders];
GO
CREATE TABLE [info].[key_value_placeholders]
(
    [id] NVARCHAR(32),
    [placeholder_name] NVARCHAR(50) NULL,
    [placeholder_nickname] NVARCHAR(50) NULL,
    [placeholder_email] NVARCHAR(50) NULL,
    [placeholder_phone] NVARCHAR(50) NULL,
    [placeholder_birthday] NVARCHAR(50) NULL,
    [placeholder_country] NVARCHAR(50) NULL,
    [placeholder_password] NVARCHAR(50) NULL,
    [placeholder_confirm_password] NVARCHAR(50) NULL,
    [label_id] NVARCHAR(32) NOT NULL,
    FOREIGN KEY ([label_id]) REFERENCES [info].[key_value_labels]([id]),
    CONSTRAINT [tuwsp_PK_key_value_placeholders] PRIMARY KEY ([id])
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[info].[signups]') AND type in (N'U'))
DROP TABLE [info].[signups];
GO
CREATE TABLE [info].[signups]
(
    [id] NVARCHAR(32),
    [title] NVARCHAR(50) NOT NULL,
    [url] NVARCHAR(255) NOT NULL,
    [method] NVARCHAR(10) NOT NULL,
    [name] NVARCHAR(70),
    [nick_name] NVARCHAR(50),
    [email] NVARCHAR(255) NOT NULL UNIQUE,
    [phone] INT,
    [password] NVARCHAR(250) NOT NULL,
    [confirm_password] NVARCHAR(250) NOT NULL,
    [form_id] NVARCHAR(32) NOT NULL,
    FOREIGN KEY ([form_id]) REFERENCES [info].[forms]([id]),
    CONSTRAINT [tuwsp_PK_signups] PRIMARY KEY ([id])
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[info].[auths]') AND type in (N'U'))
DROP TABLE [info].[auths]
GO
CREATE TABLE [info].[auths]
(
    [id] NVARCHAR(32),
    [email] NVARCHAR(255) NOT NULL UNIQUE,
    [created_at] DATETIME2 NOT NULL,
    [password] NVARCHAR(255) NOT NULL,
    [role] CHAR(6),
    [signup_id] NVARCHAR(32) NOT NULL,
    FOREIGN KEY ([signup_id]) REFERENCES [info].[signups]([id]),
    CONSTRAINT [tuwsp_PK_auths] PRIMARY KEY (id)
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[info].[logins]') AND type in (N'U'))
DROP TABLE [info].[logins];
GO
CREATE TABLE [info].[logins]
(
    [id] NVARCHAR(32),
    [title] NVARCHAR(50) NOT NULL,
    [url] NVARCHAR(255) NOT NULL,
    [method] NVARCHAR(10) NOT NULL,
    [email] NVARCHAR(250) NOT NULL UNIQUE,
    [password] NVARCHAR(32) NOT NULL,
    [created_at] DATETIME2,
    [log_out] DATETIME2,
    [auth_id] NVARCHAR(32) NOT NULL,
    FOREIGN KEY ([auth_id]) REFERENCES [info].[auths]([id]),
    [form_id] NVARCHAR(32) NOT NULL,
    FOREIGN KEY ([form_id]) REFERENCES [info].[forms]([id]),
    CONSTRAINT [tuwsp_PK_logins] PRIMARY KEY ([id])
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[url].[query_values]') AND type in (N'U'))
DROP TABLE [url].[info_users];
GO
CREATE TABLE [url].[info_users]
(
    [id] NVARCHAR(32),
    [phone] INT NOT NULL UNIQUE,
    [country] NVARCHAR(20) NOT NULL,
    [cod_country] CHAR(3) NOT NULL,
    [user_id] NVARCHAR(32),
    FOREIGN KEY([user_id]) REFERENCES [url].[users]([id]),
    CONSTRAINT [tuwsp_PK_info_users] PRIMARY KEY ([id])
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[url].[query_values]') AND type in (N'U'))
DROP TABLE [url].[query_values];
GO
CREATE TABLE [url].[query_values]
(
    [number] INT NOT NULL,
    [id] TINYINT NOT NULL,
    [value_param] NTEXT NOT NULL,
    [user_id] NVARCHAR(32) NOT NULL,
    FOREIGN KEY([user_id]) REFERENCES [url].[users]([id]),
    CONSTRAINT [tuwsp_PK_query_values] PRIMARY KEY ([number])
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[url].[query_keys]') AND type in (N'U'))
DROP TABLE [url].[query_keys];
GO
CREATE TABLE [url].[query_keys]
(
    [number] INT NOT NULL,
    [id] TINYINT NOT NULL,
    [key_param] NVARCHAR(255) NOT NULL,
    [url_id] NVARCHAR(32) NOT NULL,
    FOREIGN KEY([url_id]) REFERENCES [url].[urls]([id]),
    CONSTRAINT [tuwsp_PK_query_keys] PRIMARY KEY ([number])
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[url].[endpoints]') AND type in (N'U'))
DROP TABLE [url].[endpoints];
GO
CREATE TABLE [url].[endpoints]
(
    [number] INT NOT NULL,
    [id] TINYINT NOT NULL,
    [endpoint] NVARCHAR(255) NOT NULL,
    [url_id] NVARCHAR(32) NOT NULL,
    FOREIGN KEY([url_id]) REFERENCES [url].[urls]([id]),
    CONSTRAINT [tuwsp_PK_endponts] PRIMARY KEY ([number])
);
GO
IF  EXISTS (SELECT * FROM sys.objects WHERE object_id = OBJECT_ID(N'[info].[dashboards]') AND type in (N'U'))
DROP TABLE [info].[dashboards];
GO
CREATE TABLE [info].[dashboards]
(
    [id] NVARCHAR(32),
    [title] NVARCHAR(50) NOT NULL,
    [app] NVARCHAR(20),
    [key_value] NVARCHAR(32) NULL,
    FOREIGN KEY ([key_value]) REFERENCES [info].[key_value_labels]([id]),
    [owner] NVARCHAR(32) NOT NULL,
    FOREIGN KEY ([owner]) REFERENCES [info].[auths]([id]),
    CONSTRAINT [tuwsp_PK_dashboards] PRIMARY KEY ([id])
);
GO