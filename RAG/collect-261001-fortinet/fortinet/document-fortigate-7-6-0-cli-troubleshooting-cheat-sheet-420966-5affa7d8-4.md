---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8-4
title: "document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8.md
source_anchor: ""
source_lines: [273, 356]
sha256: 82885f16ce3a4eabfee0eb5343e4fa7a3168ce8c90917f3d0e765cb1797ff4c2
---

# document-fortigate-7-6-0-cli-troubleshooting-cheat-sheet-420966-5affa7d8

| diagnose switch-controller switch-info loop-guard | Show managed FortiSwitch loop-guard status. | 
| diagnose switch-controller switch-info dhcp-snooping | Show managed FortiSwitch DHCP snooping interface list. | 
| diagnose switch-controller switch-info arp-inspection | Show managed FortiSwitch ARP inspection interface list. | 
| diagnose switch-controller switch-info option82-mapping | Show managed FortiSwitch DHCP option 82 mapping information. | 
| diagnose switch-controller switch-info 802.1X | Show managed FortiSwitch port 802.1X status. | 
| diagnose switch-controller switch-info 802.1X-dacl | Show managed FortiSwitch port 802.1X dynamic ACL status. | 
| diagnose switch-controller switch-info mac-limit-violations | Show managed FortiSwitch violated MACs information. | 
| diagnose switch-controller switch-info flow-tracking | Show managed FortiSwitch flow information. | 
| diagnose switch-controller switch-info mirror | Show managed FortiSwitch mirror information. | 
| diagnose switch-controller switch-info ip-source-guard | Show managed FortiSwitch source guard information in hardware. | 
| diagnose switch-controller switch-info rpvst | Show managed FortiSwitch STP port information when inter-operating with rapid PVST network. | 
| execute switch-controller get-conn-status <FortiSwitch-SN> | Show FortiSwitch connection status. | 
| execute switch-controller get-physical-conn standard <FortiSwitch-SN> | Show FortiLink connectivity graph. | 
| execute switch-controller diagnose-connection <FortiSwitch-SN> | Show FortiSwitch connection diagnostics. | 
Managed FortiAPs
| Command | Description | 
|---|---|
| diagnose wireless-controller wlac -c wtp diagnose wireless-controller wlac -d wtp | Show information about the FortiAP devices. | 
| diagnose wireless-controller wlac -c sta diagnose wireless-controller wlac -d sta | Show information about the wireless clients connected to the FortiAP devices. | 
| diagnose wireless-controller wlac help | Show a list of debug options available for the wireless controller. | 
| diagnose wireless-controller wlac sta_filter diagnose wireless-controller wlac sta_filter clear diagnose wireless-controller wlac sta_filter <aa:bb:cc:dd:ee:ff> 255 diagnose debug enable | Start real-time debugging of a wireless client/station that connects to the FortiAP.  | 
| diagnose wireless-controller wlac -c vap | Show virtual access point information, including its MAC address, BSSID, SSID, the interface name, and the IP address of the APs that are broadcasting it. | 
| diagnose wireless-controller wlac wtp_filter diagnose wireless-controller wlac wtp_filter clear diagnose wireless-controller wlac wtp_filter <FAP-SN> 0-<x.x.x.x>:5246 255 diagnose debug application cw_acd 0x7ff | Show the wireless termination point (WTP), or FortiAP, debugging on the wireless controller if FortiAP is failing to connect to FortiGate.  | 
Other services
High availability
| Command | Description | 
|---|---|
| diagnose system ha status get system ha status | Show HA status and information. | 
| execute ha manage <index> <username> | Log into and manage a specific HA member. | 
| diagnose sys ha checksum cluster | Show checksum information of all cluster members. | 
| diagnose sys ha checksum show <vdom> | Show detailed checksum information for a VDOM. | 
| diagnose sys ha checksum recalculate | Recalculate HA checksums. | 
| diagnose sys ha recalculate-extfile-signature | Recalculate HA external files signatures. | 
| diagnose sys ha reset-uptime | Reset the HA uptime. This is used to test failover. | 
| diagnose debug application hatalk -1 diagnose debug application hasync -1 diagnose debug application harelay -1 diagnose debug enable | Start real-time debugging of HA daemons. | 
| diagnose sys ha history read | Show HA history. | 
| execute ha synchronize stop execute ha synchronize start | Manually start and stop HA synchronization. | 
ZTNA
|  | The WAD daemon handles proxy related processing. The FortiClient NAC daemon (fcnacd) handles FortiGate to EMS connectivity. | 
| Command | Description | 
|---|---|
| diagnose endpoint fctems test-connectivity <EMS> | Verify FortiGate to FortiClient EMS connectivity. | 
| execute fctems verify <EMS> | Verify the FortiClient EMS's certificate. | 
| diagnose test application fcnacd 2 | Dump the EMS connectivity information. | 
| diagnose debug app fcnacd -1 diagnose debug enable | Run real-time FortiClient NAC daemon debugs. | 
| diagnose endpoint ec-shm list <ip> <mac> <EMS_serial_number> <EMS_tenant_id> | Show the endpoint record list. Optionally, add filters. | 
| diagnose endpoint lls-comm send ztna find-uid <uid> <EMS_serial_number> <EMS_tenant_id> | Query endpoints by client UID, EMS serial number, and EMS tenant ID. | 
| diagnose endpoint lls-comm send ztna find-ip-vdom <ip> <vdom> | Query endpoints by the client IP-VDOM pair. | 
| diagnose wad dev query-by uid <uid> <EMS_serial_number> <EMS_tenant_id> | Query from WAD diagnose command by UID, EMS serial number, and EMS tenant ID. | 
| diagnose wad dev query-by ipv4 <ip> | Query from WAD diagnose command by IP address. | 
| diagnose firewall dynamic list | List EMS security posture tags and all dynamic IP and MAC addresses. | 
| diagnose test application fcnacd 7 diagnose test application fcnacd 8 | Check the FortiClient NAC daemon ZTNA and route cache. | 
| diagnose wad worker policy list | Display statistics associated with application gateway rules. | 
| diagnose wad debug enable category all diagnose wad debug enable level verbose diagnose debug enable | Run real-time WAD debugs. | 
| diagnose debug reset | Reset debugs when completed | 
Logging
| Command | Description | 
|---|---|
| diagnose log test | Generate logs for testing. | 
| execute log filter <filter> | Set log filters. | 
| execute log filter | Show log filters. | 
| exec log display | Show filtered logs. | 
| execute log delete | Delete filtered logs. | 
| diagnose debug application miglogd -1 diagnose debug enable | Start real-time debugging of logging process miglogd. | 
| execute log fortianalyzer test-connectivity | Test connectivity between FortiGate and FortiAnalyzer. | 
Traffic shaping
| Command | Description | 
|---|---|
| diagnose firewall shaper traffic-shaper list | Show configured traffic shapers. | 
| diagnose firewall shaper traffic-shaper stats list | Show traffic shaper statistics. | 
SIP session helper
| Command | Description | 
|---|---|
| diagnose sys sip status | Show SIP status. | 
| diagnose sys sip mapping list | Show SIP mapping list. | 
| diagnose sys sip dialog list | Show SIP dialogue list. | 
| diagnose debug application sip -1 diagnose debug enable | Start real-time SIP debugging. | 
SIP ALG
| Command | Description | 
|---|---|
| diagnose sys sip-proxy calls list | Show list of active SIP proxy calls. | 
| diagnose sys sip-proxy stats | Show SIP proxy statistics. | 
| diagnose sys sip-proxy session list | Show SIP proxy session list. | 
| diagnose debug application sip -1 diagnose debug enable | Start real-time SIP debugging. |
