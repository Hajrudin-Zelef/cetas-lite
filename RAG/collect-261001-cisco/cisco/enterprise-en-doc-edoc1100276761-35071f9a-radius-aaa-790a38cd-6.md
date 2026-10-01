---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd-6
title: "enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd"
domain: cisco
role: reference
task: reference
actors: ["China", "Huawei"]
dates: []
keywords: ["ascend", "voice"]
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd.md
source_anchor: ""
source_lines: [180, 221]
sha256: 12efeed2e4506b221a23dd59605751e6108d238353f43330b5876b6070d010f4
---

# enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd

| 81 | Tunnel-Private-Group-ID | string | Tunnel private group ID, which is used to deliver user VLANs.  NOTE:  If users are authorized with VLANs based on VLAN pools in the Tunnel-Private-Group-ID attribute, both the Tunnel-Type and Tunnel-Medium-Type attributes must be configured, and their values must be 13 and 6 respectively. The VLAN pool is configured as the service VLAN for wireless users. If the authorized VLAN is in the VLAN pool, the VLAN can be switched when an IP address fails to be obtained. If the authorized VLAN is not in the VLAN pool, the VLAN cannot be switched. Authorization can be performed based on the VLAN ID, VLAN description, VLAN name, and VLAN pool, which are listed in descending order of priority. Both wired and wireless users support authorization based on the VLAN pool. In V200R013C00 and later versions, wireless users support authorization based on the VLAN pool. When the Tunnel-Private-Group-ID attribute is used to authorize a VLAN pool, the HW-VoiceVlan attribute cannot be used to authorize a voice VLAN. Otherwise, the two attributes do not take effect. For details, see Licensing Requirements and Limitations for NAC Unified Mode. | 
| 82 | Tunnel-Assignment-Id | string | ID allocated to a tunnel. | 
| 83 | Tunnel-Preference | string | Tunnel priority. | 
| 85 | Acct-Interim-Interval | integer | Interim accounting interval, in seconds. It is recommended that the interval be at least 600 seconds. The value ranges from 60 to 3932100. | 
| 87 | NAS-Port-Id | string | Port of the NAS that is authenticating the user. The NAS-Port-Id attribute has the following formats:   NOTE:  If this attribute carries Chinese characters, it cannot be delivered using the radius-attribute set attribute-name attribute-value command. | 
| 88 | Framed-Pool | string | Address pool, which is only included in the Access-Accept packet. It is used as authorization information in Efficient VPN. | 
| 89 | Chargeable-User-Identity | string | Charging ID delivered by the server. To configure a device to support this attribute, run the radius-server support chargeable-user-identity [ not-reject ] command. | 
| 90 | Tunnel-Client-Auth-Id | string | Client tunnel ID used for authentication during tunnel setup. | 
| 91 | Tunnel-Server-Auth-Id | string | Server tunnel ID used for authentication during tunnel setup. | 
| 95 | NAS-IPv6-Address | ipaddr | IPv6 address carried in the authentication request packet sent by the NAS. Both the NAS-IPv6-Address and NAS-IP-Address fields can be included in a packet. | 
| 96 | Framed-Interface-Id | string | IPv6 interface identifier to be configured for the user. | 
| 97 | Framed-IPv6-Prefix | ipaddr | IPv6 prefix to be configured for the user. | 
| 98 | Login-IPv6-Host | ipaddr | User IPv6 address carried in the RADIUS authentication request of an IPv6 administrator. | 
| 101 | Error-Cause | integer | Cause of the authorization failure:  | 
| 121 | Input-Peak-Rate | integer | Uplink peak rate. | 
| 122 | Input-Average-Rate | integer | Average uplink rate. | 
| 123 | Input-Basic-Rate | integer | Uplink basic rate. | 
| 124 | Output-Peak-Rate | integer | Downlink peak rate. | 
| 125 | Output-Average-Rate | integer | Average downlink rate. | 
| 126 | Output-Basic-Rate | integer | Downlink basic rate. | 
| 135 | Ascend-Client-Primary-Dns | ipaddr | DNS server address 1. | 
| 136 | Ascend-Client-Secondary-Dns | ipaddr | DNS server address 2. | 
| 168 | Framed-IPv6-Address | ipaddr | IPv6 address of the user. | 
| 192 | Remote-IP-Host | ipaddr | Remote IP host. | 
| 195 | HW-SecurityStr | string | Security information of users in EAP relay authentication. | 
RADIUS is a fully extensible protocol. The No. 26 attribute (Vendor-Specific) defined in RFC2865 can be used to extend RADIUS for implementing functions not supported by standard RADIUS attributes. Table 1-11 describes Huawei proprietary RADIUS attributes.
Extended RADIUS attributes contain the vendor ID of the device. The vendor ID of Huawei is 2011.
| Attribute No. | Attribute Name | Attribute Type | Description | 
|---|---|---|---|
| 26-1 | HW-Input-Peak-Information-Rate | integer | Peak information rate (PIR) at which the user accesses the NAS, which is the maximum rate of traffic that can pass through an interface. The value is a 4-byte integer, in bit/s. The minimum value is 64. The HW-Input-Peak-Information-Rate must be higher than or equal to the HW-Input-Committed-Information-Rate. The default HW-Input-Peak-Information-Rate is equal to the HW-Input-Committed-Information-Rate. | 
| 26-2 | HW-Input-Committed-Information-Rate | integer | Committed information rate (CIR) at which the user accesses the NAS, which is the allowed average rate of traffic that can pass through an interface. The value is a 4-byte integer, in bit/s. The minimum value is 64.  NOTE:  This attribute must be specified when the rate of packets sent from the user to the NAS is limited. | 
| 26-3 | HW-Input-Committed-Burst-Size | integer | Committed burst size (CBS) at which the user accesses the NAS, which is the average volume of burst traffic that can pass through an interface. The value is a 4-byte integer, in bits. The minimum value is 80000. | 
| 26-4 | HW-Output-Peak-Information-Rate | integer | Peak information rate at which the NAS connects to the user. The value is a 4-byte integer, in bit/s. The minimum value is 64. The HW-Output-Peak-Information-Rate must be higher than or equal to the HW-Output-Committed-Information-Rate. The default HW-Output-Peak-Information-Rate is equal to the HW-Output-Committed-Information-Rate. | 
| 26-5 | HW-Output-Committed-Information-Rate | integer | Committed information rate at which the NAS connects to the user. The value is a 4-byte integer, in bit/s. The minimum value is 64.  NOTE:  This attribute must be specified when the rate of packets sent from the NAS to the user is limited. | 
| 26-6 | HW-Output-Committed-Burst-Size | integer | Committed burst size at which the NAS connects to the user. The value is a 4-byte integer, in bits. The minimum value is 80000. | 
| 26-15 | HW-Remanent-Volume | integer | Remaining traffic. The unit is KB. | 
| 26-17 | HW-Subscriber-QoS-Profile | string | Name of the QoS profile.  NOTE:    | 
| 26-18 | HW-UserName-Access-Limit | integer | Maximum number of users who are allowed to access the network using the same user name. The limit is indicated by a particular numeric value as follows:   NOTE:  This attribute can be carried only in Access-Accept packets. | 
| 26-26 | HW-Connect-ID | integer | Index of a user connection. | 
| 26-28 | HW-FTP-Directory | string | Initial directory of an FTP user. | 
| 26-29 | HW-Exec-Privilege | integer | Privilege level of an administrator (such as a Telnet user). The value ranges from 0 to 15. The priority that is greater than or equal to 16 is ineffective. | 
| 26-31 | HW-Qos-Data | string | Name of the QoS profile. The maximum length of the name is 31 bytes. The RADIUS server uses this field to deliver the QoS profile for traffic policing. The QoS profile must exist on the device and traffic policing has been configured using the car (QoS profile view) command.  NOTE:  If the server delivers both the uplink or downlink bandwidth limit (equivalent to the RADIUS attribute HW-Input-Committed-Information-Rate or HW-Output-Committed-Information-Rate) and the RADIUS attribute HW-Qos-Data for user authorization, only the uplink or downlink bandwidth limit take effect. | 
