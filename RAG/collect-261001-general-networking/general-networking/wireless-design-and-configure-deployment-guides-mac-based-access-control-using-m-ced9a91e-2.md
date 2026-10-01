---
id: collect-261001-general-networking/general-networking/wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e-2
title: "wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e.md
source_anchor: ""
source_lines: [55, 104]
sha256: 41479cbdfdd07f55c17c6ea1cda96a6070b3aa619d26ab24622f150e33469282
---

# wireless-design-and-configure-deployment-guides-mac-based-access-control-using-m-ced9a91e

 You may generate a self-signed certificate for testing/lab purposes. However, clients will not trust a self-signed certificate and you will need to disable server validation to connect.
We do not recommend this option for production deployment, due to dramatically reduced security.
Once a certificate is acquired, refer to the Import a Certificate section of Microsoft's documentation for instructions on how to import a certificate.
Add Access points as RADIUS Clients on the NPS Server
In this scenario, access points communicate with clients and receive their domain credentials, which the access point then forwards to NPS. In order for the NPS to process an access point's RADIUS access-request message, you must first add the access point as a RADIUS client/authenticator by its IP address. Since only gateway access points have an IP address on the LAN, all gateway access points in the network must be added to NPS as RADIUS clients.
To quickly gather all gateway access point's LAN IP addresses, navigate to Wireless > Monitor > Access points in dashboard, ensure that the "LAN IP" column has been added to the table, and take note of all LAN IPs listed. Access points with a LAN IP of "N/A" are repeaters and they do not need to be added as RADIUS clients:
Once a list of gateway access point's LAN IPs has been gathered, please refer to Microsoft's documentation for instructions on adding each access point as a client in NPS. Take note of the shared secret configured in NPS, which is referenced in the dashboard.
Note: To save time, entire subnets can also be added to NPS as RADIUS clients, and any requests coming from that subnet will be processed by NPS. This is only recommended if all access points are on their own management VLAN and subnet, to reduce security risks.
In many cases each RADIUS authenticator must be added to the RADIUS authentication server such as Microsoft NPS or Cisco ISE. For VPN concentration and concentrated Layer 3 roaming SSIDs, only concentrators would need to be added to the RADIUS authentication server.
Configure a Policy in NPS to Support PEAP-MSCHAPv2
NPS must be configured to support PEAP-MSCHAPv2 as its authentication method.
This is accomplished in three steps, outlined below for NPS in Windows Server 2008:
- Create an NPS Policy
- Change the Policy Process Order
- Disable Auto Remediation
Creating an NPS Policy
- Open the Network Policy Server console.
- Select NPS(Local), so you see the Getting Started pane.
- Select RADIUS server for 802.1X Wireless or Wired Connections in the Standard Configuration drop down.
- Click Configure 802.1X to begin the Configure 802.1X Wizard.
- When the Select 802.1X Connections Type window appears select the radio button Secure Wireless Connections and type a Name: for your policy or use the default. Click Next.
- Verify the access points you added as RADIUS clients on the Specify 802.1X switches window. Click Next.
- For Configure an Authentication Method select Microsoft: Protected EAP (PEAP).
- Click Configure to review the Edit Protected EAP Properties. The server certificate should be in the Certificate issued drop down. Make sure Enable Fast Reconnect is checked and EAP type is Secure password (EAP-MSCHAPv2). Click OK. Click Next.
- When the Specify User Groups window appears click Add.
- Type or find the Domain Users group. This group should be located in the same domain as your RADIUS server.
Note: If RADIUS is being used for Machine Authentication, find the Domain Computers group instead.
- When the group is added click OK. Click Next.
- Click Next on Configure a Virtual LAN (VLAN) window.
- When then Completing New IEEE 802.1X Secure Wired and Wireless Connections and RADIUS clients appears click Finish.
Change the Policy Process Order
- Navigate to Policies>Connection Request Policies. Right click the wireless policy and Move Up so it is process first.
- Navigate to Policies>Network Policies. Right click the wireless policy and Move Up so it is process first.
Disable Auto Remediation
- Navigate to Policies>Network Policies. Right click the wireless policy and select Properties.
- On the Setting tab for the policy uncheck the box Enable auto-remediation of client computers and click OK.
The following image outlines an example of an NPS policy that supports user authentication with PEAP-MSCHAPv2:
(Optional) Deploy a PEAP Wireless Profile using Group Policy
For a seamless user experience, it may be ideal to deploy a PEAP wireless profile to domain computers so users can easily associate with the SSID. Though optional for user auth, this is strongly recommended for machine authentication.
The following instructions explain how to push a PEAP wireless profile to domain computers using a GPO, on a Domain Controller running Windows Server 2008:
- Open the domain Group Policy Management snap-in.
- Create a new GPO or use an existing GPO.
- Edit the GPO and navigate to Computer Configuration > Policies > Windows Settings > Security Settings > Public Key Policies > Wireless Network (IEEE 801.X) Policies.
- Right Click Wireless Network (IEEE 801.X) Policies and choose Create a New Windows Vista Policy.
- Provide a Vista Policy Name.
- Click Add for Connect to available networks.
- Choose Infrastructure.
- On the Connection tab, provide a Profile Name and enter the SSID of the wireless network for Network Name(s). Click Add.
- Click the Security tab. Configure the following:
    
