---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-3-cli-reference-268798926-config-system-pppoe-interface-7ca2dd07
title: "document-fortigate-7-6-3-cli-reference-268798926-config-system-pppoe-interface-7ca2dd07"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-3-cli-reference-268798926-config-system-pppoe-interface-7ca2dd07.md
source_anchor: ""
source_lines: [1, 78]
sha256: ae725df791651b696b72d547052d8c3b8797acd86d6664b454903d99b4a16276
---

# document-fortigate-7-6-3-cli-reference-268798926-config-system-pppoe-interface-7ca2dd07

config system pppoe-interface
config system pppoe-interface
Configure the PPPoE interfaces.
config system pppoe-interface
    Description: Configure the PPPoE interfaces.
    edit <name>
        set ac-name {string}
        set auth-type [auto|pap|...]
        set device {string}
        set dial-on-demand [enable|disable]
        set disc-retry-timeout {integer}
        set idle-timeout {integer}
        set ipunnumbered {ipv4-address}
        set ipv6 [enable|disable]
        set lcp-echo-interval {integer}
        set lcp-max-echo-fails {integer}
        set padt-retry-timeout {integer}
        set password {password}
        set pppoe-egress-cos [cos0|cos1|...]
        set pppoe-unnumbered-negotiate [enable|disable]
        set service-name {string}
        set username {string}
    next
end
                                            config system pppoe-interface
| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| ac-name | PPPoE AC name. | string | Maximum length: 63 |  | 
| auth-type | PPP authentication type to use. | option | - | auto | 
|  |  |  |  |  | 
| device | Name for the physical interface. | string | Maximum length: 15 |  | 
| dial-on-demand | Enable/disable dial on demand to dial the PPPoE interface when packets are routed to the PPPoE interface. | option | - | disable | 
|  |  |  |  |  | 
| disc-retry-timeout | PPPoE discovery init timeout value in (0-4294967295 sec). | integer | Minimum value: 0 Maximum value: 4294967295 | 1 | 
| idle-timeout | PPPoE auto disconnect after idle timeout (0-4294967295 sec). | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| ipunnumbered | PPPoE unnumbered IP. | ipv4-address | Not Specified | 0.0.0.0 | 
| ipv6 | Enable/disable IPv6 Control Protocol (IPv6CP). | option | - | disable | 
|  |  |  |  |  | 
| lcp-echo-interval | Time in seconds between PPPoE Link Control Protocol (LCP) echo requests. | integer | Minimum value: 0 Maximum value: 32767 | 5 | 
| lcp-max-echo-fails | Maximum missed LCP echo messages before disconnect. | integer | Minimum value: 0 Maximum value: 32767 | 3 | 
| name | Name of the PPPoE interface. | string | Maximum length: 15 |  | 
| padt-retry-timeout | PPPoE terminate timeout value in (0-4294967295 sec). | integer | Minimum value: 0 Maximum value: 4294967295 | 1 | 
| password | Enter the password. | password | Not Specified |  | 
| pppoe-egress-cos | CoS in VLAN tag for outgoing PPPoE/PPP packets. | option | - | cos0 | 
|  |  |  |  |  | 
| pppoe-unnumbered-negotiate | Enable/disable PPPoE unnumbered negotiation. | option | - | enable | 
|  |  |  |  |  | 
| service-name | PPPoE service name. | string | Maximum length: 63 |  | 
| username | User name. | string | Maximum length: 64 |  | 
| Option | Description | 
|---|---|
| auto | Automatically choose the authentication method. | 
| pap | PAP authentication. | 
| chap | CHAP authentication. | 
| mschapv1 | MS-CHAPv1 authentication. | 
| mschapv2 | MS-CHAPv2 authentication. | 
| Option | Description | 
|---|---|
| enable | Enable dial on demand. | 
| disable | Disable dial on demand. | 
| Option | Description | 
|---|---|
| enable | Enable IPv6CP. | 
| disable | Disable IPv6CP. | 
| Option | Description | 
|---|---|
| cos0 | CoS 0. | 
| cos1 | CoS 1. | 
| cos2 | CoS 2. | 
| cos3 | CoS 3. | 
| cos4 | CoS 4. | 
| cos5 | CoS 5. | 
| cos6 | CoS 6. | 
| cos7 | CoS 7. | 
| Option | Description | 
|---|---|
| enable | Enable PPPoE unnumbered negotiation. | 
| disable | Disable PPPoE unnumbered negotiation. |
