---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625-2
title: "docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "license", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625.md
source_anchor: ""
source_lines: [69, 130]
sha256: cd0d2c5f7cf4cab50183e5d4e2257706f6c0eac144c956a4c7db04805048f932
---

# docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625

- Block bogon networks: This option blocks traffic from IP addresses that are not yet assigned by IANA. This is also a security measure, as it prevents malicious traffic from reaching your network.
- IPv4 Configuration Type: This setting specifies how the IPv4 address for the interface will be configured. You can choose between static or DHCP.
- IPv6 Configuration Type: This setting specifies how the IPv6 address for the interface will be configured. You can choose between static or DHCP.
- MAC address: This is the MAC address of the interface. You can leave this blank if you are not sure what it is.
- Promiscuous mode: This setting allows the interface to receive all packets, even those that are not addressed to it. This can be useful for troubleshooting purposes, but it should be disabled unless necessary.
- MTU: This is the maximum transmission unit for the interface. This is the size of the largest packet that can be sent over the interface.
- MSS: This is the maximum segment size for TCP connections. This is the size of the largest TCP segment that can be sent over the interface.
- Dynamic gateway policy: This setting specifies whether the interface should use a dynamic gateway. A dynamic gateway is a gateway that is automatically configured by the router.
- Static IPv4 configuration: This section allows you to configure a static IPv4 address for the interface.
- IPv4 Upstream Gateway: This is the IP address of the gateway for the interface. The gateway is the device that routes traffic between your network and the rest of the internet.
If you want clients on this VLAN to get IPs automatically, go to Services → DHCPv4 → [your VLAN interface], check Enable DHCP server on…, set a pool (e.g., 192.168.10.100–192.168.10.200), and save/apply.
Figure 6. VLAN Interface Configuration
Figure 7. VLAN Interface Configuration-2
7. Set Up Firewall Rules for the VLAN Network
By default, OPNsense blocks all traffic on newly created interfaces, including VLANs.
To enable communication, you need to configure firewall rules for your VLAN interface. Firewall rules determine which traffic is allowed or denied based on parameters such as source, destination, port, and protocol. Properly configuring these rules ensures secure and controlled access between your VLAN and other networks. To define and apply specific firewall rules for your VLAN, follow the steps below.
- 
Navigate to Firewall -> Rules .
- 
Select the tab corresponding to your VLAN interface.
- 
Click Add to create a new rule.Figure 8. Creating a Firewall Rule
- 
Configure the rule to allow required traffic (e.g., TCP/UDP, HTTP/HTTPS, ICMP). Figure 9. Edit Firewall Rule Figure 10. Edit Firewall Rule-2
- 
Click Save and Apply Changes to activate the rule. Figure 11. Applying Firewall Rule Changes
You can create more granular rules later to improve security.
8. Configure the Switch Port for the VLAN (if applicable)
If your network uses a managed switch, make sure the port connecting to OPNsense is properly configured to handle VLAN tags with the following steps.
- Log in to your managed switch's interface.
- Locate the port connected to OPNsense.
- Set the port as Tagged (Trunk) for the VLAN ID you configured.
- Optionally set other ports as Untagged (Access ) for end devices that will connect to the VLAN.
This ensures that VLAN-tagged traffic is properly recognized and routed.
9. Test the VLAN Connectivity with a Device
Testing confirms whether the VLAN works as expected before deploying it widely.
- Connect a device (e.g., laptop or VM) to a switch port assigned to the VLAN.
- Ensure it receives an IP address from the correct VLAN subnet.
- Test connectivity (e.g., ping the gateway or browse the internet).
- Verify that firewall rules are functioning as intended.
This helps identify and fix any misconfigurations early.
10. Apply and Save All Configuration Changes
Finalizing your changes ensures they are active and persistent across reboots with the following steps.
- Ensure that you click Apply Changes on any page where changes were made.
- Optionally, back up your configuration via System > Configuration > Backups .
Your VLAN is now fully operational and integrated into the network.
Configuring VLANs on OPNsense empowers network administrators with improved performance, security, and scalability. By effectively segregating and managing traffic, businesses can streamline their network operations and protect sensitive data.
If you're familiar with configuring VLANs in OPNsense and are now transitioning to Zenarmor, you'll find that Zenarmor offers a powerful feature called Exempted VLANs & Networks. This feature allows you to define specific VLANs and IP/Network addresses that are exempted from Zenarmor processing. Essentially, any traffic associated with these exempted VLANs and addresses bypasses Zenarmor's packet processing entirely, being directly forwarded at the interface level. The key distinction from policy-based whitelisting is that these addresses won't generate any activity reports, ensuring a seamless experience.
A particularly beneficial aspect is that devices within the exempted VLANs and networks are excluded from Zenarmor's license count, meaning they won't contribute to license calculation. This is a handy feature to keep your licensing strategy precise and effective.
However, it's important to note that the Exempted VLANs & Networks feature is available exclusively in premium Zenarmor Editions.
To configure Exempted VLANs & Networks in Zenarmor, follow these straightforward steps:
- Open your OPNsense web UI and navigate to the Zenarmor section.
- From the left-hand sidebar, select the Settings menu.
- Look for the Exempted VLANs & Networks option and click on it.
- Add VLAN ID by clicking on the Exempt VLAN ID button.
For the best and most reliable service and support, sign up with Zenarmor today. Experience seamless VLAN implementation and bolstered network defense to safeguard your organization's critical assets. Don't wait; take the next step toward a robust and secure network environment by signing up with Zenarmor!
What is a VLAN?
A Virtual Local Area Network (VLAN) is a logical subdivision of a physical network, designed to group devices together as if they were on the same local segment, regardless of their physical location. Instead of relying solely on separate physical switches and cabling, a VLAN uses logical segmentation at the switch level, where specific ports or devices are assigned to a unique VLAN ID defined by the IEEE 802.1Q standard. This VLAN ID is embedded into Ethernet frames to distinguish traffic between different virtual networks.
By creating separate broadcast domains, VLANs prevent unnecessary broadcast traffic from spreading across the entire network. This not only reduces network congestion but also enhances security by isolating traffic between departments, user groups, or services. For example, a finance department’s devices can be placed in one VLAN, completely separated from the guest Wi-Fi VLAN, even if they share the same physical hardware.
The use of VLANs provides significant security benefits, as unauthorized devices outside a given VLAN cannot directly communicate with its members. Additionally, VLANs improve network flexibility, administrators can reorganize device groupings without altering physical connections, and allow centralized management by applying policies, firewall rules, or Quality of Service (QoS) settings to specific VLANs.
In an OPNsense environment, VLANs are a powerful tool for traffic separation and policy enforcement. By assigning each VLAN to its own virtual interface, administrators can apply firewall rules, routing policies, and monitoring configurations per segment, achieving both structured organization and enhanced security in multi-department or multi-tenant networks.
How do VLANs Work in an OPNsense Environment?
