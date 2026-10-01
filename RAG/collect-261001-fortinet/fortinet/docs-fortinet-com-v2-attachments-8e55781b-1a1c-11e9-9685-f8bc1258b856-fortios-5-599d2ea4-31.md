---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-31
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [4088, 4259]
sha256: 104d288098c01930cbaa9f1e76dd71bfd54ad42aefced9c25a5f45f44eb2b956
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

FGCPHAwith802.3adaggregatedinterfaces FGCPconfigurationexamplesandtroubleshooting
l port11interfacevirtualMAC:00-09-0f-09-00-02
l port12interfacevirtualMAC:00-09-0f-09-00-03
l port13interfacevirtualMAC:00-09-0f-09-00-04
l port14interfacevirtualMAC:00-09-0f-09-00-05
l port15interfacevirtualMAC:00-09-0f-09-00-06
l port16interfacevirtualMAC:00-09-0f-09-00-07
l port17interfacevirtualMAC:00-09-0f-09-00-08
l port18interfacevirtualMAC:00-09-0f-09-00-09
l port19interfacevirtualMAC:00-09-0f-09-00-0a
l port2interfacevirtualMAC:00-09-0f-09-00-0b
l port20interfacevirtualMAC:00-09-0f-09-00-0c
l port3interfacevirtualMAC:00-09-0f-09-00-0d
l port4interfacevirtualMAC:00-09-0f-09-00-0e
l port5interfacevirtualMAC:00-09-0f-09-00-0f
l port6interfacevirtualMAC:00-09-0f-09-00-10
l port7interfacevirtualMAC:00-09-0f-09-00-11
l port8interfacevirtualMAC:00-09-0f-09-00-12
l port9interfacevirtualMAC:00-09-0f-09-00-13
Toreconnectsooner,youcanupdatetheARPtableofyourmanagement PCbydeletingtheARP
tableentryfortheFortiGate (orjustdeletingallarptableentries).Youmaybeabletodeletethearp
tableofyourmanagement PCfromacommandpromptusingacommandsimilartoarp -d.
Youcanusetheget hardware nic (ordiagnose hardware deviceinfo nic)CLI
commandtoviewthevirtualMACaddressofanyFortiGate interface. Forexample,usethefollowing
commandtoviewtheport1interfacevirtualMACaddress(Current_HWaddr)andtheport1
permanentMACaddress(Permanent_ HWaddr):
get hardware nic port1
.
.
.
MAC: 00:09:0f:09:00:00
Permanent_HWaddr: 02:09:0f:78:18:c9
.
.
.
7. Poweroff thefirstFortiGate.
8. RepeatthesestepsforthesecondFortiGate.
SetthesecondFortiGate hostnameto:
New Name FGT_ha_2
To connect the cluster to the network
1. Connecttheport1andport2interfacesofFGT_ha_1andFGT_ha_2toaswitchconnectedtotheInternet.
Configuretheswitchsothat theport1andport2ofFGT_ha_1makeupanaggregatedinterfaceand
port1andport2ofFGT_ha_2makeupasecondaggregatedinterface.
120 HighAvailability
Fortinet TechnologiesInc.

FGCPconfigurationexamplesandtroubleshooting FGCPHAwith802.3adaggregatedinterfaces
2. Connecttheport3andport4interfacesofFGT_ha_1andFGT_ha_2toaswitchconnectedtotheinternal
network.
Configuretheswitchsothat theport3andport4ofFGT_ha_1makeupanaggregatedinterfaceand
port3andport4ofFGT_ha_2makeupanotheraggregatedinterface.
3. Connecttheport5interfacesofFGT_ha_1andFGT_ha_2together. YoucanuseacrossoverEthernetcableor
regularEthernetcablesandaswitch.
4. Connecttheport5interfacesoftheclusterunitstogether. YoucanuseacrossoverEthernetcableorregular
Ethernetcablesandaswitch.
5. Powerontheclusterunits.
Theunitsnegotiate tochoosetheprimaryunitandthesubordinateunit. Thisnegotiation occurswith
nouserinterventionandnormallytakeslessthanaminute.
Whennegotiation iscomplete, theclusterisreadytobeconfiguredforyournetwork.
To view cluster status
Usethefollowing stepstoviewtheclusterdashboardandclustermemberslisttoconfirmthat theclusterunits
areoperatingasacluster.
1. Viewthesystemdashboard.
TheHAStatusdashboardwidgetdisplayshowlongtheclusterhasbeenoperating(Uptime)andthe
time sincethelastfailoveroccurred(StateChanged)YoucanhoverovertheStateChangedtime to
seetheeventthat causedthestatechange.YoucanalsoclickontheHAStatusdashboardwidgetto
configureHAsettingsortogetalistingofthemostrecentHAeventsrecordedbythecluster.
2. Goto System > HA toviewtheclustermemberslist.
Thelistshowsbothclusterunits, theirhostnames,theirrolesinthecluster,andtheirdevicepriorities.
Youcanusethislisttoconfirmthat theclusterisoperatingnormally.Forexample,ifthelistshows
onlyoneclusterunitthentheotherunithasleft theclusterforsomereason.
To troubleshoot the cluster configuration
SeeTroubleshootingHAclustersonpage138totroubleshootthecluster.
To add basic configuration settings and the aggregate interfaces
Usethefollowing stepstoaddafewbasicconfigurationsettings.
1. LogintotheclusterGUI.
2. Goto System > Administrators.
3. Edit adminandselect Change Password.
4. Enterandconfirmanewpassword.
5. Select OK.
6. Goto Network > Static Routesandtemporarilydeletethedefault route.
Youcannotaddaninterfacetoaaggregatedinterfaceifanysettings(suchasthedefault route)are
configuredforit.
7. Goto Network > Interfacesandselect Create New > Interfacetoaddtheaggregateinterfacetoconnecttothe
Internet.
8. Set Typeto 802.3ad AggregateandconfiguretheaggregateinterfacetobeconnectedtotheInternet:
HighAvailability
Fortinet TechnologiesInc.
121

FGCPHAwith802.3adaggregatedinterfaces FGCPconfigurationexamplesandtroubleshooting
Name Port1_Port2
Interface Members port1,port2
IP/Network Mask 172.20.120.141/24
9. Select OK.
10. Select Create New > Interfacetoaddtheaggregateinterfacetoconnecttotheinternalnetwork.
11. Set Typeto 802.3ad AggregateandconfiguretheaggregateinterfacetobeconnectedtotheInternet:
Name Port3_Port4
Interface Members port3,port4
IP/Netmask 10.11.101.100/24
Administrative Access HTTPS,PING,SSH
12. Select OK.
ThevirtualMACaddressesoftheFortiGate interfaceschangetothefollowing. Notethat port1and
port2bothhavetheport1virtualMACaddressandport3andport4bothhavetheport3virtualMAC
address:
l port1interfacevirtualMAC:00-09-0f-09-00-00
l port10interfacevirtualMAC:00-09-0f-09-00-01
l port11interfacevirtualMAC:00-09-0f-09-00-02
l port12interfacevirtualMAC:00-09-0f-09-00-03
l port13interfacevirtualMAC:00-09-0f-09-00-04
l port14interfacevirtualMAC:00-09-0f-09-00-05
l port15interfacevirtualMAC:00-09-0f-09-00-06
l port16interfacevirtualMAC:00-09-0f-09-00-07
l port17interfacevirtualMAC:00-09-0f-09-00-08
l port18interfacevirtualMAC:00-09-0f-09-00-09
l port19interfacevirtualMAC:00-09-0f-09-00-0a
l port2interfacevirtualMAC:00-09-0f-09-00-00 (sameasport1)
l port20interfacevirtualMAC:00-09-0f-09-00-0c
l port3interfacevirtualMAC:00-09-0f-09-00-0d
l port4interfacevirtualMAC:00-09-0f-09-00-0d (sameasport3)
l port5interfacevirtualMAC:00-09-0f-09-00-0f
l port6interfacevirtualMAC:00-09-0f-09-00-10
l port7interfacevirtualMAC:00-09-0f-09-00-11
l port8interfacevirtualMAC:00-09-0f-09-00-12
l port9interfacevirtualMAC:00-09-0f-09-00-13
13. ConnecttotheCLIandenterthefollowing commandtodisablesendingLACPpacketsfromthesubordinateunit:
config system interface
edit Port1_Port2
set lacp-ha-slave disable
122 HighAvailability
Fortinet TechnologiesInc.

FGCPconfigurationexamplesandtroubleshooting FGCPHAwith802.3adaggregatedinterfaces
next
edit Port3_Port4
set lacp-ha-slave disable
end
14. Goto Network > Static Routes.
15. Addthedefault route.
Destination IP/Mask 0.0.0.0/0.0.0.0
Gateway 172.20.120.2
Device Port1_Port2
Distance 10
16. Select OK.
To configure HA port monitoring for the aggregate interfaces
1. Goto System > HA.
2. Intheclustermemberslist, edittheprimaryunit.
3. Configurethefollowing portmonitoring fortheaggregateinterfaces:
Port Monitor
Port1_Port2 Select
Port3_Port4 Select
4. Select OK.
Configuring active-passive HA cluster that includes aggregate interfaces - CLI
TheseproceduresassumeyouarestartingwithtwoFortiGates withfactorydefault settings.
To configure the FortiGates for HA operation
1. RegisterandapplylicensestotheFortiGate. Thisincludes FortiCloud activationand FortiClientlicensing,and
enteringalicensekeyifyoupurchasedmorethan10 Virtual Domains(VDOMS).AlloftheFortiGates ina
clustermusthavethesameleveloflicensing.
2. Install anythird-partycertificatesontheFortiGate.
3. ChangethehostnameforthisFortiGate:
config system global
set hostname FGT_ha_1
end
4. ConfigureHAsettings.
config system ha
set mode a-p
set group-name example5.com
set password HA_pass_5
HighAvailability
Fortinet TechnologiesInc.
123

