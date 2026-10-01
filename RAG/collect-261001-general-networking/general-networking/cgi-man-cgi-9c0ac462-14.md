---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-9c0ac462-14
title: "cgi-man-cgi-9c0ac462"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-9c0ac462.md
source_anchor: ""
source_lines: [1773, 1919]
sha256: eb3bfa5b6b22d9e88cab7507c5d0eb0829848e117af54308c059b4bd86d2b9ae
---

# cgi-man-cgi-9c0ac462

	   pass out on $ext_if proto tcp from any to any port 80
	   # tag incoming packets as they are redirected to spamd(8). use the tag
	   # to pass those packets through the packet filter.
	   rdr on $ext_if inet proto tcp from <spammers> to port smtp \
		   tag SPAMD -> 127.0.0.1 port spamd
	   block in on $ext_if
	   pass in on $ext_if inet proto tcp tagged SPAMD
     In the example below, a router handling both address families translates an
     internal IPv4 subnet to IPv6 using the well-known 64:ff9b::/96 prefix:
	 pass in on $v4_if inet af-to inet6 from ($v6_if) to 64:ff9b::/96
     Paired  with  the	example  above, the example below can be used on another
     router handling both address families to translate back to IPv4:
	 pass in on $v6_if inet6 to 64:ff9b::/96 af-to inet from ($v4_if)
**GRAMMAR**
     Syntax for **pf.conf** in BNF:
     line	    = ( option | ether-rule | pf-rule | nat-rule | binat-rule |
		      rdr-rule | antispoof-rule | altq-rule | queue-rule |
		      trans-anchors | anchor-rule | anchor-close | load-anchor |
		      table-rule | include )
     option	    = "set" ( [ "timeout" ( timeout | "{" timeout-list "}" ) ] |
		      [ "ruleset-optimization" [ "none" | "basic" | "profile" ]] |
		      [ "optimization" [ "default" | "normal" |
		      "high-latency" | "satellite" |
		      "aggressive" | "conservative" ] ]
		      [ "limit" ( limit-item | "{" limit-list "}" ) ] |
		      [ "loginterface" ( interface-name | "none" ) ] |
		      [ "block-policy" ( "drop" | "return" ) ] |
		      [ "state-policy" ( "if-bound" | "floating" ) ]
		      [ "state-defaults" state-opts ]
		      [ "require-order" ( "yes" | "no" ) ]
		      [ "fingerprints" filename ] |
		      [ "skip on" ifspec ] |
		      [ "debug" ( "none" | "urgent" | "misc" | "loud" ) ]
		      [ "keepcounters" ] )
     ether-rule     = "ether" etheraction [ ( "in" | "out" ) ]
		      [ "quick" ] [ "on" ifspec ] [ "bridge-to" interface-name ]
		      [ etherprotospec ] [ etherhosts ] [ "l3" hosts ]
		      [ etherfilteropt-list ]
     pf-rule	    = action [ ( "in" | "out" ) ]
		      [ "log" [ "(" logopts ")"] ] [ "quick" ]
		      [ "on" ifspec ] [ route ] [ af ] [ protospec ]
		      [ hosts ] [ filteropt-list ]
     logopts	    = logopt [ "," logopts ]
     logopt	    = "all" | "matches" | "user" | "to" interface-name
     etherfilteropt-list = etherfilteropt-list etherfilteropt | etherfilteropt
     etherfilteropt = "tag" string | "tagged" string | "queue" ( string ) |
		      "ridentifier" number | "label" string
     filteropt-list = filteropt-list filteropt | filteropt
     filteropt	    = user | group | flags | icmp-type | icmp6-type | "tos" tos |
		      "af-to" af "from" ( redirhost | "{" redirhost-list "}" )
		      [ "to" ( redirhost | "{" redirhost-list "}" ) ] |
		      ( "no" | "keep" | "modulate" | "synproxy" ) "state"
		      [ "(" state-opts ")" ] |
		      "fragment" | "no-df" | "min-ttl" number | "set-tos" tos |
		      "max-mss" number | "random-id" | "reassemble tcp" |
		      fragmentation | "allow-opts" |
		      "label" string | "tag" string | [ "!" ] "tagged" string |
		      "max-pkt-rate" number "/" seconds |
		      "set prio" ( number | "(" number [ [ "," ] number ] ")" ) |
		      "max-pkt-size" number |
		      "queue" ( string | "(" string [ [ "," ] string ] ")" ) |
		      "rtable" number | "probability" number"%" | "prio" number |
		      "dnpipe" ( number | "(" number "," number ")" ) |
		      "dnqueue" ( number | "(" number "," number ")" ) |
		      "ridentifier" number |
		      "binat-to" ( redirhost | "{" redirhost-list "}" )
		      [ portspec ] [ pooltype ] |
		      "rdr-to" ( redirhost | "{" redirhost-list "}" )
		      [ portspec ] [ pooltype ] |
		      "nat-to" ( redirhost | "{" redirhost-list "}" )
		      [ portspec ] [ pooltype ] [ "static-port" ] |
		      [ ! ] "received-on" ( interface-name | interface-group )
     nat-rule	    = [ "no" ] "nat" [ "pass" [ "log" [ "(" logopts ")" ] ] ]
		      [ "on" ifspec ] [ af ]
		      [ protospec ] hosts [ "tag" string ] [ "tagged" string ]
		      [ "->" ( redirhost | "{" redirhost-list "}" )
		      [ portspec ] [ pooltype ] [ "static-port" ]
		      [ "map-e-portset" number "/" number "/" number ] ]
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
		      "for" ifspec [ af ] [ "label" string ]
		      [ "ridentifier" number ]
     table-rule     = "table" "<" string ">" [ tableopts-list ]
     tableopts-list = tableopts-list tableopts | tableopts
     tableopts	    = "persist" | "const" | "counters" | "file" string |
		      "{" [ tableaddr-list ] "}"
     tableaddr-list = tableaddr-list [ "," ] tableaddr-spec | tableaddr-spec
     tableaddr-spec = [ "!" ] tableaddr [ "/" mask-bits ]
     tableaddr	    = hostname | ifspec | "self" |
		      ipv4-dotted-quad | ipv6-coloned-hex
     altq-rule	    = "altq on" interface-name queueopts-list
		      "queue" subqueue
     queue-rule     = "queue" string [ "on" interface-name ] queueopts-list
		      subqueue
     anchor-rule    = "anchor" [ string ] [ ( "in" | "out" ) ] [ "on" ifspec ]
		      [ af ] [ protospec ] [ hosts ] [ filteropt-list ] [ "{" ]
     anchor-close   = "}"
     trans-anchors  = ( "nat-anchor" | "rdr-anchor" | "binat-anchor" ) string
		      [ "on" ifspec ] [ af ] [ "proto" ] [ protospec ] [ hosts ]
     load-anchor    = "load anchor" string "from" filename
     queueopts-list = queueopts-list queueopts | queueopts
     queueopts	    = [ "bandwidth" bandwidth-spec ] |
		      [ "qlimit" number ] | [ "tbrsize" number ] |
		      [ "priority" number ] | [ schedulers ]
     schedulers     = ( cbq-def | priq-def | hfsc-def )
     bandwidth-spec = "number" ( "b" | "Kb" | "Mb" | "Gb" | "%" )
     etheraction    = "pass" | "block"
     action	    = "pass" | "match" | "block" [ return ] | [ "no" ] "scrub"
     return	    = "drop" | "return" | "return-rst" [ "( ttl" number ")" ] |
		      "return-icmp" [ "(" icmpcode [ [ "," ] icmp6code ] ")" ] |
		      "return-icmp6" [ "(" icmp6code ")" ]
     icmpcode	    = ( icmp-code-name | icmp-code-number )
     icmp6code	    = ( icmp6-code-name | icmp6-code-number )
     ifspec	    = ( [ "!" ] ( interface-name | interface-group ) ) |
		      "{" interface-list "}"
     interface-list = [ "!" ] ( interface-name | interface-group )
		      [ [ "," ] interface-list ]
     route	    = ( "route-to" | "reply-to" | "dup-to" )
		      ( routehost | "{" routehost-list "}" )
		      [ pooltype ]
     af 	    = "inet" | "inet6"
     etherprotospec = "proto" ( proto-number | "{" etherproto-list "}" )
     etherproto-list = proto-number [ [ "," ] etherproto-list ]
     protospec	    = "proto" ( proto-name | proto-number |
		      "{" proto-list "}" )
     proto-list     = ( proto-name | proto-number ) [ [ "," ] proto-list ]
     etherhosts     = "from" macaddress "to" macaddress
     macaddress     = mac | mac "/" masklen | mac "&" mask
     hosts	    = "all" |
		      "from" ( "any" | "no-route" | "urpf-failed" | "self" | host |
		      "{" host-list "}" ) [ port ] [ os ]
		      "to"   ( "any" | "no-route" | "self" | host |
		      "{" host-list "}" ) [ port ]
     ipspec	    = "any" | host | "{" host-list "}"
     host	    = [ "!" ] ( address [ "/" mask-bits ] | "<" string ">" )
     redirhost	    = address [ "/" mask-bits ]
     routehost	    = "(" interface-name address [ "/" mask-bits ] ")"
