---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-ee37d35a-10
title: "cgi-man-cgi-ee37d35a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-ee37d35a.md
source_anchor: ""
source_lines: [1223, 1362]
sha256: 54dc937cada48cb4dadc466d1bbf74ce745e21b12e0fa4ecff687d6fd5c2faf7
---

# cgi-man-cgi-ee37d35a

	   Evaluates the filter rules in the specified *anchor*.
     *load* *anchor* <*name*> *from* <*file*>
	   Loads the rules from the specified file into the anchor *name*.
     When evaluation of the main ruleset reaches an *anchor* rule, *pf*(4) will pro-
     ceed to evaluate all rules specified in that anchor.
     Matching filter and translation rules marked with the *quick* option are  fi-
     nal  and  abort  the  evaluation of the rules in other anchors and the main
     ruleset.  If the *anchor* itself is marked with  the  *quick*	option,  ruleset
     evaluation  will  terminate  when	the  anchor  is  exited if the packet is
     matched by any rule within the anchor.
     *anchor* rules are evaluated relative to the anchor in which  they  are  con-
     tained.   For  example, all *anchor* rules specified in the main ruleset will
     reference anchor attachment points underneath the main ruleset, and  *anchor*
     rules  specified  in a file loaded from a *load* *anchor* rule will be attached
     under that anchor point.
     Rules may be contained in *anchor* attachment points which do not contain any
     rules when the main ruleset is loaded, and later such anchors can be manip-
     ulated through *pfctl*(8) without reloading the main  ruleset  or  other  an-
     chors.  For example,
	   ext_if = "kue0"
	   block on $ext_if all
	   anchor spam
	   pass out on $ext_if all
	   pass in on $ext_if proto tcp from any \
		 to $ext_if port smtp
     blocks all packets on the external interface by default, then evaluates all
     rules  in	the *anchor* named "spam", and finally passes all outgoing connec-
     tions and incoming connections to port 25.
	   # echo "block in quick from 1.2.3.4 to any" | \
		 pfctl -a spam -f -
     This loads a single rule into the *anchor*, which blocks all packets  from  a
     specific address.
     The  anchor  can  also  be populated by adding a *load* *anchor* rule after the
     *anchor* rule:
	   anchor spam
	   load anchor spam from "/etc/pf-spam.conf"
     When *pfctl*(8) loads **pf.conf**, it will also load all the rules from the  file
     */etc/pf-spam.conf* into the anchor.
     Optionally,  *anchor* rules can specify the parameter's direction, interface,
     address family, protocol and source/destination address/port using the same
     syntax as filter rules.  When parameters are used, the *anchor* rule is  only
     evaluated	for matching packets.  This allows conditional evaluation of an-
     chors, like:
	   block on $ext_if all
	   anchor spam proto tcp from any to any port smtp
	   pass out on $ext_if all
	   pass in on $ext_if proto tcp from any to $ext_if port smtp
     The rules inside *anchor* spam are only evaluated for *tcp* packets with desti-
     nation port 25.  Hence,
	   # echo "block in quick from 1.2.3.4 to any" | \
		 pfctl -a spam -f -
     will only block connections from 1.2.3.4 to port 25.
     Anchors may end with the asterisk (`*') character, which signifies that all
     anchors attached at that point should be evaluated in the alphabetical  or-
     dering of their anchor name.  For example,
	   anchor "spam/*"
     will  evaluate  each rule in each anchor attached to the **spam** anchor.  Note
     that it will only evaluate anchors that are directly attached to  the  **spam**
     anchor, and will not descend to evaluate anchors recursively.
     Since  anchors  are evaluated relative to the anchor in which they are con-
     tained, there is a mechanism for accessing the parent and ancestor  anchors
     of a given anchor.  Similar to file system path name resolution, if the se-
     quence  ".."  appears as an anchor path component, the parent anchor of the
     current anchor in the path evaluation at that point  will	become	the  new
     current anchor.  As an example, consider the following:
	   # echo ' anchor "spam/allowed" ' | pfctl -f -
	   # echo -e ' anchor "../banned" \n pass' | \
		 pfctl -a spam/allowed -f -
     Evaluation  of  the  main	ruleset  will lead into the **spam/allowed** anchor,
     which will evaluate the rules in the **spam/banned** anchor, if any, before fi-
     nally evaluating the *pass* rule.
     Filter rule *anchors* can also be loaded inline in the ruleset within a brace
     ('{' '}') delimited block.  Brace delimited blocks  may  contain  rules  or
     other  brace-delimited blocks.  When anchors are loaded this way the anchor
     name becomes optional.
	   anchor "external" on egress {
		   block
		   anchor out {
			   pass proto tcp from any to port { 25, 80, 443 }
		   }
		   pass in proto tcp to any port 22
	   }
     Since the parser specification for anchor names is a string, any  reference
     to an anchor name containing `/' characters will require double quote (`"')
     characters around the anchor name.
**TRANSLATION EXAMPLES**
     This  example  maps  incoming  requests on port 80 to port 8080, on which a
     daemon is running (because, for example, it is not run as root, and  there-
     fore lacks permission to bind to port 80).
     # use a macro for the interface name, so it can be changed easily
     ext_if = "ne3"
     # map daemon on 8080 to appear to be on 80
     rdr on $ext_if proto tcp from any to any port 80 -> 127.0.0.1 port 8080
     If  the  *pass*  modifier is given, packets matching the translation rule are
     passed without inspecting the filter rules:
     rdr pass on $ext_if proto tcp from any to any port 80 -> 127.0.0.1 \
	   port 8080
     In the example below, vlan12 is configured as  192.168.168.1;  the  machine
     translates  all  packets coming from 192.168.168.0/24 to 204.92.77.111 when
     they are going out any interface except vlan12.  This has the net effect of
     making traffic from the 192.168.168.0/24 network appear as though it is the
     Internet routable address 204.92.77.111 to nodes behind  any  interface  on
     the  router  except for the nodes on vlan12.  (Thus, 192.168.168.1 can talk
     to the 192.168.168.0/24 nodes.)
     nat on ! vlan12 from 192.168.168.0/24 to any -> 204.92.77.111
     In the example below, the machine sits between a fake internal  144.19.74.*
     network,  and a routable external IP of 204.92.77.100.  The *no* *nat* rule ex-
     cludes protocol AH from being translated.
     # NO NAT
     no nat on $ext_if proto ah from 144.19.74.0/24 to any
     nat on $ext_if from 144.19.74.0/24 to any -> 204.92.77.100
     In the example below, packets bound for one specific  server,  as	well  as
     those  generated  by  the	sysadmins are not proxied; all other connections
     are.
     # NO RDR
     no rdr on $int_if proto { tcp, udp } from any to $server port 80
     no rdr on $int_if proto { tcp, udp } from $sysadmins to any port 80
     rdr on $int_if proto { tcp, udp } from any to any port 80 -> 127.0.0.1 \
	   port 80
     This longer example uses both a NAT and a redirection.  The external inter-
     face has the address 157.161.48.183.  On localhost,  we  are  running  *ftp-*
     *proxy*(8),	waiting  for  FTP  sessions  to  be redirected to it.  The three
     mandatory anchors for *ftp-proxy*(8) are omitted from this example;	see  the
     *ftp-proxy*(8) manpage.
     # NAT
     # Translate outgoing packets' source addresses (any protocol).
     # In this case, any address but the gateway's external address is mapped.
     nat on $ext_if inet from ! ($ext_if) to any -> ($ext_if)
     # NAT PROXYING
     # Map outgoing packets' source port to an assigned proxy port instead of
     # an arbitrary port.
     # In this case, proxy outgoing isakmp with port 500 on the gateway.
     nat on $ext_if inet proto udp from any port = isakmp to any -> ($ext_if) \
	   port 500
     # BINAT
     # Translate outgoing packets' source address (any protocol).
     # Translate incoming packets' destination address to an internal machine
     # (bidirectional).
     binat on $ext_if from 10.1.2.150 to any -> $ext_if
     # RDR
