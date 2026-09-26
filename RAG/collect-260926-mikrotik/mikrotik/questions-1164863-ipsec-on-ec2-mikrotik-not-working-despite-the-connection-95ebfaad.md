---
id: collect-260926-mikrotik/mikrotik/questions-1164863-ipsec-on-ec2-mikrotik-not-working-despite-the-connection-95ebfaad
title: "questions-1164863-ipsec-on-ec2-mikrotik-not-working-despite-the-connection-95ebfaad"
domain: mikrotik
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/lot-mikrotik/forum/ipsec/questions-1164863-ipsec-on-ec2-mikrotik-not-working-despite-the-connection-95ebfaad.md
source_anchor: ""
source_lines: [1, 12]
sha256: 35b4edf346dfeba89862080b7ffc482690f84b8619af714616fd9dd3fa57ade9
---

# questions-1164863-ipsec-on-ec2-mikrotik-not-working-despite-the-connection-95ebfaad

I have 2 mikrotiks listed on AWS:
R1 - public IP 3.75.170.246 and subnet 172.168.0.0/24
R2 - public IP 18.199.145.214 and subnet 10.0.0.0/24
On AWS, in the security group I have allow all on all protocols and ports and EC2 also has source/destination disabled in EC2. I have the tunnel connected correctly, at least I think so, because the connection between them has the status "Established". The problem is that I can't ping between them at all via the private IP, via the public one, of course, without a problem. I did everything according to this guide https://www.youtube.com/watch?v=gn-uF_xQECQ When I want to ping from R2 to R1 via the private IP, I get "Connection Timeout" and I don't know why, as the firewall allows all protocols on all ports, which can be seen in the screenshots. Honestly, I don't know what the problem is, especially since the tunnel is done. I suspect that I probably have something screwed up with the Firewall, but there is one rule in it and that's it.
The problem is that the ping does not reach the second router and I get "Timeout". I've watched several guides on how to configure it, but they all look the same.
Can someone help me? I know I can use AWS IPSec to bridge the tunnel between regions, but that's not the point.
Now I changed IP 172.168.0.0 to 172.16.0.0 but it didn't help
R1 IPSec Policy R2 IPSec Policy
Installed SAs
Solved!
Problem was with IP mask, I used /24 when AWS VPC has /16
172.168.0.0is outside the RFC1918 scope, and you should not use it unless it's actually assigned to you.
