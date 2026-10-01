---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-94f1c4e2-2
title: "platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2.md
source_anchor: ""
source_lines: [77, 111]
sha256: 50dc9472fe9071574df258f047b1263311305a2ed33d0f4ac353de535d189c14
---

# platform-management-dashboard-administration-design-and-configure-architectures--94f1c4e2

BGP VPNs are utilized for Data Center Failover and load sharing. This is accomplished by placing VPN Concentrators at each Data Center. Each VPN Concentrator will utilize BGP with DC edge devices. BGP is utilized for its scalability and tuning capabilities.
More information about implementing BGP and its use cases can be found in our BGP documentation.
Auto VPN Technology Deep Dive
The Meraki WANAppliance makes use of several types of outbound communication. Configuration of the upstream firewall may be required to allow this communication.
Dashboard & Cloud
The Meraki WAN Appliance is a cloud managed networking device. As such, it is important to ensure that the necessary firewall policies are in place to allow for monitoring and configuration via the Meraki dashboard. The relevant destination ports and IP addresses can be found under the Help > Firewall Info page in the dashboard.
VPN Registry
Meraki's Auto VPN technology leverages a cloud-based registry service to orchestrate VPN connectivity. In order for successful Auto VPN connections to establish, the upstream firewall must allow the VPN concentrator to communicate with the VPN registry service. The relevant destination ports and IP addresses may vary by region, and can be found under the Help > Firewall Info page in the dashboard.
Uplink Health Monitoring
The WAN Appliance also performs periodic uplink health checks by reaching out to well-known Internet destinations using common protocols. The full behavior is outlined here. In order to allow for proper uplink monitoring, the following communications must also be allowed:
- 
    DNS test to canireachthe.net
- 
    Internet test to icmp.canireachthe.net
VPN Registry
In order to participate in Auto VPN a WAN Appliance must register with the Meraki VPN registry. The VPN registry is a cloud-based system that stores data needed to connect all WAN Appliances into an orchestrated VPN system. The VPN registry is always on and always updating in the case of a connection failure. This means no manual intervention is needed in the case of reboots, new public IP addresses hardware failovers etc. The VPN registry stores the following information for each WAN Appliance:
- 
    Subnets (for creating the VPN route table)
- 
    Uplink IP (public or private)
- 
    Public IP
The process for adding a new WAN Appliance into an infrastructure is as follows:
- 
    A new WAN Appliance reports its uplink IP address(es) and shared subnets to the registry
- 
    The information is propagated to the other WAN Appliances in the infrastructure
- 
    The WAN Appliance establishes the proper VPN tunnels 
  - 
        The WAN Appliance will try the registry-reported private uplink IP of the peer first
  - 
        If a connection to the private uplink IP of the peer fails, the WAN Appliance will try the public uplink IP of its peer
- 
        
