---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-76-2
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ascend"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-76.md
source_anchor: ""
source_lines: [13, 73]
sha256: 00e39976e5bdc39452ce7fa96f07a9cffd1ea3d3c187becf487fed7a9aea8899
---

# Overview

| Attribute | Vendor ID | Type ID | Value type | Packet type | Description | 
|---|---|---|---|---|---|
| Framed-IP-Address | 0 (standard) | 8 | ip address | Access-Accept | RFC2865 section 5.8 | 
| Framed-IP-Netmask | 0 (standard) | 9 | ip address | Access-Accept | RFC2865 section 5.9 | 
| Session-Timeout | 0 (standard) | 27 | integer (maximum value: 21474720) | Access-Accept, Access-Challenge | RFC2865 section 5.27 | 
| Idle-Timeout | 0 (standard) | 28 | integer | Access-Accept, Access-Challenge | RFC2865 section 5.28 | 
| Tunnel-Type | 0 (standard) | 64 |  | Access-Accept | RFC2868 section 3.1 RFC3580 section 3.31 | 
| Tunnel-Medium-Type | 0 (standard) | 65 |  | Access-Accept | RFC2868 section 3.2 | 
| Tunnel-Private-Group-ID | 0 (standard) | 81 | string | Access-Accept | RFC2868 section 3.6 | 
| Framed-Pool | 0 (standard) | 88 | string | Access-Accept | RFC2869 section 5.18 | 
| Framed-IPv6-Prefix | 0 (standard) | 97 | ipv6 prefix | Access-Accept | RFC3162 section 2.3 | 
| Framed-IPv6-Pool | 0 (standard) | 100 | string | Access-Accept | RFC3162 section 2.6 | 
| Delegated-IPv6-Prefix | 0 (standard) | 123 | ipv6 prefix | Access-Accept | RFC4818 | 
| Framed-IPv6-Address | 0 (standard) | 168 | ip address | Access-Accept | RFC6911 section 3.1 | 
| Mikrotik-Recv-Limit | 14988 (Mikrotik) | 1 | integer | Access-Accept | Total receive limit in bytes for the client. | 
| Mikrotik-Xmit-Limit | 14988 (Mikrotik) | 2 | integer | Access-Accept | Total transmit limit in bytes for the client. | 
| Mikrotik-Group | 14988 (Mikrotik) | 3 | string | Access-Accept | User's group for local users. HotSpot profile for HotSpot users. PPP profile for PPP users. | 
| Mikrotik-Wireless-Forward | 14988 (Mikrotik) | 4 | integer | Access-Accept | Not forward the client's frames back to the wireless infrastructure if this attribute is set to "0" (wireless only). | 
| Mikrotik-Wireless-Skip-Dot1x | 14988 (Mikrotik) | 5 | integer | Access-Accept | Disable 802.1x authentication for the particular wireless client if set to a non-zero value (wireless only). | 
| Mikrotik-Wireless-Enc-Algo | 14988 (Mikrotik) | 6 |  | Access-Accept | WEP encryption algorithm( wireless only). | 
| Mikrotik-Wireless-Enc-Key | 14988 (Mikrotik) | 7 | string | Access-Accept | WEP encryption key for the client (wireless only). | 
| Mikrotik-Rate-Limit | 14988 (Mikrotik) | 8 | string | Access-Accept | Datarate limitation for clients. The format is: rx-rate[/tx-rate] [rx-burst-rate[/tx-burst-rate] [rx-burst-threshold[/tx-burst-threshold] [rx-burst-time[/tx-burst-time] [priority] [rx-rate-min[/tx-rate-min]]]] from the point of view of the router (so "rx" is client upload, and "tx" is client download). All rates should be numbers with optional 'k' (1,000s) or 'M' (1,000,000s). If the tx-rate is not specified, the rx-rate is as tx-rate too. The same goes for tx-burst-rate and tx-burst-threshold and tx-burst-time. If both rx-burst-threshold and tx-burst-threshold are not specified (but burst-rate is specified), rx-rate and tx-rate are used as burst thresholds. If both rx-burst-time and tx-burst-time are not specified, 1s is used as default. Priority takes values 1..8, where 1 implies the highest priority, but 8 - the lowest. If rx-rate-min and tx-rate-min are not specified rx-rate and tx-rate values are used. The rx-rate-min and tx-rate-min values can not exceed rx-rate and tx-rate values. | 
| Mikrotik-Realm | 14988 (Mikrotik) | 9 | string | Access-Request | If it is set in /radius menu, it is included in every RADIUS request as Mikrotik-Realm attribute. If it is not set, the same value is sent as in the MS-CHAP-Domain attribute (if MS-CHAP-Domain is missing, Realm is not included either). | 
| Mikrotik-Host-IP | 14988 (Mikrotik) | 10 | ip address | Access-Request | The IP address of HotSpot client before Universal Client translation (the original IP address of the client). | 
| Mikrotik-Mark-Id | 14988 (Mikrotik) | 11 | string | Access-Accept | Firewall mangle chain name (HotSpot only). The MikroTik RADIUS client upon receiving this attribute creates a dynamic firewall mangle rule with action=jump chain=hotspot and jump-target equal to the attribute value. Mangle chain name can have suffixes .in or .out, which will install rule only for incoming or outgoing traffic. Multiple Mark-id attributes can be provided, but only the last ones for incoming and outgoing are used. | 
| Mikrotik-Advertise-URL | 14988 (Mikrotik) | 12 | string | Access-Accept | URL of the page with advertisements that should be displayed to clients. If this attribute is specified, advertisements are enabled automatically, including transparent proxy, even if they were explicitly disabled in the corresponding user profile. Multiple attribute instances may be sent by the RADIUS server to specify additional URLs which are chosen in a round-robin fashion. | 
| Mikrotik-Advertise-Interval | 14988 (Mikrotik) | 13 | integer | Access-Accept | The time interval between two adjacent advertisements. Multiple attribute instances may be sent by the RADIUS server to specify additional intervals. All interval values are treated as a list and are taken one by one for each successful advertisement. If the end of the list is reached, the last value is continued to be used. | 
| Mikrotik-Recv-Limit-Gigawords | 14988 (Mikrotik) | 14 | integer | Access-Accept | 4G (2^32) bytes of total receive limit (bits 32..63, when bits 0..31 are delivered in Mikrotik-Recv-Limit). | 
| Mikrotik-Xmit-Limit-Gigawords | 14988 (Mikrotik) | 15 | integer | Access-Accept | 4G (2^32) bytes of total transmit limit (bits 32..63, when bits 0..31 are delivered in Mikrotik-Recv-Limit). | 
| Mikrotik-Wireless-PSK | 14988 (Mikrotik) | 16 | string | Access-Accept |  | 
| Mikrotik-Total-Limit | 14988 (Mikrotik) | 17 | integer | Access-Accept |  | 
| Mikrotik-Total-Limit-Gigawords | 14988 (Mikrotik) | 18 | integer | Access-Accept |  | 
| Mikrotik-Address-List | 14988 (Mikrotik) | 19 | string | Access-Accept |  | 
| Mikrotik-Wireless-MPKey | 14988 (Mikrotik) | 20 | string | Access-Accept |  | 
| Mikrotik-Wireless-Comment | 14988 (Mikrotik) | 21 | string | Access-Accept |  | 
| Mikrotik-Delegated-IPv6-Pool | 14988 (Mikrotik) | 22 | string | Access-Accept | IPv6 pool used for Prefix Delegation. | 
| Mikrotik-DHCP-Option-Set | 14988 (Mikrotik) | 23 | string | Access-Accept |  | 
| Mikrotik-DHCP-Option-Param-STR1 | 14988 (Mikrotik) | 24 | string | Access-Accept |  | 
| Mikrotik-DHCP-Option-Param-STR2 | 14988 (Mikrotik) | 25 | string | Access-Accept |  | 
| Mikrotik-Wireless-VLANID | 14988 (Mikrotik) | 26 | integer | Access-Accept | VLAN ID for the client (Wireless only). | 
| Mikrotik-Wireless-VLANIDtype | 14988 (Mikrotik) | 27 |  | Access-Accept | VLAN ID type for the client (Wireless only). | 
| Mikrotik-Wireless-Minsignal | 14988 (Mikrotik) | 28 | string | Access-Accept |  | 
| Mikrotik-Wireless-Maxsignal | 14988 (Mikrotik) | 29 | string | Access-Accept |  | 
| Mikrotik-Switching-Filter | 14988 (Mikrotik) | 30 | string | Access-Accept | Allows to create dynamic switch rules when authenticating clients with dot1x server. | 

| Value | Description | 
|---|---|
| 1 | Point-to-Point Tunneling Protocol (PPTP) | 
| 2 | Layer Two Forwarding (L2F) | 
| 3 | Layer Two Tunneling Protocol (L2TP) | 
| 4 | Ascend Tunnel Management Protocol (ATMP) | 
| 5 | Virtual Tunneling Protocol (VTP) | 
| 6 | IP Authentication Header in the Tunnel-mode (AH) | 
| 7 | IP-in-IP Encapsulation (IP-IP) | 
| 8 | Minimal IP-in-IP Encapsulation (MIN-IP-IP) | 
| 9 | IP Encapsulating Security Payload in the Tunnel-mode (ESP) | 
| 10 | Generic Route Encapsulation (GRE) | 
| 11 | Bay Dial Virtual Services (DVS) | 
| 12 | IP-in-IP Tunneling | 
| 13 | Virtual LAN | 

