---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1199728-hostnames-of-my-unifi-network-equipment-40a71bb0
title: "questions-1199728-hostnames-of-my-unifi-network-equipment-40a71bb0"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1199728-hostnames-of-my-unifi-network-equipment-40a71bb0.md
source_anchor: ""
source_lines: [1, 34]
sha256: 9f4385e20ff91a2aa3780027e66e9c464c45c8cf9ebc93874d02482b4a0cca8a
---

# questions-1199728-hostnames-of-my-unifi-network-equipment-40a71bb0

I have a UDM running Unifi Network. It's connected to a little more than 10 switches and 10 APs (all Unifi). I set the IPs to static but when I do that, and I scan the network segment with Advanced IP Scanner, it doesn't return any hostnames. It only displays the IP address. Is there a way to set the hostname of Unifi switches and APs? Setting the Name field doesn't do it.
- 
        1Why do you expect an IP scanner to show the hostname? Where should it get it?vidarlo– vidarlo2026-08-27 21:34:39 +00:00Commented Aug 27 at 21:34
- 
            
            
- 
        2Do you know what DNS server is?Romeo Ninov– Romeo Ninov2026-08-28 04:26:15 +00:00Commented Aug 28 at 4:26
- 
        When you say "Advanced IP Scanner", do you mean this?jcaron– jcaron2026-08-31 09:48:27 +00:00Commented Aug 31 at 9:48
                    
                        Add a comment
                    
                 | 
            
                
            
        
         
    1 Answer 1
For the scanner, its only source for hostnames is reverse DNS. Accordingly, set up PTR records for the IP addresses in your DNS.
- 
        2Alternatively: Have your DHCP server (which is probably the UDM, which also does the DNS), have DHCP reservations for the devices. (Even if you give them a manual static ip-address.) That registers them in DHCP and DNS (so the UDM know they exist) and if for some reason one of them false back on using DHCP in stead of static it will get the same ip-address. (Actually I wouldn't even bother with static ip-addresses. Having a central administration of the ip-addresses in DHCP is so much nicer than having to update X devices individually.)Tonny– Tonny2026-08-28 12:50:26 +00:00Commented Aug 28 at 12:50
- 
            
            
- 
        If I set the switch to DHCP it does get an IP and my scanner shows the hostname. If I make it static it doesn't. I can do an IP reservation for a client PC but there doesn't seem to be a way to do that for infrastructure devices (switches and APs). At least not that I, or AI could find.Jeff– Jeff2026-08-28 13:06:55 +00:00Commented Aug 28 at 13:06
- 
        2There's no difference between infrastructure devices and PC's. That's simply some strange abstraction that unifi does, probably because they think it's a bad idea. As the vendor thinks it's a bad idea, I would probably avoid doing it.vidarlo– vidarlo2026-08-28 13:11:30 +00:00Commented Aug 28 at 13:11
- 
        @Jeff Most often, DHCP reservations hinge on the device (base) MAC address.Zac67– Zac672026-08-28 13:11:59 +00:00Commented Aug 28 at 13:11
- 
        2@jeff DHCP even with a static reservation, will likely update the DNS server. Straight static IPs can't and won't tell the DNS server anything.Criggie– Criggie2026-08-29 21:23:20 +00:00Commented Aug 29 at 21:23
