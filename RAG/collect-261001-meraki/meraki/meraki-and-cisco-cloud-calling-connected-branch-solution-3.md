---
id: collect-261001-meraki/meraki/meraki-and-cisco-cloud-calling-connected-branch-solution-3
title: "Appendix"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput", "voice"]
source: docs/RAG/collect-261001-meraki/meraki-and-cisco-cloud-calling-connected-branch-solution.md
source_anchor: ""
source_lines: [257, 403]
sha256: 31ddb63338f6bd6b38bea486614f337ee4f41794951717c6c253834edac298d6
---

# Appendix

**Step 6**. Associate the dedicated voice SSID to the voice VLAN configured in the LAN by enabling **VLAN tagging** and specifying the ID for the voice VLAN used by the Meraki MX Security Appliance and Meraki MS switches.



**Step 7**. Scroll down to **Wireless options** and set the minimum bit rate to **12 Mbps**. This is the recommended minimum bit rate for wireless networks with VoIP traffic. Disabling lower bit rates reduces overhead on the network and can result in an improved roaming experience.



**Note:** Setting the bit rate to 12 Mbps and above will prevent 802.11b clients from joining the wireless network.


**Step 8**. Navigate to the **Wireless** tab and select **Firewall & traffic shaping** and ensure that the voice-dedicated SSID is selected from the drop-down.

**Step 9**. Scroll down to **Traffic shaping rules** and create a new rule for **All VoIP & video-conferencing** traffic and set the **Per-client bandwidth limit** to **Ignore SSID per-client limit (unlimited)**.

**Step 10**. Set the **PCP** and **DSCP** tags to **6** and **46 (EF - Expedited Forwarding, Voice)** respectively.


**Note:** The DSCP tag **46 (EF - Expedited Forwarding, Voice)** maps to WMM Access Category AC-VO for Voice, Layer 2 CoS 6).


For additional information, refer to the __Wireless VoIP QoS Best Practices document__.


## **Meraki Insight**

The **Meraki Insight** product is designed to give Meraki customers an easy way to monitor the performance of **Web Applications** and **WAN Links** on their network and easily identify if any issues are likely being caused by the network or application. The goal of Meraki Insight is to provide end-to-end visibility to customers and make sure they have assurance for the mission-critical traffic of the network.


### Enabling and Disabling Meraki Insight

**To Enable Meraki Insight on a network:**

**Step 1.** Ensure the necessary licensing is available.

**Step 2.** Select the checkbox next to the network where Insight should be enabled.

**Step 3.** Click **Add network(s) to Insight** at the top left of the table to enable Meraki Insight on the selected network(s).


**To Disable Meraki Insight on a network:**

**Step 1.** Select the checkbox next to the network with Insight currently enabled.

**Step 2.** Click **Remove network(s) from Insight**

### Configuring Meraki Insight for VoIP

**After** a dashboard account and organization have been created:

**Step 1**. Log in to the Cisco Meraki dashboard.

**Step 2**. Select **VoIP Health** from the **Insight** tab.

**Step 3**. Select **Configure VoIP servers** on the pop-up.

**Step 4**. Input the **Provider name** and the **Hostname or IP**. Repeat this step to add additional VoIP servers.

**Note:** VoIP Health supports monitoring for on-premises VoIP servers that are available on the LAN or over VPN.

### Monitoring Meraki Insight for VoIP

Once you provide the servers you’d like to monitor, you’ll be able to see an org-wide view of performance:


**Step 1**. Select **VoIP Health** from the **Insight** tab.

Clicking any row will take you to a drill-down view with detailed performance information:


For further details regarding the capabilities of Meraki Insight, see the __Meraki Insight Introduction article__.

## Automation of Network Provisioning

The Meraki dashboard offers a robust set of APIs enabling automated network provisioning and deployment. Prior to utilizing the APIs, they must be enabled for an account, and a specific user must have an API key. In general, it is recommended that a dedicated dashboard administrator account be created for use of APIs, and the API key generated for that user.

### Initial Setup

In order to automate the network provisioning, some initial steps must be performed:



**Step 1**. Create the base organization (see **Creating and Configuring the Base Organization** above).

**Step 2**. Locate the organization ID for the base organization. 



**Step 3**. Configure the provisioning tool with the located organization ID.

### Customer Provisioning

Once the provisioning system has been configured, new customers can be deployed using the following steps:



 

**Step 1.** New customer order is placed, automated network provisioning begins.

**Step 2**. Create a new organization, cloning from the base organization.

**Step 3**. Locate the network ID for the network that corresponds to the blueprint network.



**Step 4**. Create a new network by cloning from the blueprint network.



**Step 5**. Add hardware to newly created network.



**Step 6**. Configure switch ports for voice services (access port and voice VLAN). As a general recommendation, you should utilize the __Action Batches API documentation__ to perform this task.



 

# Appendix

**Table 1 - Cisco Meraki devices technical specifications and sizing guide**


|  | **Micro Site**  **(Up to 7)** | **Small Site**  **(Up to 25)** | **Medium Site (25-100)** | **Large Site (100-250)** | 
| **Security & SD-WAN (MX)** |  |  |  |  | 
| MX Model | MX68CW (Branch in a Box) | MX67/MX67C | MX84 | MX100 | 
| Stateful Firewall Throughput | 450 Mbps | 450 Mbps | 500 Mbps | 750 Mbps | 
| Advanced Security Throughput | 300 Mbps | 300 Mbps | 320 Mbps | 650 Mbps | 
| WAN Interfaces (dedicated) | 2 x GbE RJ45 | 1 x GbE RJ45 | 2 x GbE RJ45 | 2 x GbE RJ45 | 
| Dual-Purpose (WAN/LAN) | - | 1 x GbE RJ45 | - | - | 
| LAN Interfaces | 10 x GbE RJ45 (2 x PoE+) | 4 x GbE RJ45***** | 8 x GbE RJ45 2 x SFP | 8 x GbE RJ45 2 x SFP | 
| LTE Failover | Yes | Yes | Yes (USB) | Yes (USB) | 
| *** One of the LAN ports is optionally available for WAN connectivity** |  |  |  |  | 
| **Switching (MS)** |  |  |  |  | 
| MS Model | MS120-8LP (Optional) | MS120-24P, MS120-48FP | MS120-48FP | MS120-48FP | 
| Interfaces | 8 x GbE | 24/48 x GbE | 48 x GbE | 48 x GbE | 
| Uplinks | 2x 1G SFP | 4 x SFP | 4 x SFP | 4 x SFP | 
| PoE/PoE+ | 67W | 370W/740W | 740W | 740W | 
| **Wireless (MR) - Optional** |  |  |  |  | 
| MR Model | - | MR33 | MR42 | MR42 | 
| Wi-Fi Generation | N/A | Wi-Fi 5 (802.11ac) | Wi-Fi 5 (802.11ac) | Wi-Fi 5 (802.11ac) | 
| Radio | - | 2x2:2 MU-MIMO | 3x3:3 MU-MIMO | 3x3:3 MU-MIMO | 
| Max Throughput | - | 1.3 Gbps | 1.9 Gbps | 1.9 Gbps | 

**Note:** The recommended models for sizing are based on models available at the time of publication for this document. This list will not be updated with new models. For the most up-to-date sizing guide information, please refer to the __Meraki MX Sizing Guide__.
