---
id: collect-261001-rattrapage/rattrapage/ms-deployment-guides-switch-templates-deployment-guide-17b68bfe-1
title: "ms-deployment-guides-switch-templates-deployment-guide-17b68bfe"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/ms-deployment-guides-switch-templates-deployment-guide-17b68bfe.md
source_anchor: ""
source_lines: [1, 49]
sha256: e14ea8463dfb8e1304abb809be0b82db29114e9a8caec9a8a79feefc07346aea
---

# ms-deployment-guides-switch-templates-deployment-guide-17b68bfe

Templates for Switching Best Practices
As a network deployment grows to span multiple sites, managing individual devices can become highly cumbersome and unnecessary. To help alleviate these operating costs, the Meraki MS switch offers the use of templates to quickly roll out new site deployments and make changes in bulk.
This guide will outline how to create and use MS switch templates in Dashboard.
Planning a Template Deployment
Before rolling out a template deployment (or enabling templates on a production network), it may be helpful to plan the "units" that make up your deployments. This involves asking questions such as:
- What are my sites? (e.g. retail location, school, branch office, etc.)
- How many switches are at each site?
- Are multiple switches at the same site configured the same way? (e.g. access switches, classroom switches, etc.)
The answers to these questions should directly affect your template deployment, specifically your use of template networks and switch templates.
Layer 3 Routing & DHCP
Layer 3 settings for networks bound to a template act as exceptions to the template. The Routing & DHCP, OSPF routing, and DHCP servers & ARP pages will be configurable on each network bound to a template and behave the same as if the network was not bound to a template. The one difference is that setting email alerts for newly detected DHCP servers on the DHCP servers & ARP page is available from the parent template's Network-wide > Configure > Alerts page and therefore applies to all networks bound to the template.
Template Networks
A "site" in network deployment terms is usually the same as a "network" in Dashboard terms; each site gets their own Dashboard network. As such, when planning multiple sites to be configured the same way, they will share a template network.
A template network is a network configuration that is shared by multiple sites/networks. Individual site networks can be bound to a template network, so changes to the template will trickle down to all bound sites. A new network can also be created based on a template, making it easy to spin up new sites of the same type.
When planning a template deployment, you should have one template network for each type of site.
Switch Templates
If multiple switches on a site share the same port configuration, they can easily be deployed and updated using switch templates. Within a template network, a switch template defines the per-port configuration for a group of switches. For example, if all sites contain multiple MS220-24 switches that are all configured identically, an administrator can set up a switch template for their MS220-24 switches. This way, whenever a new MS220-24 is installed at a site, it will automatically assume its switch template and configure its ports accordingly.
When planning a template deployment, if multiple switches share the same configuration, consider using a switch template for each type of switch.
Templates can be overwritten by a local port configuration. If the port of a template-bound switch is manually configured, that manual configuration will be "sticky" and remain on the switch, even if the template is changed.
Different MS switch models cannot share the same template.
Guidelines and Limitations
1. STP bridge priority cannot be changed on switch stacks using templates. In a network template, switch templates can be assigned STP bridge priority values from Switching > Configure > Switch settings > STP configuration. A value assigned to a switch template, however, will only propagate to the standalone switches bound to that template; switch stacks will retain the default STP priority of 32768.
If an MS switch stack must be the root bridge of an STP domain that has at least one other MS switch stack, it should be placed in a Dashboard network that is not bound to a network template for the STP priority value configured on it to take effect.
The following diagram shows the relationship between network templates and switch templates:
Configuration
The following sections walk through configuration and use of switch templates in dashboard:
Creating a Template Network
As outlined above, a template network should be created for each type of site to be deployed.
To create a template network:
- In dashboard, navigate to Organization > Monitor > Configuration templates.
- Click Create a new template.
- Select a descriptive name for your template. If this is a completely new template, select Create new and Switch template.
    
  - If this template should be based on an existing network, select Copy settings from and an existing switch network.
- Click Add:
- If you would like to bind existing networks to this new template, select those networks as Target networks and click Bind. Otherwise, click Close.
- Make sure you click on the "Save changes" button at the bottom of the page.
Multiple device types can be managed within a single template, acting as a combined network. For more information on managing non-switch devices in a template, refer to our documentation.
Once a network has been bound to this template, the template network should appear in the Network drop-down:
Configuring a Template Network
Once a template has been created, it can be configured with some global switch settings. These settings will also apply to all networks bound to this template.
To configure a template network, select the template from the network drop-down and configure settings normally under the Switching > Configure menu options. This can include setting a global management VLAN, IPv4 ACLs, port schedules, etc. Switch templates can also be configured here, as detailed below.
Please note that any configuration changes made here will apply to all bound networks.
Creating and Using Switch Templates
Switch templates can be used to bulk configure switches of the same model.
To create a switch template:
- Select the appropriate network template from the network drop-down.
- Navigate to Switching > Configure > Switch templates.
    
