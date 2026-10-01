---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-fsfz7a-one-way-ping-problem-sslvpnipsec-site-b6868e4c
title: "r-fortinet-comments-fsfz7a-one-way-ping-problem-sslvpnipsec-site-b6868e4c"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-fsfz7a-one-way-ping-problem-sslvpnipsec-site-b6868e4c.md
source_anchor: ""
source_lines: [1, 32]
sha256: 4942d503897dae076cb9083153c551bd3691ce592ba75cb05d9a24c2eac2076e
---

# r-fortinet-comments-fsfz7a-one-way-ping-problem-sslvpnipsec-site-b6868e4c

One way ping problem SSLVPN+IPSec Site 
        
        
        
    
    
    I'm experiencing a one way ping problem with my Forticlient users, who connect to a remote site's SSL VPN.
I'm at HQ, and i'm able to ping the entire remote site, except SSL VPN users.
When I RDP into a server at the remote site, I can ping that SSL VPN user just fine.
When i'm connected to the remote site SSL VPN, I can ping the HQ Subnet just fine.
- 
      SSLVPN User -> Remote Site FW -> IPSEC Tunnel -> HQ FW -> Local User (Pings OK)
- 
      Local User -> HQ FW -> IPSEC Tunnel -> Remote Site FW -> SSLVPN User (Request timed out)
- 
      Local User -> HQ FW -> IPSEC Tunnel -> Remote Site FW -> Remote site user (Pings OK)
- 
      Remote site user -> Remote Site FW -> SSLVPN User (Pings OK)
What type of configuration error would cause this?
Section des commentaires
Firewall configuration:
Remote site (in sequence):
You’re missing the auth group as destination
EDIT: I’m an idiot sandwhich
As far as I can tell, setting a group as a destination isn't possible.
And setting a group and a source isn't possible when action is IPSEC
Policy-based IPsec? In the SSL-VPN-->wan1 policy with IPsec action, check in CLI if you allow reverse direction as well.
That was it. Thanks!
Do you have a firewall rule that allows your Local User to communicate to the SSL VPN (ssl.root) interface?
See my reply
https://www.reddit.com/r/fortinet/comments/fsfz7a/one_way_ping_problem_sslvpnipsec_site/fm16fqx/
Routing, phase2 selectors, firewall policies. It's rarely anything else.Except when it's policy-based IPsec.Check all on each hop.
