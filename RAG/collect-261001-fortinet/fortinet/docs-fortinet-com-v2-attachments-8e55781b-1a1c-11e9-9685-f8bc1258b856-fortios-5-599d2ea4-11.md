---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-11
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2017-10-11"]
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [1250, 1341]
sha256: 4e5960d25e2a81caf52171f5f959d476f243ca72a284a5c9110266ed56e7d08e
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

Primaryunitselectionwithoverridedisabled(default) AnintroductiontotheFGCP
featureisjustthis, tomakesurethat theprimaryunitisalwaystheclusterunitwiththemostconnectedand
operatingmonitoredinterfaces.
Primary unit selection and age
Theclusterunitwiththehighestagevaluebecomestheprimaryunit. Theageofaclusterunitistheamount of
time sinceamonitoredinterfacefailed orisdisconnected.Ageisalsoresetwhenaclusterunitstarts(bootsup).
So,whenallclusterunitsstartupataboutthesametime, theyallhavethesameage.Agedoesnotaffect
primaryunitselectionwhenallclusterunitsstartupatthesametime. Agealsotakesprecedenceoverpriorityfor
primaryunitselection.
If alinkfailureofamonitoredinterfaceoccurs,theagevaluefortheclusterunitthat experiencesthelinkfailureis
reset.So,theclusterunitthat experiencedthelinkfailurealsohasaloweragevaluethantheotherclusterunits.
Thisreducedagedoesnoteffect primaryunitselectionbecausethenumberoflinkfailurestakesprecedence
overtheage.
If thefailed monitoredinterfaceisrestoredtheclusterunitthat hadthefailed monitoredinterfacecannotbecome
theprimaryunitbecauseitsageisstill lowerthantheageoftheotherclusterunits.
Inmostcases,thewaythat ageishandledbytheclusterreducesthenumberoftimes theclusterselectsanew
primaryunit, whichresultsinamorestableclustersinceselectinganewprimaryunithasthepotential todisrupt
traffic.
Cluster age difference margin (grace period)
Inanycluster,someoftheclusterunitsmaytakelongertostartupthanothers.Thisstartuptime differencecan
happenasaresultofanumberofissuesanddoesnotaffect thenormaloperationofthecluster.Tomakesure
that clusterunitsthat startslowercanstill becomeprimaryunits, bydefault theFGCPignoresagedifferencesof
upto5minutes(300seconds).
Inmostcases,duringnormaloperationthisagedifferencemarginorgraceperiodhelpsclustersfunctionas
expected.However,theagedifferencemargincanresultinsomeunexpectedbehaviorinsomecases:
l Duringaclusterfirmwareupgradewithuninterruptible- upgrade enabled(thedefault configuration)the
clustershouldnotselectanewprimaryunitafterthefirmwareofallclusterunitshasbeenupdated. Butsincethe
agedifferenceoftheclusterunitsismostlikelylessthan300seconds,ageisnotusedtoaffect primaryunit
selectionandtheclustermayselectanewprimaryunit.
l Duringfailovertesting whereclusterunitsarefailed overrepeatedlytheagedifferencebetweentheclusterunitswill
mostlikelybelessthan5minutes. Duringnormaloperation,ifafailoveroccurs,whenthefailed unitrejoinsthe
clusteritsagewillbeverydifferent fromtheageofthestill operatingclusterunitssotheclusterwillnotselectanew
primaryunit. However,ifaunitfailsandisrestoredinaveryshorttime theagedifferencemaybelessthan5
minutes. Asaresulttheclustermayselectanewprimaryunitduringsomefailovertesting scenarios.
Changing the cluster age difference margin
Youcanchangetheclusteragedifferencemarginusingthefollowing command:
config system ha
set ha-uptime-diff-margin 60
end
Thiscommandsetstheclusteragedifferencemarginto60seconds(1minute). Theagedifferencemarginrange
1to65535seconds.Thedefault is300seconds.
40 HighAvailability
Fortinet TechnologiesInc.

AnintroductiontotheFGCP Primaryunitselectionwithoverridedisabled(default)
Youmaywanttoreducethemarginifduringfailovertesting youdon’twanttowaitthedefault agedifference
marginof5minutes. Youmayalsowanttoreducethemargintoallowuninterruptible upgradestowork.See
Operatingclustersandvirtualclustersonpage168.
Youmaywanttoincreasetheagemarginifclusterunitstartuptime differencesarelargerthan5minutes.
Displaying cluster unit age differences
YoucanusetheCLIcommanddiagnose sys ha dump-by group todisplaytheagedifferenceofthe
unitsinacluster.Thiscommandalsodisplaysinformation aboutanumberofHA-relatedparametersforeach
clusterunit.
Forexample,consideraclusteroftwoFortiGate-600Dunits. Enteringthediagnose sys ha dump-by
group commandfromtheprimaryunitCLIdisplaysinformation similartothefollowing:
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
'FGT6HD3916801955': ha_prio/o=1/1, link_failure=0(old=0), pingsvr_failure=0,
flag=0x00000000, uptime/reset_ cnt=0/1
'FGT6HD3916800525': ha_prio/o=0/0, link_failure=0(old=0), pingsvr_failure=0,
flag=0x00000001, uptime/reset_ cnt=189/0
Thelasttwolinesoftheoutput displaystatusinformation abouteachclusterunitincludingtheuptime.The
uptime istheagedifferenceinsecondsbetweenthetwounitsinthecluster.
Intheexample,theageofthesubordinateunitis189secondsmorethantheageoftheprimaryunit. Theage
differenceislessthan5minutes(lessthan300seconds)soagehasnoaffect onprimaryunitselection.The
clusterselectedtheunitwiththehighestserialnumbertobetheprimaryunit.
If port1(themonitoredinterface)oftheprimaryunitisdisconnected,theclusterrenegotiatesandtheformer
subordinateunitbecomestheprimaryunit. WhenyoulogintothenewprimaryunitCLIandenterdiagnose
sys ha dump-by group youcouldgetresultssimilartothefollowing:
diagnose sys ha dump-by group
HA information.
group-id=0, group-name='External- HA-Cluster'
gmember_nr=2
'FGT6HD3916800525': ha_ip_idx=1, hb_packet_version=6, last_hb_jiffies=52097155,
linkfails=11, weight/o=0/0
hbdev_nr=2: port3(mac=906c..70, last_hb_jiffies=52097155, hb_lost=0), port4
(mac=906c..71, last_hb_jiffies=52097155, hb_lost=0),
HighAvailability
Fortinet TechnologiesInc.
41

