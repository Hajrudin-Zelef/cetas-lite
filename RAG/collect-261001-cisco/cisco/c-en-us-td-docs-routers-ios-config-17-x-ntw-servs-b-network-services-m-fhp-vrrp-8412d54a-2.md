---
id: collect-261001-cisco/cisco/c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp-8412d54a-2
title: "c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["preemption"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a.md
source_anchor: ""
source_lines: [99, 202]
sha256: fd36ee94f13a76f0211167430113f9b20b8534320c3a4baf2a24a698df1a58b4
---

# c-en-us-td-docs-routers-ios-config-17-x-ntw-servs-b-network-services-m-fhp-vrrp--8412d54a

An important aspect of the VRRP redundancy scheme is VRRP router priority. Priority determines the role that each VRRP router
plays and what happens if the virtual primary router fails.
If a VRRP router owns the IP address of the virtual router and the IP address of the physical interface, this router will
function as a virtual primary router.
Priority also determines if a VRRP router functions as a virtual router backup and the order of ascendancy to becoming virtual
primary router if the virtual primary router fails. You can configure the priority of each virtual router backup with a value
of 1 through 254 using the vrrppriority command.
For example, if Router A, the virtual primary router in a LAN topology, fails, an election process takes place to determine
if virtual router backups B or C should take over. If Routers B and C are configured with the priorities of 101 and 100, respectively,
Router B is elected to become virtual primary router because it has the higher priority. If Routers B and C are both configured
with the priority of 100, the virtual router backup with the higher IP address is elected to become the virtual primary router.
By default, a preemptive scheme is enabled whereby a higher priority virtual router backup that becomes available takes over
for the virtual router backup that was elected to become virtual primary router. You can disable this preemptive scheme using
the novrrppreempt command. If preemption is disabled, the virtual router backup that is elected to become virtual primary router remains as
the primary until the original virtual primary router recovers and becomes the primary again.
VRRP Advertisements
The virtual primary router sends VRRP advertisements to other VRRP routers in the same group. The advertisements communicate
the priority and state of the virtual primary router. The VRRP advertisements are encapsulated in IP packets and sent to the
IP Version 4 multicast address assigned to the VRRP group. The advertisements are sent every second by default; the interval
is configurable.
Although the VRRP protocol as per RFC 3768 does not support millisecond timers, Cisco routers allow you to configure millisecond
timers. You need to manually configure the millisecond timer values on both the primary and the backup routers. The primary
advertisement value displayed in the showvrrp command output on the backup routers is always 1 second because the packets on the backup routers do not accept millisecond
values.
Note
Use millisecond timers where absolutely necessary, with careful consideration and testing.
Millisecond values work only under favorable circumstances, and you must be aware that the use of the millisecond timer values
restricts VRRP operation to Cisco devices only.
Be cautious while using millisecond timers when SNMP traps are enabled for VRRP.
VRRP Object Tracking
Object tracking is an
independent process that manages creating, monitoring, and removing tracked
objects such as the state of the line protocol of an interface. Clients such as
the Hot Standby Router Protocol (HSRP), Gateway Load Balancing Protocol (GLBP),
and VRRP register their interest with specific tracked objects and act when the
state of an object changes.
Each tracked object
is identified by a unique number that is specified on the tracking CLI. Client
processes such as VRRP use this number to track a specific object.
The tracking process
periodically polls the tracked objects and notes any change of value. The
changes in the tracked object are communicated to interested client processes,
either immediately or after a specified delay. The object values are reported
as either up or down.
VRRP object tracking
gives VRRP access to all the objects available through the tracking process.
The tracking process allows you to track individual objects such as a the state
of an interface line protocol, state of an IP route, or the reachability of a
route.
VRRP provides an
interface to the tracking process. Each VRRP group can track multiple objects
that may affect the priority of the VRRP device. You specify the object number
to be tracked and VRRP is notified of any change to the object. VRRP increments
(or decrements) the priority of the virtual device based on the state of the
object being tracked.
How VRRP Object Tracking
Affects the Priority of a Device
The priority of a device can change dynamically if it has been configured for object tracking and the object that is being
tracked goes down. The tracking process periodically polls the tracked objects and notes any change of value. The changes
in the tracked object are communicated to VRRP, either immediately or after a specified delay. The object values are reported
as either up or down. Examples of objects that can be tracked are the line protocol state of an interface or the reachability
of an IP route. If the specified object goes down, the VRRP priority is reduced. The VRRP device with the higher priority
can now become the virtual primary device if it has the vrrppreempt command configured. See the “VRRP Object Tracking” section for more information on object tracking.
In Service Software Upgrade--VRRP
VRRP supports In Service Software Upgrade (ISSU). In Service Software Upgrade (ISSU) allows a high-availability (HA) system
to run in stateful switchover (SSO) mode even when different versions of Cisco IOS XE software are running on the active and
standby Route Processors (RPs) or line cards.
ISSU provides the ability to upgrade or downgrade from one supported Cisco IOS XE release to another while continuing to
forward packets and maintain sessions, thereby reducing planned outage time. The ability to upgrade or downgrade is achieved
by running different software versions on the active RP and standby RP for a short period of time to maintain state information
between RPs. This feature allows the system to switch over to a secondary RP running upgraded (or downgraded) software and
continue forwarding packets without session loss and with minimal or no packet loss. This feature is enabled by default.
For detailed information about ISSU, see the Cisco IOS XE In Service Software Upgrade Process document in the Cisco IOS XE High Availability Configuration Guide.
VRRP Support for Stateful
Switchover
With the introduction
of the VRRP Support for Stateful Switchover feature, VRRP is SSO aware. VRRP
can detect when a router is failing over to the secondary RP and continue in
its current group state.
SSO functions in
networking devices (usually edge devices) that support dual Route Processors
(RPs). SSO provides RP redundancy by establishing one of the RPs as the active
processor and the other RP as the standby processor. SSO also synchronizes
critical state information between the RPs so that network state information is
dynamically maintained between RPs.
Prior to being SSO
aware, if VRRP was deployed on a router with redundant RPs, a switchover of
roles between the active RP and the standby RP would result in the router
relinquishing its activity as a VRRP group member and then rejoining the group
as if it had been reloaded. The SSO--VRRP feature enables VRRP to continue its
activities as a group member during a switchover. VRRP state information
between redundant RPs is maintained so that the standby RP can continue the
router’s activities within the VRRP during and after a switchover.
This feature is
enabled by default.
To disable this feature, use the
novrrpsso command in global configuration mode.
For more information, see the Stateful Switchover
document.
How to Configure VRRP
VRRP
Customizing the behavior of VRRP is optional. Be aware that as soon as you enable a VRRP group, that group is operating.
It is possible that if you first enable a VRRP group before customizing VRRP, the router could take over control of the group
and become the virtual primary router before you have finished customizing the feature. Therefore, if you plan to customize
