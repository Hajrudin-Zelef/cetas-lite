---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-ee37d35a-4
title: "cgi-man-cgi-ee37d35a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-ee37d35a.md
source_anchor: ""
source_lines: [429, 562]
sha256: 31a0d44b1f298e8490a197ac7f106de34da9fdf2cbe9991c36bdd5aab1182c77
---

# cgi-man-cgi-ee37d35a

	   Queueing is enabled on the named interface.
     <*scheduler*>
	   Specifies which queueing scheduler to use.  Currently supported  val-
	   ues	are *cbq* for Class Based Queueing, *priq* for Priority Queueing and
	   *hfsc* for the Hierarchical Fair Service Curve scheduler.
     *bandwidth* <*bw*>
	   The maximum bitrate for all queues on an interface may  be  specified
	   using  the  *bandwidth*  keyword.  The value can be specified as an ab-
	   solute value or as a percentage of the interface bandwidth.	When us-
	   ing an absolute value, the suffixes *b*, *Kb*, *Mb*, and  *Gb*  are	used  to
	   represent  bits, kilobits, megabits, and gigabits per second, respec-
	   tively.  The value must  not  exceed  the  interface  bandwidth.   If
	   *bandwidth* is not specified, the interface bandwidth is used (but take
	   note  that  some interfaces do not know their bandwidth, or can adapt
	   their bandwidth rates).
     *qlimit* <*limit*>
	   The maximum number of packets held in the queue.  The default is 50.
     *tbrsize* <*size*>
	   Adjusts the size, in bytes, of the token bucket  regulator.	 If  not
	   specified,  heuristics  based  on the interface bandwidth are used to
	   determine the size.
     *queue* <*list*>
	   Defines a list of subqueues to create on an interface.
     In the following example, the interface dc0 should queue up to 5 Mbit/s  in
     four  second-level  queues  using	Class Based Queueing.  Those four queues
     will be shown in a later example.
	   altq on dc0 cbq bandwidth 5Mb queue { std, http, mail, ssh }
     Once interfaces are activated for queueing using the *altq* directive, a  se-
     quence  of  *queue*	directives  may  be defined.  The name associated with a
     *queue* must match a queue defined in the *altq* directive (e.g. mail), or, ex-
     cept for the *priq* *scheduler*, in a parent *queue* declaration.  The  following
     keywords can be used:
     *on* <*interface*>
	   Specifies  the interface the queue operates on.  If not given, it op-
	   erates on all matching interfaces.
     *bandwidth* <*bw*>
	   Specifies the maximum bitrate to be processed  by  the  queue.   This
	   value must not exceed the value of the parent *queue* and can be speci-
	   fied as an absolute value or a percentage of the parent queue's band-
	   width.   If	not  specified,  defaults  to 100% of the parent queue's
	   bandwidth.  The *priq* scheduler does not support bandwidth  specifica-
	   tion.
     *priority* <*level*>
	   Between  queues  a  priority level can be set.  For *cbq* and *hfsc*, the
	   range is 0 to 7 and for *priq*, the range is 0 to 15.	The default  for
	   all	is  1.	 *Priq*  queues  with  a higher priority are always served
	   first.  *Cbq* and *Hfsc* queues with a higher priority are  preferred  in
	   the case of overload.
     *qlimit* <*limit*>
	   The maximum number of packets held in the queue.  The default is 50.
     The    *scheduler*	can   get   additional	 parameters   with   <*scheduler*>
     (<*parameters*>).  Parameters are as follows:
     *default*	 Packets not matched by another queue are assigned to this  one.
		 Exactly one default queue is required.
     *red*	 Enable  RED  (Random Early Detection) on this queue.  RED drops
		 packets with a probability proportional to  the  average  queue
		 length.
     *rio*	 Enables  RIO  on this queue.  RIO is RED with IN/OUT, thus run-
		 ning RED two times more than RIO would achieve the same effect.
		 RIO is currently not supported in the GENERIC kernel.
     *ecn*	 Enables ECN (Explicit Congestion Notification) on  this  queue.
		 ECN implies RED.
     The *cbq* *scheduler* supports an additional option:
     *borrow*	 The queue can borrow bandwidth from the parent.
     The *hfsc* *scheduler* supports some additional options:
     *realtime* <*sc*>
		 The minimum required bandwidth for the queue.
     *upperlimit* <*sc*>
		 The maximum allowed bandwidth for the queue.
     *linkshare* <*sc*>
		 The bandwidth share of a backlogged queue.
     <*sc*> is an acronym for *service* *curve*.
     The  format  for  service curve specifications is (*m1*, *d*, *m2*).  *m2* controls
     the bandwidth assigned to the queue.  *m1* and *d* are optional and can be used
     to control the initial bandwidth assignment.  For the first *d*  milliseconds
     the queue gets the bandwidth given as *m1*, afterwards the value given in *m2*.
     Furthermore, with *cbq* and *hfsc*, child queues can be specified as in an *altq*
     declaration,  thus building a tree of queues using a part of their parent's
     bandwidth.
     Packets can be assigned to queues based on filter rules by using the  *queue*
     keyword.  Normally only one *queue* is specified; when a second one is speci-
     fied  it  will instead be used for packets which have a *TOS* of *lowdelay* and
     for TCP ACKs with no data payload.
     To continue the previous example, the examples below would specify the four
     referenced queues, plus a few child queues.   Interactive	*ssh*(1)	sessions
     get  priority  over bulk transfers like *scp*(1) and *sftp*(1).  The queues may
     then be referenced by filtering rules (see "PACKET FILTERING" below).
     queue std bandwidth 10% cbq(default)
     queue http bandwidth 60% priority 2 cbq(borrow red) \
	   { employees, developers }
     queue  developers bandwidth 75% cbq(borrow)
     queue  employees bandwidth 15%
     queue mail bandwidth 10% priority 0 cbq(borrow ecn)
     queue ssh bandwidth 20% cbq(borrow) { ssh_interactive, ssh_bulk }
     queue  ssh_interactive bandwidth 50% priority 7 cbq(borrow)
     queue  ssh_bulk bandwidth 50% priority 0 cbq(borrow)
     block return out on dc0 inet all queue std
     pass out on dc0 inet proto tcp from $developerhosts to any port 80 \
	   queue developers
     pass out on dc0 inet proto tcp from $employeehosts to any port 80 \
	   queue employees
     pass out on dc0 inet proto tcp from any to any port 22 \
	   queue(ssh_bulk, ssh_interactive)
     pass out on dc0 inet proto tcp from any to any port 25 \
	   queue mail
**TRANSLATION**
     Translation rules modify either the source or destination	address  of  the
     packets  associated  with	a stateful connection.	A stateful connection is
     automatically created to track packets matching such a rule as long as they
     are not blocked by the filtering section of **pf.conf**.  The	translation  en-
     gine modifies the specified address and/or port in the packet, recalculates
     IP,  TCP and UDP checksums as necessary, and passes it to the packet filter
     for evaluation.
     Since translation occurs before filtering the filter engine will see  pack-
     ets  as they look after any addresses and ports have been translated.  Fil-
     ter rules will therefore have to filter based on the translated address and
     port number.  Packets that match a translation rule are only  automatically
     passed  if  the *pass* modifier is given, otherwise they are still subject to
     *block* and *pass* rules.
     The state entry created permits *pf*(4) to keep track of the original address
     for traffic associated with that state and correctly direct return  traffic
     for that connection.
     Various types of translation are possible with pf:
     *binat*
	   A *binat* rule specifies a bidirectional mapping between an external IP
	   netblock and an internal IP netblock.
     *nat*   A  *nat*  rule  specifies  that  IP  addresses are to be changed as the
	   packet traverses the given interface.  This technique allows  one  or
	   more  IP addresses on the translating host to support network traffic
	   for a larger range of machines on an "inside" network.   Although  in
	   theory  any IP address can be used on the inside, it is strongly rec-
	   ommended that one of the address ranges defined by RFC 1918 be  used.
	   These netblocks are:
	   10.0.0.0 - 10.255.255.255 (all of net 10, i.e., 10/8)
