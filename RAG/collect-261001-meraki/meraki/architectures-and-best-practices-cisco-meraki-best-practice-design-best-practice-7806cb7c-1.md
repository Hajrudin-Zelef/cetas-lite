---
id: collect-261001-meraki/meraki/architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c-1
title: "architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-meraki/architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c.md
source_anchor: ""
source_lines: [1, 160]
sha256: e55196fe562001c62f3f0337762281e418cb204c044bf750526dafaa4720de5b
---

# architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c

MX Templates Best Practices
Click 日本語 forJapanese
As a network deployment grows to span multiple sites, managing individual devices can become highly cumbersome and unnecessary. To help alleviate these operating costs, the Meraki WAN Appliance offers the use of templates to quickly roll out new site deployments and make changes in bulk.
This guide will outline how to create and use WAN Appliance templates on the dashboard.
It should be noted that service providers or deployments that rely heavily on network management via API are encouraged to consider cloning networks instead of using templates, as the API options available for cloning currently provide more granular control than the API options available for templates.
Planning a Template Deployment for WAN Appliances
Before rolling out a template deployment (or enabling templates on a production network), it may be helpful to plan the "units" that make up your deployments. This involves asking questions such as:
- 
    What are my sites? (e.g. retail location, school, branch office, etc.)
- 
    Are the WAN Appliances going to be in HA?
- 
    Do I need local overrides?
Template Networks
A "site" in network deployment terms is usually the same as a "network" in dashboard terms; each site gets its own dashboard network. As such, when planning multiple sites to be configured the same way, they will share a template network.
A template network is a network configuration that is shared by multiple sites/networks. Individual site networks can be bound to a template network, so changes to the template will trickle down to all bound sites. A new network can also be created based on a template, making it easy to spin-up new sites of the same type.
When planning a template deployment, you should have one template network for each type of site.
Configuration
The following sections walk through the configuration and use of WAN Appliance templates in the dashboard.
Creating a Template Network
As outlined above, a template network should be created for each type of site to be deployed.
To create a template network:
- 
    In the dashboard, navigate to Organization > Monitor > Configuration templates
- 
    Choose Create a new template
- 
    Select a descriptive name for your template. If this is a completely new template, select Create new 
  - 
        If this template should be based on an existing network, select Copy settings from and select an existing Security appliance network from the drop-down menu.
- 
        
- 
    Choose Add:
- 
    If you would like to bind existing networks to this new template, select those networks as Target networks and choose Bind. Otherwise, choose Close.
Template VLAN Configuration
- 
    In the dashboard, navigate to Security & SD-WAN > Configure > Addressing & VLANs
- 
    Under Routing section, LAN setting sub-section click VLANs
- 
    Choose Add VLAN under Subnets sub-section
- 
    Select a descriptive name for your VLAN
- 
    Choose whether the subnetting should be Same or Unique for every network bound to this template. 
  - 
        If Same is chosen, all the networks bound to the template will share the exact same subnet. This is not eligible for site-to-site VPNs.
  - 
        If Unique is chosen, each network bound to the template will get a unique subnet based on the configured options. The MX does allow local VLAN overrides on templates, however, the chosen subnet needs to be from the same subnet pool assigned to the VLAN on the template and you can't override the VLAN ID. 
    - 
            Subnets are assigned randomly to each network bound to the template.
  - 
            
- 
        
For more information about template IP range VLAN allocation, reference our article on Managing Networks with Configuration Templates.
Template Static Routes
In template-based WAN Appliance deployments, static routes can be configured on the parent template and passed to child networks like other configuration parameters. The procedure for configuring a template-based static route is almost identical to the procedure for a regular network, with the exception of how next-hop IP addresses are defined as the next-hop value may be network specific.
- 
    In the dashboard, navigate to Security & SD-WAN > Configure > Addressing & VLANs > Routing > Static Routes
- 
    Choose Add Static Route
- 
    Name 
  - 
        Text description for the static route(not parsed) 
    - 
            Ex: prodWirelessNet
  - 
            
- 
        
- 
    Subnet 
  - 
        Subnet reachable via static route specified in CIDR notation 
    - 
            Ex: 10.0.10.0/24
  - 
            
- 
        
- 
    Next Hop IP 
  - 
        Next-hop IP is the IP address of the device that connects the WAN Appliance to this route. There are two methods for specifying next-hop values on template-based networks. 
    - 
            Option A: IP Assignment 
      - 
                Manually define next-hop value
      - 
                Required Info: 
        - 
                    Next-hop IP address
      - 
                    
    - 
                
    - 
            Option B: IP offset 
      - 
                Calculate next-hop IP based on network address for specified VLAN 
        - 
                    NOTE: Next Hop IP will be calculated as Network Address + Offset and not VLAN Interface IP + Offset
      - 
                    
      - 
                IP offset parameters: 
        - 
                    Select the desired VLAN from the dropdown
        - 
                    Offset value (a positive integer)
      - 
                    
      - 
                Example: 
        - 
                    VLAN configured: 10.0.254.0/30
        - 
                    VLAN Interface IP: 10.0.254.1
        - 
                    Offset: 2
        - 
                    Calculated route next-hop IP: 10.0.254.2
      - 
                    
    - 
                
  - 
            
- 
        
- 
    Active 
  - 
        The active modifier controls conditions that must be met for the WAN Appliance to deem the route usable and add the route to the local routing table. 
    - 
            Always: 
      - 
                The route will always be active in WAN Appliance's routing table
    - 
                
    - 
            While next-hop responds to ping: 
      - 
                The route is available as long as the configured next hop is responding to pings
    - 
                
    - 
            While host responds to ping: 
      - 
                The route is available as long as the configured host is responding to pings
    - 
                
  - 
            
- 
        
