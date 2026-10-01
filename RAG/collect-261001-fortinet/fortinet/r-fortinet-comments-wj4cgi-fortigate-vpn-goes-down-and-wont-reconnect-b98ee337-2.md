---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-wj4cgi-fortigate-vpn-goes-down-and-wont-reconnect-b98ee337-2
title: "config vpn ipsec phase1-interfac"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-wj4cgi-fortigate-vpn-goes-down-and-wont-reconnect-b98ee337.md
source_anchor: ""
source_lines: [3, 113]
sha256: fe3b7fdd3bcd8f42fb42f4dba5bea8d20752160cb30f0d22455e6784a97acca8
---

# config vpn ipsec phase1-interfac

    Hi,
I have a Fortigate that has an IPSec VPN setup to another FortiGate appliance.
I have the tunnel successfully established, and then randomly, the tunnel will be down and won't come back up until I reboot one device.
While the tunnel is down I have run the following tests:
- 
      Successfully ping from one device wan address to the other
- 
      Can successfully trace route from one device to the other
- 
      Run diagnose vpn ike gateway, and can see the status as connecting
- 
      Checked that IKE packets are being sent on port 500 successfully
- 
      Debug IKE and can see the following info.
I'm using IKE v2, and all my proposals and configuration is identical on both sides.
Source settings:
Destination Settings
Any ideas what would be the cause?
As stated, a reboot fixes it for a few days, then it just falls back into this state.
Cheers!
Section des commentaires
Do you have blackhole routes setup on both sides? Ran into similar issues, and flapping the port worked as well as a reboot on the remote fgt. However, without doing that the tunnel(s) would not come back up because the traffic was traversing the wrong interface. Adding the blackhole routes corrected this for me.
ETA: After adding the routes, you need to kill any active/stale sessions to get the tunnels to come back up after making those changes.
I don't have black hole routes setup. Can you please explain how to configure a little further and I'll try that out please?
Here you go: https://community.fortinet.com/t5/FortiGate/Technical-Note-Use-of-Black-hole-route-in-site-to-site-IPsec-VPN/ta-p/192526
Sorry, I thought I had posted that earlier.
Try the commands on this link:- https://community.fortinet.com/t5/FortiGate/Technical-Tip-Inbound-IPsec-traffic-dropped-due-to-layer-2/ta-p/208035 What FortiOS version are FortiGates currently on? How frequent do you get the disconnection? Is the IPsec tunnel shows Green(UP) despite no traffic is flowing? Instead of restarting the Gates, try restarting the IPsec tunnel by going to Dashboard>Network>IPsec and bring down all Phase 2.
I'm running v6.2.9 build 1234 (GA).
Disconnects happen every 2 - 5 days randomly.
The source side was showing red, but the destination green. After enabling dead peer detection to "On Idle" on both sides, now they both are red.
Don't seem to have a "Network" option in my dashboard. When I go to Monitor > IPsec monitor. I can select a tunnel, but the "Bring Down" is greyed out.
After reading through the link provided, and trying to run the "diagnose npu np6 dce 0" command.
It won't accept the "npu" portion, the command does not exist as an option after the diagnose command.
Hmm which model of FortiGate is that ? Maybe it doesnt have np6...
Any chance you can update to 6.4.x as I`m seeing some IPsec VPN issues on the release note of 6.2.9:
https://docs.fortinet.com/document/fortigate/6.2.9/fortios-release-notes/236526/known-issues
Generally NO SUITABLE IKE_SA means that the 2 Gates IPsec config (Phase 1 & 2) are not the same and hence can`t establish the tunnel. It can be Authentication(not the same pre-shared key) /Phase1(Algo,DH Groups)/Phase2 misconfiguration.
Also when I try to type the following it does not recognised the "npu-offload" command.
# config vpn ipsec phase1-interfac
edit <phase1-name>
set npu-offload disable
end
“No suitable IKE proposal” reeks of phase-1 misconfig.
Initially get rid of all the default proposals on the P1/P2 settings and choose the the proposal you actually want to use. Make sure timers are identical. AES128-GCM and SHA384 with DH 21 is a good start for very secure encryption. Cramming 4-5 proposals in IKE and leaving it up to negotiation is not a good approach.
Failing that, post up a sanitised copy of “show vpn ipsec phase1-interface” and “show vpn ipsec phase2-interface” from the CLI output.
Thanks for the suggestion, below is the config are the changes.
Source Config
FortiGate-Source# show vpn ipsec phase1-interface
config vpn ipsec phase1-interface
edit "ToDestination"
set interface "wan1"
set ike-version 2
set keylife 28800
set peertype any
set net-device disable
set proposal aes128gcm-prfsha384
set dpd on-idle
set dhgrp 21
set nattraversal disable
set remote-gw x.x.x.x
set psksecret ENC xxxxxxxxxxxxxxxxxxxxxxx
End
FortiGate-Source# show vpn ipsec phase2-interface
config vpn ipsec phase2-interface
edit "ToDestination"
set phase1name "ToDestination"
set proposal aes128gcm
set dhgrp 21
set keepalive enable
set keylifeseconds 28800
End
Destination Config
FortiGate-Destination # show vpn ipsec phase1-interface
config vpn ipsec phase1-interface
edit "ToSource"
set interface "wan1"
set ike-version 2
set keylife 28800
set peertype any
set net-device disable
set proposal aes128gcm-prfsha384
set dpd on-idle
set dhgrp 21 5
set nattraversal disable
set remote-gw x.x.x.x
set psksecret ENC xxxxxxxxxxxxxxxxxxxxxxxxxxx
End
FortiGate-Destination # show vpn ipsec phase2-interface
config vpn ipsec phase2-interface
edit "ToSource"
set phase1name "ToSource"
set proposal aes128gcm
set dhgrp 21
set keepalive enable
set keylifeseconds 28800
End
I've had exact the same issues. How is your FortiGate connected to the internet?
My FortiGate was connected to a briged G.fast router and when the IPsec tunnels disconnected I could reboot either the Forti or the Briged Router and then the tunnel came up again. So I investigated more and tryed to upgrade the FortiGate to v7.0.6 and the Firmware of the bridged router but without success.
The fix for this case was to change the bridged Router to another model. After the replacement to another model the IPsec tunnels stayed up at all time.
Thanks for your suggestion.
Indeed I do have the internet connected via a modem that is in bridge mode on the source side - this is the side I reboot when I have the issue. The connection is Fibre to the Basement, then up the building riser is copper pair. This then connects to a TP-Link Archer VR1600v router that is in bridged mode.
The LAN port of the VR1600v is connected to WAN1 of the Fortigate.
The Fortigate WAN1 is setup as DHCP and authenticates the PPoE.
Might need to try a different router to use as bridge.
Any suggestions of a good reliable model I should buy?
Thanks!
I've changed from a Zyxel XMG to Draytek Vigor166. Have no issues so far.
Just for my personal interest, could you let me know if you can bring up the IPsec tunnel when you just boot your briged router and not the FW and then try to bring-up phase 1 & 2?
I had my wan1 also as DHCP with PPPoE.
Try enabling DPD .
Now seeing this come through on the debug too.
