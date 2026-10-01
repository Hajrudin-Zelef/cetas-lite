---
id: collect-261001-general-networking/general-networking/platform-management-dashboard-administration-design-and-configure-architectures-efbeebe2-2
title: "platform-management-dashboard-administration-design-and-configure-architectures--efbeebe2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "throughput", "voice"]
source: docs/RAG/collect-261001-general-networking/platform-management-dashboard-administration-design-and-configure-architectures--efbeebe2.md
source_anchor: ""
source_lines: [81, 157]
sha256: 2a51f476306ca89b29d343ac01eafad017f80dd2b724d98a20b76f12d912dc6b
---

# platform-management-dashboard-administration-design-and-configure-architectures--efbeebe2

Furthermore, if an MX is configured for eBGP and receives a route that overlaps with our cloud connectivity network ranges, the MX’s cloud management traffic will follow that BGP route, so it is imperative that the MX, as well as it’s eBGP peer, have connectivity to everything listed on the Help > Firewall Info page in this scenario.
NAT Traversal
Finally, it is recommended to manually configure NAT traversal on a hub MX when it is in VPN concentrator mode behind an unfriendly NAT or aggressively timed CG-NAT device. The spokes that point to this hub will use the designated IP address and port, so ensure to use a public IP that is routable over the Internet. It is important that the upstream NAT device has a port forwarding rule to forward this traffic to the management IP address of this hub MX.
DC-DC Failover - Hub/DC redundancy (Disaster Recovery)
Summary
- Configure DC-DC Failover in order to protect against HW failure and DC outage / disaster.
Cisco Meraki's MX Datacenter Redundancy (DC-DC Failover) allows for network traffic sent across AutoVPN to failover between multiple geographically distributed datacenters.
DC Failover Architecture
A DC-DC failover architecture is as follows:
- One-armed VPN concentrators or NAT mode concentrators in each DC
- A subnet(s) or static route(s) advertised by two or more concentrators
- Hub & Spoke topologies
- Split or full tunnel configuration
Operation
Deploying one or more MXs to act as VPN concentrators in additional data centers provides greater redundancy for critical network services.
In a DC-DC failover design, a remote site will form VPN tunnels to all configured VPN hubs for the network. For subnets that are advertised from multiple hubs, spokes sites will send traffic to the highest priority hub that is reachable.
Concentrator Priority
Concentrator priority setting determines the order in which VPN mesh peers will prefer to connect to subnets advertised by multiple VPN concentrators.
Additional DC-DC integration data can be found in this article.
In a distributed deployment of locations connected via a site-to-site VPN, a network administrator may need to have address translation performed on traffic traversing the site-to-site VPN. A 1:1 subnet translation can be used in cases where multiple locations have the same subnet present, but both need to participate in the site-to-site VPN. Alternatively, administrators may need to conserve IP space for large deployments. For this, 1:M NAT can be used to translate entire subnets into a single IP address that is exported across the site-to-site VPN.
For additional information relating to VPN Subnet translation, please refer to this article.
MX Firmware Version
- Use the latest GA (may be different per platform)
The Cisco Meraki Dashboard allows admins to easily schedule and reschedule firmware upgrades on their networks, opt-in to beta firmware releases, view firmware changelog notes, and to set maintenance windows.
Users will only be able to upgrade to the general release and beta versions. Information about these versions can be found under Organization > Monitor > Firmware Upgrades.
MX Hub Sizing
- Sizing may change based on the traffic blend and other potential factors. It is also changing with the introduction of firmware improvements (the following is for MX 13).
- These configurations have been tested successfully with retail/restaurant locations deploying 1000-1500 tunnels* per hub with MX250/MX400 (closer to 1000) and MX450/MX600 (closer to 1500).
    
  - * Not locations, tunnels (think SD-WAN).
- Voice (and other small packet) traffic is notorious for high performance requirements and may result in a throughput and supported tunnel count that is lower than stated above.
Testing
This section captures key use cases identified to better test the MX in PoC environments.
The proposed topology for testing is detailed below. The Meraki SE and network admin will work together to refine this network architecture in the context of the POC success criteria agreed upon with the business.
The following tests should be performed:
- 
    AutoVPN Connectivity 
  - 
        Verify that AutoVPN works correctly on the Cisco Meraki MX Security appliance in a 100% Cisco Meraki environment. Use case is for Internet access, data center access.
  - 
        Ensure that solution works in full VPN and split-tunnelling configurations, delivering a ‘Branch-In-A-Box’ experience.
- 
        
- 
    AutoVPN Failover 
  - 
        Verify that MPLS (or other) fails over to AutoVPN successfully when the MPLS private WAN (or other) path fails. Make sure the MX has access to the Meraki VPN registries.
- 
        
Decommissioning Unused Sites
Meraki recommends that networks that have no further expected use be decommissioned from AutoVPN deployments by either disabling their VPN configurations, or by removing the devices in question from their networks. At larger scales, the Dashboard API can be used to ease this process.
The recommendation on this practice stems from the fact that previously learned points of contact from peers are not aged out of concentrators unless they’re rebooted. Over time - especially on concentrators that aren’t expected to have any periods of downtime - this can lead to unnecessary traffic being generated, as the concentrator reaches out to IP addresses and ports that are no longer in use, or even potentially in use by other networks.
- 
    Meraki SD-WAN (MPLS/Internet) 
  - 
        Verify that the Meraki SD-WAN service functions as designed and provides support for MPLS and Internet carriage simultaneously.
  - 
        The SD-WAN success relies on AutoVPN working correctly. Always check the communication to the VPN registry, specially when the MX has a DHCP address configured that can change. Each WAN has to reach the registry individually.
  - 
        For more information, refer to our SD-WAN Deployment Guide
- 
        
- 
    Transport Independence 
  - 
        Verify that transport independent links (e.g. MPLS, ADSL, etc) can be concurrently configured to support Auto-VPN overlay networks. Enable and configure multiple diverse uplink on the MX appliance. Configure flow preferences to pin traffic to a particular path, and/or load balancing.
- 
        
- 
    3G/4G WAN Failover 
  - 
        Verify that a failover USB 3G/4G interface can be installed, enabled and configured on the MX appliance and that traffic can be redirected over this link during a WAN interface failure condition. 
    - 
            Check the supported USB modems in our 3G/4G Cellular Failover article
  - 
            
-
