---
id: collect-261001-fortinet/fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382-8
title: "Example SSH command to connect to the FortiGate firewall"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/hegdepavankumar-fortigate-firewall-complete-guide-3d829382.md
source_anchor: ""
source_lines: [570, 660]
sha256: c4d6cfb7d6554a672a7a44a82bc856d54960e844163de02961e1dea1e8afe93b
---

# Example SSH command to connect to the FortiGate firewall

During active-active HA load balancing, the primary unit uses the configured load balancing schedule to determine which cluster unit will process a session. The primary unit stores the load-balancing information for each load-balanced session in the cluster load-balancing session table. Using the information in this table, the primary unit can then forward all of the remaining packets in each session to the appropriate cluster unit. The load balancing session table is synchronized among all cluster units.
ICMP, multicast, and broadcast sessions are never load-balanced and are always processed by the primary unit. The following sessions are only processed by the primary unit:
- IPS
- Application control
- Flow-based virus scanning
- Flow-based web filtering
- Flow-based DLP
- Flow-based email filtering
- VoIP
- IM
- P2P
- IPsec VPN
- SSL VPN
- HTTP multiplexing
- SSL offloading
- WAN optimization
- Explicit web proxy
- WCCP
In addition to load balancing, active-active HA provides the same session, device, and link failover protection as active-passive HA. If the primary unit fails, a subordinate unit becomes the primary unit and resumes operating the cluster. Active-active HA maintains as many load balanced sessions as possible after a failover by continuing to process the load balanced sessions that were being processed by the cluster units that are still operating.
If a subordinate unit fails, the primary unit redistributes the sessions that the subordinate was processing among the remaining active cluster members. If the primary unit fails, the subordinate units negotiate to select a new primary unit. The new primary unit continues to distribute packets among the remaining active cluster units.
Failover works similarly if the cluster consists of only two units. If the primary unit fails, the subordinate unit negotiates and becomes the new primary unit. If the subordinate unit fails, the primary unit processes all traffic. In both cases, the single remaining unit continues to function as a primary unit, maintaining the HA virtual MAC address for all of its interfaces.
- Enable HA: Configure active-active HA mode on both firewall units.
- Synchronize Configuration: Ensure that both units have identical configurations, including network interfaces, routing, and security policies.
- Configure Load Balancing: Define load-balancing algorithms and link load-balancing settings to distribute traffic evenly across the cluster.
- Symmetric Policies: Create symmetric firewall policies that apply to both active units, ensuring consistent enforcement of security rules.
- Policy Prioritization: Prioritize firewall policies to ensure that critical traffic is processed efficiently and does not experience delays.
- Health Monitoring: Continuously monitor the health and status of both firewall units, including CPU usage, memory utilization, and interface status.
- Alerting: Configure alert notifications to promptly notify administrators of any issues or failures within the HA cluster, allowing for quick resolution.
Sample Topology:
FGT-1 device Active-Active HA configuration, Follow the steps.
FGT-2 device HA configuration, Make sure that both are in the same group and Password inorder to synchronize.FGT-2 GUI access will not be available.
Both FGT-1 and FGT-2 are synchronized, FGT-1 will be master/primary, and FGT-2 will be secondary.
As per the priority, we have decided that FGT-1 128 is Primary and FGT-2 100 will be elected as Secondary.
FGT-1 failure occurs, and FGT-2 takes the role of Primary.
Module 3 of the FortiGate firewall course covers high availability (HA) configurations, including both active-standby and active-active setups. Here's a brief summary of the topics covered:
- Active-Standby (Theory): Explains the theory behind active-standby HA configurations, where one firewall unit serves as the primary (active) unit while the other acts as the secondary (standby) unit. In case of failure, the standby unit takes over seamlessly to ensure continuous operation.
- Active-Standby (Lab): Provides hands-on lab exercises for configuring active-standby HA on FortiGate firewall units. Students learn how to set up HA links, synchronize configurations, and test failover scenarios.
- Active-Active (Theory): Discusses the theory behind active-active HA configurations, where both firewall units actively process network traffic simultaneously. Load balancing, firewall state synchronization, and session pickup are explained in detail.
- Active-Active (Lab): Offers practical lab exercises for configuring active-active HA on FortiGate firewall units. Students learn how to configure load balancing algorithms, ensure firewall state synchronization, and test failover scenarios in an active-active setup.
Table of contents:
- Creating User and Policies
- Create Authentication Policies[Captive Portal]
- Monitor firewall Users
Creating users and policies on the FortiGate firewall allows administrators to control access to network resources and define security rules for traffic flow. Here's a detailed guide on how to create users and policies:
- Purpose: Local users are specific to the FortiGate firewall and are authenticated locally.
- Steps:
  - Access the FortiGate firewall web interface or CLI.
  - Navigate to System > Administrators > Administrators.
  - Click Create New to add a new user.
  - Enter the user details, including username, password, and privileges.
  - Save the changes.
- Purpose: FortiGate firewall supports external authentication methods such as LDAP, RADIUS, or TACACS+ for user authentication.
- Steps:
  - Configure external authentication servers under User & Device > Authentication > LDAP/Radius/TACACS+.
  - Navigate to System > Administrators > Administrators.
  - Click Create New to add a new user.
  - Select the external authentication server from the dropdown list.
  - Enter the username and assign privileges.
  - Save the changes.
- Purpose: Firewall policies control traffic flow between different network segments based on defined criteria.
- Steps:
  - Access the FortiGate firewall web interface or CLI.
  - Navigate to Policy & Objects > IPv4 Policy.
  - Click Create New to add a new policy.
  - Configure the policy details, including source and destination addresses, services, and action (allow/deny).
  - Optionally, configure security profiles such as antivirus, IPS, or web filtering.
  - Save the changes.
- Purpose: VPN policies define how VPN traffic is handled by the FortiGate firewall, including IPsec or SSL VPN tunnels.
- Steps:
  - Access the FortiGate firewall web interface or CLI.
  - Navigate to VPN > IPsec > Tunnels or VPN > SSL-VPN > Portals/Settings.
  - Click Create New to add a new VPN tunnel or portal.
  - Configure the tunnel/portal settings, including encryption algorithms, authentication methods, and access controls.
  - Save the changes.
Sample Topology:
Now we are configuring Captive Portal or User Authentication.
Steps:
- we need to create the Local user to be authenticated before accessing the web server.
- Create the user and assign that user to a particular group, or make it default as a Guest-group.
Select User Type —> Local user
Create the login credentials and disable the two-factor authentication for now.:
Enable the User account status and assign a default group or create a new one.
Our new user has been created and assigned to the default group.
Here we are authenticating LAN users to access the DMZ web browser. below image shows without Captive portal or firewall authentication.
Assign created Policy to LAN interface,
Steps:
- go to Network —> Interface
- enable the security mode and captive portal for Local authentication for Guest group.
We are going to access DMZ web server but this time we don't get direct DMZ login page, before accessing we have to be authenticated by Firewall
I have entered valid credentials so I can able to access the DMZ server.
Getting DMZ Home Page after Successfully Authenticated.
