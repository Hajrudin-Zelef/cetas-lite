---
id: collect-261001-fortinet/fortinet/en-snmp-snmp-fortinet-61b79530
title: "en-snmp-snmp-fortinet-61b79530"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/en-snmp-snmp-fortinet-61b79530.md
source_anchor: ""
source_lines: [1, 58]
sha256: 94c554e6f1feeaa9e1273a25bddef4d5fe318f2aa073daf4e3d8a6a6d225a394
---

# en-snmp-snmp-fortinet-61b79530

2026/05/04 17:10 1/4 SNMP activation on a Fortinet ﬁrewall
Esia Wiki - https://wiki.esia-sa.com/
SNMP activation on a Fortinet ﬁrewall
This tutorial has been made available to the entire Esia community thanks to the contribution of our
partner Rcarré.
Their website: https://www.rcarre.com
Via the WEB interface
Once you have logged in, you will be taken to the ﬁrewall dashboard as shown in the image below.
Click on “System” and then on “SNMP” to go to the SNMP conﬁguration page. As shown below:

Last update: 2026/05/04 15:19 en:snmp:snmp_fortinet https://wiki.esia-sa.com/en/snmp/snmp_fortinet
https://wiki.esia-sa.com/ Printed on 2026/05/04 17:10
Tick the “Enable” box and enter the description, location and contact. Then click on “Apply”. Now you
need to create the SNMP community. Just below the “Apply” button, click on “Create New”.
On the page that appears, enter the SNMP community, the IP address of your Esia server or your unity
in the HOST ﬁeld and tick the boxes as shown below. Then click on “Apply”.

2026/05/04 17:10 3/4 SNMP activation on a Fortinet ﬁrewall
Esia Wiki - https://wiki.esia-sa.com/
Now you need to authorise the SNMP protocol on the LAN interface of your ﬁrewall. To do this, go to
the “Network” menu and then “Interface”. Then tick the SNMP box in “Restrict Access”.
Click “Apply” to save the conﬁguration.
SNMP is now enabled on your Fortigate ﬁrewall.
Via CLI/SSH
Once connected via SSH, you can type the following commands to activate SNMP. You will obviously
need to adapt the description/contact/location ﬁelds.
config system snmp sysinfo
    set status enable
    set description "ce que je veux"
    set contact-info "absent"

Last update: 2026/05/04 15:19 en:snmp:snmp_fortinet https://wiki.esia-sa.com/en/snmp/snmp_fortinet
https://wiki.esia-sa.com/ Printed on 2026/05/04 17:10
    set location "Liège"
end
Now that SNMP has been activated, we need to conﬁgure the SNMP community using the following
commands:
config system snmp community
    edit 1
        set name "public"
        config hosts
            edit 1
            next
        end
        set events cpu-high mem-low log-full intf-ip vpn-tun-up vpn-tun-down
ha-switch ha-hb-failure ips-signature ips-anomaly av-virus av-oversize av-
pattern av-fragmented fm-if-change bgp-established bgp-backward-transition
ha-member-up ha-member-down ent-conf-change av-conserve av-bypass av-
oversize-passed av-oversize-blocked ips-pkg-update ips-fail-open faz-
disconnect wc-ap-up wc-ap-down
    next
end
Don't forget to change the community to “public”.
From:
https://wiki.esia-sa.com/ - Esia Wiki
Permanent link:
https://wiki.esia-sa.com/en/snmp/snmp_fortinet
Last update: 2026/05/04 15:19
