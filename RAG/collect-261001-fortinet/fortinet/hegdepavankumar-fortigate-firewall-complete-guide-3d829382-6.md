---
id: collect-261001-fortinet/fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382-6
title: "Example SSH command to connect to the FortiGate firewall"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382.md
source_anchor: ""
source_lines: [405, 512]
sha256: 5ed3f64f1d5581968c8d891413fcc184add62a56e9828ea240e0b586c3aa125e
---

# Example SSH command to connect to the FortiGate firewall

- if you want to see the logs, enable the Log Allowd Traffic. - “All Sessions”
- also, traffic should be allowed on both sides, so we have to configure the reverse Policy inorder to get the communication. WAN —> LAN (only limited access)
NAT is a technique used to modify network address information in packet headers while in transit through a router or firewall. It serves several purposes, including conserving IP addresses, enabling connectivity between different network types, and enhancing network security by hiding internal IP addresses.
- Types of NAT:
  - Source NAT (SNAT): Modifies the source IP address of outgoing packets, typically used for outbound internet access.
  - Destination NAT (DNAT): Modifies the destination IP address of incoming packets, commonly used for inbound services such as web servers or email servers.
Static NAT maps a public IP address to a private IP address on a one-to-one basis, allowing external hosts to initiate connections to internal hosts.
config firewall ippool
    edit "public_pool"
        set type static
        set address <public_ip_range>
    next
end
config firewall vip
    edit "static_nat"
        set extintf "wan1"
        set extip <public_ip>
        set mappedip <private_ip>
    next
end
Port forwarding redirects traffic from a specific port on the firewall's public IP address to an internal IP address and port.
config firewall vip
    edit "port_forwarding"
        set extintf "wan1"
        set extip <public_ip>
        set mappedip <private_ip>
        set protocol <protocol>
        set extport <public_port>
        set mappedport <private_port>
    next
end
Assigning a static IP address to a device ensures consistency and predictability in network configurations, particularly for devices requiring consistent access or services.
config system interface
    edit <interface_name>
        set ip <ip_address> <subnet_mask>
        set allowaccess <access_options>
    next
end
Configuring IP addresses on interfaces enables communication between different network segments and defines the gateway for traffic leaving the subnet.
config system interface
    edit <interface_name>
        set ip <ip_address> <subnet_mask>
        set allowaccess <access_options>
    next
end
By visiting Policy & Objects —> LAN-WAN policy
- NAT should be enabled.
- Instead of Interface IP —> Use Dynamic IP Pool —> by clicking
- we can able to select the types of NAT we want to perform.
- Overload, One-to-One, Fixed Port Range - choose as per your requirement and apply it.
Enter the IP address You want NAT.
If you want to Map the original protocol number with custom - you can do it by configuring the Protocol Option. Port Mapping.
Certainly! FortiGate Virtual Wire (VW) is a feature that allows you to transparently insert security services, such as firewall policies and intrusion prevention systems (IPS), into the network without changing the IP addressing or topology. It operates at Layer 2 of the OSI model, meaning it doesn't require IP addresses to be changed, making it ideal for scenarios where IP addressing cannot be modified easily.
Here's a breakdown of the key theoretical aspects of FortiGate Virtual Wire:
1. Layer 2 Operation:
- Virtual Wire operates at Layer 2 (Data Link Layer) of the OSI model, which means it deals with MAC addresses rather than IP addresses. This allows the FortiGate firewall to seamlessly intercept and inspect traffic passing through it without requiring any IP address changes.
- Since Virtual Wire operates at Layer 2, it can't perform routing or NAT (Network Address Translation). Instead, it forwards packets based on MAC addresses.
2. Transparent Traffic Inspection:
- Virtual Wire enables the insertion of security services, such as firewall policies, intrusion prevention systems (IPS), and antivirus scanning, into the network path without disrupting normal network operations.
- Traffic passing through the Virtual Wire is transparently inspected by the FortiGate firewall, which can enforce security policies and detect and mitigate threats in real-time.
3. In-line Deployment:
- In Virtual Wire deployment, the FortiGate firewall sits in-line between two network segments, intercepting traffic as it passes through.
- It typically involves configuring two physical interfaces on the FortiGate firewall—one for inbound traffic (ingress interface) and the other for outbound traffic (egress interface).
4. Traffic Forwarding and Filtering:
- Once traffic enters the Virtual Wire, it is forwarded to the appropriate egress interface based on the configured security policies.
- The firewall inspects the traffic according to predefined security rules, including firewall policies, IPS signatures, antivirus scans, and other security profiles.
- If the traffic matches any security policy, the firewall takes the specified action (e.g., allow, deny, log).
5. VLAN Support:
- Virtual Wire supports VLANs, allowing you to segment traffic within the Virtual Wire deployment.
- You can assign VLAN IDs to the Virtual Wire configuration to handle tagged VLAN traffic between network segments.
6. Simplified Deployment and Management:
- Virtual Wire simplifies the deployment of security services by eliminating the need for complex network reconfigurations.
- It also simplifies management by providing a transparent way to insert security services into the network path, reducing operational overhead and minimizing disruption to network operations.
Here's a detailed explanation of the concept along with configuration steps:
1. Security Policy Configuration:
Create security policies to define how traffic is handled by the Virtual Wire pair. This includes specifying the source and destination zones, as well as the security profiles to be applied (e.g., IPS, antivirus).
config firewall policy
    edit 1
        set srcintf "port2"
        set dstintf "port3"
        set action accept
        ...
    next
end
- srcintf : Specifies the source interface.
- dstintf : Specifies the destination interface.
- action : Defines the action to be taken on the traffic (e.g., accept, deny).
2. Monitoring and Logging:
Configure logging and monitoring to track traffic passing through the Virtual Wire for security analysis and troubleshooting purposes.
config log
    set status enable
    ...
end
3. Testing and Verification:
Test the Virtual Wire configuration to ensure that traffic is being inspected and forwarded correctly without any disruptions to network connectivity.
This configuration enables the FortiGate unit to operate in Virtual Wire mode, transparently inspecting and filtering traffic between two network segments without requiring any changes to IP addressing or network topology.
Sample Topology:
To configure Virtual Wire, go to Interface --> Create New:
Now new Virtual Pair Interface is Configured:
As per the Fortigate we have to configure the Firewall Virtual Wire Pair Policy, go to Policy & Objects --> Firewall Virtual Wire Pair Policy --> create bidirectional policy:
Now is the Time to Initiate the traffic towards the internet, all the traffic will be available in the firewall, Logs --> Forwarded Traffic:
Check the logs to verify the Source and destination information, Traffic from PC to Internet.:
- To get Internet access to the PC, we have configured Static NAT on the Router[Edge_R].
- Without NAT we are not able to access the internet, the ISP drops the packet.
- Because Private IPs are not routable in ISPs.
In Module 2, we covered essential topics related to configuring interfaces and firewall policies on FortiGate firewall. Here's a summary of the topics covered:
- Basic Interface Configuration: Explained how to configure interfaces on FortiGate firewall, including setting IP addresses, subnet masks, and access permissions.
- Configure Static and Dynamic Routing: Detailed the configuration of static and dynamic routing protocols such as OSPF and BGP on FortiGate firewall to enable efficient traffic forwarding.
