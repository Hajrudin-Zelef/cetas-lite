---
id: collect-261001-automatisation-infra/automatisation-infra/github-georgybi-eve-ng-labs-welcome-to-the-eve-ng-labs-catalog-this-repository-is-a-compre-3
title: "github-georgybi-eve-ng-labs-welcome-to-the-eve-ng-labs-catalog-this-repository-is-a-comprehensive-co"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "mit license"]
source: docs/RAG/collect-261001-automatisation-infra/github-georgybi-eve-ng-labs-welcome-to-the-eve-ng-labs-catalog-this-repository-is-a-comprehensive-co.md
source_anchor: ""
source_lines: [318, 464]
sha256: 0711a63a006d7c6e9fbc9529d7aa10332822c107f640f040e2f6a9f8fced40b9
---

# github-georgybi-eve-ng-labs-welcome-to-the-eve-ng-labs-catalog-this-repository-is-a-comprehensive-co

- **Router (Edge-R):** Equipped with two subinterfaces on its g0/0 interface:
  - Subinterface for VLAN 10 (IT) with IP address 10.1.1.100
  - Subinterface for VLAN 20 (HR) with IP address 20.1.1.100
- **Switch:** Configured with two VLANs:
  - VLAN 10 for the IT department
  - VLAN 20 for the HR department
- **PCs:**  - **PC-1 and PC-2:** Located in VLAN 10 (IT) with IP addresses in the 10.1.1.0/24 subnet
  - **PC-3 and PC-4:** Located in VLAN 20 (HR) with IP addresses in the 20.1.1.0/24 subnet

**Objectives:**

1. 
**VLAN Configuration:**
  - Configure VLAN 10 and VLAN 20 on the switch.
  - Assign IP subnet addresses to VLANs, ensuring VLAN 10 uses the 10.1.1.0/24 subnet and VLAN 20 uses the 20.1.1.0/24 subnet.
2. 
**Router Configuration:**
  - Establish a trunk link between the switch and the router.
  - Configure the router with subinterfaces for VLAN 10 and VLAN 20 to enable routing between the VLANs.
  - Assign the IP addresses 10.1.1.100 and 20.1.1.100 to the router’s subinterfaces.
3. 
**PC Configuration:**
  - Assign appropriate IP addresses and subnet masks to PCs within each VLAN:
    - PC-1 and PC-2 should be configured with IP addresses in the 10.1.1.0/24 subnet.
    - PC-3 and PC-4 should be configured with IP addresses in the 20.1.1.0/24 subnet.
4. Assign appropriate IP addresses and subnet masks to PCs within each VLAN:
5. 
**Connectivity Verification:**
  - Verify that PCs within the same VLAN can communicate with each other.
  - Test connectivity between PCs in different VLANs to ensure successful inter-VLAN routing.
  - Ensure that devices are able to communicate across VLANs through the router.
6. 
**Documentation:**
  - Document the configuration steps for the switch, router, and PCs.
  - Provide a comprehensive guide that includes network diagrams, configuration commands, and troubleshooting steps to support future maintenance and problem resolution.

**Deliverables:**

- A functional inter-VLAN routing setup demonstrating communication between devices in different VLANs.
- Detailed configuration documentation for switch VLANs, router subinterfaces, and PC IP settings.
- Verification results showing successful communication within and between VLANs.

**Constraints:**

- Ensure that the configuration maintains network segmentation and security while enabling required communication.
- The design should be scalable and easily adjustable for additional VLANs or changes in network topology.

**Objective:**

CDP is an essential protocol for Cisco network devices, simplifying network management and troubleshooting by providing a straightforward way to discover and display information about directly connected Cisco equipment. Proper configuration and usage of CDP can greatly enhance network visibility and operational efficiency.

You are tasked with configuring and verifying Cisco Discovery Protocol (CDP) on a network consisting of multiple Cisco devices. The network diagram shows the connectivity between various routers and switches. Your objective is to ensure that CDP is properly configured and operational across all devices to facilitate network discovery and troubleshooting.

1. 
**Enable CDP Globally on All Devices** :
  - Verify if CDP is enabled globally on each device.
  - If CDP is not enabled, enable it globally.
2. 
**Enable CDP on All Relevant Interfaces** :
  - Identify all the interfaces that connect to other Cisco devices.
  - Ensure CDP is enabled on these interfaces.
3. 
**Verify CDP Neighbor Information** :
  - Check the CDP neighbor table on each device.
  - Confirm that each device can see its directly connected neighbors.
4. 
**Collect and Document Neighbor Information** :
  - Document the details of each neighbor as shown in the CDP neighbor table.
  - Include information such as device name, local interface, neighbor interface, and capabilities.
5. 
**Troubleshoot CDP Issues** :
  - If any device does not display its neighbors correctly, verify the physical connectivity.
  - Ensure that CDP is not disabled on any necessary interfaces.
  - Check for any potential issues that might be blocking CDP packets (e.g., VLAN configuration, interface errors).
6. 
**Validate CDP Information Accuracy** :
  - Cross-check the discovered CDP information with the physical network diagram.
  - Ensure that all connections match the expected topology.
7. 
**Maintain CDP Configuration** :
  - Implement best practices for CDP configuration, such as setting the CDP timer and holdtime appropriately.
  - Regularly review and update the CDP neighbor information to reflect any network changes.
8. 
**Security Considerations** :
  - Evaluate the security implications of CDP in your network.
  - If necessary, disable CDP on interfaces connected to untrusted networks or devices.

**Objective:**

LLDP is a vendor-neutral protocol used by network devices to advertise their identity, capabilities, and neighbors on a local area network (LAN).

Use the Link Layer Discovery Protocol (LLDP) to discover and verify the network topology.

1. Enable LLDP on all network devices (routers, switches, IP phones, and VPC).
2. Verify connectivity and discover neighboring devices using LLDP.

- 
**Switch**
  - Connects to:
    - **VPC4** on port`eth0`
    - **IP_Phone-1** on port`eth0`
    - **Router-1** on port`Gi0/2`
- Connects to:
- 
**Router-1**
  - Connects to:
    - **Switch** on port`Gi0/2`
    - **IP_Phone-2** on port`eth0`
    - **Router-2** on port`Gi0/1`
- Connects to:
- 
**Router-2**
  - Connects to:
    - **Router-1** on port`Gi0/0`
- Connects to:

1. Enable LLDP globally.
2. Enable LLDP on interfaces `Gi0/0` ,`Gi0/1` , and`Gi0/2` .

1. Enable LLDP globally.
2. Enable LLDP on interface `Gi0/0` .

- Ensure LLDP is enabled if necessary through the phone's configuration interface (most IP phones support LLDP automatically).

- Since VPC typically does not support LLDP natively, verify the connection through the switch’s LLDP neighbor table.

1. Verify LLDP neighbors on the switch.
2. Verify LLDP neighbors on Router-1.
3. Verify LLDP neighbors on Router-2.
4. Verify LLDP neighbors on IP Phones through their settings or documentation.

By performing these steps, ensure all devices are properly discovered and that the topology is correctly identified according to the LLDP advertisements. This process assists in network troubleshooting, documentation, and ensuring proper connectivity across the network.

- Configure the VTP modes as per the provided diagram
- assign the modes as mentioned.
- Test the behavior of each mode, specifically checking the Configuration Revision (CR) number for transparent mode.
- Finally, manually configure all VLAN settings on the switch in transparent mode.

We welcome contributions! If you have a lab you’d like to add or improvements to suggest, please submit a pull request.

If you encounter any issues, please open an issue in this repository. We will address it as soon as possible.

This project is licensed under the MIT License. See the LICENSE file for more details.

```
Make sure to replace `https://github.com/yourusername/eve-ng-labs-catalog.git` with the actual URL of your repository and update the paths to your screenshots accordingly. This will make your repository more attractive and user-friendly on GitHub.
```
