---
id: collect-261001-meraki/meraki/meraki-and-cisco-collaboration-teleworker-solution-for-businesses-5
title: "meraki-and-cisco-collaboration-teleworker-solution-for-businesses"
domain: meraki
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["agent", "throughput"]
source: docs/RAG/collect-261001-meraki/meraki-and-cisco-collaboration-teleworker-solution-for-businesses.md
source_anchor: ""
source_lines: [610, 711]
sha256: 5a83b2f7f7e4772d7e7cf7531392e51f84acb1d920bfcae579cf1cb4be9d0845
---

# meraki-and-cisco-collaboration-teleworker-solution-for-businesses

**Step 9. (macOS)** Follow the steps to activate APNS through Apple. Download and install the Agent manually or add Systems Manager as an App (in the **Add app** window, select **macOS**, then **SM agent** and click **Next**).


## **Meraki Insight**

The **Meraki Insight** product is designed to give Meraki customers an easy way to monitor the performance of **Web Applications** and **WAN Links** on their network and easily identify if any issues are likely being caused by the network or application. The goal of Meraki Insight is to provide end-to-end visibility to the customers and make sure they have assurance for the mission-critical traffic of the network.


### Enabling and Disabling Meraki Insight

**To Enable** Meraki Insight on a network:

**Step 1.** Ensure the necessary Licensing is available

**Step 2.** Select the checkbox next to the Network where Insight should be enabled

**Step 3. Click '**Add network(s) to Insight' at the top left of the table to Enable Meraki Insight on the selected Network(s)


**To Disable** Meraki Insight on a network:

**Step 1.** Select the checkbox next to the Network with Insight currently enabled

**Step 2.** Click 'Remove network(s) from Insight'

### Configuring Meraki Insight for VoIP

**After** a dashboard account and organization have been created:


**Step 1**. Log in to the Cisco Meraki dashboard


**Step 2**. Select **VoIP Health** from the **Insight** tab


**Step 3**. Select **Configure VoIP servers** on the pup-up


**Step 4**. Input the **Provider name** and the **Hostname or IP**.  Repeat this step to add additional VoIP Servers

### Monitoring Meraki Insight for VoIP

Once you provide the servers you’d like to monitor, you’ll be able to see an org-wide view of performance:


**Step 1**. Select **VoIP Health** from the **Insight** tab.



Clicking on any row will take you to a drill-down view with detailed performance information:



For further details regarding the capabilities of Meraki Insight, see the Meraki Insight Introduction article.

## **Appendix**

### Table 1 – Cisco Meraki devices technical specifications and sizing guide


|  | **Teleworker Site** **(with LTE Backup) Z models** |  | **Teleworker Site** **(with LTE Backup) MX models** |  | 
| Z-Series Model | Z3 | Z3C | MX64, 65, 67, 68 | MX67c / 68c models | 
| Stateful Firewall Throughput | 100 Mbps | 100 Mbps | MX64/65 250 Mbps MX67/68 450 Mbps | MX64/65 250 Mbps MX67/68 450 Mbps | 
| Maximum VPN Throughput | 50 Mbps | 50 Mbps | MX64/65 100 Mbps MX67/68 200 Mbps | MX64/65 100 Mbps MX67/68 200 Mbps | 
| WAN Interfaces (Dedicated) | 1 x GbE RJ45 1 x USB (cellular failover) | 1 x GbE RJ45 1 x Integrated CAT 3 LTE Cellular Modem (cellular failover) 1 x USB (cellular failover) | MX64 1 GbE Dedicated RJ45 MX65 2 GbE Dedicated RJ45 MX67 1 GbE Dedicated RJ45 MX68 2 GbE Dedicated RJ45 1 x USB (cellular failover) | MX67 1 Dedicated RJ45 MX68 2 Dedicated RJ45 1 x Integrated CAT 6 LTE Cellular Modem (cellular failover) 1 x USB (cellular failover) | 
| LAN Interfaces | 4 x GbE RJ45 1 x PoE (802.3af, 15.5W) | 4 x GbE RJ45 1 x PoE (802.3af, 15.5W) | MX64 4 GbE RJ45 (1 Opt WAN) MX65 12 GbE RJ45 (2 x PoE+) MX67 4 GbE RJ45 (1 Opt WAN) MX68 12 GbE RJ45 (2 x PoE+) | MX64 4 GbE RJ45 (1 Opt WAN) MX65 12 GbE RJ45 (2 x PoE+) MX67 4 GbE RJ45 (1 Opt WAN) MX68 12 GbE RJ45 (2 x PoE+) | 
| Integrated Wireless | 4 SSIDs WiFi 5 (802.11a/b/g/n/ac, 2.4/5.0GHz, 2x2 MU-MIMO) | 4 SSIDs WiFi 5 (802.11a/b/g/n/ac, 2.4/5.0GHz, 2x2 MU-MIMO) | 4 SSIDs WiFi 5 (802.11a/b/g/n/ac, 2.4/5.0GHz, 2x2 MU-MIMO) | 4 SSIDs WiFi 5 (802.11a/b/g/n/ac, 2.4/5.0GHz, 2x2 MU-MIMO) | 


**Note:** The recommended models for sizing and are based on models available at the time of publication for this document. This list will not be updated with new models. For the most up-to-date sizing guide information, please refer to the Meraki Z-Series and MX-series Data Sheets.


### Table 2 – Systems Manager Application Configuration


| **Windows Apps** | **Cisco Jabber** | **Webex Meetings** | **Webex Teams** | 
| Name | Cisco Jabber | Cisco Webex Meetings | Webex Teams | 
| Identifier | Random UUID __Here__ | Random UUID __Here__ | Random UUID __Here__ | 
| Icon URL | __Link__ | __Link__ | __Link__ | 
| Vendor | Cisco Systems, Inc. | Cisco Systems, Inc. | Cisco Systems, Inc. | 
| Version | 12.8.0.51973 | 40.2.7.7 | 3.0.15036.0 | 
| Description | Cisco Jabber | Cisco Webex Meetings | Webex Teams | 
| App Download/URL | __Download__ | __File URL__ | __File URL__ | 
| Install Arguments |  | ACCEPT_EULA=TRUE ALLUSERS=1 |  | 


| **macOS Apps** | **Cisco Jabber** | **Webex Meetings** | **Webex Teams** | 
| Name | Cisco Jabber | Cisco Webex Meetings | Webex Teams | 
| Identifier | com.cisco.jabber | com.webex.meetings | com.webex.teams | 
| Icon URL | __Link__ | __Link__ | __Link__ | 
| Vendor | Cisco Systems, Inc. | Cisco Systems, Inc. | Cisco Systems, Inc. | 
| Version | 12.8.0 | 2003.0506.4002.7 | 3.0.15015.0 | 
| Description | Cisco Jabber | Cisco Webex Meetings | Webex Teams | 
| App Download/URL | __Download__ | __File URL__ | __File URL__ | 
| Install Arguments |  |  |  | 

**Note:** The versions and URLs are based on models available at the time of publication for this document. This list will not be updated with new software releases. For the most up-to-date release information, please consult documentation for the specific collaboration product.

- ##### Highest rated (rating)
- ##### Recently updated (date updated)
- ##### Recently added (date created)
