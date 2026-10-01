---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/how-to-port-forward-on-unifi-devices-488a0645-1
title: "how-to-port-forward-on-unifi-devices-488a0645"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/how-to-port-forward-on-unifi-devices-488a0645.md
source_anchor: ""
source_lines: [1, 32]
sha256: c72e5ae3a3fbfd3354dec9e787662b1b04b18a0e08ffe274b7115c9e5aac255e
---

# how-to-port-forward-on-unifi-devices-488a0645

UniFi port forwarding takes about a minute to set up, and that is the problem with it. It is the fastest way to expose something on your network to the internet, which means it is also the fastest way to expose something you did not mean to. This guide covers where port forwarding lives in the current UniFi Network application, how to set up a rule, how to lock it down with the zone-based firewall so the whole world isn’t knocking on it, and the two alternatives I’d reach for first.
One thing before the steps. I’ve said this in my UniFi security videos, and it still holds: you should port forward as little as you possibly can. If what you’re after is remote access to your own stuff, a WireGuard VPN on your UniFi gateway gets you there in a few clicks and forces you to authenticate before anything is reachable. Port forwarding is for the cases where that doesn’t work, like a game server your friends connect to, a Plex library you share with family, or a service that has to answer to the public.
Where UniFi Port Forwarding Lives in Network 10
This is the part that changed, and it’s why half the tutorials you’ll find (including the previous version of this one) send you to a menu that isn’t there anymore. In UniFi Network 10, port forwarding moved into the Policy Engine along with firewall policies, NAT rules, and routes. The path is Settings, then Policy Engine, then Policy Table. Click Create New Policy and choose Port Forwarding as the policy type. The Routing page that used to hold port forwards is now just BGP and OSPF.
If you’re still on Network 9, the same form lives under Settings, then Routing, then Port Forwarding. The fields are identical either way, so the steps below apply to both. You can check your version at the bottom of the Settings sidebar.
How to Port Forward on a UniFi Gateway
These steps cover UniFi gateway port forwarding on any model running the Network application: the Cloud Gateway line, the Dream Machines, the Dream Router, and the UniFi Express. Before you start, give the device you’re forwarding to a fixed IP address, either a DHCP reservation in UniFi or a static address on the device, because a port forward pointed at an address that changes tomorrow is a port forward that stops working tomorrow.
- Open Settings, go to Policy Engine, open the Policy Table, and click Create New Policy. Pick Port Forwarding.
- Fill in the rule using the fields below.
- Save it. UniFi creates the matching firewall policy for you, and the rule is live immediately.
- Name: whatever tells you what this rule is for six months from now. “Plex for family” beats “rule 1”.
- WAN Interface: only matters if you have two WANs. Pick the one your public traffic arrives on, or all of them.
- WAN Port: the port the outside world connects to. It does not have to match the internal port, and for anything that gets scanned constantly (SSH, RDP, anything on 22 or 3389), it should not.
- From: Any opens it to the entire internet. Limited restricts it to an IP address, a range, or a subnet. Use Limited every time you can, and I’ll get to what to do when you can’t.
- Forward IP Address: the local address of the device you’re exposing. The Select Device link pulls it from your client list, which avoids typos.
- Forward Port: the port the service is actually listening on inside your network.
- Protocol: TCP, UDP, or both. Check the service’s documentation and pick only what it needs, because “both” is one more thing exposed for no reason.
- Syslog Logging: turn it on. Logs are how you find out later who’s been trying the door.
For example, forwarding Plex looks like WAN port 32400, forward port 32400, protocol TCP, pointed at the Plex server’s fixed IP. A Minecraft server is 25565 TCP. Test it from outside your network (your phone on mobile data, not your WiFi) before you tell anyone it’s ready, because testing from inside doesn’t prove anything.
What UniFi Does in the Firewall When You Save
The moment you save a port forward, UniFi creates a firewall policy that allows that traffic in. Under the zone-based firewall, it shows up in the Policy Table as a second row, named something like “Allow Port Forward”, with a source zone of External and a destination zone of wherever the target device lives. In the screenshot above, a forward to a device on the Untrusted network produced an External to Untrusted allow on port 443. That auto-created policy is read-only, which trips people up when they go looking for a way to tighten it. You do not edit it. You add your own policy next to it, and that is the next section.
It’s also worth knowing that the port forward itself appears as a policy of type Port Forwarding with the action Translate, alongside NAT rules. If you ever want to audit what’s exposed, filter the Policy Table by the Port Forwarding type and you’ll see every rule in one list. Do this periodically, and if there’s anything in that list you can’t explain, that’s the thing to fix first.
Locking It Down With the Zone-Based Firewall
The From field limits access to IP addresses, which is great when you know them and useless when you do not. A family member’s home connection has a dynamic IP that will change, and a public service has no list at all. This is where the zone-based firewall earns its keep, because it can restrict a forwarded port in ways the rule itself can’t.
The easiest and most effective restriction is by country. If you’re sharing Plex with family who all live in the same country you do, create a firewall policy in the External zone that blocks the forwarded port from everywhere except that country. Nobody in your family is connecting from overseas, so nobody overseas needs the port. That one policy removes the vast majority of the automated scanning you’d otherwise see in the logs. If you’re limiting a family member instead, the other approach is to look up their ISP’s IP blocks and allow those ranges in the From field, which survives their address changing without opening the port to everyone.
Two things go together with the port forward itself:
- Put the exposed device on its own VLAN in the DMZ zone. The DMZ zone exists for exactly this. It’s a set of policies that lets the device talk to the internet and nothing internal, so if the service is compromised, the attacker is standing in an empty room instead of your LAN. It’s not wrong to use another isolated zone for it, since a zone is just a defined set of rules, but the DMZ is the one designed for a public server.
- Turn on intrusion detection and prevention. It’s in Settings under CyberSecure, and with a forwarded port you want it on Notify and Block, not just Notify. A port forward is the one place on your network where you’ve invited traffic in, so it’s where IDS/IPS pays for itself.
My guide to UniFi firewall rules goes through the zone matrix in detail if you’re setting up zones for the first time.
Two Settings That Port Forward Without Telling You
An empty port forwarding list does not mean nothing is exposed. Two settings will open ports on your behalf, and I’d turn both off.
- UPnP. Universal Plug and Play lets devices on your network ask the gateway to open ports for them, with no approval from you. A NAS, a game console, or a media server can port forward itself and you’d never see it in your list. The way I think about it: an unlocked front door you know about is bad, but an unlocked front door you don’t know about is worse. Turn it off in Settings under Internet. If a device genuinely needs a port, add the forward yourself so it’s in the list.
