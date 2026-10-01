---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-12-administration-guide-785501-forticlient-as-dialup-clie-4d444a34-1
title: "document-fortigate-7-4-12-administration-guide-785501-forticlient-as-dialup-clie-4d444a34"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-12-administration-guide-785501-forticlient-as-dialup-clie-4d444a34.md
source_anchor: ""
source_lines: [1, 81]
sha256: fb64dd9138f06b5a14dbe110be53bdf02181a8f862bb38457c4c68dc7617e845
---

# document-fortigate-7-4-12-administration-guide-785501-forticlient-as-dialup-clie-4d444a34

FortiClient as dialup client
FortiClient as dialup client
This is a sample configuration of dialup IPsec VPN with FortiClient as the dialup client.
You can configure dialup IPsec VPN with FortiClient as the dialup client using the GUI or CLI.
If multiple dialup IPsec VPNs are defined for the same dialup server interface, each phase1 configuration must define a unique peer ID to distinguish the tunnel that the remote client is connecting to. When a client connects, the first IKE message that is in aggressive mode contains the client's local ID. FortiGate matches the local ID to the dialup tunnel referencing the same Peer ID, and the connection continues with that tunnel.
For VPN tunnels using IKEv2, network-overlay and network-id should be configured in place of peer-id:
config vpn ipsec phase1-interface
    edit <vpn-tunnel-name>
        set network-overlay enable
        set network-id <ID>
    next
end
                                            FortiClient EMS 7.4.4 and later support configuring the Network ID on the IPsec VPN Phase 1 settings in the GUI. For earlier versions, refer to the XML reference guide for IKE settings for details on configuring the <networkid> setting.
|  | Starting with FortiClient 7.4.4, IKEv1 is no longer supported on the client. Therefore, plan accordingly when choosing your IKE version. Use IKEv2 if you plan on deploying FortiClient 7.4.4 and later. Also, FortiClient 7.4.4 does not support IPv6. Use FortiClient 7.4.6 or later. | 
To configure IPsec VPN with FortiClient as the dialup client on the GUI:
- 
                                                    Configure a user and user group. 
  - 
                                                            Go to User & Authentication > User Definition to create a local user vpnuser1.
  - 
                                                            Go to User & Authentication > User Groups to create a group vpngroup with the member vpnuser1.
- 
                                                            
- 
                                                    Go to VPN > IPsec Wizard and configure the following settings for VPN Setup: 
  - 
                                                            Enter a VPN name.
  - 
                                                            For Template Type, select Remote Access.
  - 
                                                            For Remote Device Type, select Client-based > FortiClient.
  - 
                                                            Click Next.
- 
                                                            
- 
                                                    Configure the following settings for Authentication: 
  - 
                                                            For Incoming Interface, select wan1.
  - 
                                                            For Authentication Method, select Pre-shared Key.
  - 
                                                            In the Pre-shared Key field, enter your-psk as the key.
  - 
                                                            From the User Group dropdown list, select vpngroup.
  - 
                                                            Click Next.
- 
                                                            
- 
                                                    Configure the following settings for Policy & Routing: 
  - 
                                                            From the Local Interface dropdown menu, select lan.
  - 
                                                            Configure the Local Address as local_network.
  - 
                                                            Configure the Client Address Range as 10.10.2.1-10.10.2.200.
  - 
                                                            Keep the default values for the Subnet Mask, DNS Server, Enable IPv4 Split tunnel, and Allow Endpoint Registration.
  - 
                                                            Click Next.
- 
                                                            
- 
                                                    Adjust the Client Options as needed, then click Create.
- 
                                                    Optionally, define a unique Peer ID in the phase1 configuration: 
  - 
                                                            Go to VPN > IPsec Tunnels and edit the just created tunnel.
  - 
                                                            Click Convert To Custom Tunnel.
  - 
                                                            In the Authentication section, click Edit.
  - 
                                                            Under Peer Options, set Accept Types to Specific peer ID.
  - 
                                                            In the Peer ID field, enter a unique ID, such as dialup1.
  - 
                                                            Click OK.
- 
                                                            
