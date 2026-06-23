# Distributed Job Queue - V1 Specification

## Overview

The goal of V1 is to build a simple job queue system that allows producers to submit jobs and workers to process them.

This version focuses on correctness and simplicity. Advanced messaging concepts such as retries, acknowledgements, delayed jobs, priorities, persistence, exchanges, and distributed nodes are intentionally excluded.

---

# Goals

* Accept jobs from producers
* Store jobs in memory
* Allow workers to fetch jobs
* Ensure a job is assigned to only one worker
* Track job status throughout its lifecycle
* Expose simple HTTP APIs

---

# System Architecture

```text
Producer
    |
    | HTTP
    v
Queue Server
    |
    | HTTP
    v
Worker
```

---

# Job Lifecycle

```text
PENDING
    |
    v
PROCESSING
    |
    +----------+
    |          |
    v          v
COMPLETED   FAILED
```

---

# Job Model

```go
type Job struct {
    ID        string
    Payload   string
    Status    string
    CreatedAt time.Time
}
```

## Status Values

| Status     | Description            |
| ---------- | ---------------------- |
| PENDING    | Waiting for processing |
| PROCESSING | Assigned to a worker   |
| COMPLETED  | Successfully processed |
| FAILED     | Processing failed      |

---

# Queue Behavior

## FIFO Ordering

Jobs should be assigned in the same order they were created.

Example:

```text
Job1
Job2
Job3
```

Expected assignment order:

```text
Job1
Job2
Job3
```

---

## Single Delivery

A job may only be assigned to one worker at a time.

Example:

```text
Worker A -> Job1
```

Worker B must not receive Job1 while it is in PROCESSING state.

---

# API Specification

## Create Job

### Request

```http
POST /jobs
```

### Body

```json
{
  "payload": "send email to user"
}
```

### Response

```json
{
  "job_id": "123"
}
```

### Status Code

```http
201 Created
```

---

## Get Job By ID

### Request

```http
GET /jobs/{id}
```

### Response

```json
{
  "id": "123",
  "payload": "send email to user",
  "status": "PENDING",
  "created_at": "2026-06-14T12:00:00Z"
}
```

### Status Code

```http
200 OK
```

---

## Fetch Next Job

Used by workers.

### Request

```http
GET /jobs/next
```

### Success Response

```json
{
  "id": "123",
  "payload": "send email to user"
}
```

### Status Code

```http
200 OK
```

### No Jobs Available

```http
204 No Content
```

Behavior:

* Find oldest PENDING job
* Mark as PROCESSING
* Return job to worker

---

## Mark Job Completed

### Request

```http
POST /jobs/{id}/complete
```

### Response

```json
{
  "message": "job completed"
}
```

### Status Code

```http
200 OK
```

Behavior:

* Change status from PROCESSING to COMPLETED

---

## Mark Job Failed

### Request

```http
POST /jobs/{id}/fail
```

### Response

```json
{
  "message": "job failed"
}
```

### Status Code

```http
200 OK
```

Behavior:

* Change status from PROCESSING to FAILED

---

## Queue Statistics

### Request

```http
GET /stats
```

### Response

```json
{
  "pending": 12,
  "processing": 2,
  "completed": 50,
  "failed": 3
}
```

### Status Code

```http
200 OK
```

---

# Storage

## In-Memory Storage

V1 stores all jobs in memory.

Example:

```go
map[string]*Job
```

or

```go
[]Job
```

No persistence is required.

Restarting the server may result in loss of all jobs.

---

# Worker Behavior

Worker loop:

```text
while true
    fetch next job
    process job
    mark completed
```

Example:

```text
Processing Job #123
Processing Job #124
Processing Job #125
```

---

# Out Of Scope

The following features are explicitly excluded from V1:

* Retries
* Acknowledgements
* Visibility timeouts
* Delayed jobs
* Priority queues
* Dead letter queues
* Exchanges
* Routing keys
* AMQP support
* Kafka compatibility
* Persistent storage
* Database integration
* Distributed nodes
* Cluster membership
* Authentication
* Authorization

These features belong in future versions.

---

# Success Criteria

The implementation is considered complete when:

1. Producers can submit jobs.
2. Jobs are stored in memory.
3. Workers can fetch jobs.
4. Jobs transition through valid states.
5. A job is never assigned to multiple workers simultaneously.
6. Queue statistics can be viewed through an API.
7. End-to-end processing can be demonstrated with multiple jobs and workers.
   """
