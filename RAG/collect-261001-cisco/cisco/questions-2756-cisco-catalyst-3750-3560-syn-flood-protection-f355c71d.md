---
id: collect-261001-cisco/cisco/questions-2756-cisco-catalyst-3750-3560-syn-flood-protection-f355c71d
title: "questions-2756-cisco-catalyst-3750-3560-syn-flood-protection-f355c71d"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-2756-cisco-catalyst-3750-3560-syn-flood-protection-f355c71d.md
source_anchor: ""
source_lines: [1, 13]
sha256: 70b2d08a9e80480dc7d2a9eba3422add1bd9e4d4312aaf16fd762de716d194ea
---

# questions-2756-cisco-catalyst-3750-3560-syn-flood-protection-f355c71d

How could one mitigate SYN FLOOD DOS on Catalyst 3750/3560 as it has no control plane protection?
1 Answer 1
3750 does have some internally priority on what it prefers to punt when congested, but it's not configurable.
So you should rely on common best practices, that is on all your network edges you should have iACL (infrastructure ACL). In iACL you'd allow UDP highports, ICMP to infrastructure network addresses and drop rest. This way ping and traceroute work, but infrastructure cannot be attacked.
iACL should be complemented by policing the allowed traffic to small acceptable rates.
This way when external party is attacking addresses on your 3750, it'll be dropped by network border in the edge.
iACL usually is 100% static so it's low maintenance, as it'll only include infrastructure addresses (loopback, core links).
This will still leave wide open cases where your router is facing customer LAN directly, like when LAN is 192.0.2.0/24 and 3750 has 192.0.2.1 then usually 192.0.2.1 would not be covered by iACL and can be attacked.
Solution for those devices is either to invest on device with proper CoPP capabilities or maintain dynamic iACL always adding the router's customer facing address there.
If you only face customers via link-networks (/30 or /31) solution is much cleaner, you just omit advertising the link-network and add static /32 route for the CPE side, this way external to this router parties cannot attack the router, as they won't have route.
Alternative solution to same issue is to use non-continuous ACL entry in iACL, if your CPE link-network is 198.51.100.0/24 in iACL you could do 'deny ip any 198.51.100.0 0.0.0.254' then all the even addresses would be allowed and odd addresses denied, so if CPE is even and 3750 is odd, all current and future links are protected without updating iACL.
- 
        1Pretty sure you meant 198.51.100.0 0.0.0.254Marco Marzetti– Marco Marzetti2013-08-13 07:08:41 +00:00Commented Aug 13, 2013 at 7:08
