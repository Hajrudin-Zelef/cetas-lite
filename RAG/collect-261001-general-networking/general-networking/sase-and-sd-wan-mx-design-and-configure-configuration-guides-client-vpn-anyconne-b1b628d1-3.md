---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1-3
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1.md
source_anchor: ""
source_lines: [91, 127]
sha256: 24b544db501906f6e5bfaace99b9fd8a8422da4b31a24c9a11ce40ad1b786f7a
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-b1b628d1

After connection, the user should see their local network subnet added as a non secure routes (destinations that should be accessed locally not via the VPN tunnel)
Session Timeout
The client session timeout can be configured using one of the predefined values (8 hours, 1 day, 7 days). Or, you can use the custom option and specify up to a maximum of 256 hours.
Group Policies
The need for access control over remote access connections cannot be over-emphasized. While some administrators use multiple address pools to segment users, others use VLAN tagging to existing subnets. From a Client VPN standpoint, multiple subnets or separate VLANs do not provide access control in itself. What segments users from talking to each other or other network resources is the presence and the enforcement of access rules. For example, if users are in different VLANs and access policies are not enforced somewhere, users could access anything. 
 
AnyConnect on the MX does not support multiple VLANs or address pools for Client VPN users. However, the MX supports the application and enforcement of policies to AnyConnect users on authentication. It is also important to note that, from a Client VPN standpoint on the MX, having users on the same subnet does not mean they are in the same VLAN. Users are assigned a /32 address (one address) from the pool configured on dashboard. Group Policies can then be used to limit users on the same AnyConnect subnet from talking to each other or other resources on the network.
Default Group Policy
Administrators can apply a global group policy to all users connecting through AnyConnect by selecting a configured policy from the default Group Policy drop-down menu. Group policies can be configured via Network-wide > Configure > Group policies. Refer to Creating and Applying Group Policies for more details.
 
Note: If a default group policy set and group policy with Filter-ID is also enabled, the Filter-ID policy passed by the RADIUS server will take precedence over the default group policy.  
Group Policies with RADIUS Filter-ID
AnyConnect supports the application of dashboard-configured group policies to AnyConnect users when authenticating with RADIUS. This is achieved using the RADIUS Filter-ID attribute. To set this up on your MX:
- 
    Create group policies under Network-wide > Configure > Group Policies. Specify rules within the policy. Multiple group policies can be mapped to different user groups on the RADIUS server. In this example, we are matching CONTRACTOR policy to CONTRACTOR user group. 
 
 
- 
    Enable the Filter-ID option on the dashboard. This option is only configurable if you are authenticating with a RADIUS server. 
- 
    Configure the RADIUS server to send an attribute in its accept message containing the name of a group policy configured in dashboard (as a String). Commonly, the Filter-ID attribute will be used for this purpose. The screenshot below shows a network policy in Windows Network Policy Server (NPS), configured to pass the name of a dashboard group policy ("CONTRACTOR") within the Filter-ID attribute:
The RADIUS server is configured with the group policy "CONTRACTOR" defined on dashboard. When a user in the group successfully authenticates, the "CONTRACTOR" group policy name for the authenticated user will be sent in the RADIUS accept message, allowing the MX to apply the requested policy to the user. The group policy name sent by the RADIUS server must match verbatim what is configured on the dashboard for policies to apply correctly. Currently, policies do not show up on Network-wide > Monitor > Clients list page if you have only a security appliance in your dashboard network, however, If you have a combined network, the policy will show under the 802.1X policy column.
Client VPN Connections  
Client view: 
You can see client stats and connection details by clicking on the graph in the bottom-left corner of the client. 
Clients can also see available routes on the Route Details tab. Secure routes are accessible by the client over the VPN while nonsecure routes are not accessible by the client over the VPN. Nonsecure routes are visible when split-tunneling is configured. 
Connection logs can be found under the Message History tab.
Dashboard view:
After configuring client VPN, to see how many users are connected to your network, navigate to Network-wide > Monitor > Clients. All AnyConnect clients will be seen with the AnyConnect icon. You can filter by client VPN using the search menu.
 
Note: The MAC address seen on the client list is is not the actual MAC address of the AnyConnect client. Instead, the displayed address is pseudo-randomly generated, using the provided username as its base. For example, each time someone connects using the name xyz.test@example.com, an entry will show up as active on the clients list with the same given MAC address.
In the event that multiple devices are connected simultaneously with the same set of credentials, the data seen on the list will reflect the most recently connected device.
Depending on your operating system the output of ipconfig/ifconfig may show a default gateway of 0.0.0.0 and subnet mask of 255.255.255.255. This is expected behavior and does not affect the functionality of AnyConnect. 
Event Logging
To see all available events, navigate to Network-wide > Monitor > Event Log and filter the Event type include field by AnyConnect.
To see log-on and log-off events, go to Network-Wide > Monitor > Event log and filter by VPN client connected and VPN client disconnected.
 
