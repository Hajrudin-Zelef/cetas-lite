---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-48
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [6489, 6605]
sha256: b8dd92851031622be91a52819dff6f886cdb1d538950d118bb6d5ac8b532dffb
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

Operatingclustersandvirtualclusters ClustersandSNMP
Formatting cluster unit hard disks (log disks)
If youneedtoformat theharddisk(alsocalledlogdiskordiskstorage)ofoneormoreclusterunitsyoushould
disconnecttheunitfromtheclusterandusetheexecute formatlogdisk commandtoformat thecluster
unitharddiskthenaddtheunitbacktothecluster.
Forinformation abouthowtoremoveaunitfromaclusterandadditback,seeDisconnectingaclusterunitfroma
clusteronpage200andAddingadisconnectedFortiGate backtoitsclusteronpage201.
Onceyouaddtheclusterunitwiththeformatted logdiskbacktotheclusteryoushouldmakeittheprimaryunit
beforeremovingotherunitsfromtheclustertoformat theirlogdisksandthenaddthem backtothecluster.
Clusters and SNMP
YoucanuseSNMPtomanageaclusterbyconfiguringaclusterinterfaceforSNMPadministrative access.Using
anSNMPmanageryoucangetclusterconfigurationandstatusinformation andreceivetraps.
YouconfigureSNMPforaclusterinthesamewayasconfiguringSNMPforastandaloneFortiGate. SNMP
configurationchangesmadetotheclusteraresharedbyallclusterunits.
EachclusterunitsendsitsowntrapsandSNMPmanagersystemscanuseSNMPgetcommandstoqueryeach
clusterunitseparately.TosetSNMPgetqueriestoeachclusterunityoumustcreateaspecialgetcommandthat
includestheserialnumberoftheclusterunit.
AlternativelyyoucanusetheHAreservedmanagement interfacefeaturetogiveeachclusterunitadifferent
management IPaddress.ThenyoucancreateanSNMPgetcommandforeachclusterunitthat justincludesthe
management IPaddressanddoesnothavetoincludetheserialnumber.
SNMP get command syntax for the primary unit
Normally,togetconfigurationandstatusinformation forastandaloneFortiGate orforaprimaryunit, anSNMP
managerwoulduseanSNMPgetcommandstogettheinformation inaMIBfield. TheSNMPgetcommand
syntaxwouldbesimilartothefollowing:
snmpget -v2c -c <community_name> <address_ipv4> {<OID> | <MIB_field>}
where:
<community_ name> isanSNMPcommunity nameaddedtotheFortiGate configuration. Youcanaddmore
thanonecommunity nametoaFortiGate SNMPconfiguration. Themostcommonlyusedcommunity nameis
public.
<address_ipv4> istheIPaddressoftheFortiGate interfacethat theSNMPmanagerconnectsto.
{<OID> | <MIB_field>} istheobjectidentifier (OID)fortheMIBfield ortheMIBfield nameitself. TheHA
MIBfieldsandOIDsarelistedbelow:
HighAvailability
Fortinet TechnologiesInc.
181

ClustersandSNMP Operatingclustersandvirtualclusters
SNMP field names and OIDs
MIB field OID Description
fgHaSystemMode .1.3.6.1.4.1.12356.101.13.1.1.0 HAmode(standalone,a-a,ora-p)
fgHaGroupId .1.3.6.1.4.1.12356.101.13.1.2.0 TheHAgroupIDoftheclusterunit.
fgHaPriority .1.3.6.1.4.1.12356.101.13.1.3.0 TheHApriorityoftheclusterunit. Default
128.
fgHaOverride .1.3.6.1.4.1.12356.101.13.1.4.0 WhetherHAoverrideisdisabledorenabled
fortheclusterunit.
fgHaAutoSync .1.3.6.1.4.1.12356.101.13.1.5.0 Whetherautomatic HAsynchronizationis
disabledorenabled.
fgHaSchedule .1.3.6.1.4.1.12356.101.13.1.6.0 TheHAloadbalancingschedule.Setto
noneunlessoperatingina-pmode.
fgHaGroupName .1.3.6.1.4.1.12356.101.13.1.7.0 TheHAgroupname.
fgHaStatsIndex .1.3.6.1.4.1.12356.101.13.2.1.1.1.1
Anindexvaluethat identifies the
FortiGates inanHA cluster.Theindexis
always1fortheFortiGate that receivesthe
HAget. TheotherFortiGate(s)inthe
clusterwillhaveanindexof2,3,or4.For
example,ifyougetthestatsindexfromthe
primaryFortiGate, theprimaryFortiGate
willhaveastatsindexof1andthebackup
FortiGate willhaveastatsindexof2.If you
getthestatsindexfromthebackupunit,
thebackupunitwillhaveastatsindexof1
andtheprimaryunitwillhaveastatsindex
of2.
fgHaStatsSerial .1.3.6.1.4.1.12356.101.13.2.1.1.2.1 Theserialnumberoftheclusterunit.
fgHaStatsCpuUsage .1.3.6.1.4.1.12356.101.13.2.1.1.3.1 Theclusterunit’scurrentCPUusage.
fgHaStatsMemUsage .1.3.6.1.4.1.12356.101.13.2.1.1.4.1 Theclusterunit’scurrentMemoryusage.
fgHaStatsNetUsage .1.3.6.1.4.1.12356.101.13.2.1.1.5.1 Theclusterunit’scurrentNetwork
bandwidthusage.
fgHaStatsSesCount .1.3.6.1.4.1.12356.101.13.2.1.1.6.1 Theclusterunit’scurrentsessioncount.
fgHaStatsPktCount .1.3.6.1.4.1.12356.101.13.2.1.1.7.1 Theclusterunit’scurrentpacketcount.
fgHaStatsByteCount .1.3.6.1.4.1.12356.101.13.2.1.1.8.1 Theclusterunit’scurrentbytecount.
182 HighAvailability
Fortinet TechnologiesInc.

Operatingclustersandvirtualclusters ClustersandSNMP
MIB field OID Description
fgHaStatsIdsCount .1.3.6.1.4.1.12356.101.13.2.1.1.9.1 ThenumberofattacksreportedbytheIPS
fortheclusterunit.
fgHaStatsAvCount .1.3.6.1.4.1.12356.101.13.2.1.1.10.1 Thenumberofvirusesreportedbythe
antivirussystemfortheclusterunit.
fgHaStatsHostname .1.3.6.1.4.1.12356.101.13.2.1.1.11.1 Thehostnameoftheclusterunit.
To get the HA priority for the primary unit
Thefollowing SNMPgetcommandgetstheHApriorityfortheprimaryunit. Thecommunity nameispublic.
TheIPaddressoftheclusterinterfaceconfiguredforSNMPmanagement accessis10.10.10.1. TheHApriority
MIBfield isfgHaPriorityandtheOIDforthisMIBfield is1.3.6.1.4.1.12356.101.13.1.3.0 Thefirstcommanduses
theMIBfield nameandthesecondusestheOID:
snmpget -v2c -c public 10.10.10.1 fgHaPriority
snmpget -v2c -c public 10.10.10.1 1.3.6.1.4.1.12356.101.13.1.3.0
SNMP get command syntax for any cluster unit
Togetconfigurationstatusinformation foraspecificclusterunit(fortheprimaryunitorforanysubordinateunit),
theSNMPmanagermustaddtheserialnumberoftheclusterunittotheSNMPgetcommandafterthe
community name. Thecommunity nameandtheserialnumberareseparatedwithadash.Thesyntaxforthis
SNMPgetcommandwouldbe:
snmpget -v2c -c <community_name>-<fgt_serial> <address_ipv4> {<OID> | <MIB_field>}
where:
<community_ name> isanSNMPcommunity nameaddedtotheFortiGate configuration. Youcanaddmore
thanonecommunity nametoaFortiGate SNMPconfiguration. Allunitsintheclusterhavethesamecommunity
name. Themostcommonlyusedcommunity nameispublic.
<fgt_serial> istheserialnumberofanyclusterunit. Forexample,FGT4002803033172.Youcanspecifythe
serialnumberofanyclusterunit, includingtheprimaryunit, togetinformation forthat unit.
<address_ipv4> istheIPaddressoftheFortiGate interfacethat theSNMPmanagerconnectsto.
{<OID> | <MIB_field>} istheobjectidentifier (OID)fortheMIBfield ortheMIBfield nameitself.
If theserialnumbermatchestheserialnumberofasubordinateunit, theSNMPgetrequestissentovertheHA
heartbeatlinktothesubordinateunit. Afterprocessingtherequest,thesubordinateunitsendsthereplybackover
theHAheartbeatlinkbacktotheprimaryunit. TheprimaryunitthenforwardstheresponsebacktotheSNMP
manager.
If theserialnumbermatchestheserialnumberoftheprimaryunit, theSNMPgetrequestisprocessedbythe
primaryunit. Youcanactuallyaddaserialnumbertothecommunity nameofanySNMPgetrequest.But
normallyyouonlyneedtodothisforgetting information fromasubordinateunit.
To get the CPU usage for a subordinate unit
Thefollowing SNMPgetcommandgetstheCPUusageforasubordinateunitinaFortiGate-5001SXcluster.The
subordinateunithasserialnumberFG50012205400050.Thecommunity nameispublic.TheIPaddressof
HighAvailability
Fortinet TechnologiesInc.
183

