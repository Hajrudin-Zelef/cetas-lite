---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-ee37d35a-12
title: "cgi-man-cgi-ee37d35a"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2006-10-30"]
keywords: ["parameters", "throughput"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-ee37d35a.md
source_anchor: ""
source_lines: [1511, 1647]
sha256: 7b6ed01b87c57a7523fac02e523fdd17e7665927fff1b440056c0c61e826892b
---

# cgi-man-cgi-ee37d35a

     tableopts	    = "persist" | "const" | "file" string |
		      "{" [ tableaddr-list ] "}"
     tableaddr-list = tableaddr-list [ "," ] tableaddr-spec | tableaddr-spec
     tableaddr-spec = [ "!" ] tableaddr [ "/" mask-bits ]
     tableaddr	    = hostname | ipv4-dotted-quad | ipv6-coloned-hex |
		      interface-name | "self"
     altq-rule	    = "altq on" interface-name queueopts-list
		      "queue" subqueue
     queue-rule     = "queue" string [ "on" interface-name ] queueopts-list
		      subqueue
     anchor-rule    = "anchor" [ string ] [ ( "in" | "out" ) ] [ "on" ifspec ]
		      [ af ] [ protospec ] [ hosts ] [ "{" ]
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
     action	    = "pass" | "block" [ return ] | [ "no" ] "scrub"
     return	    = "drop" | "return" | "return-rst" [ "( ttl" number ")" ] |
		      "return-icmp" [ "(" icmpcode [ [ "," ] icmp6code ] ")" ] |
		      "return-icmp6" [ "(" icmp6code ")" ]
     icmpcode	    = ( icmp-code-name | icmp-code-number )
     icmp6code	    = ( icmp6-code-name | icmp6-code-number )
     ifspec	    = ( [ "!" ] interface-name ) | "{" interface-list "}"
     interface-list = [ "!" ] interface-name [ [ "," ] interface-list ]
     route	    = ( "route-to" | "reply-to" | "dup-to" )
		      ( routehost | "{" routehost-list "}" )
		      [ pooltype ]
     af 	    = "inet" | "inet6"
     protospec	    = "proto" ( proto-name | proto-number |
		      "{" proto-list "}" )
     proto-list     = ( proto-name | proto-number ) [ [ "," ] proto-list ]
     hosts	    = "all" |
		      "from" ( "any" | "no-route" | "urpf-failed" | "self" | host |
		      "{" host-list "}" | "route" string ) [ port ] [ os ]
		      "to"   ( "any" | "no-route" | "self" | host |
		      "{" host-list "}" | "route" string ) [ port ]
     ipspec	    = "any" | host | "{" host-list "}"
     host	    = [ "!" ] ( address [ "/" mask-bits ] | "<" string ">" )
     redirhost	    = address [ "/" mask-bits ]
     routehost	    = "(" interface-name [ address [ "/" mask-bits ] ] ")"
     address	    = ( interface-name | "(" interface-name ")" | hostname |
		      ipv4-dotted-quad | ipv6-coloned-hex )
     host-list	    = host [ [ "," ] host-list ]
     redirhost-list = redirhost [ [ "," ] redirhost-list ]
     routehost-list = routehost [ [ "," ] routehost-list ]
     port	    = "port" ( unary-op | binary-op | "{" op-list "}" )
     portspec	    = "port" ( number | name ) [ ":" ( "*" | number | name ) ]
     os 	    = "os"  ( os-name | "{" os-list "}" )
     user	    = "user" ( unary-op | binary-op | "{" op-list "}" )
     group	    = "group" ( unary-op | binary-op | "{" op-list "}" )
     unary-op	    = [ "=" | "!=" | "<" | "<=" | ">" | ">=" ]
		      ( name | number )
     binary-op	    = number ( "<>" | "><" | ":" ) number
     op-list	    = ( unary-op | binary-op ) [ [ "," ] op-list ]
     os-name	    = operating-system-name
     os-list	    = os-name [ [ "," ] os-list ]
     flags	    = "flags" ( [ flag-set ] "/"  flag-set | "any" )
     flag-set	    = [ "F" ] [ "S" ] [ "R" ] [ "P" ] [ "A" ] [ "U" ] [ "E" ]
		      [ "W" ]
     icmp-type	    = "icmp-type" ( icmp-type-code | "{" icmp-list "}" )
     icmp6-type     = "icmp6-type" ( icmp-type-code | "{" icmp-list "}" )
     icmp-type-code = ( icmp-type-name | icmp-type-number )
		      [ "code" ( icmp-code-name | icmp-code-number ) ]
     icmp-list	    = icmp-type-code [ [ "," ] icmp-list ]
     tos	    = "tos" ( "lowdelay" | "throughput" | "reliability" |
		      [ "0x" ] number )
     state-opts     = state-opt [ [ "," ] state-opts ]
     state-opt	    = ( "max" number | "no-sync" | timeout |
		      "source-track" [ ( "rule" | "global" ) ] |
		      "max-src-nodes" number | "max-src-states" number |
		      "max-src-conn" number |
		      "max-src-conn-rate" number "/" number |
		      "overload" "<" string ">" [ "flush" ] |
		      "if-bound" | "floating" )
     fragmentation  = [ "fragment reassemble" | "fragment crop" |
		      "fragment drop-ovl" ]
     timeout-list   = timeout [ [ "," ] timeout-list ]
     timeout	    = ( "tcp.first" | "tcp.opening" | "tcp.established" |
		      "tcp.closing" | "tcp.finwait" | "tcp.closed" |
		      "udp.first" | "udp.single" | "udp.multiple" |
		      "icmp.first" | "icmp.error" |
		      "other.first" | "other.single" | "other.multiple" |
		      "frag" | "interval" | "src.track" |
		      "adaptive.start" | "adaptive.end" ) number
     limit-list     = limit-item [ [ "," ] limit-list ]
     limit-item     = ( "states" | "frags" | "src-nodes" ) number
     pooltype	    = ( "bitmask" | "random" |
		      "source-hash" [ ( hex-key | string-key ) ] |
		      "round-robin" ) [ sticky-address ]
     subqueue	    = string | "{" queue-list "}"
     queue-list     = string [ [ "," ] string ]
     cbq-def	    = "cbq" [ "(" cbq-opt [ [ "," ] cbq-opt ] ")" ]
     priq-def	    = "priq" [ "(" priq-opt [ [ "," ] priq-opt ] ")" ]
     hfsc-def	    = "hfsc" [ "(" hfsc-opt [ [ "," ] hfsc-opt ] ")" ]
     cbq-opt	    = ( "default" | "borrow" | "red" | "ecn" | "rio" )
     priq-opt	    = ( "default" | "red" | "ecn" | "rio" )
     hfsc-opt	    = ( "default" | "red" | "ecn" | "rio" |
		      linkshare-sc | realtime-sc | upperlimit-sc )
     linkshare-sc   = "linkshare" sc-spec
     realtime-sc    = "realtime" sc-spec
     upperlimit-sc  = "upperlimit" sc-spec
     sc-spec	    = ( bandwidth-spec |
		      "(" bandwidth-spec number bandwidth-spec ")" )
**FILES**
     */etc/hosts* 	     Host name database.
     */etc/pf.conf*	     Default location of the ruleset file.
     */etc/pf.os* 	     Default location of OS fingerprints.
     */etc/protocols*	     Protocol name database.
     */etc/services*	     Service name database.
     */usr/share/examples/pf*  Example rulesets.
**BUGS**
     Due  to  a  lock order reversal (LOR) with the socket layer, the use of the
     *group* and *user* filter parameter in conjuction with  a  Giant-free	netstack
     can   result   in	 a  deadlock.	A  workaround  is  available  under  the
     *debug.pfugidhack* sysctl which is automatically enabled when a *user* /  *group*
     rule is added or *log* *(user)* is specified.
     Route  labels are not supported by the FreeBSD *route*(4) system.  Rules with
     a route label do not match any traffic.
**SEE ALSO**
     *altq*(4), *carp*(4),	*icmp*(4),  *icmp6*(4),  *ip*(4),  *ip6*(4),  *pf*(4),  *pfsync*(4),
     *route*(4),	*tcp*(4),  *udp*(4),  *hosts*(5), *pf.os*(5), *protocols*(5), *services*(5),
     *ftp-proxy*(8), *pfctl*(8), *pflogd*(8), *route*(8)
**HISTORY**
     The **pf.conf** file format first appeared in OpenBSD 3.0.
FreeBSD 8.0			October 30, 2006		      *PF.CONF*(5)

NAME | DESCRIPTION | STATEMENT ORDER | MACROS | TABLES | OPTIONS | TRAFFIC NORMALIZATION | QUEUEING/ALTQ | TRANSLATION | PACKET FILTERING | PARAMETERS | ROUTING | POOL OPTIONS | STATE MODULATION | SYN PROXY | STATEFUL TRACKING OPTIONS | OPERATING SYSTEM FINGERPRINTING | BLOCKING SPOOFED TRAFFIC | FRAGMENT HANDLING | ANCHORS | TRANSLATION EXAMPLES | FILTER EXAMPLES | GRAMMAR | FILES | BUGS | SEE ALSO | HISTORY

Want to link to this manual page? Use this URL:

<https://man.freebsd.org/cgi/man.cgi?query=pf.conf&manpath=FreeBSD+8.0-RELEASE>
