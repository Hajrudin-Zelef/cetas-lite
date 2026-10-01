---
id: collect-261001-meraki/meraki/blog-network-security-secure-remote-access-through-meraki-z3-teleworker-296f7cba
title: "blog-network-security-secure-remote-access-through-meraki-z3-teleworker-296f7cba"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/blog-network-security-secure-remote-access-through-meraki-z3-teleworker-296f7cba.md
source_anchor: ""
source_lines: [1, 49]
sha256: a905d5a7a35d58e7151a8c97ad282d021c72e07fe7a1062f649b7a4eba9fcdb1
---

# blog-network-security-secure-remote-access-through-meraki-z3-teleworker-296f7cba

As remote work has surged due to the COVID-19 pandemic, many IT administrators have begun setting up remote access solutions for their employees. One of the common options available in the market that allows IT to extend the company LAN is the Meraki Z3 Teleworker Gateway.
This relatively new Meraki appliance effectively creates a VPN tunnel from the homes of remote workers back to the organizational headquarters or branch office. It’s a fairly simple item to deploy, even for the non-technical employee.
The use of the Meraki Z-series appliances, however, raises new security concerns, since nearly any device – like your child’s laptop – can connect to the appliance and tunnel back to the corporate network. This vulnerability thus requires that you be able to extend the same level of network access control provided on the company LAN to employees’ homes, and ultimately the Meraki Z3 appliance itself. After all, the Z3 should allow only corporate devices to connect and ensure non-authorized devices will be blocked.
Overlaying Remote Access Control Over a Meraki Z3 Teleworker Gateway
Portnox CLEAR delivers enterprise NAC as a cloud service. The platform’s cloud-delivered functionality offers improved scalability and accessibility, making it ideal for organizations with large remote workforces. As we’ll detail below, Portnox CLEAR can extend the same network visibility and control implemented on the LAN to the Meraki Z3 Teleworker Gateway network:
- One-click Cloud-Based Radius Server setup.
- 802.1X authentication:
- 
  - EAP-TLS: Corporate PKI or CLEAR root CA.
  - EAP-PEAP / EAP-TTLS: User Credentials.
- Authentication repositories integrations: on-premises Active Directory, Azure AD, G Suite, OKTA or CLEAR internal (based on email address).
- Continuous endpoint risk awareness for Windows, MAC OS X and Linux. Some examples: verify antivirus is running and updated, firewall is on, required certificates, peripherals devices not connected, required or forbidden application, geo-location awareness.
- Real time risk-based actions: alert and deny access based on end point risk assessment.
- Real time end point remediation capabilities. some examples: ensuring firewall is always on, terminate process, disable internet sharing, run login or periodic script.
How it Works: Setting up Portnox CLEAR with Your Meraki Z3
Integrating your Meraki Z3 Teleworker Gateway with Portnox can be completed in just a few minutes. In the following steps, we will configure the Meraki Z3 to be secured and protected based on RADIUS and 802.1X authentication:
Preliminary Actions
Before configuring Meraki Z3, you need to do the following:
- Verify your organization is registered on Portnox CLEAR Cloud Services (start your unlimited free account at any time or take advantage of our temporary offer of full Portnox CLEAR version free for three months).
- In the CLEAR portal, go to Settings > Services and expand CLEAR RADIUS Service.
- If the Enable Cloud RADIUS checkbox is not checked, click Edit and check the Enable Cloud RADIUS checkbox.
- Note the RADIUS server details required when configuring Meraki Z3 access:
- 
  - Cloud RADIUS IP – this is the IP address of the CLEAR RADIUS server
  - Authentication port
  - Shared Secret – this is the RADIUS client shared secret
Allow Access to Wired Networks in CLEAR
The next step is to allow, in the CLEAR portal, access to wired networks you will be securing.
- Navigate in the portal to Groups.
- Edit the default “Unassigned” group or create a new security group.
- Whether you are creating or editing a group, in Group Settings check the Enable wired access using 802.1x authentication for devices in this group check box.
Configuring the Meraki Z3 Teleworker – Wired Ports
Next, we configure the Meraki Z3 Teleworker wired ports to be secured and protected based on CLEAR RADIUS authentication:
- In the Meraki portal, navigate to Teleworker gateway > Configure > Addressing & VLANs, and verify that the VLANs are enabled in the Routing section.
- In the Per-port VLAN Settings, edit the relevant port/s:
A. Set the Enabled option to Enabled.
B. Set the Type to Access.
C. Select the relevant VLAN.
D. Select the Access policy type: 802.1x or MAC authentication bypass.
E. In the Radius servers, click add radius server and enter the following CLEAR RADIUS server details, which you noted in Preliminary Actions, step 4:
- 
  - 
    - 
      - In host, enter the Cloud RADIUS IP.
      - In port, enter the Authentication port number.
      - In secret, enter the Shared Secret.
  - 
- 
F. Click Update.
