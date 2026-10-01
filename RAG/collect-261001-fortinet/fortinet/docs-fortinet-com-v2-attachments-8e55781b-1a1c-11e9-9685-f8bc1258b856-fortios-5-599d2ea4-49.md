---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-49
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["license", "licenses"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [6606, 6731]
sha256: 773ed200e4ae66a35cdbcaaf43fd8073ba462862961ab0c56267241850c24ea9
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

ClustersandSNMP Operatingclustersandvirtualclusters
theFortiGate interfaceis10.10.10.1. TheHAstatustableMIBfield isfgHaStatsCpuUsageandtheOIDforthis
MIBfield is1.3.6.1.4.1.12356.101.13.2.1.1.3.1. ThefirstcommandusestheMIBfield nameandthesecond
usestheOIDforthistable:
snmpget -v2c -c public-FG50012205400050 10.10.10.1 fgHaStatsCpuUsage
snmpget -v2c -c public-FG50012205400050 10.10.10.1 1.3.6.1.4.1.12356.101.13.2.1.1.3.1
FortiGate SNMPrecognizesthecommunity namewithsyntax<community_ name>-<fgt_serial>.When
theprimaryunitreceivesanSNMPgetrequestthat includesthecommunity namefollowedbyserialnumber,the
FGCPextractstheserialnumberfromtherequest.ThentheprimaryunitredirectstheSNMPgetrequesttothe
clusterunitwiththat serialnumber.If theserialnumbermatchestheserialnumberoftheprimaryunit, theSNMP
getisprocessedbytheprimaryunit.
Getting serial numbers of cluster units
Thefollowing SNMPgetcommandsusetheMIBfield namefgHaStatsSerial.<index>togettheserialnumberof
eachclusterunit. Where<index>istheclusterunit’sclusterindexand1istheclusterindexoftheprimaryunit, 2
istheclusterindexofthefirstsubordinateunit, and3istheclusterindexofthesecondsubordinateunit.
TheOIDforthisMIBfield is1.3.6.1.4.1.12356.101.13.2.1.1.2.1 .Thecommunity nameispublic.
TheIPaddressoftheFortiGate interfaceis10.10.10.1.
ThefirstcommandusestheMIBfield nameandthesecondusestheOIDforthistableandgetstheserial
numberoftheprimaryunit:
snmpget -v2c -c public 10.10.10.1 fgHaStatsSerial.1
snmpget -v2c -c public 10.10.10.1 1.3.6.1.4.1.12356.101.13.2.1.1.2.1
ThesecondcommandusestheMIBfield nameandthesecondusestheOIDforthistableandgetstheserial
numberofthefirstsubordinateunit:
snmpget -v2c -c public 10.10.10.1 fgHaStatsSerial.2
snmpget -v2c -c public 10.10.10.1 1.3.6.1.4.1.12356.101.13.2.2.2
SNMP get command syntax - reserved management interface enabled
Togetconfigurationandstatusinformation foranyclusterunitwhereyouhaveenabledtheHAreserved
management interfacefeatureandassignedIPaddressestothemanagement interfaceofeachclusterunit, an
SNMPmanagerwouldusethefollowing getcommandsyntax:
snmpget -v2c -c <community_name> <mgmt_address_ipv4> {<OID> | <MIB_field>}
where:
<community_ name> isanSNMPcommunity nameaddedtotheFortiGate configuration. Youcanaddmore
thanonecommunity namestoaFortiGate SNMPconfiguration. Themostcommonlyusedcommunity nameis
public.
<mgmt_address_ipv4> istheIPaddressoftheFortiGate HAreservedmanagement interfacethat theSNMP
managerconnectsto.
{<OID> | <MIB_field>} istheobjectidentifier (OID)fortheMIBfield ortheMIBfield nameitself. Tofind
OIDsandMIBfield namesseeyourFortiGate’s onlinehelp.
184 HighAvailability
Fortinet TechnologiesInc.

Operatingclustersandvirtualclusters AddingFortiClient licensestoacluster
Adding FortiClientlicenses to a cluster
EachFortiGate inaclustermusthaveitsownFortiClient license.ContactyourresellertopurchaseFortiClient
licensesforalloftheFortiGates inyourcluster.
WhenyoureceivethelicensekeysyoucanvisittheFortinet SupportwebsiteandaddtheFortiClient licensekeys
toeachFortiGate. Then,aslongastheclustercanconnecttotheInternet eachclusterunitreceivesits
FortiClient licensekeyfromtheFortiGuardnetwork.
Adding FortiClient licenses to cluster units with a reserved management interface
Youcanalsousethefollowing stepstomanuallyaddlicensekeystoyourclusterunitsfromtheGUIorCLI. Your
clustermustbeconnectedtotheInternet andyoumusthaveconfiguredareservedmanagement interfacefor
eachclusterunit.
1. LogintotheGUIofeachclusterunitusingitsreservedmanagement interfaceIPaddress.
2. Gotothe License Informationdashboardwidgetandbeside FortiClientselect Enter License.
3. Enterthelicensekeyandselect OK.
4. Confirmthat thelicensehasbeeninstalledandthecorrectnumberofFortiClients arelicensed.
5. Repeatforalloftheclusterunits.
Youcanalsousethereservedmanagement IPaddresstologintoeachclusterunitCLIandusefollowing
commandtoaddthelicensekey:
execute FortiClient-NAC update-registration-license <license-key>
YoucanconnecttotheCLIsofeachclusterunitusingtheirreservedmanagement IPaddress.
Adding FortiClient licenses to cluster units with no reserved management interface
If youhavenotsetupreservedmanagement IPaddressesforyourclusterunits, youcanstill addFortiClient
licensekeystoeachclusterunit. Youmustlogintotheprimaryunitandthenusetheexecute ha manage
commandtoconnecttoeachclusterunitCLI. Forexample,usethefollowing stepstoaddaFortiClient license
keyaclusterofthreeFortiGates:
1. LogintotheprimaryunitCLIandenterthefollowing commandtoconfirmtheserialnumberoftheprimaryunit:
get system status
2. AddtheFortiClient licensekeyforthat serialnumbertotheprimaryunit:
execute FortiClient-NAC update-registration-license <license-key>
YoucanalsousetheGUItoaddthelicensekeytotheprimaryunit.
3. Enterthefollowing commandtologintothefirstsubordinateunit:
execute ha manage 1
4. Enterthefollowing commandtoconfirmtheserialnumberoftheclusterunitthat youhaveloggedinto:
get system status
5. AddtheFortiClient licensekeyforthat serialnumbertotheclusterunit:
execute FortiClient-NAC update-registration-license <license-key>
6. Enterthefollowing commandtologintothesecondsubordinateunit:
HighAvailability
Fortinet TechnologiesInc.
185

Clustermemberslist Operatingclustersandvirtualclusters
execute ha manage 2
7. Enterthefollowing commandtoconfirmtheserialnumberoftheclusterunitthat youhaveloggedinto:
get system status
8. AddtheFortiClient licensekeyforthat serialnumbertotheclusterunit:
execute FortiClient-NAC update-registration-license <license-key>
Viewing FortiClient license status and active FortiClient users for each cluster unit
ToviewFortiClient licensestatusandFortiClient information foreachclusterunityoumustlogintoeachcluster
unit’sGUIorCLI. Youcandothisbyconnectingtoeachclusterunit’sreservedmanagement interfaceiftheyare
configured.If youhavenotconfiguredreservedmanagement interfacesyoucanusetheexecutehamanage
commandtologintoeachclusterunitCLI.
FromtheGUI, viewFortiClient LicensestatusfromtheLicenseInformation dashboardwidgetandselect Details
todisplaythelistofactiveFortiClient usersconnectingthroughthat clusterunit. Youcanalsoseeactive
FortiClient usersbygoingto User & Device > Monitor > FortiClient.
FromtheCLIyoucanusetheexecute FortiClient {list | info} commandtodisplayFortiClient
licensestatusandactiveFortiClient users.
Forexample,usethefollowing commandtodisplaytheFortiClient licensestatusoftheclusterunitthat youare
loggedinto:
execute forticlient info
Maximum FortiClient connections: unlimited.
Licensed connections: 114
NAC: 114
WANOPT: 0
Test: 0
Other connections:
IPsec: 0
SSLVPN: 0
Usethefollowing commandtodisplaythelistofactiveFortiClient usersconnectingthroughtheclusterunit. The
output showsthetime theconnectionwasestablished,thetypeofFortiClient connection,thenameofthe
device,theusernameofthepersonconnecting,theFortiClient ID, thehostoperatingsystem,andthesourceIP
addressofthesession.
execute forticlient list
TIMESTAMP TYPE CONNECT-NAME USER CLIENT-ID HOST-OS SRC-IP
20141017 09:13:33 NAC Gordon-PC Gordon 11F76E902611484A942E31439E428C5CMicrosoft
Windows 7 , 64-bit Service Pack 1 (build 7601) 172.20.120.10
20141017 09:11:55 NAC Gordon-PC 11F76E902611484A942E31439E428C5CMicrosoft Windows 7 ,
64-bit Service Pack 1 (build 7601) 172.20.120.10
20141017 07:27:11 NAC Desktop11 Richie 9451C0B8EE3740AEB7019E920BB3761BMicrosoft
Windows 7, 64-bit Service Pack 1 (build 7601) 172.20.120.20
Cluster members list
Todisplaytheclustermemberslist, logintoanoperatingclusterandgoto System > HA.
186 HighAvailability
Fortinet TechnologiesInc.

