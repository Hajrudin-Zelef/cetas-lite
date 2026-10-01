---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-16750417515159-unifi-gateway-setting-up-sd-wan-with-unifi-site-b04f4403
title: "hc-en-us-articles-16750417515159-unifi-gateway-setting-up-sd-wan-with-unifi-site-b04f4403"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-16750417515159-unifi-gateway-setting-up-sd-wan-with-unifi-site-b04f4403.md
source_anchor: ""
source_lines: [1, 84]
sha256: e14325a89d90973edffc14c79029edfb4e51f9d432623aa142c4af7cacad8743
---

# hc-en-us-articles-16750417515159-unifi-gateway-setting-up-sd-wan-with-unifi-site-b04f4403

UniFi Gateway - Setting Up SD-WAN with UniFi Site Manager and Fabrics
UniFi SD-WAN simplifies the setup of Site-to-Site VPN tunnels between UniFi Gateways, enabling seamless resource and application sharing across multiple sites.
Comparing Topologies
SD-WAN supports both Hub-and-Spoke and Mesh topologies. The table below highlights the key differences between these configurations.
|  | Hub & Spoke | Mesh | 
| Description | Centralized architecture where all branch sites connect through a central hub or multiple hubs. | Decentralized architecture where all branch sites connect directly to one another. | 
| Best Use-Cases | Ideal for organizations with centralized data centers or resources and those needing IP-based whitelisting for cloud access. | Suited for organizations with distributed resources that require direct sharing between all locations. | 
| Scalability | Supports up to 1,000 tunnels. | Supports up to 20 sites. | 
| Security | Centralized traffic management at hubs with isolation options. | Direct traffic flow between sites; requires individual firewall rules at each site. | 
| Redundancy | Each spoke supports up to 4 active VPN tunnels with the hub; failover hubs can be added for even more redundancy. | Each site has only one active tunnel per connection. Failover tunnels require time to re-establish during Internet outages. | 
| Flexibility | Highly customizable with features like NAT for overlapping subnets, custom routing advertisements, and load balancing across hubs. | Limited customization; NAT is unsupported, meaning overlapping subnets cannot exist within an SD-WAN group. | 
Hub & Spoke
Requirements
- 
Hub: At least one device with a public IP address:
  - Cloud Gateways: EFG, UDM Pro Max, UDM SE, UDM Pro, UCG Fiber, UCG Industrial, or UDW.
  - Independent Gateways: UXG-Fiber, UXG-Enterprise, or UXG-Pro managed with a CloudKey, UniFi OS Server or Official UniFi Hosting.
- Spoke: Most Cloud Gateways (excluding Express) or Independent Gateway managed with a CloudKey, UniFI OS Server, or Official UniFi Hosting.
- All hubs and spokes must share the same UI Account Owner or managed within the same Fabric by Fabric Admins.
- UniFi Network Application version 9.0.108 or newer.
- UniFi (Cloud) Gateway version 4.1.3 or newer.
Configuring Hub & Spoke
- Navigate to Settings > SD-WAN on the UniFi Site Manager.
- Select Hub & Spoke as the deployment type and name the SD-WAN group.
- 
Choose the Hub Topology:
  - Single: All spokes connect to the same central hub.
  - Failover: All spokes connect to the same central hub with failover to a backup hub.
  - Distributed: Manually specify which hub each spoke connects to.
- 
Select the Spoke-to-Hub VPN Architecture:
  - Max Resiliency: Up to 4 simultaneous VPN tunnels (independent tunnels between each Hub and Spoke WAN). No interruptions during failover.
  - Redundant: Supports up to 2 simultaneous VPN tunnels (each Spokes’ Primary and Secondary WAN connects to the Hub Primary, and Secondary, respectively). There will not be any interruptions during failover events.
  - Scalable: One tunnel per spoke. Temporary interruption during failover.
- Add Networks and/or Routes to each hub:
  - Networks: Automatically advertise routes for the specified networks.
  - Routes: Share non-local subnets (e.g., other Site-to-Site VPN connections) or manually define summary routes.
- Assign the Primary VPN WAN and (optional) WAN Failover for each hub.
- Configure Spoke Networks and WANs:
  - 
Auto-Scale and NAT Spoke VPNs: Enable when spokes have overlapping subnets. This automatically creates a Source NAT rule to translate traffic from a spoke into a unique /24 subnet before routing it to the hub. When enabled, sessions can only be initiated by the spoke. See Overlapping Subnets and NAT below to learn more.
    - If disabled, select the Networks you want to share with the Hub.
  - 
Isolate Spokes: Blocks all traffic between spokes by auto-generating Firewall rules on the hub. Disable this option and Auto-Scale if spokes need to communicate with each other.
    - Note: This requires Zone-Based Firewalling to be available on your gateway. View more details here.
  - Standardize WAN Settings: Makes it so all Spokes use the same WAN interface as the Primary or Failover VPN interface.
  - Specify the Primary VPN WAN and the WAN Failover for each hub.
- 
Auto-Scale and NAT Spoke VPNs: Enable when spokes have overlapping subnets. This automatically creates a Source NAT rule to translate traffic from a spoke into a unique /24 subnet before routing it to the hub. When enabled, sessions can only be initiated by the spoke. See Overlapping Subnets and NAT below to learn more.
Overlapping Subnets and NAT Configuration
Overlapping subnets typically prevent communication because traffic cannot be differentiated by origin or destination. Auto-Scale and NAT Spoke VPNs solves this by automatically applying a Source NAT rule, translating traffic to appear as if it originates from a unique, non-overlapping subnet. By default, this enables one-way communication from the spoke to the hub, allowing clients on a spoke to access hub resources, download files, or proxy to cloud resources. However, spoke resources will not be sharable unless you manually configure a Destination NAT rule on the spoke.
As an example, consider a spoke assigned the 172.16.1.0/24 route with Auto-Scale and NAT Spoke VPNs enabled, and a server at 192.168.50.5 on the local 192.168.50.0/24 subnet. You can make it accessible with the following DNAT rule applied to the Spoke:
- Name: DNAT from Hub-and-Spoke to LAN
- Protocol: All
- Interface: SD-WAN VPN Tunnel
- Destination: 172.16.1.5
- Translated IP Address: 192.168.50.5
A similar rule must be created for each resource on the spoke that needs to be accessible. If you need to expose a large number of local resources, we recommend designing subnets to avoid overlap. View instructions here.
To learn more about NAT rules, visit Network Address Translation.
Hub VPN Tunnel Capacity
| Model | SD-WAN VPN Tunnels | 
| Enterprise Firewall Core (EF Core) | 1,000 | 
| Enterprise Fortress Gateway (EFG) | 1,000 | 
| Gateway Enterprise (UXG Enterprise) | 1,000 | 
| Dream Machine Beast (UDM Beast) | 200 | 
| Dream Machine Pro Max (UDM Pro Max) | 200 | 
| Dream Machine Special Edition (UDM SE) | 100 | 
| Dream Machine Pro (UDM Pro) | 100 | 
| Cloud Gateway Fiber (UCG Fiber) | 100 | 
| Cloud Gateway Industrial (UCG Industrial) | 100 | 
| Dream Wall (UDW) | 100 | 
| Gateway Pro (UXG Pro) | 100 | 
Mesh
Requirements
- Any Cloud Gateway or Independent Gateway managed with a CloudKey, UniFi OS Server, or Official UniFi Hosting.
- At least one gateway must have a public IP address.
- All participating gateways must have the same UI Account Owner
Configuring Mesh
- Navigate to Settings > SD-WAN on the UniFi Site Manager.
- Select Mesh as the deployment type and name the SD-WAN group.
- Choose up to 20 site to be a part of the mesh connection.
- Select the networks from each site that will be shared.
  - If networks have overlapping subnets, follow the instructions here.
- Click Connect.
