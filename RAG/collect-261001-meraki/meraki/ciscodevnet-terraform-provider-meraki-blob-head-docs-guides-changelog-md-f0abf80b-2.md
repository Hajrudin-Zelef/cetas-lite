---
id: collect-261001-meraki/meraki/ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b-2
title: "ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["licenses", "voice"]
source: docs/RAG/collect-261001-meraki/ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b.md
source_anchor: ""
source_lines: [64, 155]
sha256: 436be3fa79433d592575ff27835f0ad8a8fd3c486a9efe27683b68331916c74c
---

# ciscodevnet-terraform-provider-meraki-blob-head-docs-guides-changelog-md-f0abf80b

- Add local_auth_fallback_cache_timeout ,local_auth_fallback_enabled andlocal_auth_fallback_server_ca_certificate_contents attributes tomeraki_wireless_ssid resources and data sources
- Add meraki_wireless_zigbee_device resource and data sources
- Add meraki_wireless_zigbee_door_lock resource and data sources
- Add meraki_wireless_zigbee resource and data sources
- Add meraki_devices bulk resource
- Add meraki_switch_link_aggregations bulk resource
- Add meraki_organization_extensions_thousand_eyes_network resource and data sources
- Reset adaptive_policy_group_id attribute ofmeraki_switch_ports resource when destroying the resource
- Add support for importing specific item IDs of bulk resources
- Make force_delete attribute ofmeraki_network_group_policies resource "importable"
- Fix issue with attempting to delete non-existent items from bulk resources
- Add support for more than 2 WAN links to meraki_appliance_traffic_shaping_uplink_selection resource
- Add multi_exit_discriminator ,weight andpath_prepend attributes tomeraki_appliance_vpn_bgp resource
- Add support for custom-role access privilege tomeraki_organization_saml_role resource
- Add vrf_name attribute tomeraki_switch_routing_multicast_rendezvous_point resource
- Fix bulk resource import of item IDs
- Fix documentation categories for bulk resources
- Add meraki_switch_ports resource
- Add meraki_appliance_ports resource
- Add meraki_appliance_ssids resource
- Add meraki_organization_licenses resource
- Add meraki_sensor_mqtt_brokers resource
- Add meraki_wireless_ssids resource
- Add meraki_organization_policy_objects resource
- Add meraki_organization_policy_object_groups resource
- Add meraki_switch_routing_interfaces resource
- Add meraki_switch_stack_routing_interfaces resource
- Add meraki_sm_admin_roles resource
- Add meraki_organization_saml_idps resource
- Add meraki_organization_adaptive_policies resource
- Add meraki_organization_adaptive_policy_groups resource
- Add meraki_wireless_ssid_identity_psks resource
- Add meraki_wireless_rf_profiles resource
- Add meraki_network_webhook_payload_templates resource
- Add meraki_switch_stack_routing_static_routes resource
- Add meraki_switch_routing_multicast_rendezvous_points resource
- Add meraki_switch_access_policies resource
- Add meraki_sensor_alerts_profiles resource
- Add meraki_network_meraki_auth_users resource
- Add meraki_network_group_policies resource
- Add meraki_appliance_vlans resource
- Add meraki_appliance_traffic_shaping_custom_performance_classes resource
- Add meraki_appliance_rf_profiles resource
- Add meraki_appliance_prefix_delegated_statics resource
- Add meraki_switch_routing_static_routes resource
- Change items attribute of bulk data sources from typeList toSet
- Add requests_per_second provider configuration attribute
- Add group_active_active_tunnel ,group_failover_direct_to_internet ,group_number ,is_route_based ,network_ids ,peer_id ,priority_in_group ,sla_policy_id andebgp_neighbor_* attributes tomeraki_appliance_third_party_vpn_peers resource and data source
- Make name attribute oforganization_brnading_policy resource mandatory
- Add guest_group_policy_id ,guest_sgt_id ,radius_authentication_mode ,radius_critical_auth_data_group_policy_id ,radius_critical_auth_data_sgt_id ,radius_critical_auth_voice_group_policy_id ,radius_critical_auth_voice_sgt_id ,radius_failed_auth_group_policy_id ,radius_failed_auth_sgt_id andradius_pre_authentication_group_policy_id attributes tomeraki_switch_access_policy resource and data sources
- Add meraki_appliance_vpn_site_to_site_ipsec_peers_slas resource and data source
- Fix idempotency issue with move_map_marker attribute ofmeraki_device resource, link
- Add unsupported radius_das_clients_ips andradius_das_clients_shared_secret attributes tomeraki_wireless_ssid resource and data source, link
- Delete lists when destroying meraki_appliance_connectivity_monitoring_destinations ,meraki_appliance_firewall_multicast_forwarding ,meraki_appliance_organization_security_intrusion ,meraki_appliance_sdwan_internet_policies ,meraki_appliance_third_party_vpn_peers ,meraki_appliance_traffic_shaping_uplink_selection ,meraki_pliance_traffic_shaping_vpn_exclusions ,meraki_cellular_gateway_connectivity_monitoring_destinations ,meraki_network_syslog_servers ,meraki_switch_access_control_lists andmeraki_switch_dscp_to_cos_mappings resources
- Reset access_policy_type ,adaptive_policy_group_id andprofile_enabled attributes ofmeraki_switch_port resource when destroying the resource
- Fix API format of definitions[].value attribute ofmeraki_appliance_traffic_shaping_rules resource, link
- Fix API format of l7_firewall_rules[].value attribute ofmeraki_network_group_policy resource, link
- Fix API format of traffic_shaping_rules[].definitions[].value attribute ofmeraki_network_group_policy resource, link
- Fix API format of rules[].value attribute ofmeraki_wireless_ssid_l7_firewall_rules resource, link
- Fix meraki_appliance_sdwan_internet_policies deletion doing nothing, link
- Handle HTTP error code 400 correctly when trying to read non-existent resources
- Make force_delete attribute ofmeraki_network_group_policy resource "importable"
- Add meraki_network_alerts_settings resource and data source
- Fix issue with not handling paginated responses correctly, link
- Add force_delete attribute tomeraki_network_group_policy resource, link
- Add details_by_device attribute tomeraki_network_device_claim resource
- Add mac_whitelist_limit attribute tomeraki_switch_port resource and data sources
- Add adaptive_policy_group_id attribute tomeraki_wireless_ssid resource and data source
- Add meraki_wireless_location_scanning resource and data source
- Add meraki_wireless_location_scanning_receiver resource and data sources
- Fix idempotency issue with syslog_default_rule attribute ofmeraki_appliance_vpn_firewall_rules resource, link
- Fix idempotency issue with syslog_default_rule attribute ofmeraki_appliance_l3_firewall_rules resource, link
- Fix idempotency issue with syslog_default_rule attribute ofmeraki_appliance_inbound_firewall_rules resource, link
- Add meraki_network_device_claim_vmx resource
- Add meraki_appliance_dns_local_profile_assignments resource
- Add meraki_appliance_dns_local_record resource and data sources
- Add meraki_appliance_dns_split_profile resource and data sources
- Add meraki_appliance_dns_split_profile_assignments resource
- Add meraki_sm_admin_role resource and data sources
- Add meraki_sm_target_group resource and data sources
- Add warning if name query option of data source is being used and multiple objects with the same name exist
- BREAKING CHANGE: Rename switch_port_ids attribute ofmeraki_switch_organization_ports_profiles_automation resource and data source toport_ids
- Add MV84X settings to meraki_camera_quality_retention_profile resource and data source
- Add minimum_password_length attribute tomeraki_organization_login_security resource and data source
- Add include_sensor_url andmessage attributes tomeraki_sensor_alerts_profile resource and data source
- Make name attribute ofmeraki_switch_routing_interface resource and data source required
- Add mode attribute tomeraki_wireless_network_electronic_shelf_label resource and data source
- Change type of floor_number attribute ofmeraki_network_floor_plan resource and data source frominteger tofloat
- Add meraki_wireless_ssid_firewall_isolation_allowlist_entry resource and data sources
- Remove encryption_enabled andencryption_certificate_id attributes frommeraki_network_syslog_servers resource and data source
- Add meraki_appliance_firewall_multicast_forwarding resource and data source
- Add meraki_appliance_dns_local_profile resource and data sources
