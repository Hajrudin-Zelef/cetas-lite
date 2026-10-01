---
id: collect-261001-cisco/cisco/questions-54639-cisco-3172-asks-for-lan-enterprise-services-pkg-license-to-enabl-59fd065f
title: "questions-54639-cisco-3172-asks-for-lan-enterprise-services-pkg-license-to-enabl-59fd065f"
domain: cisco
role: reference
task: reference
actors: ["Intel"]
dates: ["2016-08-02", "2016-09-02", "2016-25-01"]
keywords: ["license", "copyright", "ethernet", "intel", "licenses", "memory", "open source"]
source: docs/RAG/collect-261001-cisco/questions-54639-cisco-3172-asks-for-lan-enterprise-services-pkg-license-to-enabl-59fd065f.md
source_anchor: ""
source_lines: [1, 90]
sha256: dabd739b49ed1b89abc05b1041e9083a4db758301da282b0afbd0caea883c52c
---

# questions-54639-cisco-3172-asks-for-lan-enterprise-services-pkg-license-to-enabl-59fd065f

I have several Nexus 3100 switches in which I use OSPF and PIM. Cisco says that these features are included in the LAN_BASE_SERVICES_PKG which is included by default in the switches.
Switch1# sh license
license_xxxxxxxxx_5_1.lic:
SERVER this_host ANY
VENDOR cisco
FEATURE LAN_BASE_SERVICES_PKG cisco 1.0 permanent uncounted \
        VENDOR_STRING=<LIC_SOURCE>NEXUS_PRODUCTION</LIC_SOURCE><SKU>N3K-BAS1K9</SKU> \
        HOSTID=VDH=xxxxxxxxxxx \
        NOTICE=<LicFileID>201609xxxxxxx00</LicFileID><LicLineID>1</LicLineID><PAK>N3K-C3172TQ-XLxxxxxxxx</PAK> \
        SIGN=xxxxxxx
Switch1# sh license usage
Feature                      Ins  Lic   Status Expiry Date Comments
                                 Count
--------------------------------------------------------------------------------
16P_UPG_PKG                   No    -   Unused             -
32P_LIC_PKG                   No    -   Unused             -
NXOS_OE_PKG                   No    -   Unused             -
24P_N3K_LIC_PKG               No    -   Unused             -
48P_N3K_UPG_PKG               No    -   Unused             -
LAN_BASE_SERVICES_PKG         Yes   -   Unused Never       -
LAN_ENTERPRISE_SERVICES_PKG   No    -   Unused             -
--------------------------------------------------------------------------------
Although, if I see the log I find these messages:
2018 Nov  8 12:47:05.724 Switch1 [11231]: !!!!!!  WARNING: 'LAN_ENTERPRISE_SERVICES_PKG' LICENSE NOT FOUND ON THE SYSTEM  !!!!!!. You have tried to enable a licensed feature [ospf] without installing the 'LAN_ENTERPRISE_SERVICES_PKG'
license, which has not been found on the system.
2018 Nov  8 12:47:05.724 Switch1 [11231]: If you have enabled this feature in error, please disable the feature. If you have not purchased the license, please contact your Cisco account team to order the license. It is in violation of
 Cisco's Terms of Use to use this feature without purchasing the license.
2018 Nov  8 12:47:05.724 Switch1 [11231]: !!!!!!  WARNING: 'LAN_ENTERPRISE_SERVICES_PKG' LICENSE NOT FOUND ON THE SYSTEM  !!!!!!. You have tried to enable a licensed feature [pim] without installing the 'LAN_ENTERPRISE_SERVICES_PKG' l
icense, which has not been found on the system.
2018 Nov  8 12:47:05.724 Switch1 [11231]: If you have enabled this feature in error, please disable the feature. If you have not purchased the license, please contact your Cisco account team to order the license. It is in violation of
 Cisco's Terms of Use to use this feature without purchasing the license.
The Cisco NX-OS licensing guide says:
  Layer 3 Base Services Package
LAN_BASE_SERVICES_PKG
N3K-BAS1K9
Static routing
RIPv2
EIGRP stub
OSPFv2 (limited routes)
PIMv2 (sparse mode)
Note
The Cisco Nexus 3164Q, 31128PQ, 3232C, and 3264Q switches and the Cisco Nexus 3100 Series switches in N9K mode run Cisco Nexus 9000 Series software and follow the Cisco Nexus 9000 license model. Therefore, these switches require the Layer 3 Enterprise Services Package (rather than the Layer 3 Base Services Package) to support routing protocols.
But I have confirmed that I'm in N3K mode:
Switch1# sh system switch-mode
!WARNING: switch-mode is not configured. Running in default  mode.
system switch-mode n3k
The output of show version is:
Cisco Nexus Operating System (NX-OS) Software
TAC support: http://www.cisco.com/tac
Copyright (C) 2002-2016, Cisco and/or its affiliates.
All rights reserved.
The copyrights to certain works contained in this software are
owned by other third parties and used and distributed under their own
licenses, such as open source.  This software is provided "as is," and unless
otherwise stated, there is no warranty, express or implied, including but not
limited to warranties of merchantability and fitness for a particular purpose.
Certain components of this software are licensed under
the GNU General Public License (GPL) version 2.0 or
GNU General Public License (GPL) version 3.0  or the GNU
Lesser General Public License (LGPL) Version 2.1 or
Lesser General Public License (LGPL) Version 2.0.
A copy of each such license is available at
http://www.opensource.org/licenses/gpl-2.0.php and
http://opensource.org/licenses/gpl-3.0.html and
http://www.opensource.org/licenses/lgpl-2.1.php and
http://www.gnu.org/licenses/old-licenses/library.txt.
Software
  BIOS: version 3.5.0
  NXOS: version 7.0(3)I3(1)
  BIOS compile time:  01/25/2016
  NXOS image file is: bootflash:///nxos.7.0.3.I3.1.bin
  NXOS compile time:  2/8/2016 20:00:00 [02/09/2016 05:18:17]
Hardware
  cisco Nexus 3172T Chassis
  Intel(R) Core(TM) i3-3227U CPU @ 2.50GHz with 8153496 kB of memory.
  Processor Board ID FOC20332ZW0
  Device name: Switch1
  bootflash:   14900224 kB
  usb1:               0 kB (expansion flash)
Kernel uptime is 4 day(s), 0 hour(s), 44 minute(s), 31 second(s)
Last reset
  Reason: Unknown
  System version: 7.0(3)I3(1)
  Service:
plugin
  Core Plugin, Ethernet Plugin
Active Package(s):
I'm experiencing problems with PIM and would like to be sure that they are not related with this.
Has anyone experienced this problem?
show versionfeature ospfin config mode?
