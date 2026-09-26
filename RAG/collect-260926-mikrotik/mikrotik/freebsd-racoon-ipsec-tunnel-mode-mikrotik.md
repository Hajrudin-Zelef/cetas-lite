---
id: collect-260926-mikrotik/mikrotik/freebsd-racoon-ipsec-tunnel-mode-mikrotik
title: "freebsd-racoon-ipsec-tunnel-mode-mikrotik"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/freebsd-racoon-ipsec-tunnel-mode-mikrotik.md
source_anchor: ""
source_lines: [1, 87]
sha256: 59296c6e2ec53b38369119d3ac71cd51d549920ed1c4140ef39054c1a2da3d3a
---

# freebsd-racoon-ipsec-tunnel-mode-mikrotik

OK, so let’s start by a good news - you can forget all what I wrote about the incompatibility in my first response, as I have confused a **gif** tunnel with a **vti** one, due to the title of the topic title which mentions IPsec tunnel mode alone.

According to this article, in its IPv4 over IPv4 mode, the FreeBSD’s **gif** (referring to RFC 2893, which is an evolution of RFC 1933, as an inspiration) is compatible with the **ipip** tunnel, standardized by RFC 2003. No cross-reference between these two RFCs can be found in either of them, however the packet format differs in just the payload protocol type (4 for IPv4/RFC 2003 and 41 for IPv6/RFC 2893); I didn’t dive into all those ICMP and TTL handling details and loop prevention mechanisms.

So given that both A.B.C.D and W.X.Y.Z are directly routable (no NAT between them), it would be best to first check this layer of compatibility alone, before diving into IPsec.

At Mikrotik side, you add the tunnel interface using

*/interface ipip add name=freebsd-gif0 local-address=W.X.Y.Z remote-address=A.B.C.D !keepalive*

The “logical addresses” as FreeBSD refers to them may be added using

*/ip address add address=192.168.2.1/32 network=192.168.1.1/32 interface=freebsd-gif0*



For GRE, you would just use */interface gre* instead of */interface ipip* in the step above (at RouterOS side). Except the IP protocol to be matched by the policy (the tunneling protocol itself), the IPsec layer settings will be the same regardless whether the inner tunnel used will be GIF/IPIP or GRE.

In RouterOS configuration, a Phase 1 proposal mostly matches an */ip ipsec profile*. Hence

```
remote W.X.Y.Z {
        nat_traversal off;
        exchange_mode main;
        proposal {
                encryption_algorithm des;
                hash_algorithm sha1;
                authentication_method pre_shared_key;
                dh_group 2;
        }
}
```

translates into

*/ip ipsec profile add name=FreeBSD enc-algorithm=des hash-algorithm=sha1 nat-traversal=no
/ip ipsec peer add name=FreeBSD address=A.B.C.D local-address=W.X.Y.Z exchange-mode=main profile=FreeBSD
/ip ipsec identity add peer=FreeBSD auth-method=pre-shared-key secret=verycomplexsecret*

Phase 2 proposal is called */ip ipsec proposal*, so

```
sainfo anonymous {
      pfs_group 2;
      lifetime time 1 hour;
      encryption_algorithm des;
      authentication_algorithm hmac_sha1;
      compression_algorithm deflate;
}
```

translates into

*/ip ipsec proposal add name=FreeBSD pfs-group=modp1024 lifetime=1h enc-algorithms=des auth-algorithms=sha1*

There is no way to control use of compression algorithm in RouterOS so I suspect it is not used at all. It seems that it must be specified it in Racoon configuration but it is used only as an proposed option, not a mandatory requirement.

Lastly, in RouterOS, the policy for both SAs is configured using a single line, where the *src-address* represents the local IP or subnet and the *dst-address* represents the remote one. So

```
spdadd A.B.C.D/32 W.X.Y.Z/32 ipencap -P out ipsec esp/tunnel/A.B.C.D-W.X.Y.Z/require;
spdadd W.X.Y.Z/32 A.B.C.D/32 ipencap -P in ipsec esp/tunnel/W.X.Y.Z-A.B.C.D/require;
```

translates into

*/ip ipsec policy add src-address=W.X.Y.Z dst-address=A.B.C.D ipsec-protocols=esp tunnel=yes proposal=FreeBSD protocol=ipencap level=require peer=FreeBSD*

However, as you have public IP addresses at both machines and the gif/ipip/gre tunnels are linked to them (A.B.C.D, W.X.Y.Z), you can save some packet space by using a *transport* mode of IPsec - to do that, you need to change the above to

*spdadd A.B.C.D/32 W.X.Y.Z/32 ipencap -P out ipsec esp/tunnel**transport**/A.B.C.D-W.X.Y.Z/require;
spdadd W.X.Y.Z/32 A.B.C.D/32 ipencap -P in ipsec esp/tunnel**transport**/W.X.Y.Z-A.B.C.D/require;*

and

*/ip ipsec policy add src-address=W.X.Y.Z dst-address=A.B.C.D ipsec-protocols=esp tunnel=yes**no** proposal=FreeBSD protocol=ipencap level=require peer=FreeBSD*



For GRE, the protocol field would be set to *gre* rather than *ipencap*.




This part is a bit confusing. In Mikrotik IP address configuration, the *network* parameter of an */ip address* row must either be the subnet into which the IP address itself fits, or a /32 address representing the remote end of the tunnel to the routing (so that you could specify it as a *gateway* of a route rather than the tunnel interface itself). So to reach the whole 192.168.1.0/24 via the tunnel, you have to manually add a route:

*/ip route add dst-address=192.168.1.0/24 gateway=192.168.1.1*
