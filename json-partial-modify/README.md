# Problem

Modify part of JSON message.

Input:

```json
{
    "id": 1,
    "name": "Agent Emulator",
    "url": "https://jira.com/agent",
    "createdAt": "2020-11-30T17:36:00Z",
    "tasks": [
        {
            "taskId": 1,
            "url": "/1"
        },
        {
            "taskId": 3,
            "url": "/3"
        },
        {
            "taskId": 5,
            "url": "/5"
        },
        {
            "taskId": 2,
            "url": "/2"
        }
    ]
}
```

Modify tasks' urls to full url: from `/1` to `https://jira.com/agent/1`.

# Benchmarks

```
BenchmarkByStd-8
   55593	     21087 ns/op	    5898 B/op	     132 allocs/op
BenchmarkByMapStructure-8
   30272	     39272 ns/op	   10404 B/op	     232 allocs/op
BenchmarkByMapStructurePartial-8
   39292	     29995 ns/op	    7987 B/op	     179 allocs/op
BenchmarkBySjson-8
  144561	      8551 ns/op	    1968 B/op	      21 allocs/op
```

# Output

```
C:\sk\prog\prog-go> go run json-partial-modify\main.go
Source JSON:    
{
        "id": 1,
        "name": "Agent Emulator",
        "url": "https://jira.com/agent",
        "createdAt": "2020-11-30T17:36:00Z",
        "tasks": [
                {
                        "taskId": 1,
                        "url": "/1"
                },
                {
                        "taskId": 3,
                        "url": "/3"
                },
                {
                        "taskId": 5,
                        "url": "/5"
                },
                {
                        "taskId": 2,
                        "url": "/2"
                }
        ]
}

==== byStd
{
  "createdAt": "2020-11-30T17:36:00Z",
  "id": 1,
  "name": "Agent Emulator",
  "tasks": [
    {
      "taskId": 1,
      "url": "https://jira.com/agent/1"
    },
    {
      "taskId": 3,
      "url": "https://jira.com/agent/3"
    },
    {
      "taskId": 5,
      "url": "https://jira.com/agent/5"
    },
    {
      "taskId": 2,
      "url": "https://jira.com/agent/2"
    }
  ],
  "url": "https://jira.com/agent"
}

==== byMapStructure
{
  "createdAt": "2020-11-30T17:36:00Z",
  "id": 1,
  "name": "Agent Emulator",
  "tasks": [
    {
      "taskId": 1,
      "url": "https://jira.com/agent/1"
    },
    {
      "taskId": 3,
      "url": "https://jira.com/agent/3"
    },
    {
      "taskId": 5,
      "url": "https://jira.com/agent/5"
    },
    {
      "taskId": 2,
      "url": "https://jira.com/agent/2"
    }
  ],
  "url": "https://jira.com/agent"
}

==== byMapStructurePartial
{
  "createdAt": "2020-11-30T17:36:00Z",
  "id": 1,
  "name": "Agent Emulator",
  "tasks": [
    {
      "taskId": 1,
      "url": "https://jira.com/agent/1"
    },
    {
      "taskId": 3,
      "url": "https://jira.com/agent/3"
    },
    {
      "taskId": 5,
      "url": "https://jira.com/agent/5"
    },
    {
      "taskId": 2,
      "url": "https://jira.com/agent/2"
    }
  ],
  "url": "https://jira.com/agent"
}

==== bySjson
{
  "id": 1,
  "name": "Agent Emulator",
  "url": "https://jira.com/agent",
  "createdAt": "2020-11-30T17:36:00Z",
  "tasks": [
    {
      "taskId": 1,
      "url": "https://jira.com/agent/1"
    },
    {
      "taskId": 3,
      "url": "https://jira.com/agent/3"
    },
    {
      "taskId": 5,
      "url": "https://jira.com/agent/5"
    },
    {
      "taskId": 2,
      "url": "https://jira.com/agent/2"
    }
  ]
}
```