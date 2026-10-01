---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd-7
title: "enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd"
domain: cisco
role: reference
task: reference
actors: []
dates: ["1970-01-01"]
keywords: ["agent", "voice"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd.md
source_anchor: ""
source_lines: [222, 254]
sha256: 660a8ae54534dd03cc42ea7bdabf9aa8984197a02f1c561b74922183113d71d2
---

# enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd

| 26-33 | HW-VoiceVlan | integer | Voice VLAN authorization flag. The value 1 indicates that the authorized VLAN is the voice VLAN. This attribute is used with VLAN authorization attributes.  NOTE:  After the authentication mode multi-share command is configured in the authentication profile, voice VLAN authorization is not supported. When the Tunnel-Private-Group-ID attribute is used to authorize a VLAN pool, the HW-VoiceVlan attribute cannot be used to authorize a voice VLAN. Otherwise, the two attributes do not take effect. | 
| 26-35 | HW-ProxyRdsPkt | integer | Whether a RADIUS server is a proxy server:  | 
| 26-59 | HW-NAS-Startup-Time-Stamp | integer | NAS start time, represented by the number of seconds elapsed since 00:00:00 of January 1, 1970. | 
| 26-60 | HW-IP-Host-Address | string | User IP address and MAC address carried in authentication and accounting packets, in the format A.B.C.D hh:hh:hh:hh:hh:hh. The IP address and MAC address are separated by a space. If the user's IP address is detected to be invalid during authentication, the IP address is set to 255.255.255.255. | 
| 26-61 | HW-Up-Priority | integer | 802.1p priority of upstream packets. | 
| 26-62 | HW-Down-Priority | integer | 802.1p priority of downstream packets. | 
| 26-75 | HW-Primary-WINS | ipaddr | Primary WINS server address delivered by the RADIUS server after a user is successfully authenticated. | 
| 26-76 | HW-Second-WINS | ipaddr | Secondary WINS server address delivered by the RADIUS server after a user is successfully authenticated. | 
| 26-77 | HW-Input-Peak-Burst-Size | integer | Upstream peak rate, in bit/s. The minimum value is 10000. | 
| 26-78 | HW-Output-Peak-Burst-Size | integer | Downstream peak rate, in bit/s. The minimum value is 10000. | 
| 26-82 | HW-Data-Filter | string | Used by the RADIUS server to deliver IPv4 or IPv6 ACL rules to users. ACL rules can be delivered in two modes: delivering ACL rules through DACL groups and delivering ACL rules directly. There are old and new attribute formats for ACL rules. Compared with the old attribute format, the new attribute format shortens the length of an ACL rule. Using a DACL group to deliver ACL rules saves more ACL resources than delivering ACL rules directly. The users in the same DACL group share the ACL resources in the group, whereas each user occupies ACL resources when ACL rules are directly delivered.  NOTE:   Directly delivering ACL rules in new attribute format (fields in square brackets are optional) The attribute format is as follows: $number permit/deny [ protocol ] [ direction ip-address [ port ] ] [ direction ip-address [ port ] ]. The fields are described as follows:  The following examples are the attribute values entered on the server: $1 permit dst 10.0.239.192/26 $2 permit udp src any 8080 $3 permit icmp echo dst 10.1.1.1/24 $5 deny Directly delivering ACL rules in old attribute format The attribute format is acl number key1 key-value1... keyN key-valueN permit/deny. The fields are described as follows: All keywords are case-insensitive. All keywords are separated from keyword values using spaces. The location of keywords is not fixed. The keywords permit and deny can be placed after number or the whole command line.  The following examples are the attribute values entered on the server: acl 10005 deny acl 10006 tcp-dstport 5080 permit acl 10007 dest-ip 10.11.11.2 dest-ipmask 32 permit acl 10008 dest-ip 10.11.11.3 dest-ipmask 32 udp-dstport 5070 permit acl 10009 dest-ip 11.11.11.2 dest-ipmask 32 udp-dstport 5070 udp-dstport-end 5080 deny Delivering ACL rules using DACL groups The format of ACL rules in a DACL group can be the new or old format. The new format is recommended. When the device is connected to a Cisco ISE server, an ACL rule starts with the number sign (#). The following examples are the attribute values entered on the server: $1 dacl-group-name example $2 permit dst 10.0.239.192/26 $3 permit udp src any 8080 $4 deny | 
| 26-94 | HW-VPN-Instance | string | VPN instance name delivered by the RADIUS server after a user is successfully authenticated. It specifies the VPN to which the user belongs. When the RADIUS server delivers this attribute, it also needs to deliver the standard RADIUS attribute 8 (Framed-IP-Address). | 
| 26-135 | HW-Client-Primary-DNS | ipaddr | Primary DNS address delivered by the RADIUS server after a user is successfully authenticated. | 
| 26-136 | HW-Client-Secondary-DNS | ipaddr | Secondary DNS address delivered by the RADIUS server after a user is successfully authenticated. | 
| 26-138 | HW-Domain-Name | string | Name of the domain used for user authentication. This attribute can be the domain name contained in a user name or the name of a forcible domain. | 
| 26-141 | HW-AP-Information | string | AP's MAC address used for STA authentication, in H-H-H format. H is a 4-digit hexadecimal number. | 
| 26-142 | HW-User-Information | string | User security check information delivered by the RADIUS server to an Extensible Authentication Protocol over LAN (EAPoL) user to notify the user of items that require security checks. | 
| 26-146 | HW-User-Policy | string | Service scheme name. A service scheme contains user authorization information and policies. | 
| 26-153 | HW-Access-Type | integer | User access type carried in the Access-Request or Accounting-Request packets that the device sends to the RADIUS server:  | 
| 26-155 | HW-URL-Flag | integer | Whether a Uniform Resource Locator (URL) needs to be forcibly pushed when it is used together with another attribute, for example, HW-Portal-URL:  | 
| 26-156 | HW-Portal-URL | string | Forcibly pushed URL. The maximum length is 247 bytes. If information delivered by the RADIUS server matches the configured URL template, the URL configured in the template is used. Otherwise, the character string delivered by the RADIUS server is used. | 
| 26-157 | HW-Terminal-Type | string | Terminal type of a user. To configure this attribute, run the device-type device-name command. | 
| 26-158 | HW-DHCP-Option | hex | DHCP Option, encapsulated in Type-Length-Value (TLV) format. A packet may contain multiple HW-DHCP-Option attributes to carry Option information. Only Option 82 can be delivered. | 
| 26-159 | HW-HTTP-UA | string | User-Agent information in Hypertext Transfer Protocol (HTTP) packets. | 
| 26-160 | HW-UCL-Group | integer | Index of a UCL group.  NOTE:  This attribute cannot be authorized together with ACLs. Otherwise, only the authorized ACLs take effect. | 
| 26-161 | HW-Forwarding-VLAN | string | Delivers the Internet Service Provider (ISP) VLAN for user packet forwarding.  NOTE:  This attribute is supported only by the following: X series cards. | 
| 26-162 | HW-Forwarding-Interface | string | Outbound interface for forwarding user packets.  NOTE:  This attribute is supported only by the following: X series cards. | 
| 26-163 | HW-LLDP | string | LLDP information. A packet can contain multiple HW-LLDP-Info attributes to carry different options. The meanings of different options are as follows:  | 
| 26-164 | HW-CDP | string | CDP information. A packet can contain multiple HW-CDP attributes to carry different options. The meanings of different options are as follows:  | 
| 26-166 | HW-Acct-ipv6-Input-Octets | integer | Number of upstream bytes in an IPv6 flow. The unit can be byte, kilobyte, megabyte, or gigabyte. | 
| 26-167 | HW-Acct-ipv6-Output-Octets | integer | Number of downstream bytes in an IPv6 flow. The unit can be byte, kilobyte, megabyte, or gigabyte. | 
| 26-168 | HW-Acct-ipv6-Input-Packets | integer | Number of upstream packets in an IPv6 flow. | 
| 26-169 | HW-Acct-ipv6-Output-Packets | integer | Number of downstream packets in an IPv6 flow. | 
