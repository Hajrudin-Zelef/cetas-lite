---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-setup-remote-forticlient-using-dial-up-ipsec-with-ta-ff051ae2
title: "t5-fortigate-technical-tip-setup-remote-forticlient-using-dial-up-ipsec-with-ta--ff051ae2"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-setup-remote-forticlient-using-dial-up-ipsec-with-ta--ff051ae2.md
source_anchor: ""
source_lines: [1, 5]
sha256: ece865c016c1c01282385740e678ffddccf16aa21c72f79760591b09205ce33f
---

# t5-fortigate-technical-tip-setup-remote-forticlient-using-dial-up-ipsec-with-ta--ff051ae2

| Description | This article describes how to set up an IPsec dial-up tunnel for remote users with split tunneling on 7.6 FortiGate GUI. | 
| Scope | FortiGate 7.6.0. | 
| Solution | In this particular example, IPsec dial tunnel is configured with Split tunneling enabled which means clients' traffic, which is configured through the firewall policy, will be routed over the IPsec VPN tunnel. First, create a new address for the LAN network: **Policy & Objects -> Addresses**  and select 'Create New'.   **Configure the IPsec Dial-Up VPN using FortiClient:** Navigate to **VPN -> VPN tunnels -> Create new -> IPsec tunnel from template.** Give it a name and set the template to **Remote Access.**    Select the relevant **VPN client type.**  In this example, FortiClient is used.  Assign the IP range that the devices will get upon **connecting through FortiClient.**   Select the **Authentication method** .Here, the pre-shared key is utilized. This pre-shared key will also be used under **FortiClient settings later as well.**  Select the User group which will be connecting through FortiClient for remote access.    Now, select the incoming interface on which the IPsec traffic will come and select the Local interface and Local address:    Create a tunnel. A summary page will appear, listing the objects that have been added to the FortiGate’s configuration by the wizard.    **Note:** If more than one IPsec dial-up VPN is configured on the same server interface (WAN in this case), each phase1 of those VPNs needs to be defined with a unique peer ID to distinguish between the multiple tunnels that the remote client is going to connect with.  Refer to this article for enabling Peer ID: Technical Tip: How to use Peer IDs to select an IPSec dialup tunnel on a FortiGate configured with multiple dialup tunnels.  Check the firewall address created by the wizard. To do so, navigate to **Policy & Objects -> Addresses.**    Check the VPN interface created by the wizard. To do so, navigate to **Network -> Interfaces.**   Check the IPv4 firewall policy created by the wizard. To do so, navigate to **Policy and Objects -> Firewall policy.**    **FortiClient configuration:** To add the VPN connection, navigate to **Remote Access and select 'Add a new connection'.** Set the VPN to 'IPsec VPN' and set 'Remote Gateway' to the 'FortiGate IP address (wan IP)'. Set 'Authentication Method' to' Pre-Shared Key'.    If the user is ever unable to connect, **double-check the advanced settings** under the FortiClient and**match them under the IPsec settings on FortiGate** : they must be the same on both sides.  After matching the settings, the user should be able to connect through FortiClient.   Verify the routes received on the Windows machine, and run the following command on the command prompt on Windows:  **route print**    **Related document:**  FortiClient as a dial-up client | 

Enter your E-mail address. We'll send you an e-mail with instructions to reset your password.
