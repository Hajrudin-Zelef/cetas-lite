---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546-13
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546.md
source_anchor: ""
source_lines: [544, 634]
sha256: 961d68b9a8f0a864869943831f93f2530ea16ec05eb662c731ec555b4a7164a4
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-16-6-configuration-gu-46a80546

                                 Private VLAN—You can assign a client to a private VLAN.
- 
                                 				
                                 Network Edge Access Topology (NEAT)—MAB and NEAT are mutually exclusive. You cannot enable MAB when NEAT is enabled on an interface, and you should not enable NEAT when MAB is enabled on an interface.
Cisco IOS Release 12.2(55)SE and later supports filtering of verbose MAB system messages
Network Admission Control Layer 2 IEEE 802.1x Validation
The switch supports the Network Admission Control (NAC) Layer 2 IEEE 802.1x validation, which checks the antivirus condition or posture of endpoint systems or clients before granting the devices network access. With NAC Layer 2 IEEE 802.1x validation, you can do these tasks:
-  
                                 		  
                                 Download the Session-Timeout RADIUS attribute (Attribute[27]) and the Termination-Action RADIUS attribute (Attribute[29]) from the authentication server.
-  
                                 		  
                                 Set the number of seconds between re-authentication attempts as the value of the Session-Timeout RADIUS attribute (Attribute[27]) and get an access policy against the client from the RADIUS server.
-  
                                 		  
                                 Set the action to be taken when the switch tries to re-authenticate the client by using the Termination-Action RADIUS attribute (Attribute[29]). If the value is the DEFAULT or is not set, the session ends. If the value is RADIUS-Request, the re-authentication process starts.
- 
                                 		  
                                 Set the list of VLAN number or name or VLAN group name as the value of the Tunnel Group Private ID (Attribute[81]) and the preference for the VLAN number or name or VLAN group name as the value of the Tunnel Preference (Attribute[83]). If you do not configure the Tunnel Preference, the first Tunnel Group Private ID (Attribute[81]) attribute is picked up from the list.
- 
                                 		  
                                 View the NAC posture token, which shows the posture of the client, by using the show authentication privileged EXEC command.
-  
                                 		  
                                 Configure secondary private VLANs as guest VLANs.
Configuring NAC Layer 2 IEEE 802.1x validation is similar to configuring IEEE 802.1x port-based authentication except that you must configure a posture token on the RADIUS server.
Flexible Authentication Ordering
- 
                                    		  
                                    dot1X—IEEE 802.1X authentication is a Layer 2 authentication method.
- 
                                    		  
                                    mab—MAC-Authentication Bypass is a Layer 2 authentication method.
- 
                                    		  
                                    webauth—Web authentication is a Layer 3 authentication method.
Using this feature, you can control which ports use which authentication methods, and you can control the failover sequencing of methods on those ports. For example, MAC authentication bypass and 802.1x can be the primary or secondary authentication methods, and web authentication can be the fallback method if either or both of those authentication attempts fail.
- 
                                    		  
                                    multi-auth—Multiauthentication allows one authentication on a voice VLAN and multiple authentications on the data VLAN.
- 
                                    		  
                                    multi-domain—Multidomain authentication allows two authentications: one on the voice VLAN and one on the data VLAN.
Open1x Authentication
Open1x authentication allows a device access to a port before that device is authenticated. When open authentication is configured, a new host can pass traffic according to the access control list (ACL) defined on the port. After the host is authenticated, the policies configured on the RADIUS server are applied to that host.
You can configure open authentication with these scenarios:
- 
                                 		  
                                 Single-host mode with open authentication–Only one user is allowed network access before and after authentication.
- 
                                 		  
                                 MDA mode with open authentication–Only one user in the voice domain and one user in the data domain are allowed.
- 
                                 		  
                                 Multiple-hosts mode with open authentication–Any host can access the network.
- 
                                 		  
                                 Multiple-authentication mode with open authentication–Similar to MDA, except multiple hosts can be authenticated. 
 Note
 If open authentication is configured, it takes precedence over other authentication controls. This means that if you use the authentication open interface configuration command, the port will grant access to the host irrespective of the authentication port-control interface configuration command. 
Multidomain Authentication
The switch supports multidomain authentication (MDA), which allows both a data device and voice device, such as an IP phone (Cisco or non-Cisco), to authenticate on the same switch port. The port is divided into a data domain and a voice domain.
| Note | For all host modes, the line protocol stays up before authorization when port-based authentication is configured. | 
MDA does not enforce the order of device authentication. However, for best results, we recommend that a voice device is authenticated before a data device on an MDA-enabled port.
Follow these guidelines for configuring MDA:
-  
                                 		  
                                 You must configure a switch port for MDA.
-  
                                 		  
                                 You must configure the voice VLAN for the IP phone when the host mode is set to multidomain.
-  
                                 		  
                                 Voice VLAN assignment on an MDA-enabled port is supported Cisco IOS Release 12.2(40)SE and later.
-  
                                 		  
                                 To authorize a voice device, the AAA server must be configured to send a Cisco Attribute-Value (AV) pair attribute with a value of device-traffic-class=voice. Without this value, the switch treats the voice device as a data device.
-  
                                 		  
                                 The guest VLAN and restricted VLAN features only apply to the data devices on an MDA-enabled port. The switch treats a voice device that fails authorization as a data device.
-  
                                 		  
                                 If more than one device attempts authorization on either the voice or the data domain of a port, it is error disabled.
-  
                                 		  
                                 Until a device is authorized, the port drops its traffic. Non-Cisco IP phones or voice devices are allowed into both the data and voice VLANs. The data VLAN allows the voice device to contact a DHCP server to obtain an IP address and acquire the voice VLAN information. After the voice device starts sending on the voice VLAN, its access to the data VLAN is blocked.
-  
                                 		  
                                 A voice device MAC address that is binding on the data VLAN is not counted towards the port security MAC address limit.
-  
                                 		  
