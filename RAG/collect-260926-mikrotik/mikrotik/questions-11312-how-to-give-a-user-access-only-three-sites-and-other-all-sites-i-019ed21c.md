---
id: collect-260926-mikrotik/mikrotik/questions-11312-how-to-give-a-user-access-only-three-sites-and-other-all-sites-i-019ed21c
title: "questions-11312-how-to-give-a-user-access-only-three-sites-and-other-all-sites-i-019ed21c"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-11312-how-to-give-a-user-access-only-three-sites-and-other-all-sites-i-019ed21c.md
source_anchor: ""
source_lines: [1, 27]
sha256: 2ae4bd88d7582600f1ed908760b1eafae2e6cf87430f7dd9a71922dd2c0cada0
---

# questions-11312-how-to-give-a-user-access-only-three-sites-and-other-all-sites-i-019ed21c

Mikrotik router Model :450g
like as,Access lists :
1.www.google.com
2.www.yahoo.com
3.www.banglalinkgsm.com
Denny list : All sites without access listed sites .
You can follow this procedure:
Use Firewall > Filter Rules.
For www.google.com access :
1.add a new rule
2.chain: forward
3.Src.Address : LAN Network IP
4.Dst.Address: www.google.com
5.in.Interface : LAN Port Name
6.action: accept
7.apply >OK
For www.yahoo.com access :
4.Dst.Address: www.yahoo.com
For www.banglalinkgsm.com access :
2.chain:forward
4.Dst.Address: www.banglalinkgsm.com
For Other sites access :
3.in.Interface : LAN Port Name
4.action: drop
5.apply >OK
The solution described above does not work, because www.google.com does not have fixe IP address.
Solution is this: https://wiki.mikrotik.com/wiki/Use_host_names_in_firewall_rules
