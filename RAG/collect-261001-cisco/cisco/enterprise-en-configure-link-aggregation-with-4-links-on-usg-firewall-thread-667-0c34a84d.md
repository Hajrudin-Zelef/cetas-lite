---
id: collect-261001-cisco/cisco/enterprise-en-configure-link-aggregation-with-4-links-on-usg-firewall-thread-667-0c34a84d
title: "enterprise-en-configure-link-aggregation-with-4-links-on-usg-firewall-thread-667-0c34a84d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-configure-link-aggregation-with-4-links-on-usg-firewall-thread-667-0c34a84d.md
source_anchor: ""
source_lines: [1, 21]
sha256: 4606e02e52e5527c16ce21099acfcc661f74e126f01c378cda49c189d447a3ef
---

# enterprise-en-configure-link-aggregation-with-4-links-on-usg-firewall-thread-667-0c34a84d

In this example, I will be showing how to configure link aggregation with 4 links on the USG series firewall with GUI.
Step 1:
Login to the firewall with username and password.
Step 2:
Choose the Network option and select Interface
Step 3:
Click on Add from the interface list option and add the following options.
Interface Name can be anything. (Choose own name)
Type should be Aggregation Interface.
The zone should be Trunk since we are configuring the Local network.
The mode should be Routing.
The interface can be selected from the list of interfaces available.
Assign a static IP for the aggregation.
Before choosing interfaces, make sure the Interface is in NONE Zone.
Once the interface is in NONE Zone only we can select the interface to configure aggregation.
Once selected all required interfaces, click on Ok to save.
Now the aggregated interface looks like this.
Step 4:
Commit and save the changes.
You can check the Interface link status by going to the Main menuàDevice information.
Or you can run Display eth-trunk command to check the status.
