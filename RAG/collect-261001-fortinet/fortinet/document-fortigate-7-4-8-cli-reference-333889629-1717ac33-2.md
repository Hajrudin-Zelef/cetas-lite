---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-cli-reference-333889629-1717ac33-2
title: "document-fortigate-7-4-8-cli-reference-333889629-1717ac33"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "asic"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-cli-reference-333889629-1717ac33.md
source_anchor: ""
source_lines: [186, 282]
sha256: cc2ec972fa539bff228e6fbbee2d29ba3fabab12264cf9c157c578de4406db29
---

# document-fortigate-7-4-8-cli-reference-333889629-1717ac33

| Parameter | Description | Type | Size | Default | 
|---|---|---|---|---|
| action | Policy action (accept/deny/ipsec). | option | - | deny | 
|  |  |  |  |  | 
| anti-replay | Enable/disable anti-replay check. | option | - | enable | 
|  |  |  |  |  | 
| application-list | Name of an existing Application list. | string | Maximum length: 35 |  | 
| auth-cert | HTTPS server certificate for policy authentication. | string | Maximum length: 35 |  | 
| auth-path | Enable/disable authentication-based routing. | option | - | disable | 
|  |  |  |  |  | 
| auth-redirect-addr | HTTP-to-HTTPS redirect address for firewall authentication. | string | Maximum length: 63 |  | 
| auto-asic-offload * | Enable/disable policy traffic ASIC offloading. | option | - | enable | 
|  |  |  |  |  | 
| av-profile | Name of an existing Antivirus profile. | string | Maximum length: 35 |  | 
| block-notification | Enable/disable block notification. | option | - | disable | 
|  |  |  |  |  | 
| captive-portal-exempt | Enable to exempt some users from the captive portal. | option | - | disable | 
|  |  |  |  |  | 
| capture-packet * | Enable/disable capture packets. | option | - | disable | 
|  |  |  |  |  | 
| casb-profile | Name of an existing CASB profile. | string | Maximum length: 35 |  | 
| cifs-profile | Name of an existing CIFS profile. | string | Maximum length: 35 |  | 
| comments | Comment. | var-string | Maximum length: 1023 |  | 
| custom-log-fields <field-id> | Custom fields to append to log messages for this policy. Custom log field. | string | Maximum length: 35 |  | 
| decrypted-traffic-mirror | Decrypted traffic mirror. | string | Maximum length: 35 |  | 
| delay-tcp-npu-session | Enable TCP NPU session delay to guarantee packet order of 3-way handshake. | option | - | disable | 
|  |  |  |  |  | 
| diameter-filter-profile | Name of an existing Diameter filter profile. | string | Maximum length: 35 |  | 
| diffserv-copy | Enable to copy packet's DiffServ values from session's original direction to its reply direction. | option | - | disable | 
|  |  |  |  |  | 
| diffserv-forward | Enable to change packet's DiffServ values to the specified diffservcode-forward value. | option | - | disable | 
|  |  |  |  |  | 
| diffserv-reverse | Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. | option | - | disable | 
|  |  |  |  |  | 
| diffservcode-forward | Change packet's DiffServ to this value. | user | Not Specified |  | 
| diffservcode-rev | Change packet's reverse (reply) DiffServ to this value. | user | Not Specified |  | 
| disclaimer | Enable/disable user authentication disclaimer. | option | - | disable | 
|  |  |  |  |  | 
| dlp-profile | Name of an existing DLP profile. | string | Maximum length: 35 |  | 
| dnsfilter-profile | Name of an existing DNS filter profile. | string | Maximum length: 35 |  | 
| dsri | Enable DSRI to ignore HTTP server responses. | option | - | disable | 
|  |  |  |  |  | 
| dstaddr <name> | Destination IPv4 address and address group names. Address name. | string | Maximum length: 79 |  | 
| dstaddr-negate | When enabled dstaddr specifies what the destination address must NOT be. | option | - | disable | 
|  |  |  |  |  | 
| dstaddr6 <name> | Destination IPv6 address name and address group names. Address name. | string | Maximum length: 79 |  | 
| dstaddr6-negate | When enabled dstaddr6 specifies what the destination address must NOT be. | option | - | disable | 
|  |  |  |  |  | 
| dstintf <name> | Outgoing (egress) interface. Interface name. | string | Maximum length: 79 |  | 
| dynamic-shaping | Enable/disable dynamic RADIUS defined traffic shaping. | option | - | disable | 
|  |  |  |  |  | 
| email-collect | Enable/disable email collection. | option | - | disable | 
|  |  |  |  |  | 
| emailfilter-profile | Name of an existing email filter profile. | string | Maximum length: 35 |  | 
| fec | Enable/disable Forward Error Correction on traffic matching this policy on a FEC device. | option | - | disable | 
|  |  |  |  |  | 
| file-filter-profile | Name of an existing file-filter profile. | string | Maximum length: 35 |  | 
| firewall-session-dirty | How to handle sessions if the configuration of this firewall policy changes. | option | - | check-all | 
|  |  |  |  |  | 
| fixedport | Enable to prevent source NAT from changing a session's source port. | option | - | disable | 
|  |  |  |  |  | 
| fsso-agent-for-ntlm | FSSO agent to use for NTLM authentication. | string | Maximum length: 35 |  | 
| fsso-groups <name> | Names of FSSO groups. Names of FSSO groups. | string | Maximum length: 511 |  | 
| geoip-anycast | Enable/disable recognition of anycast IP addresses using the geography IP database. | option | - | disable | 
|  |  |  |  |  | 
| geoip-match | Match geography address based either on its physical location or registered location. | option | - | physical-location | 
|  |  |  |  |  | 
| groups <name> | Names of user groups that can authenticate with this policy. Group name. | string | Maximum length: 79 |  | 
| http-policy-redirect | Redirect HTTP(S) traffic to matching transparent web proxy policy. | option | - | disable | 
|  |  |  |  |  | 
| icap-profile | Name of an existing ICAP profile. | string | Maximum length: 35 |  | 
| identity-based-route | Name of identity-based routing rule. | string | Maximum length: 35 |  | 
| inbound | Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN. | option | - | disable | 
|  |  |  |  |  | 
| inspection-mode | Policy inspection mode (Flow/proxy). Default is Flow mode. | option | - | flow | 
|  |  |  |  |  | 
| internet-service | Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. | option | - | disable | 
|  |  |  |  |  | 
| internet-service-custom <name> | Custom Internet Service name. Custom Internet Service name. | string | Maximum length: 79 |  | 
| internet-service-custom-group <name> | Custom Internet Service group name. Custom Internet Service group name. | string | Maximum length: 79 |  | 
| internet-service-group <name> | Internet Service group name. Internet Service group name. | string | Maximum length: 79 |  | 
| internet-service-name <name> | Internet Service name. Internet Service name. | string | Maximum length: 79 |  | 
| internet-service-negate | When enabled internet-service specifies what the service must NOT be. | option | - | disable | 
|  |  |  |  |  | 
| internet-service-src | Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. | option | - | disable | 
|  |  |  |  |  | 
| internet-service-src-custom <name> | Custom Internet Service source name. Custom Internet Service name. | string | Maximum length: 79 |  | 
| internet-service-src-custom-group <name> | Custom Internet Service source group name. Custom Internet Service group name. | string | Maximum length: 79 |  | 
| internet-service-src-group <name> | Internet Service source group name. Internet Service group name. | string | Maximum length: 79 |  | 
| internet-service-src-name <name> | Internet Service source name. Internet Service name. | string | Maximum length: 79 |  | 
| internet-service-src-negate | When enabled internet-service-src specifies what the service must NOT be. | option | - | disable | 
|  |  |  |  |  | 
| internet-service6 | Enable/disable use of IPv6 Internet Services for this policy. If enabled, destination address and service are not used. | option | - | disable | 
|  |  |  |  |  | 
| internet-service6-custom <name> | Custom IPv6 Internet Service name. Custom Internet Service name. | string | Maximum length: 79 |  | 
| internet-service6-custom-group <name> | Custom Internet Service6 group name. Custom Internet Service6 group name. | string | Maximum length: 79 |  | 
| internet-service6-group <name> | Internet Service group name. Internet Service group name. | string | Maximum length: 79 |  | 
