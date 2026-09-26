---
id: collect-260926-mikrotik/mikrotik/questions-994803-how-to-forbid-dhcp-server-to-lease-specific-ip-address-on-mikro-c621e6c0
title: "questions-994803-how-to-forbid-dhcp-server-to-lease-specific-ip-address-on-mikro-c621e6c0"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-994803-how-to-forbid-dhcp-server-to-lease-specific-ip-address-on-mikro-c621e6c0.md
source_anchor: ""
source_lines: [1, 25]
sha256: f92f9c2056416454e7e07a76de8035f58844c139dbfc3f056d63a5d0054b8ced
---

# questions-994803-how-to-forbid-dhcp-server-to-lease-specific-ip-address-on-mikro-c621e6c0

My Mikrotik router is running a DHCP server in /24 network, I would like to prohibit it to lease some IP addresses (I would like to use them statically on a couple of devices). Are there any ways to do it?
- 
        2Usually router devices allow to specify the range of addresses assigned by DHCP.Piotr P. Karwasz– Piotr P. Karwasz2019-12-08 17:30:02 +00:00Commented Dec 8, 2019 at 17:30
- 
            
            
- 
        That in not quite what I meant unfortuantly. The idea was to exclude several addresses form the ip range to assign.Ivan– Ivan2019-12-08 19:26:33 +00:00Commented Dec 8, 2019 at 19:26
- 
        1You can assign those static addresses from outside the DHCP range.Piotr P. Karwasz– Piotr P. Karwasz2019-12-08 19:38:48 +00:00Commented Dec 8, 2019 at 19:38
- 
        Setting of range did the job. Thanks.Ivan– Ivan2019-12-08 21:38:15 +00:00Commented Dec 8, 2019 at 21:38
                    
                        Add a comment
                    
                 | 
            
                
            
        
         
    2 Answers 2
- you can make several pools and use next-pool to exclude some addresses
- make static lease for this ip addres to nonexistent mac adress
I've solved the task as it was suggested in the comments: by creating a default-dhcp address pool. The "Pool" section could be found via Winbox software: IP -> Pool submenu.
