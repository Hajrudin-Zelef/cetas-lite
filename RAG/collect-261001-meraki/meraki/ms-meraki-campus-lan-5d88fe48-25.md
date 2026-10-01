---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-25
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [1186, 1232]
sha256: 04bae2fc8af2a157a269b83a9fbfb5bd938a1ba714072e0676743ce1b5639ba8
---

# ms-meraki-campus-lan-5d88fe48

| MS Switch Family | MS Switch Model | Minimum Firmware Required | 
| MS200 series | MS210 | MS 14.15 | 
|  | MS225 | MS 14.15 | 
|  | MS250 | MS 14.15 | 
| MS300 series | MS350 | MS 14.15 | 
|  | MS355 | MS 14.15 | 
| MS400 series | MS410 | MS 14.15 | 
|  | MS425 | MS 14.15 | 
|  | MS450 | MS 14.15 | 
- For more information on the supported AP models and firmware, please refer to the following guide. Details provided in the below table:
| MR Family | MR models | Minimum Firmware Required | 
| WiFi-5 Wave 2 (802.11ac Wave 2) | MR20, MR30H, MR33, MR42, MR42E, MR52, MR53, MR53E, MR70, MR74, MR84 | MR27.6 | 
| Wi-Fi 6 (802.11ax) | MR45, MR55, MR36, MR46, MR46E, MR56, MR76, MR86 | MR26.7 | 
Some MR44s, and MR46s are not yet supported by SecureConnect on firmware versions MS 14.18 and older. Please contact Meraki support to check for compatibility if you have any of these two models in your network.
- The management VLAN used by SecureConnect when configuring a port connected to an MR is the VLAN being used by the switch as its management VLAN at the time. This VLAN may differ from the user-configured management VLAN because, when unable to obtain an IP in the configured management VLAN, an MS switch will try to use the other VLANs for management connectivity.
- SecureConnect does not apply to LACP aggregate group ports. If an MR access-point that does not support LACP is plugged into a switchport which is part of an LACP aggregate group, the switchport will be disabled by LACP
- MR access-points that do support LACP, when plugged into a switchport configured as a part of an LACP aggregate group will continue to function as they would if SecureConnect was disabled.
- Supported APs will start off only being able to reach dashboard on the switch management VLAN. The APs will have 3 attempts of 5 seconds each to authenticate If this authentication fails, the switch's port will fall into a restricted state. (e.g. Wireless clients connected are unable to browse, OR switchport shows that SecureConnect has failed)
SecureConnect can fail if the AP and switch are in different organizations, or if the AP is not claimed in inventory
- The following table provides details of the behaviour and the port configuration associated with the different SecureConnect swtichport states:
| State | State details | Port configuration | 
| Disabled | SecureConnect is not enabled in the network | Switchport retains the last user-defined configuration settings. | 
| Enabled | SecureConnect is enabled in the network but the switchport is not connected to a SecureConnect capable MR access-point. | Switchport retains the last user-defined configuration settings. | 
| In Progress | A SecureConnect MR access-point is connected to the switchport but it has not yet completed the authentication process. While the switchport is in this state, the MR communicates with the Dashboard to download the required security certificates along with any user-defined configuration, and attempts to authenticate itself. If it is the first time that the connected MR has being plugged into a SecureConnect enabled switchport since it was claimed in the Dashboard Organization, the port may remain in this state for an extended period as the MR is issued the security certificate. | SecureConnect enforced switchport configuration: Type : Trunk Native VLAN : Switch Management VLAN Allowed VLANs : Switch Management VLAN only Access Policy : Not applicable (SecureConnect) Traffic restrictions to allow only communication between the MR and the Meraki Dashboard. The remaining user-defined switchport settings are retained. | 
| Authenticated | The MR has been successfully authenticated via Meraki Auth, using the MR’s security certificate, and has been verified to belong to the same Dashboard Organization as the switch. | SecureConnect enforced switchport configuration: Type : Trunk Native VLAN : Switch Management VLAN Allowed VLANs : All VLANs Access Policy : Not applicable (SecureConnect) The remaining user-defined switchport settings are retained. | 
| Restricted | The MR has either failed to authenticate or the authentication process resulted in a timeout. | SecureConnect enforced switchport configuration: Type : Trunk Native VLAN : Switch Management VLAN Allowed VLANs : Switch Management VLAN only Access Policy : Not applicable (SecureConnect) Traffic restrictions to allow only communication between the MR and the Meraki Dashboard. The remaining user-defined switchport settings are retained. | 
- SecureConnect-capable MR access point connected to an MS switch enabled for SecureConnect should not be configured with LAN IP VLAN number. While the other LAN IP settings can be configured, the VLAN field should be left blank (as shown below)
MS390 Specific Guidance
- SecureConnect is not yet supported on MS390 platforms
Please refer to the firmware changelog for guidance on when this feature will be introduced for MS390 platforms
Pre-requisites for Secure Connect: (In addition to above MS and MR platform notices)
- The MR access-point and the MS switch should be directly connected to support SecureConnect
- The switchport on which the MR is connected should be enabled
- The switchport must be configured for PoE if the MR is not using a power injector
Please note that the management VLAN used by SecureConnect when configuring a port connected to an MR is the VLAN being used by the switch as its management VLAN at the time. This VLAN may differ from the user-configured management VLAN because, when unable to obtain an IP in the configured management VLAN, an MS switch will try to use the other VLANs for management connectivity.
Multi Dwelling Units (MDUs)
In some situations, such as for IoT and for multi-dwelling unit (MDU) deployments, the access layer is often augmented with additional cascaded switches. For MDU deployments the devices may be small distributed access switches that are hanging off your access layer (i.e. daisy-chained) or even as an extension of the access layer itself. It is therefore important to remember that these switches will be an extension of your layer 2 domain and therefore your STP domain. The recommended design for these switches depends on the use cases implemented and the downstream devices requiring connectivity as this will also instigate specific port settings such as port security and port mode.
General Guidelines
- It is recommended to deploy MDUs as close as possible to the downstream devices
- It is recommended to avoid inter-connecting between the MDU units using direct links (i.e. They should communicate via the upstream access switch)
- In the event that you have to interconnect between the MDU switches, please remember to configure STP Loop Guard and UDLD on both sides of the inter-connecting link(s)
- It is recommended (where possible) to use multiple uplinks grouped in an Ether-Channel to multiple switches in your access stack
- It is recommended to deploy smaller MDU switches serving a single VLAN rather than large switches serving multiple VLANs as this will simplify the troubleshooting process
- Typically, these MDU switches will require access to DHCP in VLAN 1 for ZTP (Unless configured manually)
- It is recommended to configure the downstream port connecting an MDU switch in access mode (unless the MDU switch requires a management IP in the designated management VLAN)
- Ensure that you configure STP Root Guard on downstream ports connecting the MDU switches
The reason you want to apply STP Root Guard as opposed to STP BPDU guard is that it is more likely that the MDU switches will be sending BPDUs (e.g. if you turn on STP) which in case of having STP BPDU Guard will cause the port to shutdown rather than the likelihood of the MDU switch being configured with a wrong Bridge Priority which in case of having STP Root Guard will cause the port to go in ErrDisabled state.
