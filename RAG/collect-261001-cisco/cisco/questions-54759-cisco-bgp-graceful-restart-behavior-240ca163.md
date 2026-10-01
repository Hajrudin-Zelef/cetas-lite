---
id: collect-261001-cisco/cisco/questions-54759-cisco-bgp-graceful-restart-behavior-240ca163
title: "questions-54759-cisco-bgp-graceful-restart-behavior-240ca163"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-54759-cisco-bgp-graceful-restart-behavior-240ca163.md
source_anchor: ""
source_lines: [1, 57]
sha256: ee5270f2a4dad872e329888427f3c904271d70bb727f539ed83bbb40985f8c93
---

# questions-54759-cisco-bgp-graceful-restart-behavior-240ca163

In our environment we generally have ASR-1000X-2s everywhere peering via iBGP/eBGP. These routers have no redundant RPs, hence cannot keep forwarding traffic while the router reboots or crashes. As such, this is a clear example of a router that's only NSF-aware (or graceful-restart-aware) but not capable.
The reason I enabled this is because, from RFC 4724:
  In addition, even if the speaker does not have the ability to preserve its forwarding state for any address family during BGP restart, it is still recommended that the speaker advertise the Graceful Restart Capability to its peer (as mentioned before this is done by not including any in the advertised capability). There are two reasons for doing this. The first is to indicate its intention of generating the End-of-RIB marker upon the completion of its initial routing updates, as doing this would be useful for routing convergence in general. The second is to indicate its support for a peer which wishes to perform a graceful restart.
So what I would expect to see in the "show ip bgp neighbor " command, regarding Graceful Restart, would be something like the following:
BGP neighbor is <IP>,  remote AS <AS>, internal link
[...]
  Neighbor capabilities:
[...]
    Graceful Restart Capability: advertised and received
      Remote Restart timer is 120 seconds
      Address families advertised by peer:
        none
Basically, GR is negotiated, but no address family is specified, effectively only using the EOR marker for routing convergence improvements.
Instead, here's what the router specifies:
BGP neighbor is <IP>,  remote AS <AS>, internal link
[...]
  Neighbor capabilities:
[...]
    Graceful Restart Capability: advertised and received
      Remote Restart timer is 120 seconds
      Address families advertised by peer:
        IPv4 Unicast (was not preserved, VPNv4 Unicast (was not preserved
My assumption is that the 'was not preserved' in the parentesis refers to the most recent restart of the neighbor, and it means that when the neighbor re-established the BGP connection, the GR Capability for IPv4 and VPNv4 AFIs did not set the "Forwarding bit" as specified by the GR-RFC:
  Once the session is re-established, if the "Forwarding State" bit for a specific address family is not set in the newly received Graceful Restart Capability, or if a specific address family is not included in the newly received Graceful Restart Capability, or if the Graceful Restart Capability is not received in the re-established session at all, then the Receiving Speaker MUST immediately remove all the stale routes from the peer that it is retaining for that address family.
Clearly the Forwarding State bit is never going to be set by this type of router, due to hardware limitations. Here's my concern though: What happens when the router reboots, and the neighboring routers keep forwarding packets to this router because the GR-capabily did specify IPv4 and VPNv4 AFI/SAFI? This would clearly cause impact as traffic would be blackholed.
I will try to simulate this and see how it behaves, I'll report back, but any info you have it would be greatly appreciated.
EDIT:
So I did some tests. I got a packet capture going on between 2 of the routers. They both already had a BGP session going on, so I simply HARD-cleared the session from 1 side (7.0.0.1). The other side, 7.0.0.2, first sent an OPEN message with the following GR capability info:
Capability: Graceful Restart capability
    Type: Graceful Restart capability (64)
    Length: 6
    Restart Timers: 0x0078
        0... .... .... .... = Restart: No
        .... 0000 0111 1000 = Time: 120
    AFI: IPv4 (1)
    SAFI: Unicast (1)
    Flag: 0x00
        0... .... = Preserve forwarding state: No
Notice the AFI/SAFI for IPv4, which according to RFC means that this router supports GR for this AF.
And here is the GR capability info from the other side, where I reset the BGP peering from (7.0.0.1):
Capability: Graceful Restart capability
    Type: Graceful Restart capability (64)
    Length: 6
    Restart Timers: 0x8078, Restart
        1... .... .... .... = Restart: Yes
        .... 0000 0111 1000 = Time: 120
    AFI: IPv4 (1)
    SAFI: Unicast (1)
    Flag: 0x00
        0... .... = Preserve forwarding state: No
An interesting this to note here is that this router sets the Restart Bit even though it didn't technically restart as a whole. I assume this is useful, since:
  When set (value 1), this bit indicates that the BGP speaker has restarted, and its peer MUST NOT wait for the End-of-RIB marker from the speaker before advertising routing information to the speaker.
So it allows neighboring peers to immediately send Updates to this router because BGP restarted, without waiting for an EOR packet.
Anyway, when I actually restarted one of the routers, through an induced crash, GR clearly did not work. The 2 routers have an iBGP peering via loopbacks, hence when the loopback is not reachable, the routes for that neighbor are not used because of "next-hop-self" and the next-hop no longer being reachable. So clearly I lose access to the routes that the restarting neighbor is advertising, even though the RFC clearly states:
  When the Receiving Speaker detects termination of the TCP session for a BGP session with a peer that has advertised the Graceful Restart Capability, it MUST retain the routes received from the peer for all the address families that were previously received in the Graceful Restart Capability and MUST mark them as stale routing information.
So the routes should have been simply marked as "stale" but still usable.
So I think the software either has a bug or it just doesn't follow the RFC. As long as we have only this type / model of routers we should be fine, though a different platform might take the GR-capability advertised by this kind of router seriously and blackhole traffic for as long as this router reboots, which obviously could be a problem...
