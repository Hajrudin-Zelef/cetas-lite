---
id: collect-261001-fortinet/fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-se-d297db14-2
title: "Configure administrator timeout"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/shayan-heydarikhah-sheynshield-blob-head-fortinet-fortigate-cheatsheet-system-se-d297db14.md
source_anchor: ""
source_lines: [249, 573]
sha256: e78c78c6459ab0500f1f4a23eae21a6a9f644aa6d1e659d280b9d6d0acd7b7c0
---

# Configure administrator timeout

FortiOS supports up to 10 failed attempts, while the lockout duration is configured in seconds.
GUI:
Log & Report
    >
System Events
    >
General System Events
Useful for identifying:
- Failed logins
- Successful logins
- Lockouts
- Administrative activity
- Authentication anomalies
Auxiliary Session allows FortiGate to create additional session state when traffic associated with an existing connection uses a different incoming or outgoing interface/path.
It is particularly relevant to:
- ECMP
- Asymmetric routing
- Multiple WAN links
- Load balancing
- Policy-based routing
- SD-WAN
- ADVPN
- Complex multi-path topologies
Fortinet describes auxiliary sessions as a mechanism for handling changes in incoming, outgoing, or return interfaces in environments such as ECMP and load balancing.
| FortiOS | Auxiliary Session | 
|---|---|
| 6.0 and earlier | Not supported | 
| 6.2.0–6.2.2 | Permanently enabled | 
| 6.2.3+ | Disabled by default; can be enabled | 
config system settings
    set auxiliary-session enable
end
Disable:
config system settings
    set auxiliary-session disable
end
FortiOS 7.2 CLI documents the default as disabled.
IPv4:
config system settings
    set asymroute enable
end
IPv6:
config system settings
    set asymroute6 enable
end
ICMP-specific controls also exist:
set asymroute-icmp enable
set asymroute6-icmp enable
FortiOS documents these as separate controls for IPv4, IPv6 and ICMP asymmetric routing.
These are related but not identical.
Asymmetric Routing
        |
        +---- Traffic uses different paths
        |
        v
Auxiliary Session
        |
        +---- Helps FortiGate maintain session handling
              when interfaces/paths change
asymroute controls acceptance/handling of asymmetric routing.
auxiliary-session controls creation/use of auxiliary sessions for changing traffic paths.
             ISP-1
               |
               v
Client ---- FortiGate ---- Server
               ^
               |
             ISP-2
Suppose:
Forward path  = ISP-1
Return path   = ISP-2
The original session was created through one interface, but return traffic arrives through another.
Without auxiliary sessions, FortiGate may need to modify/refresh the session state.
With auxiliary sessions:
Main Session
      |
      +---- Auxiliary Session
                   |
                   v
             Alternate Path
This can allow traffic to continue while preserving appropriate session handling.
A simplified FortiGate packet-processing model:
             Incoming Packet
                    |
                    v
                  CPU
                    |
             Session Lookup
                    |
             +------+------+
             |             |
           New           Existing
             |             |
             v             v
       Create Session   Session Check
             |             |
             v             v
       Policy Lookup    NPU Eligibility
             |             |
             v             v
        Routing/PBR       NPU
                           |
                           v
                    Fast Forwarding
The CPU is responsible for control-plane/session-establishment work such as:
- Initial packet processing
- Session lookup
- New session creation
- Policy decisions
- Routing decisions
- Session-state management
- Control-plane functions
On supported FortiGate platforms, the NPU can offload eligible data-plane processing.
Typical benefit:
CPU
 |
 | session establishment
 v
NPU
 |
 | high-speed forwarding
 v
Output Interface
NPU offloading does not mean the CPU disappears from the session.
The CPU generally performs the initial/session-control work, while the NPU can handle subsequent eligible packet forwarding.
When traffic changes interface/path:
Existing Session
      |
      v
Interface Changed
      |
      v
Session becomes DIRTY
      |
      v
Session/interface update
      |
      v
CPU processing
This can become expensive in environments with:
- ECMP
- Interface flapping
- Multiple WAN paths
- Frequent path changes
Existing Session
      |
      v
Path/interface changed
      |
      v
Auxiliary Session
      |
      v
NPU can continue eligible forwarding
This can reduce the need to repeatedly modify the original session.
A session may appear as:
state = dirty
A dirty session indicates that the session state requires additional processing/update.
Typical triggers include:
- Path change
- Incoming interface change
- Return interface change
- Routing changes
- ECMP behavior
- Asymmetric traffic
Example session output:
npu info:
flag=0x91/0x81
offload=8/8
Both directions show offload.
Another example:
npu info:
flag=0x91/0x00
offload=8/0
Only one direction is offloaded.
offload=x/y
x = one direction
y = reverse direction
Always interpret the exact output according to the FortiOS/platform version.
Client
  |
port1
  |
FortiGate
  |
port3
  |
Server
Return:
Server
  |
port3
  |
FortiGate
  |
port1
  |
Client
Return traffic can use the original session/path.
FortiGate evaluates the applicable route/PBR/SD-WAN decision.
If the best path is:
port1
traffic uses port1.
If the best path is:
port2
traffic may use port2.
When both PBR and SD-WAN-related routing decisions are involved, the exact decision process depends on the configuration and FortiOS version.
Do not memorize "PBR always wins" without checking the specific FortiOS routing decision order.
Original:
Session:
port1 ---> FortiGate ---> port3
Return arrives:
port4 ---> FortiGate
The existing session may become:
DIRTY
FortiGate updates the session/interface state.
High traffic or repeated path changes can increase CPU processing.
FortiGate can create:
Original Session
       +
Auxiliary Session
allowing traffic to continue using the alternate path.
Original:
port1 ---> FortiGate ---> port3
New incoming packet:
port2 ---> FortiGate
Existing Session
      |
      v
DIRTY
      |
      v
Session update
Existing Session
      |
      v
Auxiliary Session
      |
      v
Continue forwarding
Original:
Client
  |
port1
  |
FortiGate
  |
port3
  |
Server
A new route appears:
Server -> port4
If the original route remains valid, an established session may continue using the original path rather than immediately moving to the newly preferred route.
A route/interface change can trigger session handling that refreshes the path/interface information.
Auxiliary sessions should not be enabled blindly.
Fortinet specifically warns that routing symmetry should be considered in topologies such as:
- SD-WAN Hub-and-Spoke
- ADVPN
- Other environments where return traffic is expected to be symmetric
Symmetric traffic expected?
        |
        +---- YES ---> Prefer disabling auxiliary-session
        |
        +---- NO ----> Evaluate auxiliary-session
The correct setting depends on the actual topology and traffic behavior.
diagnose sys session list
Useful information includes:
state
dev
npu_state
npu info
offload
vlan
qid
diagnose sys session list
Search for:
state=dirty
or:
state = dirty
DIRTY
  |
  +-- Path/interface changed
  +-- Session state requires update
  +-- Possible CPU processing increase
Look for:
npu info:
offload=8/8
or:
offload=8/0
8/8 = both directions offloaded
8/0 = one direction offloaded
Example:
reflect info 0:
dev=37->38/38->37
npu_state=0x000400
npu info:
flag=0x91/0x00
offload=8/0
...
total reflect session num: 1
The exact fields are platform/version dependent.
Use the output to correlate:
Session
   |
   +-- Interface
   +-- NPU state
   +-- Offload state
   +-- Auxiliary/reflect information
                 ISP-1
              192.168.254.2
                   |
                 port1
                   |
             +-----------+
             |  FGT-1    |
             +-----------+
                   |
                 port3
                   |
             192.168.20.0/24
                   |
                Server
                 ISP-2
              192.168.200.1
                   |
