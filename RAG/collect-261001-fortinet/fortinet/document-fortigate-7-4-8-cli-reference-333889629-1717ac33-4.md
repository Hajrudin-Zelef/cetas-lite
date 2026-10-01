---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-cli-reference-333889629-1717ac33-4
title: "document-fortigate-7-4-8-cli-reference-333889629-1717ac33"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["asic"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-cli-reference-333889629-1717ac33.md
source_anchor: ""
source_lines: [382, 526]
sha256: 0ac1c4eb038672df35c73877dce0dda6ca5a7ad42b4089712bb1beadc0637157
---

# document-fortigate-7-4-8-cli-reference-333889629-1717ac33

| srcaddr-negate | When enabled srcaddr specifies what the source address must NOT be. | option | - | disable | 
|  |  |  |  |  | 
| srcaddr6 <name> | Source IPv6 address name and address group names. Address name. | string | Maximum length: 79 |  | 
| srcaddr6-negate | When enabled srcaddr6 specifies what the source address must NOT be. | option | - | disable | 
|  |  |  |  |  | 
| srcintf <name> | Incoming (ingress) interface. Interface name. | string | Maximum length: 79 |  | 
| ssh-filter-profile | Name of an existing SSH filter profile. | string | Maximum length: 35 |  | 
| ssh-policy-redirect | Redirect SSH traffic to matching transparent proxy policy. | option | - | disable | 
|  |  |  |  |  | 
| ssl-ssh-profile | Name of an existing SSL SSH profile. | string | Maximum length: 35 | no-inspection | 
| status | Enable or disable this policy. | option | - | enable | 
|  |  |  |  |  | 
| tcp-mss-receiver | Receiver TCP maximum segment size (MSS). | integer | Minimum value: 0 Maximum value: 65535 | 0 | 
| tcp-mss-sender | Sender TCP maximum segment size (MSS). | integer | Minimum value: 0 Maximum value: 65535 | 0 | 
| tcp-session-without-syn | Enable/disable creation of TCP session without SYN flag. | option | - | disable | 
|  |  |  |  |  | 
| timeout-send-rst | Enable/disable sending RST packets when TCP sessions expire. | option | - | disable | 
|  |  |  |  |  | 
| tos | ToS (Type of Service) value used for comparison. | user | Not Specified |  | 
| tos-mask | Non-zero bit positions are used for comparison while zero bit positions are ignored. | user | Not Specified |  | 
| tos-negate | Enable negated TOS match. | option | - | disable | 
|  |  |  |  |  | 
| traffic-shaper | Traffic shaper. | string | Maximum length: 35 |  | 
| traffic-shaper-reverse | Reverse traffic shaper. | string | Maximum length: 35 |  | 
| users <name> | Names of individual users that can authenticate with this policy. Names of individual users that can authenticate with this policy. | string | Maximum length: 79 |  | 
| utm-status | Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy. | option | - | disable | 
|  |  |  |  |  | 
| uuid | Universally Unique Identifier (UUID; automatically assigned but can be manually reset). | uuid | Not Specified | 00000000-0000-0000-0000-000000000000 | 
| videofilter-profile | Name of an existing VideoFilter profile. | string | Maximum length: 35 |  | 
| virtual-patch-profile | Name of an existing virtual-patch profile. | string | Maximum length: 35 |  | 
| vlan-cos-fwd | VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest. | integer | Minimum value: 0 Maximum value: 7 | 255 | 
| vlan-cos-rev | VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest. | integer | Minimum value: 0 Maximum value: 7 | 255 | 
| vlan-filter | VLAN ranges to allow | user | Not Specified |  | 
| voip-profile | Name of an existing VoIP (voipd) profile. | string | Maximum length: 35 |  | 
| vpntunnel | Policy-based IPsec VPN: name of the IPsec VPN Phase 1. | string | Maximum length: 35 |  | 
| waf-profile | Name of an existing Web application firewall profile. | string | Maximum length: 35 |  | 
| wanopt * | Enable/disable WAN optimization. | option | - | disable | 
|  |  |  |  |  | 
| wanopt-detection * | WAN optimization auto-detection mode. | option | - | active | 
|  |  |  |  |  | 
| wanopt-passive-opt * | WAN optimization passive mode options. This option decides what IP address will be used to connect server. | option | - | default | 
|  |  |  |  |  | 
| wanopt-peer * | WAN optimization peer. | string | Maximum length: 35 |  | 
| wanopt-profile * | WAN optimization profile. | string | Maximum length: 35 |  | 
| wccp | Enable/disable forwarding traffic matching this policy to a configured WCCP server. | option | - | disable | 
|  |  |  |  |  | 
| webcache * | Enable/disable web cache. | option | - | disable | 
|  |  |  |  |  | 
| webcache-https * | Enable/disable web cache for HTTPS. | option | - | disable | 
|  |  |  |  |  | 
| webfilter-profile | Name of an existing Web filter profile. | string | Maximum length: 35 |  | 
| webproxy-forward-server | Webproxy forward server name. | string | Maximum length: 63 |  | 
| webproxy-profile | Webproxy profile name. | string | Maximum length: 63 |  | 
| ztna-device-ownership | Enable/disable zero trust device ownership. | option | - | disable | 
|  |  |  |  |  | 
| ztna-ems-tag <name> | Source ztna-ems-tag names. Address name. | string | Maximum length: 79 |  | 
| ztna-ems-tag-secondary <name> | Source ztna-ems-tag-secondary names. Address name. | string | Maximum length: 79 |  | 
| ztna-geo-tag <name> | Source ztna-geo-tag names. Address name. | string | Maximum length: 79 |  | 
| ztna-policy-redirect | Redirect ZTNA traffic to matching Access-Proxy proxy-policy. | option | - | disable | 
|  |  |  |  |  | 
| ztna-status | Enable/disable zero trust access. | option | - | disable | 
|  |  |  |  |  | 
| ztna-tags-match-logic | ZTNA tag matching logic. | option | - | or | 
|  |  |  |  |  | 
| Option | Description | 
|---|---|
| accept | Allows session that match the firewall policy. | 
| deny | Blocks sessions that match the firewall policy. | 
| ipsec | Firewall policy becomes a policy-based IPsec VPN policy. | 
| Option | Description | 
|---|---|
| enable | Enable anti-replay check. | 
| disable | Disable anti-replay check. | 
| Option | Description | 
|---|---|
| enable | Enable authentication-based routing. | 
| disable | Disable authentication-based routing. | 
| Option | Description | 
|---|---|
| enable | Enable auto ASIC offloading. | 
| disable | Disable ASIC offloading. | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
| Option | Description | 
|---|---|
| enable | Enable exemption of captive portal. | 
| disable | Disable exemption of captive portal. | 
| Option | Description | 
|---|---|
| enable | Enable capture packets. | 
| disable | Disable capture packets. | 
| Option | Description | 
|---|---|
| enable | Enable TCP NPU session delay in order to guarantee packet order of 3-way handshake. | 
| disable | Disable TCP NPU session delay in order to guarantee packet order of 3-way handshake. | 
| Option | Description | 
|---|---|
| enable | Enable DSCP copy. | 
| disable | Disable DSCP copy. | 
| Option | Description | 
|---|---|
| enable | Enable setting forward (original) traffic Diffserv. | 
| disable | Disable setting forward (original) traffic Diffserv. | 
| Option | Description | 
|---|---|
| enable | Enable setting reverse (reply) traffic DiffServ. | 
| disable | Disable setting reverse (reply) traffic DiffServ. | 
| Option | Description | 
|---|---|
| enable | Enable user authentication disclaimer. | 
| disable | Disable user authentication disclaimer. | 
| Option | Description | 
|---|---|
| enable | Enable DSRI. | 
| disable | Disable DSRI. | 
| Option | Description | 
|---|---|
| enable | Enable destination address negate. | 
| disable | Disable destination address negate. | 
| Option | Description | 
|---|---|
| enable | Enable IPv6 destination address negate. | 
| disable | Disable IPv6 destination address negate. | 
| Option | Description | 
|---|---|
| enable | Enable dynamic RADIUS defined traffic shaping. | 
| disable | Disable dynamic RADIUS defined traffic shaping. | 
| Option | Description | 
|---|---|
| enable | Enable email collection. | 
| disable | Disable email collection. | 
| Option | Description | 
|---|---|
| enable | Enable Forward Error Correction. | 
| disable | Disable Forward Error Correction. | 
| Option | Description | 
|---|---|
| check-all | Flush all current sessions accepted by this policy. These sessions must be started and re-matched with policies. | 
| check-new | Continue to allow sessions already accepted by this policy. | 
| Option | Description | 
|---|---|
| enable | Enable setting. | 
| disable | Disable setting. | 
