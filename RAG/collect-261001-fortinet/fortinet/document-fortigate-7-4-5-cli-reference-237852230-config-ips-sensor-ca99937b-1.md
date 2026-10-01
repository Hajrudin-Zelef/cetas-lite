---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-5-cli-reference-237852230-config-ips-sensor-ca99937b-1
title: "document-fortigate-7-4-5-cli-reference-237852230-config-ips-sensor-ca99937b"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-5-cli-reference-237852230-config-ips-sensor-ca99937b.md
source_anchor: ""
source_lines: [1, 164]
sha256: bbfb55fe688ad0c6c47cce77a3f112a6a581ea4c7efe46bc4fdd6cc7d90b720f
---

# document-fortigate-7-4-5-cli-reference-237852230-config-ips-sensor-ca99937b

config ips sensor
config ips sensor
Configure IPS sensor.
config ips sensor
    Description: Configure IPS sensor.
    edit <name>
        set block-malicious-url [disable|enable]
        set comment {var-string}
        config entries
            Description: IPS sensor filter.
            edit <id>
                set action [pass|block|...]
                set application {user}
                set cve <cve-entry1>, <cve-entry2>, ...
                set default-action [all|pass|...]
                set default-status [all|enable|...]
                config exempt-ip
                    Description: Traffic from selected source or destination IP addresses is exempt from this signature.
                    edit <id>
                        set dst-ip {ipv4-classnet}
                        set src-ip {ipv4-classnet}
                    next
                end
                set last-modified {user}
                set location {user}
                set log [disable|enable]
                set log-attack-context [disable|enable]
                set log-packet [disable|enable]
                set os {user}
                set protocol {user}
                set quarantine [none|attacker]
                set quarantine-expiry {user}
                set quarantine-log [disable|enable]
                set rate-count {integer}
                set rate-duration {integer}
                set rate-mode [periodical|continuous]
                set rate-track [none|src-ip|...]
                set rule <id1>, <id2>, ...
                set severity {user}
                set status [disable|enable|...]
                set vuln-type <id1>, <id2>, ...
            next
        end
        set extended-log [enable|disable]
        set replacemsg-group {string}
        set scan-botnet-connections [disable|block|...]
    next
end
                                            config ips sensor
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| block-malicious-url | Enable/disable malicious URL blocking. | option | - | disable | 
|  |  |  |  |  | 
| comment | Comment. | var-string | Maximum length: 255 |  | 
| extended-log | Enable/disable extended logging. | option | - | disable | 
|  |  |  |  |  | 
| name | Sensor name. | string | Maximum length: 35 |  | 
| replacemsg-group | Replacement message group. | string | Maximum length: 35 |  | 
| scan-botnet-connections | Block or monitor connections to Botnet servers, or disable Botnet scanning. | option | - | disable | 
|  |  |  |  |  | 
| Option | Description | 
|---|---|
| disable | Disable malicious URL blocking. | 
| enable | Enable malicious URL blocking. | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
| Option | Description | 
|---|---|
| disable | Do not scan connections to botnet servers. | 
| block | Block connections to botnet servers. | 
| monitor | Log connections to botnet servers. | 
config entries
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| action | Action taken with traffic in which signatures are detected. | option | - | default | 
|  |  |  |  |  | 
| application | Operating systems to be protected. Use all for every application and other for unlisted application. | user | Not Specified | all | 
| cve <cve-entry> | List of CVE IDs of the signatures to add to the sensor. CVE IDs or CVE wildcards. | string | Maximum length: 19 |  | 
| default-action | Signature default action filter. | option | - | all | 
|  |  |  |  |  | 
| default-status | Signature default status filter. | option | - | all | 
|  |  |  |  |  | 
| id | Rule ID in IPS database. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| last-modified | Filter by signature last modified date. Formats: before <date>, after <date>, between <start-date> <end-date>. | user | Not Specified |  | 
| location | Protect client or server traffic. | user | Not Specified | all | 
| log | Enable/disable logging of signatures included in filter. | option | - | enable | 
|  |  |  |  |  | 
| log-attack-context | Enable/disable logging of attack context: URL buffer, header buffer, body buffer, packet buffer. | option | - | disable | 
|  |  |  |  |  | 
| log-packet | Enable/disable packet logging. Enable to save the packet that triggers the filter. You can download the packets in pcap format for diagnostic use. | option | - | disable | 
|  |  |  |  |  | 
| os | Operating systems to be protected. Use all for every operating system and other for unlisted operating systems. | user | Not Specified | all | 
| protocol | Protocols to be examined. Use all for every protocol and other for unlisted protocols. | user | Not Specified | all | 
| quarantine | Quarantine method. | option | - | none | 
|  |  |  |  |  | 
| quarantine-expiry | Duration of quarantine. Requires quarantine set to attacker. | user | Not Specified | 5m | 
| quarantine-log | Enable/disable quarantine logging. | option | - | enable | 
|  |  |  |  |  | 
| rate-count | Count of the rate. | integer | Minimum value: 0 Maximum value: 65535 | 0 | 
| rate-duration | Duration (sec) of the rate. | integer | Minimum value: 1 Maximum value: 65535 | 60 | 
| rate-mode | Rate limit mode. | option | - | continuous | 
|  |  |  |  |  | 
| rate-track | Track the packet protocol field. | option | - | none | 
|  |  |  |  |  | 
| rule <id> | Identifies the predefined or custom IPS signatures to add to the sensor. Rule IPS. | integer | Minimum value: 0 Maximum value: 4294967295 |  | 
| severity | Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity. | user | Not Specified | all | 
| status | Status of the signatures included in filter. Only those filters with a status to enable are used. | option | - | default | 
|  |  |  |  |  | 
| vuln-type <id> | List of signature vulnerability types to filter by. Vulnerability type ID. | integer | Minimum value: 0 Maximum value: 4294967295 |  | 
| Option | Description | 
|---|---|
| pass | Pass or allow matching traffic. | 
| block | Block or drop matching traffic. | 
| reset | Reset sessions for matching traffic. | 
| default | Pass or drop matching traffic, depending on the default action of the signature. | 
| Option | Description | 
|---|---|
| all | Selects signatures with any default action. | 
| pass | Selects signatures with default action 'pass'. | 
| block | Selects signatures with default action 'block'. | 
| Option | Description | 
|---|---|
| all | Selects signatures with any default status. | 
| enable | Selects signatures enabled by default. | 
| disable | Selects signatures disabled by default. | 
| Option | Description | 
|---|---|
| disable | Disable logging of selected rules. | 
| enable | Enable logging of selected rules. | 
| Option | Description | 
|---|---|
| disable | Disable logging of detailed attack context. | 
| enable | Enable logging of detailed attack context. | 
| Option | Description | 
|---|---|
| disable | Disable packet logging of selected rules. | 
| enable | Enable packet logging of selected rules. | 
| Option | Description | 
|---|---|
| none | Quarantine is disabled. | 
| attacker | Block all traffic sent from attacker's IP address. The attacker's IP address is also added to the banned user list. The target's address is not affected. | 
| Option | Description | 
|---|---|
| disable | Disable quarantine logging. | 
| enable | Enable quarantine logging. | 
| Option | Description | 
|---|---|
| periodical | Allow configured number of packets every rate-duration. | 
| continuous | Block packets once the rate is reached. | 
| Option | Description | 
|---|---|
| none | none | 
| src-ip | Source IP. | 
| dest-ip | Destination IP. | 
| dhcp-client-mac | DHCP client. | 
| dns-domain | DNS domain. | 
| Option | Description | 
|---|---|
| disable | Disable status of selected rules. | 
| enable | Enable status of selected rules. | 
| default | Default. | 
config exempt-ip
