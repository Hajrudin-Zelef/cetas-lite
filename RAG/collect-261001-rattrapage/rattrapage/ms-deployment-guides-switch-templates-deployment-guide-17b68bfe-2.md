---
id: collect-261001-rattrapage/rattrapage/ms-deployment-guides-switch-templates-deployment-guide-17b68bfe-2
title: "ms-deployment-guides-switch-templates-deployment-guide-17b68bfe"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/ms-deployment-guides-switch-templates-deployment-guide-17b68bfe.md
source_anchor: ""
source_lines: [50, 90]
sha256: ecfb849621e6bae39a06b571873d5097aca522bcd6ac5c98aa8ca05eaca5c016
---

# ms-deployment-guides-switch-templates-deployment-guide-17b68bfe

  - If no templates have been created, a Create switch template window will appear. Otherwise, click Create switch template.
- Select a descriptive name and the switch model to be configured.
- Click Save:
- Click on the newly created template.
- To configure the template, select View ports on this switch template.
- The following port configuration page can be used identically to the normal switch port configuration page:
- Navigate back to the switch template under Switching > Configure > Switch templates > Template name.
- Click Bind switches.
- Select one or more switches, then click Bind to profile:
All bound switches will now use the port configuration set in the switch templates. Any changes made in the template will now affect all bound switches.
Local Overrides
Once a switch has been bound to a template, it can still be configured normally through Dashboard. Any port configuration changes made directly on the switch will override the template configuration, and be reported as a local override. If there is a need to override template configuration on a large scale, it is recommended to keep the switch port count (with local overrides) per network up to 2500. For better performance on loading the switchport page, reduce the count to 2000.
In the example below, the bound switch was directly configured to have a custom VLAN set on port 3. In the template network, under Switching > Configure > Switch templates, this configuration change is shown under the Local overrides column:
Auto-Binding Switches
When a new network is bound to an existing template with at least one switch template, the option is available to "auto-bind" switches to that template's templates. This dramatically reduces the amount of work necessary to set up a new site; if templates exist for every switch model to be deployed, the only necessary configuration is binding the network to the template.
To auto-bind switches in a new network:
- Create a new switch network.
- Add devices to the network as normal.
- Navigate to Organization > Configuration templates > Template name.
- Click Bind additional networks.
- Select the newly created network, and check Auto-bind target devices.
- Click Bind:
All switches in the new network will now automatically be bound to their appropriate templates.
Auditing a Template Deployment
Though templates can be used to consolidate most switch configuration, there may be exceptions for individual ports or settings that necessitate local configuration changes on the switch. Since local configuration changes override template and template configurations, it is important to keep track of any local configuration changes across the organization.
To view local configuration overrides of a template:
- Navigate to Organization > Configuration templates > Template name.
- The Local overrides column will show if any networks are overriding the template configuration:
To view local configuration overrides of a switch template:
- Select the appropriate template from the network drop-down.
- Navigate to Switching > Configure > Switch templates > Template name.
- The Local overrides column will show if any switches are overriding the template configuration:
Managing Scheduled Firmware Upgrades in Template Deployments
Firmware Updates are scheduled for a Parent template network but occur at the individual network level. The time zone used for these updates is determined by the Local time zone setting found under Network-wide > Configure > General for each specific network.
When a new network is created and bound to a template, or when an existing network is bound to a template, its Local time zone setting will initially inherit the value configured in the parent template network.
Therefore, if your template deployment includes switches across multiple time zones and you wish to perform Scheduled Firmware Updates in the local time zone of each site, you must explicitly configure the Local time zone for each individual bound network. This ensures that each network has its correct local time zone, independent of the template, allowing for precise local scheduling of firmware updates. To achieve this, your deployment should be structured such that switches belonging to different time zones are located in separate networks (e.g., a unique network per site location, or a unique network per region assuming the devices in that region are all in the same time zone).
To configure the local time zone for a specific network:
- Select the appropriate network (the individual site or region network, not the parent template network) from the network drop-down menu in Dashboard.
- Navigate to Network-wide > Configure > General.
- Under Local time zone, select the correct Time Zone for the devices in the selected network:
 
