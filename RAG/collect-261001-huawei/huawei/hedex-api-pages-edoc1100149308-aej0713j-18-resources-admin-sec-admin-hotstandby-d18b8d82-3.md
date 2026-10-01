---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100149308-aej0713j-18-resources-admin-sec-admin-hotstandby-d18b8d82-3
title: "Configure a security policy to allow intranet users to access the Internet."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100149308-aej0713j-18-resources-admin-sec-admin-hotstandby--d18b8d82.md
source_anchor: ""
source_lines: [57, 57]
sha256: bea5a142ce2c2e45238688cc1c5abeebd60d0b1275c264538d8874d433787e29
---

# Configure a security policy to allow intranet users to access the Internet.

| #  hrp enable  hrp interface GigabitEthernet 0/0/7 remote 10.10.0.2 # interface GigabitEthernet 0/0/1  ip address 10.2.0.1 255.255.255.0  vrrp vrid 1 virtual-ip 1.1.1.1 255.255.255.0 active # interface GigabitEthernet 0/0/3  ip address 10.3.0.1 255.255.255.0  vrrp vrid 2 virtual-ip 10.3.0.3 active # interface GigabitEthernet 0/0/7  ip address 10.10.0.1 255.255.255.0 # firewall zone trust  set priority 85  add interface GigabitEthernet 0/0/3 # firewall zone untrust  set priority 5  add interface GigabitEthernet 0/0/1 # firewall zone dmz  set priority 50  add interface GigabitEthernet 0/0/7 #  ip route-static 0.0.0.0 0.0.0.0 1.1.1.10 #     nat address-group group1  route enable   section 0 1.1.1.2 1.1.1.5 #     security-policy    rule name trust_to_untrust   source-zone trust     destination-zone untrust   source-address 10.3.0.0 24   action permit     #     nat-policy    rule name policy_nat1   source-zone trust   destination-zone untrust   source-address 10.3.0.0 16    action source-nat address-group group1 | #  hrp enable  hrp interface GigabitEthernet 0/0/7 remote 10.10.0.1 # interface GigabitEthernet 0/0/1  ip address 10.2.0.2 255.255.255.0  vrrp vrid 1 virtual-ip 1.1.1.1 255.255.255.0 standby # interface GigabitEthernet 0/0/3  ip address 10.3.0.2 255.255.255.0  vrrp vrid 2 virtual-ip 10.3.0.3 standby # interface GigabitEthernet 0/0/7  ip address 10.10.0.2 255.255.255.0 #     firewall zone trust  set priority 85  add interface GigabitEthernet 0/0/3 #     firewall zone untrust  set priority 5  add interface GigabitEthernet 0/0/1 #     firewall zone dmz      set priority 50       add interface GigabitEthernet0/0/7 #  ip route-static 0.0.0.0 0.0.0.0 1.1.1.10 #     nat address-group group1   route enable  section 0 1.1.1.2 1.1.1.5 #     security-policy    rule name trust_to_untrust   source-zone trust     destination-zone untrust   source-address 10.3.0.0 24   action permit     #     nat-policy    rule name policy_nat1   source-zone trust   destination-zone untrust   source-address 10.3.0.0 16    action source-nat address-group group1 |
