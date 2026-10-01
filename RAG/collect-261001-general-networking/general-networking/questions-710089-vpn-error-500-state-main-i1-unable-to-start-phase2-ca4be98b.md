---
id: collect-261001-general-networking/general-networking/questions-710089-vpn-error-500-state-main-i1-unable-to-start-phase2-ca4be98b
title: "leftxauthclient=yes"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-710089-vpn-error-500-state-main-i1-unable-to-start-phase2-ca4be98b.md
source_anchor: ""
source_lines: [1, 44]
sha256: 6cffff16f0013626330721f532cd9724269b792b10e75ce764038a1a9de4c5ee
---

# leftxauthclient=yes

i'm trying to set up a site to site vpn to a fortigate 60c from a CentOS 7 with openswan, the error i get everytime is the following
000 #1: "office":500 STATE_MAIN_I1 (sent MI1, expecting MR1);
EVENT_v1_RETRANSMIT in 8s; nodpd; idle; import:admin initiate
000 #1: pending Phase 2 for "office" replacing #0
my configuration files
office.conf
conn office
left=%defaultroute              # Your local linux machine IP
leftsubnet=192.168.3.0/24       # The subnet of your local Linux machine
leftid=@openswan               # Same as given in Sonicwall
leftnexthop=%defaultroute
#    leftxauthclient=yes
right=mrt.mx          # Sonicwall VPN IP
rightsubnet=192.168.1.0/24     # Sonicwall LAN subnet
rightid=office          # Sonicwall Unique Identifier
#    rightxauthserver=yes
#    keyingtries=0
#    pfs=yes
auto=start
auth=esp
esp=3DES-SHA1                 
ike=3DES-SHA1
ikelifetime=1800s
authby=secret
aggrmode=no
#    leftmodecfgclient=yes
dpddelay=30
dpdtimeout=60
ipsec.conf
  GNU nano 2.3.1                              Fichero: /etc/ipsec.conf                                                                      
version 2.0     # conforms to second version of ipsec.conf specification
# basic configuration
config setup
# Debug-logging controls:  "none" for (almost) none, "all" for lots.
# klipsdebug=none
# plutodebug="control parsing"
# For Red Hat Enterprise Linux and Fedora, leave protostack=netkey
protostack=netkey
nat_traversal=yes
interfaces=%defaultroute
oe=off
# Enable this if you see "failed to find any available worker"
nhelpers=0
include /etc/ipsec.d/*.conf
