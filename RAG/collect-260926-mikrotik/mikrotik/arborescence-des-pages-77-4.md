---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-77-4
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "ascend"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-77.md
source_anchor: ""
source_lines: [181, 322]
sha256: feca31ba20cb4c98edeeaa4940ea3b6e9ac55128e772985e120d316bc328c61c
---

# Summary

## Stop and Interim-Update Accounting-Request packet

Additionally to the accounting start request, the following messages will contain the following attributes:

- **Acct-Session-Time** - connection uptime in seconds
- **Acct-Input-Octets** - bytes received from the client
- **Acct-Input-Gigawords** - 4G (2^32) bytes received from the client (bits 32..63, when bits 0..31 are delivered in Acct-Input-Octets)
- **Acct-Input-Packets** - nubmer of packets received from the client
- **Acct-Output-Octets** - bytes sent to the client
- **Acct-Output-Gigawords** - 4G (2^32) bytes sent to the client (bits 32..63, when bits 0..31 are delivered in Acct-Output-Octets)
- **Acct-Output-Packets** - number of packets sent to the client

RADIUS accounting messages in RouterOS represent output as traffic from client point of view, for example, for VPN user "test" output is user upload.

## **Stop Accounting-Request packet**

These packets will, additionally to the Interim Update packets, have:

- **Acct-Terminate-Cause** - session termination cause (see RFC 2866 ch. 5.10)

## **Change of Authorization**

RADIUS disconnect and Change of Authorization (according to RFC3576) are supported as well. These attributes may be changed by a CoA request from the RADIUS server:

- **Mikrotik-Group**
- **Mikrotik-Recv-Limit**
- **Mikrotik-Xmit-Limit**
- **Mikrotik-Rate-Limit**
- **Ascend-Data-Rate** (only if Mikrotik-Rate-Limit is not present)
- **Ascend-XMit-Rate** (only if Mikrotik-Rate-Limit is not present)
- **Mikrotik-Mark-Id**
- **Filter-Id**
- **Mikrotik-Advertise-Url**
- **Mikrotik-Advertise-Interval**
- **Session-Timeout**
- **Idle-Timeout**
- **Port-Limit**

Note that it is not possible to change IP address, pool or routes that way - for such changes a user must be disconnected first.

# MikroTik Specific RADIUS Attribute Numeric Values

| Name | VendorID | Value | RFC | 
|---|---|---|---|
| **MIKROTIK_RECV_LIMIT** | 14988 | 1 |  | 
| **MIKROTIK_XMIT_LIMIT** | 14988 | 2 |  | 
| **MIKROTIK_GROUP** | 14988 | 3 |  | 
| **MIKROTIK_WIRELESS_FORWARD** | 14988 | 4 |  | 
| **MIKROTIK_WIRELESS_SKIPDOT1X** | 14988 | 5 |  | 
| **MIKROTIK_WIRELESS_ENCALGO** | 14988 | 6 |  | 
| **MIKROTIK_WIRELESS_ENCKEY** | 14988 | 7 |  | 
| **MIKROTIK_RATE_LIMIT** | 14988 | 8 |  | 
| **MIKROTIK_REALM** | 14988 | 9 |  | 
| **MIKROTIK_HOST_IP** | 14988 | 10 |  | 
| **MIKROTIK_MARK_ID** | 14988 | 11 |  | 
| **MIKROTIK_ADVERTISE_URL** | 14988 | 12 |  | 
| **MIKROTIK_ADVERTISE_INTERVAL** | 14988 | 13 |  | 
| **MIKROTIK_RECV_LIMIT_GIGAWORDS** | 14988 | 14 |  | 
| **MIKROTIK_XMIT_LIMIT_GIGAWORDS** | 14988 | 15 |  | 
| **MIKROTIK_WIRELESS_PSK** | 14988 | 16 |  | 
| **MIKROTIK_TOTAL_LIMIT** | 14988 | 17 |  | 
| **MIKROTIK_TOTAL_LIMIT_GIGAWORDS** | 14988 | 18 |  | 
| **MIKROTIK_ADDRESS_LIST** | 14988 | 19 |  | 
| **MIKROTIK_WIRELESS_MPKEY** | 14988 | 20 |  | 
| **MIKROTIK_WIRELESS_COMMENT** | 14988 | 21 |  | 
| **MIKROTIK_DELEGATED_IPV6_POOL** | 14988 | 22 |  | 
| **MIKROTIK_DHCP_OPTION_SET** | 14988 | 23 |  | 
| **MIKROTIK_DHCP_OPTION_PARAM_STR1** | 14988 | 24 |  | 
| **MIKROTIK_DHCP_OPTION_PARAM_STR2** | 14988 | 25 |  | 
| **MIKROTIK_WIRELESS_VLANID** | 14988 | 26 |  | 
| **MIKROTIK_WIRELESS_VLANIDTYPE** | 14988 | 27 |  | 
| **MIKROTIK_WIRELESS_MINSIGNAL** | 14988 | 28 |  | 
| **MIKROTIK_WIRELESS_MAXSIGNAL** | 14988 | 29 |  | 

# All Supported Attribute Numeric Values

| Name | VendorID | Value | RFC | 
|---|---|---|---|
| **Acct-Authentic** |  | 45 | RFC 2866 | 
| **Acct-Delay-Time** |  | 41 | RFC 2866 | 
| **Acct-Input-Gigawords** |  | 52 | RFC 2869 | 
| **Acct-Input-Octets** |  | 42 | RFC 2866 | 
| **Acct-Input-Packets** |  | 47 | RFC 2866 | 
| **Acct-Interim-Interval** |  | 85 | RFC 2869 | 
| **Acct-Output-Gigawords** |  | 53 | RFC 2869 | 
| **Acct-Output-Octets** |  | 43 | RFC 2866 | 
| **Acct-Output-Packets** |  | 48 | RFC 2866 | 
| **Acct-Session-Id** |  | 44 | RFC 2866 | 
| **Acct-Session-Time** |  | 46 | RFC 2866 | 
| **Acct-Status-Type** |  | 40 | RFC 2866 | 
| **Acct-Terminate-Cause** |  | 49 | RFC 2866 | 
| **Ascend-Client-Gateway** | 529 | 132 |  | 
| **Ascend-Data-Rate** | 529 | 197 |  | 
| **Ascend-Xmit-Rate** | 529 | 255 |  | 
| **Called-Station-Id** |  | 30 | RFC 2865 | 
| **Calling-Station-Id** |  | 31 | RFC 2865 | 
| **CHAP-Challenge** |  | 60 | RFC 2866 | 
| **CHAP-Password** |  | 3 | RFC 2865 | 
| **Class** |  | 25 | RFC 2865 | 
| **Filter-Id** |  | 11 | RFC 2865 | 
| **Framed-Compression** |  | 53 | RFC 2865 | 
| **Framed-IP-Address** |  | 8 | RFC 2865 | 
| **Framed-IP-Netmask** |  | 9 | RFC 2865 | 
| **Framed-IPv6-Prefix** |  | 97 | RFC 3162 | 
| **Framed-Mtu** |  | 52 | RFC 2869 | 
| **Framed-Pool** |  | 88 | RFC 2869 | 
| **Framed-Protocol** |  | 7 | RFC 2865 | 
| **Framed-Route** |  | 22 | RFC 2865 | 
| **Framed-Routing** |  | 50 | RFC 2865 | 
| **Idle-Timeout** |  | 28 | RFC 2865 | 
| **MS-CHAP-Challenge** | 311 | 11 | RFC 2548 | 
| **MS-CHAP-Domain** | 311 | 10 | RFC 2548 | 
| **MS-CHAP-Response** | 311 | 1 | RFC 2548 | 
| **MS-CHAP2-Response** | 311 | 25 | RFC 2548 | 
| **MS-CHAP2-Success** | 311 | 26 | RFC 2548 | 
| **MS-MPPE-Encryption-Policy** | 311 | 7 | RFC 2548 | 
| **MS-MPPE-Encryption-Types** | 311 | 8 | RFC 2548 | 
| **MS-MPPE-Recv-Key** | 311 | 17 | RFC 2548 | 
| **MS-MPPE-Send-Key** | 311 | 16 | RFC 2548 | 
| **NAS-Identifier** |  | 32 | RFC 2865 | 
| **NAS-Port** |  | 5 | RFC 2865 | 
| **NAS-IP-Address** |  | 4 | RFC 2865 | 
| **NAS-Port-Id** |  | 87 | RFC 2869 | 
| **NAS-Port-Type** |  | 61 | RFC 2865 | 
| **Port-Limit** |  | 62 | RFC 2865 | 
| **Redback-Agent-Remote-Id** | 2352 | 96 |  | 
| **Redback-Agent-Circuit-Id** | 2352 | 97 |  | 
| **Service-Type** |  | 6 | RFC 2865 | 
| **Session-Timeout** |  | 27 | RFC 2865 | 
| **User-Name** |  | 1 | RFC 2865 | 
| **User-Password** |  | 2 | RFC 2865 | 
| **WISPr-Bandwidth-Max-Down** | 14122 | 8 | wi-fi.org | 
| **WISPr-Bandwidth-Max-Up** | 14122 | 7 | wi-fi.org | 
| **WISPr-Bandwidth-Min-Down** | 14122 | 6 | wi-fi.org | 
| **WISPr-Bandwidth-Min-Up** | 14122 | 5 | wi-fi.org | 
| **WISPr-Location-Id** | 14122 | 1 | wi-fi.org | 
| **WISPr-Location-Name** | 14122 | 2 | wi-fi.org | 
| **WISPr-Logoff-URL** | 14122 | 3 | wi-fi.org | 
| **WISPr-Redirection-URL** | 14122 | 4 | wi-fi.org | 
| **WISPr-Session-Terminate-Time** | 14122 | 9 | wi-fi.org | 
| **WISPr-Session-Terminate-End-Of-Day** | 14122 | 10 | wi-fi.org | 
| **WISPr-Billing-Class-Of-Service** | 14122 | 11 | wi-fi.org |
