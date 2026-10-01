---
id: collect-261001-fortinet/fortinet/poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5-1
title: "poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent", "distribution"]
source: docs/RAG/collect-261001-fortinet/poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5.md
source_anchor: ""
source_lines: [1, 176]
sha256: 1f8675e392c01a241c83d720f1b8510d56800e1878ea0fe86fbf23acf566837f
---

# poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5

| subcategory | FortiGate System | 
|---|---|
| layout | fortios | 
| page_title | FortiOS: fortios_system_interface | 
| description | Configure interfaces. | 
Configure interfaces.
resource "fortios_system_interface" "example" {
  allow_append = true
  vdomparam    = "root"
  name = "TESTACCINT"
  type = "loopback"
  ip   = "169.254.72.255/32"
  vdom = "root"
}
resource "fortios_system_interface" "example2" {
  allow_append = true
  vdomparam    = "root"
  name = "TESTV6"
  type = "loopback"
  ip   = "169.254.71.255/32"
  vdom = "root"
  secondary_ip = "enable"
  ipv6 {
    ip6_address = "2001:beef::01/128"
  }
  secondaryip {
    ip = "3.3.3.3/32"
  }
}
- 
vdomparam - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
- 
allow_append - If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution! Requiresname to be defined.
- 
dynamic_sort_table -true orfalse , set this parameter totrue when using dynamic for_each + toset to configure and sort sub-tables, if set totrue static sub-tables must be ordered.
- 
ac_name - PPPoE server name.
- 
aggregate - Aggregate interface.
- 
aggregate_type - Type of aggregation. Valid values:physicalvxlan .
- 
algorithm - Frame distribution algorithm. Valid values:L2L3L4Source-MAC .
- 
alias - Alias will be displayed with the interface name to make it easier to distinguish.
- 
allowaccess - Permitted types of management access to this interface. Valid values:pinghttpssshsnmphttptelnetfgfmradius-acctprobe-responsefabricftmspeed-test .
- 
ap_discover - Enable/disable automatic registration of unknown FortiAP devices. Valid values:enabledisable .
- 
arpforward - Enable/disable ARP forwarding. Valid values:enabledisable .
- 
auth_cert - HTTPS server certificate. This attribute must reference one of the following datasources:vpn.certificate.local.name .
- 
auth_portal_addr - Address of captive portal.
- 
auth_type - PPP authentication type to use. Valid values:autopapchapmschapv1mschapv2 .
- 
auto_auth_extension_device - Enable/disable automatic authorization of dedicated Fortinet extension device on this interface. Valid values:enabledisable .
- 
bandwidth_measure_time - Bandwidth measure time.
- 
bfd - Bidirectional Forwarding Detection (BFD) settings. Valid values:globalenabledisable .
- 
bfd_desired_min_tx - BFD desired minimal transmit interval.
- 
bfd_detect_mult - BFD detection multiplier.
- 
bfd_required_min_rx - BFD required minimal receive interval.
- 
broadcast_forticlient_discovery - Enable/disable broadcasting FortiClient discovery messages. Valid values:enabledisable .
- 
broadcast_forward - Enable/disable broadcast forwarding. Valid values:enabledisable .
- 
cli_conn_status - CLI connection status.
- 
color - Color of icon on the GUI.
- 
dedicated_to - Configure interface for single purpose. Valid values:nonemanagement .
- 
defaultgw - Enable to get the gateway IP from the DHCP or PPPoE server. Valid values:enabledisable .
- 
description - Description.
- 
detected_peer_mtu - MTU of detected peer (0 - 4294967295).
- 
detectprotocol - Protocols used to detect the server. Valid values:pingtcp-echoudp-echo .
- 
detectserver - Gateway's ping server for this IP.
- 
device_identification - Enable/disable passively gathering of device identity information about the devices on the network connected to this interface. Valid values:enabledisable .
- 
device_user_identification - Enable/disable passive gathering of user identity information about users on this interface. Valid values:enabledisable .
- 
devindex - Device Index.
- 
dhcp_classless_route_addition - Enable/disable addition of classless static routes retrieved from DHCP server. Valid values:enabledisable .
- 
dhcp_client_identifier - DHCP client identifier.
- 
dhcp_relay_agent_option - Enable/disable DHCP relay agent option. Valid values:enabledisable .
- 
dhcp_relay_interface - Specify outgoing interface to reach server. This attribute must reference one of the following datasources:system.interface.name .
- 
dhcp_relay_interface_select_method - Specify how to select outgoing interface to reach server. Valid values:autosdwanspecify .
- 
dhcp_relay_ip - DHCP relay IP address.
- 
dhcp_relay_link_selection - DHCP relay link selection.
- 
dhcp_relay_request_all_server - Enable/disable sending of DHCP requests to all servers. Valid values:disableenable .
- 
dhcp_relay_service - Enable/disable allowing this interface to act as a DHCP relay. Valid values:disableenable .
- 
dhcp_relay_type - DHCP relay type (regular or IPsec). Valid values:regularipsec .
- 
dhcp_renew_time - DHCP renew time in seconds (300-604800), 0 means use the renew time provided by the server.
- 
disc_retry_timeout - Time in seconds to wait before retrying to start a PPPoE discovery, 0 means no timeout.
- 
disconnect_threshold - Time in milliseconds to wait before sending a notification that this interface is down or disconnected.
- 
distance - Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
- 
dns_server_override - Enable/disable use DNS acquired by DHCP or PPPoE. Valid values:enabledisable .
- 
dns_server_protocol - DNS transport protocols. Valid values:cleartextdotdoh .
- 
drop_fragment - Enable/disable drop fragment packets. Valid values:enabledisable .
- 
drop_overlapped_fragment - Enable/disable drop overlapped fragment packets. Valid values:enabledisable .
- 
eap_ca_cert - EAP CA certificate name. This attribute must reference one of the following datasources:certificate.ca.name .
- 
eap_identity - EAP identity.
- 
eap_method - EAP method. Valid values:tlspeap .
- 
eap_password - EAP password.
- 
eap_supplicant - Enable/disable EAP-Supplicant. Valid values:enabledisable .
- 
eap_user_cert - EAP user certificate name. This attribute must reference one of the following datasources:certificate.local.name .
- 
egress_cos - Override outgoing CoS in user VLAN tag. Valid values:disablecos0cos1cos2cos3cos4cos5cos6cos7 .
- 
egress_shaping_profile - Outgoing traffic shaping profile. This attribute must reference one of the following datasources:firewall.shaping-profile.profile-name .
- 
eip - External IP.
- 
estimated_downstream_bandwidth - Estimated maximum downstream bandwidth (kbps). Used to estimate link utilization.
- 
estimated_upstream_bandwidth - Estimated maximum upstream bandwidth (kbps). Used to estimate link utilization.
- 
explicit_ftp_proxy - Enable/disable the explicit FTP proxy on this interface. Valid values:enabledisable .
- 
explicit_web_proxy - Enable/disable the explicit web proxy on this interface. Valid values:enabledisable .
- 
external - Enable/disable identifying the interface as an external interface (which usually means it's connected to the Internet). Valid values:enabledisable .
- 
fail_action_on_extender - Action on FortiExtender when interface fail. Valid values:soft-restarthard-restartreboot .
- 
fail_alert_method - Select link-failed-signal or link-down method to alert about a failed link. Valid values:link-failed-signallink-down .
- 
fail_detect - Enable/disable fail detection features for this interface. Valid values:enabledisable .
- 
fail_detect_option - Options for detecting that this interface has failed. Valid values:detectserverlink-down .
- 
fortilink - Enable FortiLink to dedicate this interface to manage other Fortinet devices. Valid values:enabledisable .
- 
fortilink_backup_link - FortiLink split interface backup link.
- 
fortilink_neighbor_detect - Protocol for FortiGate neighbor discovery. Valid values:lldpfortilink .
- 
fortilink_split_interface - Enable/disable FortiLink split interface to connect member link to different FortiSwitch in stack for uplink redundancy. Valid values:enabledisable .
- 
