---
id: collect-261001-general-networking/general-networking/wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e-3
title: "wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e.md
source_anchor: ""
source_lines: [105, 153]
sha256: 5395ccdb64d26bb31e44cf8612a037fec5609409e3e022023ddbd2de2937f98f
---

# wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e

  - Authentication: WPA2-Enterprise or WPA-Enterprise
  - Encryption: AES or TKIP
  - Network Authentication Method: Microsoft: Protected EAP (PEAP)
  - Authentication mode: Computer Authentication (for machine auth)
- 
    Click Properties.
- 
    For Trusted Root Certification Authorities select the check box next to the appropriate Certificate Authorities and click OK.
- 
    Click OK to close out and click Apply on wireless policy page to save the settings.
- 
    Apply the GPO to the domain or OU containing the domain member computers (refer to Microsoft's Deploying Group Policy documentation for details).
Dashboard Configuration
When a RADIUS Server is added to the configuration, it will disconnect current clients connected to the SSID.
Once a RADIUS server has been set up with the appropriate requirements to support authentication, the following instructions explain how to configure an SSID to support WPA2-Enterprise, and authenticate against the RADIUS server:
- In dashboard, navigate to Wireless > Configure > Access control.
- Select your desired SSID from the SSID drop down (or navigate to Wireless > Configure > SSIDs to create a new SSID first).
- For Security choose Enterprise with my RADIUS server.
- Under RADIUS click Add server
- Enter the following information in the table:
- Host IP or FQDN*: IP address or FQDN of your RADIUS server, reachable from the access points.
- Auth port: UDP port of the RADIUS server listens on for Access-Requests; 1812 by default.
- Secret: RADIUS client shared secret
6. Click the Save button.
*The network and all the access points must be running MR28.0+ to support FQDN.
Aside from the RADIUS server requirements outlined above, all authenticating access points will need to be able to contact the IP address and port specified in dashboard. Make sure that your access points all have network connectivity to the RADIUS server and no firewalls are preventing access.
VLAN Tagging Options
Dashboard offers a number of options to tag client traffic from a particular SSID with a specific VLAN tag. Most commonly, the SSID will be associated with a VLAN ID (see VLAN Tagging on MR Access Points), so all client traffic from that SSID will be sent on that VLAN.
With RADIUS integration, a VLAN ID can be embedded within the RADIUS server's response. This allows for dynamic VLAN assignment based on the RADIUS server's configuration. Please refer to our documentation regarding Per-User VLAN tagging for configuration specifics.
Testing RADIUS from Dashboard
Dashboard has a built-in RADIUS test utility, to ensure that all access points (at least those broadcasting the SSID using RADIUS) can contact the RADIUS server:
- Navigate to Wireless > Configure > Access control.
- Ensure that WPA2-Enterprise was already configured based on the Dashboard Configuration section of this article.
- Under RADIUS servers, click the Test button for the desired server.
- Enter the credentials of a user account in the Username and Password fields.
- Click Begin test.
- The window will show progress of testing from each access point in the network, and then present a summary of the results at the end.
- APs passed: Access points that were online and able to successfully authenticate using the credentials provided.
- APs failed: Access points that were online but unable to authenticate using the credentials provided. Ensure the server is reachable from the access points and the access points are added as clients on the RADIUS server.
- APs unreachable: Access points that were not online and thus could not be tested with.
When using the RADIUS Test tool on the Dashboard and an AP fails the test, the pop-up displays a link to a list of all the alerting APs, including those alerting due to 802.1X failures and other issues. A delay of up to 5 minutes between an AP failure due to an 802.1X failure and the Dashboard reporting an alert is expected.
RADIUS Accounting
Optionally, RADIUS accounting can be configured on an SSID that's using WPA2-Enterprise with RADIUS authentication. When RADIUS accounting is configured, "start" and "stop" accounting messages are sent from the access point to the specified RADIUS accounting server.
The following instructions explain how to configure RADIUS accounting on an SSID:
- Navigate to Wireless > Configure > Access control and select the desired SSID from the dropdown menu.
- Under RADIUS accounting servers, click Add a server.
Note: Multiple servers can be added for failover. RADIUS messages will be sent to these servers in a top-down order.
- Enter the details for: 
    
