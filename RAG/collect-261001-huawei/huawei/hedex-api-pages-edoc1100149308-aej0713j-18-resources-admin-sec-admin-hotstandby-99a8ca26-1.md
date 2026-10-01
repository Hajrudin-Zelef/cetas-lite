---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100149308-aej0713j-18-resources-admin-sec-admin-hotstandby-99a8ca26-1
title: "Configure a security policy to allow intranet users to access the Internet."
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100149308-aej0713j-18-resources-admin-sec-admin-hotstandby--99a8ca26.md
source_anchor: ""
source_lines: [1, 54]
sha256: 549d7fea06f41aead29dcdf6ac1bac6d05030610840431e79e476473c24368ac
---

# Configure a security policy to allow intranet users to access the Internet.

This section provides a CLI example of configuring hot standby in load balancing mode in which the service interfaces of the firewalls work at Layer 3 and connect to switches in upstream and downstream directions.
As shown in Figure 1, the service interfaces of the FWs work at Layer 3 and are directly connected to switches. The upstream switch is connected to the carrier network, and the public IP addresses the carrier assigns to the enterprise are 1.1.1.3 and 1.1.1.4. The FWs are expected to work in load balancing mode. Normally, both FW_A and FW_B forward traffic. If either FW fails, the other FW forwards all traffic to ensure service continuity.
| FW_A | FW_B | 
|---|---|
| # Set IP addresses for the interfaces on FWs. |  | 
| <FW_A> system-view  [FW_A] interface GigabitEthernet 0/0/1 [FW_A-GigabitEthernet0/0/1] ip address 10.2.0.1 24 [FW_A-GigabitEthernet0/0/1] quit  [FW_A] interface GigabitEthernet 0/0/3 [FW_A-GigabitEthernet0/0/3] ip address 10.3.0.1 24 [FW_A-GigabitEthernet0/0/3] quit  [FW_A] interface GigabitEthernet 0/0/7 [FW_A-GigabitEthernet0/0/7] ip address 10.10.0.1 24 [FW_A-GigabitEthernet0/0/7] quit  | <FW_B> system-view  [FW_B] interface GigabitEthernet 0/0/1 [FW_B-GigabitEthernet0/0/1] ip address 10.2.0.2 24 [FW_B-GigabitEthernet0/0/1] quit  [FW_B] interface GigabitEthernet 0/0/3 [FW_B-GigabitEthernet0/0/3] ip address 10.3.0.2 24 [FW_B-GigabitEthernet0/0/3] quit  [FW_B] interface GigabitEthernet 0/0/7 [FW_B-GigabitEthernet0/0/7] ip address 10.10.0.2 24 [FW_B-GigabitEthernet0/0/7] quit  | 
| # Assign the interfaces to security zones on FWs. |  | 
| [FW_A] firewall zone untrust [FW_A-zone-untrust] add interface GigabitEthernet 0/0/1 [FW_A-zone-untrust] quit [FW_A] firewall zone trust [FW_A-zone-trust] add interface GigabitEthernet 0/0/3 [FW_A-zone-trust] quit  [FW_A] firewall zone dmz [FW_A-zone-dmz] add interface GigabitEthernet 0/0/7 [FW_A-zone-dmz] quit  | [FW_B] firewall zone untrust [FW_B-zone-untrust] add interface GigabitEthernet 0/0/1 [FW_B-zone-untrust] quit [FW_B] firewall zone trust [FW_B-zone-trust] add interface GigabitEthernet 0/0/3 [FW_B-zone-trust] quit  [FW_B] firewall zone dmz [FW_B-zone-dmz] add interface GigabitEthernet 0/0/7 [FW_B-zone-dmz] quit  | 
| # Create a default route with next hop 1.1.1.10 on FWs. |  | 
| [FW_A] ip route-static 0.0.0.0 0.0.0.0 1.1.1.10 | [FW_B] ip route-static 0.0.0.0 0.0.0.0 1.1.1.10 | 
To implement load balancing, configure two VRRP groups on each service interface and set the status of one VRRP group to Active and the other to Standby.
| FW_A | FW_B | 
|---|---|
| # Configure VRRP groups 1 and 2 on upstream service interface GE0/0/1 of FW_A and set the status of VRRP group 1 to Active and status of VRRP group 2 to Standby. Configure VRRP groups 1 and 2 on upstream service interface GE0/0/1 of FW_B and set the status of VRRP group 1 to Standby and status of VRRP group 2 to Active. Note that if the interface IP address resides on a different subnet from the address of the VRRP group, you need to specify a subnet mask when setting the address of the VRRP group. |  | 
| [FW_A] interface GigabitEthernet 0/0/1 [FW_A-GigabitEthernet0/0/1] vrrp vrid 1 virtual-ip 1.1.1.3 24 active [FW_A-GigabitEthernet0/0/1] vrrp vrid 2 virtual-ip 1.1.1.4 24 standby [FW_A-GigabitEthernet0/0/1] quit | [FW_B] interface GigabitEthernet 0/0/1 [FW_B-GigabitEthernet0/0/1] vrrp vrid 1 virtual-ip 1.1.1.3 24 standby [FW_B-GigabitEthernet0/0/1] vrrp vrid 2 virtual-ip 1.1.1.4 24 active [FW_B-GigabitEthernet0/0/1] quit | 
| # Configure VRRP groups 3 and 4 on downstream service interface GE0/0/3 of FW_A and set the status of VRRP group 3 to Active and status of VRRP group 4 to Standby. Configure VRRP groups 3 and 4 on downstream service interface GE0/0/3 of FW_B and set the status of VRRP group 3 to Standby and status of VRRP group 4 to Active. |  | 
| [FW_A] interface GigabitEthernet 0/0/3 [FW_A-GigabitEthernet0/0/3] vrrp vrid 3 virtual-ip 10.3.0.3 active [FW_A-GigabitEthernet0/0/3] vrrp vrid 4 virtual-ip 10.3.0.4 standby [FW_A-GigabitEthernet0/0/3] quit  | [FW_B] interface GigabitEthernet 0/0/3 [FW_B-GigabitEthernet0/0/3] vrrp vrid 3 virtual-ip 10.3.0.3 standby [FW_B-GigabitEthernet0/0/3] vrrp vrid 4 virtual-ip 10.3.0.4 active [FW_B-GigabitEthernet0/0/3] quit  | 
| FW_A | FW_B | 
|---|---|
| # Configure quick session backup on both FWs in case of inconsistent forward and return packet paths. |  | 
| [FW_A] hrp mirror session enable | [FW_B] hrp mirror session enable | 
| # Specify the heartbeat interface and enable hot standby on FWs. |  | 
| [FW_A] hrp interface GigabitEthernet 0/0/7 remote 10.10.0.2  [FW_A] hrp enable  | [FW_B] hrp interface GigabitEthernet 0/0/7 remote 10.10.0.1  [FW_B] hrp enable  | 
# Configure a security policy to allow intranet users to access the Internet.
HRP_M[FW_A] security-policy
HRP_M[FW_A-policy-security] rule name trust_to_untrust  
HRP_M[FW_A-policy-security-rule-trust_to_untrust] source-zone trust
HRP_M[FW_A-policy-security-rule-trust_to_untrust] destination-zone untrust
HRP_M[FW_A-policy-security-rule-trust_to_untrust] action permit
HRP_M[FW_A-policy-security-rule-trust_to_untrust] source-address 10.3.0.0 24
HRP_M[FW_A-policy-security-rule-trust_to_untrust] quit
HRP_M[FW_A-policy-security] quit  
# Configure a NAT policy to translate source addresses on subnet 10.3.0.0/24 to an IP address in the NAT address pool (1.1.2.5 to 1.1.2.8) when intranet users access the Internet.
HRP_M[FW_A] nat address-group group1
HRP_M[FW_A-address-group-group1] section 0 1.1.2.5 1.1.2.8
HRP_M[FW_A-address-group-group1] route enable
HRP_M[FW_A-address-group-group1] quit
HRP_M[FW_A] nat-policy
HRP_M[FW_A-policy-nat] rule name policy_nat1  
HRP_M[FW_A-policy-nat-rule-policy_nat1] source-zone trust
HRP_M[FW_A-policy-nat-rule-policy_nat1] destination-zone untrust
HRP_M[FW_A-policy-nat-rule-policy_nat1] source-address 10.3.0.0 24 
HRP_M[FW_A-policy-nat-rule-policy_nat1] action source-nat address-group group1
HRP_M[FW_A-policy-nat-rule-policy_nat1] quit
HRP_M[FW_A-policy-nat] quit 
# To prevent port conflicts in address translation on the FWs in load balancing mode, configure available port ranges respectively on FW_A and FW_B. The configuration on FW_A is as follows:
HRP_M[FW_A] hrp nat resource primary-group
After the command is executed on FW_A, FW_B will automatically back it up and convert it to the hrp nat resource secondary-group command.
# Add the three interfaces of the switches to the same VLANs accordingly. For configuration commands, refer to related documents of the switches.
# On some intranet PCs, specify the IP address (10.3.0.3) of VRRP group 3 as the default gateway address and on some other intranet PCs, specify the IP address (10.3.0.4) of VRRP group 4 as the default gateway address to implement load balancing of intranet traffic.
# Configure equal-cost routes to the NAT address pool, with the next hops being the virtual IP addresses of VRRP group 1 and VRRP group 2.
Run the display vrrp command on FW_A and FW_B to check the status information about the interfaces in the VRRP group. If the following information is displayed, the VRRP group is successfully created.
| FW_A | FW_B | 
|---|---|
