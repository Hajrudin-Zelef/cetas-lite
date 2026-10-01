---
id: collect-261001-general-networking/general-networking/questions-1838556-adguardhome-dns-not-working-for-wireguard-vpn-clients-on-opnse-192d3df8
title: "AdguardHome DNS not working for Wireguard VPN clients on OpnSense"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1838556-adguardhome-dns-not-working-for-wireguard-vpn-clients-on-opnse-192d3df8.md
source_anchor: ""
source_lines: [1, 12]
sha256: ccdaae7a3d92986a8472e586655a8a749886f23327c85a5ce26dfa7bfb72cfaa
---

# AdguardHome DNS not working for Wireguard VPN clients on OpnSense

*Score : 2 | Source : https://superuser.com/questions/1838556/adguardhome-dns-not-working-for-wireguard-vpn-clients-on-opnsense*

For my home network I have a x86 System running OpnSense. I use Wireguard to connect external devices with my network (set up according to the official documentation). This setup worked fine until I installed the AdGuard Home plugin on OpnSense. Now VPN devices cannot resolve host names any more. While I do see the requests from this devices in the AdGuard Home log (with response code NOERROR and the correct IP addresses) the responses do not seem to reach their destination. Using a dedicated AdGuard Home server running on a different system than OpnSense does work though.

---

### Reponse — score 1

Not sure if this is the correct/best solution, but I found out that when using the dedicated server (in my case 192.168.0.226) the actual packets where addressed from 192.168.0.266 to 192.168.1.2 (IP of the VPN client). But when using the AdGuard plugin, the packets where addressed from 192.168.1.1 (instead of 192.168.0.1) to 192.168.1.2.
So changing the bind_hosts from 0.0.0.0 to 192.168.0.1 in the AdGuard Home configuration file solved my problem.
