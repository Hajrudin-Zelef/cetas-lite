---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd-8
title: "enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd"
domain: cisco
role: reference
task: reference
actors: ["China"]
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd.md
source_anchor: ""
source_lines: [255, 283]
sha256: 1efeee31f7cf95c1a916d58513a17ca02b0feba13258daf02644b159bd84ac65
---

# enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd

| 26-170 | HW-Acct-ipv6-Input-Gigawords | integer | This attribute specifies the number of times that more than 4 GB upstream packets are carried in an IPv6 flow. This attribute is usually used with the HW-Acct-ipv6-Input-Octets attribute. | 
| 26-171 | HW-Acct-ipv6-Output-Gigawords | integer | This attribute specifies the number of times that more than 4 GB downstream packets are carried in an IPv6 flow. This attribute is usually used with the HW-Acct-ipv6-Output-Octets attribute. | 
| 26-173 | HW-Redirect-ACL | string | Redirect IPv4 ACL. Redirection is performed for only the users matching the ACL rules. The ACL number or ACL name can be delivered. The ACL name must start with a character.  NOTE:  The value range of acl-number is from 3000 to 3999 for wired users and from 3000 to 3031 for wireless users. After the authentication mode multi-share command is configured in the authentication profile, redirect ACL authorization is not supported. | 
| 26-174 | HW-Tariff-Ipv6-Input-Octets | string | Number of upstream bytes in the IPv6 traffic at the specified tariff level sent to the accounting server. This field is included in the accounting packets. The unit is Byte, KByte, MByte, or GByte. | 
| 26-175 | HW-Tariff-Ipv6-Output-Octets | string | Number of downstream bytes in the IPv6 traffic at the specified tariff level sent to the accounting server. This field is included in the accounting packets. The unit is Byte, KByte, MByte, or GByte. | 
| 26-176 | HW-Tariff-Ipv6-Input-Gigawords | string | How many times is the number of upstream bytes in the IPv6 traffic at the specified tariff level larger than 4 GB. This field and the HW-Tariff-Ipv6-Input-Octets field specify the number of upstream bytes in the IPv6 traffic at the specified tariff level. | 
| 26-177 | HW-Tariff-Ipv6-Output-Gigawords | string | How many times is the number of downstream bytes in the IPv6 traffic at the specified tariff level larger than 4 GB. This field and the HW-Tariff-Ipv6-Input-Octets field specify the number of downstream bytes in the IPv6 traffic at the specified tariff level. | 
| 26-178 | HW-IPv6-Redirect-ACL | string | Redirection IPv6 ACL. Redirection is performed for only the users matching the ACL rules. The ACL number or ACL name can be delivered. The ACL name must start with a character.  NOTE:   | 
| 26-201 | HW-User-Extend-Info | string | Extended user information. This attribute is contained in authentication and accounting request packets. A packet can contain multiple HW-User-Extend-Info attributes. The following describes extended user information:  This attribute applies only to MAC address authentication and Portal authentication. | 
| 26-202 | HW-MUD-URL | string | Identity information of iConnect terminals. The information is carried in an authentication request packet. The RADIUS server determines whether a terminal is an iConnect terminal based on the HW-MUD-URL attribute. If the terminal is an iConnect terminal, the RADIUS server queries the corresponding authorization policy based on the user account and encapsulates the authorization policy into an authentication response packet. For an iConnect terminal, you are advised to configure a redirection-related RADIUS attribute (such as HW-Redirect-ACL or HW-Portal-URL). In this way, the iConnect terminal will be redirected to a URL to download an EAP-TLS certificate after being authenticated successfully. After the certificate is downloaded, EAP-TLS authentication is triggered for the terminal.  NOTE:  The HW-MUD-URL attribute can be used only in wireless scenarios. The HW-MUD-URL attribute is supported only on the X series cards. The RADIUS server must be iMaster NCE-Campus. If a client sends an EAPoL-Start packet to trigger authentication, iConnect-URL is not carried during RADIUS authentication of an iConnect terminal. | 
| 26-203 | HW-VIP-Level-ID | integer | User priority. The value is 0 or 1. A larger value indicates a higher priority. The default value is 0. | 
| 26-204 | HW-SAC-Profile | string | SAC profile name.  NOTE:  This attribute is supported only by the S12700E. | 
| 26-237 | HW-Web-Authen-Info | string | Information sent from the portal server via the device (which transparently transmits the information) to the RADIUS server. For example, a user selects the authentication-free option and time information for next login, based on which the RADIUS server saves the MAC address of the user for a period of time. Upon the next login of the user, the login page is not displayed. Instead, MAC address authentication is preferentially used. This attribute can be used for transparent transmission in complex modes such as EAP. | 
| 26-238 | HW-Ext-Specific | string | User extended attributes:   NOTE:  During RADIUS CoA dynamic authorization, when the value of the user-command field is 1, 2, or 3, other authorization attributes are not supported. The user-dscp-in and user-dscp-out attributes cannot be authorized to wireless users in direct forwarding mode. This attribute applies only to NAC users. Pay attention to the following points if the value of the user-command field in the RADIUS attribute HW-Ext-Specific(26-238) carried in a CoA packet sent by the RADIUS server is 2 or 3:  | 
| 26-239 | HW-User-Access-Info | string | User context profile information.  NOTE:  The attribute can be authorized only after the access-context profile enable command is run. | 
| 26-240 | HW-Access-Device-Info | string | The authentication and accounting request packets carry the IP addresses, MAC addresses, and port numbers of access switches in policy association. The format is ip=A.B.C.D;mac=XXXX-XXXX-XXXX;slot=XX;subslot=XXX;port=XXX;vlanid=XXXX; interfaceName=port. | 
| 26-244 | HW-Reachable-Detect | string | Server reachability detection information. Authentication packets carrying this attribute are server detection packets. | 
| 26-247 | HW-Tariff-Input-Octets | string | Number of upstream bytes at the specified tariff level sent to the accounting server. This field is included in the accounting packets. The unit can be byte, kilobyte, megabyte, or gigabyte. The format is Tariff level:Number of upstream bytes. An accounting packet can contain the traffic of at most 8 tariff levels. | 
| 26-248 | HW-Tariff-Output-Octets | string | Number of downstream bytes at the specified tariff level sent to the accounting server. This field is included in the accounting packets. The unit can be byte, kilobyte, megabyte, or gigabyte. The format is Tariff level:Number of downstream bytes. An accounting packet can contain the traffic of at most 8 tariff levels. | 
| 26-249 | HW-Tariff-Input-Gigawords | string | Number of times larger the number of upstream bytes at the specified tariff level is than 4G. This field and the HW-Tariff-Input-Octets field specify the number of upstream bytes at the specified tariff level. | 
| 26-250 | HW-Tariff-Output-Gigawords | string | Number of times larger the number of downstream bytes at the specified tariff level is than 4G. This field and the HW-Tariff-Output-Octets field specify the number of downstream bytes at the specified tariff level. | 
| 26-251 | HW-IPv6-Filter-ID | string | Number of a user IPv6 ACL. The value range is 2000 to 3999 (for wired users) or 2000 to 3031 (for wireless users).  NOTE:    | 
| 26-253 | HW-Framed-IPv6-Address | ipaddr | IPv6 address to be configured for the user. | 
| 26-254 | HW-Version | string | Software version of the device.  NOTE:  If this attribute carries Chinese characters, it cannot be delivered using the radius-attribute set attribute-name attribute-value command. | 
| 26-255 | HW-Product-ID | string | NAS product name. | 
| Attribute No. | Attribute Name | Attribute Type | Description | 
|---|---|---|---|
| MICROSOFT-16 | MS-MPPE-Send-Key | string | This attribute indicates the MPPE sending key. | 
| MICROSOFT-17 | MS-MPPE-Recv-Key | string | This attribute indicates the MPPE receiving key. | 
