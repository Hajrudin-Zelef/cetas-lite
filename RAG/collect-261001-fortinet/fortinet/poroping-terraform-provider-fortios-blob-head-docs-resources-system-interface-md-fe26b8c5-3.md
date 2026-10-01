---
id: collect-261001-fortinet/fortinet/poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5-3
title: "poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-fortinet/poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5.md
source_anchor: ""
source_lines: [361, 492]
sha256: 60e144d51c09fbcafd34697bf399b525e20bdae3b91522564b01420244f0ebf6
---

# poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5

stp_ha_secondary - Control STP behavior on HA secondary. Valid values:disableenablepriority-adjust .
- 
stpforward - Enable/disable STP forwarding. Valid values:enabledisable .
- 
stpforward_mode - Configure STP forwarding mode. Valid values:rpl-all-ext-idrpl-bridge-ext-idrpl-nothing .
- 
subst - Enable to always send packets from this interface to a destination MAC address. Valid values:enabledisable .
- 
substitute_dst_mac - Destination MAC address that all packets are sent to from this interface.
- 
swc_first_create - Initial create for switch-controller VLANs.
- 
swc_vlan - Creation status for switch-controller VLANs.
- 
switch - Contained in switch.
- 
switch_controller_access_vlan - Block FortiSwitch port-to-port traffic. Valid values:enabledisable .
- 
switch_controller_arp_inspection - Enable/disable FortiSwitch ARP inspection. Valid values:enabledisable .
- 
switch_controller_dhcp_snooping - Switch controller DHCP snooping. Valid values:enabledisable .
- 
switch_controller_dhcp_snooping_option82 - Switch controller DHCP snooping option82. Valid values:enabledisable .
- 
switch_controller_dhcp_snooping_verify_mac - Switch controller DHCP snooping verify MAC. Valid values:enabledisable .
- 
switch_controller_dynamic - Integrated FortiLink settings for managed FortiSwitch. This attribute must reference one of the following datasources:switch-controller.fortilink-settings.name .
- 
switch_controller_feature - Interface's purpose when assigning traffic (read only). Valid values:nonedefault-vlanquarantinerspanvoicevideonacnac-segment .
- 
switch_controller_igmp_snooping - Switch controller IGMP snooping. Valid values:enabledisable .
- 
switch_controller_igmp_snooping_fast_leave - Switch controller IGMP snooping fast-leave. Valid values:enabledisable .
- 
switch_controller_igmp_snooping_proxy - Switch controller IGMP snooping proxy. Valid values:enabledisable .
- 
switch_controller_iot_scanning - Enable/disable managed FortiSwitch IoT scanning. Valid values:enabledisable .
- 
switch_controller_learning_limit - Limit the number of dynamic MAC addresses on this VLAN (1 - 128, 0 = no limit, default).
- 
switch_controller_mgmt_vlan - VLAN to use for FortiLink management purposes.
- 
switch_controller_nac - Integrated FortiLink settings for managed FortiSwitch. This attribute must reference one of the following datasources:switch-controller.fortilink-settings.name .
- 
switch_controller_netflow_collect - NetFlow collection and processing. Valid values:disableenable .
- 
switch_controller_rspan_mode - Stop Layer2 MAC learning and interception of BPDUs and other packets on this interface. Valid values:disableenable .
- 
switch_controller_source_ip - Source IP address used in FortiLink over L3 connections. Valid values:outboundfixed .
- 
switch_controller_traffic_policy - Switch controller traffic policy for the VLAN. This attribute must reference one of the following datasources:switch-controller.traffic-policy.name .
- 
system_id - Define a system ID for the aggregate interface.
- 
system_id_type - Method in which system ID is generated. Valid values:autouser .
- 
tcp_mss - TCP maximum segment size. 0 means do not change segment size.
- 
trunk - Enable/disable VLAN trunk. Valid values:enabledisable .
- 
trust_ip_1 - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
- 
trust_ip_2 - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
- 
trust_ip_3 - Trusted host for dedicated management traffic (0.0.0.0/24 for all hosts).
- 
trust_ip6_1 - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
- 
trust_ip6_2 - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
- 
trust_ip6_3 - Trusted IPv6 host for dedicated management traffic (::/0 for all hosts).
- 
type - Interface type. Valid values:physicalvlanaggregateredundanttunnelvdom-linkloopbackswitchhard-switchvap-switchwl-meshfext-wanvxlangenevehdlcswitch-vlanemac-vlanssllan-extension .
- 
username - Username of the PPPoE account, provided by your ISP.
- 
vdom - Interface is in this virtual domain (VDOM). This attribute must reference one of the following datasources:system.vdom.name .
- 
vindex - Switch control interface VLAN ID.
- 
vlan_protocol - Ethernet protocol of VLAN. Valid values:8021q8021ad .
- 
vlanforward - Enable/disable traffic forwarding between VLANs on this interface. Valid values:enabledisable .
- 
vlanid - VLAN ID (1 - 4094).
- 
vrf - Virtual Routing Forwarding ID.
- 
vrrp_virtual_mac - Enable/disable use of virtual MAC for VRRP. Valid values:enabledisable .
- 
wccp - Enable/disable WCCP on this interface. Used for encapsulated WCCP communication between WCCP clients and servers. Valid values:enabledisable .
- 
weight - Default weight for static routes (if route has no weight configured).
- 
wins_ip - WINS server IP.
- 
client_options - DHCP client options. The structure ofclient_options block is documented below.
The client_options block contains:
- code - DHCP client option code.
- id - ID.
- ip - DHCP option IPs.
- type - DHCP client option type. Valid values:hexstringipfqdn .
- value - DHCP client option value.
- dhcp_snooping_server_list - Configure DHCP server access list. The structure ofdhcp_snooping_server_list block is documented below.
The dhcp_snooping_server_list block contains:
- name - DHCP server name.
- server_ip - IP address for DHCP server.
- egress_queues - Configure queues of NP port on egress path. The structure ofegress_queues block is documented below.
The egress_queues block contains:
- cos0 - CoS profile name for CoS 0. This attribute must reference one of the following datasources:system.isf-queue-profile.name .
- cos1 - CoS profile name for CoS 1. This attribute must reference one of the following datasources:system.isf-queue-profile.name .
- cos2 - CoS profile name for CoS 2. This attribute must reference one of the following datasources:system.isf-queue-profile.name .
- cos3 - CoS profile name for CoS 3. This attribute must reference one of the following datasources:system.isf-queue-profile.name .
- cos4 - CoS profile name for CoS 4. This attribute must reference one of the following datasources:system.isf-queue-profile.name .
- cos5 - CoS profile name for CoS 5. This attribute must reference one of the following datasources:system.isf-queue-profile.name .
- cos6 - CoS profile name for CoS 6. This attribute must reference one of the following datasources:system.isf-queue-profile.name .
- cos7 - CoS profile name for CoS 7. This attribute must reference one of the following datasources:system.isf-queue-profile.name .
- fail_alert_interfaces - Names of the FortiGate interfaces to which the link failure alert is sent. The structure offail_alert_interfaces block is documented below.
The fail_alert_interfaces block contains:
- name - Names of the non-virtual interface. This attribute must reference one of the following datasources:system.interface.name .
- ipv6 - IPv6 of interface. The structure ofipv6 block is documented below.
The ipv6 block contains:
- autoconf - Enable/disable address auto config. Valid values:enabledisable .
- cli_conn6_status - CLI IPv6 connection status.
- dhcp6_client_options - DHCPv6 client options. Valid values:rapidiapdiana .
- dhcp6_information_request - Enable/disable DHCPv6 information request. Valid values:enabledisable .
- dhcp6_prefix_delegation - Enable/disable DHCPv6 prefix delegation. Valid values:enabledisable .
- dhcp6_prefix_hint - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
- dhcp6_prefix_hint_plt - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
- dhcp6_prefix_hint_vlt - DHCPv6 prefix hint valid life time (sec).
- dhcp6_relay_ip - DHCPv6 relay IP address.
- dhcp6_relay_service - Enable/disable DHCPv6 relay. Valid values:disableenable .
