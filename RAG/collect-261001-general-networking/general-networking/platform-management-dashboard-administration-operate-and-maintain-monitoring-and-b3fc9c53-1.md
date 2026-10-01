---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-operate-and-maintain-monitoring-and-b3fc9c53-1
title: "platform-management-dashboard-administration-operate-and-maintain-monitoring-and-b3fc9c53"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-operate-and-maintain-monitoring-and-b3fc9c53.md
source_anchor: ""
source_lines: [1, 51]
sha256: de83ee4e9483c7471754c12e6b34d1691c1daccd1841a13cf2188a554e9fdfef
---

# platform-management-dashboard-administration-operate-and-maintain-monitoring-and-b3fc9c53

Click 日本語 for Japanese

## Overview

This article provides a list of most common syslog event types, description of each event, and a sample output of each log.



### Meraki MX Security Appliance


| **Event type** | **Description** | **Sample Syslog Message** | 
| events (Auto VPN) | vpn connectivity change | 1380664922.583851938 MX84 events type=vpn_connectivity_change vpn_type='site-to-site' peer_contact='98.68.191.209:51856' peer_ident='2814ee002c075181bb1b7478ee073860' connectivity='false' | 
| events (Auto VPN) | vpn connectivity change | 1380664994.337961231 MX84 events type=vpn_connectivity_change vpn_type='site-to-site' peer_contact='98.68.191.209:51856' peer_ident='2814ee002c075181bb1b7478ee073860' connectivity='true' | 
| events | uplink connectivity change | Dec 6 08:46:12 192.168.1.1 1 1386337584.254756845 MX84 events Cellular connection down | 
| events | uplink connectivity change | Dec 6 08:45:24 192.168.1.1 1 1386337535.803931423 MX84 events failover to wan1 | 
| events | uplink connectivity change | Dec 6 08:43:43 192.168.1.1 1 1386337435.108107268 MX84 events failover to cellular | 
| events | uplink connectivity change | Dec 6 08:41:44 192.168.1.1 1 1386337316.207232138 MX84 events Cellular connection up | 
| events | dhcp no offers | Sep 11 16:12:41 192.168.10.1 1 1599865961.535491111 MX84 events dhcp no offers for mac A4:83:E7:XX:XX:XX host = 192.168.10.1 | 
| events | dhcp lease | Sep 11 16:05:15 192.168.10.1 1 1599865515.687171503 MX84 events dhcp lease of ip 192.168.10.68 from server mac E0:CB:BC:0F:XX:XX for client mac 8C:16:45:XX:XX:XX from router 192.168.10.1 on subnet 255.255.255.0 with dns 8.8.8.8, 8.8.4.4 | 
| urls | HTTP GET requests | 1374543213.342705328 MX84 urls src=192.168.1.186:63735 dst=69.58.188.40:80 mac=58:1F:AA:CE:61:F2 request: GET https://... | 
| Events | Content filtering block | 1787058176.597410862 MX67 events content_filtering_block url='https://www.facebook.com/...' category0='User-defined Blacklist' server='157.240.225.35:443' client_mac='XX:XX:XX:XX:XX:XX' | 
| L7_firewall | l7_firewall | 1787058162.285901262 MX67 l7_firewall src=10.10.10.3 dst=104.21.4.210 protocol=tcp sport=59526 dport=443 decision=blocked | 
| flows *(deprecated, this has been updated in MX18.101 and newer to "firewall")*  | L3 FW rule matched | 1374543986.038687615 MX84 flows src=192.168.1.186 dst=8.8.8.8 mac=58:1F:AA:CE:61:F2 protocol=udp sport=55719 dport=53 pattern: allow all | 
| firewall cellular_firewall vpn_firewall | L3 FW rule matched | 1374543986.038687615 MX84 firewall src=192.168.1.186 dst=8.8.8.8 mac=58:1F:AA:CE:61:F2 protocol=udp sport=55719 dport=53 pattern: allow all | 
| ids-alerts | ids signature matched | 1377449842.514782056 MX84 ids-alerts signature=129:4:1 priority=3 timestamp=1377449842.512569 direction=ingress protocol=tcp/ip src=74.125.140.132:80 | 
| ids-alerts | ids signature matched | 1377448470.246576346 MX84 ids-alerts signature=119:15:1 priority=2 timestamp=1377448470.238064 direction=egress protocol=tcp/ip src=192.168.111.254:56240 | 
| security_event ids_alerted | ids signature matched | signature=1:28423:1 priority=1 timestamp=1468531589.810079 dhost=98:5A:EB:E1:81:2F direction=ingress protocol=tcp/ip src=151.101.52.238:80 dst=192.168.128.2:53023 decision=blocked action=rst message: EXPLOIT-KIT Multiple exploit kit single digit exe detection | 
| security_event security_filtering_file_scanned | Malicious file blocked by amp | url=http://www.eicar.org/download/eicar.com.txt src=192.168.128.2:53150 dst=188.40.238.250:80 mac=98:5A:EB:E1:81:2F name='EICAR:EICAR_Test_file_not_a_virus-tpd' sha256=275a021bbfb6489e54d471899f7db9d1663fc695ec2fe2a2c4538aabf651fd0f disposition=malicious action=block | 
| security_event security_filtering_disposition_change | File issued retrospective malicious disposition | name=EICAR:EICAR_Test_file_not_a_virus-tpd sha256=275a021bbfb6489e54d471899f7db9d1663fc695ec2fe2a2c4538aabf651fd0f disposition=malicious action=allow | 
| events (post MX 15.12) | Establishing Phase 1 (IKE_SA) tunnel | VPN: <remote-peer-2\|12> IKE_SA remote-peer-2[12] established between 192.168.13.5[192.168.13.5]...192.168.13.2[192.168.13.2] | 
| events (post MX 15.12) | Establishing Phase 2 (Child_SA) tunnel | VPN: <remote-peer-2\|12> CHILD_SA net-2{1478} established with SPIs cd94e190(inbound) c2b06071(outbound) and TS 192.168.12.0/24 === 192.168.13.0/24 | 
| events (post MX 15.12) | Destroying Phase 1 (IKE_SA) tunnel | VPN: <remote-peer-2\|12> deleting IKE_SA remote-peer-2[12] between 192.168.13.5[192.168.13.5]...192.168.13.2[192.168.13.2] | 
| events (post MX 15.12) | Destroying Phase 2 (Child_SA) tunnel | VPN: <remote-peer-2\|12> closing CHILD_SA net-2{1478} with SPIs cd94e190(inbound) (0 bytes) c2b06071(outbound) (0 bytes) and TS 192.168.12.0/24 === 192.168.13.0/24 | 
| events | AnyConnect VPN general (various msgs) | 1720051390.733639600 labs_appliance events type=anyconnect_vpn_general msg= 'AnyConnect server is started. ' | 
| events | AnyConnect VPN authentication success | 1720045578.339796505 labs_appliance events type=anyconnect_vpn_auth_success msg= 'Peer IP=192.168.0.1 Peer port=57096 AAA[7]: AAA authentication successful ' | 
| events | AnyConnect VPN authentication failure | 1720051237.124589040 labs_appliance events type=anyconnect_vpn_auth_failure msg= 'Peer IP=192.168.0.1Peer port[8748] AAA[8]: AAA authenticate failed retval=7 - Authentication failure ' | 
| events | AnyConnect VPN session manager (various msgs) | 1720045578.340434385 labs_appliance events type=anyconnect_vpn_session_manager msg= 'Sess-ID[7] Peer IP=192.168.0.1 User[miles@meraki.net]: Session connected. Session Type: TLS ' | 
| events | AnyConnect VPN Connect | 1720045578.495767745 labs_appliance events anyconnect_vpn_connect user id 'miles@meraki.net' local ip 192.168.5.224 connected from 192.168.0.1 | 
| events | AnyConnect VPN Disconnect | 1720045578.515109505 labs_appliance events anyconnect_vpn_disconnect user id 'miles@meraki.net' local ip 192.168.5.135 connected from 192.168.0.1 | 
| events (pre MX 15.12) | purging ISAKMP-SA | 1578424543.894083034 labs_appliance events Site-to-site VPN: purging ISAKMP-SA spi=9d1bb66d7ddc5cf0:d98cd0ed59e82f13 | 
| events (pre MX 15.12) | ISAKMP-SA deleted | 1578424543.918665436 labs_appliance events Site-to-site VPN: ISAKMP-SA deleted 172.24.23.6[4500]-172.24.23.10[4500] spi:9d1bb66d7ddc5cf0:d98cd0ed59e82f13 | 
| events (pre MX 15.12) | IPsec-SA request queued due to no phase 1 found | 1578424549.917669303 labs_appliance events Site-to-site VPN: IPsec-SA request for 172.24.23.10 queued due to no phase1 found | 
| events (pre MX 15.12) | failed to get sainfo | 1578426208.829677788 labs_Z1 events Site-to-site VPN: failed to get sainfo | 
| events (pre MX 15.12) | failed to pre-process ph2 packet | 1578426208.915091184 labs_Z1 events Site-to-site VPN: failed to pre-process ph2 packet (side: 1, status: 1) | 
| events (pre MX 15.12) | phase2 negotiation failed due to time up waiting for phase1 | 1578424408.321445408 labs_appliance events Site-to-site VPN: phase2 negotiation failed due to time up waiting for phase1. ESP 172.24.23.10[0]->172.24.23.6[0] | 
| events (pre MX 15.12) | initiate new phase 1 negotiation | 1578424549.931720602 labs_appliance events Site-to-site VPN: initiate new phase 1 negotiation: 172.24.23.6[500]<=>172.24.23.10[500] | 
| events (pre MX 15.12) | ISAKMP-SA established | 1578424550.965202127 labs_appliance events Site-to-site VPN: ISAKMP-SA established 172.24.23.6[4500]-172.24.23.10[4500] spi:fb903f191f1c7566:4dc90bd31c7884c1 | 
| events (pre MX 15.12) | initiate new phase 2 negotiation | 1578424550.975495647 labs_appliance events Site-to-site VPN: initiate new phase 2 negotiation: 172.24.23.6[4500]<=>172.24.23.10[4500] | 
| events (pre MX 15.12) | IPsec-SA established | 1578424551.120459981 labs_appliance events Site-to-site VPN: IPsec-SA established: ESP/Tunnel 172.24.23.6[4500]->172.24.23.10[4500] spi=241280704(0xe61a6c0) | 

