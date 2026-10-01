---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100149308-aej0713j-18-resources-admin-sec-admin-hotstandby-d18b8d82-1
title: "Configure a security policy to allow intranet users to access the Internet."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100149308-aej0713j-18-resources-admin-sec-admin-hotstandby--d18b8d82.md
source_anchor: ""
source_lines: [1, 43]
sha256: 44083cd435fbb23c5567d2dac0ba40423ea5e1b6ac5a0bf37deec80acf919719
---

# Configure a security policy to allow intranet users to access the Internet.

This section provides a CLI example of configuring hot standby in active/standby mode in which the service interfaces of the firewalls work at Layer 3 and connect to switches in upstream and downstream directions.
On the network shown in Figure 1, the service interfaces of two FWs work at Layer 3 and are directly connected to switches. The upstream switch is connected to the carrier network, and the public IP address the carrier assigns to the enterprise is 1.1.1.1. The FWs are expected to work in active/standby mode. Normally, traffic is forwarded by FW_A. If FW_A is faulty, FW_B takes over to ensure service continuity.
| FW_A | FW_B | 
|---|---|
| # Set IP addresses for the interfaces on FWs. |  | 
| <FW_A> system-view [FW_A] interface GigabitEthernet 0/0/1 [FW_A-GigabitEthernet0/0/1] ip address 10.2.0.1 24 [FW_A-GigabitEthernet0/0/1] quit  [FW_A] interface GigabitEthernet 0/0/3 [FW_A-GigabitEthernet0/0/3] ip address 10.3.0.1 24 [FW_A-GigabitEthernet0/0/3] quit  [FW_A] interface GigabitEthernet 0/0/7 [FW_A-GigabitEthernet0/0/7] ip address 10.10.0.1 24 [FW_A-GigabitEthernet0/0/7] quit  | <FW_B> system-view [FW_B] interface GigabitEthernet 0/0/1 [FW_B-GigabitEthernet0/0/1] ip address 10.2.0.2 24 [FW_B-GigabitEthernet0/0/1] quit  [FW_B] interface GigabitEthernet 0/0/3 [FW_B-GigabitEthernet0/0/3] ip address 10.3.0.2 24 [FW_B-GigabitEthernet0/0/3] quit  [FW_B] interface GigabitEthernet 0/0/7 [FW_B-GigabitEthernet0/0/7] ip address 10.10.0.2 24 [FW_B-GigabitEthernet0/0/7] quit  | 
| # Assign the interfaces to security zones on FWs. |  | 
| [FW_A] firewall zone trust [FW_A-zone-trust] add interface GigabitEthernet 0/0/3 [FW_A-zone-trust] quit  [FW_A] firewall zone dmz [FW_A-zone-dmz] add interface GigabitEthernet 0/0/7 [FW_A-zone-dmz] quit  [FW_A] firewall zone untrust [FW_A-zone-untrust] add interface GigabitEthernet 0/0/1 [FW_A-zone-untrust] quit | [FW_B] firewall zone trust [FW_B-zone-trust] add interface GigabitEthernet 0/0/3 [FW_B-zone-trust] quit  [FW_B] firewall zone dmz [FW_B-zone-dmz] add interface GigabitEthernet 0/0/7 [FW_B-zone-dmz] quit  [FW_B] firewall zone untrust [FW_B-zone-untrust] add interface GigabitEthernet 0/0/1 [FW_B-zone-untrust] quit | 
| # Create a default route with next hop 1.1.1.10 on FWs. |  | 
| [FW_A] ip route-static 0.0.0.0 0.0.0.0 1.1.1.10 | [FW_B] ip route-static 0.0.0.0 0.0.0.0 1.1.1.10 | 
| FW_A | FW_B | 
|---|---|
| # Configure VRRP group 1 on upstream service interface GE0/0/1 of FW_A and set the VRRP group status to Active. Configure VRRP group 1 on upstream service interface GE0/0/1 of FW_B and set the VRRP group status to Standby. Note that if the interface IP address resides on a different subnet from the address of the VRRP group, you need to specify a subnet mask when setting the address of the VRRP group. |  | 
| [FW_A] interface GigabitEthernet 0/0/1 [FW_A-GigabitEthernet0/0/1] vrrp vrid 1 virtual-ip 1.1.1.1 24 active [FW_A-GigabitEthernet0/0/1] quit | [FW_B] interface GigabitEthernet 0/0/1 [FW_B-GigabitEthernet0/0/1] vrrp vrid 1 virtual-ip 1.1.1.1 24 standby [FW_B-GigabitEthernet0/0/1] quit | 
| # Configure VRRP group 2 on downstream service interface GE0/0/3 of FW_A and set the VRRP group status to Active. Configure VRRP group 2 on downstream service interface GE0/0/3 of FW_B and set the VRRP group status to Standby. |  | 
| [FW_A] interface GigabitEthernet 0/0/3 [FW_A-GigabitEthernet0/0/3] vrrp vrid 2 virtual-ip 10.3.0.3 active [FW_A-GigabitEthernet0/0/3] quit  | [FW_B] interface GigabitEthernet 0/0/3 [FW_B-GigabitEthernet0/0/3] vrrp vrid 2 virtual-ip 10.3.0.3 standby [FW_B-GigabitEthernet0/0/3] quit  | 
| FW_A | FW_B | 
|---|---|
| [FW_A] hrp interface GigabitEthernet 0/0/7 remote 10.10.0.2  [FW_A] hrp enable  | [FW_B] hrp interface GigabitEthernet 0/0/7 remote 10.10.0.1  [FW_B] hrp enable  | 
# Configure a security policy to allow intranet users to access the Internet.
HRP_M[FW_A] security-policy
HRP_M[FW_A-policy-security] rule name trust_to_untrust  
HRP_M[FW_A-policy-security-rule-trust_to_untrust] source-zone trust
HRP_M[FW_A-policy-security-rule-trust_to_untrust] destination-zone untrust
HRP_M[FW_A-policy-security-rule-trust_to_untrust] source-address 10.3.0.0 24
HRP_M[FW_A-policy-security-rule-trust_to_untrust] action permit
HRP_M[FW_A-policy-security-rule-trust_to_untrust] quit
HRP_M[FW_A-policy-security] quit  
# Configure a NAT policy to translate source addresses on subnet 10.3.0.0/16 to an IP address in the NAT address pool (1.1.1.2 to 1.1.1.5) when intranet users access the Internet.
HRP_M[FW_A] nat address-group group1
HRP_M[FW_A-address-group-group1] section 0 1.1.1.2 1.1.1.5
HRP_M[FW_A-address-group-group1] route enable
HRP_M[FW_A-address-group-group1] quit
HRP_M[FW_A] nat-policy
HRP_M[FW_A-policy-nat] rule name policy_nat1
HRP_M[FW_A-policy-nat-rule-policy_nat1] source-zone trust
HRP_M[FW_A-policy-nat-rule-policy_nat1] destination-zone untrust
HRP_M[FW_A-policy-nat-rule-policy_nat1] source-address 10.3.0.0 16 
HRP_M[FW_A-policy-nat-rule-policy_nat1] action source-nat address-group group1
Configure equal-cost routes to FW, with the next hop being the virtual IP addresses of VRRP group 1.
Run the display vrrp command on FW_A and FW_B to check the status information about the interfaces in the VRRP group. If the following information is displayed, the VRRP group is successfully created.
| FW_A | FW_B | 
|---|---|
