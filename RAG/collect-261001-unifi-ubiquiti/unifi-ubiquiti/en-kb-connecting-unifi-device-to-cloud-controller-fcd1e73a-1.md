---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/en-kb-connecting-unifi-device-to-cloud-controller-fcd1e73a-1
title: "en-kb-connecting-unifi-device-to-cloud-controller-fcd1e73a"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/en-kb-connecting-unifi-device-to-cloud-controller-fcd1e73a.md
source_anchor: ""
source_lines: [1, 56]
sha256: 366d6ee9f3148937e6e408f099f62d2d4872ffbb379cfbbb3553509c4def5a1e
---

# en-kb-connecting-unifi-device-to-cloud-controller-fcd1e73a

Connect UniFi Device to Cloud Controller via Sophos Firewall
If a UniFi Access Point or switch is behind a Sophos Firewall and the UniFi Network Application is at another site, the device needs a reachable Inform path to the controller. The key requirements are a narrow firewall rule for TCP 8080 and an Inform URL supplied through DNS, SSH, or DHCP Option 43.
This guide primarily applies to Official UniFi Hosting, CloudKeys, and self-hosted UniFi Network Servers outside the device VLAN. UniFi Cloud Gateways can sometimes discover devices directly within their own routed environment. For help choosing an operating model, see UniFi Controller – Managing Access Points and Switches.
Customers with a Sophos Firewall subscription can use the Avanet Cloud Controller for up to five UniFi devices free of charge.
Check the requirements and choose the right method
Before adoption, the following must be in place:
- The UniFi device receives an IP address, a gateway, and a working DNS server through DHCP.
- The controller FQDN or controller IP and the correct Inform URL are known. With Official UniFi Hosting, use Copy Inform URL.
- The device network has an allowed path to the UniFi Network Application over TCP 8080 .
- For the SSH method, TCP 22 is required only from the admin client to the UniFi device.
- Existing DNS, DHCP, and firewall values and the current Inform URL have been recorded. Create dedicated objects for the test instead of overwriting existing entries.
The appropriate method depends on the environment:
- DNS entry unifi : useful when several new devices on the same network should automatically find the correct controller.
- SSH with set-inform : a targeted manual method for a single device with a known IP address.
- DHCP Option 43: suitable when the DHCP scope should distribute the controller address centrally to several devices.
- UniFi Mobile App: an alternative when the smartphone is on the same device VLAN.
- Zero-Touch Provisioning: for supported devices and operating models; current compatibility is listed in the Ubiquiti ZTP documentation.
If Sophos Firewall and UniFi switches provide VLANs together, first check the mapping described in Configure a VLAN on Sophos Firewall and a UniFi Switch.
Create the Sophos Firewall rule for adoption
For an external controller, the rule from the UniFi device network to the controller should be as narrow as possible. A specific example:
- Under Hosts and services > FQDN host , create an object such asUniFi_Controller with the FQDNcontroller.example.com . The background to DNS resolution is explained in Use FQDN hosts correctly on Sophos Firewall.
- If no suitable service object exists, create UniFi-Inform underHosts and services > Services with ProtocolTCP , Source port1:65535 , and Destination port8080 .
- Under Rules and policies > Firewall rules , create a rule namedUniFi devices to controller with Action set toAccept .
- For Source zones, select the actual zone of the device VLAN and, for Source networks and devices, select an object such as UniFi_Devices for192.168.10.0/24 . A host or device-group object is narrower than the entire/24 example network and is preferable once the devices have fixed addresses.
- For a cloud controller, select WAN as the Destination zones andUniFi_Controller as the Destination networks. For a local controller in another VLAN, select its internal zone.
- Under Services, select UniFi-Inform , enable Log firewall traffic, and place the rule above general block rules. Check its position afterwards because SFOS evaluates rules from top to bottom.
- If the UniFi environment uses STUN for adoption and device communication, create a second service object named UniFi-STUN with ProtocolUDP , Source port1:65535 , and Destination port3478 , and add it to the same narrow rule.
TCP 8080 is the key device and Inform path: Ubiquiti lists it as ingress from the application host’s perspective, so the new session on Sophos Firewall runs from the device to the controller. Ubiquiti lists UDP 3478 for STUN in both directions; the stateful firewall already allows replies to a session initiated by the device. The Both designation is not a reason to publish the device network from the WAN. A controller-initiated reverse path requires a separately planned routing or VPN design that is explicitly required by the specific operating model.
TCP 443, or TCP 8443 for a self-hosted UniFi Network Server, is used for administrator access and should not be added indiscriminately to the rule from the device VLAN to the controller. UDP 10001 is used for local L2 discovery and does not need to be opened to the WAN for cross-site L3 adoption.
The device must also be able to reach DNS and NTP, but these connections belong to the designated DNS and time servers and not automatically in the controller rule. If a self-hosted controller is behind a remote gateway, TCP 8080 must be forwarded to it or made reachable through a VPN. Double NAT without a reachable path does not work.
Adopt through the DNS name unifi
UniFi Network Devices try to resolve the name unifi for L3 adoption. If Sophos Firewall serves as the DNS server on the device network, create the entry as follows:
- Open Network > DNS .
- Scroll to DNS host entry and select Add.
- Enter unifi as the Host/Domain name.
- Set Entry type to IP address , enter the reachable controller IP as the IP address, and save. Leave Publish on WAN turned off; the internal short name should not be answered on the WAN.
- Under Administration > Device access , check that DNS is allowed only for the required device zone. Configure Device Access on Sophos Firewall explains how to restrict this securely.
Configure and test DNS host entries on Sophos Firewall explains how to build a static entry with TTL, a client test, and a clear boundary to DNS request routes.
A test client on the same VLAN should then return the expected controller IP. In this example, 192.168.10.1 is the Sophos Firewall:
nslookup unifi 192.168.10.1
Then restart the UniFi device or renew its DHCP lease. It should appear as ready for adoption in UniFi Network. If it remains invisible, use the Sophos Firewall Log Viewer to check whether the connection from the device network to the controller over TCP 8080 is allowed.
If an internal DNS server rather than Sophos Firewall supplies the responses, the unifi entry must be created there. An entry on the firewall does not help if the devices use a different DNS server.
Adopt through SSH and set-inform
SSH is suitable when a single device needs to be assigned to a specific controller. The device must already have an IP address and be reachable from the admin client on the internal management or device network. Do not publish TCP 22 from the internet to the device.
Current UniFi Network Devices that have not yet been adopted use ui as the default user name and password. Older devices may still use ubnt/ubnt. If the device was already adopted, only the credentials stored under Device SSH Authentication in UniFi Network apply.
With the example address 192.168.10.20, the login command is:
ssh ui@192.168.10.20
After login, set the Inform URL:
set-inform http://controller.example.com:8080/inform
The device should then appear in UniFi Network and can be adopted. If adoption does not complete or the device subsequently shows Disconnected, run the same set-inform command again. Ubiquiti documents that this repetition may be necessary; it is not a reason to broaden the firewall rule.
When using the DNS method with the short name unifi, the command can also be:
set-inform http://unifi:8080/inform
On classic UniFi Access Points and switches, the following command often shows the current Inform URL and connection status:
info
This diagnostic command depends on the model and firmware. First use help to check whether the device offers it. If it does not, the device status in UniFi Network and the firewall logs are more reliable checks.
Use DHCP Option 43 as an alternative
