---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-23
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "voice"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [1034, 1104]
sha256: 48ceb818bb678612ddc0e8ff58e63a7e4a5224181345b2177ef8024163402590
---

# ms-meraki-campus-lan-5d88fe48

  - Tunnel -Medium-Type: Choose 802 (Includes all 802 media plus Ethernet canonical format) for the Attribute value Commonly used for 802.1X
  - Tunnel-Private-Group-ID: Choose String and enter the VLAN desired (ex. "500"). This string will specify the VLAN ID 500.
  - Tunnel-Type: Choose Attribute value Commonly used for 802.1X and select Virtual LANs (VLANs).
* Dynamic VLAN Assignment is not supported on the voice VLAN/domain
- Guest VLANs can be used to allow unauthorized devices access to limited network resources
Guest VLANs is not supported on the voice VLAN/domain
- With Failed Authentication VLAN, A client device connecting to a switchport controlled by an access-policy can be placed in the failed authentication VLAN if the RADIUS server denies its access request (e.g. non-compliance with network security requirements)
Failed Authentication VLAN is only supported in the Single Host, Multi Host and Multi Domain modes
Access policies using Multi Auth mode are not supported.
- When the Re-authentication Interval (time in seconds) is specified, the switch will periodically attempt authentication for clients connected to switchports with access policies. This is recommended to provide a better security policy by periodically validating client authentication in a network, but also the re-authentication timer enables the recovery of clients placed in the Failed Authentication because of incomplete provisioning of credentials.
- 
    With Suspend Re-authentication when RADIUS servers are unreachable, Periodic re-authentication of clients can be an issue when RADIUS servers are unreachable. The Suspend Re-authentication when RADIUS servers are unreachable disables the re-authentication process when none of the RADIUS servers are reachable.
Suspend re-authentication when RADIUS servers are unreachable,' is not a configurable option on the MS390 series switches. An MS390 switch will automatically ignore this config, and will always suspend client re-authentication, if it loses connectivity with the RADIUS server
- With Critical Authentication VLAN, it can be used to provide network connectivity to client devices connecting on switchports controlled by an access-policy when all the RADIUS servers for that policy are unreachable or fail to respond to the authentication request on time. (i.e. Critical authentication VLAN ensures that these clients are still able access the business-critical resources, by placing them in separate VLAN). This also allows network administrators better control the network access available to clients when their identities cannot be established using RADIUS.
The critical data and critical voice VLANs should not be the same
Configuring Critical Authentication VLAN or Failed Authentication VLAN under an access policy may affect its existing Guest VLAN behavior. Please consult the Interoperability and backward compatibility section of this document for details.
- 
    With Suspend port bounce, You can STOP bouncing the clients placed in the Critical Authentication VLAN when any of the Radius servers are restored (Determined by the Radius Testing process). The switch does this by bouncing (turning off and on) the switchports on which these clients are connected. If required, this port-bounce action can be stopped by enabling the Suspend port bounce option (i.e. the clients will be retained in the Critical Authentication VLAN until a re-authentication for these clients is manually triggered)
MS 14 is the minimum firmware version required for the following configuration options:
- Failed Authentication VLAN
- Re-authentication Interval,
- Suspend Re-authentication when RADIUS servers are unreachable
- Critical Authentication VLANs
- Suspend port bounce
MS390 Special Guidance
- MS390s support RADIUS CoA & URL-Redirect as of MS15
Interoperability and Backward Compatibility
If Critical and/or Failed Authentication VLANs are specified in an Access Policy, the Guest VLAN functionality gets modified to ensure backward-compatibility and inter-op between the configured VLANs. Please refer to the Interoperability and backward-compatibility table below for more details on this.
The following matrix shows the remediation VLAN, in any, that client device would be placed in for the different combinations of the remediation VLAN configuration options and the RADIUS authentication result.
| Configured options | Authentication result |  |  | 
|  | EAP timeout (for 802.1X policies only) | RADIUS timeout (server unreachable) | Authentication Fail (access-reject) | 
| Guest (existing behavior) | Guest VLAN | Guest VLAN | Access denied 1 | 
| Failed | Access denied | Access denied | Failed Auth VLAN | 
| Critical | Access denied | Critical Auth VLAN | Access denied | 
| Guest and Failed | Guest VLAN | Guest VLAN | Failed Auth VLAN | 
| Guest and Critical | Guest VLAN | Critical Auth VLAN | Access denied 1 | 
| Critical and Failed | Access denied | Critical Auth VLAN | Failed Auth VLAN | 
| Guest, Failed and Critical | Guest VLAN | Critical Auth VLAN | Failed Auth VLAN | 
1 When using hybrid authentication without increase access speed (concurrent-auth), a client failing both 802.1X and MAB authentication will also be placed in the Guest VLAN
Cisco ISE Integration Guidance
- Meraki MS platforms can integrate with Cisco ISE for authentication and posture
- It is important to understand the compatability when integrating with Cisco ISE. Please refer to the below table:
| Feature | Integration Status | Notes | 
|---|---|---|
| Profiling | Full |  | 
| BYOD | Full |  | 
| Posture | Full |  | 
| AAA | Partial | dACL is not supported, Use Group Policy ACL instead (beta) | 
| Guest | Partial | Local Web Authentication not supported | 
| TrustSec | Partial | Adaptive Policy with MS390s. Sync requires a docker container. Also please refer to this guide for Hybrid Campus LAN with Adaptive Policy | 
| MDM | Partial | Only with Meraki Systems Manager | 
| Guest Originating URL | Not Supported |  | 
Named VLAN Profiles
General Guidance
Named VLAN profiles are currently in closed beta testing. Please reach out to Meraki support to have it enabled.
- Named VLAN Profiles work along with 802.1X RADIUS authentication to assign authenticated users and devices to specific VLANs according to a VLAN name rather than an integer number (e.g. Use case of having multiple sites with different VLAN ID numbers for same functional group of users and devices)
Named Profiles Scaling Considerations:
Each profile can include up to 1024 VLAN name to ID mappings, and each VLAN name can be up to 32 characters long. The VLAN profile name itself has a 255 character limit.
You can also map more than one VLAN ID number to a VLAN name using commas or hyphens to separate non-contiguous and contiguous ranges (e.g. 100,200,120-130)
- In order to use named VLAN profiles, an access policy must be first configured and assigned to switchports to authenticate users and devices connecting to those ports
The RADIUS server must be configured to send three attributes to the switch as part of the RADIUS Access-Accept message sent to the switch as a result of a successful 802.1X authentication. These attributes tell the switch which VLAN name to assign to the session for that user or device. The required attributes are:
- [64] Tunnel-type = VLAN
- [65] Tunnel-Medium-Type = 802
- [81] Tunnel-Private-Group-ID = <vlan name>
- 
    If the RADIUS server returns a name value that is not defined in the VLAN profiles, the switchport will fail-closed and the client device will not be able to access the network
- 
    When using multi-auth mode, make sure to have matching VLAN information (i.e. same VLAN) for all subsequent hosts or they will be denied access to the port
- 
    Please make sure to enable Named VLAN Profiles for each network otherwise settings will not take effect
- 
