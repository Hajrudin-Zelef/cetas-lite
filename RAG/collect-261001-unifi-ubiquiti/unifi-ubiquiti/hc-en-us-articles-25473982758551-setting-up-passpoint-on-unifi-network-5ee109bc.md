---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-25473982758551-setting-up-passpoint-on-unifi-network-5ee109bc
title: "hc-en-us-articles-25473982758551-setting-up-passpoint-on-unifi-network-5ee109bc"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-25473982758551-setting-up-passpoint-on-unifi-network-5ee109bc.md
source_anchor: ""
source_lines: [1, 23]
sha256: a93fc740a221eb8d76eeed7ef378be21c665db83270d2c3801a9b304dc5ebcbf
---

# hc-en-us-articles-25473982758551-setting-up-passpoint-on-unifi-network-5ee109bc

Setting Up Passpoint on UniFi Network
Passpoint, built on the 802.11u standard, simplifies public WiFi connectivity by enabling seamless, secure connections for guest devices across multiple venues without the need for guest portals or complex connection processes. Third-party Passpoint providers such as Google Orion, IronWifi, and OpenRoaming seamlessly integrate with UniFi Network. 
Many clients obtain Passpoint profiles to join compatible networks via a cellular carrier, distributed with their SIM. Clients with a valid SIM simply need to be within range of a corresponding Passpoint network to automatically join the network. Other Passpoint profiles can be obtained via an app or website, and once the profile is added to a device, clients will automatically join any matching network, no matter where they are. For example, a hospitality provider's Passpoint profile can provide connectivity at their properties around the world, without requiring any further client action.
Requirements
To configure Passpoint in UniFi Network, ensure the following:
- UniFi Network version 8.4.54 or higher
- Access Point (AP) firmware 6.6.75 / 7.0.63 or higher
- A configured RADIUS profile in UniFi Network.
Please note:
- A third party RADIUS server is currently required for clients to authenticate against.
- Passpoint-enabled SSIDs do not support captive portals or redirects.
- To enable cellular data offload on your Passpoint WiFi network, make sure you have correct MCC/MNC information for your region's cellular carriers. These are often provided by your Passpoint provider.
Configuring RADIUS
- Navigate to Settings > Profiles > RADIUS.
- Click Create New.
- Configure your authentication and accounting servers according to your provider's guidelines.
Configuring your WiFi Network
- Go to Settings > WiFi > Create New.
- Set Hotspot 2.0 to Passpoint.
- Configure the settings provided by your Passpoint provider. Ensure that WPA2 or WPA3 Enterprise is selected.
- Under RADIUS Profile, select the RADIUS profile you configured earlier.
- If required by your provider, adjust the NAS ID settings under the RADIUS Profile section.
Once the network configuration is saved and provisioned on your APs, compatible devices should automatically connect.
