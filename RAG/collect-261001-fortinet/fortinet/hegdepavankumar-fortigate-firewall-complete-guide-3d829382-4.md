---
id: collect-261001-fortinet/fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382-4
title: "Example SSH command to connect to the FortiGate firewall"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber", "ethernet"]
source: docs/RAG/collect-261001-fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382.md
source_anchor: ""
source_lines: [148, 285]
sha256: c1f29aa291088875d013d0b2e07194438ae622b6d57038a529acd3a62f0d4479
---

# Example SSH command to connect to the FortiGate firewall

- Session Limits: Administration profiles can define session limits to control the number of concurrent administrative sessions allowed per user or group.
- Timeouts: Profiles specify session timeouts to automatically log out inactive administrators, enhancing security and resource management.
- Creation: Administration profiles are created and configured through the FortiGate firewall's web-based management GUI or Command Line Interface (CLI).
- Naming Conventions: Profiles are assigned unique names for easy identification and management.
- Assignment to Administrators: Once created, administration profiles are assigned to individual administrators or administrative groups based on their roles and responsibilities.
- Multiple Profiles: Administrators can be assigned multiple profiles to accommodate complex access requirements.
Click on Administrator: System —> Administrator
By default Super_Admin
Set the username/password also select as a Local user:
Create the new Administrator Profile for the new user
Select the Permission Which you want to give and click Ok.
Note:
The idle timeout period is the amount of time that an administrator will stay logged in to the GUI without any activity. This is to prevent someone from accessing the FortiGate if the management PC is left unattended. By default, it is set to five minutes.
Select newly created Profile and Click OK.
Newly created Admin :
Administrator Profile Hierachy:
The module provided a basic introduction to FortiGate firewall, covering various aspects of the product:
- Understanding Features of FortiGate: Explains the features and capabilities of FortiGate firewall, highlighting its advanced threat protection, network segmentation, and secure connectivity.
- FortiGuard Queries & Packages: Discusses FortiGuard services, including threat intelligence and security updates provided by Fortinet, enhancing the effectiveness of the firewall in detecting and preventing threats.
- UTM Firewalls Features: Describes UTM (Unified Threat Management) features of FortiGate, which include firewall, intrusion prevention, antivirus, web filtering, and application control, offering comprehensive protection against various cyber threats.
- Platform Design and Architecture: Explores the design and architecture of FortiGate firewall, including its processing units, security services, networking components, and management features.
- About CLI: Provides an overview of the FortiGate Command Line Interface (CLI), which allows administrators to configure, monitor, and troubleshoot the firewall using text-based commands.
- Getting Mgmt GUI Access: Details the steps to access the management GUI (Graphical User Interface) of FortiGate firewall, allowing administrators to configure and manage the firewall through a web-based interface.
- About Administration Profiles: Discusses administration profiles in FortiGate, which define access rights and privileges for administrators or administrative groups, ensuring secure and efficient management of the firewall.
Table of contents:
- Basic Interface Configuration
- configure static and dynamic routing
- Configuring DHCP
- Basic Firewall Policies
- Network Address Translation - Fortigate
- Virtual Wire configuration
Configuring interfaces on a FortiGate firewall is essential for establishing network connectivity and defining traffic flow. Here's a detailed guide on how to perform basic interface configuration using commands:
Before configuring interfaces, establish a connection to the FortiGate firewall using SSH or through the console port directly connected to the firewall device.
Example SSH command to connect to the FortiGate firewall
ssh admin@<firewall_ip_address>
Enter configuration mode to make changes to the firewall's configuration. You will need to enter the global configuration context to configure interfaces.
# Enter global configuration mode
config system global
FortiGate firewalls have physical interfaces (e.g., Ethernet ports) that connect to the network. Configure the desired physical interfaces with appropriate IP addresses and other settings.
# Example command to configure physical interface
edit system interface
    edit <interface_name>
        set ip <ip_address> <subnet_mask>
    next
end
If VLANs (Virtual Local Area Networks) are used to segment the network, configure VLAN interfaces and assign them to the desired physical interfaces.
# Example command to configure VLAN interface
edit system interface
    edit <vlan_interface_name>
        set vlanid <vlan_id>
        set ip <ip_address> <subnet_mask>
    next
end
Virtual interfaces such as loopback interfaces can be configured for various purposes, such as management or routing.
# Example command to configure loopback interface
edit system interface
    edit <loopback_interface_name>
        set ip <ip_address> <subnet_mask>
    next
end
Specify the default gateway for the firewall to enable outbound traffic routing to external networks.
# Example command to configure default gateway
config router static
    edit 1
        set gateway <gateway_ip_address>
end
Save the configuration changes to persist them across reboots.
# Save configuration
end
- Interface Naming: Use meaningful names for interfaces to easily identify their purpose and location.
- Security Policies: After configuring interfaces, create firewall policies to control traffic flow between interfaces and enforce security rules.
- Monitoring: Regularly monitor interface status and traffic to detect any issues or anomalies.
Sample Lab topology:
The Management Interface Configurations we have done through CLI:
Configure the as per the below image:
Steps:
- Add Alias LAN or any meaningful name to identify the LAN
- Set the role as LAN as per our Topology.
- Select Manual and assign IP for interface(port2)
- Allow the only required Protocols
- Make sure the Interface status is Enabled.
- Save the configurations by clicking OK.
NOTE: Also we can configure the interfaces via CLI
# LAN port2 interface
config system interface
edit port2
set mode static
set ip 10.1.1.100/24
set allowaccess ping
set alias "LAN"
set role lan
end
# WAN port3 interface
config system interface
edit port3
set mode static
set ip 192.168.1.100/24
set allowaccess ping
set alias "WAN"
set role wan
end
# DMZ port4 interface
config system interface
edit port4
set mode static
set ip 172.16.1.100/24
set allowaccess ping
set alias "DMZ"
set role dmz
end
# To see the configuration on CLI
show system interface
Configure the remaining WAN & DMZ interfaces same as the previous one.
Routing is a critical function in network devices like FortiGate firewalls, enabling the forwarding of traffic between different networks. Here's a detailed guide on how to configure static and dynamic routing using commands:
Before configuring routing, establish a connection to the FortiGate firewall using SSH or through the console port directly connected to the firewall device.
# Example SSH command to connect to the FortiGate firewall
ssh admin@<firewall_ip_address>
Enter configuration mode to make changes to the firewall's configuration. You will need to enter the global configuration context to configure routing.
# Enter global configuration mode
config system global
Static routes are manually configured routes that define the next-hop IP address for destinations not directly connected to the firewall.
# Example command to configure a static route
config router static
    edit 1
        set dst <destination_network> <subnet_mask>
        set gateway <next_hop_ip_address>
end
FortiGate firewalls support dynamic routing protocols such as OSPF (Open Shortest Path First) and BGP (Border Gateway Protocol) for dynamic route exchange and network convergence.
# Example command to configure OSPF
config router ospf
    set router-id <router_id>
    config area
        edit <area_id>
            set network <area_network> <area_subnet_mask>
    end
    config redistribute connected
        set status enable
    end
