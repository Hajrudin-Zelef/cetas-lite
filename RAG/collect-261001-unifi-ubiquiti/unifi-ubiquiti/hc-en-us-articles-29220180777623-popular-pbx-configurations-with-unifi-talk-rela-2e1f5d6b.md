---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-29220180777623-popular-pbx-configurations-with-unifi-talk-rela-2e1f5d6b
title: "hc-en-us-articles-29220180777623-popular-pbx-configurations-with-unifi-talk-rela-2e1f5d6b"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-29220180777623-popular-pbx-configurations-with-unifi-talk-rela-2e1f5d6b.md
source_anchor: ""
source_lines: [1, 43]
sha256: 6dbfc58a77443f49bba9e84f091e3c028b17e07568cfc02462b9318387e1da7c
---

# hc-en-us-articles-29220180777623-popular-pbx-configurations-with-unifi-talk-rela-2e1f5d6b

Popular PBX Configurations with UniFi Talk Relay
UniFi Talk Relay has prepared configuration profiles to make it easier for IT admin to configure Talk VoIP phones in standalone mode.
Note: To configure PBX, you must be using UniFi Talk Relay, a service designed for remote management of UniFi Talk Gen 3 phones adopted to third-party PBX systems. If you’re seeking a plug-and-play solution that also supports your own SIP provider, see UniFi Talk for more information.
Nextiva
- Navigate to the Voice page and open the Phones & Devices tab.
- Click Add Device.
- Select Phone not listed (Generic SIP) during the phone setup step.
- Copy the configuration details and click Save and Continue.
- Assign the phone to a user.
- Enter the configuration details in Talk Relay:
  - SIP Username: Paste into the Username field.
  - Domain: Paste into the PBX Domain field.
  - Authentication Name: Paste into the Register Name field.
  - Authentication Password: Paste into the Password field.
RingCentral
- Log in to the Admin Portal.
- Navigate to Phone System > Phones & Devices > User Phones.
- If the device is not listed, click Add Device and select Existing Phone.
- Click Set Up and Provision, then choose Set up manually using SIP.
- Enter the configuration details in Talk Relay:
  - SIP Domain: Paste into the PBX Domain field.
  - Remote SIP Port: Paste into the PBX Port field.
  - User Name: Paste into the Username field.
  - Password: Paste into the Password field.
  - Authorization ID: Paste into the Register Name field.
  - Outbound Proxy: Enable and paste into the Proxy Host field.
  - Outbound Proxy Port: Paste into the Proxy Port field.
voip.ms
- (Optional) Create subaccounts under Subaccounts > Create Sub Account for your users.
- Enter the configuration details in Talk Relay:
  - Register Name and Username: Use the format #####_subaccount (replace with your SIP main account or subaccount).
  - Password: Paste into the Password field.
  - Codecs: Select PCMU only.
  - PBX Domain: Choose the server closest to your location from this list.
CallCentric
- Create or modify an extension on the Extensions page.
- Enter the configuration details in Talk Relay:
  - sip.callcentric.net: Paste into the PBX Domain field.
  - 5060: Paste into the PBX Port field.
  - SIP Username: Paste into the Username field.
  - SIP Password: Paste into the Password field.
  - sip.callcentric.net: Enable Outbound Proxy and paste into the Proxy Host field.
  - 5060: Paste into the Proxy Port field.
