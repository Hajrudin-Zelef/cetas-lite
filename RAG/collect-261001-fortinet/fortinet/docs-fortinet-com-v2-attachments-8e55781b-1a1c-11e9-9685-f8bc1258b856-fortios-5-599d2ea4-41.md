---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-41
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [5512, 5664]
sha256: a33cd13d48c78955a582a09fb23b22d1e86438eb0ec13e30083177c91493aef8
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

FullmeshHA Examplefull meshHAconfiguration
Full mesh HA configuration
ThetwoFortiGates (FGT_ha_1andFGT_ha_2)canbeoperatinginNATortransparentmode. Asidefromthe
standardHAsettings, theFortiGate configurationincludesthefollowing:
l Theport5andport6interfacesconfiguredasheartbeatinterfaces.Afull meshHAconfigurationalsoincludes
redundantHAheartbeatinterfaces.
l Theport1andport2interfacesaddedtoaredundantinterface. Port1istheactivephysicalinterfaceinthis
redundantinterface. Tomaketheport1interfacetheactivephysicalinterfaceitshouldappearabovetheport2
interfaceintheredundantinterfaceconfiguration.
l Theport3andport4interfacesaddedtoaredundantinterface. Port3istheactivephysicalinterfaceinthis
redundantinterface. Tomaketheport3interfacetheactivephysicalinterfaceitshouldappearabovetheport4
interfaceintheredundantinterfaceconfiguration.
Full mesh switch configuration
Thefollowing redundantswitchconfigurationisrequired:
l Tworedundantswitches(Sw3andSw4)connectedtotheinternalnetwork.Establishan802.1Q(Dot1Q)or
interswitch-link(ISL)connectionbetweenthem.
l Tworedundantswitches(Sw1andSw2)connectedtotheInternet. Establishan802.1Q(Dot1Q)orinterswitch-link
(ISL)connectionbetweenthem.
Full mesh network connections
Makethefollowing physicalnetworkconnectionsforFGT_ha_1:
l Port1toSw1(active)
l Port2toSw2(inactive)
l Port3toSw3(active)
l Port4toSw4(inactive)
Makethefollowing physicalnetworkconnectionsforFGT_ha_2:
l Port1toSw2(active)
l Port2toSw1(inactive)
l Port3toSw4(active)
l Port4toSw3(inactive)
How packets travel from the internal network through the full mesh cluster and to the
Internet
If theclusterisoperatinginactive-passivemodeandFGT_ha_2istheprimaryunit, allpacketstakethefollowing
pathfromtheinternalnetworktotheinternet:
1. FromtheinternalnetworktoSw4.Sw4istheactiveconnectiontoFGT_ha_2;whichistheprimaryunit. The
primaryunitreceivesallpackets.
2. FromSw4totheFGT_ha_2port3interface. ActiveconnectionbetweenSw4andFGT_ha_2.Port3istheactive
memberoftheredundantinterface.
HighAvailability
Fortinet TechnologiesInc.
157

Examplefull meshHAconfiguration FullmeshHA
3. FromFGT_ha_2port3toFGT_ha_2port1.ActiveconnectionbetweenFGT_ha_2andSw2.Port1istheactive
memberoftheredundantinterface.
4. FromSw2totheexternalrouterandtheInternet.
Configuring full-mesh HA - GUI
EachclusterunitmusthavethesameHAconfiguration.
To configure the FortiGates for HA operation
1. RegisterandapplylicensestotheFortiGate.
2. Onthe System Informationdashboardwidget, beside Host Nameselect Change.
3. EnteranewHostNameforthisFortiGate.
New Name FGT_ha_1
4. GotoSystem > HA andchangethefollowing settings.
Mode Active-Active
Group Name Rexample1.com
Password RHA_pass_1
Heartbeat Interface
Enable Priority
port5 Select 50
port6 Select 50
158 HighAvailability
Fortinet TechnologiesInc.

FullmeshHA Examplefull meshHAconfiguration
5. Select OK.
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
6. Poweroff thefirstFortiGate.
7. RepeatthesestepsforthesecondFortiGate.
SetthesecondFortiGate hostnameto:
HighAvailability
Fortinet TechnologiesInc.
159

Examplefull meshHAconfiguration FullmeshHA
New Name FGT_ha_2
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
Usethefollowing stepstoviewtheclusterdashboardandclustermemberslisttoconfirmthat theclusterunits
areoperatingasacluster.
1. Viewthesystemdashboard.
TheSystemInformation dashboardwidgetshowsthe Cluster Name(Rexample1.com)andthehost
namesandserialnumbersofthe Cluster Members.TheUnitOperationwidgetshowsmultiple
clusterunits.
2. Goto System > HA toviewtheclustermemberslist.
Thelistshowstwoclusterunits, theirhostnames,theirrolesinthecluster,andtheirpriorities.You
canusethislisttoconfirmthat theclusterisoperatingnormally.
To troubleshoot the cluster configuration
If theclustermemberslistandthedashboarddoesnotdisplayinformation forbothclusterunitstheFortiGates
arenotfunctioning asacluster.SeeExamplefull meshHAconfigurationonpage156totroubleshootthecluster.
To add basic configuration settings and the redundant interfaces
Usethefollowing stepstoaddafewbasicconfigurationsettings.
1. LogintotheclusterGUI.
2. Goto System > Administrators.
3. Edit adminandselect Change Password.
160 HighAvailability
Fortinet TechnologiesInc.

