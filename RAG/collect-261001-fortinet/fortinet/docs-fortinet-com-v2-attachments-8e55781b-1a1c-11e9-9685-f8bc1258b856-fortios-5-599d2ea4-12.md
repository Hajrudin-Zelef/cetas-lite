---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-12
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2017-10-11"]
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [1342, 1436]
sha256: ce9a4ebae7d83aa53e654f6c3546cd56b8a2067708979682d961e5d230e510cc
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

Primaryunitselectionwithoverridedisabled(default) AnintroductiontotheFGCP
'FGT6HD3916801195': ha_ip_idx=0, hb_packet_version=6, last_hb_jiffies=0,
linkfails=0, weight/o=0/0
vcluster_nr=1
vcluster_0: start_time=1507754642 (2017-10-11 13:44:02), state/o/chg_ time=2(work)/2
(work)/1507754644 (2017-10-11 13:44:04)
'FGT6HD3916800525': ha_prio/o=1/1, link_failure=0(old=0), pingsvr_failure=0,
flag=0x00000000, uptime/reset_ cnt=0/1
'FGT6HD3916801955': ha_prio/o=0/0, link_failure=0(old=0), pingsvr_failure=0,
flag=0x00000001, uptime/reset_ cnt=236/0
Thecommandresultsshowthat theageofthenewprimaryunitis236secondshigherthantheageofthenew
subordinateunit.
If port1oftheformerprimaryunitisreconnectedtheclusterwillonceagainmakethistheprimaryunitbecause
theagedifferencewillstill belessthan300seconds.WhenyoulogintotheprimaryunitCLIandenter
diagnose sys ha dump-by group yougetresultssimilartothefollowing:
diagnose sys ha dump-by group
HA information.
group-id=0, group-name='External- HA-Cluster'
gmember_nr=2
'FGT6HD3916800525': ha_ip_idx=1, hb_packet_version=6, last_hb_jiffies=52097155,
linkfails=11, weight/o=0/0
hbdev_nr=2: port3(mac=906c..70, last_hb_jiffies=52097155, hb_lost=0), port4
(mac=906c..71, last_hb_jiffies=52097155, hb_lost=0),
'FGT6HD3916801195': ha_ip_idx=0, hb_packet_version=6, last_hb_jiffies=0,
linkfails=0, weight/o=0/0
vcluster_nr=1
vcluster_0: start_time=1507754642 (2017-10-11 13:44:02), state/o/chg_ time=2(work)/2
(work)/1507754644 (2017-10-11 13:44:04)
'FGT6HD3916800525': ha_prio/o=1/1, link_failure=0(old=0), pingsvr_failure=0,
flag=0x00000000, uptime/reset_ cnt=0/1
'FGT6HD3916801955': ha_prio/o=0/0, link_failure=0(old=0), pingsvr_failure=0,
flag=0x00000001, uptime/reset_ cnt=-236/0
Resetting the age of all cluster units
Insomecases,agedifferencesamongclusterunitscanresultinthewrongclusterunitorthewrongvirtualcluster
becomingtheprimaryunit. Forexample,ifaclusterunitsettoahighpriorityreboots,that unitwillhavealower
agethanotherclusterunitswhenitrejoinsthecluster.Sinceagetakesprecedenceoverpriority,thepriorityof
thisclusterunitwillnotbeafactorinprimaryunitselection.
Thisproblemalsoaffects virtualclusterVDOMpartitioning inasimilarway.Afterarebootofoneoftheunitsina
virtualclusterconfiguration, traffic forallVDOMscouldcontinuetobeprocessedbytheclusterunitthat didnot
reboot.Thiscanhappenbecausetheageofbothvirtualclustersontheunitthat didnotrebootisgreaterthat the
ageofbothvirtualclustersontheunitthat rebooted.
Onewaytoresolvethisissueistorebootalloftheclusterunitsatthesametime sothat theageofallofthe
clusterunitsisreset.However,rebootingclusterunitsmayinterruptoratleastslowdowntraffic. If youwould
rathernotrebootalloftheclusterunitsyoucaninsteadusethefollowing commandtoresettheageofindividual
clusterunits.
diagnose sys ha reset-uptime
42 HighAvailability
Fortinet TechnologiesInc.

AnintroductiontotheFGCP Primaryunitselectionwithoverridedisabled(default)
Thiscommandresetstheageofaunitbacktozerosothat ifnootherunitintheclusterwasresetatthesame
time, itwillnowhavethelowestage.Youwouldusethiscommandtoresettheageoftheclusterunitthat is
currentlytheprimaryunit. Sinceitwillhavethelowestage,theotherunitintheclusterwillhavethehighestage
andcanthenbecometheprimaryunit.
Thediagnose sys ha reset-uptime commandshouldonlybeusedasa
temporarysolution. ThecommandresetstheHAageinternallyanddoesnotaffect the
uptime displayedforclusterunitsusingthediagnose sys ha dump-by all-
vcluster commandortheuptime displayedontheDashboardorclustermembers
list. Tomakesuretheactualuptime forclusterunitsisthesameastheHAageyou
shouldreboottheclusterunitsduringamaintenancewindow.
Primary unit selection and device priority
Aclusterunitwiththehighestdeviceprioritybecomestheprimaryunitwhentheclusterstartsuporrenegotiates.
Bydefault, thedevicepriorityforallclusterunitsis128.Youcanchangethedeviceprioritytocontrolwhich
FortiGate becomestheprimaryunitduringclusternegotiation. Allotherfactorsthat influenceprimaryunit
selectioneithercannotbeconfigured(ageandserialnumber)oraresynchronizedamongallclusterunits
(interfacemonitoring). Youcansetadifferent devicepriorityforeachclusterunit. Duringnegotiation, ifall
monitoredinterfacesareconnected,andallclusterunitsentertheclusteratthesametime (orhavethesame
age),theclusterwiththehighestdeviceprioritybecomestheprimaryunit.
Ahigherdeviceprioritydoesnotaffect primaryunitselectionforaclusterunitwiththemostfailed monitored
interfacesorwithanagethat ishigherthanallotherclusterunitsbecausefailed monitoredinterfacesandageare
usedtoselectaprimaryunitbeforedevicepriority.
Increasingthedevicepriorityofaclusterunitdoesnotalwaysguaranteethat thisclusterunitwillbecomethe
primaryunit. Duringclusteroperation,aneventthat mayaffect primaryunitselectionmaynotalwaysresultinthe
clusterrenegotiating. Forexample,whenaunitjoinsafunctioning cluster,theclusterwillnotrenegotiate. Soifa
unitwithahigherdevicepriorityjoinsaclusterthenewunitbecomesasubordinateunituntil thecluster
renegotiates.
Enablingtheoverride HACLIkeywordmakeschangesindeviceprioritymore
effective bycausingtheclustertonegotiate moreoften tomakesurethat theprimary
unitisalwaystheunitwiththehighestdevicepriority.Formoreinformation about
override,seePrimaryunitselectionwithoverridedisabled(default)onpage37.
Controlling primary unit selection by changing the device priority
Yousetadifferent devicepriorityforeachclusterunittocontroltheorderinwhichclusterunitsbecomethe
primaryunitwhentheprimaryunitfails.
TochangethedevicepriorityfromtheGUIgoto System > HA andchangethe Device priority.
Enterthefollowing CLIcommandtochangethedevicepriorityto200:
config system ha
set priority 200
end
Thedevicepriorityisnotsynchronizedamongclusterunits. Inafunctioning clusteryoucanchangethedevice
priorityofanyunitinthecluster.Wheneveryouchangethedevicepriorityofaclusterunit, whenthecluster
negotiates, theunitwiththehighestdeviceprioritybecomestheprimaryunit.
HighAvailability
Fortinet TechnologiesInc.
43

