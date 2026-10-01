---
id: collect-261001-fortinet/fortinet/poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5-4
title: "poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5.md
source_anchor: ""
source_lines: [493, 600]
sha256: edb7ca8b6dca3acd677431fac4e67b2c93579fdf7ac59fb8eee3d37b1478d613
---

# poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5

- dhcp6_relay_source_interface - Enable/disable use of address on this interface as the source address of the relay message. Valid values:disableenable .
- dhcp6_relay_type - DHCPv6 relay type. Valid values:regular .
- icmp6_send_redirect - Enable/disable sending of ICMPv6 redirects. Valid values:enabledisable .
- interface_identifier - IPv6 interface identifier.
- ip6_address - Primary IPv6 address prefix. Syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx.
- ip6_allowaccess - Allow management access to the interface. Valid values:pinghttpssshsnmphttptelnetfgfmfabric .
- ip6_default_life - Default life (sec).
- ip6_delegated_prefix_iaid - IAID of obtained delegated-prefix from the upstream interface.
- ip6_dns_server_override - Enable/disable using the DNS server acquired by DHCP. Valid values:enabledisable .
- ip6_hop_limit - Hop limit (0 means unspecified).
- ip6_link_mtu - IPv6 link MTU.
- ip6_manage_flag - Enable/disable the managed flag. Valid values:enabledisable .
- ip6_max_interval - IPv6 maximum interval (4 to 1800 sec).
- ip6_min_interval - IPv6 minimum interval (3 to 1350 sec).
- ip6_mode - Addressing mode (static, DHCP, delegated). Valid values:staticdhcppppoedelegated .
- ip6_other_flag - Enable/disable the other IPv6 flag. Valid values:enabledisable .
- ip6_prefix_mode - Assigning a prefix from DHCP or RA. Valid values:dhcp6ra .
- ip6_reachable_time - IPv6 reachable time (milliseconds; 0 means unspecified).
- ip6_retrans_time - IPv6 retransmit time (milliseconds; 0 means unspecified).
- ip6_send_adv - Enable/disable sending advertisements about the interface. Valid values:enabledisable .
- ip6_subnet - Subnet to routing prefix. Syntax: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx.
- ip6_upstream_interface - Interface name providing delegated information. This attribute must reference one of the following datasources:system.interface.name .
- nd_cert - Neighbor discovery certificate. This attribute must reference one of the following datasources:certificate.local.name .
- nd_cga_modifier - Neighbor discovery CGA modifier.
- nd_mode - Neighbor discovery mode. Valid values:basicSEND-compatible .
- nd_security_level - Neighbor discovery security level (0 - 7; 0 = least secure, default = 0).
- nd_timestamp_delta - Neighbor discovery timestamp delta value (1 - 3600 sec; default = 300).
- nd_timestamp_fuzz - Neighbor discovery timestamp fuzz factor (1 - 60 sec; default = 1).
- ra_send_mtu - Enable/disable sending link MTU in RA packet. Valid values:enabledisable .
- unique_autoconf_addr - Enable/disable unique auto config address. Valid values:enabledisable .
- vrip6_link_local - Link-local IPv6 address of virtual router.
- vrrp_virtual_mac6 - Enable/disable virtual MAC for VRRP. Valid values:enabledisable .
- dhcp6_iapd_list - DHCPv6 IA-PD list. The structure ofdhcp6_iapd_list block is documented below.
The dhcp6_iapd_list block contains:
- iaid - Identity association identifier.
- prefix_hint - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
- prefix_hint_plt - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
- prefix_hint_vlt - DHCPv6 prefix hint valid life time (sec).
- ip6_delegated_prefix_list - Advertised IPv6 delegated prefix list. The structure ofip6_delegated_prefix_list block is documented below.
The ip6_delegated_prefix_list block contains:
- autonomous_flag - Enable/disable the autonomous flag. Valid values:enabledisable .
- delegated_prefix_iaid - IAID of obtained delegated-prefix from the upstream interface.
- onlink_flag - Enable/disable the onlink flag. Valid values:enabledisable .
- prefix_id - Prefix ID.
- rdnss - Recursive DNS server option.
- rdnss_service - Recursive DNS service option. Valid values:delegateddefaultspecify .
- subnet - Add subnet ID to routing prefix.
- upstream_interface - Name of the interface that provides delegated information. This attribute must reference one of the following datasources:system.interface.name .
- ip6_extra_addr - Extra IPv6 address prefixes of interface. The structure ofip6_extra_addr block is documented below.
The ip6_extra_addr block contains:
- prefix - IPv6 address prefix.
- ip6_prefix_list - Advertised prefix list. The structure ofip6_prefix_list block is documented below.
The ip6_prefix_list block contains:
- autonomous_flag - Enable/disable the autonomous flag. Valid values:enabledisable .
- onlink_flag - Enable/disable the onlink flag. Valid values:enabledisable .
- preferred_life_time - Preferred life time (sec).
- prefix - IPv6 prefix.
- rdnss - Recursive DNS server option.
- valid_life_time - Valid life time (sec).
- dnssl - DNS search list option. The structure ofdnssl block is documented below.
The dnssl block contains:
- domain - Domain name.
- vrrp6 - IPv6 VRRP configuration. The structure ofvrrp6 block is documented below.
The vrrp6 block contains:
- accept_mode - Enable/disable accept mode. Valid values:enabledisable .
- adv_interval - Advertisement interval (1 - 255 seconds).
- preempt - Enable/disable preempt mode. Valid values:enabledisable .
- priority - Priority of the virtual router (1 - 255).
- start_time - Startup time (1 - 255 seconds).
- status - Enable/disable VRRP. Valid values:enabledisable .
- vrdst6 - Monitor the route to this destination.
- vrgrp - VRRP group ID (1 - 65535).
- vrid - Virtual router identifier (1 - 255).
- vrip6 - IPv6 address of the virtual router.
- member - Physical interfaces that belong to the aggregate or redundant interface. The structure ofmember block is documented below.
The member block contains:
- interface_name - Physical interface name. This attribute must reference one of the following datasources:system.interface.name .
- secondaryip - Second IP address of interface. The structure ofsecondaryip block is documented below.
The secondaryip block contains:
- allowaccess - Management access settings for the secondary IP address. Valid values:pinghttpssshsnmphttptelnetfgfmradius-acctprobe-responsefabricftmspeed-test .
- detectprotocol - Protocols used to detect the server. Valid values:pingtcp-echoudp-echo .
- detectserver - Gateway's ping server for this IP.
- gwdetect - Enable/disable detect gateway alive for first. Valid values:enabledisable .
- ha_priority - HA election priority for the PING server.
- id - ID.
- ip - Secondary IP address of the interface.
- ping_serv_status - PING server status.
- security_groups - User groups that can authenticate with the captive portal. The structure ofsecurity_groups block is documented below.
The security_groups block contains:
- name - Names of user groups that can authenticate with the captive portal. This attribute must reference one of the following datasources:user.group.name .
- tagging - Config object tagging. The structure oftagging block is documented below.
The tagging block contains:
- category - Tag category. This attribute must reference one of the following datasources:system.object-tagging.category .
- name - Tagging entry name.
- tags - Tags. The structure oftags block is documented below.
The tags block contains:
- name - Tag name. This attribute must reference one of the following datasources:system.object-tagging.tags.name .
- vrrp - VRRP configuration. The structure ofvrrp block is documented below.
The vrrp block contains:
- accept_mode - Enable/disable accept mode. Valid values:enabledisable .
- adv_interval - Advertisement interval (1 - 255 seconds).
- ignore_default_route - Enable/disable ignoring of default route when checking destination. Valid values:enabledisable .
- preempt - Enable/disable preempt mode. Valid values:enabledisable .
- priority - Priority of the virtual router (1 - 255).
- start_time - Startup time (1 - 255 seconds).
- status - Enable/disable this VRRP configuration. Valid values:enabledisable .
- version - VRRP version. Valid values:23 .
- vrdst - Monitor the route to this destination.
