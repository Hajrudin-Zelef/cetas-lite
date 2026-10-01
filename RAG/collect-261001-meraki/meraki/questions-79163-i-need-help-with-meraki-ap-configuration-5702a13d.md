---
id: collect-261001-meraki/meraki/questions-79163-i-need-help-with-meraki-ap-configuration-5702a13d
title: "questions-79163-i-need-help-with-meraki-ap-configuration-5702a13d"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-meraki/questions-79163-i-need-help-with-meraki-ap-configuration-5702a13d.md
source_anchor: ""
source_lines: [1, 33]
sha256: 2072e185732432f3950069d171a1af5b0954e9159fa4bdc50e1a7c88c77e974b
---

# questions-79163-i-need-help-with-meraki-ap-configuration-5702a13d

Hello everyone i have this port gi2/0/41 was configured for work station and port gi1/0/18 is configured for Meraki Ap , i need to correct the configuration in port gi2/0/41 and make is as the same port gi1/0/18, i do not want effect the network or cause issue and i'm new in network environment
BLDG-B-IDF-01#show run int gi1/0/18
Building configuration...
Current configuration : 119 bytes
!
interface GigabitEthernet1/0/18
 description MERAKI AP
 switchport trunk native vlan 103
 switchport mode trunk
end
BLDG-B-IDF-01#show run int gi2/0/41
Building configuration...
Current configuration : 617 bytes
!
interface GigabitEthernet2/0/41
 description MERAKI AP
 switchport access vlan 110
 switchport trunk native vlan 103
 switchport mode trunk
 switchport nonegotiate
 switchport voice vlan 922
 switchport port-security maximum 5
 switchport port-security violation restrict
 switchport port-security aging time 1
 switchport port-security aging type inactivity
 switchport port-security
 storm-control broadcast level 2.00 1.00
 auto qos trust dscp
 spanning-tree portfast
 spanning-tree bpduguard enable
 service-policy input AutoQos-4.0-Trust-Dscp-Input-Policy
 service-policy output AutoQos-4.0-Output-Policy
end
