---
id: collect-261001-general-networking/general-networking/questions-2796-problem-with-remote-access-fixed-l2l-and-dynamic-l2l-ipsec-workin-6926e62a
title: "questions-2796-problem-with-remote-access-fixed-l2l-and-dynamic-l2l-ipsec-workin-6926e62a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-2796-problem-with-remote-access-fixed-l2l-and-dynamic-l2l-ipsec-workin-6926e62a.md
source_anchor: ""
source_lines: [1, 44]
sha256: 3b6c65c71e1f53f23b0008daa6d99b3774b8c0f4fa28e45a1381350e10c7998d
---

# questions-2796-problem-with-remote-access-fixed-l2l-and-dynamic-l2l-ipsec-workin-6926e62a

I have an issue with a mixture of remote access, L2L and dynamic L2L tunnels on an ASA5540 running 8.2
Here is a snip of the relevant config:-
crypto dynamic-map outside-crypto-dynamic-map 10 match address outside-crypto-dynamic-map-10
crypto dynamic-map outside-crypto-dynamic-map 10 set transform-set ESP-3DES-MD5
crypto dynamic-map outside-crypto-dynamic-map 20 set transform-set ESP-3DES-MD5
crypto map outside-crypto-map 201 match address outside-crypto-map-201
crypto map outside-crypto-map 201 set peer X.X.X.X
crypto map outside-crypto-map 201 set transform-set ESP-3DES-MD5
crypto map outside-crypto-map 202 match address outside-crypto-map-202
crypto map outside-crypto-map 202 set peer Y.Y.Y.Y
crypto map outside-crypto-map 202 set transform-set ESP-AES256-SHA
crypto map outside-crypto-map 65535 ipsec-isakmp dynamic outside-crypto-dynamic-map
crypto map outside-crypto-map interface outside
I have a number of remote sites that use dynamic IPs. The LAN subnets for these are in an ACL "outside-crypto-dynamic-map-10".
These match fine based on this line:-
crypto dynamic-map outside-crypto-dynamic-map 10 match address outside-crypto-dynamic-map-10
I have other "static" L2L tunnels that work fine per 201 and 202 in the config above.
With my remote access (Cisco VPN Client) users, unless I have the following line, they won't connect:-
crypto dynamic-map outside-crypto-dynamic-map 20 set transform-set ESP-3DES-MD5
If I try to add a "match address" statement to that sequence (like below), then remote access stops working (where "vpc-client-subnet" is an ACL containing the subnet from the IP pool used for the remote access clients) as it cannot find a matching local/remote.
crypto dynamic-map outside-crypto-dynamic-map 20 match address vpc-client-subnet
The problem is that I have a number of L2L endpoints out there (that I no longer want connected but have no control over) that are still configured (on the remote end) with the PSK used by the dynamic L2L peers. I have removed their LAN subnets from "outside-crypto-dynamic-map-10" (as I don't want them connected any more) so they don't match on the dynamic-map seq 10 anymore, however they are still able to successfully complete phase 2 against "outside-crypto-dynamic-map 20" and it seems as if this end just accepts whatever the remote proposes as the local/remote for the SA.
I am not in a position to change the PSK. I can't add a "match address" to the "outside-crypto-dynamic-map 20" as it will prevent remote access clients from connecting, however if I don't, then it acts as a catch-all for other peers that otherwise know the PSK.
Ideally I could add "match address" to seq 20, so that remote access users will match against it, but the "old" seq 10 peers will not. Alternatively, is there a way to prevent the ASA from accepting the remote peers idea of the local/remote for the SA when it has matched against a map that doesn't have a "match address" statement?
EDIT:
The relevant tunnel group config:-
tunnel-group DefaultL2LGroup ipsec-attributes
 pre-shared-key *****
 peer-id-validate nocheck
 isakmp keepalive threshold 30 retry 5
tunnel-group x.x.x.x type ipsec-l2l
tunnel-group x.x.x.x ipsec-attributes
 pre-shared-key *****
 isakmp keepalive threshold 30 retry 5
tunnel-group ravpn type remote-access
tunnel-group ravpn general-attributes
 address-pool ip-pool-ravpn
 authentication-server-group Edirectory
 default-group-policy ravpn
tunnel-group ravpn ipsec-attributes
 pre-shared-key *****
 isakmp keepalive threshold 30 retry 2
tunnel-group-map default-group DefaultL2LGroup
There are multiple tunnel-groups for the "static" tunnels (per the x.x.x.x example above). The remote access users hit the "ravpn" tunnel-group, and the dynamic ones all match the "DefaultL2LGroup". The PSK is not the same between these.
