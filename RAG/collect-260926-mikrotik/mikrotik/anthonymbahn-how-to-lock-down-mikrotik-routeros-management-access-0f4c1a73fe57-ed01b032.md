---
id: collect-260926-mikrotik/mikrotik/anthonymbahn-how-to-lock-down-mikrotik-routeros-management-access-0f4c1a73fe57-ed01b032
title: "anthonymbahn-how-to-lock-down-mikrotik-routeros-management-access-0f4c1a73fe57-ed01b032"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2026-09"]
keywords: ["advisory", "incident"]
source: docs/RAG/lot-mikrotik/RouterOS/anthonymbahn-how-to-lock-down-mikrotik-routeros-management-access-0f4c1a73fe57-ed01b032.md
source_anchor: ""
source_lines: [1, 39]
sha256: 6b1ac5cdf82fa66e54831466edb96bfafa0a5321dbe04ffcf7faa8b2beec06be
---

# anthonymbahn-how-to-lock-down-mikrotik-routeros-management-access-0f4c1a73fe57-ed01b032

How to Lock Down MikroTik RouterOS Management Access
MikroTik RouterOS is powerful, cheap and everywhere, which is exactly why its management interfaces end up exposed. The September 2026 MikroTrick chain, an unauthenticated SSH takeover that attackers were exploiting before the advisory landed, is only the latest reminder that a ‘s control plane is a target in its own right. This guide is a concrete checklist for locking down RouterOS management access so that the next protocol bug is a on your schedule instead of someone else’s intrusion. Every command below is standard RouterOS configuration; adapt interface and address names to your environment, and test on a device you can physically reach before you apply it remotely.
Step 1: Patch and check for tampering first
Hardening a compromised device just protects the attacker’s foothold. Before anything else, upgrade to a fixed build for your channel: 6.49.21 or 7.23.4 on long-term, 7.24.2 on stable. From RouterOS 7.23.4 onward the system inspects the configuration at boot and raises a flagged state if it detects signs of unauthorized access. Read that state before you trust the device by printing the device-mode status with the command:
If it reports flagged, RouterOS has already disabled configuration it recognized as malicious and is blocking new scheduler, proxy and tunnel entries until you intervene. Do not clear the flag reflexively. Audit users, scripts, scheduler tasks and tunnels by hand first, then clear it deliberately, which requires a physical button press or a hard reboot:
Step 2: Take management services off untrusted networks
The single most valuable change is to stop offering management services to the internet. RouterOS exposes SSH, WebFig (www and www-ssl), Winbox, Telnet, FTP and the API through the service menu. Disable everything you do not use, and the MikroTik advisory specifically calls out SSH, WWW/WWW-SSL and -test as services to pull off untrusted networks during this incident.
Get Anthony Bahn’s stories in your inbox
Join Medium for free to get updates from this writer.
The bandwidth-test server is a good example of an on-by-default service that most production routers never need, and one of the September flaws abused it for an unauthenticated crash. Turn it off, along with the other convenience services that quietly listen unless you rely on them:
- /tool bandwidth-server set enabled=no
- /ip proxy set enabled=no
- /ip socks set enabled=no
- /ip upnpAllows devices to automatically configure port forwarding-convenient but a security risk. set enabled=no
- /ip dns set allow-remote-requests=no
Restrict layer-2 discovery so a device cannot be managed over MAC from arbitrary segments:
- /tool mac-server set allowed-interface-list=none
- /tool mac-server mac-winbox set allowed-interface-list=none
- /tool mac-server ping set enabled=no
Step 3: Filter the input chain
Disabling a service protects that service; a protects the whole device. RouterOS filters traffic destined for the router itself in the input chain, and the goal is simple: accept management only from inside your network, and drop everything else that arrives from the internet. Start from interface lists so the rules read clearly:
- /interface list add name=WAN
- /interface list add name=LAN
- /interface list member add interface=ether1 list=WAN
- /interface list member add interface=bridge list=LAN
Then build the input chain to accept existing connections and drop anything that is not coming from the LAN:
- /ip firewall filter add action=accept chain=input connection-state=established,related,untracked
- /ip firewall filter add action=accept chain=input protocol=icmp
- /ip firewall filter add action=drop chain=input in-interface-list=!LAN
That final drop rule is the one that matters. With it in place, an exposed SSH or WebFig service is no longer reachable from a WAN interface at all, which is why a firewalled RouterOS device shrugged off attacks that took over exposed ones. Apply the same treatment to the IPv6 input chain if the device has a routable address; an attacker does not care which protocol reaches the port.
Step 4: Harden the accounts and the crypto
Reduce what an attacker gains even if they do reach the login. Replace the default admin account with a named administrator and a long, mixed password, then disable the default, and require modern SSH algorithms so weak ciphers and key exchanges are off the table:
- /user add name=myname password=your-strong-password group=full
- /user disable admin
- /ip ssh set strong-crypto=yes
Step 5: Make it repeatable and watch it
One hardened router is a good afternoon; a hardened fleet is a policy. Export each device’s configuration, keep the exports in version control, and diff them on a schedule so an added “ops” user or an unexpected scheduler entry shows up as a change rather than a surprise. Forward logs off the device, because an attacker who owns the router owns its local log. The account-creation signature from the RouterOS attacks, an entry crediting the session as ssh:-2 from the attacker’s address, is only useful if the logs live somewhere the attacker cannot edit.
None of this is exotic, and that is the point. The MikroTrick chain, explained in our companion article on how SSH servers get tricked into skipping authentication, needed an exposed SSH port to work. The organizations that never gave it one turned a critical advisory into a routine upgrade. Treat management access the way our guide on how to restrict a web admin interface to trusted IP addresses treats web consoles: private by default, reachable only from where you administer it, and boring to attack. That is the outcome you are buying with an afternoon of configuration, and it holds against the next bug as well as this one.
Originally published at https://www.anthonybahn.com.
