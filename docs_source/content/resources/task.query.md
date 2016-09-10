---
title: "task.query"
slug: "task-query"
date: "2016-09-08T23:26:25-07:00"
menu:
  main:
    parent: resources
---




## Example

```hcl
task.query "hostname" {
  query = "hostname"
}

file.content "hostname data" {
  destination = "hostname.txt"
  content     = "{{lookup `task.query.hostname.Status.Stdout`}}"
}

```


## Parameters

- `interpreter` (string)

  
- `query` (string)

  
- `check_flags` (list of strings)

  
- `exec_flags` (list of strings)

  
- `timeout` (duration string)

  
- `dir` (string)

  
- `env` (map of string to string)

  
