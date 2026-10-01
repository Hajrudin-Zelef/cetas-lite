---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd-5
title: "enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd"
domain: cisco
role: reference
task: reference
actors: ["China"]
dates: ["1970-01-01"]
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd.md
source_anchor: ""
source_lines: [151, 179]
sha256: 64230008554cdfc25faaa5272f3e8087f05192b4f9e4169d8e42bd96d62fd507
---

# enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd

| 29 | Termination-Action | integer | What action the NAS should take when the specified service is completed:   NOTE:  This attribute is valid only for 802.1X and MAC address authentication users. When the authentication point is deployed on a VLANIF interface, MAC address authenticated users do not support the authorization of Termination-Action=1. When the RADIUS server delivers only this attribute, the value of attribute 27 Session-Timeout is set to 3600s (for 802.1X authentication users) or 1800s (for MAC address authentication users) by default. | 
| 30 | Called-Station-Id | string | Number of the NAS:   NOTE:  If this attribute carries Chinese characters, it cannot be delivered using the radius-attribute set attribute-name attribute-value command. | 
| 31 | Calling-Station-Id | string | Identification number of the client. Generally, it is the MAC address of the client.  NOTE:  If this attribute carries Chinese characters, it cannot be delivered using the radius-attribute set attribute-name attribute-value command. | 
| 32 | NAS-Identifier | string | String identifying the NAS device originating the Access-Request. By default, the attribute value is the host name of the NAS device. You can change the attribute value to the VLAN ID of the user or the MAC address of the AP using the radius-server nas-identifier-format { hostname \| vlan-id \| ap-info } command.  NOTE:  If this attribute carries Chinese characters, it cannot be delivered using the radius-attribute set attribute-name attribute-value command. | 
| 40 | Acct-Status-Type | integer | Accounting-Request type:  | 
| 41 | Acct-Delay-Time | integer | Number of seconds the client has been trying to send the accounting packet (excluding the network transmission time). | 
| 42 | Acct-Input-Octets | integer | Number of bytes in upstream traffic, corresponding to the lower 32 bits in the data structure for storing the upstream traffic. Contents of this attribute and the RADIUS attribute 52 (Acct-Input-Gigawords) compose the upstream traffic. The traffic unit must be the same as that of the RADIUS server and can be Byte, KByte, MByte, and GByte. To set the traffic unit for each RADIUS server, run the radius-server traffic-unit command. By default, the unit is Byte. | 
| 43 | Acct-Output-Octets | integer | Number of bytes in downstream traffic, corresponding to the lower 32 bits in the data structure for storing the downstream traffic. Contents of this attribute and the RADIUS attribute 53 (Acct-Output-Gigawords) compose the downstream traffic. The traffic unit must be the same as that of the RADIUS server and can be Byte, KByte, MByte, and GByte. To set the traffic unit for each RADIUS server, run the radius-server traffic-unit command. By default, the unit is Byte. | 
| 44 | Acct-Session-Id | string | Accounting session ID. The Accounting-Start, Interim-Accounting, and Accounting-Stop packets of the same accounting session must have the same session ID. The format of this attribute is: Host name (7 bits) + Slot ID (2 bits) + Subcard number (1 bit) + Port number (2 bits) + Outer VLAN ID (4 bits) + Inner VLAN ID (5 bits) + Central Processing Unit (CPU) Tick (6 bits) + User ID prefix (2 bits) + User ID (5 bits). | 
| 45 | Acct-Authentic | integer | User authentication mode:  | 
| 46 | Acct-Session-Time | integer | How long (in seconds) the user has received service.  NOTE:  If the administrator modifies the system time after the user goes online, the online time calculated by the device may be incorrect. | 
| 47 | Acct-Input-Packets | integer | Number of incoming packets. | 
| 48 | Acct-Output-Packets | integer | Number of outgoing packets. | 
| 49 | Acct-Terminate-Cause | integer | Cause of a terminated session:  | 
| 50 | Acct-Multi-Session-Id | string | Accounting ID, which is used in multi-link session scenarios. | 
| 52 | Acct-Input-Gigawords | integer | Number of times the number of bytes in upstream traffic is greater than 4 GB (2^32), corresponding to the higher 32 bits in the data structure for storing the upstream traffic. Contents of this attribute and the RADIUS attribute 42 (Acct-Input-Octets) compose the upstream traffic. The traffic unit must be the same as that of the RADIUS server and can be Byte, KByte, MByte, and GByte. To set the traffic unit for each RADIUS server, run the radius-server traffic-unit command. By default, the unit is Byte. | 
| 53 | Acct-Output-Gigawords | integer | Number of times the number of bytes in downstream traffic is greater than 4 GB (2^32), corresponding to the higher 32 bits in the data structure for storing the downstream traffic. Contents of this attribute and the RADIUS attribute 43 (Acct-Output-Octets) compose the downstream traffic. The traffic unit must be the same as that of the RADIUS server and can be Byte, KByte, MByte, and GByte. To set the traffic unit for each RADIUS server, run the radius-server traffic-unit command. By default, the unit is Byte. | 
| 55 | Event-Timestamp | integer | Time when an Accounting-Request packet is generated, represented by is the number of seconds elapsed since 00:00:00 of January 1, 1970. | 
| 56 | Egress-VLANID | integer | VLAN with a specified ID that an interface is added to in tagged or untagged mode. The VLAN ID is converted into a 6-digit hexadecimal number. If the number contains fewer than 6 digits, the number is prefixed with zeros. Then, the 6-digit hexadecimal number is prefixed with 0x31 in tagged mode or with 0x32 in untagged mode, which generates an 8-digit hexadecimal number. Finally, the 8-digit hexadecimal number is converted into a decimal number as the value configured on the RADIUS server. For example, the value used to authorize VLAN 48 to which an interface is added in untagged mode is 838860848.  NOTE:  The Egress-VLANID can be authorized only after the authentication mode multi-share command is configured in an authentication profile. | 
| 58 | Egress-VLAN-Name | string | VLAN with a specified name or description that an interface is added to in tagged or untagged mode. The VLAN name is prefixed with 1 in tagged mode or 2 in untagged mode. For example, the value used to authorize VLAN 48 to which an interface is added in untagged mode is 2vlan48. | 
| 60 | CHAP-Challenge | string | Challenge field in CHAP authentication. This field is generated by the NAS for Message Digest algorithm 5 (MD5) calculation. | 
| 61 | NAS-Port-Type | integer | NAS port type.  | 
| 64 | Tunnel-Type | integer | Protocol type of the tunnel. The value is fixed as 13, indicating VLAN. | 
| 65 | Tunnel-Medium-Type | integer | Medium type used on the tunnel. The value is fixed as 6, indicating Ethernet. | 
| 66 | Tunnel-Client-Endpoint | string | Tunnel client address. | 
| 67 | Tunnel-Server-Endpoint | string | Tunnel server address. | 
| 75 | Password-Retry | integer | Number of password retries. | 
| 79 | EAP-Message | string | Encapsulates Extended Access Protocol (EAP) packets so that RADIUS supports EAP authentication. When an EAP packet is longer than 253 bytes, the packet is encapsulated into multiple attributes. A RADIUS packet can carry multiple EAP-Message attributes. | 
| 80 | Message-Authenticator | string | Authenticates and verifies authentication packets to prevent spoofing packets. This attribute is used only when RADIUS supports EAP authentication. | 
