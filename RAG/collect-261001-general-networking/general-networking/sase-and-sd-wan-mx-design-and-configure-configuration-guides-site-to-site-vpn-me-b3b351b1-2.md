---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1-2
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["latency", "training"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1.md
source_anchor: ""
source_lines: [27, 73]
sha256: 33a6f519965807bddde5006368e9ae80a5853761942aa61fe8e1756d77cb550a
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-site-to-site-vpn-me-b3b351b1

  - 43.192.139.128/25
  - 43.196.13.128/25
Ports used for IPsec tunneling:
- Source UDP port range 32768-61000
- Destination UDP port range 32768-61000
The VPN connection can be monitored under Security & SD-WAN > Monitor > VPN status page. The status of each WAN Appliance is displayed, along with their exported subnets, latency, connectivity and routing decisions that are being made over the Auto VPN domain in near real-time.
Auto VPN vs Non-Meraki Site-to-Site VPN
- Auto VPN is a VPN connection between/among the WAN Appliances in different networks of the same Meraki dashboard organization.
- Non-Meraki site-to-site VPN is used when you form a VPN tunnel with a third-party/non-Meraki device or when you establish a VPN connection with a Meraki WAN Appliance in a different dashboard organization.
- Like Non-Meraki Site-to-Site VPN, Auto VPN has encryption, authentication and a key. The traffic is encrypted using an AES cipher. However, all of this is transparent to users and does not need to be (and cannot be) modified.
Auto VPN: A Component of Meraki SD-WAN
| SD-WAN Characteristics | Meraki SD-WAN Component | 
|---|---|
| Support for VPNs | Meraki Auto VPN | 
| Multiple connection types (MPLS, Internet, LTE, etc.) | WAN Appliance uplink options allow for multiple connection type. | 
| Dynamic path selection (allows for load sharing across WAN connection) | WAN Appliances can perform uplink load balancing across WAN connections | 
| Simple WAN Configurations Interface (Must support zero-touch provisioning at a branch, should be easy to set up) | Meraki dashboard & API configuration interfaces | 
Auto VPN, as a component of SD-WAN, transitions the manual steps for setting the VPN tunnel into a simple automated process. It takes only a few clicks and makes it easy to deploy and manage an SD-WAN environment. It gives resilience, security and application optimization. It has automatic VPN route generation using the IKE/IPSec-like tunnels and all this is done in the Meraki cloud.
If you have two uplinks on your WAN Appliance, Auto VPN as a component of SD-WAN allows you to decide the flow preferences within the VPN tunnel under Security & SD-WAN > Configure > SD-WAN & traffic shaping page > Uplink selection > Multi-Uplink Auto VPN. Multi-Uplink Auto VPN allows you to create a VPN tunnel with flow preferences over both the uplinks.
If Multi-Uplink Auto VPN is disabled, the tunnel will be formed over the primary WAN link and will failover to the secondary if the primary fails.
Auto VPN Configuration
To enable site-to-site VPN between WAN appliances, simply login to the Meraki dashboard and navigate to the Security & SD-WAN > Configure > Site-to-site VPN page, and select Hub or Spoke and save the page. That's all that is required to enable VPN connectivity. Auto VPN takes care of all connection settings and brokers the connections immediately.
Note that Auto VPN is a simple opt-in process. You can think of the WAN Appliances dashboard organization an existing VPN hub and spoke mesh topology environment, and every WAN Appliance that has Auto VPN turned on is simply choosing to participate in that mesh. By default, all hubs contact all other hubs, and all spokes contact specified hubs. Additional configuration options can be found below.
Auto VPN Configuration Details
Enable Auto VPN by defining how the WAN Appliance will communicate with the rest of the Auto VPN domain
If the WAN Appliance is configured as a Hub, it will build VPN tunnels to all other Hub WAN Appliances in the Auto VPN domain (in the same dashboard organization). It will also build VPN tunnels to all Spoke WAN Appliances in the Auto VPN domain that have this WAN Appliance configured as a hub. If all WAN Appliances in the Auto VPN domain are configured as Hub then the Auto VPN has a full mesh topology.
If the WAN Appliance is configured as a Spoke, it will build tunnels to only the WAN Appliances that are configured as its Hubs. If the majority of WAN Appliances in the Auto VPN domain are configured as Spoke with only a few key locations (such as data centers or headquarters) configured as hubs, then the Auto VPN environment has a hub-and-spoke topology.
Full Tunnel or Split Tunnel
By default all WAN Appliances in the Auto VPN domain (dashboard organization) will only send traffic to an Auto VPN peer if the traffic is destined for a subnet contained within the Auto VPN domain. This is often referred to as 'split-tunnelling,' meaning that VPN-subnet-bound traffic is sent over VPN, and other traffic is routed normally via the primary WAN Appliance WAN uplink. If an organization wants to route all traffic (including traffic not contained within the Auto VPN domain) through a specific hub site, this is referred to as 'full-tunneling.'
Note that full-tunneling only affects client data and all Meraki management traffic will egress directly via the primary WAN regardless.
To configure full-tunneling in a full mesh topology simply define an Exit hub from the WAN Appliances in the Auto VPN domain.
To configure full-tunneling in a hub-and-spoke topology, simply associate a ‘Default route’ with one or more hub WAN Appliances:
Choose which subnets (local networks) to export over VPN
Earmark which locally defined or available subnets are to be exported to the Auto VPN domain. Select Enabled under VPN Mode to define available subnets over the VPN.
Then save the changes so the WAN Appliance fetches the configurations from the cloud.
Troubleshooting Auto VPN
Learn more with these free online training courses on the Meraki Learning Hub:
Configurations that Use Automatic NAT Traversal
- Site-to-site VPN: When MX Series WAN appliances are configured to create a VPN between two or more sites within the same Organization.
- Teleworker VPN: When MR Series access points have an SSID configured to tunnel all traffic back to a central concentrator. This concentrator can be either an MX Series WAN appliance in passthrough/concentrator mode, or a VM (Virtual Machine) concentrator.
Troubleshooting VPN Registration
When using VPN functionality to securely tunnel traffic between Cisco Meraki devices, such as the WAN Appliance Site-to-site VPN, or Access Point Teleworker VPN, the devices must first register with the Dashboard VPN registry. This allows their connections between each other to be dynamic, and automatically establish without manual configuration. However, sometimes issues can occur with this process, which will be discussed in this article.
In order to ensure connectivity, each Meraki node sends a keepalive message to the VPN Registry every 10 seconds. If more than 6 keepalives are not received by the registry, that node is marked as disconnected. For information on how connectivity to the VPN Registry works, please read the article on Automatic NAT Traversal.
Both Meraki peers must be in communication with the VPN registry in order to get the correct information to form a valid VPN tunnel. If one Meraki device, such as an MX WAN appliance, is able to reach the VPN registry, but the intended peer WAN Appliance is not, the tunnel will not form. A common occurrence of this is when an upstream firewall blocks VPN registry communication on UDP port 9350-9381. This issue is explained in the section VPN Registry Disconnected.
If the appliance/concentrator is successfully connected to the VPN registry, but is disconnected from another VPN peer, refer to the article on troubleshooting VPN connections between peers.
VPN Registry Disconnected
When the "VPN Registry: Disconnected" message appears on the Security & SD-WAN > Monitor > VPN status page for WAN Appliance networks, it indicates that the appliance has been unable to establish connectivity with the VPN registry. This means that a firewall or other upstream device is either preventing traffic from reaching the VPN registry, or from returning to the appliance.
