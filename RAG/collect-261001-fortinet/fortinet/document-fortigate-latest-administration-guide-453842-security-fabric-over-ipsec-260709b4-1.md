---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4-1
title: "document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4.md
source_anchor: ""
source_lines: [1, 151]
sha256: 6a2d39df032c7c23512757fb6eb8ccf645319e0ee26262acf467cc137c3a2497
---

# document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4

Security Fabric over IPsec VPN
                                                
                                            
                                            This is an example of configuring Security Fabric over IPsec VPN.
Sample topology
This sample topology shows a downstream FortiGate (HQ2) connected to the root FortiGate (HQ1) over IPsec VPN to join Security Fabric.
Sample configuration
To configure the root FortiGate (HQ1):
- 
                                                    Configure the interface: 
  - 
                                                            Go to Network > Interfaces.
  - 
                                                            Edit port2: 
    - 
                                                                    Set Role to WAN.
    - 
                                                                    For the interface connected to the internet, set the IP/Network Mask to 10.2.200.1/255.255.255.0
  - 
                                                                    
  - 
                                                            Edit port6: 
    - 
                                                                    Set Role to DMZ.
    - 
                                                                    For the interface connected to FortiAnalyzer, set the IP/Network Mask to 192.168.8.250/255.255.255.0
  - 
                                                                    
- 
                                                            
- 
                                                    Configure the static route to connect to the internet: 
  - 
                                                            Go to Network > Static Routes and click Create New or Create New > IPv4 Static Route. 
    - 
                                                                    Set Destination to 0.0.0.0/0.0.0.0.
    - 
                                                                    Set Interface to port2.
    - 
                                                                    Set Gateway Address to 10.2.200.2.
  - 
                                                                    
  - 
                                                            Click OK.
- 
                                                            
- 
                                                    Configure the IPsec VPN: 
  - 
                                                            Go to VPN > IPsec Wizard. 
    - 
                                                                    Set Name to To-HQ2.
    - 
                                                                    Set Template Type to Custom.
    - 
                                                                    Click Next.
    - 
                                                                    Set Authentication to Method.
    - 
                                                                    Set Pre-shared Key to 123456.
  - 
                                                                    
  - 
                                                            Leave all other fields in their default values and click OK.
- 
                                                            
- 
                                                    Configure the IPsec VPN interface IP address which will be used to form Security Fabric: 
  - 
                                                            Go to Network > Interfaces.
  - 
                                                            Edit To-HQ2: 
    - 
                                                                    Set Role to LAN.
    - 
                                                                    Set the IP/Network Mask to 10.10.10.1/255.255.255.255.
    - 
                                                                    Set Remote IP/Network Mask to 10.10.10.3/255.255.255.0.
  - 
                                                                    
- 
                                                            
- 
                                                    Configure the IPsec VPN local and remote subnets: 
  - 
                                                            Go to Policy & Objects > Addresses.
  - 
                                                            Click Create New 
    - 
                                                                    Set Name to To-HQ2_remote_subnet_2.
    - 
                                                                    Set Type to Subnet.
    - 
                                                                    Set IP/Network Mask to 10.10.10.3/32.
  - 
                                                                    
  - 
                                                            Click OK.
  - 
                                                            Click Create New 
    - 
                                                                    Set Name to To-HQ2_local_subnet_1.
    - 
                                                                    Set Type to Subnet.
    - 
                                                                    Set IP/Network Mask to 192.168.8.0/24.
  - 
                                                                    
  - 
                                                            Click OK.
  - 
                                                            Click Create New 
    - 
                                                                    Set Name to To-HQ2_remote_subnet_1.
    - 
                                                                    Set Type to Subnet.
    - 
                                                                    Set IP/Network Mask to 10.1.100.0/24.
  - 
                                                                    
  - 
                                                            Click OK.
- 
                                                            
- 
                                                    Configure the IPsec VPN static routes: 
  - 
                                                            Go to Network > Static Routes.
  - 
                                                            Click Create New or Create New > IPv4 Static Route. 
    - 
                                                                    For Named Address, select Type and select To-HQ2_remote_subnet_1.
    - 
                                                                    Set Interface to To-HQ2.
 Click OK.
  - 
                                                                    
  - 
                                                            Click Create New or Create New > IPv4 Static Route. 
    - 
                                                                    For Named Address, select Type and select To-HQ2_remote_subnet_1.
    - 
                                                                    Set Interface to Blackhole.
    - 
                                                                    Set Administrative Distance to 254.
  - 
                                                                    
  - 
                                                            Click OK.
- 
                                                            
