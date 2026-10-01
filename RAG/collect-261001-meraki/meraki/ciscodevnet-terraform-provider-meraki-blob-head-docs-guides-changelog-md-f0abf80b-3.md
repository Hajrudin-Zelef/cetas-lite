---
id: collect-261001-meraki/meraki/ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b-3
title: "ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "licenses"]
source: docs/RAG/collect-261001-meraki/ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b.md
source_anchor: ""
source_lines: [156, 253]
sha256: 627c022f9ac2bda6778e4604e667e9e7825876399d976a65ab96c3eff3dd2b57
---

# ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b

- Do not import default rule for meraki_appliance_cellular_firewall_rules ,meraki_appliance_inbound_cellular_firewall_rules ,meraki_appliance_inbound_firewall_rules ,meraki_appliance_l3_firewall_rules ,meraki_appliance_vpn_firewall_rules andmeraki_switch_access_control_lists resource
- Do not import default and local LAN access rules for meraki_wireless_ssid_l3_firewall_rules resource
- Add value_countries attribute tomeraki_appliance_l7_firewall_rules resource and data source
- BREAKING CHANGE: Remove application_value attribute frommeraki_wireless_ssid_traffic_shaping_rules resource and data source
- Add subnet_nat_is_allowed ,nat_enabled andnat_remote_subnet attributes tomeraki_appliance_site_to_site_vpn resource and data source
- Add video_settings_mv53_x_quality andvideo_settings_mv53_x_resolution attributes tomeraki_camera_quality_retention_profile resource and data sources
- Add self_registration_authorization_type andself_registration_enabled attributes tomeraki_wireless_ssid_splash_settings resource and data source
- Add application_value attribute tomeraki_wireless_ssid_traffic_shaping_rules resource and data source
- Add meraki_organization_auth_radius_server resource and data sources
- Add meraki_camera_custom_analytics resource and data source
- Add meraki_camera_quality_retention resource and data source
- Add meraki_camera_quality_retention_profile resource and data sources
- Add meraki_camera_role resource and data sources
- Add meraki_camera_sense resource and data source
- Add meraki_camera_video_settings resource and data source
- Add meraki_camera_wireless_profile resource and data sources
- Add meraki_camera_device_wireless_profiles resource and data source
- Add meraki_insight_monitored_media_server resource and data source
- Add meraki_sensor_alerts_profile resource and data source
- Add meraki_sensor_mqtt_broker resource and data sources
- Add meraki_sensor_relationships resource and data source
- Add meraki_sensor_network_relationships data source
- Delete all rules when destroying meraki_appliance_cellular_firewall_rules ,meraki_appliance_inbound_cellular_firewall_rules ,meraki_appliance_inbound_firewall_rules ,meraki_appliance_l3_firewall_rules ,meraki_appliance_l7_firewall_rules ,meraki_appliance_one_to_many_nat_rules ,meraki_appliance_one_to_one_nat_rules ,meraki_appliance_port_forwarding_rules ,meraki_appliance_traffic_shaping_rules ,meraki_appliance_vpn_firewall_rules ,meraki_wireless_ssid_l3_firewall_rules ,meraki_wireless_ssid_l7_firewall_rules ,meraki_wireless_ssid_traffic_shaping_rules resources
- Add floor_number attribute tomeraki_network_floor_plan resource and data sources
- Add local_status_page_authentication_username attribute tomeraki_network_settings resource and data source
- Add encryption attributes to meraki_network_syslog_servers resource and data source
- Add log andtcp_established attributes tomeraki_organization_adaptive_policy_acl resource and data sources
- Add stackwise_virtual_is_dual_active_detector andstackwise_virtual_is_stack_wise_virtual_link attributes tomeraki_switch_ports data source
- Add radius_radsec_tls_tunnel_timeout attribute tomeraki_wireless_ssid resource and data source
- Add meraki_cellular_gateway_connectivity_monitoring_destinations resource and data sources
- Add meraki_cellular_gateway_dhcp resource and data sources
- Add meraki_cellular_gateway_lan resource and data sources
- Add meraki_cellular_gateway_port_forwarding_rules resource and data sources
- Add meraki_cellular_gateway_subnet_pool resource and data sources
- Add meraki_cellular_gateway_uplink resource and data sources
- Add provider configuration to define a list of HTTP error codes to retry on
- Add ip_version attribute tomeraki_wireless_ssid_l3_firewall_rules resource and data source
- Add meraki_appliance_vmx_authentication_token resource
- Configure default settings when deleting wireless SSID using meraki_wireless_ssid resource
- Add fixed_ip_assignments attribute tomeraki_appliance_vlan resource and data sources
- Add public_hostname attribute tomeraki_appliance_third_party_vpn_peers resource and data source
- Add treat_these_traffic_types_as_one_threshold attribute tomeraki_switch_storm_control resource and data source
- Add dhcp_boot_filename ,dhcp_boot_next_server ,dns_nameservers ,vpn_nat_subnet ,dhcp_relay_server_ips andreserved_ip_ranges attributes tomeraki_appliance_vlan resource and data source
- Add meraki_organization_early_access_features_opt_in resource and data sources
- Add meraki_switch_organization_ports_profile resource and data sources
- Add meraki_switch_organization_ports_profiles_automation resource and data sources
- Add meraki_appliance_vlan_dhcp resource
- Add meraki_organization_license resource and data source
- Add meraki_network_firmware_upgrades resource and data source
- Add meraki_organization_licenses data source
- Add meraki_appliance_ports data source
- Add meraki_appliance_ssids data source
- Add meraki_switch_ports data source
- Add meraki_wireless_ssids data source
- Add meraki_organization_early_access_features data source
- Add meraki_organization_branding_policy resource and data sources
- Add meraki_organization_branding_policies_priorities resource and data source
- Add meraki_network_client_splash_authorization_status resource and data source
- Add meraki_network_devices data source
- Add meraki_organization_devices data source
- Add meraki_organization_firmware_upgrades data source
- Add meraki_organization_inventory_devices data source
- Add meraki_network_policies_by_client data source
- Add meraki_network_vlan_profile_assignments_by_device data source
- Add meraki_appliance_firewalled_service resource and data source
- Add meraki_appliance_inbound_cellular_firewall_rules resource and data source
- Add meraki_appliance_inbound_firewall_rules resource and data source
- Add meraki_appliance_l3_firewall_rules resource and data source
- Add meraki_appliance_l7_firewall_rules resource and data source
- Add meraki_appliance_one_to_many_nat_rules resource and data source
- Add meraki_appliance_one_to_one_nat_rules resource and data source
- Add meraki_appliance_port_forwarding_rules resource and data source
- Add meraki_appliance_firewall_settings resource and data source
- Add meraki_appliance_port resource and data source
- Add meraki_appliance_vlans_settings resource and data source
- Add meraki_appliance_prefix_delegated_static resource and data source
- Add meraki_appliance_prefix_delegated_statics data source
- Add meraki_appliance_radio_settings resource and data source
- Add meraki_appliance_rf_profile resource and data source
- Add meraki_appliance_rf_profiles data source
- BREAKING CHANGE: Rename per_ssid_settingsXX_* attributes ofmeraki_wireless_rf_profile resource and data sources toper_ssid_settings_XX_*
- Add meraki_appliance_sdwan_internet_policies resource
- Add meraki_appliance_network_security_intrusion resource and data source
- Add meraki_appliance_organization_security_intrusion resource and data source
- Add meraki_appliance_security_malware resource and data source
- Add meraki_appliance_settings resource and data source
- Add meraki_appliance_single_lan resource and data source
- Add meraki_appliance_ssid resource and data source
- Add meraki_appliance_static_route resource and data source
- Add meraki_appliance_static_routes data source
- Add meraki_appliance_vlan resource and data source
- Add meraki_appliance_vlans data source
- Add meraki_appliance_traffic_shaping resource and data source
- Add meraki_appliance_traffic_shaping_custom_performance_class resource and data source
- Add meraki_appliance_traffic_shaping_custom_performance_classes data source
- Add meraki_appliance_traffic_shaping_rules resource and data source
- Add meraki_appliance_traffic_shaping_uplink_bandwidth resource and data source
- Add meraki_appliance_traffic_shaping_uplink_selection resource and data source
