---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-unifi-adoption-ap-failed-e232b048
title: "blog-unifi-adoption-ap-failed-e232b048"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-unifi-adoption-ap-failed-e232b048.md
source_anchor: ""
source_lines: [1, 41]
sha256: 0e6cdcc41526003fe8e67ef96c76621fd06d8721ebdc079a79fa36c7cba4cb76
---

# blog-unifi-adoption-ap-failed-e232b048

UniFi adoption AP failed: How to fix it
If you are going to deploy a UniFi network, you'll need to ingreate some access points (APs) into your UniFi Controller. This process is quite easy, but sometimes the integration/adoption does not go as planned. When the adoption of your AP fails, it's annoying to figure out why it happend.
In this guide, I'll quickly demonstrate the easiest way to troubleshoot the issues and to get going again. I'll outline some reasons why an AP adoption might fail provide some hands-on steps to resolve it.
Reasons for AP adoption failures
AP adoption is the process by which new UniFi wireless access points are recognized and managed by the UniFi Controller. The process involves the AP connecting to the network, the controller recognizing the AP, and then the controller applying configurations to the AP. Simple as that.
Adoption failures can be caused by quite a few factors:
- Network issues: Incorrect network settings, such as DHCP errors or IP conflicts.
- Communication barriers: Firewalls or network security settings that block communication ports.
- Firmware issues: Mismatched firmware versions between the AP and the UniFi Controller (this happens a lot).
- Distance and connectivity: Physical connectivity issues or being too far from other devices.
Troubleshooting Guide
If you encounter an issue with adopting an AP, follow these steps to troubleshoot and resolve the problem:
Step 1: Check network connectivity
- Verify cabling and connections: Make sure that the AP is properly connected to your network via Ethernet and is receiving power (PoE or adapter based).
- Confirm DHCP functionality: Make sure that your DHCP server is functioning correctly and that the AP is receiving an IP address. A static IP may also be set for troubleshooting purposes.
Step 2: Verify controller accessibility
- Ping the AP: From a device on the same network, try pinging the AP’s IP address to test connectivity.
- Check controller pettings: Make sure the UniFi Controller is set to accept and manage new APs, and isn’t restricted by IP or MAC filtering.
Step 3: Inspect firewall and security settings
- Firewall rules: Check if any firewall settings on your network or computer might be blocking communication between the AP and the UniFi Controller.
- Ports: Make sure that necessary ports are open. UniFi APs generally require ports 8080, 8443, 8880, and 6789 to be open for proper communication.
Step 4: Update Firmware
Make sure that both the UniFi Controller and the AP are running compatible firmware versions. If not, here is how to update your UniFi Controller to the right version.
Step 5: Adopt Via SSH
If standard adoption methods fail:
- SSH into AP: Use an SSH client to connect directly to the AP using its IP address. Default credentials are typically "ubnt" for both username and password, unless changed.
- Set Inform URL: Once SSH'd in, execute the command:
Replaceset-inform http://[controller_ip]:8080/inform
[controller_ip] with the IP address of your UniFi Controller. This forces the AP to attempt adoption with the specified controller (for more info about how to set inform in UniFi, check this guide.
Step 6: Factory Reset the AP
If all else fails:
- Reset the AP: Perform a factory reset by holding the reset button on the AP for about 10 seconds.
- Re-adopt: After resetting, try adopting the AP again through the UniFi Controller.
Final thoughts
UniFi AP adoption issues are annoying. If you run into them, follow the steps above to systematically identify and address the causes. It's important to make sure that your network expansions and updates proceed smoothly and without delays.
If you're running larger networks and want some assitence to make them more efficient and secure, consider a managed-hosting provider like UniHosted. We provide managed UniFi Controller to improve your network efficiency. Especially, if you are an MSP or IT-service business looking for reliable, secure and scalable UniFi network deployments, this is something for you.
Related guides
Keep reading
- Understanding and managing DHCP: A step by step guideStep-by-step guide on configuring DHCP settings Read guide
- How to configure firewall rules for UniFi Controller: A step by step guideDetailed guide on how to configure firewall rules for UniFi Controllers Read guide
- How to update your UniFi Controller safelyStep-by-step guide to update UniFi Controller — covers both self-hosted and cloud-hosted methods with backup tips and best practices included. Read guide
