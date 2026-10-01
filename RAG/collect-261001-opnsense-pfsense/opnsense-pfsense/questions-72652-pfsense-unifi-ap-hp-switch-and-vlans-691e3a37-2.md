---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-72652-pfsense-unifi-ap-hp-switch-and-vlans-691e3a37-2
title: "questions-72652-pfsense-unifi-ap-hp-switch-and-vlans-691e3a37"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-72652-pfsense-unifi-ap-hp-switch-and-vlans-691e3a37.md
source_anchor: ""
source_lines: [99, 172]
sha256: 2a521b1699579061587c4492ea0d1d3ef2d51fd5e6130a9696f5a91da088fb0d
---

# questions-72652-pfsense-unifi-ap-hp-switch-and-vlans-691e3a37

VLAN 200 (no_vpn_vlan)
switch-2520G# show vlan 200
 Status and Counters - VLAN Information - VLAN 200
  VLAN ID : 200    
  Name : no_vpn_vlan         
  Status : Port-based  Voice : No 
  Jumbo : No 
  Port Information Mode     Unknown VLAN Status    
  ---------------- -------- ------------ ----------
  1                Untagged Learn        Up        
  2                Untagged Learn        Down      
  3                Untagged Learn        Down      
  4                Untagged Learn        Down      
  5                Untagged Learn        Down      
  6                Untagged Learn        Down      
  7                Untagged Learn        Down      
  8                Untagged Learn        Down      
  9                Untagged Learn        Down      
  10               Untagged Learn        Down      
  11               Untagged Learn        Up        
  12               Untagged Learn        Down      
  13               Untagged Learn        Up        
  14               Untagged Learn        Down      
  15               Untagged Learn        Down      
  16               Untagged Learn        Up        
  17               Untagged Learn        Down      
  18               Untagged Learn        Down      
  19               Untagged Learn        Down      
  20               Untagged Learn        Down      
  21               Untagged Learn        Up        
  22               Untagged Learn        Down      
  23               Untagged Learn        Up        
  24               Untagged Learn        Down      
VLAN port assignment
  Port   DEFAULT_VLAN  no_vpn_vlan    |  Port   DEFAULT_VLAN  no_vpn_vlan
  ---- + ------------  ------------   |  ---- + ------------  ------------
  1    | No            Untagged       |  13   | No            Untagged
  2    | No            Untagged       |  14   | No            Untagged
  3    | No            Untagged       |  15   | No            Untagged
  4    | No            Untagged       |  16   | No            Untagged
  5    | No            Untagged       |  17   | No            Untagged
  6    | No            Untagged       |  18   | No            Untagged
  7    | No            Untagged       |  19   | No            Untagged
  8    | No            Untagged       |  20   | No            Untagged
  9    | No            Untagged       |  21   | No            Untagged
  10   | No            Untagged       |  22   | No            Untagged
  11   | No            Untagged       |  23   | No            Untagged
  12   | No            Untagged       |  24   | No            Untagged
UPDATE-2/SOLUTION (as provided by Zac67):
Change IP addresses for both VLANS: Switch Configuration - Internet (IP) Service
  Default Gateway :                
  Default TTL     : 64   
  Arp Age         : 20  
          VLAN           IP Config     IP Address       Subnet Mask
  -------------------- + ----------  ---------------  ---------------
  DEFAULT_VLAN         | Manual      192.168.1.1      255.255.255.0
  no_vpn_vlan          | Manual      192.168.2.1      255.255.255.0
Correct VLAN tags:
               Switch Configuration - VLAN - VLAN Port Assignment
  Port   DEFAULT_VLAN  no_vpn_vlan    |  Port   DEFAULT_VLAN  no_vpn_vlan
  ---- + ------------  ------------   |  ---- + ------------  ------------
  1    | Untagged      Tagged         |  13   | Untagged      No
  2    | Untagged      No             |  14   | Untagged      No
  3    | Untagged      No             |  15   | Untagged      No
  4    | Untagged      No             |  16   | Untagged      No
  5    | Untagged      No             |  17   | Untagged      No
  6    | Untagged      No             |  18   | Untagged      No
  7    | Untagged      No             |  19   | Untagged      No
  8    | Untagged      No             |  20   | Untagged      No
  9    | Untagged      No             |  21   | Untagged      No
  10   | Untagged      No             |  22   | Untagged      No
  11   | Untagged      No             |  23   | Untagged      Tagged
  12   | Untagged      No             |  24   | Untagged      No
192./162/.1.66above a typo or deliberate? Note that 192.162.0.0/22 is a Russian network and not private.{}).
