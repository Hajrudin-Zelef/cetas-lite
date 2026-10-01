---
id: collect-261001-fortinet/fortinet/index-php-document-fortigate-7-4-0-cli-reference-58620-1eb4bca6
title: "config system session-ttl"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/index-php-document-fortigate-7-4-0-cli-reference-58620-1eb4bca6.md
source_anchor: ""
source_lines: [1, 31]
sha256: 205778d5e4a5676648c9aa87056300e40e8746b4522636f994e6ce6aeaf84c93
---

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
                                            
                                            | Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| default | Default timeout. | user | Not Specified |  | 

| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| id | Table entry ID. | integer | Minimum value: 0 Maximum value: 65535 | 0 | 
| protocol | Protocol. | integer | Minimum value: 0 Maximum value: 255 | 0 | 
| start-port | Start port number. | integer | Minimum value: 0 Maximum value: 65535 | 0 | 
| end-port | End port number. | integer | Minimum value: 0 Maximum value: 65535 | 0 | 
| timeout | Session timeout (TTL). | user | Not Specified |  |
