---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-ee37d35a-11
title: "cgi-man-cgi-ee37d35a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-ee37d35a.md
source_anchor: ""
source_lines: [1363, 1510]
sha256: c641a7820865b4e42be2b0d7b1d780ce6f82a6a37da7c4a538c29110aa802518
---

     # Translate incoming packets' destination addresses.
     # As an example, redirect a TCP and UDP port to an internal machine.
     rdr on $ext_if inet proto tcp from any to ($ext_if) port 8080 \
	   -> 10.1.2.151 port 22
     rdr on $ext_if inet proto udp from any to ($ext_if) port 8080 \
	   -> 10.1.2.151 port 53
     # RDR
     # Translate outgoing ftp control connections to send them to localhost
     # for proxying with ftp-proxy(8) running on port 8021.
     rdr on $int_if proto tcp from any to any port 21 -> 127.0.0.1 port 8021
     In  this  example,  a NAT gateway is set up to translate internal addresses
     using a pool of public addresses (192.0.2.16/28) and to  redirect	incoming
     web server connections to a group of web servers on the internal network.
     # NAT LOAD BALANCE
     # Translate outgoing packets' source addresses using an address pool.
     # A given source address is always translated to the same pool address by
     # using the source-hash keyword.
     nat on $ext_if inet from any to any -> 192.0.2.16/28 source-hash
     # RDR ROUND ROBIN
     # Translate incoming web server connections to a group of web servers on
     # the internal network.
     rdr on $ext_if proto tcp from any to any port 80 \
	   -> { 10.1.2.155, 10.1.2.160, 10.1.2.161 } round-robin
**FILTER EXAMPLES**
     # The external interface is kue0
     # (157.161.48.183, the only routable address)
     # and the private network is 10.0.0.0/8, for which we are doing NAT.
     # use a macro for the interface name, so it can be changed easily
     ext_if = "kue0"
     # normalize all incoming traffic
     scrub in on $ext_if all fragment reassemble
     # block and log everything by default
     block return log on $ext_if all
     # block anything coming from source we have no back routes for
     block in from no-route to any
     # block packets whose ingress interface does not match the one in
     # the route back to their source address
     block in from urpf-failed to any
     # block and log outgoing packets that do not have our address as source,
     # they are either spoofed or something is misconfigured (NAT disabled,
     # for instance), we want to be nice and do not send out garbage.
     block out log quick on $ext_if from ! 157.161.48.183 to any
     # silently drop broadcasts (cable modem noise)
     block in quick on $ext_if from any to 255.255.255.255
     # block and log incoming packets from reserved address space and invalid
     # addresses, they are either spoofed or misconfigured, we cannot reply to
     # them anyway (hence, no return-rst).
     block in log quick on $ext_if from { 10.0.0.0/8, 172.16.0.0/12, \
	   192.168.0.0/16, 255.255.255.255/32 } to any
     # ICMP
     # pass out/in certain ICMP queries and keep state (ping)
     # state matching is done on host addresses and ICMP id (not type/code),
     # so replies (like 0/0 for 8/0) will match queries
     # ICMP error messages (which always refer to a TCP/UDP packet) are
     # handled by the TCP/UDP states
     pass on $ext_if inet proto icmp all icmp-type 8 code 0
     # UDP
     # pass out all UDP connections and keep state
     pass out on $ext_if proto udp all
     # pass in certain UDP connections and keep state (DNS)
     pass in on $ext_if proto udp from any to any port domain
     # TCP
     # pass out all TCP connections and modulate state
     pass out on $ext_if proto tcp all modulate state
     # pass in certain TCP connections and keep state (SSH, SMTP, DNS, IDENT)
     pass in on $ext_if proto tcp from any to any port { ssh, smtp, domain, \
	   auth }
     # Do not allow Windows 9x SMTP connections since they are typically
     # a viral worm. Alternately we could limit these OSes to 1 connection each.
     block in on $ext_if proto tcp from any os {"Windows 95", "Windows 98"} \
	   to any port smtp
     # IPv6
     # pass in/out all IPv6 traffic: note that we have to enable this in two
     # different ways, on both our physical interface and our tunnel
     pass quick on gif0 inet6
     pass quick on $ext_if proto ipv6
     # Packet Tagging
     # three interfaces: $int_if, $ext_if, and $wifi_if (wireless). NAT is
     # being done on $ext_if for all outgoing packets. tag packets in on
     # $int_if and pass those tagged packets out on $ext_if.  all other
     # outgoing packets (i.e., packets from the wireless network) are only
     # permitted to access port 80.
     pass in on $int_if from any to any tag INTNET
     pass in on $wifi_if from any to any
     block out on $ext_if from any to any
     pass out quick on $ext_if tagged INTNET
     pass out on $ext_if proto tcp from any to any port 80
     # tag incoming packets as they are redirected to spamd(8). use the tag
     # to pass those packets through the packet filter.
     rdr on $ext_if inet proto tcp from <spammers> to port smtp \
	     tag SPAMD -> 127.0.0.1 port spamd
     block in on $ext_if
     pass in on $ext_if inet proto tcp tagged SPAMD
**GRAMMAR**
     Syntax for **pf.conf** in BNF:
     line	    = ( option | pf-rule | nat-rule | binat-rule | rdr-rule |
		      antispoof-rule | altq-rule | queue-rule | trans-anchors |
		      anchor-rule | anchor-close | load-anchor | table-rule | )
     option	    = "set" ( [ "timeout" ( timeout | "{" timeout-list "}" ) ] |
		      [ "ruleset-optimization" [ "none" | "basic" | "profile" ]] |
		      [ "optimization" [ "default" | "normal" |
		      "high-latency" | "satellite" |
		      "aggressive" | "conservative" ] ]
		      [ "limit" ( limit-item | "{" limit-list "}" ) ] |
		      [ "loginterface" ( interface-name | "none" ) ] |
		      [ "block-policy" ( "drop" | "return" ) ] |
		      [ "state-policy" ( "if-bound" | "floating" ) ]
		      [ "require-order" ( "yes" | "no" ) ]
		      [ "fingerprints" filename ] |
		      [ "skip on" ( interface-name | "{" interface-list "}" ) ] |
		      [ "debug" ( "none" | "urgent" | "misc" | "loud" ) ] )
     pf-rule	    = action [ ( "in" | "out" ) ]
		      [ "log" [ "(" logopts ")"] ] [ "quick" ]
		      [ "on" ifspec ] [ "fastroute" | route ] [ af ] [ protospec ]
		      hosts [ filteropt-list ]
     logopts	    = logopt [ "," logopts ]
     logopt	    = "all" | "user" | "to" interface-name
     filteropt-list = filteropt-list filteropt | filteropt
     filteropt	    = user | group | flags | icmp-type | icmp6-type | tos |
		      ( "no" | "keep" | "modulate" | "synproxy" ) "state"
		      [ "(" state-opts ")" ] |
		      "fragment" | "no-df" | "min-ttl" number |
		      "max-mss" number | "random-id" | "reassemble tcp" |
		      fragmentation | "allow-opts" |
		      "label" string | "tag" string | [ ! ] "tagged" string |
		      "queue" ( string | "(" string [ [ "," ] string ] ")" ) |
		      "rtable" number | "probability" number"%"
     nat-rule	    = [ "no" ] "nat" [ "pass" [ "log" [ "(" logopts ")" ] ] ]
		      [ "on" ifspec ] [ af ]
		      [ protospec ] hosts [ "tag" string ] [ "tagged" string ]
		      [ "->" ( redirhost | "{" redirhost-list "}" )
		      [ portspec ] [ pooltype ] [ "static-port" ] ]
     binat-rule     = [ "no" ] "binat" [ "pass" [ "log" [ "(" logopts ")" ] ] ]
		      [ "on" interface-name ] [ af ]
		      [ "proto" ( proto-name | proto-number ) ]
		      "from" address [ "/" mask-bits ] "to" ipspec
		      [ "tag" string ] [ "tagged" string ]
		      [ "->" address [ "/" mask-bits ] ]
     rdr-rule	    = [ "no" ] "rdr" [ "pass" [ "log" [ "(" logopts ")" ] ] ]
		      [ "on" ifspec ] [ af ]
		      [ protospec ] hosts [ "tag" string ] [ "tagged" string ]
		      [ "->" ( redirhost | "{" redirhost-list "}" )
		      [ portspec ] [ pooltype ] ]
     antispoof-rule = "antispoof" [ "log" ] [ "quick" ]
		      "for" ( interface-name | "{" interface-list "}" )
		      [ af ] [ "label" string ]
     table-rule     = "table" "<" string ">" [ tableopts-list ]
     tableopts-list = tableopts-list tableopts | tableopts
