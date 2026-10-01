---
id: collect-261001-meraki/meraki/meraki-and-cisco-collaboration-teleworker-solution-for-businesses-2
title: "meraki-and-cisco-collaboration-teleworker-solution-for-businesses"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-meraki/meraki-and-cisco-collaboration-teleworker-solution-for-businesses.md
source_anchor: ""
source_lines: [133, 287]
sha256: 785bebc3f59285fefa07da79a9f03e034b3fff3c71c5382a7653e1418368ff76
---

# meraki-and-cisco-collaboration-teleworker-solution-for-businesses

**Step 3**. Select a **Create a new template** and select **Create new** then make you provide a name that clearly describes its purpose (e.g. “Template-Teleworker-VPN”) and click the **Add** button. Click **Close** and then **Save changes**


Note: There is currently a limitation in the Meraki dashboard where you are not able to change the Regulatory Domain for the Wireless settings on the Teleworker gateway if you create the network as a Combined network – ensure that you select the Security appliance setting to avoid this limitation.


**Step 4**. Select the newly created template (e.g. “Template-Teleworker-VPN”) from the **Network** dropdown tab. We will make all our configuration changes in this template and apply the template to all future created networks.

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


Template Teleworker Appliance Configuration (Z1/Z3/Z4 or MX 6x)


The configuration steps described in this section are specific to the Meraki Z-Series Teleworker gateway. The Z-Series line has built-in wireless and VPN capabilities and is used in all the Teleworker deployments as outlined in **Table 1 of the Appendix**. Use the following steps to configure this portion of the reference network:


**Step 1**. Navigate to the **Security & SD-WAN** tab and select **Addressing & VLANs**


**Step 2**. From the **Deployment Settings** section of the page ensure that the **Routed** and **MAC address** options are selected for the **Mode** and **Client tracking** settings respectively



Note: The Routed mode allows the Teleworker gateway to act as a layer 3 gateway for the different subnets created and all Internet bound client traffic will be translated (NAT). More information on this is available in the MX Addressing and VLANs article.


**Step 3**. Scroll down to the **Routing** section and enable the use of VLANs


Note: Teleworker gateways define the VLANs and subnets that exist in the remote network. It is best practice to ensure traffic is segregated by creating dedicated VLANs for voice and data traffic.


**Step 4**. Click on the **Add VLAN** button to create each of the VLANs needed for remote networks.


For the Subnet here, we are able have dashboard automatically create non-overlapping subnets out of a larger supernet. This avoids the need to have to configure each networks IP addressing as the networks are created.

 



- 
    **Name** : Use a name that describes the purpose of the VLAN (e.g. “Voice VLAN”)
- 
    **Subnet** : Defines the subnet to be mapped to the VLAN. Use Classless Inter-Domain Routing (CIDR) notation.  We can configure Unique here as in the example above. Each new network that the template is added to will provision a unique /29 subnet out of the larger 192.168.100.0/22 network
- 
    **VLAN ID** : The assignment of the VLAN tag. Use a number between 1 and 4096
- 
    **Group Policy** : Specifies the group policy to be applied to traffic within the subnet


**Step 5**. Click **Preview** then **Update** after entering the VLAN information and confirm VLAN creation. Repeat the process for each of the VLANs needed.  Note that the VLAN Interface IP column will be “Auto-generated” and the Appliance will take the first IP each time a subnet is provisioned.



**Step 6**. Navigate to the **Security & SD-WAN** tab and select **DHCP**. Each of the created subnets will have its own DHCP server configuration, ensure the **Run a DHCP server** drop-down option is selected for all VLANs.


**Step 7**. Scroll down to the DHCP server configuration for the Voice VLAN. In order to allow IP phones to retrieve configuration and a hosted phone firmware, add custom DHCP options entries with the following configuration:

- 
    **DNS namesservers:** Enter the DNS servers for your organization for internal name resolution
- 
    **Option** : Select**Custom** from the drop-down to configure an unlisted DHCP option
- 
    **Code** : Enter the number for the DHCP option (e.g. 150 for TFTP server address)
- 
    **Type** : Select**Text**
- 
    **Value** : Enter the URL or IP address information



Note: It is recommended to use at least two of the custom DHCP options to eliminate the need for manual phone provisioning. The sequence Cisco IP phones (MPP) use when booting up is: 66, 160, 159, 150. Enterprise (Non-MPP) phones use DHCP option 150 and 66.


**Step 8**. Provide the appropriate configuration settings for the Data VLAN, and any other VLAN that DHCP will be configured for.:

- 
    **DNS namesservers:** Enter the DNS servers for your organization for internal name resolution
- 
    **Option** : Select**Custom** from the drop-down to configure an unlisted DHCP option


**Step 9**. Navigate to the **Security & SD-WAN** Scroll down to **SD-WAN & traffic shaping** and configure a new rule for **All VoIP & video conferencing** traffic. Refer to the image below for additional details on how to configure the rule.



 

#### Template Wireless Configuration (Z1/Z3/Z4 or MX6X)

This section covers the wireless configuration piece of the reference network. It is highly recommended to perform a pre-install RF survey as outlined in the __Wireless VoIP QoS Best Practices__ document. Follow the next steps to configure the wireless settings of the blueprint network.


**Step 1**. Configure a dedicated SSID for voice by navigating to the **Security & SD-WAN** tab and selecting **Wireless settings**


**Step 2**. Change the **Status** of SSID 1 to **Enabled**. Rename the default SSID. Use a name that describes its purpose (e.g. “CiscoVoIP”)


**Step 3**. Change **VLAN assignment** to the Voice VLAN that you created above.


**Step 4**. Change the SSID **Security** to **WPA2 PSK**. Enter a WPA Key for clients to connect to the SSID, and change WPA encryption mode to **WPA2 only**.  If 802.1x is desired for wireless, select **WPA2 Enterprise** then under **Authentication** select **My RADIUS server** and input the organizations RADIUS servers IP addresses.



This concludes the Meraki Teleworker gateway wireless portion of the configuration for the network template.

#### New Network Creation (Z1/Z3/Z4 or MX6X)

**After** the Template-Teleworker-VPN template is created we can now create new networks and apply the template we just created:


**Step 1**. Log in to the Cisco Meraki dashboard as an organization administrator


**Step 2**. Select **Create a new network** from the **Network** dropdown tab


**Step 3**. Enter a **Network name** and make sure it clearly describes its purpose (e.g. “Teleworker-VPN-Miles”) and click the **Add** button


