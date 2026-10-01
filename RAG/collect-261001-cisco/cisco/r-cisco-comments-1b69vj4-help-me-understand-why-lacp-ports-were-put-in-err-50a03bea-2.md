---
id: collect-261001-cisco/cisco/r-cisco-comments-1b69vj4-help-me-understand-why-lacp-ports-were-put-in-err-50a03bea-2
title: "r-cisco-comments-1b69vj4-help-me-understand-why-lacp-ports-were-put-in-err-50a03bea"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/r-cisco-comments-1b69vj4-help-me-understand-why-lacp-ports-were-put-in-err-50a03bea.md
source_anchor: ""
source_lines: [3, 103]
sha256: 956ffadca11cd4978dea12637b8e63e78f396adb721f974e3d1b57da8e12522f
---

# r-cisco-comments-1b69vj4-help-me-understand-why-lacp-ports-were-put-in-err-50a03bea

    Hi
we have a pair of stacked C9500s with a port on each configured in an Access Port channel connecting to a physically separate ring network based on Hirschmann switches.
For the second time now, all of a sudden, the switch has put both ethernet ports in err_disabled - the first time, months after initial configuration, and Saturday, months after the previous error.
A simple shut / no shut solves the issue, but I would like to understand the root cause of this problem and avoid it in the future, since this connection is very crucial.
I am hypothesizing a BPDU issue in occasion of something happening on their network such as a change in the redundant ring - but the logs are not very explicit:
Mar  2 09:54:47.609: %PM-4-ERR_DISABLE: channel-misconfig error detected on Po17, putting Twe1/0/17 in err-disable state
Mar  2 09:54:47.613: %PM-4-ERR_DISABLE: channel-misconfig error detected on Po17, putting Twe2/0/17 in err-disable state
Mar  2 09:54:47.617: %PM-4-ERR_DISABLE: channel-misconfig error detected on Po17, putting Po17 in err-disable state
Mar  2 09:54:48.609: %LINEPROTO-5-UPDOWN: Line protocol on Interface TwentyFiveGigE1/0/17, changed state to down
Mar  2 09:54:48.614: %LINEPROTO-5-UPDOWN: Line protocol on Interface TwentyFiveGigE2/0/17, changed state to down
Mar  2 09:54:48.614: %LINEPROTO-5-UPDOWN: Line protocol on Interface Port-channel17, changed state to down
Mar  2 09:54:49.613: %LINK-3-UPDOWN: Interface TwentyFiveGigE1/0/17, changed state to down
Mar  2 09:54:49.616: %LINK-3-UPDOWN: Interface Port-channel17, changed state to down
Mar  2 09:54:49.616: %LINK-3-UPDOWN: Interface TwentyFiveGigE2/0/17, changed state to down
Mar  2 10:44:42.569: %SEC_LOGIN-5-LOGIN_SUCCESS: Login Success [localport: 22] at 11:44:42 CEST Sat Mar 2 2024
Mar  2 10:50:14.170: %LINK-5-CHANGED: Interface TwentyFiveGigE1/0/17, changed state to administratively down
Mar  2 10:50:14.180: %LINK-5-CHANGED: Interface TwentyFiveGigE2/0/17, changed state to administratively down
Mar  2 10:50:14.217: %LINK-5-CHANGED: Interface Port-channel17, changed state to administratively down
Mar  2 10:52:49.770: %LINK-3-UPDOWN: Interface Port-channel17, changed state to down
Mar  2 10:52:49.771: %LINK-3-UPDOWN: Interface TwentyFiveGigE1/0/17, changed state to down
Mar  2 10:52:49.773: %LINK-3-UPDOWN: Interface TwentyFiveGigE2/0/17, changed state to down
Mar  2 10:52:54.570: %LINK-3-UPDOWN: Interface TwentyFiveGigE1/0/17, changed state to up
Mar  2 10:52:54.620: %LINK-3-UPDOWN: Interface TwentyFiveGigE2/0/17, changed state to up
Mar  2 10:52:56.459: %LINEPROTO-5-UPDOWN: Line protocol on Interface TwentyFiveGigE1/0/17, changed state to up
Mar  2 10:52:56.468: %LINEPROTO-5-UPDOWN: Line protocol on Interface TwentyFiveGigE2/0/17, changed state to up
Mar  2 10:52:57.457: %LINK-3-UPDOWN: Interface Port-channel17, changed state to up
Mar  2 10:52:58.458: %LINEPROTO-5-UPDOWN: Line protocol on Interface Port-channel17, changed state to up
There is no port-security enabled on these ports:
Port Security              : Disabled
Port Status                : Secure-down
Violation Mode             : Shutdown
Aging Time                 : 0 mins
Aging Type                 : Absolute
SecureStatic Address Aging : Disabled
Maximum MAC Addresses      : 1
Total MAC Addresses        : 0
Configured MAC Addresses   : 0
Sticky MAC Addresses       : 0
Last Source Address:Vlan   : 0000.0000.0000:0
Security Violation Count   : 0
And ErrDisable reason are set to default:
ErrDisable Reason            Detection        Mode
-----------------            ---------        ----
arp-inspection               Enabled          port
bpduguard                    Enabled          port
channel-misconfig            Enabled          port
community-limit              Enabled          port
dhcp-rate-limit              Enabled          port
dtp-flap                     Enabled          port
evpn-mh-core-isolation       Enabled          port
gbic-invalid                 Enabled          port
iif-reg-failure              Enabled          port
inline-power                 Enabled          port
invalid-policy               Enabled          port
l2ptguard                    Enabled          port
link-flap                    Enabled          port
link-monitor-failure         Enabled          port
loopback                     Enabled          port
loopdetect                   Enabled          port
lsgroup                      Enabled          port
oam-remote-failure           Enabled          port
mac-limit                    Enabled          port
pagp-flap                    Enabled          port
port-mode-failure            Enabled          port
pppoe-ia-rate-limit          Enabled          port
psecure-violation            Enabled          port/vlan
security-violation           Enabled          port
sfp-config-mismatch          Enabled          port
sgacl_limitation:enforcem    Enabled          port
sgacl_limitation:multiple    Enabled          port
storm-control                Enabled          port
udld                         Enabled          port
psp                          Enabled          port
dual-active-recovery         Enabled          port
evc-lite input mapping fa    Enabled          port
vsl-and-non-vsl-port-pair    Enabled          port
fasthello-and-non-fasthel    Enabled          port
mvrp                         Enabled          port
mrp-miscabling               Enabled          port
Does anyone have an idea? Unfortunately I did not know what else to do other than shut / no shut.
Section des commentaires
The error would indicate that the port channel configuration is different on them, I would compare the configuration on each.
If that were the case, it wouldn't error out once every 4 months would it?
All they sent me is a screenshot of the Hirschmann web gui where they enable link aggregation, not much - so I was hoping for someone to have a bit of experience with them.
Spanning tree has etherchannel guard misconfig in addition to the generic ErrDisable cause. This might be a reason for Cisco to disable the ports.
[Edit] found a description for this feature on a blog: "This feature is somewhat of a hack. If you keep receiving BPDUs from several MAC addresses, the feature will assume that you have a bundling problem and shut down the port.:
On the Hirschmann side, are the ports configured as plain LACP access, no funky industrial protocols active?
You might consider to deactivate the ErrDisable Reason if it doesn't impact other functionality.
The ports themselves are, but behind that switch there's a double redundant ring that uses their own protocols, and a whole host of things configured with industrial protocols as well.
I know I can set a short recovery timer for the specific error, and I did, but is there a way to disable errdisable channel-misconfig entirely?
Can it be disabled by port or only for the whole switch?
What are your port configs?
The 2 x access ports and the port channel?
Sorry I just implied it, it's literally just two ports in access with a port channel.
Speed is left to auto with a 1G transceiver, same as the other side.
In that case mate, a nice first step would be to enable bpdu filter on the access ports and the access-port-channel.
That just discards the bpdu's from the other switch, which is what I would bet is sending bpdu's and your switch is seeing the misconfig and getting its knickers in a twist.
It says misconfig error, but you haven’t shown us any port configs.
Sorry I just implied it, it's literally just two ports in access with a port channel.
Speed is left to auto with a 1G transceiver, same as the other side.
It’s hard to say for sure. But technically, your ports aren’t set to access. They’re set to dynamic because you haven’t typed “switchport mode access”. I have had ports flip from access to trunk and back before because they were left dtp dynamic like yours.
