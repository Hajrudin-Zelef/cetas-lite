---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-89
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [11073, 11238]
sha256: 131d8d5eccaa1035a11c29edad135528e94afba51625ac2e9f16748c72ef045f
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

FortiGate SessionLifeSupportProtocol(FGSP) Basicexampleconfiguration
Basic example configuration
Thefollowing configurationexampleshowshowtoconfigurebasicFGSPforthetwopeerFortiGates shown
below.
l Thehostnamesofpeersarepeer_1andpeer_2.
l Bothpeersareconfiguredwithtwovirtualdomains:rootandvdom_1.
l Allsessionsprocessedbyvdom_1aresynchronized.
l Thesynchronizationlinkinterfaceisport3whichisintherootvirtualdomain.
l TheIPaddressofport3onpeer_1is10.10.10.1.
l TheIPaddressofport3onpeer_2is10.10.10.2.
Alsoonbothpeers,port1andport2areaddedtovdom_1.Onpeer_1theIPaddressofport1issetto
192.168.20.1 andtheIPaddressofport2issetto172.110.20.1. Onpeer_2theIPaddressofport1issetto
192.168.20.2 andtheIPaddressofport2issetto172.110.20.2.
Example FGSP network configuration
To configure FGSP:
1. Configuretheloadbalancerorroutertosendallsessionstopeer_1.
2. Configuretheloadbalancerorroutertosendalltraffic topeer_2ifpeer_1fails.
3. UsenormalFortiGate configurationstepsonpeer_1:
l Enablevirtualdomainconfiguration.
l Addthevdom_1virtualdomain.
l Addport1andport2tothevdom_1virtualdomainandconfiguretheseinterfaces.
l SettheIPaddressofport1to192.168.20.1.
l SettheIPaddressofport2to172.110.20.1.
HighAvailability
Fortinet TechnologiesInc.
297

Basicexampleconfiguration FortiGate SessionLifeSupportProtocol(FGSP)
l SettheIPaddressofport3to10.10.10.1.
l Addroutemodesecuritypoliciesbetweenport1andport2tovdom_1.
4. Enterthefollowing commandstoconfiguresessionsynchronizationforpeer_1:
config system cluster-sync
edit 1
set peerip 10.10.10.2
set peervd root
set syncvd vdom_1
end
5. UsenormalFortiGate configurationstepsonpeer_2:
l Enablevirtualdomainconfiguration.
l Addthevdom_1virtualdomain.
l Addport1andport2tothevdom_1virtualdomainandconfiguretheseinterfaces.
l SettheIPaddressofport1to192.168.20.2.
l SettheIPaddressofport2to172.110.20.2.
l SettheIPaddressofport3to10.10.10.1.
l Addroutemodesecuritypoliciesbetweenport1andport2tovdom_1.
6. Enterthefollowing commandtoconfiguresessionsynchronizationforpeer_1:
config system cluster-sync
edit 1
set peerip 10.10.10.1
set peervd root
set syncvd vdom_1
end
To add filters:
Youcanaddafilter tothisbasicconfigurationifyouonlywanttosynchronizesomeTCPsessions.Forexample
youcanenterthefollowing commandtoaddafilter sothat onlyHTTPsessionsaresynchronized:
config system cluster-sync
edit 1
config filter
set service HTTP
end
end
Youcanalsoaddafilter tocontrolthesourceanddestination addressesoftheIPv4packetsthat are
synchronized.Forexample,youcanenterthefollowing toaddafilter sothat onlysessionswithsourceaddresses
intherange10.10.10.100 to10.10.10.200 aresynchronized:
config system cluster-sync
edit 1
config filter
set srcaddr 10.10.10.100 10.10.10.200
end
end
Youcanalsoaddafilter tocontrolthesourceanddestination addressesoftheIPv6packetsthat are
synchronized.Forexample,youcanenterthefollowing toaddafilter sothat onlysessionswithdestination
addressesintherange2001:db8:0:2::/64 aresynchronized:
config system cluster-sync
edit 1
298 HighAvailability
Fortinet TechnologiesInc.

FortiGate SessionLifeSupportProtocol(FGSP) VerifyingtheFGSPconfigurationandsynchronization
config filter
set dstaddr6 2001:db8:0:2::/64
end
end
To synchronize TCP sessions:
Enterthefollowing tosynchronizeTCPsessionsandsetthesynchronizationlink(heartbeatdevice):
config system ha
set hbdev "port3" 50
set session-pickup enable
end
To synchronize UDP and ICMP sessions:
Enterthefollowing toaddsynchronizationofUDPandICMPsessionstothisconfiguration:
config system ha
set session-pickup enable
set session-pickup-connectionless enable
end
Verifying the FGSP configurationand synchronization
Youcanusethefollowing diagnosecommandstoverifythat theFGSPanditssynchronizationfunctionsare
operatingcorrectly.
FGSP configuration summary and status
Enterthefollowing commandtodisplayasummaryoftheFGSPconfigurationandsynchronizationstatus:
diagnose sys session sync
sync_ctx: sync_started=1, sync_tcp=1, sync_others=1,
sync_expectation=1, sync_redir=0, sync_nat=1, stdalone_sessync=0.
sync: create=12:0, update=0, delete=0:0, query=14
recv: create=14:0, update=0, delete=0:0, query=12
ses pkts: send=0, alloc_fail=0, recv=0, recv_err=0 sz_err=0
nCfg_sess_sync_num=5, mtu=16000
sync_filter:
1: vd=0, szone=0, dzone=0, saddr=0.0.0.0:0.0.0.0,daddr=0.0.0.0:0.0.0.0,
sync_started=1 showsthat synchronizationisworking.If thisissetto0thensomethingisnotcorrectwith
sessionsynchronizationandsynchronizationhasnotbeenabletostartbecauseofit.
sync_tcp=1,sync_others=1,sync_expectation=1 ,andsync_nat=1 showthat theFGSPhasbeen
configuredtosynchronizeTCP,connectionless,asymmetric, andNATsessions.
sync: create=12:0 andrecv: create=14:0 showthat thisFortiGate hassynchronized12sessionsto
itspeerandhasreceived14sessionsfromitspeer.
sync_filter showstheconfiguredFGSPfilter. Inthiscasenofilter hasbeencreatedsoallsessionsare
synchronized.
HighAvailability
Fortinet TechnologiesInc.
299

Synchronizingtheconfiguration FortiGate SessionLifeSupportProtocol(FGSP)
vd=0 indicatesthat rootVDOMsessionsaresynchronized.
Verifying that sessions are synchronized
Enterthecommanddiagnose sys session list todisplayinformation aboutthesessionsbeing
processedbytheFortiGate. Inthecommandoutput lookforsessionsthat shouldbesynchronizedandmakesure
theycontainoutput linesthat includesynced forexample,state=log may_dirty ndr synced,to
confirmthat theyarebeingsynchronizedbytheFGSP.
diagnose sys session list
session info: proto=6 proto_state=05 duration=469 expire=0 timeout=3600
flags=00000000 sockflag=00000000 sockport=21 av_idx=0 use=4
origin-shaper=
reply-shaper=
per_ip_shaper=
ha_id=0 policy_dir=0 tunnel=/
state=log may_dirty ndr synced
statistic(bytes/packets/allow_err): org=544/9/1 reply=621/7/0 tuples=2
orgin->sink: org pre->post, reply pre->post dev=46->45/45->46
gwy=10.2.2.1/10.1.1.1
hook=pre dir=org act=noop 192.168.1.50:45327->172.16.1.100:21(0.0.0.0:0)
hook=post dir=reply act=noop 172.16.1.100:21->192.168.1.50:45327(0.0.0.0:0)
pos/(before,after) 0/(0,0), 0/(0,0)
misc=0 policy_id=1 id_policy_id=0 auth_info=0 chk_client_info=0 vd=0
serial=00002deb tos=ff/ff ips_view=1 app_list=2000 app=16427
dd_type=0 dd_mode=0
per_ip_bandwidth meter: addr=192.168.1.50,bps=633
Synchronizingthe configuration
YoucanconfiguresynchronizationfromonestandaloneFortiGate toanotherstandaloneFortiGate
(standalone- config-sync).Withtheexceptionofsomeconfigurationsthat donotsync,therestofthe
configurationsaresynced,suchasfirewallpolicies,firewalladdresses,andUTMprofiles.
ThisoptionisusefulinsituationswhenyouneedtosetupFGSPpeers,orwhenyouwanttoquicklydeploy
severalFortiGates withthesameconfigurations.Youcansetupstandalone- config-sync formultiple
members.
Bydefault, configurationsynchronizationisdisabled.Youcanenterthefollowing toenableit:
config system ha
set standalone-config-sync enable
end
standalone- config-sync isanindependentfeatureandshouldbeusedwith
cautionastherearesomelimitations. Fortinet recommendsdisablingitoncethe
configurationshavebeensyncedover.
YoumustenterthiscommandonalloftheFortiGates inthegroup.Whenyouenablesynchronizingthe
configuration, FGCPprimaryunitselectionisusedtoselectaprimary(ormaster)FortiGate (seePrimaryunit
selectionwithoverridedisabled(default)onpage37).TheotherFortiGates inthedeploymentbecomebackup
300 HighAvailability
Fortinet TechnologiesInc.

