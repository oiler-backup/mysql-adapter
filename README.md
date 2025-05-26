# MySQL Adapter Helm Chart

## Overview

The MySQL Adapter is a Helm chart designed to deploy the `mysql-adapter` service within a Kubernetes cluster. This service facilitates database backup and restoration operations for the `oiler-backup` system. It includes a scheduler component responsible for managing these operations via Kubernetes CronJobs and Jobs, and it communicates through a gRPC API.

## Components

### Scheduler

The Scheduler is a critical component that manages database backups and restorations. It uses Kubernetes CronJobs to schedule periodic backups and Jobs to execute one-time restorations. The Scheduler exposes a gRPC API for external systems to interact with these operations.

#### Key Features

- **Backup Management**: Creates and updates CronJobs for regular database backups.
- **Restore Management**: Initiates Jobs to perform one-time database restorations.
- **gRPC API**: Provides endpoints for backup and restore operations.

#### Configuration

The Scheduler is configured through environment variables, which can be set via the Helm chart's `values.yaml` file.

- **SYSTEM_NAMESPACE**: The namespace where the Kubernetes Operator core resides.
- **BACKUPER_VERSION**: The Docker image version for the backuper.
- **RESTORER_VERSION**: The Docker image version for the restorer.
- **PORT**: The gRPC port for the Scheduler.

### Backuper

The Backuper component is responsible for performing the actual database backup operations. It uses the MySQL CLI tool to dump the database and store the backup file in an S3 bucket.

#### Key Features

- **Database Dump**: Uses the `mysqldump` command to create a backup of the MySQL database.
- **S3 Storage**: Uploads the backup file to an S3 bucket using provided credentials.

To learn more about the Backuper, refer to its [README](/backuper/README.md).

### Restorer

The Restorer component is responsible for restoring a MySQL database from a backup file stored in an S3 bucket. It utilizes the `mysql` CLI tool to restore the database.

#### Key Features

- **Database Connection**: Establishes a connection to the MySQL database using provided credentials.
- **Backup Download**: Downloads the specified backup file from the S3 bucket.
- **Database Restoration**: Restores the database from the downloaded backup file using the `mysql` command.
- **Metrics Reporting**: Reports the status of the restoration operation, including success and duration.

To learn more about the Restorer, refer to its [README](/restorer/README.md).

## Installation

To install the MySQL Adapter Helm chart, follow these steps:

1. Add the repository:
   ```bash
   helm repo add my-repo https://my-repo.com/charts