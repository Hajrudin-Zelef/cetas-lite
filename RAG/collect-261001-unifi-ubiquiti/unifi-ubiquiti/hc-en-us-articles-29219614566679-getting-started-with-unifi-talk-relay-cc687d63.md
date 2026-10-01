---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-29219614566679-getting-started-with-unifi-talk-relay-cc687d63
title: "hc-en-us-articles-29219614566679-getting-started-with-unifi-talk-relay-cc687d63"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "packaging"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-29219614566679-getting-started-with-unifi-talk-relay-cc687d63.md
source_anchor: ""
source_lines: [1, 74]
sha256: 54ca58ee151c75584791a7f25c52abef39c02ca0725fdf10aa39c35fb4342a97
---

# hc-en-us-articles-29219614566679-getting-started-with-unifi-talk-relay-cc687d63

Getting Started with UniFi Talk Relay
UniFi Talk Relay allows seamless remote management of UniFi Talk Gen 3 phones adopted to third-party PBX systems. This guide provides a step-by-step walkthrough to set up UniFi Talk Relay with UniFi’s hosting service, enabling administrators to securely and efficiently manage their Talk phone fleet from anywhere.
If you’re seeking a plug-and-play solution that also supports your own SIP provider, see UniFi Talk for more information.
Requirements
Before you begin, ensure that you have the following:
- Official UniFi Hosting Subscription
- UniFi Talk Gen3 Phones
Standard Setup Method
The standard configuration method is designed for real-time setup and configuration of Talk Phones. See Pre-Provisioning with ZTP Codes below for more information on automating your deployment rollouts.
- 
Prepare UniFi Talk Relay:
Ensure you have an Official UniFi Hosting instance that is running UniFi Talk Relay. To view installed applications, go to Site Manager, open your Official UniFi Hosting instance, and go to Settings > Control Plane.
- 
Connect Your Talk Gen3 Phones:
 Plug the phones into a PoE switch for power and Internet connectivity.
  - Note: Phones must be in a Factory Default State during setup.
- 
Enable Third-Party PBX:
 Once the phones boot, select the button on the phones’ screen for Third-Party PBX.
  - The "Third-Party PBX" button triggers a firmware update, which is required before phones can be configured with Talk Relay.
- 
Scan the QR Code:
Use a cell phone to scan the QR code displayed on the phone and follow the prompts to complete the provisioning process.
- 
Configure Your PBX Server:
 In the Talk Relay application, go to Devices.
  - Verify your new phone is now showing.
  - Click Add in the phone’s Configuration column to open its device panel.
  - Beside PBX, click New.
  - Select your PBX provider from the available templates.
  - Enter any additional information requested, such as register name, password, proxy, and host.
  - Click Save & Sync.
Pre-Provisioning with ZTP Codes
ZTP Codes allow you to pre-configure phones so that they can be automatically provisioned with all PBX and management settings once they are powered on and connected to a network.
Pre-Provisioning
- 
Prepare UniFi Talk Relay:
Ensure you have an Official UniFi Hosting instance that is running UniFi Talk Relay. To view installed applications, go to Site Manager, open your Official UniFi Hosting instance, and go to Settings > Control Plane.
- 
Preconfigure PBX Profiles (Optional):
Navigate to Settings > PBX and create profiles for your third-party PBX providers. This allows you to apply configurations to devices in advance.
- 
Enter ZTP Code:
 Click the plus button in the upper right corner and select Add Devices. You can find the device’s ZTP Code either in the Device’s packaging, or on its screen when it is plugged in.
  - Note: If you have already used the ZTP code to add phones to your Site Manager inventory, they can be assigned to this specific site by clicking Add from Inventory.
- 
Apply PBX Configuration: Navigate to the Configuration Tab of your device’s settings and:
  - Select a preconfigured PBX profile, or create a new one by entering the required details (register name, password, proxy, and host).
  - Save the configuration. Once the device powers on and connects to the Internet, it will retrieve the configuration automatically from the Talk Relay.
Deploying
- 
Connect Your Talk Gen3 Phones:
Plug the phones into a PoE switch for power and Internet connectivity.
- 
Enable Third-Party PBX:
 Once the phones boot, select the button for Third-Party PBX.
  - To configure Zoom Phone, see the instructions here.
Troubleshooting
If devices are not appearing in the Devices tab:
- Verify the ZTP code is entered correctly.
- Make sure ZTP service is enabled on your Talk phone by tapping on the Navigation captive button on your Talk phone’s screen, going into Settings > System > Zero Touch Provisioning.
- Check your network’s firewall settings to ensure devices can communicate with UniFi’s Cloud servers.
- Restart your Talk phones and attempt to pair them again.
- If your Talk phone is displaying the status “Needs attention,” the phone has likely failed to register with the PBX provider. This may be due to firewall settings on your network or an incorrectly entered PBX configuration.
Account Configuration Status
At the top of your Talk phone’s screen, you will see the phone’s registration status:
- Green: Phone has active connection with the PBX provider.
- Red: Phone has failed to establish a connection with the PBX provider.
Standalone Setup (Without Official UniFi Hosting)
Standalone mode allows each Talk phone to be manually configured, directly from its touch screen, without any remote management or monitoring capabilities. To do this:
- Reset your UniFi Talk phones to factory settings if they have been previously configured.
- Power the phone on and choose Third-Party PBX setup on the phone’s screen.
- Tap Manual Configuration.
- Enter the configuration details obtained from your PBX provider.
