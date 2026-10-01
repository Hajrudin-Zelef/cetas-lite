---
id: collect-261001-meraki/meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9-6
title: "fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-meraki/fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9.md
source_anchor: ""
source_lines: [492, 586]
sha256: 4b05f282c4b005aa2e73ebeb8ae8e02219e2c9f6caf3ff3263585be67f751ccf
---

# fr-pp-integrations-plugin-packs-procedures-network-cisco-meraki-restapi-0563f0f9

| --filter-organization-id | Filter devices by organization ID (can be a regexp). | 
| --filter-organization-name | Filter devices by organization name (can be a regexp). | 
| --filter-tags | Filter devices by tags (can be a regexp). | 
| --add-switch-ports | Add switch port statuses and traffic. | 
| --filter-switch-port | Filter switch port (can be a regexp). | 
| --skip-clients | Don't monitor clients traffic on device. | 
| --skip-performance | Don't monitor appliance performance score. | 
| --skip-connections | Don't monitor connection stats. | 
| --skip-traffic-disconnect-port | Skip port traffic counters if port status is disconnected. | 
| --unknown-status | Define the conditions to match for the status to be UNKNOWN. You can use the following variables: %{status}, %{display} | 
| --warning-status | Threshold. | 
| --critical-status | Threshold. | 
| --unknown-link-status | Define the conditions to match for the status to be UNKNOWN. You can use the following variables: %{link_status}, %{display} | 
| --warning-link-status | Threshold. | 
| --critical-link-status | Threshold. | 
| --unknown-port-status | Define the conditions to match for the status to be UNKNOWN. You can use the following variables: %{port_status}, %{port_enabled}, %{display} | 
| --warning-port-status | Threshold. | 
| --critical-port-status | Threshold. | 
| --warning-connections-assoc | Threshold. | 
| --critical-connections-assoc | Threshold. | 
| --warning-connections-auth | Threshold. | 
| --critical-connections-auth | Threshold. | 
| --warning-connections-dhcp | Threshold. | 
| --critical-connections-dhcp | Threshold. | 
| --warning-connections-dns | Threshold. | 
| --critical-connections-dns | Threshold. | 
| --warning-connections-success | Threshold. | 
| --critical-connections-success | Threshold. | 
| --warning-link-latency | Threshold in milliseconds. | 
| --critical-link-latency | Threshold in milliseconds. | 
| --warning-link-loss | Threshold in percentage. | 
| --critical-link-loss | Threshold in percentage. | 
| --warning-links-ineffective | Threshold. | 
| --critical-links-ineffective | Threshold. | 
| --warning-load | Threshold. | 
| --critical-load | Threshold. | 
| --warning-port-traffic-in | Threshold in b/s. | 
| --critical-port-traffic-in | Threshold in b/s. | 
| --warning-port-traffic-out | Threshold in b/s. | 
| --critical-port-traffic-out | Threshold in b/s. | 
| --warning-total-alerting | Threshold. | 
| --critical-total-alerting | Threshold. | 
| --warning-total-offline | Threshold. | 
| --critical-total-offline | Threshold. | 
| --warning-total-offline-prct | Threshold in percentage. | 
| --critical-total-offline-prct | Threshold in percentage. | 
| --warning-total-online | Threshold. | 
| --critical-total-online | Threshold. | 
| --warning-total-online-prct | Threshold in percentage. | 
| --critical-total-online-prct | Threshold in percentage. | 
| --warning-traffic-in | Threshold in b/s. | 
| --critical-traffic-in | Threshold in b/s. | 
| --warning-traffic-out | Threshold in b/s. | 
| --critical-traffic-out | Threshold in b/s. | 
| Option | Description | 
|---|---|
| --filter-network-name | Filter network name (can be a regexp). | 
| --filter-organization-id | Filter networks by organization ID (can be a regexp). | 
| --filter-organization-name | Filter networks by organization name (can be a regexp). | 
| --warning-connections-assoc | Threshold. | 
| --critical-connections-assoc | Threshold. | 
| --warning-connections-auth | Threshold. | 
| --critical-connections-auth | Threshold. | 
| --warning-connections-dhcp | Threshold. | 
| --critical-connections-dhcp | Threshold. | 
| --warning-connections-dns | Threshold. | 
| --critical-connections-dns | Threshold. | 
| --warning-connections-success | Threshold. | 
| --critical-connections-success | Threshold. | 
| --warning-traffic-in | Threshold in b/s. | 
| --critical-traffic-in | Threshold in b/s. | 
| --warning-traffic-out | Threshold in b/s. | 
| --critical-traffic-out | Threshold in b/s. | 
| Option | Description | 
|---|---|
| --filter-network-name | Filter VPN tunnels by network name (can be a regexp). | 
| --filter-organization-id | Filter VPN tunnels by organization ID (can be a regexp). | 
| --filter-organization-name | Filter VPN tunnels by organization name (can be a regexp). | 
| --filter-device-serial | Filter VPN tunnels by device serial (can be a regexp). | 
| --filter-vpn-type | Filter VPN tunnels by VPN type (can be a regexp). | 
| --filter-vpn-name | Filter VPN tunnels by VPN name (can be a regexp). | 
| --unknown-device-status | Define the conditions to match for the status to be UNKNOWN (default: '%{deviceStatus} =~ /offline/i'). You can use the following variables: %{deviceStatus}, %{deviceSerial}, %{deviceMode} | 
| --warning-device-status | Threshold. | 
| --critical-device-status | Threshold. | 
| --unknown-vpn-status | Define the conditions to match for the status to be UNKNOWN. You can use the following variables: %{vpnStatus}, %{vpnName}, %{vpnType}, %{deviceStatus}, %{deviceSerial} | 
| --warning-vpn-status | Threshold. | 
| --critical-vpn-status | Threshold. | 
| --warning-total-unreachable | Threshold. | 
| --critical-total-unreachable | Threshold. | 
Pour un mode, la liste de toutes les options disponibles et leur signification peut être
affichée en ajoutant le paramètre --help à la commande :
/usr/lib/centreon/plugins/centreon_cisco_meraki_restapi.pl \
	--plugin=network::cisco::meraki::cloudcontroller::restapi::plugin \
	--mode=networks \
	--help
