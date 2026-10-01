---
id: collect-261001-mikrotik/mikrotik/2022-12-19-ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-mikrotik-routeros-f93e83b0-3
title: "IKEv2 / IPsec Site-to-Site VPN Fortinet FortiGate <-> Mikrotik RouterOS"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-mikrotik/2022-12-19-ikev2-ipsec-site-to-site-vpn-fortinet-fortigate-mikrotik-routeros-f93e83b0.md
source_anchor: ""
source_lines: [132, 187]
sha256: 7f81a0f414f15c8a45378751dfc2490c9cb4eb208525308f5b02d6e84924d355
---

# IKEv2 / IPsec Site-to-Site VPN Fortinet FortiGate <-> Mikrotik RouterOS

  DPD sent/recv: 00000000/00000000
And then print the details of each phase 2 configuration on their own:
FG60F (root) # get vpn ipsec tunnel name vpn-to-mikrotik
gateway
  name: 'vpn-to-mikrotik'
  local-gateway: 10.0.0.1:0 (static)
  remote-gateway: 10.0.0.2:0 (static)
  dpd-link: on
  mode: ike-v2
  interface: 'wan2' (6)
  rx  packets: 0  bytes: 0  errors: 0
  tx  packets: 0  bytes: 0  errors: 0
  dpd: on-demand/negotiated  idle: 20000ms  retry: 3  count: 0
  selectors
    name: 'vpn-to-mikrotik'
    auto-negotiate: enable
    mode: tunnel
    src: 0:192.168.100.0/255.255.255.0:0
    dst: 0:192.168.200.0/255.255.255.0:0
    SA
      lifetime/rekey: 3600/2815   
      mtu: 1280
      tx-esp-seq: 1
      replay: enabled
      qat: 0
      inbound
        spi: 83391146
        enc:  aes-gc  fe03aea03aa08dee28c9a4ad2022d55113ea72be2078e0d5c77c2eacafe5ccb7d87e3611
        auth:   null  
      outbound
        spi: 09387f45
        enc:  aes-gc  6cb5756673b296440c9a0e428662abd80e8e444e97eb2ba8777fd2f9c9ff6a6e033a5f77
        auth:   null  
      NPU acceleration: none
    SA
      lifetime/rekey: 3600/2844   
      mtu: 1280
      tx-esp-seq: 1
      replay: enabled
      qat: 0
      inbound
        spi: 83391147
        enc:  aes-gc  7e4d12c5363b3aec6c29a9872e50ad65f4e1969e8dbc2db753853c620fd70e0448fd3339
        auth:   null  
      outbound
        spi: 03b2b08e
        enc:  aes-gc  034b284ccfdd7bdfb7b6aeb81f9c8f9a94fc5eb03e595fc0966011c46aa88b6637b06854
        auth:   null  
      NPU acceleration: none
For the Mikrotik RouterOS site the monitoring is done in the „IP –> IPsec –> Active Peers“ using Winbox or the WebUI:
You can get the same output using the commandline:
[admin@MikroTik] > /ip/ipsec/active-peers/print
Columns: ID, STATE, UPTIME, PH2-TOTAL, REMOTE-ADDRESS
# ID        STATE        UPTIME  PH2-TOTAL  REMOTE-ADDRESS
0 10.0.0.1  established  1m32s           1  10.0.0.1      
[admin@MikroTik] >
