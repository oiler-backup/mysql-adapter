# Scheduler

## Overview

The Scheduler is a service responsible for managing database backups and restorations. It uses Kubernetes CronJobs for scheduling backups and Jobs for executing restorations. The service provides a gRPC API for interacting with these operations.

## Components

### BackupServer

`BackupServer` is the main server that handles gRPC requests for backup and restore operations.

#### Methods

- **Backup**: Creates a CronJob to schedule regular backups.
- **Update**: Updates an existing CronJob with new configuration.
- **Restore**: Creates a Job to perform a one-time database restoration.

### JobsCreator

`JobsCreator` is an interface that defines methods for creating and updating Kubernetes resources.

#### Methods

- **CreateCronJob**: Creates a CronJob in Kubernetes.
- **UpdateCronJob**: Updates an existing CronJob with new environment variables.
- **CreateJob**: Creates a Job in Kubernetes.

### JobsStub

`JobsStub` is a utility that builds Kubernetes Job and CronJob objects from configuration data.

#### Methods

- **BuildBackuperCronJob**: Builds a CronJob for database backups.
- **BuildRestorerJob**: Builds a Job for database restorations.

### Config

`Config` stores configuration settings for the Scheduler.

#### Fields

- **SystemNamespace**: Namespace of the Kubernetes Operator core.
- **BackuperVersion**: Docker image version for the backuper.
- **RestorerVersion**: Docker image version for the restorer.
- **Port**: gRPC port for the Scheduler.

## Configuration

Configuration for the Scheduler is loaded from environment variables. Required fields include `SYSTEM_NAMESPACE`. Default values are provided for `BACKUPER_VERSION`, `RESTORER_VERSION`, and `PORT`.

### Example Environment Variables

```bash
export SYSTEM_NAMESPACE=test-system
export BACKUPER_VERSION=myorg/my-backuper:latest
export RESTORER_VERSION=myorg/my-restorer:latest
export PORT=8080