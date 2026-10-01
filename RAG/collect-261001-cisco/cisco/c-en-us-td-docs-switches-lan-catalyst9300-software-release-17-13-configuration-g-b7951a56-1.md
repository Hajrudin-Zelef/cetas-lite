---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56.md
source_anchor: ""
source_lines: [1, 46]
sha256: e0431de8f86173127d669fcd621769038f512286895d59f12e86db8cf76fa52a
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-17-13-configuration-g-b7951a56

Prerequisites for Port Security
If you try to set the maximum value to a number less than the number of secure addresses already configured on an interface, the command is rejected.
The maximum number of secure MAC addresses that you can configure on a switch is set by the maximum number of available MAC addresses allowed in the system. This number is the total of available MAC addresses, including those used for other Layer 2 functions and any other secure MAC addresses configured on interfaces.
Port Security is not supported on EtherChanel interfaces.
Port Security is not supported on private VLAN ports.
We recommend that you do not enable port security on an 802.1X authenticator interface.
When port-security is disabled on a port, the 802.1X sessions on the port get removed, because the aging timer and inactivity type is still configured. To ensure that the 802.1X sessions are not removed, when disabling port-security, disable the aging timer and inactivity type by removing the following commands:
switchport port-security aging time 1
switchport port-security aging type inactivity
If the inactivity timer is required, see the section "Enabling and Configuring Port Security Aging".
You can use the port security feature to restrict input to an interface by limiting and identifying MAC addresses of the stations allowed to access the port. When you assign secure MAC addresses to a secure port, the port does not forward packets with source addresses outside the group of defined addresses. If you limit the number of secure MAC addresses to one and assign a single secure MAC address, the workstation attached to that port is assured the full bandwidth of the port.
If a port is configured as a secure port and the maximum number of secure MAC addresses is reached, when the MAC address of a station attempting to access the port is different from any of the identified secure MAC addresses, a security violation occurs. Also, if a station with a secure MAC address configured or learned on one secure port attempts to access another secure port, a violation is flagged.
The switch cannot physically detect the link state for devices connected through a Cisco IP phone. When a port disconnects from the switch, the system immediately removes the dynamic secure MAC addresses, even if the port-security aging inactivity is configured to retain those addresses.
The switch supports these types of secure MAC addresses:
Static secure MAC addresses—These are manually configured by using the switchport port-security mac-address mac-address interface configuration command, stored in the address table, and added to the switch running configuration.
Dynamic secure MAC addresses—These are dynamically configured, stored only in the address table, and removed when the switch restarts.
Sticky secure MAC addresses—These can be dynamically learned or manually configured, stored in the address table, and added to the running configuration. If these addresses are saved in the configuration file, when the switch restarts, the interface does not need to dynamically reconfigure them.
The following table shows the default settings for the MAC address table.
| Table 1.  Default Settings for the MAC Address |  | 
|---|---|
| Feature | Default Setting | 
|---|---|
| Aging time | 300 seconds | 
| Dynamic addresses | Automatically learned | 
| Static addresses | None configured | 
With multiple MAC addresses supported on all ports, you can connect any port on the device to other network devices. The device provides dynamic addressing by learning the source address of packets it receives on each port and adding the address and its associated port number to the address table. As devices are added or removed from the network, the device updates the address table, adding new dynamic addresses and aging out those that are not in use.
The aging interval is globally configured. However, the device maintains an address table for each VLAN, and STP can accelerate the aging interval on a per-VLAN basis.
The device sends packets between any combination of ports, based on the destination address of the received packet. Using the MAC address table, the device forwards the packet only to the port associated with the destination address. If the destination address is on the port that sent the packet, the packet is filtered and not forwarded. The device always uses the store-and-forward method: complete packets are stored and checked for errors before transmission.
You can configure an interface to convert the dynamic MAC addresses to sticky secure MAC addresses and to add them to the running configuration by enabling sticky learning. The interface converts all the dynamic secure MAC addresses, including those that were dynamically learned before sticky learning was enabled, to sticky secure MAC addresses. All sticky secure MAC addresses are added to the running configuration.
The sticky secure MAC addresses do not automatically become part of the configuration file, which is the startup configuration used each time the switch restarts. If you save the sticky secure MAC addresses in the configuration file, when the switch restarts, the interface does not need to relearn these addresses. If you do not save the sticky secure addresses, they are lost.
If sticky learning is disabled, the sticky secure MAC addresses are converted to dynamic secure addresses and are removed from the running configuration.
It is a security violation when one of these situations occurs:
The maximum number of secure MAC addresses have been added to the address table, and a station whose MAC address is not in the address table attempts to access the interface.
An address learned or configured on one secure interface is seen on another secure interface in the same VLAN.
Running diagnostic tests with port security enabled.
You can configure the interface for one of three violation modes, based on the action to be taken if a violation occurs:
protect—when the number of secure MAC addresses reaches the maximum limit allowed on the port, packets with unknown source addresses are dropped until you remove a sufficient number of secure MAC addresses to drop below the maximum value or increase the number of maximum allowable addresses. You are not notified that a security violation has occurred.
| Note | We do not recommend configuring the protect violation mode on a trunk port. The protect mode disables learning when any VLAN reaches its maximum limit, even if the port has not reached its maximum limit. | 
restrict—when the number of secure MAC addresses reaches the maximum limit allowed on the port, packets with unknown source addresses are dropped until you remove a sufficient number of secure MAC addresses to drop below the maximum value or increase the number of maximum allowable addresses. In this mode, you are notified that a security violation has occurred. An SNMP trap is sent, a syslog message is logged, and the violation counter increments.
shutdown—a port security violation causes the interface to become error-disabled and to shut down immediately, and the port LED turns off. When a secure port is in the error-disabled state, you can bring it out of this state by entering the errdisable recovery cause psecure-violation global configuration command, or you can manually re-enable it by entering the shutdown and no shut down interface configuration commands. This is the default mode.
shutdown vlan—Use to set the security violation mode per-VLAN. In this mode, the VLAN is error disabled instead of the entire port when a violation occurs
This table shows the violation mode and the actions taken when you configure an interface for port security.
| Table 2. Security Violation Mode Actions |  |  |  |  |  |  | 
|---|---|---|---|---|---|---|
| Violation Mode | Traffic is forwarded 1 | Sends SNMP trap | Sends syslog message | Displays error message 2 | Violation counter increments | Shuts down port | 
|---|---|---|---|---|---|---|
