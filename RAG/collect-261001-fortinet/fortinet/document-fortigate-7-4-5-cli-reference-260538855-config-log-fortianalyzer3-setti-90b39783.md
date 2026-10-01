---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-5-cli-reference-260538855-config-log-fortianalyzer3-setti-90b39783
title: "document-fortigate-7-4-5-cli-reference-260538855-config-log-fortianalyzer3-setti-90b39783"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-5-cli-reference-260538855-config-log-fortianalyzer3-setti-90b39783.md
source_anchor: ""
source_lines: [1, 136]
sha256: 370c5535cb19e4ccea29b9d19b3dd88e1909662609512281a4901d777d1c345a
---

# document-fortigate-7-4-5-cli-reference-260538855-config-log-fortianalyzer3-setti-90b39783

config log fortianalyzer3 setting
config log fortianalyzer3 setting
Global FortiAnalyzer settings.
config log fortianalyzer3 setting
    Description: Global FortiAnalyzer settings.
    set access-config [enable|disable]
    set alt-server {string}
    set certificate {string}
    set certificate-verification [enable|disable]
    set conn-timeout {integer}
    set enc-algorithm [high-medium|high|...]
    set fallback-to-primary [enable|disable]
    set hmac-algorithm {option}
    set interface {string}
    set interface-select-method [auto|sdwan|...]
    set ips-archive [enable|disable]
    set max-log-rate {integer}
    set monitor-failure-retry-period {integer}
    set monitor-keepalive-period {integer}
    set preshared-key {string}
    set priority [default|low]
    set reliable [enable|disable]
    set serial <name1>, <name2>, ...
    set server {string}
    set server-cert-ca {string}
    set source-ip {string}
    set ssl-min-proto-version [default|SSLv3|...]
    set status [enable|disable]
    set upload-day {user}
    set upload-interval [daily|weekly|...]
    set upload-option [store-and-upload|realtime|...]
    set upload-time {user}
end
                                            config log fortianalyzer3 setting
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| access-config | Enable/disable FortiAnalyzer access to configuration and data. | option | - | enable | 
|  |  |  |  |  | 
| alt-server | Alternate FortiAnalyzer. | string | Maximum length: 127 |  | 
| certificate | Certificate used to communicate with FortiAnalyzer. | string | Maximum length: 35 |  | 
| certificate-verification | Enable/disable identity verification of FortiAnalyzer by use of certificate. | option | - | enable | 
|  |  |  |  |  | 
| conn-timeout | FortiAnalyzer connection time-out in seconds (for status and log buffer). | integer | Minimum value: 1 Maximum value: 3600 | 10 | 
| enc-algorithm | Configure the level of SSL protection for secure communication with FortiAnalyzer. | option | - | high | 
|  |  |  |  |  | 
| fallback-to-primary | Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. | option | - | enable | 
|  |  |  |  |  | 
| hmac-algorithm | OFTP login hash algorithm. | option | - | sha256 | 
|  |  |  |  |  | 
| interface | Specify outgoing interface to reach server. | string | Maximum length: 15 |  | 
| interface-select-method | Specify how to select outgoing interface to reach server. | option | - | auto | 
|  |  |  |  |  | 
| ips-archive | Enable/disable IPS packet archive logging. | option | - | enable | 
|  |  |  |  |  | 
| max-log-rate | FortiAnalyzer maximum log rate in MBps (0 = unlimited). | integer | Minimum value: 0 Maximum value: 100000 | 0 | 
| monitor-failure-retry-period | Time between FortiAnalyzer connection retries in seconds (for status and log buffer). | integer | Minimum value: 1 Maximum value: 86400 | 5 | 
| monitor-keepalive-period | Time between OFTP keepalives in seconds (for status and log buffer). | integer | Minimum value: 1 Maximum value: 120 | 5 | 
| preshared-key | Preshared-key used for auto-authorization on FortiAnalyzer. | string | Maximum length: 63 |  | 
| priority | Set log transmission priority. | option | - | default | 
|  |  |  |  |  | 
| reliable | Enable/disable reliable logging to FortiAnalyzer. | option | - | disable | 
|  |  |  |  |  | 
| serial <name> | Serial numbers of the FortiAnalyzer. Serial Number. | string | Maximum length: 79 |  | 
| server | The remote FortiAnalyzer. | string | Maximum length: 127 |  | 
| server-cert-ca | Mandatory CA on FortiGate in certificate chain of server. | string | Maximum length: 79 |  | 
| source-ip | Source IPv4 or IPv6 address used to communicate with FortiAnalyzer. | string | Maximum length: 63 |  | 
| ssl-min-proto-version | Minimum supported protocol version for SSL/TLS connections. | option | - | default | 
|  |  |  |  |  | 
| status | Enable/disable logging to FortiAnalyzer. | option | - | disable | 
|  |  |  |  |  | 
| upload-day | Day of week (month) to upload logs. | user | Not Specified |  | 
| upload-interval | Frequency to upload log files to FortiAnalyzer. | option | - | daily | 
|  |  |  |  |  | 
| upload-option | Enable/disable logging to hard disk and then uploading to FortiAnalyzer. | option | - | 5-minute | 
|  |  |  |  |  | 
| upload-time | Time to upload logs (hh:mm). | user | Not Specified |  | 
| Option | Description | 
|---|---|
| enable | Enable FortiAnalyzer access to configuration and data. | 
| disable | Disable FortiAnalyzer access to configuration and data. | 
| Option | Description | 
|---|---|
| enable | Enable identity verification of FortiAnalyzer by use of certificate. | 
| disable | Disable identity verification of FortiAnalyzer by use of certificate. | 
| Option | Description | 
|---|---|
| high-medium | Encrypt logs using high and medium encryption algorithms. | 
| high | Encrypt logs using high encryption algorithms. | 
| low | Encrypt logs using all encryption algorithms. | 
| Option | Description | 
|---|---|
| enable | Enable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. | 
| disable | Disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. | 
| Option | Description | 
|---|---|
| sha256 | Use SHA256 as HMAC algorithm. | 
| Option | Description | 
|---|---|
| auto | Set outgoing interface automatically. | 
| sdwan | Set outgoing interface by SD-WAN or policy routing rules. | 
| specify | Set outgoing interface manually. | 
| Option | Description | 
|---|---|
| enable | Enable IPS packet archive logging. | 
| disable | Disable IPS packet archive logging. | 
| Option | Description | 
|---|---|
| default | Set FortiAnalyzer log transmission priority to default. | 
| low | Set FortiAnalyzer log transmission priority to low. | 
| Option | Description | 
|---|---|
| enable | Enable reliable logging to FortiAnalyzer. | 
| disable | Disable reliable logging to FortiAnalyzer. | 
| Option | Description | 
|---|---|
| default | Follow system global setting. | 
| SSLv3 | SSLv3. | 
| TLSv1 | TLSv1. | 
| TLSv1-1 | TLSv1.1. | 
| TLSv1-2 | TLSv1.2. | 
| TLSv1-3 | TLSv1.3. | 
| Option | Description | 
|---|---|
| enable | Enable logging to FortiAnalyzer. | 
| disable | Disable logging to FortiAnalyzer. | 
| Option | Description | 
|---|---|
| daily | Upload log files to FortiAnalyzer once a day. | 
| weekly | Upload log files to FortiAnalyzer once a week. | 
| monthly | Upload log files to FortiAnalyzer once a month. | 
| Option | Description | 
|---|---|
| store-and-upload | Log to hard disk and then upload to FortiAnalyzer. | 
| realtime | Log directly to FortiAnalyzer in real time. | 
| 1-minute | Log directly to FortiAnalyzer at least every 1 minute. | 
| 5-minute | Log directly to FortiAnalyzer at least every 5 minutes. |
