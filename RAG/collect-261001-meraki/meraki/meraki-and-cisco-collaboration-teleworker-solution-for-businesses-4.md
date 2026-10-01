---
id: collect-261001-meraki/meraki/meraki-and-cisco-collaboration-teleworker-solution-for-businesses-4
title: "meraki-and-cisco-collaboration-teleworker-solution-for-businesses"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "distribution", "voice"]
source: docs/RAG/collect-261001-meraki/meraki-and-cisco-collaboration-teleworker-solution-for-businesses.md
source_anchor: ""
source_lines: [450, 609]
sha256: ea8443ab2767c8a342a8febce6115bf0a1205a78da525a48b65a8b0c50688f63
---

# meraki-and-cisco-collaboration-teleworker-solution-for-businesses

- 
    **Name** : Use a name that describes the purpose of the VLAN (e.g. “Voice VLAN”)
- 
    **Subnet** : Defines the subnet to be mapped to the VLAN. Use Classless Inter-Domain Routing (CIDR) notation
- 
    **VLAN Interface IP** : The IP address within the subnet that the Teleworker Gateway will use
- 
    **VLAN ID** : The assignment of the VLAN tag. Use a number between 1 and 4096
- 
    **Group Policy** : Specifies the group policy to be applied to traffic within the subnet


**Step 5**. Click **Preview** then **Update** after entering the VLAN information and confirm VLAN creation. Repeat the process for each of the VLANs needed.



**Step 6**. Navigate to the **Teleworker Gateway** tab and select **DHCP**. Each of the created subnets will have its own DHCP server configuration, ensure the **Run a DHCP server** drop-down option is selected for all VLANs.


**Step 7**. Scroll down to the DHCP server configuration for the Voice VLAN. In order to allow IP phones to retrieve configuration and a hosted phone firmware, add custom DHCP options entries with the following configuration:

- 
    **Option** : Select**Custom** from the drop-down to configure an unlisted DHCP option
- 
    **Code** : Enter the number for the DHCP option (e.g. 150 for TFTP server address)
- 
    **Type** : Select**Text**
- 
    **Value** : Enter the URL or IP address information




 

Note: It is recommended to use at least two of the custom DHCP options to eliminate the need for manual phone provisioning. The sequence Cisco IP phones (MPP) use when booting up is: 66, 160, 159, 150. Enterprise (Non-MPP) phones use DHCP option 150 and 66.


**Step 8**. Scroll down to **Traffic shaping rules** and configure a new rule for **All VoIP & video conferencing** traffic. Refer to the image below for additional details on how to configure the rule.



#### Blueprint Teleworker Wireless Configuration (Z1/Z3/Z4 or MX6X)

This section covers the wireless configuration piece of the reference network. It is highly recommended to perform a pre-install RF survey as outlined in the __Wireless VoIP QoS Best Practices__ document. Follow the next steps to configure the wireless settings of the blueprint network.


**Step 1**. Configure a dedicated SSID for voice by navigating to the **Teleworker gateway** tab and selecting **Wireless settings**


**Step 2**. Change the **Status** of SSID 1 to **Enabled**. Rename the default SSID. Use a name that describes its purpose (e.g. “CiscoVoIP”)


**Step 3**. Change **VLAN assignment** to the Voice VLAN that you created above.


**Step 4**. Change the SSID **Security** to **WPA2 PSK**. Enter a WPA Key for clients to connect to the SSID, and change WPA encryption mode to **WPA2 only**.  If 802.1x is desired for wireless, select **WPA2 Enterprise** then under **Authentication** select **My RADIUS server** and input the organizations RADIUS servers IP addresses 



#### Blueprint New Network Creation (Z1/Z3/Z4 or MX6X)

**After** the blueprint network is created, we can now create new networks by cloning them from the Blueprint network:


**Step 1**. Log in to the Cisco Meraki dashboard as an organization administrator


**Step 2**. Select **Create a new network** from the **Network** dropdown


**Step 3**. Enter a **Network name** and make sure it clearly describes its purpose (e.g. “Teleworker-VPN-Miles”) and click the **Add** button


**Step 4**. Set the **Network type** to **Security appliance**, change **Network configuration** set to **Clone from existing network** and select the **Blueprint Network** from the list.


**Step 5.** Select a “Teleworker Gateway or WAN Appliance” from the bottom to add to this network. Then scroll down and click **Create network**.




#### Blueprint Teleworker Wireless Configuration (Wireless Only)

This section covers the wireless configuration piece of the reference network. Depending on the user equipment being deployed, it may be desired to create two VLAN’s, one dedicated to Data (laptop, mobile devices etc) and a second dedicated to wireless phones. It is highly recommended to perform a pre-install RF survey as outlined in the __Wireless VoIP QoS Best Practices__ document.  Below are the steps to configure your Wireless only configuration for the Blueprint network.


**Step 1**. Configure a dedicated SSID for voice by navigating to the **Wireless** tab and selecting **SSIDs**


**Step 2**. Change the **Status** of SSID 1 to **Enabled**. Rename the SSID. Use a name that describes its purpose (e.g. “CiscoTeleworker” or “CiscoVoIP” etc)


**Step 3**. Change the SSID **Security** to **WPA2 PSK**. Enter a WPA Key for clients to connect to the SSID, and change WPA encryption mode to **WPA2 only**.  If 802.1x is desired for wireless, select **WPA2 Enterprise** then under **Authentication** select **My RADIUS server** and input the organizations RADIUS servers IP addresses.


**Step 4**. Change **Addressing and traffic** to the VPN: tunnel data to a concentrator. Select the **Concentrator** you want to tunnel the data to ( you will need to have a Meraki VPN Concentrator in your organization to select).  Select **Full tunnel: tunnel all traffic.**  Once the template is applied to the network you can “Test connectivity”



**Step 5**. Configure security for your SSID(s) by navigating to the **Wireless** tab and selecting **Firewall & traffic shaping.**  Select the SSID to have changes applied.  You can set firewall rules (L3/4/7 based rules)  By Default voice is prioritized.


**Step 6**. Ensure that the Concentrator has the required DHCP configurations in place for phones to get their needed configurations.



#### Blueprint New Network Creation (Wireless only)

**After** the blueprint network is created, we can now create new networks by cloning them from the Blueprint network:


**Step 1**. Log in to the Cisco Meraki dashboard as an organization administrator


**Step 2**. Select **Create a new network** from the **Network** dropdown


**Step 3**. Enter a **Network name** and make sure it clearly describes its purpose (e.g. “Teleworker-Wireless-VPN-Miles”) and click the **Add** button


**Step 4**. Set the **Network type** to **Wireless**, change **Network configuration** set to **Clone from existing network** and select the **Blueprint Network** from the list.


**Step 5.** Select an "Access Point” from the bottom to add to this network. Then scroll down and click **Create network**.

### **Meraki Systems Manager (Unified Endpoint Management)**

If your deployment includes soft clients (Webex Meetings, Webex Teams, Jabber), you can simplify the deployment and management of these endpoints by deploying Meraki Systems Manager to manage the endpoints, and to deploy the soft clients to them. This section walks through the process of creating the Systems Manager network and basic configuration of these features.

#### Network Creation

**After** a dashboard account and organization have been created:


**Step 1**. Log in to the Cisco Meraki dashboard as an organization administrator


**Step 2**. Select **Create a new network** from the **Network** dropdown


**Step 3**. Enter a **Network name** and make sure it clearly describes its purpose (e.g. “Systems Manager”) and click the **Add** button


**Step 4**. Leave **Network type** set to **EMM (Systems Manager)**, leave **Network configuration** set to **Default Meraki configuration**, then scroll down and click **Create network**.


**Step 5**. Select **Apps** from the **Systems Manager** tab.


**Step 6**. Click **Add app**, then select the App platform. See **Table 2 of the Appendix** for application details.


**Step 7**. Select **Devices** from the **Systems Manager** tab. Click the type of device corresponding with the device that you want to manage.



**Step 8. (Windows)** Download the Agent (required for software distribution via Systems Manager) and Install it.


