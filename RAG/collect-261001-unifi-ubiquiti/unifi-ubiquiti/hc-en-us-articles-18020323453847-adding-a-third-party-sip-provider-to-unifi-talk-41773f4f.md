---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-18020323453847-adding-a-third-party-sip-provider-to-unifi-talk-41773f4f
title: "hc-en-us-articles-18020323453847-adding-a-third-party-sip-provider-to-unifi-talk-41773f4f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-18020323453847-adding-a-third-party-sip-provider-to-unifi-talk-41773f4f.md
source_anchor: ""
source_lines: [1, 112]
sha256: a7011e57387bb152343321d517a5c641bb0ea07352d22116342f835d4e0a07be
---

# hc-en-us-articles-18020323453847-adding-a-third-party-sip-provider-to-unifi-talk-41773f4f

Adding a Third-Party SIP Provider to UniFi Talk
If you currently subscribe to a third-party SIP provider, you can use your existing service with your UniFi Talk deployment without needing to purchase an official UniFi Talk subscription. Third-party SIP configuration is possible only with unlocked UniFi Talk phones. This guide will walk you through configuring SIP Trunk and connecting your provider to UniFi Talk.
To simplify the setup process, we have prepared templates for the following SIP providers: Breeze Connect, Deutsche Telekom, easybell, DIDWW, Telnyx, and VoIP.ms. While you can configure UniFi Talk with any SIP provider for maximum flexibility, please note that our support team will not be able to assist with SIP providers that are not listed.
Create & Configure SIP Trunking
Set Up a New SIP Trunk
Before configuring UniFi Talk, set up a new SIP trunk through your SIP provider's dashboard. Note that these instructions are general; consult your provider's documentation for specific details.
- 
Create a new SIP trunk:
  - Set the authentication credentials (username and password).
- 
Add your router's public IP address to the IP Access Control List (ACL):
  - If required, add a new entry for your router's public IP address. Depending on your provider’s requirements, you may need to enter this IP address, use CIDR notation, or include a port number (e.g., 1.2.3.4:6767).
  - We recommend specifying port 6767, as it is used by UniFi Talk when the Static Signaling Port setting is enabled.
- 
Assign Direct Inward Dialing (DID) numbers:
  - Ensure that the DID numbers you want to use with UniFi Talk are assigned to the newly created SIP trunk.
Configure UniFi Talk
- 
Enable Static Signaling Port:
  - Navigate to System Settings in UniFi Talk and enable the Static Signaling Port to use port 6767 for SIP signaling.
- Select Add Third-Party SIP Provider and fill out the necessary fields:
  - Custom Fields: Reference your provider’s ITSPs directory or user documentation.
  - Destination Countries: Select specific countries for routing through this SIP provider.
  - Handle All Outgoing Calls By Default: Enable if you want all outgoing calls routed through this SIP provider.
  - DID Numbers & IP Address Range: Enter the IP addresses of the servers used by your SIP provider.
Assign DID Numbers
Once your SIP provider is connected and your DID numbers are imported, assign them to your users. For more information, refer to our Managing Users guide.
Note: If you're using a third-party SIP provider, they will be responsible for maintaining E911 compliance. Contact your provider for guidance on meeting these requirements.
Adjust Port Forwarding Rules
Some third-party SIP configurations require the Static Signaling Port setting to be enabled. Follow these steps to adjust port forwarding rules:
- 
Enable Static Signaling Port:
  - In the Talk application, go to Settings > System Settings > Static Signaling Port and check the box.
- 
Create Port Forwarding Rule:
  - If your routing tasks are handled by UniFi, go to Network application > Settings > Security > Port Forwarding and create a rule:
    - Interface: Both
    - From: Any
    - Port: 6767
    - Forward IP: Your UniFi Console’s IP address (e.g., 192.168.1.1)
    - Forward Port: 6767
    - Protocol: Both
- If your routing tasks are handled by UniFi, go to Network application > Settings > Security > Port Forwarding and create a rule:
- 
Forward Port 6767 from Upstream Routers:
  - If you have another router upstream, forward incoming traffic on port 6767 from that router to your UniFi Console.
- 
Create Firewall Rule:
  - If your firewall rules are managed by UniFi, go to Network application > Settings > Security > Traffic & Firewall Rules and create a rule:
    - Rule Type: Advanced
    - Type: Internet Local
    - Action: Accept
    - Protocol: All
    - Source Type: Port/IP Group
    - Address Group: Any
    - Port Group: Any
    - Destination Type: Port/IP Group
    - Address Group: Any
    - Port Group: Select New
    - Profile Name: Enter a name
    - Port: 6767, select Add
    - Create
- If your firewall rules are managed by UniFi, go to Network application > Settings > Security > Traffic & Firewall Rules and create a rule:
Configuring Specific SIP Providers
Telnyx SIP
- 
Create a new SIP Connection:
  - In your Telnyx portal, go to Voice > SIP Trunking > Create SIP Connection. Choose Credentials as your SIP Type.
- 
Save Credentials:
  - Write down the Username and Password generated by Telnyx.
- 
Buy a DID Number:
  - In the Telnyx portal, go to Numbers > Search & Buy Numbers.
- 
Assign DID Number to SIP Profile:
  - Under My Numbers, select your DID number, go to the Voice tab, and assign a SIP Connection.
- 
Configure Outbound Voice Profile (OVP):
  - In Voice > Outbound Voice Profiles, configure the profile and assign it to your SIP trunk.
- 
Account Verification:
  - If using Telnyx in Europe, pass the account verification steps in your Telnyx portal.
Troubleshooting:
- If you cannot make or receive calls, check the SIP Trunk configuration on the Telnyx site and ensure details match those in UniFi Talk.
Useful Telnyx articles:
DIDWW SIP
- 
Create a new SIP Connection:
  - In the DIDWW portal, navigate to Voice > Inbound Trunks. Use the default settings, noting the Preferred server.
- 
Find Public WAN IP:
  - Locate your UniFi OS console’s public WAN IP address in the UniFi OS settings.
- 
Enter WAN IP in DIDWW Portal:
  - Enter your console’s WAN IP address into the DIDWW portal’s Host field.
- 
Buy and Assign a DID Number:
  - In the DIDWW portal, go to Phone Numbers > My Numbers and assign a DID number to a SIP Trunk.
- 
Configure Outbound SIP Trunk:
  - In Voice > Outbound Trunks, create a new SIP trunk, choose Credentials and IP-based authentication, and enter your console’s WAN IP address.
- 
Save Credentials:
  - In the DIDWW portal, find the credentials by clicking the keychain icon in the Outbound Trunk tab.
- 
Enter Credentials in UniFi Talk:
  - Enter the Outbound Username and Password, and the Outbound Host information into the UniFi Talk application.
Account Verification:
- If using DIDWW in Europe, pass the account verification steps in your DIDWW portal.
Troubleshooting:
- If you cannot make or receive calls, check the SIP Trunk configuration on the DIDWW site and ensure details match those in UniFi Talk.
