---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-25
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["licenses"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [3177, 3339]
sha256: d8f696a24557b2be1498bdb4d91b0241f6647d56c46708ef9f075206b7928e6e
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

FGCPconfigurationexamplesandtroubleshooting Active-activeHAclusterintransparentmode
set hostname FGT_ha_2
end
10. ConfigureHAsettings.
config system ha
set mode a-a
set group-name example2.com
set password HA_pass_2
end
TheFortiGate negotiatestoestablishanHAcluster.Youmaytemporarilylosenetworkconnectivity
withtheFortiGate astheHAclusternegotiatesandbecausetheFGCPchangestheMACaddressof
theFortiGate interfaces.
Toreconnectsooner,youcanupdatetheARPtableofyourmanagement PCbydeletingtheARP
tableentryfortheFortiGate (orjustdeletingallarptableentries).Youmaybeabletodeletethearp
tableofyourmanagement PCfromacommandpromptusingacommandsimilartoarp -d.
11. DisplaytheHAconfiguration(optional).
get system ha
group-id : 0
group-name : example2.com
mode : a-a
password : *
hbdev : "port3" 50 "port4" 50
session-sync-dev :
route-ttl : 10
route-wait : 0
route-hold : 10
sync-config : enable
encryption : disable
authentication : disable
hb-interval : 2
hb-lost-threshold : 20
hello-holddown : 20
arps : 5
arps-interval : 8
session-pickup : disable
update-all-session-timer: disable
session-sync-daemon-number: 1
link-failed-signal : disable
uninterruptible-upgrade: enable
ha-mgmt-status : disable
ha-eth-type : 8890
hc-eth-type : 8891
l2ep-eth-type : 8893
ha-uptime-diff-margin: 300
vcluster2 : disable
vcluster-id : 1
override : disable
priority : 128
schedule : round-robin
monitor :
pingserver-monitor-interface:
pingserver-failover-threshold: 0
pingserver-slave-force-reset: enable
pingserver-flip-timeout: 60
HighAvailability
Fortinet TechnologiesInc.
93

Active-activeHAclusterintransparentmode FGCPconfigurationexamplesandtroubleshooting
vdom : "root"
schedule : round-robin
12. Poweroff theFortiGate.
To connect the cluster to the network
1. Connecttheport1interfacesofFGT_ha_1andFGT_ha_2toaswitchconnectedtotheInternet.
2. Connecttheport2interfacesofFGT_ha_1andFGT_ha_2toaswitchconnectedtotheinternalnetwork.
3. Connecttheport3interfacesofFGT_ha_1andFGT_ha_2together. YoucanuseacrossoverEthernetcableor
regularEthernetcablesandaswitch.
4. Connecttheport4interfacesoftheclusterunitstogether. YoucanuseacrossoverEthernetcableorregular
Ethernetcablesandaswitch.
5. Powerontheclusterunits.
Theunitsstartandnegotiate tochoosetheprimaryunitandthesubordinateunit. Thisnegotiation
occurswithnouserinterventionandnormallytakeslessthanaminute.
Whennegotiation iscompletetheclusterisreadytobeconfiguredforyournetwork.
To connect to the cluster CLI and switch the cluster to transparentmode
1. Determinewhichclusterunitistheprimaryunit.
l Usethenull-modem cableandserialconnectiontore-connecttotheCLIofoneoftheclusterunits.
l Enterthecommandget system status.
l If thecommandoutput includesCurrent HA mode: a-a, master,theclusterunitsareoperatingasa
clusterandyouhaveconnectedtotheprimaryunit. ContinuewithStep2.
l If thecommandoutput includesCurrent HA mode: a-a, backup,youhaveconnectedtoasubordinate
unit. Connecttotheotherclusterunit, whichshouldbetheprimaryunitandcontinuewithStep2.
If thecommandoutput includesCurrent HA mode: standalone ,thecluster
unitisnotoperatinginHAmode.
2. Changetotransparentmode.
config system settings
set opmode transparent
set manageip 192.168.20.3/24
set gateway 192.168.20.1
end
TheclusterswitchestotransparentMode, andyouradministration sessionisdisconnected.
YoucannowconnecttotheclusterCLIusingSSHtoconnecttotheclusterinternalinterfaceusingthe
management IPaddress(192.168.20.3).
To view cluster status
Usethefollowing stepstoviewclusterstatusfromtheCLI.
1. Determinewhichclusterunitistheprimaryunit.
94 HighAvailability
Fortinet TechnologiesInc.

FGCPconfigurationexamplesandtroubleshooting FortiGate-5000active-activeHAclusterwithFortiClient licenses
l Usethenull-modem cableandserialconnectiontore-connecttotheCLIofoneoftheclusterunits.
l Enterthecommandget system status.
l If thecommandoutput includesCurrent HA mode: a-a, master,theclusterunitsareoperatingasa
clusterandyouhaveconnectedtotheprimaryunit. ContinuewithStep"Active-activeHAclusterintransparent
mode"onpage84.
l If thecommandoutput includesCurrent HA mode: a-a, backup,youhaveconnectedtoasubordinate
unit. Connectthenull-modem cabletotheotherclusterunit, whichshouldbetheprimaryunitandcontinue
withStep2.
If thecommandoutput includesCurrent HA mode: standalone ,thecluster
unitisnotoperatinginHAmodeandyoushouldreviewyourHAconfiguration.
2. Enterthefollowing commandtoconfirmtheHAconfigurationofthecluster:
get system ha status
HA Health Status: OK
Model: FortiGate- XXXX
Mode: HA A-P
Group: 0
Debug: 0
Cluster Uptime: 7 days 00:30:26
.
.
.
Youcanusethiscommandtoconfirmthat theclusterishealthyandoperatingnormally,some
information abouttheclusterconfiguration, andinformation abouthowlongtheclusterhasbeen
operating.Information notshowninthisexampleincludeshowtheprimaryunitwasselected,
configurationsynchronizationstatus, usagestatsforeachclusterunit, heartbeatstatus, andthe
relativeprioritiesoftheclusterunits.
To troubleshoot the cluster configuration
If theclustermemberslistandthedashboarddonotdisplayinformation forbothclusterunitstheFortiGates are
notfunctioning asacluster.SeeTroubleshootingHAclustersonpage138totroubleshootthecluster.
To add a password for the admin administrative account
1. Addapasswordfortheadminadministrative account.
config system admin
edit admin
set password <psswrd>
end
FortiGate-5000 active-active HA cluster with FortiClientlicenses
ThissectiondescribeshowtoconfigureanHAclusterofthreeFortiGate-5001Dunitsthat connectaninternal
networktotheInternet. TheFortiGate-5001DunitseachhaveaFortiClient licenseinstalledonthem tosupport
HighAvailability
Fortinet TechnologiesInc.
95

FortiGate-5000active-activeHAclusterwithFortiClient licenses FGCPconfigurationexamplesandtroubleshooting
FortiClient profiles.
Normallyitisrecommendedthat youaddFortiClient licensestotheFortiGates beforesetting upthecluster.This
example;however,describeshowtoapplyFortiClient licensestotheFortiGates inanoperatingcluster.
Example network topology
Thefollowing diagramshowsanHAclusterconsistingofthreeFortiGate-5001D clusterunits(hostnamesslot-3,
slot-4,andslot-5)installedinaFortiGate-5000serieschassiswithtwoFortiController-5003Bunitsforheartbeat
communication betweentheclusterunits. TheclusterappliessecurityfeaturesincludingFortiClient profilesto
datatraffic passingthroughit.
TheclusterismanagedfromtheinternalnetworkusingtheFortiGate-5001Dmgmt1 interfacesconfiguredasHA
reservedmanagement interfaces.Usingthesereservedmanagement interfacestheoverallclustercanbe
managedandclusterunitscanbemanagedindividually.Individualmanagement accesstoeachclusterunit
makessomeoperations,suchasinstalling FortiClient licenses,easierandalsoallowsyoutoviewstatusofeach
clusterunit.
Thereservedmanagement interfaceofeachclusterunithasadifferent IPaddressandretainsitsownMAC
address.Theclusterdoesnotchangethereservedmanagement interfaceMACaddress.
Example network topology
Bydefault base1andbase2areusedforheartbeatcommunication betweentheFortiGates. Tousethebase1
andbase2interfacesfortheHAheartbeat, theexampledescribeshowtodisplaythebackplaneinterfacesonthe
96 HighAvailability
Fortinet TechnologiesInc.

