---
id: collect-261001-cisco/cisco/c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b-5
title: "c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b"
domain: cisco
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["agent", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b.md
source_anchor: ""
source_lines: [233, 285]
sha256: f5884157ab8b575a532472cc4fe1e4f064a08e099fc55d5c2f112e25a0e592ca
---

# c-en-us-td-docs-security-asa-asa917-configuration-general-asa-917-general-config-ce78b47b

                                 VLAN-only—SNMP uses logical statistics for ifInOctets and ifOutOctets.
The examples in the following table show the differences in SNMP traffic statistics. Example 1 shows the difference in physical and logical output statistics for the show interface command and the show traffic command. Example 2 shows output statistics for a VLAN-only interface for the show interface command and the show traffic command. The example shows that the statistics are close to the output that appears for the show traffic command.
| Table 6. SNMP Traffic Statistics for Physical and VLAN                                     		Interfaces |  | 
|---|---|
| Example 1 | Example 2 | 
|---|---|
|  ciscoasa# show interface GigabitEthernet3/2 interface GigabitEthernet3/2 description fullt-mgmt nameif mgmt security-level 10 ip address 10.7.14.201 255.255.255.0 management-only  ciscoasa# show traffic (Condensed output)  Physical Statistics GigabitEthernet3/2: received (in 121.760 secs) 36 packets       3428 bytes 0 pkts/sec      28 bytes/sec  Logical Statistics mgmt: received (in 117.780 secs) 36 packets       2780 bytes 0 pkts/sec      23 bytes/sec  The following examples show the SNMP output statistics for the management interface and the physical interface. The ifInOctets value is close to the physical statistics output that appears in the show traffic command output but not to the logical statistics output. ifIndex of the mgmt interface:  IF_MIB::ifDescr.6 = Adaptive Security Appliance ‘mgmt’ interface  ifInOctets that corresponds to the physical interface statistics:  IF-MIB::ifInOctets.6 = Counter32:3246   |  ciscoasa# show interface GigabitEthernet0/0.100 interface GigabitEthernet0/0.100 vlan 100 nameif inside security-level 100 ip address 10.7.1.101 255.255.255.0 standby 10.7.1.102  ciscoasa# show traffic inside received (in 9921.450 secs) 1977 packets       126528 bytes 0 pkts/sec      12 bytes/sec transmitted (in 9921.450 secs) 1978 packets       126556 bytes 0 pkts/sec      12 bytes/sec  ifIndex of VLAN inside:  IF-MIB::ifDescr.9 = Adaptive Security Appliance ‘inside’ interface IF-MIB::ifInOctets.9 = Counter32: 126318   | 
SNMP Version 3 Overview
SNMP Version 3 provides security enhancements that are not available in SNMP Version 1 or Version 2c. SNMP Versions 1 and 2c transmit data between the SNMP server and SNMP agent in clear text. SNMP Version 3 adds authentication and privacy options to secure protocol operations. In addition, this version controls access to the SNMP agent and MIB objects through the User-based Security Model (USM) and View-based Access Control Model (VACM). The ASA also supports the creation of SNMP groups and users, as well as hosts, which is required to enable transport authentication and encryption for secure SNMP communications.
Security Models
For configuration purposes, the authentication and privacy options are grouped together into security models. Security models apply to users and groups, which are divided into the following three types:
- 
                                    NoAuthPriv—No Authentication and No Privacy, which means that no security is applied to messages.
- 
                                    AuthNoPriv—Authentication but No Privacy, which means that messages are authenticated.
- 
                                    AuthPriv—Authentication and Privacy, which means that messages are authenticated and encrypted.
SNMP Groups
An SNMP group is an access control policy to which users can be added. Each SNMP group is configured with a security model, and is associated with an SNMP view. A user within an SNMP group must match the security model of the SNMP group. These parameters specify what type of authentication and privacy a user within an SNMP group uses. Each SNMP group name and security model pair must be unique.
SNMP Users
SNMP users have a specified username, a group to which the user belongs, authentication password, encryption password, and authentication and encryption algorithms to use. The authentication algorithm options are SHA-1, SHA-224, SHA-256 HMAC, and SHA-384. The encryption algorithm options are 3DES and AES (which is available in 128, 192, and 256 versions). When you create a user, you must associate it with an SNMP group. The user then inherits the security model of the group.
| Note | When configuring an SNMP v3 user account, ensure that the length of authentication algorithm is equal to or greater than the length of encryption algorithm. | 
SNMP Hosts
An SNMP host is an IP address to which SNMP notifications and traps are sent. To configure SNMP Version 3 hosts, along with the target IP address, you must configure a username, because traps are only sent to a configured user. SNMP target IP addresses and target parameter names must be unique on the ASA. Each SNMP host can have only one username associated with it. To receive SNMP traps, after you have added the snmp-server host command, make sure that you configure the user credentials on the NMS to match the credentials for the ASA.
| Note | You can add up to 4000 hosts. However, only 128 of this number can be for traps. | 
Implementation Differences Between the ASA and Cisco IOS Software
The SNMP Version 3 implementation in the ASA differs from the SNMP Version 3 implementation in the Cisco IOS software in the following ways:
- 
                                    		  
                                    The local-engine and remote-engine IDs are not configurable. The local engine ID is generated when the ASA starts or when a context is created.
- 
                                    		  
                                    No support exists for view-based access control, which results in unrestricted MIB browsing.
- 
                                    		  
                                    Support is restricted to the following MIBs: USM, VACM, FRAMEWORK, and TARGET.
- 
                                    		  
                                    You must create users and groups with the correct security model.
- 
                                    		  
                                    You must remove users, groups, and hosts in the correct sequence.
- 
                                    		  
                                    Use of the snmp-server host command creates an ASA rule to allow incoming SNMP traffic.
SNMP Syslog Messaging
SNMP generates detailed syslog messages that are numbered 212nnn. Syslog messages indicate the status of SNMP requests, SNMP traps, SNMP channels, and SNMP responses from the ASA or ASASM to a specified host on a specified interface.
For detailed information about syslog messages, see the syslog messages guide.
| Note | SNMP polling fails if SNMP syslog messages exceed a high rate (approximately 4000 per second). | 
Application Services and Third-Party Tools
For information about SNMP support, see the following URL:
http://www.cisco.com/en/US/tech/tk648/tk362/tk605/tsd_technology_support_sub-protocol_home.html
For information Refer to SNMP Chapter in Cisco ASA Series General Operations ASDM Configuration Guide, see the following URL:
