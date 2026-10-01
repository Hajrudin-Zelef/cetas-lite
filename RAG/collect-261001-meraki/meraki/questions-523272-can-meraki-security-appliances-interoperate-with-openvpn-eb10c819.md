---
id: collect-261001-meraki/meraki/questions-523272-can-meraki-security-appliances-interoperate-with-openvpn-eb10c819
title: "questions-523272-can-meraki-security-appliances-interoperate-with-openvpn-eb10c819"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-523272-can-meraki-security-appliances-interoperate-with-openvpn-eb10c819.md
source_anchor: ""
source_lines: [1, 24]
sha256: a91db194f9fef11e851d486ac726c22260854e729c6040fb6eb1a50fa41397f1
---

# questions-523272-can-meraki-security-appliances-interoperate-with-openvpn-eb10c819

I know this is a long shot, but has Meraki said anything about interoperating with SSL vpns such as openvpn?
- 
        Ok, please educate me on how to improve. Down votes without any reason aren't as helpful as ones that have them.user67327– user673272013-07-15 17:04:53 +00:00Commented Jul 15, 2013 at 17:04
- 
        I think this was a perfectly reasonable question (I found this thread because I was wondering the same). Maybe the downvotes came from security elitists (the types who would tell people to "RTFM" rather than at least giving a pointer where to find more information).Ville– Ville2017-08-14 19:58:16 +00:00Commented Aug 14, 2017 at 19:58
                    
                        Add a comment
                    
                 | 
            
                
            
        
         
    1 Answer 1
Just going to quote from the docs here:
  Currently, MX series support for third-party VPN interoperability requires the following:
- Preshared keys (no certificates)
- LAN static routes (no routing protocol for the VPN interface)
- Phase 1 (IKE Policy): 3DES, SHA1, DH group 2, lifetime 8 hours
- Phase 2 (IPsec Rule): Any of 3DES, DES, or AES; either MD5 or SHA1; PFS disabled; lifetime 8 hours
So, no, no OpenVPN support. If you need OpenVPN support, I suggest you contact your Meraki rep.
- 
        I'll give this the answer, but I was hoping (and still am hoping) that someone had heard about their plans. I wasn't clear about wanting plans rather than current features. Juniper has an SSL VPN appliance and I was hoping they would ante up to compete.user67327– user673272013-07-15 02:08:27 +00:00Commented Jul 15, 2013 at 2:08
