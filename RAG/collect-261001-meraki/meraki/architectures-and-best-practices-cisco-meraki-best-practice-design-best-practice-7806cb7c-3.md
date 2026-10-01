---
id: collect-261001-meraki/meraki/architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c-3
title: "architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c.md
source_anchor: ""
source_lines: [246, 287]
sha256: 72d04cbff8947ad5251cfa1a15623cc27709c36df121a9ef789e857008ca98b2
---

# architectures-and-best-practices-cisco-meraki-best-practice-design-best-practice-7806cb7c

- 
    Once you are done with configuring the criteria to apply the policy and the policy, choose Save
Note:
The "add host" button under Security & SD-WAN> Configure > Traffic shaping > Flow preferences, gives an option to enter a value between 1-254. The following example illustrates this behavior:
If we have a subnet /26 from 10.0.0.0/8, there would be 4 possible 4th octets: x.x.x.0/26 = .0-.63
x.x.x.64/26 = .64-.127
x.x.x.128/26 = .128-.191
x.x.x.192/26 = .192-.255
Since the template needs to be applicable to ALL networks tied to it, it uses an offset.
If we were to specify the .1 host, this would be the equivalent of .1, .65, .129, and .193 depending on the given network tied to the template.
Local Overrides
Once a WAN Appliance network has been bound to a template, some options can still be configured normally through the dashboard. Any local configuration changes made directly on the WAN Appliance network will override the template configuration.
In the example below, the bound WAN Appliance was directly configured to have a custom Default VLAN. This change can be made in the template network, under Security & SD-WAN > Configure > Addressing & VLANs:
If a network is removed from a template, local overrides will automatically be lost as well as any template related configuration. The WAN Appliance will automatically get the configuration from the network it is on.
Note: Auto VPN hubs should not be added to templates. Dashboard will not allow a WAN Appliance configured as a spoke to be peered with a hub in a network bound to a template. When configuring a spoke (bound or not bound to a template), hubs in templates are not available(exit hub or not) in the hub selection section. However, hubs in templates can still form AutoVPN meshes with other hubs in the same Organization.
Note: Static Route local overrides are not supported at this moment for WAN Appliance networks bound to templates.
DHCP Exceptions
The Meraki WAN Appliance provides a fully-featured DHCP service that can be enabled and configured on each VLAN individually. When bound to a template, local overrides can be made to the DHCP configurations under Security & SD-WAN > Configure > DHCP.
Forwarding Rules Overrides
To override forwarding rules, navigate under Security & SD-WAN > Configure > Firewall > Forwarding rules overrides.
Active-Active AutoVPN
To override the uplink selection rules for Active-Active AutoVPN, navigate to Security & SD-WAN > Configure > SD-WAN & traffic shaping > Active-Active AutoVPN.
Templates with MXs of Different Port Counts
You can toggle the LAN2 port between LAN and Internet, through Uplink configuration under the Local status tab on the Local Status Page.
Client VPN
To manage client VPN users, navigate under Network-wide > Configure > Users
Note: To manage Client VPN users across all networks bound to a template, you can do so in the User Management section of the Security & SD-WAN > Client VPN page of said template. If you would like to manage Client VPN users for a specific network bound to that template, you can do so in the Network-wide > Configure > Users page of that network
For more detailed guidance on the above Client VPN setup, please refer to the Client VPN Overview document.
Performing MX Templates Firmware Upgrades
Firmware upgrades scheduled on the template will automatically be applied to the child networks’ network local timezone.
As a best practice, make sure that each MX has the correct local time zone configuration under Network-wide > Configure > General.
WAN Appliance Replacement Walkthrough
Below are instructions for how to copy configurations from a failed WAN Appliane bound to a template.
- 
    On the Organization > Configure > Inventory page, claim the new WAN Appliance.
- 
    Navigate to the network that has the faulty WAN Appliance and remove it under Security & SD-WAN > Monitor > Appliance Status > Remove appliance from network
- 
    Add the replacement WAN Appliance to the same network by navigating to Network-wide > Configure > Add devices
- 
    Select the network and choose Add devices.
For more information on replacing a WAN Appliance, refer to our MX Cold Swap article.
