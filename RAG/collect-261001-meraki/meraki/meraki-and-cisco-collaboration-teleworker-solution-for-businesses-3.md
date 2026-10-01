---
id: collect-261001-meraki/meraki/meraki-and-cisco-collaboration-teleworker-solution-for-businesses-3
title: "meraki-and-cisco-collaboration-teleworker-solution-for-businesses"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-meraki/meraki-and-cisco-collaboration-teleworker-solution-for-businesses.md
source_anchor: ""
source_lines: [288, 449]
sha256: 42d9269332d6cbd479b53ac42c50a0c410a25c784c6cfbf30a319ccd79a8e323
---

# meraki-and-cisco-collaboration-teleworker-solution-for-businesses

**Step 4**. Set the **Network type** to **Security appliance**, change **Network configuration** set to **Bind to template** and select the Template-Teleworker-VPN from the list. **Meraki configuration**, then scroll down and click **Create network**.


**Step 5.** Select a “Teleworker Gateway or WAN Appliance” from the bottom to add to this network. Then scroll down and click **Create network**.



#### Template Wireless Configuration (Wireless Only)

This section covers the wireless configuration piece of the reference network. Depending on the user equipment being deployed, it may be desired to create two VLAN’s, one dedicated to Data (laptop, mobile devices etc) and a second dedicated to wireless phonesIt is highly recommended to perform a pre-install RF survey as outlined in the __Wireless VoIP QoS Best Practices__ document.


**Step 1**. Configure a dedicated SSID for voice by navigating to the **Wireless** tab and selecting **SSIDs**


**Step 2**. Change the **Status** of SSID 1 to **Enabled**. Rename the SSID. Use a name that describes its purpose (e.g. “CiscoTeleworker” or “CiscoVoIP” etc)


**Step 3**. Change the SSID **Security** to **WPA2 PSK**. Enter a WPA Key for clients to connect to the SSID, and change WPA encryption mode to **WPA2 only**.  If 802.1x is desired for wireless, select **WPA2 Enterprise** then under **Authentication** select **My RADIUS server** and input the organizations RADIUS servers IP addresses.


**Step 4**. Change **Addressing and traffic** to the VPN: tunnel data to a concentrator. Select the **Concentrator** you want to tunnel the data to ( you will need to have a Meraki VPN Concentrator in your organization to select).  Select **Full tunnel: tunnel all traffic.**  Once the template is applied to the network you can “Test connectivity”

 


**Step 5**. Configure security for your SSID(s) by navigating to the **Wireless** tab and selecting **Firewall & traffic shaping.**  Select the SSID to have changes applied.  You can set firewall rules (L3/4/7 based rules)  By Default voice is prioritized.


**Step 6**. Ensure that the Concentrator has the required DHCP configurations in place for phones to get their needed configurations.



#### New Network Creation (Wireless only)

**After** the Template-Teleworker-VPN template is created we can now create new networks and apply the template we just created:


**Step 1**. Log in to the Cisco Meraki dashboard as an organization administrator


**Step 2**. Select **Create a new network** from the **Network** dropdown tab


**Step 3**. Enter a **Network name** and make sure it clearly describes its purpose (e.g. “Teleworker-VPN-Wireless-Miles”) and click the **Add** button


**Step 4**. Set the **Network type** to **Wireless**, change **Network configuration** set to **Bind to template** and select the Template-Teleworker-VPN from the list. **Meraki configuration**, then scroll down and click **Create network**.


**Step 5.** Select an Access Point from the bottom to add to this network. Then scroll down and click **Create network**.


This concludes the steps for the WAN Appliance, Teleworker Gateway, and Wireless portion of the configuration for the network template.

## **Setup and General Provisioning - Blueprint Network Option**

In order to deploy this solution using the best practices provided here for Blueprint Network Method, the following workflows must be performed in order by the administrator. The initial setup consists of...

1. 
    Configuring the general settings for the organization
2. 
    Creating the blueprint network



**After initial setup**, when the base organization containing the blueprint network is available, administrators can follow the below workflow for deploying new Teleworker sites.

1. 
    Create a new network by cloning the blueprint network
2. 
    Add devices to the newly created network



The next section will walk through the process of configuring the base organization as well as the blueprint network.

### Creating the Blueprint Network

This section walks through the process of creating the blueprint network that will serve as the base configuration for the network infrastructure and will be used as a starting point for new customer deployments. The blueprint network represents the Cisco recommended settings for successful deployments.

#### Network Creation

**After** a dashboard account and organization have been created:


**Step 1**. Log in to the Cisco Meraki dashboard as an organization administrator


**Step 2**. Select **Create a new network** from the **Network** dropdown


**Step 3**. Enter a **Network name** and make sure it clearly describes its purpose (e.g. “Blueprint Teleworker Network”) and click the **Add** button


**Step 4**. Leave **Network type** set to **Security appliance**, leave **Network configuration** set to **Default Meraki configuration**, then scroll down and click **Create network**.


Note: There is currently a limitation in the Meraki dashboard where you are not able to change the Regulatory Domain for the Wireless settings on the Teleworker gateway if you create the network as a Combined network – ensure that you select the Security appliance setting to avoid this limitation.

#### Network-wide (General) Configuration

For the purposes of consistency, it is recommended to configure the general settings that are common across Teleworker locations in the blueprint network. Below is a list of the most relevant settings available when selecting **General** under the **Network-wide** tab.

- 
    **Country/Region** sets the country of network. Regulatory domain is set based on this
- 
    **Local time zone** sets the time zone of the network
- 
    **Traffic analysis** specifies the level of detail desired for traffic analysis


For added visibility, use the **Detailed** option for **Traffic analysis**. This enables collection of detailed information about the destinations which can be useful for troubleshooting.



 

Additionally, the **Traffic analysis** section allows for the definition of specific destinations to track and to build a custom pie chart. The destinations can be defined based on a **HTTP hostname**, **Port**, **IP range** or **IP range & port**.



 

When a custom pie is configured and there is matching traffic to the destinations, traffic usage information will be available when selecting **Clients** from the **Network-wide** tab.



 

The following section goes through the Teleworker gateway specific configuration.


Blueprint Teleworker appliance Configuration (Z1/Z3/Z4 or MX6X)


The configuration steps described in this section are specific to the Meraki Z-Series Teleworker gateway. The Z-Series line has built-in wireless and VPN capabilities and is used in all the Teleworker deployments as outlined in **Table 1 of the Appendix**. Use the following steps to configure this portion of the reference network:


**Step 1**. Navigate to the **Teleworker gateway** tab and select **Addressing & VLANs**


**Step 2**. From the **Deployment Settings** section of the page ensure that the **Routed** and **MAC address** options are selected for the **Mode** and **Client tracking** settings respectively



 

Note: The Routed mode allows the Teleworker gateway to act as a layer 3 gateway for the different subnets created and all Internet bound client traffic will be translated (NAT). More information on this is available in the MX Addressing and VLANs article.


**Step 3**. Scroll down to the **Routing** section and enable the use of VLANs


Note: Teleworker gateways define the VLANs and subnets that exist in the remote network. It is best practice to ensure traffic is segregated by creating dedicated VLANs for voice and data traffic.


**Step 4**. Click on the **Add VLAN** button to create each of the VLANs needed for remote networks.




