---
id: collect-261001-automatisation-infra/automatisation-infra/hegdepavankumar-eve-ng-labs-3a58a934-2
title: "hegdepavankumar-eve-ng-labs-3a58a934"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "latency", "throughput"]
source: docs/RAG/collect-261001-automatisation-infra/hegdepavankumar-eve-ng-labs-3a58a934.md
source_anchor: ""
source_lines: [150, 262]
sha256: 15729050908e462c194013718050b403ec314333b6b9146402e0dc8eb323ae1d
---

# hegdepavankumar-eve-ng-labs-3a58a934

- IP addressing and routing configurations must be efficient and well-documented to facilitate troubleshooting and management.
Deliverables:
- A detailed network diagram illustrating the access, distribution, and core layers, including connections and configurations.
- An IP addressing scheme and VLAN assignments.
- HSRP configuration details for the core switches.
- Routing configuration for internet access.
- Documentation of all configurations and design decisions.
📁Download Topology : click here 🔫
Problem Statement: Simulating a Spine-Leaf Network Architecture
Objective:
Design and simulate a spine-leaf network architecture for a high-performance data center to achieve scalability, low latency, and high throughput. This simulation will provide insights into traffic management and the interactions between spine switches, leaf switches, and end devices.
Requirements:
- 
Network Design: 
  - Implement a spine-leaf network topology consisting of multiple spine switches and leaf switches.
  - Ensure that each leaf switch is connected to every spine switch to create a non-blocking network with equal bandwidth paths.
  - Designate a sufficient number of spine and leaf switches to meet the performance and scalability goals.
- 
Performance Metrics: 
  - Achieve low latency and high throughput across the network.
  - Optimize traffic management to prevent bottlenecks and ensure efficient data transfer between end devices.
- 
Scalability: 
  - Design the network to easily scale by adding additional spine or leaf switches without significant changes to the existing infrastructure.
  - Include mechanisms to handle increased traffic and device density effectively.
- 
Simulation Goals: 
  - Use network simulation tools to model and analyze the performance of the spine-leaf architecture.
  - Evaluate key performance indicators such as latency, throughput, and packet loss.
  - Simulate various traffic patterns and workloads to test the network’s efficiency and reliability.
- 
Traffic Management: 
  - Implement and test traffic distribution strategies to ensure balanced load across the spine and leaf switches.
  - Configure Quality of Service (QoS) policies as needed to prioritize critical traffic and manage network resources effectively.
- 
End Devices: 
  - Include a variety of end devices (nodes) in the simulation to represent typical data center workloads.
  - Assess the performance impact of different types of end devices and their interaction with the spine-leaf network.
Constraints:
- Ensure that the simulation accurately reflects real-world scenarios and network conditions.
- The design should be cost-effective and feasible with available hardware and software resources.
- Provide documentation of the simulation setup, configurations, and results to support analysis and decision-making.
Deliverables:
- A detailed network diagram of the spine-leaf architecture including spine and leaf switch configurations.
- Simulation results highlighting performance metrics such as latency, throughput, and traffic distribution.
- Analysis of the network's ability to scale and handle various traffic patterns.
- Recommendations for optimizing the spine-leaf network based on simulation findings.
📁Download Topology : click here 🔫
Objective:
The goal is to ensure all devices are properly configured to communicate within their subnets and across the network, providing seamless connectivity to the Internet and internal resources.
You have been tasked with configuring a small office/home office (SOHO) network. The network diagram provided outlines the structure of the network, including various devices, subnets, and their interconnections. The goal is to ensure all devices are properly configured to communicate within their subnets and across the network, providing seamless connectivity to the Internet and internal resources.
- Router Configuration:
  - The router connects the internal network to the Internet.
  - The WAN interface (Gi0/1) should be configured to connect to the Internet using the subnet 192.168.1.0/24.
  - The LAN interface (Gi0/0) should be configured with the subnet 192.168.20.0/24.
- FortiGate Firewall Configuration:
  - The firewall serves as the primary security device between the internal network and the router.
  - Configure port1 to connect to the switch using the subnet 192.168.20.0/24.
  - Configure port2 to connect to the router using the subnet 192.168.20.0/24.
- Switch Configuration:
  - The switch connects multiple devices within the internal network.
  - Ensure all connected devices can communicate within the subnet 192.168.10.0/24.
  - Interfaces Gi0/0 to Gi0/3 should be configured to connect to PCs and workstations.
  - Interfaces Gi1/0 to Gi1/3 should be configured for other devices like laptops, IP phones, and tablets.
  - Interface Gi2/0 should be configured for the Wireless Access Point.
- Wireless Access Point Configuration:
  - The Wireless Access Point (WAP) should provide wireless connectivity for devices within the subnet 192.168.10.0/24.
- Device Configurations:
  - Each device on the network must be assigned an IP address within the appropriate subnet:
    - PCs, Workstations, IP Phones, and Tablets: 192.168.10.0/24
    - Wireless Laptops: 192.168.10.0/24 via WAP
- Each device on the network must be assigned an IP address within the appropriate subnet:
- Additional Requirements:
  - Ensure proper VLAN configurations if necessary to separate different types of traffic.
  - Implement DHCP where applicable to automate IP address assignment for devices.
  - Enable appropriate firewall rules to allow necessary traffic while blocking unauthorized access.
📁Download Topology : click here 🔫
VLAN Configuration and Testing - to check the basic knowledge of switching.
You are tasked with setting up a VLAN on two switches to ensure proper communication between devices on the same VLAN. The specific requirements are as follows:
- Create VLAN 10:
- VLAN 10 needs to be created on both Switch 1 and Switch 2.
- Configure Trunk Ports:
- The G0/0 interface on both Switch 1 and Switch 2 should be configured as trunk ports. These trunk ports will carry traffic for VLAN 10.
- Assign Ports to VLAN 10:
- Both Switch 1 and Switch 2 should have specific ports assigned to VLAN 10. These ports will be used by devices that need to communicate with each other.
- Test Communication:
- Ensure that devices connected to the ports assigned to VLAN 10 on Switch 1 and Switch 2 can communicate with each other.
Tasks Breakdown
- VLAN Creation:
- On Switch 1, create VLAN 10.
- On Switch 2, create VLAN 10.
- Trunk Port Configuration:
- On Switch 1, configure interface G0/0 as a trunk port.
- On Switch 2, configure interface G0/0 as a trunk port.
- **Assign Ports to VLAN 10:
- On Switch 1, assign the desired ports to VLAN 10.
- On Switch 2, assign the desired ports to VLAN 10.
- Testing:
- Connect two devices to the assigned ports on Switch 1 and Switch 2, respectively.
- Verify that the devices can communicate with each other, indicating that VLAN 10 is correctly configured and operational across both switches.
📁Download Topology : click here 🔫
📝 Problem Statement: Problem Statement: Inter-VLAN Routing Solution with Router-on-a-Stick Configuration
Objective:
Design and implement an inter-VLAN routing solution using the router-on-a-stick configuration to enable communication between VLANs on a network. The goal is to configure and verify connectivity between VLANs, ensuring seamless communication while maintaining network segmentation and security.
Topology:
- Router (Edge-R): Equipped with two subinterfaces on its g0/0 interface:
  - Subinterface for VLAN 10 (IT) with IP address 10.1.1.100
  - Subinterface for VLAN 20 (HR) with IP address 20.1.1.100
- Switch: Configured with two VLANs:
  - VLAN 10 for the IT department
  - VLAN 20 for the HR department
- PCs:
  - PC-1 and PC-2: Located in VLAN 10 (IT) with IP addresses in the 10.1.1.0/24 subnet
