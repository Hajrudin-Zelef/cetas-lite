---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-ncvr4l-vpn-ipsec-tunnel-74c6a424
title: "r-fortinet-comments-ncvr4l-vpn-ipsec-tunnel-74c6a424"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-ncvr4l-vpn-ipsec-tunnel-74c6a424.md
source_anchor: ""
source_lines: [1, 42]
sha256: a0addb430b7990b4487fbdb84faedce89ef0c0363a50904d4057f394b147bbeb
---

# r-fortinet-comments-ncvr4l-vpn-ipsec-tunnel-74c6a424

VPN IPSec Tunnel 
        
    Hello everyone,
I was given a task to set up an lab on my VMWare to configure the VPN IPSec Tunnel on Fortigate-VM64 using 7.0. I have used almost all the resources on google and Youtube, but I am still unable to get it working. I am supposed to get a Remote Client to connect via FortiClient to access an internal network.
After setting up the tunnel using the IPSec Wizard, I will be automatically given a private IP for the VPN, Tunnel interface (usually 169.254.x.x) and when I try to create a static route for it, I will get an error saying, "Gateway IP undefined could be unreachable, It is not in any subnet of the interface VPN". So I have tried many different IP Addresses, but will still always get this error.
According to my VPN Event logs, I would always get this " progress IPsec phase 1 failure" and " IPsec phase 1 error". I have checked my phase 1 and phase 2 settings and they are exactly the same so I googled and researched for other solutions, but none of them seem to be able to stop this same error.
I will really appreciate any help given at all because honestly, my boss is about to kick my ass out the door. Thank you so much in advance.
Section des commentaires
Hi everyone,
Thanks for your help! I really appreciate it! I have attached my configuration for your reference.
https://imgur.com/lbKrqnc
https://imgur.com/vV6VdFA
https://imgur.com/83Peu5S
https://imgur.com/mIfm0GP
https://imgur.com/qnFxsZv
https://imgur.com/9ERw82X
I just labbed this up and you didn't follow the link. Your phase 2 selectors should be 0.0.0.0 on both sides after the wizard is done.
The debug output would have told you that your phase 2 is the problem by the way.
Hi! thanks for trying!
I changed the phase 2 selectors to 0.0.0.0 but I am still getting the same error code "IPSec Phase 1 interface down"
https://imgur.com/sMBXsBj
https://imgur.com/VX4L73w
https://imgur.com/43nB6Cz
https://imgur.com/L2TU6ff
Edit: I have seen a few other topics on the VPN related issues due to the eval version (which is what I am using). Would that be the same for this instance?
https://docs.fortinet.com/document/fortigate/6.0.0/cookbook/589121/ipsec-vpn-with-forticlient
This has a simple IPsec configuration for clients. Just delete everything you've done and start fresh.
The fact that you are getting an APIPA address tells me you either messed up somewhere or you actually want to do that, which I think isn't the case.
Hello,
Thanks for responding. I restarted and followed the steps in the link. However, I still get the error "IPSec Phase 1 interface down" and the IP Address given for the tunnel was still 169.254.x.x.
Can't say much more without seeing your config on both sides if you have followed the link.
Doubt it's an issue with 7.0 however, which you shouldn't run to begin with.
Is the destination IP of the remote gateway — the “public” IP of the other appliance — reachable?
What 169 addresses are assigned to each? That’s about the only thing that should differ in the configs — those are the private IPs used after the tunnel handshake from the external interfaces of the appliances.
Something isn’t matching up on both sides.
Cipher suites? Lifetime? IP addresses (Phase 2 selectors)? Everything must match up on both sides. In my experience, IPSec just works unless something doesn’t match up, or it’s an invalid config to begin with.
Start with Phase 1. If that doesn’t work, something foundational is wrong, and that should be the most simple part.
Since it’s a lab, can you share more info? Configs, network addresses, log events, etc.
The logs should tell you at least something about why Phase 1 isn’t working. (Maybe I’m wrong and a debug session is necessary.)
Both my phase 1 and phase 2 options are exactly the same hence I really have no idea why it can't connect to Phase 2.
Post your config so we can have a look
Did you intentionally set this up to assign an IP in the APIPA range? If not, that could suggest some issue. IIRC I've seen SSL-VPN end up with an APIPA IP if the IP to be assigned conflicts with a subnet of another interface on the client PC. Maybe the same happens with IPsec.
