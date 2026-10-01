---
id: collect-261001-general-networking/general-networking/hc-en-us-articles-204909754-remote-adoption-layer-3-aeddad25
title: "hc-en-us-articles-204909754-remote-adoption-layer-3-aeddad25"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/hc-en-us-articles-204909754-remote-adoption-layer-3-aeddad25.md
source_anchor: ""
source_lines: [1, 50]
sha256: 5da7ccde5c58bd9aa3cc5dbf03fe81b9fea62f1d818f6b5f135b4f38309486ad
---

# hc-en-us-articles-204909754-remote-adoption-layer-3-aeddad25

Remote Adoption (Layer 3)
Layer 3 adoption is the process of adopting a UniFi device to a remote or cloud-hosted UniFi Network Application. This is only recommended for advanced users, or those adopting devices when using Official UniFi Hosting.
Note: Refer to Device Adoption if you are unsure whether you require remote adoption.
Requirements
The UniFi device and the UniFi Network Application must have unrestricted connectivity over TCP Port 8080.
Note: a UniFi device will cycle through all of these inform options and will try to be adopted by each, including the ZTP code. Whichever is the first to resolve and respond will adopt the device.
Methods
Mobile App (Existing Networks)
The UniFi Mobile App (iOS / Android) can streamline adoption if your network is already broadcasting WiFi on the same VLAN as the unadopted UniFi device.
- Ensure the device has a solid white LED indicating it is in a factory default state.
- Connect your mobile device to WiFi associated with the same VLAN as the unadopted device. If this isn't possible, refer to the SSH method below.
- Open your UniFi Network Mobile App and connect to site you want to adopt your device.
- Your device should appear for adoption in the device list.
SSH (New Networks)
If WiFi is not available, follow these steps:
- Ensure your device is in a factory default state.
- Connect a laptop directly to the network using a wired connection.
- SSH into the device.
- 
Issue the following command: set-inform http://ip-of-host:8080/inform
- Visit the UniFi Site Manager or UniFi Mobile App and click on the device appearing as "ready to adopt."
Note: If you are using Official UniFi Hosting, you can visit your UniFi Network dashboard and click Copy Inform URL to paste into the command above.
Other (Advanced)
DHCP Option 43
This option leverages your DHCP server to inform your UniFi device of the location of the machine running the Network Application remotely. Those with a UniFi Gateway can easily accomplish this by entering the IP address of the remote Network Application in Option 43 Application Host Address field located under Settings > Networks > [Select VLAN] > Option 43.
For those using a third-party gateway or DHCP server, we recommend consulting your manufacturer’s documentation to learn more. For most third party firewalls, DHCP Option 43 is encoded in hex, and there are two methods of specifying the machine running UniFi Network / UniFi OS:
- IP Address:
  - Start with 0104 , then add your static IP (encoded in hex).
  - Example: 192.168.3.10 would be entered as 0104 c0a8030a (without spaces).
- Start with 
- Fully Qualified Domain Name (FQDN):
  - Start with 02 then the length of the FQDN in bytes (hex), then the inform URL (converted to hex).
  - Example: http://192.168.3.10:8080/inform, which is 31 bytes (1f) long, would be entered as 02 1f 687474703a2f2f3139322e3136382e332e31303a383038302f696e666f726d
- Start with 
- For additional redundancy, both the IP and FQDN could be specified in the same DHCP option:
  - Example: 01 04 c0a8030a 02 1f 687474703a2f2f3139322e3136382e332e31303a383038302f696e666f726d (without spaces)
- Example: 
DNS
You'll need to configure your DNS server to resolve unifi to your remote UniFi Network installation.
There are two methods of specifying the machine running UniFi Network / UniFi OS:
- IP Address: http://ip-address:8080/inform
- Fully Qualified Domain Name (FQDN): http://FQDN:8080/inform
Troubleshooting
If you are experiencing difficulty, then your UniFi device and UniFi Network application likely cannot communicate over TCP Port 8080. This must be addressed differently depending on the location of your UniFi Network Application:
- If the UniFi application and unadopted device are on separate VLANs behind the same local gateway, ensure inter-VLAN routing is enabled and TCP Port 8080 is not blocked on the firewall.
- If you are using Official UniFi Hosting or your UniFi Console has a public IP address, ensure your local firewall/gateway allows your unadopted UniFi device to communicate with it over TCP Port 8080.
- If the UniFi Network Application (often running on a CloudKey) is running behind a remote gateway and the gateway has a public IP address, then you must Port Forward TCP 8080 from the gateway to the IP address associated with the CloudKey.
- If the UniFi Network Application (often running on a CloudKey) is running behind a remote gateway, but the gateway has a private IP address (it is double NAT'd), this will not work. Contact your ISP for assistance.
Zero-Touch Provisioning (ZTP)
ZTP is a seamless and efficient method for adopting Layer 3 UniFi devices. This is currently supported by U7 Pro Max APs, with more devices coming soon. For detailed instructions, visit Zero-Touch Provisioning.
