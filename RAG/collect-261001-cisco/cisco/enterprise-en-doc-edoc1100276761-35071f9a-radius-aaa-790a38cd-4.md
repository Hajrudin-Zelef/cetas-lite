---
id: collect-261001-cisco/cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd-4
title: "enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd"
domain: cisco
role: reference
task: reference
actors: ["China", "Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd.md
source_anchor: ""
source_lines: [113, 150]
sha256: feb39830df0d9e9ac3181ab1ff9e1b502233b30598fb074cf81a4f521109aec2
---

# enterprise-en-doc-edoc1100276761-35071f9a-radius-aaa-790a38cd

Each service provided by the NAS to a user constitutes a session, with the beginning of the session defined as the point where service is first provided and the end of the session defined as the point where service is ended.
The match methods are as follows:
any method
The device performs a match check between an attribute and user information on the device. The priority for identifying the RADIUS attributes used by the users is as follows: Acct-Session-ID (44) > Calling-Station-Id (31) > Framed-IP-Address (8). The device searches for the attributes in the request packet based on the priority, and performs a match check between the first found attribute and user information on the device. If the attribute is successfully matched, the device responds with an ACK packet; otherwise, the device responds with a NAK packet.
all method
The device performs a match check between all attributes and user information on the device. The device identifies the following RADIUS attributes used by the users: Acct-Session-ID (44), Calling-Station-Id (31), Framed-IP-Address (8), and User-Name (1). The device performs a match check between all the preceding attributes in the Request packet and user information on the device. If all the preceding attributes are successfully matched, the device responds with an ACK packet; otherwise, the device responds with a NAK packet.
When the CoA-Request or DM-Request packet from the RADIUS server fails to match user information on the device, the device describes the failure cause using the error code in the CoA-NAK or DM-NAK packet. For the error code description, see Table 1-8 and Table 1-9.
| Name | Value | Description | 
|---|---|---|
| RD_DM_ERRCODE_MISSING_ATTRIBUTE | 402 | The request packet lacks key attributes, so that the integrity check of the RADIUS attributes fails. | 
| RD_DM_ERRCODE_INVALID_REQUEST | 404 | Parsing the attributes in the request packet fails. | 
| RD_DM_ERRCODE_INVALID_ATTRIBUTE_VALUE | 407 | The request packet contains attributes that are not supported by the device or do not exist, so that the attribute check fails. Contents of the authorization check include VLAN, ACL, CAR, number of the ACL used for redirection, and whether Huawei RADIUS extended attributes RD_hw_URL_Flag and RD_hw_Portal_URL can be authorized to the interface-based authenticated user. Errors that may occur are as follows:  | 
| RD_DM_ERRCODE_SESSION_CONTEXT_NOT_FOUND | 503 | The session request fails. The cause includes:  | 
| RD_DM_ERRCODE_RESOURCES_UNAVAILABLE | 506 | This error code is used for other authorization failures. | 
RFC2865, RFC2866, and RFC3576 define standard RADIUS attributes that are supported by all mainstream vendors. For details, see Table 1-10.
| Attribute No. | Attribute Name | Attribute Type | Description | 
|---|---|---|---|
| 1 | User-Name | string | User name for authentication. The user name format can be user name@domain name or user name.  NOTE:  If this attribute carries Chinese characters, it cannot be delivered using the radius-attribute set attribute-name attribute-value command. | 
| 2 | User-Password | string | User password for authentication, which is only valid for the Password Authentication Protocol (PAP). | 
| 3 | CHAP-Password | string | User password for authentication, which is only valid for the Challenge Handshake Authentication Protocol (CHAP). | 
| 4 | NAS-IP-Address | ipaddr | Internet Protocol (IP) address of the NAS carried in authentication request packets. By default, the attribute value is the source IP address of the authentication request packets sent by the NAS. You can change the attribute value to the specified IP address on the NAS or the IP address of the AP using the radius-attribute nas-ip { ip-address \| ap-info } command. | 
| 5 | NAS-Port | integer | Physical port number of the network access server that is authenticating the user, which is in either of the following formats:  | 
| 6 | Service-Type | integer | Service type of the user to be authenticated:  | 
| 7 | Framed-Protocol | integer | Encapsulation protocol of Frame services:  | 
| 8 | Framed-IP-Address | ipaddr | User IP address. | 
| 9 | Framed-IP-Netmask | ipaddr | User IP address mask. This field must be used with the Framed-IP-Address field. | 
| 11 | Filter-Id | string | UCL group name, user group name, or IPv4 Access Control List (ACL) ID.  NOTE:   | 
| 12 | Framed-MTU | integer | MTU value of the data link between the NAS and server. For example, in 802.1X Extensible Authentication Protocol (EAP) authentication, the NAS specifies the maximum length of the EAP packet in this attribute. An EAP packet larger than the link MTU may be lost.  NOTE:  This attribute is carried only in EAP relay mode. | 
| 14 | Login-IP-Host | ipaddr | Management user IP address:  | 
| 15 | Login-Service | integer | Service to use to connect the user to the login host:   NOTE:  An attribute can contain multiple service types. | 
| 18 | Reply-Message | string | This attribute determines whether a user is authenticated:  | 
| 19 | Callback-Number | string | Information sent from the authentication server and to be displayed to a user, such as a mobile number. | 
| 22 | Framed-Route | string | Routing information provided by the RADIUS server to users, in format Destination/Mask NextHop Metric, for example, 192.168.1.0/24 192.168.1.1 1. If the NextHop value is 0.0.0.0, the user IP address is used as the next hop address. The device can obtain only one Metric value. If the attribute delivered by the RADIUS server contains multiple Metric values, the device obtains only the first one.  NOTE:  Currently, the Framed-Route attribute cannot provide routing information for PPPoE users on the device. | 
| 24 | State | string | If the RADIUS server sends a RADIUS Access-Challenge packet carrying this attribute to a device, the subsequent RADIUS Access-Request packets sent from the device must carry this attribute with the same value. This attribute is sent by the server to the client in an Access-Accept packet that also includes a Termination-Action attribute with the value of RADIUS-Request. If the NAS performs the termination action by sending a new Access-Request packet upon the termination of the current session, it must include the State attribute unchanged in that Access-Request packet. | 
| 25 | Class | string | If the RADIUS server sends a RADIUS Access-Accept packet carrying the Class attribute to the NAS, the subsequent RADIUS Accounting-Request packets sent from the NAS must carry the Class attribute with the same value. | 
| 26 | Vendor-Specific | string | Vendor-specific attribute. For details, see Table 1-11. A packet can carry one or more private attributes. Each private attribute contains one or more sub-attributes. | 
| 27 | Session-Timeout | integer | In the Access-Accept packet, this attribute indicates the maximum number of seconds a user should be allowed to remain connected. In the Access-Challenge packet, this attribute indicates the duration for which EAP authentication users are reauthenticated. When the value of this attribute is 0:   NOTE:  This attribute is only valid for 802.1X, MAC address, Portal, and PPPoE authentication users. When the RADIUS server delivers only this attribute, the value of attribute 29 Termination-Action is set to 0 (users are forced offline) by default. | 
| 28 | Idle-Timeout | integer | Maximum number of consecutive seconds of idle connection the user is allowed before termination of the session or prompt.  NOTE:   | 
