CREATE PROCEDURE [dbo].[ExecuteFullDatabaseBackup]
AS

DECLARE @Location NVARCHAR(MAX) = 'C:\DatabaseBackups\BuildSystem_' + 
    CONVERT(VARCHAR(8),GETDATE(),112) + 
    REPLACE(CONVERT(VARCHAR(8),GETDATE(),108), ':','_') + 
    '.Bak'

BACKUP DATABASE [BuildSystem]  
TO DISK = @Location
   WITH FORMAT,  
      MEDIANAME = 'C_DatabaseBackups',  
      NAME = 'Full Backup of BuildSystem';  

RETURN 0
