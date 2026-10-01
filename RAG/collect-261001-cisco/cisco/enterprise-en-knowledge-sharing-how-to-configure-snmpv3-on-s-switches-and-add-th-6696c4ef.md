---
id: collect-261001-cisco/cisco/enterprise-en-knowledge-sharing-how-to-configure-snmpv3-on-s-switches-and-add-th-6696c4ef
title: "Configure an ACL to permit only the NMS 10.1.1.1 to access the switch."
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-knowledge-sharing-how-to-configure-snmpv3-on-s-switches-and-add-th-6696c4ef.md
source_anchor: ""
source_lines: [1, 58]
sha256: 5715d28a1b23372b3d0cc9081bb8923488bbb56dd8fe6659a58fd5c6682c1e92
---

# Configure an ACL to permit only the NMS 10.1.1.1 to access the switch.

As you know, Simple Network Management Protocol (SNMP) is a standard network management protocol widely used on TCP/IP networks. SNMP manages network elements by using a central computer (a network management station, NMS) running network management software.
SNMP is available in three versions. SNMPv1 is the initial version of SNMP. It provides authentication based on community names. SNMPv1 has low security, and can return only a few error codes. SNMPv2c issued by IETF is the second release of SNMP. SNMPv2c has enhancements to standard error codes, data types (Counter 64 and Counter 32), and operations including GetBulk and Inform. SNMPv2c does not improve the security, so IETF issued SNMPv3 that provides User Security Module (USM)-based authentication and encryption and View-based Access Control Model (VACM)-based access control.
SNMPv3 is applicable to networks of various scales, especially
networks that have strict security requirements and can be managed
only by authorized network administrators. For example, SNMPv3 can
be used if data between the NMS and managed device needs to be transmitted
over a public network.That's why I choose this version as an example for my guide.
As shown in figure, the NMS server manages all devices on the network. The network is large and is prone to attack. Therefore, devices on the network use SNMPv3 to communicate with the NMS server. A new switch is added to the network. The network administrator wants to use the existing network resources to manage the new switch and quickly locate and rectify network faults.
The configuration roadmap is as follows:
Configure SNMPv3 on the switch so that the NMS running SNMPv3 can manage the switch.
Configure access control so that only the NMS with the specified IP address can perform read/write operations on the specified MIB objects of the switch.
Configure a user group and user based on which the switch permits access of the NMS.
Configure a trap host and enable the switch to automatically send traps to the NMS.
Add the switch to the NMS. The user group and user configured on the switch must be the same as those used by the NMS; otherwise, the NMS cannot manage the switch.
<HUAWEI> system-view
[HUAWEI] sysname Switch
[Switch] snmp-agent sys-info version v3 //By default, SNMPv3 is supported. If SNMPv3 is not disabled, skip this command.
# Configure an ACL to permit only the NMS 10.1.1.1 to access the switch.
[Switch] acl 2001
[Switch-acl-basic-2001] rule permit source 10.1.1.1 0
[Switch-acl-basic-2001] rule deny
[Switch-acl-basic-2001] quit
# Configure the MIB view to specify the MIB objects that can be accessed by the NMS.
[Switch] snmp-agent mib-view included isoview iso   //Configure the MIB view isoview to access the iso subtree.
# Configure the user group group001, set the security level to privacy, and configure access control to restrict the access of NMS to the switch.
[Switch] snmp-agent group v3 group001 privacy write-view isoview acl 2001 //Configure a user group and apply an ACL to make the access control function take effect.
# Configure the user user001.
[Switch] snmp-agent usm-user v3 user001 group group001 //Configure a user name and add the user name to the specified user group.
# Set the authentication password to Authe@1234.
[Switch] snmp-agent usm-user v3 user001 authentication-mode sha //Set an authentication mode to improve security of the managed switch running SNMPv3.
Please configure the authentication password (8-64)                             
Enter Password:                  //Enter the authentication password. It is Authe@1234 in this example.
Confirm Password:                //Confirm the password. It is Authe@1234 in this example.
# Set the encryption password to Priva@1234.
[Switch] snmp-agent usm-user v3 user001 privacy-mode aes128 //Set an encryption mode to improve security of the managed switch running SNMPv3.
Please configure the privacy password (8-64)                                    
Enter Password:                  //Enter the encryption password. It is Priva@1234 in this example.
Confirm Password:                //Confirm the password. It is Priva@1234 in this example.
[Switch] snmp-agent trap enable
Warning: All switches of SNMP trap/notification will be open. Continue? [Y/N]:y //Enable all trap functions on the switch. By default, only some trap functions are enabled. You can run the display snmp-agent trap all command to check trap status.
[Switch] snmp-agent target-host trap address udp-domain 10.1.1.1 params securityname user001 v3 privacy //Configure a trap host. By default, traps are sent by UDP port 162. The security name must be the same as the user name; otherwise, the NMS cannot manage the device.
# Log in to eSight and choose . Set SNMP parameters based on Table 1 and click OK. A switch is added to and can be managed by eSight. The switch will proactively send trap messages to eSight.
| Table 1 SNMP parameters |  | 
|---|---|
| Parameter | Setting | 
|---|---|
| Select discovery protocol | SNMP | 
| IP address | 10.1.1.2 | 
| SNMP | Edit SNMP parameters | 
| Version | V3 | 
| Security name | user001 | 
| Port | 161 | 
| Authentication protocol | HMAC_SHA | 
| Authentication password | Authe@1234 | 
| Privacy protocol | AES_128 | 
| Encryption password | Priva@1234 | 
The parameter settings on the NMS and switch must be the same; otherwise, the switch cannot be added to the NMS.
If authentication is required for remote logins to the switch, Telnet parameters need to be set so that the NMS can manage the switch. In this example, administrators can remotely log in to the switch using Telnet, password authentication is used, and the password is huawei2012.
