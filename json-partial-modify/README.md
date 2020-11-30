# Problem

Modify only partial of JSON.

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


# Output

```
C:\sk\prog\prog-go>go run json-partial-modify\main.go
go: finding module for package github.com/tidwall/sjson
go: finding module for package github.com/tidwall/gjson
go: finding module for package github.com/mitchellh/mapstructure
go: finding module for package github.com/tidwall/pretty        
go: found github.com/mitchellh/mapstructure in github.com/mitchellh/mapstructure v1.4.0
go: found github.com/tidwall/gjson in github.com/tidwall/gjson v1.6.3
go: found github.com/tidwall/pretty in github.com/tidwall/pretty v1.0.2
go: found github.com/tidwall/sjson in github.com/tidwall/sjson v1.1.2
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