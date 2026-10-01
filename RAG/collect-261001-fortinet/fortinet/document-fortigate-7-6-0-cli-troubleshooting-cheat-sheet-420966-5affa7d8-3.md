---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8-3
title: "document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "parameters"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8.md
source_anchor: ""
source_lines: [190, 272]
sha256: 8d1eba118e8b572dbe7083c76acf6056890644530264bec802f9b75318507f62
---

# document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8

| get router info multicast pim sparse-mode rp-mapping | Show RP to group mapping information. | 
| get router info multicast pim sparse-mode table | Show sparse-mode routing table. | 
| diagnose ip router pim-sm events enable diagnose ip router pim-sm all enable diagnose ip router pim-sm level info diagnose debug enable | Start real-time debugging of PIM sparse mode. | 
SD-WAN
| Command | Description | 
|---|---|
| diagnose sys sdwan health-check status | Show SD-WAN health check statistics. | 
| diagnose sys sdwan service4 diagnose sys sdwan service6 | Show SD-WAN rules in control plane. | 
| diagnose sys sdwan member | Show SD-WAN members. | 
| diagnose firewall proute list | Show SDWAN rule and policy routes in the data plane. | 
| diagnose sys link-monitor status diagnose sys link-monitor interface <interface> | Show link monitoring statistics. | 
| diagnose debug application link-monitor -1 diagnose debug enable | Start real-time link monitor debugging. | 
| diagnose test application lnkmtd 1 diagnose test application lnkmtd 2 diagnose test application lnkmtd 3 | Show link monitoring statistics. | 
Authentication
| Command | Description | 
|---|---|
| diagnose firewall auth filter <filter> | Set the filter used to list entries. | 
| diagnose firewall auth list | List filtered, authenticated IPv4 users. | 
| diagnose wad user list | List current users authenticated by proxy (wad daemon). | 
| diagnose debug application fnbamd -1 diagnose debug application authd -1 diagnose debug enable | Start real-time debugging for remote and local authentication. | 
| diagnose test authserver <auth_protocol> <server_name> <user> <password> | Test authentication directly from the CLI. Caution: The password is visible in clear text; be careful when capture this command to a log file. | 
| diagnose test authserver ldap <server_name> <user> <password> | Test user authentication using an LDAP server. Caution: The password is visible in clear text; be careful when capture this command to a log file. | 
| diagnose test authserver radius <server_name> <auth_type> <user> <password> | Test user authentication using a Radius server. Caution: The password is visible in clear text; be careful when capture this command to a log file. | 
| diagnose debug fsso-polling detail diagnose debug fsso-polling summary | Show information about the polls from FortiGate to DC. | 
| diagnose debug fsso-polling user diagnose debug authd fsso list | Show FSSO logged on users when Fortigate polls the DC. | 
| diagnose debug application fssod -1 diagnose debug application smbcd -1 diagnose debug enable | Start real-time debugging when the FortiGate is used for FSSO polling. | 
| diagnose debug fsso-polling refresh-user execute fsso refresh | Refresh the current logged on FSSO users and refresh the list. Caution: This command can cause an outage, use it carefully. | 
| diagnose debug authd fsso server-status | Show current status of connection between FortiGate and the collector agent. | 
| diagnose debug application authd 8256 diagnose debug enable | Start real-time debugging for the connection between FortiGate and the collector agent. | 
| diagnose debug authd fsso refresh-logons | Resend the logged-on users list to FortiGate from the collector agent. | 
| diagnose debug application authd 8256 diagnose debug enable | Start real-time debugging for the connection between FortiGate and the collector agent. | 
| diagnose debug application samld -1 diagnose debug enable | Start real-time SAML debugging. | 
VPN
IPsec
| Command | Description | 
|---|---|
| diagnose vpn ike gateway list | Show IPsec phase 1 information. | 
| diagnose vpn tunnel list | Show IPsec phase 2 information. | 
| get vpn ipsec tunnel summary get vpn ipsec tunnel details | Show summary and detailed information about IPsec tunnels. | 
| diagnose vpn tunnel flush | Flush all Phase2 tunnel SAs (Security Associations). | 
| diagnose vpn tunnel flush <name> [name] | Flush one or more specific Phase2 tunnels by name. | 
| diagnose vpn ike gateway <clear \| flush> | Clear/flush IKE gateways (Phase1). Apply diagnose vpn ike gateway filter to filter on specific gateways. | 
| diagnose vpn ike gateway <clear \| flush> name <name> | Clear/flush a specific IKE gateway (Phase1) by name. | 
| diagnose vpn ike gateway filter | Use various options to filter the IKE gateways. | 
| diagnose vpn ipsec status | Show information about encryption counters. | 
| diagnose vpn ike log filter <filter> | Set a filter for IKE daemon debugs. | 
| diagnose debug application ike -1 diagnose debug enable | Start real-time debugging of IKE daemon with the filter set. | 
| diagnose vpn ike restart | Restart the IKE process. | 
| diagnose vpn ike counts diagnose vpn ike routes diagnose vpn ike errors diagnose vpn ike stats diagnose vpn ike status diagnose vpn ike crypto | Show other information, such as IKE counts, routes, errors, and statistics. | 
SSL VPN
|  | SSL VPN web mode has become Agentless VPN, and SSL VPN tunnel mode is no longer supported in 7.6.3 and later. Therefore, SSL VPN related debug commands may not work as expected. | 
| Command | Description | 
|---|---|
| diagnose vpn ssl debug-filter list | Show any filters that are set for SSL VPN debug. | 
| diagnose vpn ssl debug-filter clear | Clear any filters that are set for SSL VPN daemon debug. | 
| diagnose vpn ssl debug-filter <filter> | Set a filter for SSL VPN debugs. | 
| diagnose debug application sslvpn -1 diagnose debug enable | Start SSL VPN debugs for traffic that the filter is applied to. | 
| diagnose vpn ssl list get vpn ssl monitor execute vpn sslvpn list | Show the current SSL VPN sessions for both web and tunnel mode. | 
| diagnose vpn ssl statistics diagnose vpn ssl mux-stat | Show the SSL VPN statistics. | 
| execute vpn sslvpn list | Show all SSL VPN web and tunnel mode connections. | 
| execute vpn sslvpn del-tunnel | Disconnect the users from tunnel mode SSL VPN connection. | 
| execute vpn sslvpn del-web | Disconnect the users from web mode SSL VPN connection. | 
Managed devices
Managed FortiSwitches
|  | The successful execution of commands for managed FortiSwitches requires that the feature is available on the FortiSwitch device itself. See the FortiSwitchOS Feature Matrix. | 
|  | Enter ? to view additional options or parameters required to obtain the required information in thediagnose switch-controller switch-info commands. | 
| Command | Description | 
|---|---|
| diagnose switch-controller switch-info mac-table | Show managed FortiSwitch MAC address list. | 
| diagnose switch-controller switch-info port-stats | Show managed FortiSwitch port statistics. | 
| diagnose switch-controller switch-info trunk status | Show managed FortiSwitch trunk information. | 
| diagnose switch-controller switch-info mclag | Show MCLAG related information from FortiSwitch. | 
| diagnose switch-controller switch-info poe | Show POE-related information. | 
| diagnose switch-controller switch-info lldp | Show LLDP-related information. | 
| diagnose switch-controller switch-info port-properties | Show managed FortiSwitch port properties. | 
| diagnose switch-controller switch-info acl-counters | Show managed FortiSwitch port ACL counters information. | 
| diagnose switch-controller switch-info pdu-counters-list | Show managed FortiSwitch pdu-counters information. | 
| diagnose switch-controller switch-info flapguard | Show managed FortiSwitch flapguard information. | 
| diagnose switch-controller switch-info qos-stats | Show managed FortiSwitch QoS statistics. | 
| diagnose switch-controller switch-info modules | Show modules related information from FortiSwitch. | 
| diagnose switch-controller switch-info stp | Show managed FortiSwitch STP instance status. | 
| diagnose switch-controller switch-info bpdu-guard-status | Show managed FortiSwitch STP BPDU guard status. | 
| diagnose switch-controller switch-info igmp-snooping | Show managed FortiSwitch IGMP snooping information. | 
