---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-5-cli-reference-105110478-config-system-ntp-7f8cbc74
title: "config system ntp"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-5-cli-reference-105110478-config-system-ntp-7f8cbc74.md
source_anchor: ""
source_lines: [1, 130]
sha256: e37a688c8f3d4c90b30bf56c13e1e403a218106aaae2143d102fea6bc5432f77
---

# config system ntp

# config system ntp

Configure system NTP information.

```
config system ntp
    Description: Configure system NTP information.
    set authentication [enable|disable]
    set interface <interface-name1>, <interface-name2>, ...
    set key {password}
    set key-id {integer}
    set key-type [MD5|SHA1|...]
    config ntpserver
        Description: Configure the FortiGate to connect to any available third-party NTP server.
        edit <id>
            set authentication [enable|disable]
            set interface {string}
            set interface-select-method [auto|sdwan|...]
            set ip-type [IPv6|IPv4|...]
            set key {password}
            set key-id {integer}
            set key-type [MD5|SHA1|...]
            set ntpv3 [enable|disable]
            set server {string}
        next
    end
    set ntpsync [enable|disable]
    set server-mode [enable|disable]
    set source-ip {ipv4-address}
    set source-ip6 {ipv6-address}
    set syncinterval {integer}
    set type [fortiguard|custom]
end
```
                                            ## config system ntp

| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| authentication | Enable/disable authentication. | option | - | disable | 
|  |  |  |  |  | 
| interface `<interface-name>` | FortiGate interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services. Interface name. | string | Maximum length: 79 |  | 
| key | Key for authentication. | password | Not Specified |  | 
| key-id | Key ID for authentication. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| key-type | Key type for authentication (MD5, SHA1, SHA256). | option | - | MD5 | 
|  |  |  |  |  | 
| ntpsync | Enable/disable setting the FortiGate system time by synchronizing with an NTP Server. | option | - | disable | 
|  |  |  |  |  | 
| server-mode | Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server. | option | - | disable | 
|  |  |  |  |  | 
| source-ip | Source IP address for communication to the NTP server. | ipv4-address | Not Specified | 0.0.0.0 | 
| source-ip6 | Source IPv6 address for communication to the NTP server. | ipv6-address | Not Specified | :: | 
| syncinterval | NTP synchronization interval. | integer | Minimum value: 1 Maximum value: 1440 | 60 | 
| type | Use the FortiGuard NTP server or any other available NTP Server. | option | - | fortiguard | 
|  |  |  |  |  | 

| Option | Description | 
|---|---|
| *enable* | Enable authentication. | 
| *disable* | Disable authentication. | 

| Option | Description | 
|---|---|
| *MD5* | Use MD5 to authenticate the message. | 
| *SHA1* | Use SHA1 to authenticate the message. | 
| *SHA256* | Use SHA256 to authenticate the message. | 

| Option | Description | 
|---|---|
| *enable* | Enable synchronization with NTP Server. | 
| *disable* | Disable synchronization with NTP Server. | 

| Option | Description | 
|---|---|
| *enable* | Enable FortiGate NTP Server Mode. | 
| *disable* | Disable FortiGate NTP Server Mode. | 

| Option | Description | 
|---|---|
| *fortiguard* | Use the FortiGuard NTP server. | 
| *custom* | Use any other available NTP server. | 

### config ntpserver

| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| authentication | Enable/disable authentication. | option | - | disable | 
|  |  |  |  |  | 
| id | NTP server ID. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| interface | Specify outgoing interface to reach server. | string | Maximum length: 15 |  | 
| interface-select-method | Specify how to select outgoing interface to reach server. | option | - | auto | 
|  |  |  |  |  | 
| ip-type | Choose to connect to IPv4 or/and IPv6 NTP server. | option | - | Both | 
|  |  |  |  |  | 
| key | Key for MD5(NTPv3)/SHA1(NTPv4)/SHA256(NTPv4) authentication. | password | Not Specified |  | 
| key-id | Key ID for authentication. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| key-type | Select NTP authentication type. | option | - | MD5 | 
|  |  |  |  |  | 
| ntpv3 | Enable to use NTPv3 instead of NTPv4. | option | - | disable | 
|  |  |  |  |  | 
| server | IP address or hostname of the NTP Server. | string | Maximum length: 63 |  | 

| Option | Description | 
|---|---|
| *enable* | Enable authentication. | 
| *disable* | Disable authentication. | 

| Option | Description | 
|---|---|
| *auto* | Set outgoing interface automatically. | 
| *sdwan* | Set outgoing interface by SD-WAN or policy routing rules. | 
| *specify* | Set outgoing interface manually. | 

| Option | Description | 
|---|---|
| *IPv6* | Enable look up for IPv6 NTP server. | 
| *IPv4* | Enable look up for IPv4 NTP server. | 
| *Both* | Enable look up for both IPv4 and IPv6 NTP server. | 

| Option | Description | 
|---|---|
| *MD5* | Enable MD5(NTPv3) authentication. | 
| *SHA1* | Enable SHA1(NTPv4) authentication. | 
| *SHA256* | Enable SHA256(NTPv4) authentication. | 

| Option | Description | 
|---|---|
| *enable* | Enable NTPv3. | 
| *disable* | Disable NTPv3 (use NTPv4). |
