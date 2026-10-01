---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100149308-aej0713j-18-resources-admin-sec-admin-hotstandby-99a8ca26-3
title: "Configure a security policy to allow intranet users to access the Internet."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100149308-aej0713j-18-resources-admin-sec-admin-hotstandby--99a8ca26.md
source_anchor: ""
source_lines: [63, 68]
sha256: d5847fa9e6f70be44b974212e5c48476a04e0c36efacfae55f5a771ef2ffa1db
---

# Configure a security policy to allow intranet users to access the Internet.

| HRP_M<FW_A> display firewall session table   Current Total Sessions : 1   icmp  VPN: public --> public 10.3.0.10:0[1.1.2.5:10298] --> 1.1.1.10:2048   | HRP_S<FW_B> display firewall session table   Current Total Sessions : 1   icmp  VPN:public --> public  Remote 10.3.0.10:0[1.1.2.5:10298] --> 1.1.1.10:2048 | 
The command output shows that sessions tagged with Remote are created on FW_B, indicating that sessions are successfully backed up after you configure hot standby.
Run the ping 1.1.1.10 -t command on the PC, pull out the cable from GE0/0/1 on FW_A, and then check whether active/standby switchover is performed and whether ping packets are discarded. Insert the cable back to GE0/0/1 on FW_A and check again whether active/standby switchover is performed and whether ping packets are discarded.
| FW_A | FW_B | 
|---|---|
| #  hrp enable  hrp interface GigabitEthernet 0/0/7 remote 10.10.0.2  hrp mirror session enable  hrp nat resource primary-group # interface GigabitEthernet 0/0/1  ip address 10.2.0.1 255.255.255.0  vrrp vrid 1 virtual-ip 1.1.1.3 255.255.255.0 active  vrrp vrid 2 virtual-ip 1.1.1.4 255.255.255.0 standby # interface GigabitEthernet 0/0/3  ip address 10.3.0.1 255.255.255.0  vrrp vrid 3 virtual-ip 10.3.0.3 active  vrrp vrid 4 virtual-ip 10.3.0.4 standby # interface GigabitEthernet 0/0/7  ip address 10.10.0.1 255.255.255.0 # firewall zone trust  set priority 85  add interface GigabitEthernet 0/0/3 # firewall zone dmz    set priority 50     add interface GigabitEthernet0/0/7 # firewall zone untrust  set priority 5     add interface GigabitEthernet 0/0/1 #  ip route-static 0.0.0.0 0.0.0.0 1.1.1.10 #      nat address-group group1    route enable   section 0 1.1.2.5 1.1.2.8 #     security-policy    rule name trust_to_untrust   source-zone trust     destination-zone untrust   source-address 10.3.0.0 24   action permit     #     nat-policy    rule name policy_nat1   source-zone trust   destination-zone untrust   source-address 10.3.0.0 24    action source-nat address-group group1 | #  hrp enable  hrp interface GigabitEthernet 0/0/7 remote 10.10.0.1  hrp mirror session enable  hrp nat resource secondary-group # interface GigabitEthernet 0/0/1  ip address 10.2.0.2 255.255.255.0  vrrp vrid 1 virtual-ip 1.1.1.3 255.255.255.0 standby  vrrp vrid 2 virtual-ip 1.1.1.4 255.255.255.0 active # interface GigabitEthernet 0/0/3  ip address 10.3.0.2 255.255.255.0  vrrp vrid 3 virtual-ip 10.3.0.3 standby  vrrp vrid 4 virtual-ip 10.3.0.4 active # interface GigabitEthernet 0/0/7  ip address 10.10.0.2 255.255.255.0 # firewall zone trust  set priority 85  add interface GigabitEthernet 0/0/3 # firewall zone dmz    set priority 50     add interface GigabitEthernet0/0/7 # firewall zone untrust  set priority 5   add interface GigabitEthernet 0/0/1 #  ip route-static 0.0.0.0 0.0.0.0 1.1.1.10 #      nat address-group group1    route enable   section 0 1.1.2.5 1.1.2.8 #     security-policy    rule name trust_to_untrust   source-zone trust     destination-zone untrust   source-address 10.3.0.0 24   action permit     #     nat-policy    rule name policy_nat1   source-zone trust   destination-zone untrust   source-address 10.3.0.0 24    action source-nat address-group group1 |
