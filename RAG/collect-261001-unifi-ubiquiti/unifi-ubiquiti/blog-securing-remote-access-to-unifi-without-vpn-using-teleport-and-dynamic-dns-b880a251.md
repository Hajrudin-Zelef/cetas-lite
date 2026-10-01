---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/blog-securing-remote-access-to-unifi-without-vpn-using-teleport-and-dynamic-dns-b880a251
title: "blog-securing-remote-access-to-unifi-without-vpn-using-teleport-and-dynamic-dns-b880a251"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["OpenAI", "SpaceX"]
dates: []
keywords: ["chatgpt", "fine-tuning"]
source: docs/RAG/collect-261001-unifi-ubiquiti/blog-securing-remote-access-to-unifi-without-vpn-using-teleport-and-dynamic-dns-b880a251.md
source_anchor: ""
source_lines: [1, 86]
sha256: 7988ad56d524c5fe25f21e6ae6a8516523cf430db08ccddaac64c66f4689833b
---

# blog-securing-remote-access-to-unifi-without-vpn-using-teleport-and-dynamic-dns-b880a251

Securing Remote Access to UniFi Without VPN Using Teleport and Dynamic DNS
If you need to manage your UniFi network remotely, maybe you’re offsite or supporting a client, but you don’t want to wrestle with traditional VPN, Teleport and Dynamic DNS (DDNS) is an elegant combo. Together they give you secure, seamless remote access without manual port forwarding or expensive static IPs.
In this post, we’ll explain how both fit together, walk through the setup, go into firewall fine-tuning, and touch on what to watch out for. By the end, you’ll be able to safely manage your network on the go.
Let’s dive in !!
Before we dive in, please don't self‑host your UniFi Controller if you take care of client networks. Sooner or later this will cause issues! It's fine for home users, but definitely not recommended for IT service businesses and MSPs. If you want secure, reliable and a scalable hosting solution check out UniHosted.
Why this combo: teleport + DDNS
First, let’s talk about the individual pieces:
- Teleport gives you zero‑config VPN access built into UniFi. It uses a one‑time invite to build a secure WireGuard tunnel, and reconnects automatically, even across NAT or behind CGNAT networks. There’s no need to open ports or manage certs.
- Dynamic DNS (DDNS) ensures your remote controller is reachable under a consistent hostname, even when your ISP changes your public IP. Services like No-IP, DuckDNS, and DynDNS automate IP updates, so you never have to guess or track IP manually.
When you use them together, you get a stable hostname and a secure tunnel to your network, no port forwarding, no VPN appliances, no static IPs. Just launch Teleport and connect through your DDNS address.
Step 1: Set up Dynamic DNS
First, configure a hostname that always points to your WAN IP:
- Register with your DDNS provider (e.g. myuniunifi.noip.com ).
- In UniFi Network navigate to Settings → Internet → Dynamic DNS.
- Choose your provider, enter hostname and credentials.
- UniFi will now update your public IP automatically—no extra hardware or scripts needed.
A DDNS address means you can always hit your UniFi instance without memorizing IPs.
Step 2: Enable and invite Teleport VPN
Teleport gives fast and secure remote access:
- Confirm Remote Access is toggled on under Settings → Console.
- Go to Settings → Teleport & VPN and turn Teleport on.
- Generate an Invite, valid for 24 hours and one-time use.
- Paste the link into WiFiman (desktop or mobile) or scan the QR code.
- Accept the invite. Your device gains a WireGuard address in 100.64.x.x .
- Teleport reconnects automatically whenever your device is online, no extra setup.
With the tunnel in place, remote tools (dashboard, SSH, SMB) connect as if you were local.
Step 3: Access through DDNS over the tunnel
Now your remote device has a secure tunnel even when outside. Next, configure UniFi to force DNS lookups of your DDNS hostname over that tunnel:
- On Teleport, traffic is routed via the UniFi gateway.
- The DDNS name resolves to your current public IP, but Teleport handles the relay internally.
- This makes connecting to the controller as easy as pointing your browser to https://myuniunifi.noip.com:8443 .
No port forwarding, no public exposure, just secure modular access.
Step 4: Harden with firewall rules
Teleport tunnels grant broad LAN access by default. To lock it down:
- Navigate to Settings → Firewall & Security → Rules → LAN IN.
- 
Create a DROP rule targeting Teleport’s subnet (e.g. 100.64.0.0/10 ).
- 
Add ALLOW rules beforehand, specifying source, port, and destination:
  - GUI (TCP 8443, UDP)
  - SSH (TCP 22)
  - SMB (TCP 445) or custom services
- Reorder rules so ALLOWs come before DROP.
- Save changes and test connectivity from a remote device.
By whitelisting access, you keep Teleport safe and scoped.
Step 5: Internal DNS considerations
By default, Teleport devices use public DNS. That means:
- Internal hostnames might not resolve correctly.
- You may need to connect via IP or manually assign DNS on clients.
Workarounds include:
- Setting static DHCP DNS entries in UniFi to force internal DNS use.
- Using a split-horizon DNS setup (advanced).
- Relying on IPs for remote management until internal DNS works better.
UniFi Identity still gets better, future releases may support internal DNS propagation over Teleport.
Step 6: Manage invite lifecycle
Teleport invites expire in 24 hours and single-use:
- Generate a new invite per device or reissue regularly.
- Within UI Network, go to Devices → Teleport Clients, find the remote client, and click Revoke to force disconnect.
This gives you control: remove access in one click if the device is lost or compromised.
Step 7: Use cases & examples
MSP remote support You set up Teleport and DDNS on a client's UDM. Send invites for your laptop and phone. Weeks later, their Wi-Fi dies, no time to wait. Open your remote app, connect, jump into the dashboard, reboot devices, and fix the issue, without a VPN or extra steps.
Remote home management
You're out camping. Your UDR gets sticky or APs blink offline. You don’t need a VPN. Use Teleport, point to myuniunifi.duckdns.org, connect, and manage cameras or check alerts from your phone.
Roaming network You run a cabin shop with seasonal IP changes. DDNS keeps your host reachable without new IPs each time. Teleport secures access, and you lock down remote access via firewall rules.
Step 8: Troubleshooting
- Home ISP CGNAT: If real public IP isn't assigned, DDNS hits internal IPs that won’t work. Check for CGNAT and consider bypass mode or ask ISP for a public IP.
- Invite expired: Remember invitations are one-time, 24hrs. Easily solve by generating a new one.
- Service unresponsive: Confirm your firewall allows only whitelisted traffic. Drop rules must come after ALLOWs.
- No internal DNS: Use IP or set DHCP static DNS entries until your UniFi version supports internal DNS over Teleport.
Step 9: Scaling remote access
Teleport + DDNS is ideal for small teams or MSPs with a few sites. For larger operations:
- Rotate keys per admin for accountability
- Enforce 2FA for reverse access
- Automate invite generation via API
- Use monitoring and logging to audit usage
- Consider shared resource servers or proxies in cloud
For high-demand or enterprise needs, combine with site-to-site IPSec or managed VPN solutions. But simplicity has its value.
Final thoughts
If you need secure remote controller access without the complexity of traditional VPN, this combo is a winner. It’s easy, secure, and flexible, plus you can lock it down tightly with firewall rules and one-click revocation.
For MSPs or anyone managing multiple sites, self-hosted DDNS plus per-device Teleport setup begins to scale, but it's a bit manual. That’s where UniHosted steps in: every controller we host ships ready with DDNS, Teleport, firewall rules, backups, updates, and monitoring built in. No scripts, no VPN puzzles, just turnkey, scalable remote access.
If you're ready to simplify remote access and skip hosting pain, check out UniHosted.
Related guides
Keep reading
- UniFi Teleport VPN: how to set it up (Windows, Mac and mobile)UniFi Teleport is Ubiquiti's invitation-based WireGuard VPN. How to enable it, invite users, connect from Windows, Mac or a phone, and where it stops. Read guide
- How to optimize UniFi for remote work with Starlink and TeleportSet up a reliable remote work network with Starlink and UniFi by optimizing settings, enabling Teleport VPN, and managing remote access smoothly. Read guide
- How to Access UniFi Controller from Anywhere Without Port ForwardingRemotely manage your UniFi Controller without port forwarding using secure options like Teleport, WireGuard, or reverse proxies. Ask ChatGPT Read guide
