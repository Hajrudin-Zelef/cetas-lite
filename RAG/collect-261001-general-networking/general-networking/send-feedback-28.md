---
id: collect-261001-general-networking/general-networking/send-feedback-28
title: "config system session-ttl"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/send-feedback-28.md
source_anchor: ""
source_lines: [1, 36]
sha256: 22ae251999896abc13f45315e031618d14be63f6a17be9377185ce4956f40e71
---

# config system session-ttl

# config system session-ttl

Configure global session TTL timers for this FortiGate.

```
config system session-ttl
    Description: Configure global session TTL timers for this FortiGate.
    set default {user}
    config port
        Description: Session TTL port.
        edit <id>
            set protocol {integer}
            set start-port {integer}
            set end-port {integer}
            set timeout {user}
        next
    end
end
```
                                            ## config system session-ttl

| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| default | Default timeout. | user | Not Specified |  | 

### config port

| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| id | Table entry ID. | integer | Minimum value: 0 Maximum value: 65535 | 0 | 
| protocol | Protocol. | integer | Minimum value: 0 Maximum value: 255 | 0 | 
| start-port | Start port number. | integer | Minimum value: 0 Maximum value: 65535 | 0 | 
| end-port | End port number. | integer | Minimum value: 0 Maximum value: 65535 | 0 | 
| timeout | Session timeout (TTL). | user | Not Specified |  |
