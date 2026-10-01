---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-cli-reference-333889629-1717ac33-3
title: "document-fortigate-7-4-8-cli-reference-333889629-1717ac33"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["0000-00-00"]
keywords: ["agent", "asic"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-cli-reference-333889629-1717ac33.md
source_anchor: ""
source_lines: [283, 381]
sha256: 8e47ac3a5047fa2ee7e571024412eba06b6510644026176ed6c4b55b51339b0c
---

# document-fortigate-7-4-8-cli-reference-333889629-1717ac33

| internet-service6-name <name> | IPv6 Internet Service name. IPv6 Internet Service name. | string | Maximum length: 79 |  | 
| internet-service6-negate | When enabled internet-service6 specifies what the service must NOT be. | option | - | disable | 
|  |  |  |  |  | 
| internet-service6-src | Enable/disable use of IPv6 Internet Services in source for this policy. If enabled, source address is not used. | option | - | disable | 
|  |  |  |  |  | 
| internet-service6-src-custom <name> | Custom IPv6 Internet Service source name. Custom Internet Service name. | string | Maximum length: 79 |  | 
| internet-service6-src-custom-group <name> | Custom Internet Service6 source group name. Custom Internet Service6 group name. | string | Maximum length: 79 |  | 
| internet-service6-src-group <name> | Internet Service6 source group name. Internet Service group name. | string | Maximum length: 79 |  | 
| internet-service6-src-name <name> | IPv6 Internet Service source name. Internet Service name. | string | Maximum length: 79 |  | 
| internet-service6-src-negate | When enabled internet-service6-src specifies what the service must NOT be. | option | - | disable | 
|  |  |  |  |  | 
| ippool | Enable to use IP Pools for source NAT. | option | - | disable | 
|  |  |  |  |  | 
| ips-sensor | Name of an existing IPS sensor. | string | Maximum length: 35 |  | 
| ips-voip-filter | Name of an existing VoIP (ips) profile. | string | Maximum length: 35 |  | 
| logtraffic | Enable or disable logging. Log all sessions or security profile sessions. | option | - | utm | 
|  |  |  |  |  | 
| logtraffic-start | Record logs when a session starts. | option | - | disable | 
|  |  |  |  |  | 
| match-vip | Enable to match packets that have had their destination addresses changed by a VIP. | option | - | enable | 
|  |  |  |  |  | 
| match-vip-only | Enable/disable matching of only those packets that have had their destination addresses changed by a VIP. | option | - | disable | 
|  |  |  |  |  | 
| name | Policy name. | string | Maximum length: 35 |  | 
| nat | Enable/disable source NAT. | option | - | disable | 
|  |  |  |  |  | 
| nat46 | Enable/disable NAT46. | option | - | disable | 
|  |  |  |  |  | 
| nat64 | Enable/disable NAT64. | option | - | disable | 
|  |  |  |  |  | 
| natinbound | Policy-based IPsec VPN: apply destination NAT to inbound traffic. | option | - | disable | 
|  |  |  |  |  | 
| natip | Policy-based IPsec VPN: source NAT IP address for outgoing traffic. | ipv4-classnet | Not Specified | 0.0.0.0 0.0.0.0 | 
| natoutbound | Policy-based IPsec VPN: apply source NAT to outbound traffic. | option | - | disable | 
|  |  |  |  |  | 
| network-service-dynamic <name> | Dynamic Network Service name. Dynamic Network Service name. | string | Maximum length: 79 |  | 
| network-service-src-dynamic <name> | Dynamic Network Service source name. Dynamic Network Service name. | string | Maximum length: 79 |  | 
| np-acceleration * | Enable/disable UTM Network Processor acceleration. | option | - | enable | 
|  |  |  |  |  | 
| ntlm | Enable/disable NTLM authentication. | option | - | disable | 
|  |  |  |  |  | 
| ntlm-enabled-browsers <user-agent-string> | HTTP-User-Agent value of supported browsers. User agent string. | string | Maximum length: 79 |  | 
| ntlm-guest | Enable/disable NTLM guest user access. | option | - | disable | 
|  |  |  |  |  | 
| outbound | Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN. | option | - | enable | 
|  |  |  |  |  | 
| passive-wan-health-measurement | Enable/disable passive WAN health measurement. When enabled, auto-asic-offload is disabled. | option | - | disable | 
|  |  |  |  |  | 
| pcp-inbound | Enable/disable PCP inbound DNAT. | option | - | disable | 
|  |  |  |  |  | 
| pcp-outbound | Enable/disable PCP outbound SNAT. | option | - | disable | 
|  |  |  |  |  | 
| pcp-poolname <name> | PCP pool names. PCP pool name. | string | Maximum length: 79 |  | 
| per-ip-shaper | Per-IP traffic shaper. | string | Maximum length: 35 |  | 
| permit-any-host | Accept UDP packets from any host. | option | - | disable | 
|  |  |  |  |  | 
| permit-stun-host | Accept UDP packets from any Session Traversal Utilities for NAT (STUN) host. | option | - | disable | 
|  |  |  |  |  | 
| policy-expiry | Enable/disable policy expiry. | option | - | disable | 
|  |  |  |  |  | 
| policy-expiry-date | Policy expiry date (YYYY-MM-DD HH:MM:SS). | datetime | Not Specified | 0000-00-00 00:00:00 | 
| policy-expiry-date-utc | Policy expiry date and time, in epoch format. | user | Not Specified |  | 
| policyid | Policy ID (0 - 4294967294). | integer | Minimum value: 0 Maximum value: 4294967294 | 0 | 
| poolname <name> | IP Pool names. IP pool name. | string | Maximum length: 79 |  | 
| poolname6 <name> | IPv6 pool names. IPv6 pool name. | string | Maximum length: 79 |  | 
| port-preserve | Enable/disable preservation of the original source port from source NAT if it has not been used. | option | - | enable | 
|  |  |  |  |  | 
| profile-group | Name of profile group. | string | Maximum length: 35 |  | 
| profile-protocol-options | Name of an existing Protocol options profile. | string | Maximum length: 35 | default | 
| profile-type | Determine whether the firewall policy allows security profile groups or single profiles only. | option | - | single | 
|  |  |  |  |  | 
| radius-mac-auth-bypass | Enable MAC authentication bypass. The bypassed MAC address must be received from RADIUS server. | option | - | disable | 
|  |  |  |  |  | 
| redirect-url | URL users are directed to after seeing and accepting the disclaimer or authenticating. | var-string | Maximum length: 1023 |  | 
| replacemsg-override-group | Override the default replacement message group for this policy. | string | Maximum length: 35 |  | 
| reputation-direction | Direction of the initial traffic for reputation to take effect. | option | - | destination | 
|  |  |  |  |  | 
| reputation-direction6 | Direction of the initial traffic for IPv6 reputation to take effect. | option | - | destination | 
|  |  |  |  |  | 
| reputation-minimum | Minimum Reputation to take action. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| reputation-minimum6 | IPv6 Minimum Reputation to take action. | integer | Minimum value: 0 Maximum value: 4294967295 | 0 | 
| rtp-addr <name> | Address names if this is an RTP NAT policy. Address name. | string | Maximum length: 79 |  | 
| rtp-nat | Enable Real Time Protocol (RTP) NAT. | option | - | disable | 
|  |  |  |  |  | 
| schedule | Schedule name. | string | Maximum length: 35 |  | 
| schedule-timeout | Enable to force current sessions to end when the schedule object times out. Disable allows them to end from inactivity. | option | - | disable | 
|  |  |  |  |  | 
| sctp-filter-profile | Name of an existing SCTP filter profile. | string | Maximum length: 35 |  | 
| send-deny-packet | Enable to send a reply when a session is denied or blocked by a firewall policy. | option | - | disable | 
|  |  |  |  |  | 
| service <name> | Service and service group names. Service and service group names. | string | Maximum length: 79 |  | 
| service-negate | When enabled service specifies what the service must NOT be. | option | - | disable | 
|  |  |  |  |  | 
| session-ttl | TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL). | user | Not Specified |  | 
| sgt <id> | Security group tags. Security group tag (1 - 65535). | integer | Minimum value: 1 Maximum value: 65535 |  | 
| sgt-check | Enable/disable security group tags (SGT) check. | option | - | disable | 
|  |  |  |  |  | 
| src-vendor-mac <id> | Vendor MAC source ID. Vendor MAC ID. | integer | Minimum value: 0 Maximum value: 4294967295 |  | 
| srcaddr <name> | Source IPv4 address and address group names. Address name. | string | Maximum length: 79 |  | 
