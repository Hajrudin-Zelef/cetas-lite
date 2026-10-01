---
id: collect-261001-automatisation-infra/automatisation-infra/hegdepavankumar-eve-ng-labs-3a58a934-1
title: "hegdepavankumar-eve-ng-labs-3a58a934"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "ethernet"]
source: docs/RAG/collect-261001-automatisation-infra/hegdepavankumar-eve-ng-labs-3a58a934.md
source_anchor: ""
source_lines: [1, 149]
sha256: 3a64afc261df8c011b36f30c80feda9f2ff8d865a9ec0daa8b9a025deb2e7340
---

# hegdepavankumar-eve-ng-labs-3a58a934

Welcome to the EVE-NG Labs Catalog! This repository is a comprehensive collection of network simulation labs designed for network engineers, students, and professionals who want to enhance their skills in routing, switching, and firewall technologies. These labs utilize EVE-NG (Emulated Virtual Environment - Next Generation) to provide a hands-on, practical learning experience.
- Description: Dive into dynamic and static routing. This section includes labs on OSPF, EIGRP, BGP, and more.
- Devices: Cisco Routers & Switches (various models)
- Description: Learn to secure networks with industry-leading firewall technologies. This section includes labs on:
  - Cisco ASA & Firepower 🛡️
  - Palo Alto 🔐
  - Checkpoint 🔲
  - FortiGate 🏰
Each lab comes with:
- Topology Files: Topology files for EVE-NG. (Compatible with both Community and PRO)✅
- Configuration Files: Initial and final configuration files. ❌
- Problem Statements: Detailed problem statements to guide you through each lab.✅
To get started, follow these steps:
- Import Labs into EVE-NG:
By working through these labs, you will:
- Gain hands-on experience with routing and switching protocols.
- Understand the configuration and management of various firewalls.
- Develop troubleshooting skills in a simulated network environment.
- Exploring Multiple real-time Problem statement.
Objective: Configure IPv4 and IPv6 Addresses on Cisco Routers
- You are tasked with configuring both IPv4 and IPv6 addresses on a Cisco router.
- The router has two interfaces: GigabitEthernet 0/0 and GigabitEthernet 0/1.
- The IPv4 address for GigabitEthernet 0/0 should be 192.168.1.1/24 and for GigabitEthernet 0/1 should be 192.168.2.1/24.
- The IPv6 address for GigabitEthernet 0/0 should be 2001:192:168:1::1/64 and for GigabitEthernet 0/1 should be 2001:192:168:2::1/64
📁Download Topology : Click Here 🔫
Objective: Cisco network setup consisting of both Layer 2 and Layer 3 switches
- Assign an IP address to the Layer 3 switch.
- Enable routing on the Layer 3 switch.
- Create VLANs and assign them to specific ports on both Layer 2 and Layer 3 switches.
Requirements:
- Assign IP Address to Layer 3 Switch:
- Choose an appropriate IP address and subnet mask for the management interface of the Layer 3 switch.
- Enable Routing on Layer 3 Switch:
- Enable IP routing to allow the Layer 3 switch to route traffic between different VLANs.
Create VLANs:
- Define multiple VLANs (e.g., VLAN 10 for the Sales department, VLAN 20 for the HR department).
- Assign these VLANs to specific ports on both Layer 2 and Layer 3 switches.
- Ensure inter-VLAN routing is configured to allow communication between VLANs on the Layer 3 switch.
📁Download Topology : Click Here 🔫
Objective:
Design and understand a basic 2-tier network topology. This topology will include end devices (such as computers or workstations) connected to a switch, which in turn connects to a router providing access to external networks like the internet. This exercise aims to help you grasp the fundamental components and functions of a simple network architecture.
Components:
- 
End Devices: 
  - 4 Computers (PC1, PC2, PC3, PC4)
- 
Network Devices: 
  - 1 Switch (Switch1)
  - 1 Router (Router1)
- 
Network Connections: 
  - Ethernet cables to connect PCs to the switch
  - An Ethernet cable to connect the switch to the router
Network Requirements:
- 
IP Addressing: 
  - Use a private IP address range (e.g., 192.168.1.0/24).
  - Assign static IP addresses to each PC.
  - Configure the router with an appropriate IP address within the same subnet.
- 
Switch Configuration: 
  - Basic configuration to ensure all ports are active and correctly connected.
- 
Router Configuration: 
  - Configure the router with an IP address within the network subnet.
  - Set up the router to provide internet access or simulate an external network connection.
Steps to Complete the Task:
- 
Design the Topology: 
  - Draw a simple diagram with 4 PCs connected to Switch1.
  - Show a connection from Switch1 to Router1.
- 
Assign IP Addresses: 
  - Assign IP addresses to each PC. Example:
    - PC1: 192.168.1.2
    - PC2: 192.168.1.3
    - PC3: 192.168.1.4
    - PC4: 192.168.1.5
  - Assign the switch a management IP if needed (e.g., 192.168.1.1).
- Assign IP addresses to each PC. Example:
- 
Connect Devices: 
  - Physically or virtually (in a network simulator) connect PCs to the switch using Ethernet cables.
  - Connect the switch to the router using an Ethernet cable.
- 
Configure the Router: 
  - Set the router’s internal IP address to 192.168.1.1.
  - Configure the router to route traffic to an external network (e.g., simulate internet access).
- 
Verify Connectivity: 
  - Ping from one PC to another to ensure local network connectivity.
  - Ping the router from each PC to verify they can reach the gateway.
  - If an external network is configured, ping an external IP address to ensure internet connectivity.
Expected Outcomes:
- 
Local Connectivity: 
  - All PCs should be able to communicate with each other.
  - PCs should be able to reach the router.
- 
External Connectivity: 
  - If configured correctly, PCs should be able to reach an external network (e.g., ping 8.8.8.8 for internet).
Deliverables:
- 
Network Diagram: 
  - A visual representation of the 2-tier network topology.
- 
IP Address Plan: 
  - A table listing each device and its assigned IP address.
- 
Configuration Files: 
  - Any configuration settings applied to the router and switch.
- 
Connectivity Test Results: 
  - Ping test results showing successful connectivity between devices.
📁Download Topology : Click Here 🔫
Objective:
Three-tier network architecture is a design model for organizing network infrastructure into three distinct layers, each with specific functions and responsibilities. This segmentation helps improve scalability, performance, manageability, and security.
Problem Statement: Network Design and Configuration
Objective:
Design and configure a network infrastructure for a small-to-medium-sized organization that includes access, distribution, and core layers. The network should support connectivity for four PCs, provide redundancy and fault tolerance, and ensure reliable internet access.
Requirements:
- 
Access Layer: 
  - Connect four PCs (PC-1 to PC-4) to four individual access switches (A-SW-1 to A-SW-4).
  - Each access switch should be connected to two distribution layer switches for redundancy and load balancing.
- 
Distribution Layer: 
  - Implement two distribution switches (D-SW-1 and D-SW-2).
  - Configure EtherChannel with two links between the distribution switches to ensure high bandwidth and redundancy.
  - Each distribution switch should be connected to two access switches.
- 
Core Layer: 
  - Deploy two core switches (C-SW-1 and C-SW-2).
  - Ensure that both core switches are interconnected and each core switch is connected to both distribution switches.
  - Connect the core switches to the internet through a router.
- 
IP Addressing: 
  - Define an IP addressing scheme for each layer, including VLANs for different network segments and devices to ensure proper network segmentation and management.
- 
HSRP Configuration: 
  - Configure Hot Standby Router Protocol (HSRP) on the core switches to provide gateway redundancy.
  - Ensure that one core switch is designated as the active HSRP router and the other as the standby to maintain network availability in case of a failure.
- 
Internet Connectivity: 
  - Verify that the core switches have connectivity to a router that provides internet access.
  - Configure routing protocols as necessary to enable proper internet access from all network devices.
Constraints:
- Ensure that the design is scalable to accommodate potential future expansion.
- The network must provide redundancy and high availability to minimize downtime.
