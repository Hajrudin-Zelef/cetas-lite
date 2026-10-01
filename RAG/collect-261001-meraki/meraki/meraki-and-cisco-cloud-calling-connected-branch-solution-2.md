---
id: collect-261001-meraki/meraki/meraki-and-cisco-cloud-calling-connected-branch-solution-2
title: "Appendix"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-meraki/meraki-and-cisco-cloud-calling-connected-branch-solution.md
source_anchor: ""
source_lines: [112, 256]
sha256: 605c862b072d0f3ef273ee79fe83d4ae51c53ca997963131c703f232e9f97838
---

# Appendix

This section walks through the process of creating the blueprint network that will serve as the base configuration for the network infrastructure and will be used as a starting point for new customer deployments. The blueprint network represents the Cisco recommended settings for successful deployments.


**Note:** It is required for the configuration to be applied to a configuration template in order to allow for it to be cloned from the base org onto new customer organizations.

### Network Creation

**After** a dashboard account and organization have been created:

**Step 1**. Log in to the Cisco Meraki dashboard as an organization administrator.

**Step 2**. Select **Configuration templates** from the **Organization** tab.

**Step 3**. Click the **Create a new template** button.

**Step 4.** Enter a **Template** **name** and make sure it clearly describes its purpose (e.g.“Blueprint Network”) and click the **Add** button.

**Step 5.** Leave **Target networks** blank, click on the **Close** button, and **Save changes**.

**Step 6**. Navigate to the **Organization** tab and select **Configuration templates.** The newly configured template has been created and it is available for configuration.

**Step 7**. Click on the name for the template (e.g. Blueprint Network) to start configuring it.


**Note:** By default, the ability to create new networks by cloning from a configuration template (blueprint network) is not enabled. Please reach out to Meraki Support to have this functionality enabled on the base org. This process only needs to be performed once for the base org. Newly created organizations will inherit this functionality from the base org.

### Network-Wide (General) Configuration

For the purposes of consistency, it is recommended to configure the general settings that are common across customer sites in the blueprint network. Below is a list of the most relevant settings available when selecting **General** under the **Network-wide** tab.

- 
    **Country/Region** sets the country of network; regulatory domain is set based on this
- 
    **Local time zone** sets the time zone of the network
- 
    **Traffic analysis** specifies the level of detail desired for traffic analysis

For added visibility, use the **Detailed** option for **Traffic analysis**. This enables collection of detailed information about the destinations which can be useful for troubleshooting.



Additionally, the **Traffic analysis** section allows for the definition of specific destinations to track and build a custom pie chart. The destinations can be defined based on **HTTP hostname**, **Port**, **IP range,** or **IP range & port**.



When a custom pie is configured and there is matching traffic to the destinations, traffic usage information will be available when selecting **Clients** from the **Network-wide** tab.



### Security & SD-WAN (MX) Configuration

The configuration steps described in this section are specific to the Meraki MX Security appliances. The MX line has built-in security, SD-WAN capabilities, and is used in all the branch deployments as outlined in **table 1** **of the Appendix**. Use the following steps to configure this portion of the reference network:

**Step 1.** Navigate to the **Security & SD-WAN** tab and select **Addressing & VLANs**

**Step 2.** From the **Deployment Settings** section of the page, ensure that the **Routed** and **MAC address** options are selected for the **Mode** and **Client tracking** settings respectively



**Note:** The **Routed** mode allows the MX to act as a Layer 3 gateway for the different subnets created, and all internet-bound client traffic will be translated (NAT). More information on this is available in the __MX Addressing and VLANs article__.


**Step 3.** Scroll down to the **Routing** section and enable the use of VLANs.


**Note:** Security appliances (MX) define the VLANs and subnets that exist in the branch network. It is best practice to ensure traffic is segregated by creating dedicated VLANs for voice, data, and guest traffic.


**Step 4.** Click on the **Add VLAN** button to create each of the VLANs needed for branch deployments.



- 
    **Name:** Use a name that describes the purpose of the VLAN (e.g. “Voice VLAN”).
- 
    **Subnet:** Defines the subnet to be mapped to the VLAN. Use Classless Inter-Domain Routing (CIDR) notation.
- 
    **MX IP:** The IP address within the subnet that the Security Appliance (MX) will use.
- 
    **VLAN ID:** The assignment of the VLAN tag. Use a number between 1 and 4,096.
- 
    **Group Policy:** Specifies the group policy to be applied to traffic within the subnet.


**Step 5.** Click **Update** after entering the VLAN information and confirm VLAN creation. Repeat the process for each of the VLANs needed.



**Step 6.** Navigate to the **Security & SD-WAN** tab and select **DHCP**. Each of the created subnets will have its own DHCP server configuration, ensure the **Run a DHCP serve**r drop-down option is selected for all VLANs.

**Step 7.** Scroll down to the DHCP server configuration for the Voice VLAN. In order to allow IP phones to retrieve configuration and a hosted phone firmware, add custom DHCP options entries with the following configuration:

- 
    **Option:** Select**Custom** from the drop-down to configure an unlisted DHCP option.
- 
    **Code:** Enter the number for the DHCP option (e.g. 150 for TFTP server address).
- 
    **Type:** Select**Text.**
- 
    **Value:** Enter the URL or IP address information.



**Note:** It is recommended to use at least two of the custom DHCP options to eliminate the need for manual phone provisioning. The sequence Cisco IP phones (MPP) use when booting up is: 66, 160, 159, 150


**Step 8.** Navigate to the **Security & SD-WAN** tab and select **SD-WAN & Traffic shaping**. Scroll down to the **SD-WAN policies** section and add a preference by using **All VoIP & video conferencing** under traffic filters and selecting **Best for VoIP** as the preferred uplink. 

 


**Step 9.** Scroll down to **Traffic shaping rules** and configure a new rule for **All VoIP & video conferencing** traffic. Refer to the image below for additional details on how to configure the rule.



### Switching (MS) Configuration

This section walks through the Meraki MS switching configuration portion of the blueprint network.

**Step 1.** Navigate to the **Switch** tab and select **Switch settings**. 

**Step 2.** Scroll down to the **Quality of service** section of the page and create a new rule. Enter the voice VLAN ID, select **Any** protocol from the drop-down menu, and set the **DSCP** setting to **Trust incoming DSCP**.



**Note:** Switch port configuration is applied once devices are added to customer networks. This is covered in the **Customer Provisioning** section of this document.

### Wireless (MR) Configuration

This section covers the Meraki MR wireless configuration piece of the reference network and it is only applicable to site deployments that leverage Cisco wireless phones. It is highly recommended to perform a pre-install RF survey as outlined in the __Wireless VoIP QoS Best Practices document__. Follow the next steps to configure the wireless settings of the blueprint network.

**Step 1**. Configure a dedicated SSID for voice by navigating to the **Wireless** tab and selecting **SSIDs**. 

**Step 2**. Rename the default SSID. Use a name that describes its purpose (e.g. “CiscoVoIP”). 

**Step 3**. Navigate to the **Wireless** tab and select **Access control** to configure the SSID. 

**Step 4**. Configure the **Network access** section using a pre-shared key with WPA2 as the authentication method and WPA2 only as the **WPA encryption mode**. 



**Step 5**. Scroll down to **Addressing and traffic** and select **Bridge mode**. In this mode, the Meraki MR access points will bridge the client directly to the LAN, and it is recommended for Layer 2 seamless roaming.



