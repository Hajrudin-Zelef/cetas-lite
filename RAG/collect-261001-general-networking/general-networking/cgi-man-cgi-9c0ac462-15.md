---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-9c0ac462-15
title: "cgi-man-cgi-9c0ac462"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2025-11-03"]
keywords: ["ethernet", "throughput"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-9c0ac462.md
source_anchor: ""
source_lines: [1920, 2006]
sha256: adbaf4f2b043ac158e99bbc9251207e206f3df2336069315f8b9071e39209fe0
---

# cgi-man-cgi-9c0ac462

     address	    = ( interface-name | interface-group |
		      "(" ( interface-name | interface-group ) ")" |
		      hostname | ipv4-dotted-quad | ipv6-coloned-hex )
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
     tos	    = ( "lowdelay" | "throughput" | "reliability" |
		      [ "0x" ] number )
     state-opts     = state-opt [ [ "," ] state-opts ]
     state-opt	    = ( "max" number | "no-sync" | timeout | "sloppy" |
		      "source-track" [ ( "rule" | "global" ) ] |
		      "max-src-nodes" number | "max-src-states" number |
		      "max-src-conn" number |
		      "max-src-conn-rate" number "/" number |
		      "overload" "<" string ">" [ "flush" ] |
		      "if-bound" | "floating" | "pflow" )
     fragmentation  = [ "fragment reassemble" ]
     timeout-list   = timeout [ [ "," ] timeout-list ]
     timeout	    = ( "tcp.first" | "tcp.opening" | "tcp.established" |
		      "tcp.closing" | "tcp.finwait" | "tcp.closed" |
		      "sctp.first" | "sctp.opening" | "sctp.established" |
		      "sctp.closing" | "sctp.closed" |
		      "udp.first" | "udp.single" | "udp.multiple" |
		      "icmp.first" | "icmp.error" |
		      "other.first" | "other.single" | "other.multiple" |
		      "frag" | "interval" | "src.track" |
		      "adaptive.start" | "adaptive.end" ) number
     limit-list     = limit-item [ [ "," ] limit-list ]
     limit-item     = ( "states" | "frags" | "src-nodes" ) number
     pooltype	    = ( "bitmask" | "random" |
		      "source-hash" [ ( hex-key | string-key ) ] |
		      "round-robin" ) [ sticky-address | prefer-ipv6-nexthop ]
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
     include	    = "include" filename
**FILES**
     */etc/hosts*      Host name database.
     */etc/pf.conf*    Default location of the ruleset file.  The file has  to  be
		     created manually as it is not installed with a standard in-
		     stallation.
     */etc/pf.os*      Default location of OS fingerprints.
     */etc/protocols*  Protocol name database.
     */etc/services*   Service name database.
**SEE ALSO**
     *altq*(4),  *carp*(4),  *icmp*(4),  *icmp6*(4), *ip*(4), *ip6*(4), *pf*(4), *pflow*(4), *pf-*
     *sync*(4), *sctp*(4), *tcp*(4), *udp*(4), *hosts*(5),  *pf.os*(5),  *protocols*(5),  *ser-*
     *vices*(5), *ftp-proxy*(8), *pfctl*(8), *pflogd*(8)
**HISTORY**
     The **pf.conf** file format first appeared in OpenBSD 3.0.
FreeBSD 15.1			November 3, 2025		      *PF.CONF*(5)

NAME | DESCRIPTION | STATEMENT ORDER | MACROS | TABLES | OPTIONS | ETHERNET FILTERING | TRAFFIC NORMALIZATION | QUEUEING with ALTQ | QUEUEING with dummynet | TRANSLATION | PACKET FILTERING | ROUTING | POOL OPTIONS | STATE MODULATION | SYN PROXY | STATEFUL TRACKING OPTIONS | OPERATING SYSTEM FINGERPRINTING | BLOCKING SPOOFED TRAFFIC | FRAGMENT HANDLING | ANCHORS | SCTP CONSIDERATIONS | TRANSLATION EXAMPLES | COMPATIBILITY TRANSLATION EXAMPLES | FILTER EXAMPLES | GRAMMAR | FILES | SEE ALSO | HISTORY

Want to link to this manual page? Use this URL:

<https://man.freebsd.org/cgi/man.cgi?query=pf.conf&sektion=5&manpath=FreeBSD+15.1-RELEASE+and+Ports>
