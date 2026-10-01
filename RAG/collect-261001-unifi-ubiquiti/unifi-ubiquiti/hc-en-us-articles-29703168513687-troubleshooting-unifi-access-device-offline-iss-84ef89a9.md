---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-29703168513687-troubleshooting-unifi-access-device-offline-iss-84ef89a9
title: "hc-en-us-articles-29703168513687-troubleshooting-unifi-access-device-offline-iss-84ef89a9"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-29703168513687-troubleshooting-unifi-access-device-offline-iss-84ef89a9.md
source_anchor: ""
source_lines: [1, 77]
sha256: 34ecf1e7575aa8c9ee4b30f3ce310eda3269c7008b8d816546a0e47581bdb2b5
---

# hc-en-us-articles-29703168513687-troubleshooting-unifi-access-device-offline-iss-84ef89a9

Troubleshooting UniFi Access Device Offline Issue
If your UniFi Access devices (e.g., Access Control Hub or Access Reader) are constantly or intermittently offline, follow the steps below sequentially to diagnose and resolve the issue.
| Update your UniFi OS, Access application, and Access devices to the latest versions for the newest features and optimal performance. | 
Step 1: Ensure the UniFi Console and Access Devices Are on the Same VLAN
- Navigate to Access application > Devices and check the Access device's IP.
- Navigate to Network > UniFi Devices to confirm that the console and Access device are on the same VLAN.
- If they are on different VLANs, navigate to Access application > Settings > General > Network > All Networks.
  - Network Video Recorder Pro (UNVR-Pro) and Network Video Recorder (UNVR) do not support the All Networks option, but you can use their GbE RJ45 port and 10G SFP+ port to configure two separate VLANs if needed.
  - 
CloudKey+ (UCK-G2-PLUS) does not support the All Networks option and only allows a single VLAN.
    - If you cannot place the Access devices and the CloudKey+ on the same VLAN, consider using a different UniFi console that supports Access and the All Networks option.
Step 2: Verify Power Supply and PoE Connection
Check PoE requirements for Access Control Hubs:
- PoE+ for Access Ultra.
- PoE++ for Door Hub Mini, Door Hub, Elevator Hub, and Gate Hub.
- PoE++ adapter (60W) is recommended for stable power.
- If using a switch, ensure it has a sufficient PoE budget.
  - Check it in Network application > UniFi Devices > Switches > select your switch > Overview > Power Utilization.
  - Signs of insufficient PoE: Devices randomly go offline or restart unexpectedly.
Check PoE connection:
- If an Access Reader is offline, try a different PoE port on the Access Control Hub.
- If an Access Control Hub is offline, try a different PoE port on the UniFi Console.
- Try replacing the PoE cables to rule out faulty connections.
- Try replacing the PoE++ adapter.
Note: Enterprise Access Hub requires AC power; Retrofit Hub requires DC power.
Step 3: Update Application and Device Version
Update the Access application and devices to the latest versions:
- Access application: Navigate to Access application > Settings > Control Plane > Updates > Access.
- Access device: Navigate to Access application > Devices > Click to Update.
Step 4: Check Device Uptime
- Navigate to Access application > Devices > select the offline device > Overview > Uptime.
Step 5: Ping the Device to Verify Connectivity
- Enable SSH and note the SSH password in Access application > Settings > Control Plane > Console > SSH.
- 
Run the following commands on a computer connected to the UniFi Console network. ssh root@UNIFI_GATEWAY_LAN_IP 
 ping UA_Device_IPReplace UNIFI_GATEWAY_LAN_IP with your console's LAN IP. 
Replace UA_Device_IP with the actual IP of the device experiencing the issue.
- Take a screenshot of the result.
Step 6: Find the Exact Timestamp of the Offline Event
- Go to Access application > System Log > Critical and look for any logged offline events.
- Take a screenshot that clearly shows the exact timestamp and device name of the most recent entries.
Step 7: Download Device Support Files
Ensure to download the support files when the devices are online.
- Access Control Hub support file: Navigate to Access application > Devices > select your hub > Settings > Manage > Download Support File.
- Access Reader support file: Navigate to Access application > Devices > select your hub to which the reader is connected > Settings > Manage > Download Support File.
  - If an Access Reader is offline, try bringing it back online by navigating to Access application > Devices > select the hub to which the reader is connected > Overview > Paired Devices > Click to Power Cycle.
- UniFi Console support file: Navigate to Access application > Settings > Control Plane > Console > Support File > Full > Download.
Step 8: Factory Reset and Adopt the Device
Factory resetting the device will erase all existing log files. Before proceeding, ensure that you have downloaded the support files.
- Press and hold the reset button on the offline device for 10 seconds.
  - The device LED should turn steady white.
  - For Reader Pro, the screen should display: "This Reader Pro has not been set up."
- Adopt the device in Access application > Devices > Click to Adopt.
Step 9: Contact Technical Support and Provide Relevant Information
If the issue persists, contact UniFi Access Technical Support via UI Account (https://account.ui.com/) > Support. When submitting your request, be sure to include:
- The results from Steps 1 to 7, especially the exact timestamp and all relevant support files.
- Confirmation of whether you have attempted Step 8, and whether the issue persists afterward.
FAQs
How do I bring offline cameras and readers connected to Access Control Hubs back online?
You can bring offline cameras and readers back online by power cycling them through the connected Access Control Hub.
(Requires Access 3.3.10 or later) Navigate to Access application > Devices > select the hub to which the offline device is connected > Overview > Paired Devices > Click to Power Cycle.
How do I ensure my support file contains complete logs?
Support file logs are usually automatically cleared after a few days to save storage and system resources. To avoid losing critical log details, always download the support file immediately after the issue occurs.
- Access device support file: Navigate to Access application > Devices > select your device > Settings > Manage > Download Support File.
- UniFi Console support file: Navigate to Access application > Settings > Control Plane > Console > Support File > Full > Download.
In addition, certain actions will immediately clear all logs from the support file. Ensure to download it before performing any of the following:
- Restarting a device
  - UniFi Access device: Navigate to Access application > Devices > select the device > Settings > Restart.
  - UniFi Console: Navigate to Access application > Settings > Control Plane > Console > Controls > Restart.
- Removing an Access device:
  - Navigate to Access application > Devices > select the device > Settings > Remove.
- Factory resetting a device
  - UniFi Access device: Press and hold the reset button on the device for 10 seconds.
  - UniFi Console:
    - Press and hold the reset button on the device for 10 seconds.
    - 
Navigate to Access application > Settings > Control Plane > Console > Controls > Factory Reset.
