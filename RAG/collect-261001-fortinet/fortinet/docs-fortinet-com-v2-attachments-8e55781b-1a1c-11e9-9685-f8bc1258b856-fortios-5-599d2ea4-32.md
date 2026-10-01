---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-32
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [4260, 4440]
sha256: a56ae356f541b076bb707cc7caa1c99cc5618de199795591edadc076b20010df
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

FGCPHAwith802.3adaggregatedinterfaces FGCPconfigurationexamplesandtroubleshooting
set hbdev port5 50 port6 50
end
Sinceport3andport4willbeusedforanaggregatedinterface, youmustchangetheHAheartbeat
configuration.
TheFortiGate negotiatestoestablishanHAcluster.Youmaytemporarilyloseconnectivitywiththe
FortiGate astheHAclusternegotiatesandtheFGCPchangestheMACaddressoftheFortiGate
interfaces.TheMACaddressesoftheFortiGate interfaceschangetothefollowing virtualMAC
addresses:
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
4. RepeatthesestepsfortheotherFortiGate.
124 HighAvailability
Fortinet TechnologiesInc.

FGCPconfigurationexamplesandtroubleshooting FGCPHAwith802.3adaggregatedinterfaces
SettheotherFortiGate hostnameto:
config system global
set hostname FGT_ha_2
end
To connect the cluster to the network
1. Connecttheport1andport2interfacesofFGT_ha_1andFGT_ha_2toaswitchconnectedtotheInternet.
Configuretheswitchsothat theport1andport2ofFGT_ha_1makeupanaggregatedinterfaceand
port1andport2ofFGT_ha_2makeupanotheraggregatedinterface.
2. Connecttheport3andport4interfacesofFGT_ha_1andFGT_ha_2toaswitchconnectedtotheinternal
network.
Configuretheswitchsothat theport3andport4ofFGT_ha_1makeupaninterfacedandport3and
port4ofFGT_ha_2makeupanotheraggregatedinterface.
3. Connecttheport5interfacesofFGT_ha_1andFGT_ha_2together. YoucanuseacrossoverEthernetcableor
regularEthernetcablesandaswitch.
4. Connecttheport5interfacesoftheclusterunitstogether. YoucanuseacrossoverEthernetcableorregular
Ethernetcablesandaswitch.
5. Powerontheclusterunits.
Theunitsstartandnegotiate tochoosetheprimaryunitandthesubordinateunit. Thisnegotiation
occurswithnouserinterventionandnormallytakeslessthanaminute.
Whennegotiation iscompletetheclusterisreadytobeconfiguredforyournetwork.
To view cluster status
Usethefollowing stepstoviewclusterstatusfromtheCLI.
1. LogintotheCLI.
2. Enterget system status toverifytheHAstatusoftheclusterunitthat youloggedinto. Lookforthefollowing
information inthecommandoutput.
Current HA mode: a-a, master Theclusterunitsareoperatingasaclusterandyouhave
connectedtotheprimaryunit.
Current HA mode: a-a, backup Theclusterunitsareoperatingasaclusterandyouhave
connectedtoasubordinateunit.
Current HA mode: standalone TheclusterunitisnotoperatinginHAmode
3. Enterthefollowing commandtoviewthestatusofthecluster:
get system ha status
HA Health Status: OK
Model: FortiGate- XXXX
Mode: HA A-P
Group: 0
Debug: 0
Cluster Uptime: 7 days 00:30:26
.
HighAvailability
Fortinet TechnologiesInc.
125

FGCPHAwith802.3adaggregatedinterfaces FGCPconfigurationexamplesandtroubleshooting
.
.
Youcanusethiscommandtoconfirmthat theclusterishealthyandoperatingnormally,some
information abouttheclusterconfiguration, andinformation abouthowlongtheclusterhasbeen
operating.Information notshowninthisexampleincludeshowtheprimaryunitwasselected,
configurationsynchronizationstatus, usagestatsforeachclusterunit, heartbeatstatus, andthe
relativeprioritiesoftheclusterunits.
To troubleshoot the cluster configuration
SeeTroubleshootingHAclustersonpage138totroubleshootthecluster.
To add basic configuration settings and the aggregate interfaces
Usethefollowing stepstoaddafewbasicconfigurationsettingsandtheaggregateinterfaces.
1. Addapasswordfortheadminadministrative account.
config system admin
edit admin
set password <psswrd>
end
2. Temporarilydeletethedefault route.
Youcannotaddaninterfacetoanaggregateinterfaceifanysettings(suchasthedefault route)are
configuredforit. Inthisexampletheindexofthedefault routeis1.
config router static
delete 1
end
3. Addtheaggregateinterfaces:
config system interface
edit Port1_Port2
set type aggregate
set lacp-ha-slave disable
set member port1 port2
set ip 172.20.120.141/24
set vdom root
next
edit Port3_Port4
set type aggregate
set lacp-ha-slave disable
set member port3 port4
set ip 10.11.101.100/24
set vdom root
end
ThevirtualMACaddressesoftheFortiGate interfaceschangetothefollowing. Notethat port1and
port2bothhavetheport1virtualMACaddressandport3andport4bothhavetheport3virtualMAC
address:
l port1interfacevirtualMAC:00-09-0f-09-00-00
l port10interfacevirtualMAC:00-09-0f-09-00-01
l port11interfacevirtualMAC:00-09-0f-09-00-02
126 HighAvailability
Fortinet TechnologiesInc.

FGCPconfigurationexamplesandtroubleshooting ExampleHAandredundantinterfaces
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
4. Addthedefault route.
config router static
edit 1
set dst 0.0.0.0 0.0.0.0
set gateway 172.20.120.2
set device Port1_Port2
end
To configure HA port monitoring for the aggregate interfaces
1. ConfigureHAportmonitoring fortheaggregateinterfaces.
config system ha
set monitor Port1_Port2 Port3_Port4
end
Example HA and redundantinterfaces
OnFortiGate modelsthat supportityoucancombinetwoormoreinterfacesintoasingleredundantinterface. A
redundantinterfaceconsistsoftwoormorephysicalinterfaces.Traffic isprocessedbythefirstphysicalinterface
intheredundantinterface. If that physicalinterfacefails, traffic failsovertothenextphysicalinterface.
Redundantinterfacesdon’thavethebenefit ofimprovedperformancethat aggregateinterfacescanhave,but
theydoprovidefailoverifaphysicalinterfacefailsorisdisconnected.
HighAvailability
Fortinet TechnologiesInc.
127

