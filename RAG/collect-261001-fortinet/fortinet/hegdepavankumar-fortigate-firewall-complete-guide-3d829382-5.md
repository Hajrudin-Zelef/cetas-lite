---
id: collect-261001-fortinet/fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382-5
title: "Example SSH command to connect to the FortiGate firewall"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382.md
source_anchor: ""
source_lines: [286, 404]
sha256: 863c2d084797ef1593a9347871f321ec9be43dcfbb0d825248dfc7fb287c3ab0
---

# Example SSH command to connect to the FortiGate firewall

end
# Example command to configure BGP
config router bgp
    set as <autonomous_system_number>
    config neighbor
        edit <neighbor_ip_address>
            set remote-as <neighbor_as_number>
            set capability-default-originate enabl
    end
end
After configuring static and dynamic routing, verify the routing table and routing protocol status to ensure correct configuration.
# Example command to view routing table
get router info routing-table
# Example command to view OSPF neighbor status
get router info ospf neighbor
# Example command to view BGP neighbor status
get router info bgp neighbor
Save the configuration changes to persist them across reboots.
# Save configuration
end
- Route Summarization: Use route summarization to reduce the size of the routing table and optimize routing efficiency.
- Redundancy: Implement redundancy and failover mechanisms such as ECMP (Equal-Cost Multi-Path) and HA (High Availability) to ensure network availability and reliability.
- Security Policies: After configuring routing, create appropriate firewall policies to control traffic flow between networks and enforce security rules.
All the Routing Parts will be available Network Tab Section:
We can configure the Static Route towards Our Wifi Router/GW to get the internet access.
Steps:
- Go to Network —> Static Routes
- Click on Create New
Assign the Wifi Router IP address, Make it Destination 0.0.0.0/0.0.0.0 Any Any
- By default, it takes the WAN interface.
- click on OK to save the configurations.
To configure dynamic routing protocols like RIPv2, OSPF, BGP
Steps to configure RIPv2:
- select the required version
- and enter the network/subnet that you want
- if any passive interface or authentication is configured make sure that match the md5 key.
- make sure that the Hello and Hold timers are matching.
- If any redistribute want to do turn on the toggle-specific protocol.
Follow the Image with your real network and conditions.
Configuring a DHCP server pool on the LAN interface of a FortiGate firewall allows local users to obtain IP addresses automatically, simplifying network management. Here's a detailed guide on how to configure the DHCP server pool for local users:
Before configuring DHCP, establish a connection to the FortiGate firewall using SSH or through the console port directly connected to the firewall device.
# Example SSH command to connect to the FortiGate firewall
ssh admin@<firewall_ip_address>
Enter configuration mode to make changes to the firewall's configuration. You will need to enter the system interface context to configure the LAN interface.
# Enter system interface configuration mode
config system interface
If not already configured, configure the LAN interface with an IP address and subnet mask.
# Example command to configure LAN interface
edit <lan_interface_name>
    set ip <ip_address> <subnet_mask>
    set allowaccess ping https ssh
    set dhcp-server enable
    set dhcp-server-option lease-time <lease_time_in_seconds>
    set dhcp-server-option default-gateway <gateway_ip_address>
    set dhcp-server-ip-range <start_ip_address> <end_ip_address>
end
- <lan_interface_name> : Name of the LAN interface (e.g., "lan").
- <ip_address> : IP address of the LAN interface.
- <subnet_mask> : Subnet mask of the LAN interface.
- <lease_time_in_seconds> : Lease time for IP addresses (in seconds).
- <gateway_ip_address> : Default gateway IP address for DHCP clients.
- <start_ip_address> : Start IP address of the DHCP IP pool.
- <end_ip_address> : End IP address of the DHCP IP pool.
Optionally, configure DNS server settings for DHCP clients.
# Example command to configure DNS server for DHCP clients
set dhcp-server-option dns-server <dns_server_ip_address>
Save the configuration changes to persist them across reboots.
# Save configuration
end
- DHCP Lease Time: Adjust the DHCP lease time based on network requirements and usage patterns.
- IP Address Range: Ensure that the DHCP IP address range does not overlap with statically assigned IP addresses or other DHCP pools.
- DNS Configuration: Provide accurate DNS server information to DHCP clients for name resolution.
- Security: Limit access to DHCP configuration and ensure proper firewall policies are in place to protect the DHCP service from unauthorized access.
To configure the DHCP server go to Network —> Interface —> port2(LAN)
- enable the DHCP server toggle
- set the IP range by excluding the static IP- to avoid the conflict of IP's
- set the netmask and default gateway as interface IP
- set the DNS server addresses- any
- set the lease time as your requirement
- set apply to save the config.
Firewall policies on the FortiGate firewall define how traffic is allowed or denied between different network segments. Understanding basic firewall policy configurations and the theory behind rule-by-fault behavior is essential for effective network security. Here's a detailed explanation:
Firewall policies are rules that dictate the flow of traffic through the firewall. Each policy consists of conditions, actions, and security profiles. Policies are evaluated in sequence, and the first matching policy is applied to the traffic.
- Source and Destination: Specify the source and destination IP addresses or address groups for the traffic.
- Service: Define the protocol and port number or service group used by the traffic.
- Schedule: Optionally, restrict when the policy is active based on a defined schedule.
- Action: Specify whether the traffic is allowed, denied, or logged.
- Accept: Allow the traffic to pass through the firewall.
- Deny: Block the traffic and generate a log entry.
- Monitor: Log the traffic but allow it to pass through the firewall.
- Antivirus: Scan files for viruses and malware.
- Intrusion Prevention System (IPS): Detect and prevent network-based attacks.
- Web Filtering: Block access to malicious or inappropriate websites.
- Application Control: Control access to specific applications and protocols.
FortiGate firewall follows the rule by default behavior, where traffic that does not match any firewall policy is implicitly denied by default. This behavior ensures that only explicitly permitted traffic is allowed to traverse the firewall, enhancing network security.
By default, FortiGate firewall includes an implicit "Deny All" policy at the end of the policy list. This policy denies all traffic that does not match any preceding policy. Administrators can modify this behavior by adding specific allow policies above the "Deny All" policy.
- Rule Ordering: Arrange firewall policies from most specific to least specific to ensure that traffic matches the intended policy.
- Logging: Enable logging for deny actions to monitor and analyze network traffic effectively.
- Regular Review: Periodically review and update firewall policies to adapt to changes in network requirements and security threats.
config firewall policy
    edit 1
        set srcintf "internal"
        set dstintf "external"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set logtraffic all
end
To configure policy go to Policy & Objects —> Firewall Policy
- There will be one by default policy present - Implicit deny
- by clicking create new we can create a new policy on top of Implicit deny.
- It will override the Implicit deny policy.
- For example any traffic does not match with configured policy, it will discards the packet as per the implicit deny policy.
To enable the traffic for LAN —> WAN
- give the meaningful name for ‘Name’ to identify the purpose of the policy.
- make sure you have selected LAN and WAN interfaces as incoming and outgoing interfaces respectively.
- Source, Destination, and Services set to “all” Initially set it as all, for learning - once you are clear with concepts apply the specific one.
- allow the NAT.
