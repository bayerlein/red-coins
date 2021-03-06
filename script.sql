USE [RedCoins]
GO
/****** Object:  Table [dbo].[users]    Script Date: 10/25/2018 22:17:30 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[users](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[full_name] [varchar](100) NULL,
	[email] [varchar](50) NULL,
	[date_of_birth] [datetime] NULL,
	[password] [varchar](max) NULL,
	[register_date] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[transaction_type]    Script Date: 10/25/2018 22:17:30 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
SET ANSI_PADDING ON
GO
CREATE TABLE [dbo].[transaction_type](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[description] [varchar](10) NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
SET ANSI_PADDING OFF
GO
/****** Object:  Table [dbo].[bitcoin_transaction]    Script Date: 10/25/2018 22:17:30 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[bitcoin_transaction](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[amount] [decimal](18, 9) NULL,
	[total_value] [decimal](18, 9) NULL,
	[price_used] [decimal](18, 9) NULL,
	[transaction_date] [datetime] NULL,
	[transaction_type] [int] NULL,
	[user_id] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  ForeignKey [FK__bitcoin_t__trans__300424B4]    Script Date: 10/25/2018 22:17:30 ******/
ALTER TABLE [dbo].[bitcoin_transaction]  WITH CHECK ADD FOREIGN KEY([transaction_type])
REFERENCES [dbo].[transaction_type] ([id])
GO
/****** Object:  ForeignKey [FK__bitcoin_t__user___2F10007B]    Script Date: 10/25/2018 22:17:30 ******/
ALTER TABLE [dbo].[bitcoin_transaction]  WITH CHECK ADD FOREIGN KEY([user_id])
REFERENCES [dbo].[users] ([id])
GO
