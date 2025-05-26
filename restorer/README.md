# Restorer

The `restorer` package is responsible for restoring a MySQL database from a backup file stored in an S3 bucket. It utilizes the `mysql` CLI tool to restore the database. The process involves several key steps:
1. **Configuration**: The `config` package reads environment variables to configure the database connection details, S3 credentials, and other settings required for the restoration process.
2. **Database Connection**: The `Restorer` struct in the `restorer` package establishes a connection to the MySQL database using the provided credentials.
3. **Backup Download**: The `Restore` method of the `Restorer` struct downloads the specified backup file from the S3 bucket using the `s3base` package.
4. **Database Restoration**: After the backup file is downloaded locally, it is restored to the MySQL database using the `mysql` command.
5. **Metrics Reporting**: The `metricsbase` package is used to report the status of the restoration operation, including whether it was successful and the time taken to complete the restoration.

### Usage
To use the `restorer` package, you need to set the required environment variables and then call the `main` function. The `main` function initializes the logger, reads the configuration, downloads the backup from S3, restores it to the database, and reports the status.

#### Environment Variables
- `DB_HOST`: Hostname of the MySQL database.
- `DB_PORT`: Port number of the MySQL database.
- `DB_USER`: Username for the MySQL database.
- `DB_PASSWORD`: Password for the MySQL database.
- `DB_NAME`: Name of the database to restore.
- `CORE_ADDR`: URI of the Kubernetes Operator core.
- `S3_ENDPOINT`: Endpoint of the S3 service.
- `S3_ACCESS_KEY`: Access key for S3.
- `S3_SECRET_KEY`: Secret key for S3.
- `S3_BUCKET_NAME`: Name of the S3 bucket where the backup is stored.
- `BACKUP_REVISION`: Revision of the backup to restore.
- `SECURE`: Boolean flag to enable or disable TLS/SSL encryption (default: false).
