---
id: collect-261001-fortinet/fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382-3
title: "Example SSH command to connect to the FortiGate firewall"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382.md
source_anchor: ""
source_lines: [84, 147]
sha256: 7f5eb97957ddb79a156e6a61ba0e9b39cc6697578b646168edafc881098aa31c
---

# Example SSH command to connect to the FortiGate firewall

FortiGate includes physical and virtual network interfaces to connect to various network segments, enabling traffic ingress/egress and network segmentation for security and performance optimization.
FortiGate supports dynamic and static routing protocols to route traffic between different network segments efficiently and securely, ensuring optimal network performance and connectivity.
VLANs allow FortiGate to segment the network into multiple virtual LANs, isolating traffic and improving security, scalability, and performance across large and complex networks.
FortiGate provides a web-based management interface, command-line interface (CLI), and centralized management platforms (FortiManager) for configuring, monitoring, and managing firewall policies, security services, and network settings.
FortiGate logs network activity, security events, and policy violations, generating detailed reports and alerts for administrators to analyze security incidents, troubleshoot issues, and maintain compliance with regulatory requirements.
FortiGate's platform design and architecture leverage these components to deliver robust network security, performance, and scalability for modern enterprise environments.
- NETWORK PROCESSOR 7 (NP7)
- CONTENT PROCESSOR 9 (CP9)
- SECURITY PROCESSING UNIT (SP5)
Link: https://www.fortinet.com/products/fortigate/fortiasic
The FortiGate firewall Command Line Interface (CLI) provides administrators with a powerful and flexible tool for configuring, monitoring, and troubleshooting the firewall. Here's an explanation of the FortiGate firewall CLI:
The CLI is accessed using SSH or through the console port directly connected to the firewall device. It provides a text-based interface where administrators can execute commands to perform various tasks related to firewall configuration and management.
- Configuration Hierarchy: FortiGate CLI follows a hierarchical structure where configuration settings are organized into nested levels, such as system, interface, firewall policy, etc.
- Configuration Commands: Administrators can use CLI commands to view, modify, and commit configuration changes. Commands include show ,get ,set ,edit ,delete ,execute , andend .
- Status Monitoring: CLI commands provide real-time monitoring of system status, interface statistics, CPU and memory usage, VPN connections, and more.
- Diagnostic Tools: FortiGate CLI includes diagnostic tools such as ping ,traceroute ,diag sniff , anddiag debug commands to troubleshoot network connectivity issues and analyze traffic flow.
- Firewall Policies: Administrators can define and manage firewall policies using CLI commands to control traffic flow between different network segments based on source/destination IP, port, protocol, and security profiles.
- Security Profiles: CLI allows configuring security profiles such as antivirus, IPS, web filtering, and application control to enforce security policies and protect against threats.
- VPN Tunnels: CLI commands enable administrators to configure IPsec, SSL, and other types of VPN tunnels to establish secure communication channels between remote sites, users, and partners.
- VPN Monitoring: Administrators can monitor VPN connections, view tunnel status, and troubleshoot VPN-related issues using CLI commands.
- System Configuration: CLI provides commands to configure system settings, including hostname, time zone, DNS, NTP, SNMP, logging, and administrative access controls.
- User Management: Administrators can manage user accounts, authentication methods, and access permissions using CLI commands.
- Granular Control: CLI offers granular control over firewall configuration settings, allowing administrators to customize settings according to specific requirements.
- Scripting and Automation: CLI commands can be scripted and automated using shell scripts or automation tools, facilitating batch configuration changes and streamlining repetitive tasks.
- Direct Access: CLI provides direct access to firewall configuration without the need for a graphical user interface (GUI), making it suitable for advanced users and troubleshooting scenarios.
Accessing the management GUI (Graphical User Interface) of a FortiGate firewall allows administrators to configure and manage the firewall using a web-based interface. Here's how to obtain management GUI access:
- Connect to the FortiGate Firewall
First, establish a connection to the FortiGate firewall. This can be done through the console port directly connected to the firewall device or via SSH (Secure Shell) if remote access is enabled.
# Example SSH command to connect to the FortiGate firewall
ssh admin@<firewall_ip_address>
Ensure that management access is enabled on the FortiGate firewall. By default, HTTPS (HTTP over SSL) access is enabled on port 443 for management GUI access.
# Example command to enable HTTPS access
config system settings
    set admin-https-ssl-port 443
    set gui-mgmt https
    end
Configure administrative access credentials to log in to the management GUI. Ensure that the admin user has the necessary privileges to access and manage the firewall.
# Example command to configure administrative access
config system admin
    edit admin
        set password <admin_password>
    next
end
Once management access is enabled and administrative credentials are configured, access the management GUI using a web browser. Enter the IP address of the FortiGate firewall in the browser's address bar and log in with the administrative credentials.
https://<firewall_ip_address>
- Firewall Rules: Ensure that firewall rules permit traffic to the management interface (usually port 443 for HTTPS) from the IP addresses or networks that require access to the management GUI.
- Security: Use strong, unique passwords for administrative access and regularly update them to enhance security.
- Logging and Monitoring: Monitor access to the management GUI and enable logging to track administrative activities for auditing and security purposes.
Default Username: admin
Password:
Configure the new strong Password
Sample Topology:
Initial CLI Conifguration for GUI access:
Taking GUI Admin Access: http://105.0.0.254
Changing Firewall Hostname: FGT
Welcome to Fortigate Firewall Dashboard
Administration profiles in FortiGate firewall provide a flexible way to manage administrative access and privileges within the firewall. They allow administrators to define specific permissions and restrictions for different users or groups, ensuring secure and efficient management of the firewall. Here's an in-depth look at administration profiles:
Administration profiles serve as templates that define the access rights and capabilities of administrators or administrative groups. Each profile specifies the level of access to various firewall functionalities, including configuration, monitoring, and management tasks.
- Permissions: Administration profiles define permissions for different firewall functionalities, such as configuration changes, system settings, security policies, and VPN configurations.
- Granularity: Profiles can be configured with granular access controls, allowing administrators to assign specific permissions based on their roles and responsibilities.
- Authentication Methods: Administration profiles specify the authentication methods used to verify the identity of administrators, including local authentication, RADIUS, LDAP, or TACACS+.
- Authentication Servers: Profiles can be configured to authenticate users against multiple authentication servers for redundancy and flexibility.
- Role-based Access: Administration profiles support role-based access control (RBAC), allowing administrators to assign different roles with varying levels of privileges.
- Super Administrators: Super administrators have unrestricted access to all firewall functionalities and settings, while other administrators may have limited privileges based on their assigned profiles.
