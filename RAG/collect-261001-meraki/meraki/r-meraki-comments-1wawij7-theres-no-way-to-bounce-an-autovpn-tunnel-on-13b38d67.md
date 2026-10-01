---
id: collect-261001-meraki/meraki/r-meraki-comments-1wawij7-theres-no-way-to-bounce-an-autovpn-tunnel-on-13b38d67
title: "r-meraki-comments-1wawij7-theres-no-way-to-bounce-an-autovpn-tunnel-on-13b38d67"
domain: meraki
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "lean"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-1wawij7-theres-no-way-to-bounce-an-autovpn-tunnel-on-13b38d67.md
source_anchor: ""
source_lines: [1, 45]
sha256: 94458990fd6afb7a5e5181bd720be854e93d4622b48165b5940769a23b844d12
---

# r-meraki-comments-1wawij7-theres-no-way-to-bounce-an-autovpn-tunnel-on-13b38d67

theres no way to bounce an autovpn tunnel on meraki, so we ended up rebooting the MX instead (reboot bot) 
        
    question i kept coming back to: when a spoke loses its tunnel to the concentrator and doesnt come back on its own, how do you actually force it to renegotiate?
on a fortigate this is nothing. you reset the tunnel, it rebuilds, done. on meraki theres no equivalent. no cli, no api call to bring a tunnel down and up. you either wait and hope it recovers, or someone drives to the store and power cycles the box. at 1200 sites across 62 countries the second option isnt really an option.
what we ended up doing: wrote a bot that watches vpn status and reboots the MX when a store has lost both concentrator tunnels. reboot is the only lever meraki actually gives you, so we built around that instead of around the tunnel.
the logic took a couple of days. getting it to run reliably against the api took weeks, mostly rate limits and pagination behaviour that isnt obvious until you hit it at scale.
where it landed: one poll every 5 min, org level vpn statuses in a single call, target stores pulled by tag so the filtering happens meraki side, hub list cached at startup. handful of calls per scan, nowhere near the org rate limit. been running in production for a couple of months now.
one design note: it acts on a single scan rather than requiring two in a row. the gate that matters more is the cloud check. if the device cant reach the cloud at all its an isp or power problem, not something a reboot fixes, so it never fires and just raises an alert instead. reboot only happens when the box is clearly online but both tunnels are unreachable, and at that point the site is already isolated so the reboot costs nothing.
so, anyone solved this differently? genuinely curious if theres a way to force renegotiation that i missed, because rebooting a firewall to fix a tunnel still feels like the wrong shape of solution even though it works.
Section des commentaires
You can also disable and then re-enable VPN. Just ensure to wait for a config fetch after disabling the VPN before you re-enable.
huh, toggling the vpn mode is a much lighter touch than what were doing. and doing it through the api rather than by hand makes it a real option. big win is the lan staying up, when we reboot the box the switches and APs at the site get disrupted too.
one thing im stuck on though. how do you confirm the device actually pulled config? with a reboot i dont have to think about it, box comes back up and the tunnels form. with the toggle i have to set vpn mode to none, wait for the device to fetch config, confirm its up to date, then flip it back. and that fetch time isnt consistent across sites, some of ours are on pretty rough circuits.
is there something you check for, or do you just wait a fixed interval and hope?
we just wait until the dashboard says up to date again, normally within 60 sec. Thankfully this is a pretty rare problem for us - the tiny sites with ropey 4G connections we reduced to a single tunnel only and it increased reliability no end
the config toggle is clever but feels like same shape of problem as the reboot, just dressed different. you're still taking the whole vpn down and waiting for it to figure itself out instead of just bouncing the one tunnel that's dead
how fast does the re-enable actually come back in practice? i'm curious if you ever timed it vs a reboot
Absolutely correct. It's just less impactful
We toggle AutoVPN to off, wait for config to update, then back to spoke
huh, toggling the vpn mode is a much lighter touch than what were doing. and doing it through the api rather than by hand makes it a real option. big win is the lan staying up, when we reboot the box the switches and APs at the site get disrupted too.
one thing im stuck on though. how do you confirm the device actually pulled config? with a reboot i dont have to think about it, box comes back up and the tunnels form. with the toggle i have to set vpn mode to none, wait for the device to fetch config, confirm its up to date, then flip it back. and that fetch time isnt consistent across sites, some of ours are on pretty rough circuits.
is there something you check for, or do you just wait a fixed interval and hope?
You can just turn off the vpn and turn it back on (save in between, wait a minute) instead of bumping the entire box. Depending, if all the tunnels on that box are already down, there's no impact (vs rebooting the entire box that has impact)
also worth saying, when the tunnels to the concentrator are down the site cant really operate anyway, so a reboot isnt costing much on top of that. and honestly i lean on how quick meraki boxes come back up, its consistently within an acceptable window.
flipping vpn off and waiting for a config fetch might actually take longer than that in some cases.
The problem you describe, when AutoVPN fails to come back up to an MX in concentrator mode, usually only happens when the concentrator is sitting behind a device doing NAT with NAT traversal is set to "auto".
If, on the concentrator AutoVPN page, you enable NAT traversal, enter the public IP address and a UDP port, and then forward those to it from your firewall, this problem goes away.
It will be rock solid if you make this change.
we already have that set manually, public ip and udp port, not auto. still see the drops.
Are you sure the outside firewall is definately allowing that UDP port in from any source address on the Internet?
Pull the wireshark and filter isakmp to see what's wrong first.
Otherwise you could turn on another local subnet in phase 2 that normally isn't advertised. This would nudge the renogiation, and wouldn't take down any of the existing SAs.
API wise
First enable a "SA trigger" network 192.168.10.0/24
"subnets": [ { "localSubnet": "192.168.10.0/24", "useVpn": true
wait 10 seconds and then turn it off
"subnets": [ { "localSubnet": "192.168.10.0/24", "useVpn": false
This is autovpn not 3rd party
That doesn't matter for the SAs.
How dare you come in here and say something that makes sense!
I like it..
Autovpn has self healing properties - ie if the tunnel is down long enough the engine will restart itself. I'd be taking pcaps while it's down to see why it breaks. Do you have unfriendly Nat upstream? Does your wan port have a public IP or is it being nated by your ISP? Best practice is to have your ISP router in bridge mode and your wan port configured with a route able public IP
Ya most things VPN related on Meraki full on suck. Having used Sonicwalls. Palo Alto etc before with dozens of IPSec running never had issue with single tunnel impacting everything else. Have had their support tell me any changes to VPN can impact the entire stack.
yeah the "any vpn change can affect the whole stack" thing is exactly why we ended up doing this in software rather than touching config. coming from fortigate where you just reset a tunnel and move on, its a big adjustment.
ITT: people having a chat with Claude
