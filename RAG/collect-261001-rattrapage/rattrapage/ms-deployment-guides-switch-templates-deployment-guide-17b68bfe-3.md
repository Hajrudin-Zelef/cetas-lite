---
id: collect-261001-rattrapage/rattrapage/ms-deployment-guides-switch-templates-deployment-guide-17b68bfe-3
title: "ms-deployment-guides-switch-templates-deployment-guide-17b68bfe"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/ms-deployment-guides-switch-templates-deployment-guide-17b68bfe.md
source_anchor: ""
source_lines: [91, 160]
sha256: 9a256e057243f582305748e24d5c99167eb0e90df3a6015e26a96deb2c280c93
---

# ms-deployment-guides-switch-templates-deployment-guide-17b68bfe

- Click Save Changes.
Switch Replacement Walkthrough for Stacks
Below are instructions for how to copy configurations from a failed switch that is part of a stack and where the network is bound to a template.
- 
    On the Organization > Configure > Inventory page, claim the new switch and then add the new switch to the existing network.
- Bind new switch to the switch template.
- 
    Navigate to the parent template in dashboard.
- 
    Navigate to Switching > Configure > Switch templates within that template.
- 
    Click on the corresponding template.
- 
    Click the Bind switches button.
- 
    Click the checkbox next to the new switch and click the Bind to switch template button.
- Firmware upgrade for the new switch.
- 
    Provide the new switch a physical uplink connection and then power it on. The new switch needs to be brought online as a standalone device, not yet added to the stack so that it can update its firmware.
- 
    Confirm via the connectivity graph or Support that the switch has upgraded its firmware.
- 
    While the new switch upgrades, you may proceed with the below steps, stopping before Step 8 until the new switch has had a chance to upgrade.
- Obtain Current Configuration.
- 
    Navigate to Switching > Configure > Switch templates within the parent template.
- 
    Click on the template in question.
- 
    Filter in the Search switches… field for the name of the old switch.
- 
    Note the local override configuration. Save in a text editor for use in Step 5.
- 
    In the child network, navigate to the Switching > Monitor > Switch ports page.
- 
    In the Search switches… field, filter by the name of the old switch and select the below column options.
- Then, take screenshots of the port configurations or copy and paste into a spreadsheet or text editor application.
- Configure replacement switch.
- 
    On the Switching > Monitor > Switch ports page of the child network, configure the switch ports of the new switch based on the configuration gathered in Step 4.
- 
    Once complete, navigate back to the switch template details page from Step 4 and ensure that the local overrides between the old and new switch match.
- Power down the old switch.
- Unbind Old Switch from template.
- On the switch template details page, click the check box next to the old switch and then click the Unbind button.
- Add the new switch to the stack.
- 
    After confirming that the new switch has upgraded its firmware as mentioned in Step 3, power down the new switch.
- 
    In the child network, navigate to the Switching > Monitor > Switch stacks page.
- 
    Click on the stack in question.
- 
    Click the Manage members tab.
- 
    Under Add members, click the checkbox next to the new switch and then click the Add switches button.
- Remove the old switch from the stack.
- 
    In the child network, navigate to the Switching > Monitor > Switch stacks page.
- 
    Click on the stack in question.
- 
    Click the Manage members tab.
- 
    Click the checkbox next to the old switch and then click the Remove switches button.
- Physically cable and stack the new switch.
- Power on the new switch.
Additional Notes and Resources
If many networks are being deployed at once, the bulk network creation tool can be used to bind templates in bulk.
Please reference our documentation for more information on Meraki configuration templates.
