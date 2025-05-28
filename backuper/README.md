# Backuper

The `backuper` package is responsible for performing backups of a MySQL database. It utilizes the `mysqldump` CLI tool to create a backup and then uploads it to an S3 bucket. The process involves several key steps:

1. **Configuration**: The `config` package reads environment variables to configure the database connection details, S3 credentials, and other settings required for the backup process.

2. **Database Connection**: The `Backuper` struct in the `backuper` package establishes a connection to the MySQL database using the provided credentials.

3. **Backup Execution**: The `Backup` method of the `Backuper` struct executes the `mysqldump` command to create a backup of the specified database. It handles both secure (_TLS/SSL_) and insecure connections based on the configuration.

4. **S3 Upload**: After the backup is created locally, it is uploaded to an S3 bucket using the `s3base` package. The package also ensures that only the specified maximum number of backups are retained in the bucket.

5. **Metrics Reporting**: The `metricsbase` package is used to report the status of the backup operation, including whether it was successful and the time taken to complete the backup.

### Usage

To use the `backuper` package, you need to set the required environment variables and then call the `main` function. The `main` function initializes the logger, reads the configuration, performs the backup, uploads it to S3, and reports the status.

#### Environment Variables

- `DB_HOST`: Hostname of the MySQL database.
- `DB_PORT`: Port number of the MySQL database.
- `DB_USER`: Username for the MySQL database.
- `DB_PASSWORD`: Password for the MySQL database.
- `DB_NAME`: Name of the database to backup.
- `CORE_ADDR`: URI of the Kubernetes Operator core.
- `S3_ENDPOINT`: Endpoint of the S3 service.
- `S3_ACCESS_KEY`: Access key for S3.
- `S3_SECRET_KEY`: Secret key for S3.
- `S3_BUCKET_NAME`: Name of the S3 bucket to store the backup.
- `MAX_BACKUP_COUNT`: Maximum number of backups to retain in the S3 bucket.
- `SECURE`: Boolean flag to enable or disable TLS/SSL encryption (default: false).
