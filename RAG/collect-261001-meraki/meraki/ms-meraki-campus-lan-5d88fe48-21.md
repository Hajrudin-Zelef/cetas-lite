---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-21
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [934, 1009]
sha256: 2c62272019d5f80403d4e97782a1f7a4e4eab95e0083efc463be5eeea6b6a74f
---

# ms-meraki-campus-lan-5d88fe48

- It is recommended to mark your traffic as close as possible to the source. So, have your traffic marked at the SSID level using the MR traffic shaping feature. Marked traffic can be trusted on the MS platforms and will be policed based DSCP to CoS mappings mentioned below
- Configuring QoS on your Meraki switches is done at the Network level which means that it automatically applies to all of the switches in the Meraki Network
- QoS rules can be defined based on VLAN, Source port (or range), destination port (or range)
MS120 and MS125 series switches support QoS rules based on VLANs only. Port-range based rules are not supported and will be not be applied and dashboard will display an error
- An MS network has 6 configurable CoS queues labeled 0-5. Each queue is serviced using FIFO. Without QoS enabled, all traffic is serviced in queue 0 (default class) using a FIFO model. The queues are weighted as follows:
| CoS | Weight | 
| 0 (default class) | 1 | 
| 1 | 2 | 
| 2 | 4 | 
| 3 | 8 | 
| 4 | 16 | 
| 5 | 32 | 
To translate the above weights to bandwidth allocations, please refer to the following table:
| Priority | CoS | Weight (N) | Bandwidth Allocation (NTotal = Sum of weights all configured classes on dashboard) | Min BW* | Max BW* | 
| Lowest | 0 (default class) | 1 | = (1 / NTotal )% | ~ 1.5% | 100% | 
|  | 1 | 2 | = (2 / NTotal )% | ~ 3% | ~ 67% | 
|  | 2 | 4 | = (4 / NTotal )% | ~ 7% | ~ 80% | 
|  | 3 | 8 | = (8 / NTotal )% | ~ 12.5% | ~ 89% | 
|  | 4 | 16 | = (16 / NTotal )% | ~ 25% | ~ 94% | 
| Highest | 5 | 32 | = (32 / NTotal )% | ~ 50% | ~ 97% | 
* Assumption for the values above is that you will always keep the Default class (CoS value 0). Hence, the Max BW represents 2 queues. And the Min BW represents all 6 in use.
Please refer to this example to calculate the bandwidth allocated in a certain queue.
Traffic will be assigned to the default FIFO queue if one of the following is true:
- QoS is not enabled on switch
- No match to the DSCP value
- Match DSCP value which is mapped to CoS value 0
- You can edit the default DSCP to CoS mapping as well (See below)
- If you do not specify a mapping for DSCP value to CoS, the default CoS value assigned will be 0
Please note that as soon as the first QoS rule is added, the switch will begin to trust DSCP bits on incoming packets that have a DSCP to CoS mappings. This rule is invisible and processed last.
However if an incoming packet has a DSCP tag set but no matching QoS rule or DSCP to CoS mapping, it will be placed in the default queue.
MS390 Specific Guidance
- Same guidance as above
Access Policy
General Guidance
- Use Access Policy on MS platforms to authenticate devices against a Radius server
- These access policies are typically applied to ports on access-layer switches
- As of MS 9.16, changes to an existing access policy will cause a port-bounce on all ports configured for that policy
- Use Single-host mode (default) on switchports with only one client attached (if multiple devices are connected, only the first client will be allowed network access upon successful authentication)
- Use Multi-domain mode to authenticate one device in each of the data and voice VLANs. This mode is recommended for switchports connected to a phone with a device behind the phone (Authentication is independent on each VLAN and will not affect the forwarding state of each other)
MS Switches require the Cisco-AVPair: device-traffic-class=voice pairs within the Access-Accept frame to put devices on the voice VLAN
- Use Multi-Auth mode to authenticate each device connected (All hosts attached must have matching VLAN information or will be denied access, with only. one device supported in voice VLAN)
- Use Multi-Host mode to authenticate the first device connected and subsequently allowing (i.e. ignoring authentication) for all other hosts that will be granted access without authentication. This is recommended in deployments where the authenticated device acts as a point of access to the network, for example, hubs and access points
- With 802.1x the client will be prompted to provide their domain credentials which are authenticated against a Radius server (If no Access-Request is presented, the device will be placed in the Guest VLAN if defined)
- With MAC Authentication Bypass (MAB) the client's MAC address is authenticated against a Radius server (no user prompt). It is typically used to offer seamless user experience restricting the network to specific devices without having to prompt the user
- With Hybrid Authentication the client will first be prompted to provide credentials for 802.1x authentication. If that fails (e.g. no EAP received within 8 seconds) then the switch will use the client's MAC address and will be authenticated via MAB (if both methods fail, the device will be placed in Guest VLAN if defined). It is recommended to use Hybrid if not every device supports 802.1x since MAB can be used as a failover method.
Radius Attributes and Features
General Guidance
- When an access policy is configured with RADIUS server, authentication is performed using PAP. The following attributes are present in the Access-Request messages sent from MS switch to the RADIUS server:
    
  - User-Name
  - NAS-IP-Address
  - Calling-Station-Id: Contains the MAC address of the Meraki MS switch (all caps, octets separated by hyphens). Example: "AA-BB-CC-DD-EE-FF".
  - Called-Station-Id: Contains the MAC address of the Meraki MS switch (all caps, octets separated by hyphens).
  - Framed-MTU
  - NAS-Port-Type
  - EAP-Message
  - Message-Authenticator
- RADIUS traffic will always be sourced from the Management IP of the MS (even if the RADIUS Server is reachable via a configured SVI and in this instance, the RADIUS traffic would first be sent to the default gateway associated with the Management IP, which would then forward this traffic back down towards the switch to reach the RADIUS server)
- When using PEAP EAP-MSCHAPv2 on an MS switchport, if an unmanaged switch is between the supplicant (user machine) and the RADIUS client (MS) the authentication will fail. (It is possible to circumvent this by using MAC based RADIUS authentication. If one machine authenticates via MAC based RADIUS through the MS on an unmanaged switch, the machine that has authenticated will be granted access. It is a workaround and it is less secure and requires more configuration on the NPS and DC)
- Meraki MS switches support CoA for RADIUS re-authentication and disconnection as well as port bouncing (UDP/1700 is the default port used by all MS for CoA with Cisco ISE and port 3799 for many other vendors)
The CoA Request frame is a RADIUS code 43 frame. Cisco Meraki switches require that all the following attribute pairs within this frame:
- Calling-Station-ID
- Cisco-AV-Pair
    
  - subscriber:command=reauthenticate
  - audit-session-id (The Cisco audit-session-id custom AVPair is used to identify the current client session that CoA is destined for. Meraki switches learn the session ID from the original RADIUS access accept message that begins the client session)
Please see the following CoA frame as an example:
The Disconnect Request frame is a RADIUS code 40 frame. The Cisco Meraki switch will utilize the following attribute pairs within this frame:
- Cisco-AV-Pair
    
  - audit-session-id (The Cisco audit-session-id custom AVPair is used to identify the current client session that CoA is destined for. Meraki switches learn the session ID from the original RADIUS access accept message that begins the client session)
- Calling-Station-Id
Please see the following Disconnect Request frame as an example:
The Port Bounce request is a RADIUS code 43 request. The Cisco Meraki switch will utilize the following attribute pairs within this frame:
- Cisco-AV-Pair
    
