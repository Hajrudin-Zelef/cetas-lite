---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-53
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2018-03-06"]
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [7020, 7173]
sha256: 107598bd6f1205df3b337a2f1a8d8787f08db791242dd5df6993e8734c061d7c
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

ViewingclusterstatusfromtheCLI Operatingclustersandvirtualclusters
Field Description
Master selected
using
Showshowtheprimaryunitwasselectedthelastfourtimes that thecluster
negotiated. Forexample,whenaclusterfirstforms, thispartofthecommand
output couldhaveonelineshowingthat theprimaryunitistheclusterunitwith
thehighestuptime. Uptofourlinescanbeincludedastheclusternegotiatesto
chooseanewprimaryunitondifferent occasions.Eachlineincludesatime stamp
andthecriteriausedtoselecttheprimaryunit.
ses_pickup Thestatusofsessionpickup:enableordisable.
load_balance Thestatusoftheload-balance-all keyword:enableordisable.Active-
activeclustersonly.
load_balance_udp Thestatsoftheload-balance-udp keyword:enableordisable.Availableon
someFortiGate models.Active-activeclustersonly.
schedule Theactive-activeloadbalancingschedule.Active-activeclustersonly.
override Thestatusoftheoverrideoptionforthecurrentclusterunit: enableordisable.
Configuration
Status Showsiftheconfigurationsofeachoftheclusterunitsaresynchronizedornot.
System Usage
stats
Showshowbusyeachclusterunitisbydisplayingthenumberofsessionsbeing
processedbytheclusterunit, CPUusage,andmemoryusage.
HBDEV stats
Showsthestatusofeachclusterunit'sheartbeatinterfaces.Includeswhetherthe
interfacesareupordown,howmuchdatatheyhaveprocessed,aswellaserrors
found.
Master
Slave
Displaysthehostname, serialnumber,andclusterindexoftheprimaryunit
(master)andthesubordinateunits(slave).TheFortiGate withclusterindex0is
theprimaryunitandtheFortiGates withclusterindexes1to3arethebackup
units.
Theorderinwhichtheclusterunitsarelistedstartswiththeclusterunitthat you
areloggedinto.
number of
vcluster
Thenumberofvirtualclusters.If virtualdomainsarenotenabled,theclusterhas
onevirtualcluster.If virtualdomainsareenabledtheclusterhastwovirtual
clusters.
vcluster 1
vcluster 2
TheheartbeatinterfaceIPaddressoftheprimaryunitineachvirtualcluster.If
virtualdomainsarenotenabledthereisonevclusterandthisistheIPaddressof
theprimaryunit. If virtualdomainsareenabledtheneachvclusterlinewillhave
anIPaddress.If theIPaddressesarethesamethenthesameFortiGate isthe
primaryunitforbothvirtualclusters.
194 HighAvailability
Fortinet TechnologiesInc.

Operatingclustersandvirtualclusters ViewingclusterstatusfromtheCLI
Field Description
vcluster 1
Master
Slave
TheHAstate(hello,work,orstandby)andHAheartbeatIPaddressoftheprimary
unit. If virtualdomainsarenotenabled,vcluster 1 displaysinformation for
thecluster.If virtualdomainsareenabled,vcluster 1 displaysinformation for
virtualcluster1.
vcluster 1 alsoliststheprimaryunit(master)andsubordinateunits(slave)in
virtualcluster1.Thelistincludestheserialnumberandoperatingclusterindexof
eachclusterunitinvirtualcluster1.Theclusterunitthat youhaveloggedintoisat
thetopofthelist. TheFortiGate intheclusterwiththehighestserialnumber
alwayshasanoperatingclusterindexof0.OtherFortiGates intheclustergeta
higheroperatingclusterindexbasedintheirserialnumber.Whenyouusethe
execute ha manage commandtologintoanotherFortiGate youusethe
operatingclusterindextospecifytheFortiGate tologinto.
If virtualdomainsarenotenabledandyouconnecttotheprimaryunitCLI, theHA
stateoftheclusterunitinvirtualcluster1iswork.Thedisplayliststhecluster
unitsstartingwiththeprimaryunit.
If virtualdomainsarenotenabledandyouconnecttoasubordinateunitCLI, the
HAstateoftheclusterunitinvirtualcluster1isstandby.Thedisplayliststhe
clusterunitsstartingwiththesubordinateunitthat youhaveloggedinto.
If virtualdomainsareenabledandyouconnecttothevirtualcluster1primaryunit
CLI, theHAstateoftheclusterunitinvirtualcluster1iswork.Thedisplayliststhe
clusterunitsstartingwiththevirtualcluster1primaryunit.
If virtualdomainsareenabledandyouconnecttothevirtualcluster1subordinate
unitCLI, theHAstateoftheclusterunitinvirtualcluster1isstandby.Thedisplay
liststheclusterunitsstartingwiththesubordinateunitthat youareloggedinto.
vcluster 2
Master Slave
vcluster 2 onlyappearsifvirtualdomainsareenabled.vcluster 2
displaystheHAstate(hello,work,orstandby)andHAheartbeatIPaddressofthe
clusterunitthat youhaveloggedintoinvirtualcluster2.TheHAheartbeatIP
addressis169.254.0.2 ifyouareloggedintotheprimaryunitofvirtualcluster2
and169.254.0.1 ifyouareloggedintoasubordinateunitofvirtualcluster2.
vcluster 2 alsoliststheprimaryunit(master)andsubordinateunits(slave)in
virtualcluster2.Thelistincludestheclusterindexandserialnumberofeach
clusterunitinvirtualcluster2.Theclusterunitthat youhaveloggedintoisatthe
topofthelist.
If youconnecttothevirtualcluster2primaryunitCLI, theHAstateofthecluster
unitinvirtualcluster2iswork.Thedisplayliststheclusterunitsstartingwiththe
virtualcluster2primaryunit.
If youconnecttothevirtualcluster2subordinateunitCLI, theHAstateofthe
clusterunitinvirtualcluster2isstandby.Thedisplayliststheclusterunits
startingwiththesubordinateunitthat youareloggedinto.
HighAvailability
Fortinet TechnologiesInc.
195

ViewingclusterstatusfromtheCLI Operatingclustersandvirtualclusters
Get system ha status example - two FortiGates in active-passive mode
Thefollowing exampleshowsget system ha status output foraclusteroftwoFortiGate-600Dsoperating
inactive-passivemode. Theclusterishealthyandhasbeenrunningfor3hoursand26minutes. Primaryunit
selectiontookplaceonceandtheclusterhasbeenstablesincethen.
Thefollowing commandoutput wasproducedbyconnectingtotheprimaryunitCLI(hostnameEdge2-
Primary).
get system ha status
HA Health Status: OK
Model: FortiGate-600D
Mode: HA A-P
Group: 25
Debug: 0
Cluster Uptime: 0 days 03:26:00
Cluster state change time: 2018-03-06 13:16:33
Master selected using:
<2018/03/06 13:16:33> FGT6HD3916806098 is selected as the master because it has the
largest value of override priority.
<2018/03/06 12:47:58> FGT6HD3916806070 is selected as the master because it has the
largest value of override priority.
<2018/03/06 12:47:55> FGT6HD3916806098 is selected as the master because it has the
largest value of uptime.
<2018/03/06 12:47:55> FGT6HD3916806098 is selected as the master because it's the only
member in the cluster.
ses_pickup: enable, ses_pickup_delay=disable
override: disable
Configuration Status:
FGT6HD3916806098(updated 1 seconds ago): in-sync
FGT6HD3916806070(updated 2 seconds ago): in-sync
System Usage stats:
FGT6HD3916806098(updated 1 seconds ago):
sessions=141, average-cpu-user/nice/system/idle=0%/0%/0%/100%,memory=34%
FGT6HD3916806070(updated 2 seconds ago):
sessions=12, average-cpu-user/nice/system/idle=0%/0%/0%/100%,memory=33%
HBDEV stats:
FGT6HD3916806098(updated 1 seconds ago):
port3: physical/1000full,up, rx-bytes/packets/dropped/errors=45437370/71531/0/0,
tx=36186194/65035/0/0
port4: physical/1000full,up, rx-bytes/packets/dropped/errors=27843923/39221/0/0,
tx=27510707/39075/0/0
FGT6HD3916806070(updated 2 seconds ago):
port3: physical/1000full,up, rx-bytes/packets/dropped/errors=37267057/67136/0/0,
tx=46354380/73516/0/0
port4: physical/1000full,up, rx-bytes/packets/dropped/errors=28294029/40177/0/0,
tx=28536766/40208/0/0
Master: Edge2-Primary , FGT6HD3916806098, cluster index = 0
Slave : Edge2-Backup , FGT6HD3916806070, cluster index = 1
number of vcluster: 1
vcluster 1: work 169.254.0.1
Master: FGT6HD3916806098, operating cluster index = 0
Slave : FGT6HD3916806070, operating cluster index = 1
196 HighAvailability
Fortinet TechnologiesInc.

