---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625-4
title: "docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625.md
source_anchor: ""
source_lines: [174, 243]
sha256: 164137bc97dc04c758c6635fb5da5cb705a5adb63ef97bd49c8286c12a7e5481
---

# docs-network-security-tutorials-how-to-configure-vlan-on-opnsense-2b504625

- If you choose DHCP (for IPv4 or IPv6), the VLAN interface will automatically obtain its IP address and subnet from an upstream DHCP server. This option is often used when the VLAN is part of a network managed by another router or DHCP-capable device.
After saving and applying the changes, you can configure Services → ISC DHCPv4 → [VLAN Interface]  to enable OPNsense’s DHCP server for that VLAN, defining the address pool (e.g., 192.168.50.100 to 192.168.50.200) and other network options.
By assigning a proper IP subnet, each VLAN can function as a self-contained network, enabling precise traffic control, routing, and firewall policy application. This structure is essential for setups like guest networks, VoIP segments, or department-specific LANs.
How do I Set Up DHCP Server Per VLAN?
In OPNsense, you can configure a separate DHCP server for each VLAN so that devices in different network segments receive IP addresses from their own dedicated range. This ensures proper network isolation, accurate IP management, and the ability to apply unique settings such as DNS or gateway addresses per VLAN. Follow these steps to configure DHCP on a VLAN interface.
- 
Ensure the VLAN interface is enabled and has a static IP. Go to Interfaces → [VLAN Name] and verify it is enabled with a static IPv4 address (e.g., 192.168.50.1/24).
- 
Open the DHCP server settings. Navigate to Services → ISC DHCPv4 → [VLAN Interface] .
- 
Enable the DHCP server for the VLAN. Check the box Enable DHCP server on [VLAN Interface] .
- 
Define the address pool. Specify the IP range to be handed out, for example: 
  - From: 192.168.50.100
  - to: 192.168.50.200
- 
Configure optional settings. 
  - DNS servers: Enter internal or public DNS servers (e.g., 8.8.8.8).
  - Gateway: Typically the VLAN interface IP (e.g., 192.168.50.1).
  - Lease time: Adjust based on usage needs.
- 
Save and apply changes. Click Save , thenApply Changes at the top to activate the DHCP server.
If you are creating a guest network VLAN (e.g., VLAN 20), using a dedicated DHCP pool ensures guests cannot obtain IPs from your main LAN. For example, you might assign 192.168.20.1/24 to the guest VLAN interface and set the DHCP pool from 192.168.20.50 to 192.168.20.150. Combined with firewall rules blocking access to internal VLANs, this setup allows guests to have internet connectivity while maintaining security for your private network.
Can I Create Firewall Rules for VLAN Segmentation?
Yes, OPNsense allows you to create firewall rules for each VLAN interface, enabling precise control over traffic between VLANs and to the internet. By defining separate rule sets for each VLAN, you can enforce segmentation, apply security policies, and ensure that devices in one VLAN cannot access resources in another unless explicitly allowed.
In the Firewall → Rules section of OPNsense, each VLAN interface appears as a distinct tab. This means you can create tailored rules for the following puposes.
- Allowing or blocking inter-VLAN traffic (e.g., VLAN 10 to VLAN 20).
- Restricting guest network access so it only reaches the internet and not internal VLANs.
- Applying service-specific permissions, such as allowing only HTTP/HTTPS traffic from certain VLANs.
If you have a VLAN 30 for guests, you can add a default block rule preventing any traffic to private subnets (192.168.0.0/16, 10.0.0.0/8, 172.16.0.0/12). Then, add an allow rule for outbound internet traffic (e.g., destination any on ports 80 and 443). This ensures guests have internet access while your LAN remains secure.
For strong network security, it is important to place explicit block rules at the top of the rule list for sensitive networks, followed by specific allow rules for necessary services. Regularly reviewing these rules ensures there are no unintended access paths between VLANs and that your segmentation strategy remains effective over time.
What are Recommended Firewall Rules for a Guest VLAN?
When setting up a guest VLAN, the primary goal is to provide visitors with secure internet access while protecting internal network resources. The firewall rules should clearly define both traffic restrictions and permitted communication. Here are the recommended firewall rules for a typical guest VLAN.
- Deny Inter-VLAN Traffic: Block all traffic from the guest VLAN to other VLANs or internal networks. This prevents guests from accessing company servers, printers, or shared folders.
- Allow Internet Access: Permit only outbound traffic to the internet, typically HTTP, HTTPS, and DNS.
- Allow Captive Portal Access: If a captive portal is in use, ensure that access to the portal page is allowed so guests can log in or accept terms before internet access is granted.
To implement these rules in OPNsense, go to Firewall → Rules, select the guest VLAN interface, and create each rule on a separate line. Rule order matters; place blocking rules at the top and allow rules beneath them.
How to Test VLAN Isolation and Connectivity?
After configuring VLANs in OPNsense, it’s important to validate both isolation and connectivity to ensure your network segmentation is working as intended. Testing helps confirm that each VLAN can access only its permitted resources while remaining blocked from restricted areas.
The following steps outline a clear, step-by-step process to test VLAN isolation and connectivity:
- 
Ping Test for Connectivity: From a client device within the VLAN, open a terminal or command prompt. 
  - Ping the VLAN gateway IP (e.g., 192.168.10.1) to confirm local communication.
  - Ping a known external IP (e.g., 8.8.8.8) to verify internet access if it is allowed for that VLAN.
- 
Ping Test for Isolation: Attempt to ping devices in other VLAN subnets (e.g., from VLAN 10 to VLAN 20). If there is no reply, isolation rules are functioning correctly.
- 
Traceroute for Path Verification: Use traceroute (Linux/macOS) or tracert (Windows) to check the path to external and internal destinations. Ensure that traffic follows expected routes without crossing into unauthorized VLANs.
- 
Client Verification: Test with different device types (PC, mobile, IoT) to make sure VLAN tagging and firewall rules work consistently.
For guest VLANs, verify that the captive portal appears if configured.
How to Set Up a Guest Network VLAN on OPNsense?
A guest network VLAN allows you to provide internet access to visitors while keeping your internal network secure. By isolating guest devices, you can prevent unauthorized access to sensitive resources and ensure bandwidth control. The following instructions outline how to configure an OPNsense guest network step by step.
- 
Create the Guest VLAN: To begin setting up your OPNsense guest network, you must first create a dedicated VLAN for guest devices. This VLAN will serve as the isolated segment where all guest traffic is routed, ensuring separation from internal systems. 
  - Log in to the OPNsense GUI.
  - Navigate to Interfaces → Devices → VLAN .
  - Click Add , choose theparent interface (usually your LAN or trunk port), assign aVLAN ID (e.g., 30), and give it a descriptive name such as Guest_VLAN.
- 
Assign the VLAN Interface: Once the VLAN is created, it needs to be assigned to the system so it can function as a separate interface. This step also includes defining the network’s subnet for guest traffic. 
  - Go to Interfaces → Assignments .
  - Add the newly created VLAN to the list and enable it.
  - Configure the interface with a unique subnet (e.g., 192.168.30.1/24).
- Go to 
- 
Set Up DHCP for the Guest VLAN: Guests should automatically receive an IP address when they connect. Setting up DHCP ensures that IP allocation is handled dynamically, making it easier for visitors to join the network. 
  - Navigate to Services → ISC DHCPv4 → [Guest_VLAN] .
  - Enable DHCP and define the IP address pool (e.g., 192.168.30.10 to 192.168.30.100).
- Navigate to 
- 
