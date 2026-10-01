---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-45
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [6100, 6250]
sha256: 200477fc321f226caac02f6b213d4ab80b55b7dbfe3fdc31dc2d59ce3a3a80f5
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

Managingindividualclusterunitsusingareservedout-of-bandmanagement
interface
Operatingclustersandvirtual
clusters
1. LogintotheCLIofanyclusterunit.
2. Enterthefollowing commandtoenablethereservedmanagement interface, setport8asthereservedinterface,
andaddanIPv4default routeof10.11.101.2 andanIPv6default routeof2001:db8:0:2::20 forthereserved
management interface.
config system ha
set ha-mgmt-status enable
config ha-mgmt-interfaces
edit 1
set interface port8
set gateway 10.11.101.2
set gateway6 2001:db8:0:2::20
end
Thereservedmanagement interfacedefault routeisnotsynchronizedtootherclusterunits.
To change the primary unit reserved management interface configuration - GUI
YoucanchangetheIPaddressoftheprimaryunitreservedmanagement interfacefromtheprimaryunitGUI.
Configurationchangestothereservedmanagement interfacearenotsynchronizedtootherclusterunits.
1. FromaPContheinternalnetwork,browsetohttp://10.11.101.100 andlogintotheclusterGUI.
ThislogsyouintotheprimaryunitGUI.
Youcanidentify theprimaryunitfromitsserialnumberorhostnamethat appearsontheSystem
Information dashboardwidget.
2. Goto Network > Interfacesandedittheport8interfaceasfollows:
Alias primary_reserved
IP/Netmask 10.11.101.101/24
Administrative Access Ping,SSH,HTTPS,SNMP
3. Select OK.
YoucannowlogintotheprimaryunitGUIbybrowsingtohttps://10.11.101.101. Youcanalsologinto
thisprimaryunitCLIbyusinganSSHclienttoconnectto10.11.101.101.
To change subordinate unit reserved management interface configuration - CLI
Atthispointyoucannotconnecttothesubordinateunitreservedmanagement interfacebecauseitdoesnothave
anIPaddress.Instead, thisproceduredescribesconnectingtotheprimaryunitCLIandusingtheexecute ha
manage commandtoconnecttosubordinateunitCLItochangetheport8interface. Youcanalsouseaserial
connectiontotheclusterunitCLI. Configurationchangestothereservedmanagement interfacearenot
synchronizedtootherclusterunits.
1. ConnecttotheprimaryunitCLIandusetheexecute ha manage commandtoconnecttoasubordinateunit
CLI.
Youcanidentify thesubordinateunitfromisserialnumberorhostname. Thehostnameappearsin
theCLIprompt.
2. Enterthefollowing commandtochangetheport8IPaddressto10.11.101.102andsetmanagement access
toHTTPS,ping,SSH,andSNMP.
config system interface
edit port8
172 HighAvailability
Fortinet TechnologiesInc.

Operatingclustersandvirtual
clusters
Managingindividualclusterunitsusingareservedout-of-bandmanagement
interface
set ip 10.11.101.102/24
set allowaccess https ping ssh snmp
end
YoucannowlogintothesubordinateunitGUIbybrowsingtohttps://10.11.101.102. Youcanalsolog
intothissubordinateunitCLIbyusinganSSHclienttoconnectto10.11.101.102.
To configure the cluster for SNMP management using the reserved management interfaces- CLI
ThisproceduredescribeshowtoconfiguretheclustertoallowtheSNMPservertogetstatusinformation fromthe
primaryunitandthesubordinateunit. TheSNMPconfigurationissynchronizedtoallclusterunits. Tosupport
usingthereservedmanagement interfaces,youmustaddatleastoneHAdirectmanagement hosttoanSNMP
community. If yourSNMPconfigurationincludesSNMPuserswithusernamesandpasswordsyoumustalso
enableHAdirectmanagement forSNMPusers.
1. Enterthefollowing commandtoaddanSNMPcommunity calledCommunity andaddahosttothecommunity
forthereservedmanagement interfaceofeachclusterunit. ThehostincludestheIPaddressoftheSNMPserver
(10.11.101.20).
config system snmp community
edit 1
set name Community
config hosts
edit 1
set ha-direct enable
set ip 10.11.101.20
end
end
Enablingha-directinnon-HAenvironmentsmakesSNMPunusable.
2.
3. Enterthefollowing commandtoaddanSNMPuserforthereservedmanagement interface.
config system snmp user
edit 1
set ha-direct enable
set notify-hosts 10.11.101.20
end
Configureothersettingsasrequired.
To get CPU, memory, and network usage of each cluster unit using the reserved management IP
addresses
FromthecommandlineofanSNMPmanager,youcanusethefollowing SNMPcommandstogetCPU,memory
andnetworkusageinformation foreachclusterunit. Intheexamples,thecommunity nameisCommunity.The
commandsusetheMIBfield namesandOIDslistedbelow.
Enterthefollowing commandstogetCPU,memoryandnetworkusageinformation fortheprimaryunitwith
reservedmanagement IPaddress10.11.101.101 usingtheMIBfields:
snmpget -v2c -c Community 10.11.101.101 fgHaStatsCpuUsage
snmpget -v2c -c Community 10.11.101.101 fgHaStatsMemUsage
snmpget -v2c -c Community 10.11.101.101 fgHaStatsNetUsage
HighAvailability
Fortinet TechnologiesInc.
173

Managingindividualclusterunitsusingareservedout-of-bandmanagement
interface
Operatingclustersandvirtual
clusters
Enterthefollowing commandstogetCPU,memoryandnetworkusageinformation fortheprimaryunitwith
reservedmanagement IPaddress10.11.101.101 usingtheOIDs:
snmpget -v2c -c Community 10.11.101.101 1.3.6.1.4.1.12356.101.13.2.1.1.3.1
snmpget -v2c -c Community 10.11.101.101 1.3.6.1.4.1.12356.101.13.2.1.1.4.1
snmpget -v2c -c Community 10.11.101.101 1.3.6.1.4.1.12356.101.13.2.1.1.5.1
Enterthefollowing commandstogetCPU,memoryandnetworkusageinformation forthesubordinateunitwith
reservedmanagement IPaddress10.11.101.102 usingtheMIBfields:
snmpget -v2c -c Community 10.11.101.102 fgHaStatsCpuUsage
snmpget -v2c -c Community 10.11.101.102 fgHaStatsMemUsage
snmpget -v2c -c Community 10.11.101.102 fgHaStatsNetUsage
Enterthefollowing commandstogetCPU,memoryandnetworkusageinformation forthesubordinateunitwith
reservedmanagement IPaddress10.11.101.102 usingtheOIDs:
snmpget -v2c -c Community 10.11.101.102 1.3.6.1.4.1.12356.101.13.2.1.1.3.1
snmpget -v2c -c Community 10.11.101.102 1.3.6.1.4.1.12356.101.13.2.1.1.4.1
snmpget -v2c -c Community 10.11.101.102 1.3.6.1.4.1.12356.101.13.2.1.1.5.1
Adding firewall local-in policies for the dedicated HA management interface
Toaddlocal-inpolicesforthededicatedmanagement interface, enableha-mgmt-inft-only andsetintf to
any.Enablingha-mgmt-intf-only meansthelocal-inpolicyappliesonlytotheVDOMthat containsthe
dedicatedHAmanagement interface. Forexample:
config firewall local-in-policy
edit 0
set ha-mgmt-intf-only enable
set intf any
set srcaddr internal-net
set dstaddr mgmt-int
set action accept
set service HTTPS
set schedule weekdays
end
NTP over Dedicated HA management interfaces
If yousetupdedicatedmanagement interfacesoneachclusterunit, ifNTPisenabled,theprimaryunitcontacts
anNTP serverusingthededicatedmanagement interface. Systemtime isthensynchronizedtothebackupunits
throughtheHAheartbeat.
ExampleCLI:
config system interface
edit port5
set ip 172.16.79.46 255.255.255.0
end
config system ha
set group-name FGT-HA
set mode a-p
set ha-mgmt-status enable
config ha-mgmt-interfaces
edit 1
set interface port5
set gateway 172.16.79.1
174 HighAvailability
Fortinet TechnologiesInc.

