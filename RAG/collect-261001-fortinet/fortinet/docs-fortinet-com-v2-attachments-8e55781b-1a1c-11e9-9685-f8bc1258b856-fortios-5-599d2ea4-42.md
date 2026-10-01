---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-42
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [5665, 5839]
sha256: e46b6793466fa7329a7b4ead79967875aec4ac5d627a1d13de99dc0575b77d34
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

FullmeshHA Examplefull meshHAconfiguration
4. Enterandconfirmanewpassword.
5. Select OK.
6. Goto Network > Static Routesandtemporarilydeletethedefault route.
Youcannotaddaninterfacetoaredundantinterfaceifanysettings(suchasthedefault route)are
configuredforit.
7. Goto Network > Interfacesandselect Create New > Interfaceandconfiguretheredundantinterfaceto
connecttotheInternet.
Name Port1_Port2
Type Redundant
Physical Interface Members
Selected Interfaces port1,port2
IP/Netmask 172.20.120.141/24
8. Select OK.
9. Select Create Newandconfiguretheredundantinterfacetoconnecttotheinternalnetwork.
Name Port3_Port4
Type Redundant
Physical Interface Members
Selected Interfaces port3,port4
IP/Netmask 10.11.101.100/24
Administrative Access HTTPS,PING,SSH
10. Select OK.
ThevirtualMACaddressesoftheFortiGate interfaceschangetothefollowing. Noticethat port1and
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
HighAvailability
Fortinet TechnologiesInc.
161

Examplefull meshHAconfiguration FullmeshHA
l port2interfacevirtualMAC:00-09-0f-09-00-00 (sameasport1)
l port20interfacevirtualMAC:00-09-0f-09-00-0c
l port3interfacevirtualMAC:00-09-0f-09-00-0d
l port4interfacevirtualMAC:00-09-0f-09-00-0d (sameasport3)
l port5interfacevirtualMAC:00-09-0f-09-00-0f
l port6interfacevirtualMAC:00-09-0f-09-00-10
l port7interfacevirtualMAC:00-09-0f-09-00-11
l port8interfacevirtualMAC:00-09-0f-09-00-12
l port9interfacevirtualMAC:00-09-0f-09-00-13
11. Goto Router > Static > Static Routes.
12. Addthedefault route.
Destination IP/Mask 0.0.0.0/0.0.0.0
Gateway 172.20.120.2
Device Port1_Port2
Distance 10
13. Select OK.
To configure HA port monitoring for the redundant interfaces
1. Goto System > HA.
2. Intheclustermemberslist, edittheprimaryunit.
3. Enable interface monitoringthe Port1_Port2andthe Port3_Port4interfaces
4. Select OK.
Configuring full mesh HA - CLI
EachclustermusthavethesameHAconfiguration. Usethefollowing proceduretoconfiguretheFortiGates for
HAoperation.
To configure the FortiGates for HA operation
1. RegisterandapplylicensestotheFortiGate.
2. EnteranewHostNameforthisFortiGate.
config system global
set hostname FGT_ha_1
end
3. ConfigureHAsettings.
config system ha
set mode a-a
set group-name Rexample1.com
set password RHA_pass_1
set hbdev port5 50 port6 50
end
162 HighAvailability
Fortinet TechnologiesInc.

FullmeshHA Examplefull meshHAconfiguration
TheFortiGate negotiatestoestablishanHAcluster.WhenyouselectOKyoumaytemporarilylose
connectivitywiththeFortiGate astheHAclusternegotiatesandtheFGCPchangestheMACaddress
oftheFortiGate interfaces.TheMACaddressesoftheFortiGate interfaceschangetothefollowing
virtualMACaddresses:
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
l port2interfacevirtualMAC:00-09-0f-09-00-0b
l port20interfacevirtualMAC:00-09-0f-09-00-0c
l port3interfacevirtualMAC:00-09-0f-09-00-0d
l port4interfacevirtualMAC:00-09-0f-09-00-0e
l port5interfacevirtualMAC:00-09-0f-09-00-0f
l port6interfacevirtualMAC:00-09-0f-09-00-10
l port7interfacevirtualMAC:00-09-0f-09-00-11
l port8interfacevirtualMAC:00-09-0f-09-00-12
l port9interfacevirtualMAC:00-09-0f-09-00-13
Tobeabletoreconnectsooner,youcanupdatetheARPtableofyourmanagement PCbydeleting
theARPtableentryfortheFortiGate (orjustdeletingallarptableentries).Youmaybeabletodelete
thearptableofyourmanagement PCfromacommandpromptusingacommandsimilartoarp -d.
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
4. Poweroff thefirstFortiGate.
5. RepeatthesestepsforthesecondFortiGate.
SettheotherFortiGate hostnameto:
config system global
set hostname FGT_ha_2
HighAvailability
Fortinet TechnologiesInc.
163

Examplefull meshHAconfiguration FullmeshHA
end
To connect the cluster to your network
1. Makethefollowing physicalnetworkconnectionsforFGT_ha_1:
l Port1toSw1(active)
l Port2toSw2(inactive)
l Port3toSw3(active)
l Port4toSw4(inactive)
2. Makethefollowing physicalnetworkconnectionsforFGT_ha_2:
l Port1toSw2(active)
l Port2toSw1(inactive)
l Port3toSw4(active)
l Port4toSw3(inactive)
3. ConnectSw3andSw4totheinternalnetwork.
4. ConnectSw1andSw2totheexternalrouter.
5. Enable802.1Q(Dot1Q)orISLcommunication betweenSw1andSw2andbetweenSw3andSw4.
6. Powerontheclusterunits.
Theunitsstartandnegotiate tochoosetheprimaryunitandthesubordinateunit. Thisnegotiation
occurswithnouserintervention.
Whennegotiation iscompletetheclusterisreadytobeconfiguredforyournetwork.
To view cluster status
Usethefollowing stepstoviewclusterstatusfromtheCLI.
1. LogintotheCLI.
2. Enterget system status toverifytheHAstatusoftheclusterunitthat youloggedinto.
If thecommandoutput includesCurrent HA mode: a-a, master,theclusterunitsare
operatingasaclusterandyouhaveconnectedtotheprimaryunit.
If thecommandoutput includesCurrent HA mode: a-a, backup,youhaveconnectedtoa
subordinateunit.
If thecommandoutput includesCurrent HA mode: standalone theclusterunitisnot
operatinginHAmode.
3. Enterthefollowing commandtoconfirmtheHAconfigurationofthecluster:
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
164 HighAvailability
Fortinet TechnologiesInc.

