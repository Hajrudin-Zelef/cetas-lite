---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4-2
title: "document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4.md
source_anchor: ""
source_lines: [152, 294]
sha256: 4aada28d63079cb2ef15bbd4cadbfc42645bbf4feb8b0d19564f90be7d60b584
---

# document-fortigate-latest-administration-guide-453842-security-fabric-over-ipsec-260709b4

- 
                                                    Configure the IPsec VPN policies: 
  - 
                                                            Go to Policy & Objects > Firewall Policy
  - 
                                                            Click Create New. 
    - 
                                                                    Set Name to vpn_To-HQ2_local.
    - 
                                                                    Set Incoming Interface to port6.
    - 
                                                                    Set Outgoing Interface to To-HQ2.
    - 
                                                                    Set Source to To-HQ2_local_subnet_1.
    - 
                                                                    Set Destination to To-HQ2_remote_subnet_1.
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
                                                                    Set Name to vpn_To-HQ2_remote.
    - 
                                                                    Set Incoming Interface to To-HQ2.
    - 
                                                                    Set Outgoing Interface to port6.
    - 
                                                                    Set Source to To-HQ2_remote_subnet_1, To-HQ2_remote_subnet_2.
    - 
                                                                    Set Destination to To-HQ2_local_subnet_1.
    - 
                                                                    Set Schedule to Always.
    - 
                                                                    Set Service to All.
    - 
                                                                    Enable NAT.
    - 
                                                                    Set IP Pool Configuration to Use Outgoing Interface Address.
  - 
                                                                    
  - 
                                                            Click OK.
- 
                                                            
- 
                                                    Configure the Security Fabric: 
  - 
                                                            Go to Security Fabric > Fabric Connectors and double-click the Security Fabric Setup card.
  - 
                                                            Select the Settings tab, and set the Security Fabric role to Serve as Fabric Root.
  - 
                                                            Enter a Fabric name, such as Office-Security-Fabric.
  - 
                                                            Ensure Allow other Security Fabric devices to join is enabled and add VPN interface To-HQ2.
  - 
                                                            Click OK.
- 
                                                            
- 
                                                    Configure the FortiAnalyzer logging settings: 
  - 
                                                            Go to Security Fabric > Fabric Connectors and double-click the Logging & Analytics card.
  - 
                                                            Select the Settings tab, select the FortiAnalyzer tab, and set the Status to Enabled.
  - 
                                                            Enter the FortiAnalyzer IP in the Server field (192.168.8.250). The Upload option is automatically set to Real Time.
  - 
                                                            Click Refresh. The FortiAnalyzer serial number is verified.
  - 
                                                            Click OK.
- 
                                                            
To configure the downstream FortiGate (HQ2):
- 
                                                    Configure the interface: 
  - 
                                                            Go to Network > Interfaces.
  - 
                                                            Edit interface wan1: 
    - 
                                                                    Set Role to WAN.
    - 
                                                                    For the interface connected to the internet, set the IP/Network Mask to 192.168.7.3/255.255.255.0.
  - 
                                                                    
  - 
                                                            Edit interface vlan20: 
    - 
                                                                    Set Role to LAN.
    - 
                                                                    For the interface connected to local endpoint clients, set the IP/Network Mask to 10.1.100.3/255.255.255.0.
  - 
                                                                    
- 
                                                            
- 
                                                    Configure the static route to connect to the internet: 
  - 
                                                            Go to Network > Static Routes and click Create New or Create New > IPv4 Static Route. 
    - 
                                                                    Set Destination to 0.0.0.0/0.0.0.0.
    - 
                                                                    Set Interface to wan1.
    - 
                                                                    Set Gateway Address to 192.168.7.2.
  - 
                                                                    
  - 
                                                            Click OK.
- 
                                                            
- 
                                                    Configure the IPsec VPN: 
  - 
                                                            Go to VPN > IPsec Wizard. 
    - 
                                                                    Set VPN Name to To-HQ1.
    - 
                                                                    Set Template Type to Custom.
    - 
                                                                    Click Next.
    - 
                                                                    In the Network IP Address, enter 10.2.200.1.
    - 
                                                                    Set Interface to wan1.
    - 
                                                                    Set Authentication to Method.
    - 
                                                                    Set Pre-shared Key to 123456.
  - 
                                                                    
  - 
                                                            Leave all other fields in their default values and click OK.
- 
                                                            
