---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-1fcc53d9-3
title: "cgi-man-cgi-1fcc53d9"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-1fcc53d9.md
source_anchor: ""
source_lines: [307, 444]
sha256: aa410eb54d7afd4c8a38174b9abd7cb5283d422c39f308a237fe3c274fa2e418
---

# cgi-man-cgi-1fcc53d9

	   *if-bound*	States are bound to interface.
	   *floating*	States	can  match  packets  on  any interfaces (the de-
			fault).
	   For example:
		 set state-policy if-bound
     *set* *syncookies* *never* | *always* | *adaptive*
	   When **syncookies** are active, pf will answer each incoming TCP SYN with
	   a syncookie SYNACK, without allocating any resources.  Upon reception
	   of the client's ACK in response to  the  syncookie  SYNACK,	pf  will
	   evaluate the ruleset and create state if the ruleset permits it, com-
	   plete  the  three way handshake with the target host and continue the
	   connection with synproxy in place.  This allows pf  to  be  resilient
	   against  large  synflood  attacks  which  would  run  the state table
	   against its limits otherwise.  Due to the blind answers to every  in-
	   coming SYN syncookies share the caveats of synproxy, namely seemingly
	   accepting connections that will be dropped later on.
	   **never**     pf will never send syncookie SYNACKs (the default).
	   **always**    pf will always send syncookie SYNACKs.
	   **adaptive**  pf  will  enable  syncookie mode when a given percentage of
		     the state table is used up by half-open TCP connections, as
		     in, those that saw the initial SYN but  didn't  finish  the
		     three way handshake.  The thresholds for entering and leav-
		     ing syncookie mode can be specified using
			   set syncookies adaptive (start 25%, end 12%)
     *set* *state-defaults*
	   The	*state-defaults*	option sets the state options for states created
	   from rules without an explicit *keep* *state*.  For example:
		 set state-defaults no-sync
     *set* *hostid*
	   The 32-bit *hostid* identifies this firewall's state table  entries  to
	   other  firewalls  in  a  *pfsync*(4)  failover cluster.  By default the
	   hostid is set to a pseudo-random value, however it may  be  desirable
	   to  manually  configure  it,  for example to more easily identify the
	   source of state table entries.
		 set hostid 1
	   The hostid may be specified in either decimal or hexadecimal.
     *set* *require-order*
	   By default *pfctl*(8) enforces an ordering of the  statement  types  in
	   the	 ruleset  to:  *options*,  *normalization*,  *queueing*,  *translation*,
	   *filtering*.  Setting this option  to	*no*  disables  this  enforcement.
	   There  may  be  non-trivial and non-obvious implications to an out of
	   order ruleset.  Consider carefully before  disabling  the  order  en-
	   forcement.
     *set* *fingerprints*
	   Load fingerprints of known operating systems from the given filename.
	   By  default fingerprints of known operating systems are automatically
	   loaded from *pf.os*(5) in */etc* but can be overridden via  this  option.
	   Setting  this  option may leave a small period of time where the fin-
	   gerprints referenced by the currently active ruleset are inconsistent
	   until the new ruleset finishes loading.   The  default  location  for
	   fingerprints is */etc/pf.os*.
	   For example:
		 **set** **fingerprints** **"/etc/pf.os.devel"**
     *set* *skip* *on* <*ifspec*>
	   List  interfaces  for  which packets should not be filtered.  Packets
	   passing in or out on such interfaces are passed as  if  pf  was  dis-
	   abled,  i.e. pf does not process them in any way.  This can be useful
	   on loopback and other virtual interfaces, when  packet  filtering  is
	   not desired and can have unexpected effects.  For example:
		 **set** **skip** **on** **lo0**
     *set* *debug*
	   Set the debug *level* to one of the following:
	   *none* 	 Don't generate debug messages.
	   *urgent*	 Generate debug messages only for serious errors.
	   *misc* 	 Generate debug messages for various errors.
	   *loud* 	 Generate debug messages for common conditions.
     *set* *keepcounters*
	   Preserve  rule  counters  across rule updates.  Usually rule counters
	   are reset to zero on every update of the ruleset.  With  *keepcounters*
	   set	pf will attempt to find matching rules between old and new rule-
	   sets and preserve the rule counters.
**ETHERNET FILTERING**
     *pf*(4) has the ability to *block* and *pass*  packets  based  on  attributes  of
     their Ethernet (layer 2) header.
     Each  time  a packet processed by the packet filter comes in on or goes out
     through an interface, the filter rules are evaluated in  sequential  order,
     from  first  to last.  The last matching rule decides what action is taken.
     If no rule matches the packet, the default action is  to  pass  the  packet
     without creating a state.
     The following actions can be used in the filter:
     *block*
	   The	packet is blocked.  Unlike for layer 3 traffic the packet is al-
	   ways silently dropped.
     *pass*  The packet is passed; no state is created for layer 2 traffic.
   **Parameters** **applicable** **to** **layer** **2** **rules**
     The rule parameters specify the packets to which a rule applies.  A  packet
     always  comes  in	on, or goes out through, one interface.  Most parameters
     are optional.  If a parameter is specified, the rule only applies to  pack-
     ets  with matching attributes.  The matching for some parameters can be in-
     verted with the **!** operator.  Certain parameters can be expressed as  lists,
     in which case *pfctl*(8) generates all needed rule combinations.
     *in* or *out*
	   This rule applies to incoming or outgoing packets.  If neither *in* nor
	   *out* are specified, the rule will match packets in both directions.
     *quick*
	   If  a packet matches a rule which has the *quick* option set, this rule
	   is considered the last matching rule, and  evaluation  of  subsequent
	   rules is skipped.
     *on* <*ifspec*>
	   This rule applies only to packets coming in on, or going out through,
	   this  particular  interface or interface group.  For more information
	   on interface groups, see the **group** keyword in *ifconfig*(8).  *any*  will
	   match any existing interface except loopback ones.
     *bridge-to* <interface>
	   Packets  matching  this rule will be sent out of the specified inter-
	   face without further processing.
     *proto* <*protocol*>
	   This rule applies only to packets of this protocol.	Note that Ether-
	   net protocol numbers are different  from  those  used  in  *ip*(4)  and
	   *ip6*(4).
     *from* <*source*> *to* <*dest*>
	   This  rule applies only to packets with the specified source and des-
	   tination MAC addresses.
     *queue* <*queue*>
	   Packets matching this rule will be assigned to the  specified  queue.
	   See "QUEUEING" for setup details.
     *tag* <*string*>
	   Packets  matching this rule will be tagged with the specified string.
	   The tag acts as an internal marker that can be used to identify these
	   packets later on.  This can be used, for example,  to  provide  trust
	   between interfaces and to determine if packets have been processed by
	   translation	rules.	 Tags are "sticky", meaning that the packet will
	   be tagged even if the rule is not the last  matching  rule.	 Further
	   matching rules can replace the tag with a new one but will not remove
	   a  previously applied tag.  A packet is only ever assigned one tag at
	   a time.
     *tagged* <*string*>
	   Used to specify that packets must already be tagged	with  the  given
	   tag	in  order  to  match the rule.	Inverse tag matching can also be
	   done by specifying the !  operator before the tagged keyword.
**TRAFFIC NORMALIZATION**
     Traffic normalization is a broad umbrella term for aspects  of  the  packet
     filter  which  deal with verifying packets, packet fragments, spoofed traf-
     fic, and other irregularities.
   **Scrub**
     Scrub involves sanitising packet content in such a way that  there  are  no
     ambiguities  in packet interpretation on the receiving side.  It is invoked
     with the **scrub** option, added to filter rules.
