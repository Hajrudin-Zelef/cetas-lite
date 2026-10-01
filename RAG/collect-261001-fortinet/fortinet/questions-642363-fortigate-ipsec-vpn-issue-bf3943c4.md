---
id: collect-261001-fortinet/fortinet/questions-642363-fortigate-ipsec-vpn-issue-bf3943c4
title: "questions-642363-fortigate-ipsec-vpn-issue-bf3943c4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-642363-fortigate-ipsec-vpn-issue-bf3943c4.md
source_anchor: ""
source_lines: [1, 14]
sha256: 82fc8350492579963fa7032d082dffd2ff250edd44f20569b215417cb43a0e77
---

# questions-642363-fortigate-ipsec-vpn-issue-bf3943c4

Have a challenging question here.
We have a Fortigate 620B which we're trying to use to route some traffic over a VPN tunnel to a customer.
We want the traffic to go out of our interface with one of our public IPs (we have it set to NAT the address using a specific public IP address) to a public IP on the client end.
I've got a configuration that I believe SHOULD work except the traffic as it hits the link says that it's coming from our internal IP address (192.168.X.X) instead of the public IP address (x.x.x.x).
The error we get in the logs:
id=13 trace_id=368 msg="vd-root received a packet(proto=6, 192.168.XX.XX:50470->XX.XX.183.94:443) from port8. flag [S], seq 342573222, ack 0, win 8192"
id=13 trace_id=368 msg="Find an existing session, id-0a2bb411, original direction"
id=13 trace_id=368 msg="enter IPsec interface-XXX_P2P"
id=13 trace_id=368 msg="No matching IPsec selector, drop"
It drops the packet because the quick mode selector on the VPN is set to use our public IP instead of our private IP.
The challenge here is that us and our client use the same private IP space so we have to NAT both end of the traffic.
To get around that I have one Phase 1 proposal and two Phase 2 Proposals.
The routing rules work because the static route guides the packet to the VPN interface and the policy rules work but for some reason the NAT isn't working properly.
I'm happy to provide any information I can to help troubleshoot.
