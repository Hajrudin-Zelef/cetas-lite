---
id: collect-261001-meraki/meraki/ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b-1
title: "ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-meraki/ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b.md
source_anchor: ""
source_lines: [1, 63]
sha256: e2d51c254cae9a96f353125af8e363d4ae455666c40ed8956f4b173618d2162a
---

# ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b

| subcategory | Guides | 
|---|---|
| page_title | Changelog | 
| description | Changelog | 
- Add meraki_generate_appliance_vmx_authentication_token action
- Add meraki_reboot_device action
- Add meraki_blink_device_leds action
- Fix issue with spare_serial attribute ofmeraki_appliance_warm_spare resource, link
- Add enabled attribute tomeraki_appliance_static_route resource and data source, link
- Add eox attributes (eox_end_of_sale_at ,eox_end_of_support_at ,eox_status ) tomeraki_organization_inventory_devices data source
- Add multicast_to_unicast_conversion_enabled attribute tomeraki_wireless_settings resource and data source
- Add meraki_network_wireless_radio_rrm resource and data source
- Fix idempotency issue with port_ids attribute ofmeraki_switch_organization_ports_profiles_automation resource, link
- Fix fallback_profile_id andfallback_profile_name attributes ofmeraki_switch_organization_ports_profiles_automation resource not being cleared via the API when removed from config, link
- Add meraki_network_firmware_upgrades_rollback resource
- Add is_upgrade_available ,current_version ,last_upgrade , extranext_upgrade.to_version , andavailable_versions attributes tomeraki_network_firmware_upgrades data source
- Only send products.*.next_upgrade.* attributes to the API when changed for a given product inmeraki_network_firmware_upgrades resource, so other products' previously-scheduled upgrades are preserved in the config
- Preserve attributes that the API explicitly returned as null in the saved initial state used byrestore_original_state_on_destroy , so they are restored on destroy instead of being silently dropped
- Fix write-only attributes (_wo ) requiring the base sensitive attribute to also be set inmeraki_network_snmp (users[*].passphrase ),meraki_appliance_third_party_vpn_peers (peers[*].secret ),meraki_network_meraki_auth_user (password ),meraki_organization_auth_radius_server (secret ), andmeraki_wireless_location_scanning_receiver (shared_secret ), link
- Fix "Missing Resource Identity After Read" provider error for resources with custom read implementations (meraki_appliance_firewall_multicast_forwarding ,meraki_appliance_traffic_shaping_vpn_exclusions ,meraki_wireless_air_marshal_rule ,meraki_wireless_air_marshal_settings ,meraki_wireless_location_scanning ,meraki_wireless_ssid_open_roaming ,meraki_wireless_zigbee , and others) when deleted out-of-band
- Add wan1_vrf_name andwan2_vrf_name attributes tomeraki_device_management_interface resource and data source
- Fix "Missing Resource Identity After Read" provider error when a resource was first created with Terraform < 1.12 and has since been deleted out-of-band
- Fix issue with stack_ids attribute ofmeraki_network_vlan_profile_assignment resource, link
- Fix idempotency issue with vpn_traffic_uplink_preferences[].traffic_filters ,wan_traffic_uplink_preferences[].traffic_filters attributes ofmeraki_appliance_traffic_shaping_uplink_selection resource, link
- Fix "Missing Resource Identity After Update" provider error when updating resources with Terraform versions not supporting resource identity (< 1.12)
- EXPERIMENTAL: Add restore_original_state_on_destroy provider attribute to opt in to restoring the original API state of singleton resources on destroy. When enabled, the provider captures the initial state during resource creation and restores it when the resource is destroyed. This feature is experimental and may change in future releases. See the Restore State on Destroy guide for details and limitations.
- Add resource identity support for Terraform 1.12+ import blocks, with backward compatibility for older Terraform versions
- Add meraki_organization_integrations_xdr_networks resource and data source
- Add meraki_network_vlan_profile_assignment resource and data source
- Add write-only attribute support (_wo /_wo_version siblings) for sensitive string attributes, compatible with Terraform 1.11+
- Add support for policy objects (OBJ(<id>) ) and policy object groups (GRP(<id>) ) insrc_cidr anddest_cidr fields ofmeraki_appliance_l3_firewall_rules andmeraki_appliance_cellular_firewall_rules resources.link
- Add candidate_uplink_v4 ,is_switch_default_gateway ,static_v4_dns1 ,static_v4_dns2 ,uplink_v4 ,uplink_v6 ,ipv6_candidate_uplink ,ipv6_is_switch_default_gateway ,ipv6_static_v6_dns1 ,ipv6_static_v6_dns2 attributes tomeraki_switch_routing_interface resources and data sources
- Add candidate_uplink_v4 ,is_switch_default_gateway ,static_v4_dns1 ,static_v4_dns2 ,uplink_v4 ,uplink_v6 ,ipv6_candidate_uplink ,ipv6_is_switch_default_gateway ,ipv6_static_v6_dns1 ,ipv6_static_v6_dns2 attributes tomeraki_switch_stack_routing_interface resources and data sources
- Add authentication_host_mode attribute tomeraki_switch_organization_ports_profile resource and data sources
- Add items.id attribute tomeraki_appliance_vpn_site_to_site_ipsec_peers_slas resource and data source
- Mark sensitive attributes (passwords, secrets, PSKs, passphrases, tokens, SNMP community strings) to prevent exposure in plan output and logs
- Add Software Bill of Materials (SBOM) generation in SPDX and CycloneDX formats during releases
- Add sso_login_url attribute tomeraki_organization_saml_idp resources and data sources
- Add uplink_selection_candidates anduplink_selection_failback_enabled attributes tomeraki_switch_settings resource and data source
- Add meraki_wireless_ssid_open_roaming resource and data source
- Add sp_initiated_idp_id andsp_initiated_subdomain attributes tomeraki_organization_saml resource and data source
- Add stp_port_fast_trunk attribute tomeraki_switch_port resources and data sources
- Add access_control_* attributes tomeraki_wireless_ssids data source
- Add ip_version attribute tomeraki_appliance_third_party_vpn_peers resource and data source, link
- Enhance handling of DNS split profile assignments by adding checks for empty responses and refining matching logic for assignment IDs during updates - Issue #131, link
- Add undocumented ipv6_prefix_assignments[].disabled attribute tomeraki_appliance_single_vlan resource, link
- Allow pushing empty lists, link
- Add switchport option toaccess attribute ofmeraki_organization_saml_role resources, link
- Add high_speed_enabled attribute tomeraki_switch_port resources and data sources
- Add support for both value toip_version attribute ofwireless_ssid_l3_firewall_rules resource
- Fix issue with configuring DHCP relay (dhcp_handling ,dhcp_relay_server_ips atributes) inmeraki_appliance_vlan_dhcp resource, link
- Add fixed_ip_assignments andvpn_nat_subnet attributes tomeraki_appliance_vlan_dhcp resource, link
- Add mandatory undocumented ipv6_prefix_assignments[].disabled attribute tomeraki_appliance_vlan resource, link
- Add adaptive_policy_voice_group_id attribute tomeraki_switch_organization_ports_profile resource and data sources
- Add module_serial andmodule_slot attributes tomeraki_switch_ports data source
- Add radius_accounting_start_delay attribute tomeraki_wireless_ssid resources and data source
- Add meraki_wireless_mqtt_settings resource and data source
- Apply changes to default RF profiles during creation, link
- Fix issue with is_indoor_default andis_outdoor_default attributes ofmeraki_wireless_rf_profile not being applied correctly
- Add is_indoor_default andis_outdoor_default attributes tomeraki_wireless_rf_profile resource and data source
- Add privacy_link attribute tomeraki_organization_early_access_features data source
- Add mode ,ospf_settings_network_type ,switch_port_id andvrf_name attributes tomeraki_switch_routing_interface andmeraki_switch_stack_routing_interface resources and data sources
- Add vrf_leak_route_to_default_vrf andvrf_name attributes tomeraki_switch_routing_static_route andmeraki_switch_stack_routing_static_route resources and data sources
