---
id: collect-261001-fortinet/fortinet/poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5-2
title: "poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5.md
source_anchor: ""
source_lines: [177, 360]
sha256: 87fe6d45ea20960eb09f4131948b70b6861fd5f724c2ff55f93acd0b56af7a0c
---

# poroping-terraform-provider-fortios-blob-head-docs-resources-system-interface-md-fe26b8c5

fortilink_stacking - Enable/disable FortiLink switch-stacking on this interface. Valid values:enabledisable .
- 
forward_domain - Transparent mode forward domain.
- 
forward_error_correction - Configure forward error correction (FEC). Valid values:nonedisablecl91-rs-feccl74-fc-fecauto .
- 
gwdetect - Enable/disable detect gateway alive for first. Valid values:enabledisable .
- 
ha_priority - HA election priority for the PING server.
- 
icmp_accept_redirect - Enable/disable ICMP accept redirect. Valid values:enabledisable .
- 
icmp_send_redirect - Enable/disable sending of ICMP redirects. Valid values:enabledisable .
- 
ident_accept - Enable/disable authentication for this interface. Valid values:enabledisable .
- 
idle_timeout - PPPoE auto disconnect after idle timeout seconds, 0 means no timeout.
- 
ike_saml_server - Configure IKE authentication SAML server. This attribute must reference one of the following datasources:user.saml.name .
- 
inbandwidth - Bandwidth limit for incoming traffic (0 - 100000000 kbps), 0 means unlimited.
- 
ingress_cos - Override incoming CoS in user VLAN tag on VLAN interface or assign a priority VLAN tag on physical interface. Valid values:disablecos0cos1cos2cos3cos4cos5cos6cos7 .
- 
ingress_shaping_profile - Incoming traffic shaping profile. This attribute must reference one of the following datasources:firewall.shaping-profile.profile-name .
- 
ingress_spillover_threshold - Ingress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
- 
interface - Interface name. This attribute must reference one of the following datasources:system.interface.name .
- 
internal - Implicitly created.
- 
ip - Interface IPv4 address and subnet mask, syntax: X.X.X.X/24.
- 
ip_managed_by_fortiipam - Enable/disable automatic IP address assignment of this interface by FortiIPAM. Valid values:enabledisable .
- 
ipmac - Enable/disable IP/MAC binding. Valid values:enabledisable .
- 
ips_sniffer_mode - Enable/disable the use of this interface as a one-armed sniffer. Valid values:enabledisable .
- 
ipunnumbered - Unnumbered IP used for PPPoE interfaces for which no unique local address is provided.
- 
l2forward - Enable/disable l2 forwarding. Valid values:enabledisable .
- 
lacp_ha_secondary - LACP HA secondary member. Valid values:enabledisable .
- 
lacp_ha_slave - LACP HA slave. Valid values:enabledisable .
- 
lacp_mode - LACP mode. Valid values:staticpassiveactive .
- 
lacp_speed - How often the interface sends LACP messages. Valid values:slowfast .
- 
lcp_echo_interval - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
- 
lcp_max_echo_fails - Maximum missed LCP echo messages before disconnect.
- 
link_up_delay - Number of milliseconds to wait before considering a link is up.
- 
lldp_network_policy - LLDP-MED network policy profile. This attribute must reference one of the following datasources:system.lldp.network-policy.name .
- 
lldp_reception - Enable/disable Link Layer Discovery Protocol (LLDP) reception. Valid values:enabledisablevdom .
- 
lldp_transmission - Enable/disable Link Layer Discovery Protocol (LLDP) transmission. Valid values:enabledisablevdom .
- 
macaddr - Change the interface's MAC address.
- 
managed_subnetwork_size - Number of IP addresses to be allocated by FortiIPAM and used by this FortiGate unit's DHCP server settings. Valid values:32641282565121024204840968192163843276865536 .
- 
management_ip - High Availability in-band management IP address of this interface.
- 
measured_downstream_bandwidth - Measured downstream bandwidth (kbps).
- 
measured_upstream_bandwidth - Measured upstream bandwidth (kbps).
- 
mediatype - Select SFP media interface type Valid values:nonegmiisgmiisrlrcrsr2lr2cr2sr4lr4cr4sr8lr8cr8 .
- 
min_links - Minimum number of aggregated ports that must be up.
- 
min_links_down - Action to take when less than the configured minimum number of links are active. Valid values:operationaladministrative .
- 
mode - Addressing mode (static, DHCP, PPPoE). Valid values:staticdhcppppoe .
- 
monitor_bandwidth - Enable monitoring bandwidth on this interface. Valid values:enabledisable .
- 
mtu - MTU value for this interface.
- 
mtu_override - Enable to set a custom MTU for this interface. Valid values:enabledisable .
- 
name - Name.
- 
ndiscforward - Enable/disable NDISC forwarding. Valid values:enabledisable .
- 
netbios_forward - Enable/disable NETBIOS forwarding. Valid values:disableenable .
- 
netflow_sampler - Enable/disable NetFlow on this interface and set the data that NetFlow collects (rx, tx, or both). Valid values:disabletxrxboth .
- 
np_qos_profile - NP QoS profile ID.
- 
outbandwidth - Bandwidth limit for outgoing traffic (0 - 100000000 kbps).
- 
padt_retry_timeout - PPPoE Active Discovery Terminate (PADT) used to terminate sessions after an idle time.
- 
password - PPPoE account's password.
- 
ping_serv_status - PING server status.
- 
polling_interval - sFlow polling interval in seconds (1 - 255).
- 
pppoe_unnumbered_negotiate - Enable/disable PPPoE unnumbered negotiation. Valid values:enabledisable .
- 
pptp_auth_type - PPTP authentication type. Valid values:autopapchapmschapv1mschapv2 .
- 
pptp_client - Enable/disable PPTP client. Valid values:enabledisable .
- 
pptp_password - PPTP password.
- 
pptp_server_ip - PPTP server IP address.
- 
pptp_timeout - Idle timer in minutes (0 for disabled).
- 
pptp_user - PPTP user name.
- 
preserve_session_route - Enable/disable preservation of session route when dirty. Valid values:enabledisable .
- 
priority - Priority of learned routes.
- 
priority_override - Enable/disable fail back to higher priority port once recovered. Valid values:enabledisable .
- 
proxy_captive_portal - Enable/disable proxy captive portal on this interface. Valid values:enabledisable .
- 
reachable_time - IPv4 reachable time in milliseconds (30000 - 3600000, default = 30000).
- 
redundant_interface - Redundant interface.
- 
remote_ip - Remote IP address of tunnel.
- 
replacemsg_override_group - Replacement message override group.
- 
ring_rx - RX ring size.
- 
ring_tx - TX ring size.
- 
role - Interface role. Valid values:lanwandmzundefined .
- 
sample_direction - Data that NetFlow collects (rx, tx, or both). Valid values:txrxboth .
- 
sample_rate - sFlow sample rate (10 - 99999).
- 
secondary_ip - Enable/disable adding a secondary IP to this interface. Valid values:enabledisable .
- 
security_8021x_dynamic_vlan_id - VLAN ID for virtual switch.
- 
security_8021x_master - 802.1X master virtual-switch.
- 
security_8021x_mode - 802.1X mode. Valid values:defaultdynamic-vlanfallbackslave .
- 
security_exempt_list - Name of security-exempt-list.
- 
security_external_logout - URL of external authentication logout server.
- 
security_external_web - URL of external authentication web server.
- 
security_mac_auth_bypass - Enable/disable MAC authentication bypass. Valid values:mac-auth-onlyenabledisable .
- 
security_mode - Turn on captive portal authentication for this interface. Valid values:nonecaptive-portal802.1X .
- 
security_redirect_url - URL redirection after disclaimer/authentication.
- 
service_name - PPPoE service name.
- 
sflow_sampler - Enable/disable sFlow on this interface. Valid values:enabledisable .
- 
snmp_index - Permanent SNMP Index of the interface.
- 
speed - Interface speed. The default setting and the options available depend on the interface hardware. Valid values:auto10full10half100full100half100auto1000full1000auto2500auto5000auto10000full10000auto25000full25000auto40000full40000auto50000full50000auto100Gfull100Gauto200Gfull200Gauto400Gfull400Gauto .
- 
spillover_threshold - Egress Spillover threshold (0 - 16776000 kbps), 0 means unlimited.
- 
src_check - Enable/disable source IP check. Valid values:enabledisable .
- 
status - Bring the interface up or shut the interface down. Valid values:updown .
- 
stp - Enable/disable STP. Valid values:disableenable .
- 
