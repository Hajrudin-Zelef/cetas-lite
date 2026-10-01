---
id: collect-261001-general-networking/general-networking/manual-how-tos-nat-reflection-html-7747e718-2
title: "manual-how-tos-nat-reflection-html-7747e718"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-how-tos-nat-reflection-html-7747e718.md
source_anchor: ""
source_lines: [58, 107]
sha256: 854389c1ee292b484f42da7fb4c5e804b08100baece73474ba9f34484ebdf853
---

# manual-how-tos-nat-reflection-html-7747e718

- Select Hybrid Source NAT rule generation and save. That way you can have manual Source NAT rules in conjunction with automatic IP-Masquerading rules. You could also choose Manual Source NAT rule generation. Please make sure that you create your own IP-Masquerading rules with manual Source NAT enabled. Select + to create a new Source NAT rule. Interface: Select DMZ - It’s the interface of the subnet the Webserver is in.Protocol: Select TCPSource Address: Select DMZ net - It’s the alias for the DMZ Network172.16.1.0/24Source Port: Select AnyDestination Address: Input 172.16.1.1 - It’s the Webserver’s internal IPv4 address in the DMZ.Destination Port: Input 443 - Or select the aliasHTTPSTranslation/target: Select DMZ address - It’s the alias for the OPNsense Interface IPv4 address172.16.1.254 in the DMZ Network.Description: Input Hairpin NAT Rule Webserver 443
Tip
Reading the SNAT rule like a sentence makes it clearer:
If a packet is received by the OPNsense on the interface DMZ with protocol TCP from the source net 172.16.1.0/24 and the source port ANY to destination IP 172.16.1.1 and destination port 443 –> rewrite the source ip to 172.16.1.254 and answer from the OPNsense DMZ interface.
Note
Now all DMZ clients (and the Webserver itself) can reach the Webserver with its external IP.
- You need this additional SNAT rule to avoid asymmetrical traffic between clients and servers in the same layer 2 broadcast domain. TCP traffic won’t work otherwise.
Repeat Method 1 until all additional servers are reachable.
If you encounter any issues, check Troubleshooting NAT Rules for a few tips.
Warning
The following methods are not advised, but are still explained in order to prevent misconfigurations. There is more information in (Advanced) Settings.
Method 2 - Creating Automatic Port-Forward NAT (DNAT), Manual Source NAT (SNAT), and Manual firewall rules
- Go to
- Enable Reflection for Destination NAT (Port Forwards) to create automatic rules for all entries that have WAN as interface.
- Go to
- Create the NAT rule as in Method 1 - Destination NAT (Port Forward) but change the following things: 
  - Make sure that your Destination NAT (Port Forwarding) rule specifies only WAN as interface.
- Go to
- Action: Select PassInterface: Select WAN ,DMZ andLAN - Select all interfaces in which clients are that should access the webserver.Protocol: Select TCPSource: Select AnyDestination: Input 172.16.1.1 - It’s the Webserver’s internal IPv4 address in the DMZ. NAT matches before firewall.Destination port range: Input 443 - Or select the aliasHTTPSDescription: Input Reflection NAT Rule Webserver 443
- Go to
- Create the NAT rule as in Method 1 - Source NAT
Method 3 - Creating Automatic Port-Forward NAT (DNAT), Automatic Source NAT (SNAT), and Manual firewall rules
- Go to
- Enable Reflection for Destination NAT (Port Forward)s to create automatic rules for all :menuselection: Firewall –> NAT –> Destination NAT (Port Forward) that have WAN as interface.
Enable Automatic Source NAT (Outbound) for Reflection to create automatic SNAT rules.
- Go to
- Create the NAT rule as in Method 2 - Destination NAT (Port Forward)
- Go to
- Create the floating firewall rule as Method 2 - Floating
One-to-One NAT Reflection
When Reflection for 1:1 is activated, automatic Reflection NAT rules for all One-to-One NAT rules are generated.
If you want to create manual Reflection and Hairpin NAT rules, leave Reflection for 1:1 disabled and follow the steps in Method 1. The only change is not adding the WAN interface to the Destination NAT (Port Forward) rules you create. The resulting Destination NAT (Port Forward) and Source NAT rules are in addition to the existing One-to-One NAT rules.
If your Destination NAT (Port Forward) rule has 1 interface selected (e.g. LAN), the resulting Filter rule association: Add associated filter rule will appear in . If you have more than 1 interface selected, it will appear in Firewall –> Rules –> Floating.
Troubleshooting NAT Rules
Tip
- Open SSH shell:
- Display all loaded and active NAT rules:
- pfctl -s nat
- “rdr” means rules.
- “nat” means rules.
- You can also check the rules in the GUI in
Tip
- Displays all NAT rules in the OPNsense debug:
- cat /tmp/rules.debug | grep -i nat
- If there are more rules here than in pfctl -s nat , it means you forgot to hit apply somewhere.
Tip
- Look at the default drops of the firewall live log in
- Turn on logging of the NAT and Firewall rules you have created, and check if they match in . NAT rules have the label “NAT” or “RDR”. Firewall rules have their description as label.
- In “ you can check if there is a session between your internal client and your internal server, and which rule matches to it.
- Use tcpdump on the client, the opnsense and the server, and test if the traffic goes back and forth between the devices without any mistakes. Look for TCP SYN and SYN ACK. If there are only SYN then the connection isn’t established and there are mistakes in your rules.
