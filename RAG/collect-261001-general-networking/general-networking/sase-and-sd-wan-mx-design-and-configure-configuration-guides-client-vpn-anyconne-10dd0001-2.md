---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-10dd0001-2
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-10dd0001"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-10dd0001.md
source_anchor: ""
source_lines: [98, 179]
sha256: e355de9bf8948165dd047511e01949d3a439ecfe028a8046430f0dc9f5e73371
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-client-vpn-anyconne-10dd0001

Configure a RADIUS Connection Request
- 
    In the NPS server console, navigate to Policies > Connection Request Policies. Right-click the Connection Request Policies folder and select New.
- 
    In the Connection Request Policy Wizard, enter a policy name and select the network access server type Unspecified, then press Next.
- 
    Click Add to add conditions to your policy. Access request messages will need to meet these conditions to be allowed access.
- 
    From the list of conditions, select the option for NAS-Port-Type. Select VPN Virtual, and then press Next
- 
    Press Next on the next three pages of the wizard to leave the default settings intact.
- 
    Review the settings, then press Finish. 
Configure a RADIUS Network Policy.
- 
    In the left-side pane of the NPS server console, right-click the Network Policies option and select New.
- 
    In the Network Policy Wizard, enter a policy name and select the network access server type Unspecified, then press Next.
- 
    Click Add to add conditions to your policy.
- 
    From the list of conditions, select the option for Windows Groups. Click Add Groups and enter the name you would like to grant client VPN permission to.
- 
    From the list of conditions, select the option for NAS-Port-Type. Select VPN Virtual and press Next.
- 
    Leave the default settings on the Specify Access Permission page and press Next.
- 
    Deselect all checkboxes and select Unencrypted authentication (PAP, SPAP). An informational box will be displayed, press No to continue, and press Next.
- 
    Press Next on the next two pages of the wizard to leave the default settings intact.
- 
    Review the settings, then press Finish. 
Active Directory Authentication
Use this option if user authentication should be done with Active Directory domain credentials. You will need to provide the following information:
- 
    Short domain: The short name of the Active Directory domain.
- 
    Server IP: The IP address of an Active Directory server on the MX LAN.
- 
    Domain admin: The domain administrator account the MX should use to query the server.
- 
    Password: Password for the domain administrator account.
For example, considering the following scenario: Users in the domain test.company.com should be authenticated using an Active Directory server with IP 172.16.1.10. Users normally log in to the domain using the format "test/username" and you have created a domain administrator account with username "vpnadmin" and password "vpnpassword".
- 
    The Short domain would be "test"
- 
    The Server IP would be 172.16.1.10.
- 
    The Domain admin would be "vpnadmin"
- 
    The Password would be "vpnpassword"
Note: Only one AD server can be specified for authentication with AnyConnect at the moment. The MX does not support mapping group policies via Active Directory for users connecting through the client VPN. Refer to this document for more information on integrating with client VPN.
Note: TLS is required to communicate with the Domain Controller. For more details regarding TLS requirements please review our Certificate Requirements for TLS documentation.
Multi-Factor Authentication with RADIUS or Active Directory as a Proxy
MFA can be configured with your RADIUS or Active Directory server. The MFA challenge takes place between the RADIUS/Active Directory/IdP and the user. The MX will not pass any OTP or PINs between the user and RADIUS. When a user connects to the MX, they are prompted for a username and password. The MX passes these credentials to the RADIUS or AD server, which then challenges the user directly (not through the MX). The user responds to the RADIUS or AD server, possibly via push notification or other means. The RADIUS or AD server then informs the MX that the user has successfully authenticated. Only then is the user allowed access to the network. Refer to this document for more information on authentication.
RADIUS Timeout
The default RADIUS timeout is three seconds. This is how long the AnyConnect server will wait for a response from the RADIUS server before failing over to a different RADIUS server or ignoring the response entirely. To support two-factor authentication, you can increase the RADIUS timeout by modifying the RADIUS timeout field on the Cisco Secure Client Settings page. The configurable timeout range is 1 - 300 seconds.
Systems Manager Sentry Authentication
Sentry AnyConnect VPN is a special Cisco Meraki integration between MX and Systems Manager (SM) enrolled devices. This allows secure and automatic certificate-based Always-On AnyConnect VPN for SM managed devices. SM managed devices will be sent all the necessary configurations, certificates, and app settings for an Always-On VPN tunnel back to the MX. For SM enrolled devices, the end users are not interrupted with any authentication/setup steps and the VPN tunnel will open automatically.
Note: SM Sentry AnyConnect Authentication is currently in beta and requires SM or MX product teams to enable it. Please request access here or please await the general availability release.
Systems Manager Sentry Authentication Setup
On the Security & SD WAN > Client VPN > Cisco Secure Client Settings page, find the Authentication & Access section. Set the Authentication Type to Systems Manager Sentry Authentication. Then, click on Add Sentry Network and choose one or more Systems Manager networks and the desired tag scoping to allow all or a subset of the devices enrolled in SM to obtain this configuration.
Firmware requirement: MX 17.10 or higher
In the above image, all devices tagged with sentry-anyconnect-vpn in the SM network called Cisco-on-Cisco will receive the automatic Systems Manager Sentry AnyConnect VPN configuration.
Note: The SM network and MX network must be part of the same Meraki organization.
SM Sentry Technology
The SM enrolled devices will automatically be delivered the settings necessary to connect via AnyConnect to the MX host. The Meraki Cloud Certificate Authority (CA) handles all the certificate management. Administrators do not need to be burdened with the complexities of certificate management, while also obtaining all of the certificates, security, and authentication benefits. At a high level, the automatic Sentry AnyConnect VPN configuration to managed SM devices contains three main settings:
- SCEP certificate payload used for certificate-only authentication to MX via Meraki Cloud CA.
- VPN payload with AnyConnect Always-On enabled. This will automatically use the host:port configured on the MX Client VPN page.
- Cisco Secure Client application with necessary managed app configurations. This app is necessary to enable the AnyConnect VPN app's networking capabilities on devices.
iOS and iPadOS
Currently, the Sentry AnyConnect VPN beta supports iOS and iPadOS only. To request beta access: fill out this form.
Always-On VPN tunnel requires Supervision.
After the Cisco Secure Client app is installed, the device will prompt end users for notifications. This can be suppressed (automatically accepted so end users are not interrupted) by adding an optional Notification profile payload for Cisco Secure Client app in the SM > Settings page:
If end users are able to remove the Cisco Secure Client app, they can disable the Always-On VPN tunnel. To prevent end users' ability to remove the app (and thus keep the tunnel enabled), add a Home Screen Layout, set the app to be hidden via Restriction, and/or uncheck the removable option for the Cisco Secure Client app on the SM > Apps page:
To install the Cisco Secure Client app completely silently, a VPP license will need to be added. It is recommended to use VPP Device Assignment on the SM > Apps page:
Android
Support coming soon.
macOS
Support coming soon.
Windows
Support coming soon.
