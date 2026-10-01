---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-cli-reference-614401817-config-user-fsso-e75eafa4
title: "document-fortigate-7-4-8-cli-reference-614401817-config-user-fsso-e75eafa4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-cli-reference-614401817-config-user-fsso-e75eafa4.md
source_anchor: ""
source_lines: [1, 100]
sha256: 35a7792ff7ad6c23db2c1f51aedbc81628dc20bfa934dad0d8202d16dbcd1a33
---

# document-fortigate-7-4-8-cli-reference-614401817-config-user-fsso-e75eafa4

config user fsso
config user fsso
Configure Fortinet Single Sign On (FSSO) agents.
config user fsso
    Description: Configure Fortinet Single Sign On (FSSO) agents.
    edit <name>
        set group-poll-interval {integer}
        set interface {string}
        set interface-select-method [auto|sdwan|...]
        set ldap-poll [enable|disable]
        set ldap-poll-filter {string}
        set ldap-poll-interval {integer}
        set ldap-server {string}
        set logon-timeout {integer}
        set password {password}
        set password2 {password}
        set password3 {password}
        set password4 {password}
        set password5 {password}
        set port {integer}
        set port2 {integer}
        set port3 {integer}
        set port4 {integer}
        set port5 {integer}
        set server {string}
        set server2 {string}
        set server3 {string}
        set server4 {string}
        set server5 {string}
        set sni {string}
        set source-ip {ipv4-address}
        set source-ip6 {ipv6-address}
        set ssl [enable|disable]
        set ssl-server-host-ip-check [enable|disable]
        set ssl-trusted-cert {string}
        set type [default|fortinac]
        set user-info-server {string}
    next
end
                                            config user fsso
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| group-poll-interval | Interval in minutes within to fetch groups from FSSO server, or unset to disable. | integer | Minimum value: 1 Maximum value: 2880 | 0 | 
| interface | Specify outgoing interface to reach server. | string | Maximum length: 15 |  | 
| interface-select-method | Specify how to select outgoing interface to reach server. | option | - | auto | 
|  |  |  |  |  | 
| ldap-poll | Enable/disable automatic fetching of groups from LDAP server. | option | - | disable | 
|  |  |  |  |  | 
| ldap-poll-filter | Filter used to fetch groups. | string | Maximum length: 2047 | (objectCategory=group) | 
| ldap-poll-interval | Interval in minutes within to fetch groups from LDAP server. | integer | Minimum value: 1 Maximum value: 2880 | 180 | 
| ldap-server | LDAP server to get group information. | string | Maximum length: 35 |  | 
| logon-timeout | Interval in minutes to keep logons after FSSO server down. | integer | Minimum value: 1 Maximum value: 2880 | 5 | 
| name | Name. | string | Maximum length: 35 |  | 
| password | Password of the first FSSO collector agent. | password | Not Specified |  | 
| password2 | Password of the second FSSO collector agent. | password | Not Specified |  | 
| password3 | Password of the third FSSO collector agent. | password | Not Specified |  | 
| password4 | Password of the fourth FSSO collector agent. | password | Not Specified |  | 
| password5 | Password of the fifth FSSO collector agent. | password | Not Specified |  | 
| port | Port of the first FSSO collector agent. | integer | Minimum value: 1 Maximum value: 65535 | 8000 | 
| port2 | Port of the second FSSO collector agent. | integer | Minimum value: 1 Maximum value: 65535 | 8000 | 
| port3 | Port of the third FSSO collector agent. | integer | Minimum value: 1 Maximum value: 65535 | 8000 | 
| port4 | Port of the fourth FSSO collector agent. | integer | Minimum value: 1 Maximum value: 65535 | 8000 | 
| port5 | Port of the fifth FSSO collector agent. | integer | Minimum value: 1 Maximum value: 65535 | 8000 | 
| server | Domain name or IP address of the first FSSO collector agent. | string | Maximum length: 63 |  | 
| server2 | Domain name or IP address of the second FSSO collector agent. | string | Maximum length: 63 |  | 
| server3 | Domain name or IP address of the third FSSO collector agent. | string | Maximum length: 63 |  | 
| server4 | Domain name or IP address of the fourth FSSO collector agent. | string | Maximum length: 63 |  | 
| server5 | Domain name or IP address of the fifth FSSO collector agent. | string | Maximum length: 63 |  | 
| sni | Server Name Indication. | string | Maximum length: 255 |  | 
| source-ip | Source IP for communications to FSSO agent. | ipv4-address | Not Specified | 0.0.0.0 | 
| source-ip6 | IPv6 source for communications to FSSO agent. | ipv6-address | Not Specified | :: | 
| ssl | Enable/disable use of SSL. | option | - | disable | 
|  |  |  |  |  | 
| ssl-server-host-ip-check | Enable/disable server host/IP verification. | option | - | disable | 
|  |  |  |  |  | 
| ssl-trusted-cert | Trusted server certificate or CA certificate. | string | Maximum length: 79 |  | 
| type | Server type. | option | - | default | 
|  |  |  |  |  | 
| user-info-server | LDAP server to get user information. | string | Maximum length: 35 |  | 
| Option | Description | 
|---|---|
| auto | Set outgoing interface automatically. | 
| sdwan | Set outgoing interface by SD-WAN or policy routing rules. | 
| specify | Set outgoing interface manually. | 
| Option | Description | 
|---|---|
| enable | Enable automatic fetching of groups from LDAP server. | 
| disable | Disable automatic fetching of groups from LDAP server. | 
| Option | Description | 
|---|---|
| enable | Enable use of SSL. | 
| disable | Disable use of SSL. | 
| Option | Description | 
|---|---|
| enable | Enable server host/IP verification. | 
| disable | Disable server host/IP verification. | 
| Option | Description | 
|---|---|
| default | All other unspecified types of servers. | 
| fortinac | FortiNAC server. |
