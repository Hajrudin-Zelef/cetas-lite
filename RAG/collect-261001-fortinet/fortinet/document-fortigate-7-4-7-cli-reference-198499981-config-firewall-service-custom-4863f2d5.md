---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-7-cli-reference-198499981-config-firewall-service-custom-4863f2d5
title: "config firewall service custom"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-7-cli-reference-198499981-config-firewall-service-custom-4863f2d5.md
source_anchor: ""
source_lines: [1, 131]
sha256: f83a9a60e351da1465f5224ecfc675ec6d00c48da241481b11d677d2bdd17908
---

# config firewall service custom

# config firewall service custom

Configure custom services.

```
config firewall service custom
    Description: Configure custom services.
    edit <name>
        set app-category <id1>, <id2>, ...
        set app-service-type [disable|app-id|...]
        set application <id1>, <id2>, ...
        set category {string}
        set check-reset-range [disable|strict|...]
        set color {integer}
        set comment {var-string}
        set fabric-object [enable|disable]
        set fqdn {string}
        set helper [auto|disable|...]
        set icmpcode {integer}
        set icmptype {integer}
        set iprange {user}
        set protocol [TCP/UDP/SCTP|ICMP|...]
        set protocol-number {integer}
        set proxy [enable|disable]
        set sctp-portrange {user}
        set session-ttl {user}
        set tcp-halfclose-timer {integer}
        set tcp-halfopen-timer {integer}
        set tcp-portrange {user}
        set tcp-rst-timer {integer}
        set tcp-timewait-timer {integer}
        set udp-idle-timer {integer}
        set udp-portrange {user}
        set uuid {uuid}
    next
end
```
                                            ## config firewall service custom

| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| app-category `<id>` | Application category ID. Application category id. | integer | Minimum value: 0 Maximum value: 4294967295 |  | 
| app-service-type | Application service type. | option | - | disable | 
|  |  |  |  |  | 
| application `<id>` | Application ID. Application id. | integer | Minimum value: 0 Maximum value: 4294967295 |  | 
| category | Service category. | string | Maximum length: 63 |  | 
| check-reset-range | Configure the type of ICMP error message verification. | option | - | default | 
|  |  |  |  |  | 
| color | Color of icon on the GUI. | integer | Minimum value: 0 Maximum value: 32 | 0 | 
| comment | Comment. | var-string | Maximum length: 255 |  | 
| fabric-object | Security Fabric global object setting. | option | - | disable | 
|  |  |  |  |  | 
| fqdn | Fully qualified domain name. | string | Maximum length: 255 |  | 
| helper | Helper name. | option | - | auto | 
|  |  |  |  |  | 
| icmpcode | ICMP code. | integer | Minimum value: 0 Maximum value: 255 |  | 
| icmptype | ICMP type. | integer | Minimum value: 0 Maximum value: 4294967295 |  | 
| iprange | Start and end of the IP range associated with service. | user | Not Specified |  | 
| name | Custom service name. | string | Maximum length: 79 |  | 
| protocol | Protocol type based on IANA numbers. | option | - | TCP/UDP/SCTP | 
|  |  |  |  |  | 
| protocol-number | IP protocol number. | integer | Minimum value: 0 Maximum value: 254 | 0 | 
| proxy | Enable/disable web proxy service. | option | - | disable | 
|  |  |  |  |  | 
| sctp-portrange | Multiple SCTP port ranges. | user | Not Specified |  | 
| session-ttl | Session TTL (300 - 2764800, 0 = default). | user | Not Specified |  | 
| tcp-halfclose-timer | Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default). | integer | Minimum value: 0 Maximum value: 86400 | 0 | 
| tcp-halfopen-timer | Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default). | integer | Minimum value: 0 Maximum value: 86400 | 0 | 
| tcp-portrange | Multiple TCP port ranges. | user | Not Specified |  | 
| tcp-rst-timer | Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default). | integer | Minimum value: 5 Maximum value: 300 | 0 | 
| tcp-timewait-timer | Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default). | integer | Minimum value: 0 Maximum value: 300 | 0 | 
| udp-idle-timer | Number of seconds before an idle UDP connection times out (0 - 86400 sec, 0 = default). | integer | Minimum value: 0 Maximum value: 86400 | 0 | 
| udp-portrange | Multiple UDP port ranges. | user | Not Specified |  | 
| uuid | Universally Unique Identifier (UUID; automatically assigned but can be manually reset). | uuid | Not Specified | 00000000-0000-0000-0000-000000000000 | 

| Option | Description | 
|---|---|
| *disable* | Disable application type. | 
| *app-id* | Application ID. | 
| *app-category* | Applicatin category. | 

| Option | Description | 
|---|---|
| *disable* | Disable RST range check. | 
| *strict* | Check RST range strictly. | 
| *default* | Using system default setting. | 

| Option | Description | 
|---|---|
| *enable* | Object is set as a security fabric-wide global object. | 
| *disable* | Object is local to this security fabric member. | 

| Option | Description | 
|---|---|
| *auto* | Automatically select helper based on protocol and port. | 
| *disable* | Disable helper. | 
| *ftp* | FTP. | 
| *tftp* | TFTP. | 
| *ras* | RAS. | 
| *h323* | H323. | 
| *tns* | TNS. | 
| *mms* | MMS. | 
| *sip* | SIP. | 
| *pptp* | PPTP. | 
| *rtsp* | RTSP. | 
| *dns-udp* | DNS UDP. | 
| *dns-tcp* | DNS TCP. | 
| *pmap* | PMAP. | 
| *rsh* | RSH. | 
| *dcerpc* | DCERPC. | 
| *mgcp* | MGCP. | 

| Option | Description | 
|---|---|
| *TCP/UDP/SCTP* | TCP, UDP and SCTP. | 
| *ICMP* | ICMP. | 
| *ICMP6* | ICMP6. | 
| *IP* | IP. | 
| *HTTP* | HTTP - for web proxy. | 
| *FTP* | FTP - for web proxy. | 
| *CONNECT* | Connect - for web proxy. | 
| *SOCKS-TCP* | Socks TCP - for web proxy. | 
| *SOCKS-UDP* | Socks UDP - for web proxy. | 
| *ALL* | All - for web proxy. | 

| Option | Description | 
|---|---|
| *enable* | Enable setting. | 
| *disable* | Disable setting. |
