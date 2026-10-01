---
id: collect-261001-general-networking/general-networking/switching-ms-switches-design-and-configure-configuration-guides-access-control-c-12e04954
title: "switching-ms-switches-design-and-configure-configuration-guides-access-control-c-12e04954"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/switching-ms-switches-design-and-configure-configuration-guides-access-control-c-12e04954.md
source_anchor: ""
source_lines: [1, 38]
sha256: 853450b2c882a9a06df9e6180f2b14dd3d634b3ea2b7dbd46969285bfe18fdc8
---

# switching-ms-switches-design-and-configure-configuration-guides-access-control-c-12e04954

Configuring 802.1X Access Policies on MS Switches using Windows 2008 NPS
Preventing unauthorized access to data network is a critical job for a Network Administrator. This article discusses the benefits of using 802.1X access policies to secure LAN access on your Cisco Meraki MS Switches, and walks through the steps to configure your Windows 2008 NPS server, MS Switch, and your Windows and Mac clients.
Note: A trusted certificate must be installed on your NPS server in order for the switches to securely communicate with the server.
Benefits and Use Case
Often administrators overlook the importance of securing access to their switches.   Implementing 802.1X in an Ethernet environment secures your LAN by providing authentication and access control.  Authentication requires clients to provide credentials to gain access to your LAN.  Access control gives you the ability to determine the level of access a user is given on the network.  Additionally, 802.1X also provides administrators with accounting information.  Accounting gives administrators insight as to who is connected to the LAN and when  they were connected, which can be useful for network monitoring.
 
Ports in common areas make your corporate network vulnerable to access by guests and other unauthorized users. It's also possible for unauthorized users to gain physical access to switches, or gain access to a port in a non-public area. Utilizing 802.1X on your access layer switches gives you the ability to allow users who successfully authenticate access to the corporate LAN. Unauthorized users can be blocked entirely or placed on a guest VLAN that provides Internet access only.
Adding MS Switches as RADIUS clients on the NPS Server
All switches that will use 802.1X Access Policies must be added as clients on the NPS server. Below are the steps to add the switches as RADIUS clients.
- Open the NPS Server Console by going to Start > Programs > Administrative Tools >Network Policy Server.
- In the Left pane, expand the RADIUS Clients and Servers option.
- Right click the RADIUS Clients option and select New.
- Enter a Friendly Name for the MS Switch.
- Enter the the IP Address of your MS Switch.
- Create and enter a RADIUS Shared Secret (note this secret - we will need to add this to the Dashboard).
- Press OK when finished.
- Repeat these steps b - g for all switches. See the following image for a sample RADIUS client configuration.
Configuring a RADIUS NPS Policy
- In the NPS Server Console, navigate to NPS (Local).
- In the main window select RADIUS server for 802.1X Wireless or Wired Connections from the drop down box.
- Press the green button labeled Configure 802.1X.
- In the wizard dialog select the option for Secure Wired (Ethernet) Connections and enter a descriptive name (e.g. MS 802.1X Access Policies)
- On the next screen you should see your list of MS Switches that were added as RADIUS clients. Press Next to proceed.
- Select the Authentication Method Microsoft: Protected EAP (PEAP) and press Next.
- Click Add and enter the name of Windows Group you would like to give LAN access to and press Next. (e.g Domain Users)
- On the Configure Traffic Controls page leave the default settings and press Next to continue.
- Review the settings then press Finish.
Creating 802.1X Access Policy on Dashboard
- On the Dashboard navigate to Configure > Access Policies.
- Click on the link Add Access Policy in the main window then click the link to Add a server.
- Enter the IP address of the RADIUS server, the port (default is 1812), and the secret created earlier.
- Specify a Guest VLAN. This will be the VLAN unauthenticated users will be placed on.
- Click Save changes
Apply Access Policy to Switch Ports
- Navigate to Monitor > Switch Ports.
- Select the port(s) you would like to apply the access policy to and press the Edit button.
- Convert the port type from trunk to access. Note: you can only apply an Access Policy to an access port.
- From the Access Policy drop down box, select the Access Policy you created and press the Update ports button.
