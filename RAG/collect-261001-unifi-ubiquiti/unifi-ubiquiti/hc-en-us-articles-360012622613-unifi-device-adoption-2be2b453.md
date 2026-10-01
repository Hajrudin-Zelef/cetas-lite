---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-360012622613-unifi-device-adoption-2be2b453
title: "hc-en-us-articles-360012622613-unifi-device-adoption-2be2b453"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-360012622613-unifi-device-adoption-2be2b453.md
source_anchor: ""
source_lines: [1, 59]
sha256: cc88df41f410f29f9e9f2788887cfb135b5ffbe61138f2db504a3e628c34fa20
---

# hc-en-us-articles-360012622613-unifi-device-adoption-2be2b453

UniFi - Device Adoption
Adoption is the process of connecting a UniFi device—such as an Access Point (AP), camera, switch, or sensor—to the UniFi Gateway and the UniFi application that will manage it. Most UniFi devices adopt automatically with no manual configuration. This article outlines the standard plug-and-play method and then covers additional adoption methods for advanced environments.
The Easiest Way to Adopt Devices
If you're using a UniFi Cloud Gateway, adoption is simple:
- Install a UniFi device.
- Download the UniFi Mobile App for iOS or Android.
- Power on your device and connect it to the same network.
- Open the app and tap to adopt the device.
No port forwarding, firewalls, or manual commands are needed. Devices are automatically discovered by UniFi Network. We recommend this method for most users.
Other Adoption Scenarios
In some cases, such as custom VLANs, remote setups, or self-hosted controllers, you may need to follow additional steps for successful adoption.
1. Same-VLAN Adoption (CloudKey or Self-Hosted)
If the device and UniFi Network application (e.g., CloudKey or self-hosted) are on the same VLAN:
- Ensure your UniFi Host (e.g., CloudKey or self-hosted) is set up and running UniFi Network.
- Confirm the device is powered on.
- The device should appear in the UniFi app or web interface as Pending Adoption
- If not, ensure TCP: 8080 and UDP: 10001 is open between on all gateways, firewalls, and antivirus software, between the UniFi Host and the UniFi Devices.
- As a last resort, factory reset the device and try again.
2. Different VLAN or Subnet (Layer 3 Adoption)
If the device is on a different VLAN or subnet than your UniFi Network application:
- Follow the Layer 3 Adoption guide.
- Ensure routing and firewall rules allow TCP port 8080 between the device and UniFi Network.
- As a last resort, SSH into the device and run: set-inform http://<gateway-ip>:8080/inform .
 
  - You may need to repeat the command after the device appears in UniFi.
3. Remote Adoption (Different Site)
If adopting a UniFi device from a different physical location (i.e., over the internet), follow the same Layer 3 Adoption steps. Then:
- Ensure that UniFi Network is reachable over TCP port 8080 from the device. What this means depends on your deployment:
 
  - Scenario 1: If the UniFi host has its own public IP address, make sure no upstream firewall is blocking TCP port 8080.
  - Scenario 2: If the UniFi host does not have a public IP address (e.g., it’s behind a Gateway), you must port forward TCP port 8080 on the Gateway to the internal IP of the UniFi host.
  - Scenario 3: If using a VPN, the device is already on the same logical network. Just ensure that no firewall is blocking TCP 8080 traffic internally.
- As a last resort, SSH into the device and run: set-inform http://<public-ip-or-hostname>:8080/inform .
Alternatively, consider using Zero-Touch Provisioning (ZTP) for remote deployments that don’t require manual inform commands.
Applies to UniFi APs connecting wirelessly to a wired uplink AP. Not available in Standalone Mode.
- Ensure your Gateway is running UniFi Network.
- Confirm the device is powered on.
- Enable Wireless Meshing in Settings > WiFi > [Wireless Meshing].
- On the uplink AP, go to Devices > Select AP > Settings, and enable Mesh Connect.
- Place the new AP near the uplink AP to ensure a stable signal.
- See Wireless Uplink Requirements for best practices.
- As a last resort, factory reset the device and try again.
For wired cameras and devices in restricted setups:
- Factory reset the device if needed.
- Confirm the device is powered on.
- Ensure any firewall or antivirus software between the Protect Device and the UniFi Protect Host allow are not blocking network connectivity. Device adoption requires TCP: 7444, 7550, 7442 and UDP: 10001.
- Ensure VLANs and switch ports are properly configured so that the camera receives and IP address via DHCP. See VLANs and Switch Ports to learn more and VLAN Troubleshooting if you encounter an issue.
- Ensure your devices are on the same network (VLAN) as the UniFi Protect host (e.g., CloudKey or NVR).
  - If you must have cameras on a different VLAN, you must manually specify the UniFi Host on each camera: Navigate to the camera's IP address in your browser > sign in using default credentials (ui/ui) > specify the IP address of the UniFi Protect Host.
The default adoption method is via the UniFi Protect App on iOS or Android. Open the Protect app and walk near a Bluetooth enabled Protect host device and tap the pop-up when it appears on your screen.
If your device does not automatically appear, ensure:
- You have enabled Bluetooth permissions on your mobile device.
- The VLAN associated with your WiFi SSID is the same as the Protect host (e.g., CloudKey or NVR).
- Your mobile device is in close proximity to the Protect device you are adopting (within Bluetooth range).
- Your UniFi Protect device is in close proximity to the AP it is connecting to.
- There are no firewall rules blocking local network connectivity from the device you are adopting and the Protect host.
To adopt a UniFi wireless sensor like UP Sense:
- Make sure your UP Sense is near a UniFi SuperLink gateway, or a Bluetooth-enabled U6 Access Point.
- As a last resort, factory reset the device and try again.
