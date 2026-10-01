---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-54
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2018-03-06"]
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [7174, 7275]
sha256: f1be45e2973017e526910dc68c7a5d17e08886ab405f8dd084b9884a6d6eb516
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

Operatingclustersandvirtualclusters ViewingclusterstatusfromtheCLI
Thefollowing commandoutput wasproducedbyusingexecute ha manage 1tologintothesubordinate
unitCLIoftheclustershowninthepreviousexample.ThehostnameofthesubordinateunitisEdge2-
Backup.
get system ha status
HA Health Status: OK
Model: FortiGate- 600D
Mode: HA A-P
Group: 25
Debug: 0
Cluster Uptime: 0 days 03:33:04
Cluster state change time: 2018-03-06 13:16:33
Master selected using:
<2018/03/06 13:16:33> FGT6HD3916806098 is selected as the master because it
has the largest value of override priority.
<2018/03/06 12:47:58> FGT6HD3916806070 is selected as the master because it
has the largest value of override priority.
<2018/03/06 12:47:57> FGT6HD3916806098 is selected as the master because it
has the largest value of uptime.
<2018/03/06 12:47:56> FGT6HD3916806098 is selected as the master because it
has the largest value of uptime.
ses_pickup: enable, ses_pickup_delay=disable
override: disable
Configuration Status:
FGT6HD3916806070 (updated 1 seconds ago): in-sync
FGT6HD3916806098 (updated 1 seconds ago): in-sync
System Usage stats:
FGT6HD3916806070 (updated 1 seconds ago):
sessions=20, average-cpu-user/nice/system/idle=0%/0%/0%/100%, memory=34%
FGT6HD3916806098 (updated 1 seconds ago):
sessions=163, average-cpu-user/nice/system/idle=0%/0%/0%/100%, memory=34%
HBDEV stats:
FGT6HD3916806070 (updated 1 seconds ago):
port3: physical/1000full, up, rx-bytes/pack -
ets/dropped/errors=40755112/71809/0/0, tx=48104698/76943/0/0
port4: physical/1000full, up, rx-bytes/pack -
ets/dropped/errors=29804904/42302/0/0, tx=30030641/42333/0/0
FGT6HD3916806098 (updated 1 seconds ago):
port3: physical/1000full, up, rx-bytes/pack -
ets/dropped/errors=47188898/74965/0/0, tx=39680065/69723/0/0
port4: physical/1000full, up, rx-bytes/pack -
ets/dropped/errors=29338501/41347/0/0, tx=29022293/41201/0/0
Slave : Edge2-Backup , FGT6HD3916806070, cluster index = 1
Master: Edge2-Primary , FGT6HD3916806098, cluster index = 0
number of vcluster: 1
vcluster 1: standby 169.254.0.1
Slave : FGT6HD3916806070, operating cluster index = 1
Master: FGT6HD3916806098, operating cluster index = 0
About the HA operating cluster index and the execute ha manage command
Whenaclusterstartsup,ifprimaryunitselectisbasedonserialnumber,theFortiGate ClusterProtocol(FGCP)
assignsaclusterindexandanHAheartbeatIPaddresstoeachclusterunitbasedontheserialnumberofthe
HighAvailability
Fortinet TechnologiesInc.
197

ViewingclusterstatusfromtheCLI Operatingclustersandvirtualclusters
clusterunit:
l TheFGCPselectstheclusterunitwiththehighestserialnumbertobecometheprimaryunit. TheFGCPassignsa
clusterindexof0,anoperatingclusterindexof0,andanHAheartbeatIPaddressof169.254.0.1 tothisunit.
l TheFGCPassignsaclusterindexof1,anoperatingclusterindexof1,andanHAheartbeatIPaddressof
169.254.0.2 totheclusterunitwiththesecondhighestserialnumber.
l If theclustercontainsmoreunits, theclusterunitwiththethirdhighestserialnumberisassignedaclusterindexof
2,andoperatingclusterindexof2,andanHAheartbeatIPaddressof169.254.0.3, andsoon.
Youcandisplaytheclusterindexandoperatingclusterindexassignedtoeachclusterunitusingtheget
system ha status command. Whenyouusetheexecute ha manage commandyouselectaclusterunit
tologintobyenteringitsoperatingclusterindex.
TheoperatingclusterindexandHAheartbeatIPaddressonlychangeifaunitleavestheclusterorifanewunit
joinsthecluster.Whenoneoftheseeventshappens,theFGCPresetstheclusterindex,operatingclusterindex,
andHAheartbeatIPaddressofeachclusterunitaccordingtoserialnumberinthesamewayaswhenthecluster
firststartsup.
If FortiGates don'tleaveorjoin, eachclusterunitkeepsitsassignedoperatingclusterindex,andHAheartbeatIP
addresssincethesearebasedontheFortiGate serialnumber,evenastheunitstakeondifferent rolesinthe
cluster.AftertheoperatingclusterindexandHAheartbeatIPaddressesaresetaccordingtoserialnumber,the
FGCPchecksotherprimaryunitselectioncriteriasuchasdevicepriorityandmonitoredinterfaces.Checking
thesecriteriacouldresultinselectingaclusterunitwithout thehighestserialnumbertooperateastheprimary
unit.
Eveniftheclusterunitwithout thehighestserialnumbernowbecomestheprimaryunit, theoperatingcluster
indexesandHAheartbeatIPaddressesassignedtotheindividualclusterunitsdonotchange.InsteadtheFGCP
changestheclusterindextoreflectthisrolechange.Theclusterindexisalways0fortheprimaryunitand1and
higherfortheotherunitsinthecluster.Bydefault bothsetsofclusterindexesarethesame.Butifprimaryunit
selectionselectstheclusterunitthat doesnothavethehighestserialnumbertobetheprimaryunit, thenthis
clusterunitisassignedaclusterindexof0.
Using the execute ha manage command
WhenyouusetheCLIcommandexecute ha manage <index_integer> toconnecttotheCLIofanother
clusterunit, the<index_integer> that youenteristheoperatingclusterindexoftheunitthat youwantto
connectto.
Using get system ha status to display cluster indexes
YoucandisplaytheclusterindexassignedtoeachclusterunitusingtheCLIcommandget system ha
status.Thefollowing exampleshowstheinformation displayedbytheget system ha status command
foraclusterconsistingoftwoFortiGates operatinginactive-passiveHAmodewithvirtualdomainsnotenabled
andwithout virtualclustering.
get system ha status
.
.
.
Master: Edge2-Primary , FGT6HD3916806098, cluster index = 0
Slave : Edge2-Backup , FGT6HD3916806070, cluster index = 1
number of vcluster: 1
vcluster 1: work 169.254.0.1
198 HighAvailability
Fortinet TechnologiesInc.

