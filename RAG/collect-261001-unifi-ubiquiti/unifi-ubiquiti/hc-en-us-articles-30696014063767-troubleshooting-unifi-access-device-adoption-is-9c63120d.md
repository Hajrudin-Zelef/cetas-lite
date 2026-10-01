---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-30696014063767-troubleshooting-unifi-access-device-adoption-is-9c63120d
title: "hc-en-us-articles-30696014063767-troubleshooting-unifi-access-device-adoption-is-9c63120d"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-30696014063767-troubleshooting-unifi-access-device-adoption-is-9c63120d.md
source_anchor: ""
source_lines: [1, 64]
sha256: c1ca18625f93ede9268f1cf7c8be42b5ac73d1867e31ee2842aa2f047b30552b
---

# hc-en-us-articles-30696014063767-troubleshooting-unifi-access-device-adoption-is-9c63120d

Troubleshooting UniFi Access Device Adoption Issue
If your UniFi Access device (e.g., Access Control Hub or Access Reader) cannot be adopted, follow the steps below in order to diagnose and resolve the issue based on whether the device is discovered in the Access application.
Issue 1: Access Device Is Not Discovered in the Access Application
Step 1: Ensure the Access Application Is Up to Date
Navigate to Access application > Settings > Control Plane > Updates.
- Update the UniFi OS if it's not already on the latest version.
- Update the Access application if it's not already on the latest version.
Step 2: Ensure No Firewall Rules Are Blocking ICMP Traffic
Navigate to Network application > Settings > Policy Engine > Traffic & Firewall Rules.
Step 3: Ensure the UniFi Console and Access Devices Are on the Same VLAN
- Navigate to Access application > Devices and check the Access device's IP.
- Navigate to Network > UniFi Devices to confirm that the console and Access device are on the same VLAN.
- If they are on different VLANs, navigate to Access application > Settings > General > Network > All Networks.
  - Network Video Recorder Pro (UNVR-Pro) and Network Video Recorder (UNVR) do not support the All Networks option, but you can use their GbE RJ45 port and 10G SFP+ port to configure two separate VLANs if needed.
  - 
CloudKey+ (UCK-G2-PLUS) does not support the All Networks option and only allows a single VLAN.
    - If you cannot place the Access devices and the CloudKey+ on the same VLAN, consider using a different UniFi console that supports Access and the All Networks option.
Step 4: Connect all Access Devices Directly to the Console or Switch
Connect your Access Control Hubs and Access Readers directly to your console or switch using short, pre-made patch cables that have been confirmed to work.
Step 5: Ping the Access Device to Verify Connectivity
- Enable SSH and note the SSH password in Site Manager > UniFi Console > Settings > Control Plane > Console > SSH.
- 
Run the following commands on a computer connected to the UniFi Console network. ssh root@UNIFI_GATEWAY_LAN_IP
ping UA_Device_IPReplace UNIFI_GATEWAY_LAN_IP with your console's LAN IP. 
Replace UA_Device_IP with the actual IP of the device experiencing the issue.
- Take a screenshot of the result.
Step 6: Download Device Support Files
- (If available) Access Control Hub support file: Navigate to Access application > Devices > select your hub > Settings > Manage > Download Support File.
- UniFi OS support file: Navigate to Access application > Settings > Control Plane > Console > Support File > Full > Download.
Step 7: Factory Reset and Adopt the Device
Factory resetting the device will erase all existing log files. Before proceeding, ensure that you have downloaded the support files.
- Press and hold the reset button on the offline device for 10 seconds.
  - The device LED should turn steady white.
  - For Reader Pro, the screen should display: "This Reader Pro has not been set up."
- Adopt the device in Access application > Devices > Click to Adopt.
Step 8: Contact Technical Support and Provide Relevant Information
If the issue persists, contact UniFi Access Technical Support via UI Account (https://account.ui.com/) > Support. When submitting your request, be sure to include:
- The results from Steps 1 to 6.
- Confirmation of whether you have attempted Step 7, and whether the issue persists afterward.
Issue 2: Access Device Is Discovered in the Access Application, But Adoption Fails
Step 1: Verify PoE Connection and Port
- If an Access Reader is not discovered, try a different PoE port on the Access Control Hub or connect it directly to the console or switch.
- If an Access Control Hub is not discovered, try a different PoE port on the console or switch.
- Try replacing the PoE cable with a short pre-made patch cable to rule out faulty connections.
- Try replacing the PoE++ adapter.
Step 2: Ping the Access Device to Verify Connectivity
- Enable SSH and note the SSH password in Site Manager > UniFi Console > Settings > Control Plane > Console > SSH.
- 
Run the following commands on a computer connected to the UniFi Console network. ssh root@UNIFI_GATEWAY_LAN_IP
ping UA_Device_IPReplace UNIFI_GATEWAY_LAN_IP with your console's LAN IP. 
Replace UA_Device_IP with the actual IP of the device experiencing the issue.
- Take a screenshot of the result.
Step 3: Download Device Support File
UniFi OS support file: Navigate to Access application > Settings > Control Plane > Console > Support File > Full > Download.
Step 4: Factory Reset and Adopt the Device
Factory resetting the device will erase all existing log files. Before proceeding, ensure that you have downloaded the support file.
- Press and hold the reset button on the offline device for 10 seconds.
  - The device LED should turn steady white.
  - For Reader Pro, the screen should display: "This Reader Pro has not been set up."
- Adopt the device in Access application > Devices > Click to Adopt.
Step 5: Contact Technical Support and Provide Relevant Information
If the issue persists, contact UniFi Access Technical Support via UI Account (https://account.ui.com/) > Support. When submitting your request, be sure to include:
- The results from Steps 1 to 3.
- Confirmation of whether you have attempted Step 4, and whether the issue persists afterward.
