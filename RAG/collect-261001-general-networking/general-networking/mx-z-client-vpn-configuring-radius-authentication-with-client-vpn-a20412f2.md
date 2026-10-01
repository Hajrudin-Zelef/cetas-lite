---
id: collect-261001-general-networking/general-networking/mx-z-client-vpn-configuring-radius-authentication-with-client-vpn-a20412f2
title: "mx-z-client-vpn-configuring-radius-authentication-with-client-vpn-a20412f2"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/mx-z-client-vpn-configuring-radius-authentication-with-client-vpn-a20412f2.md
source_anchor: ""
source_lines: [1, 78]
sha256: b09bcd5e11b5bc720b30ebfa0c3e86df185229a6065f56ced1a118cd0ef8fb3d
---

# mx-z-client-vpn-configuring-radius-authentication-with-client-vpn-a20412f2

Configuring RADIUS Authentication with Client VPN IPsec
Cisco Meraki Client VPN can be configured to use a RADIUS server to authenticate remote users against an existing userbase.
This article outlines the configuration requirements for RADIUS-authenticated Client VPN, as well an example RADIUS configuration steps using Microsoft NPS on Windows Server 2008.
RADIUS Configuration
While any RADIUS server can be used, the following configuration requirements are necessary for Client VPN integration:
- RADIUS must be configured to allow PAP (unencrypted authentication)
Note: Communication between the client and the MX will be encapsulated within IPsec, so this does not mean that client communication is unencrypted.
- The MX/Z1's IP address must be configured on the server as a RADIUS client/authenticator, with a shared secret.
Please refer to your RADIUS server vendor's documentation for configuration specifics.
 
For troubleshooting guidance, please follow RADIUS Issue Resolution Guide.
Example RADIUS Server Configuration (Windows NPS + AD)
The following example configuration outlines how to configure an existing Windows 2008 server, running Network Policy Server (NPS) alongside Active Directory:
- Add the MX Security Appliance as a RADIUS client on the NPS server.
- Configure a RADIUS Connection Request in NPS.
- Configure a RADIUS Network policy in NPS.
Note: This configuration assumes that NPS is already running on the Windows server. Please refer to Microsoft documentation for assistance in running NPS.
Add MX Security Appliance as RADIUS clients on the NPS server
In order for the MX to act as an authenticator for RADIUS, it must be added as a client on NPS.
- Open the NPS Server Console by going to Start > Programs > Administrative Tools > Network Policy Server.
- In the Left pane, expand the RADIUS Clients and Servers option.
- Right-click the RADIUS Clients option and select New.
- Enter a Friendly Name for the MX Security Appliance or Z1 Teleworker Gateway RADIUS Client.
- Enter the IP Address of your MX Security Appliance or Z1 Teleworker Gateway. This IP will differ depending on where the RADIUS server is located:
    
  - On a local subnet - Use the IP address of the MX/Z1 on the subnet shared with the RADIUS server.
  - Over a static route - Use the IP address of the MX/Z1 on the subnet shared with the next hop.
  - Over VPN - Use the IP address of the MX/Z1 on the highest-numbered VLAN in VPN.
- Create and enter a RADIUS Shared Secret (make note of this secret - we will need to add this to the Dashboard).
Note: Currently only ASCII characters are supported for RADIUS shared secrets - Unicode characters will not work correctly.
- Press OK when finished.
For additional information or troubleshooting assistance, please refer to Microsoft documentation.
Configure a RADIUS Connection Request
- In the NPS Server Console, navigate to Policies > Connection Request Policies. Right-click the Connection Request Policies folder and select New.
- In the Connection Request Policy Wizard, enter a Policy Name and select the Network Access Server Type unspecified then press Next.
- Click Add to add conditions to your policy. Access Request messages will need to meet these conditions to be allowed access.
- From the list of conditions select the option for Framed-Protocol. Press Add and place a check next to the PPP option then press Ok.
- Click Add to add another condition and select the option for CallingStationID. Enter CLIENTVPN into the text box and press Next.
- On the next three pages of the wizard we will leave the default settings. Press Next on these pages to continue.
- Review the settings then press Finish.
Configure a RADIUS Network Policy
- In the Left pane of the NPS Server Console, right-click the Network Policies option and select New.
- In the Network Policy Wizard enter a Policy Name and select the Network Access Server type unspecified then press Next.
- Click Add to add conditions to your policy.
- From the list of conditions, select the option for Windows Groups. Click Add Groups and enter the name of Windows Group you would like to give Client VPN permission.
- Click Add to add an additional condition. Select the option for Framed Protocol, press add and check the PPP option then press Ok
- Click Add to add a final condition. Select the option for CallingStationID. Enter CLIENTVPN into the text box and press Next.
Note: Some versions of Windows Server require that the CallingStationID is omitted. If you are unable to establish connectivity, remove the CallingStationID and leave the field blank.
- Leave the default settings on the Specify Access Permission page and press Next.
- Deselect all checkboxes and select Unencrypted authentication (PAP, SPAP). An informational box will be displayed press No to continue and press Next.
- The next 2 pages of the wizard we will leave the default settings. Press Next on these pages to continue.
- Review the settings then press Finish.
If utilizing Microsoft RADIUS Authentication with Azure MFA, it is advisable to adjust the RADIUS timeout value on the AD server to 300 seconds. Timeouts related to this issue will present a reason code 9 in the Windows event log.
Dashboard Configuration
Once a RADIUS server has been configured appropriately, the following steps outline how to configure Client VPN to use RADIUS:
- Log onto the Cisco Meraki Dashboard and navigate to Security & SD-WAN > Configure > Client VPN.
- Select the option to enable the Client VPN Server.
- Set the Client VPN Subnet. This will be a unique IP subnet offered to clients connecting to the MX Security Appliance via a Client VPN connection.
- Specify the DNS servers.
- Enter a shared secret that will be used by the client devices to establish the VPN connection.
Note: This is a different value from the RADIUS shared secret.
- Select RADIUS as the Authentication method.
- Click the Add a RADIUS Server link.
    
  - Enter your RADIUS Host IP Address.
  - Enter the RADIUS Port that the MX Security Appliance will use to communicate to the NPS server. The default port is 1812.
  - Enter the RADIUS Shared Secret (established when the MX was added as an authenticator).
- Enable or disable RADIUS message-authenticator verification.
- Configure RADIUS timeout.
- Configure the RADIUS retry amount.
- Click Save changes.
Regarding CVE-2024-3596: If RADIUS message-authenticator verification is enabled for IPsec Client VPN but not on the NPS, the RADIUS server omits the Message-Authenticator(80) attribute, causing silent 802.1x authentication failures on the MX. To resolve this, either disable the feature on the MX or enable it on the NPS.
Note: IPsec and AnyConnect share the same configured RADIUS and Active directory servers. Please refer to AnyConnect on the MX Appliance.
Client Configuration
Installation of additional software is not required on client devices. The Cisco Meraki Client VPN solution uses L2TP over IPsec, which is supported by almost all device's built-in native clients.
Please refer to our Client VPN documentation for client configuration instructions.
Additional Resources
For additional information about Client VPN, please refer to the following articles:
