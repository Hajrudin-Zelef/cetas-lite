---
id: collect-260926-mikrotik/mikrotik/questions-1112515-limit-access-to-specific-host-in-my-network-over-vpn-in-mikrot-cf5fb0b7
title: "questions-1112515-limit-access-to-specific-host-in-my-network-over-vpn-in-mikrot-cf5fb0b7"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vpn/questions-1112515-limit-access-to-specific-host-in-my-network-over-vpn-in-mikrot-cf5fb0b7.md
source_anchor: ""
source_lines: [1, 15]
sha256: 0358cf228479e7c6a53be0d731aeeb966bbe0ccb05a58e8dfe22e9fdf4de370a
---

# questions-1112515-limit-access-to-specific-host-in-my-network-over-vpn-in-mikrot-cf5fb0b7

I have a Mikrotik RB2011, the mikrotik it's configured to act as a VPN Server and it's working, I need the enable to external consultant access to only one PC of my network, no other PC or IP, Please help me to reconfigure my VPN in Mikrotik Thanks
2 Answers 2
Below is the command line config for a MikroTik router. You can copy these lines in a MikroTik CLI to create a PPP user with limited access to servers.
/ppp profile add address-list=VPN_USER_client local-address=10.15.32.33 name=USER remote-address=ovpn-lan
/ip firewall address-list add address=10.0.0.10 list=VPN_USER_server
/ip firewall filter
add action=jump chain=forward jump-target=VPN_USER src-address-list=VPN_USER_client
add action=accept chain=VPN_USER dst-address-list=VPN_USER_server dst-port=3389 protocol=tcp
add action=drop chain=VPN_USER
- 
        It's the comand line config for a MikroTik router. You can copy these lines in a MikroTik CLI to create a PPP user with limited access to servers.Jānis Šteninbergs– Jānis Šteninbergs2017-07-14 07:50:23 +00:00Commented Jul 14, 2017 at 7:50
- 
        I added that to your answer so just know that including such detail in answers is important for a reader so they fully understand what they need to do with such content, etc.Vomit IT - Chunky Mess Style– Vomit IT - Chunky Mess Style2017-07-14 13:58:16 +00:00Commented Jul 14, 2017 at 13:58
This can't be done in VPN Configuration of Mikrotik. Add a filter rule in Mikrotik Firewall that allows traffic from VPN IP address of your consultant to the specified PC. Then add a filter rule which denies any access from consultant's VPN IP address to anywhere.
Consultant's VPN IP address is the IP address which Mirotik gives to VPN connection of your consultant. You can configure VPN server to give a specific IP address to a specific VPN username and password.
