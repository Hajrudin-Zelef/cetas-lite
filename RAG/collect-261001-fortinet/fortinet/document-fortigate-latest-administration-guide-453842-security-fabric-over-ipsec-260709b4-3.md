---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4-3
title: "document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4.md
source_anchor: ""
source_lines: [295, 426]
sha256: d50c84f3395b703af39138a8c580fe420933f81581d482018ae775d79f09924d
---

# document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4

- 
                                                    Configure the IPsec VPN interface IP address which will be used to form Security Fabric: 
  - 
                                                            Go to Network > Interfaces.
  - 
                                                            Edit To-HQ1: 
    - 
                                                                    Set Role to WAN.
    - 
                                                                    Set the IP/Network Mask to 10.10.10.3/255.255.255.255.
    - 
                                                                    Set Remote IP/Network Mask to 10.10.10.1/255.255.255.0.0.
  - 
                                                                    
- 
                                                            
- 
                                                    Configure the IPsec VPN local and remote subnets: 
  - 
                                                            Go to Policy & Objects > Addresses.
  - 
                                                            Click Create New 
    - 
                                                                    Set Name to To-HQ1_local_subnet_1.
    - 
                                                                    Set Type to Subnet.
    - 
                                                                    Set IP/Network Mask to 10.1.100.0/24.
  - 
                                                                    
  - 
                                                            Click OK.
  - 
                                                            Click Create New 
    - 
                                                                    Set Name to To-HQ1_remote_subnet_1.
    - 
                                                                    Set Type to Subnet.
    - 
                                                                    Set IP/Network Mask to 192.168.8.0/24.
  - 
                                                                    
  - 
                                                            Click OK.
- 
                                                            
- 
                                                    Configure the IPsec VPN static routes: 
  - 
                                                            Go to Network > Static Routes and click Create New or Create New > IPv4 Static Route. 
    - 
                                                                    For Named Address, select Type and select To-HQ1_remote_subnet_1.
    - 
                                                                    Set Interface to To-HQ1.
  - 
                                                                    
  - 
                                                            Click OK.
  - 
                                                            Click Create New or Create New > IPv4 Static Route. 
    - 
                                                                    For Named Address, select Type and select To-HQ1_remote_subnet_1.
    - 
                                                                    Set Interface to Blackhole.
    - 
                                                                    Set Administrative Distance to 254.
  - 
                                                                    
  - 
                                                            Click OK.
- 
                                                            
- 
                                                    Configure the IPsec VPN policies: 
  - 
                                                            Go to Policy & Objects > Firewall Policy and click Create New. 
    - 
                                                                    Set Name to vpn_To-HQ1_local.
    - 
                                                                    Set Incoming Interface to vlan20.
    - 
                                                                    Set Outgoing Interface to To-HQ1.
    - 
                                                                    Set Source to To-HQ1_local_subnet_1.
    - 
                                                                    Set Destination to To-HQ1_remote_subnet_1.
    - 
                                                                    Set Schedule to Always.
    - 
                                                                    Set Service to All.
    - 
                                                                    Disable NAT.
  - 
                                                                    
  - 
                                                            Click OK.
  - 
                                                            Click Create New. 
    - 
                                                                    Set Name to vpn_To-HQ1_remote.
    - 
                                                                    Set Incoming Interface to To-HQ1.
    - 
                                                                    Set Outgoing Interface to vlan20.
    - 
                                                                    Set Source to To-HQ1_remote_subnet_1.
    - 
                                                                    Set Destination to -HQ1_local_subnet_1.
    - 
                                                                    Set Schedule to Always.
    - 
                                                                    Set Service to All.
    - 
                                                                    Disable NAT.
  - 
                                                                    
  - 
                                                            Click OK.
- 
                                                            
- 
                                                    Configure the Security Fabric: 
  - 
                                                            Go to Security Fabric > Fabric Connectors and double-click the Security Fabric Setup card.
  - 
                                                            In the Settings tab, set the Security Fabric role to Join Existing Fabric. FortiAnalyzer automatically enables logging. FortiAnalyzer settings will be retrieved when the downstream FortiGate connects to the root FortiGate.
  - 
                                                            Set the Upstream FortiGate IP to 10.10.10.1.
  - 
                                                            Click OK.
- 
                                                            
