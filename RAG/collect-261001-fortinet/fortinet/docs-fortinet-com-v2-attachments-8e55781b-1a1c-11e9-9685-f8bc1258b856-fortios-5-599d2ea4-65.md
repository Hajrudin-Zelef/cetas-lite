---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5-599d2ea4-65
title: "docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4.md
source_anchor: ""
source_lines: [8333, 8457]
sha256: 13eeec972a16848e64de23e544d9b22bf1d714c09c928cb65fe834dbf8892a4c
---

# docs-fortinet-com-v2-attachments-8e55781b-1a1c-11e9-9685-f8bc1258b856-fortios-5--599d2ea4

HAandfailoverprotection Synchronizingtheconfiguration
Out of Sync Message Configuration Object
HA_SYNC_SETTING_IDSCUSTOM = 0x22
HA_SYNC_SETTING_IDS_MONITOR = 0x23 /bin/ipsmonitor
HA_SYNC_SETTING_IDS_SENSOR = 0x24 /bin/ipsengine
HA_SYNC_SETTING_NIDS_LIB = 0x25 /data/lib/libips.so
HA_SYNC_SETTING_WEBLISTS = 0x30
HA_SYNC_SETTING_CONTENTFILTER = 0x31 /data/cmdb/webfilter.bword
HA_SYNC_SETTING_URLFILTER = 0x32 /data/cmdb/webfilter.urlfilter
HA_SYNC_SETTING_FTGD_OVRD = 0x33 /data/cmdb/webfilter.fgtd- ovrd
HA_SYNC_SETTING_FTGD_LRATING = 0x34 /data/cmdb/webfilter.fgtd- ovrd
HA_SYNC_SETTING_EMAILLISTS = 0x40
HA_SYNC_SETTING_EMAILCONTENT = 0x41 /data/cmdb/spamfilter.bword
HA_SYNC_SETTING_EMAILBWLIST = 0x42 /data/cmdb/spamfilter.emailbwl
HA_SYNC_SETTING_IPBWL = 0x43 /data/cmdb/spamfilter.ipbwl
HA_SYNC_SETTING_MHEADER = 0x44 /data/cmdb/spamfilter.mheader
HA_SYNC_SETTING_RBL = 0x45 /data/cmdb/spamfilter.rbl
HA_SYNC_SETTING_CERT_CONF = 0x50 /etc/cert/cert.conf
HA_SYNC_SETTING_CERT_CA = 0x51 /etc/cert/ca
HA_SYNC_SETTING_CERT_LOCAL = 0x52 /etc/cert/local
HA_SYNC_SETTING_CERT_CRL = 0x53 /etc/cert/crl
HA_SYNC_SETTING_DB_VER = 0x55
HA_GET_DETAIL_CSUM = 0x71
HA_SYNC_CC_SIG = 0x75 /etc/cc_sig.dat
HA_SYNC_CC_OP = 0x76 /etc/cc_op
HA_SYNC_CC_MAIN = 0x77 /etc/cc_main
HighAvailability
Fortinet TechnologiesInc.
225

Synchronizingtheconfiguration HAandfailoverprotection
Out of Sync Message Configuration Object
HA_SYNC_FTGD_CAT_LIST = 0x7a /migadmin/webfilter/ublock/ftgd/
data/
Comparing checksums of cluster units
Youcanusethediagnose sys ha checksum show commandtocomparetheconfigurationchecksumsof
allclusterunits. Theoutput ofthiscommandshowschecksumslabeledglobal andall aswellaschecksums
foreachoftheVDOMsincludingtheroot VDOM.Theget system ha-nonsync-csum commandcanbe
usedtodisplaysimilarinformation; however,thiscommandisintendedtobeusedbyFortiManager.
Theprimaryunitandsubordinateunitchecksumsshouldbethesame.If theyarenotyoucanusetheexecute
ha synchronize start commandtoforceasynchronization.
Thefollowing commandoutput isfortheprimaryunitofaclusterthat doesnothavemultiple VDOMsenabled:
diagnose sys ha checksum show
is_manage_master()=1, is_root_master()=1
debugzone
global: a0 7f a7 ff ac 00 d5 b6 82 37 cc 13 3e 0b 9b 77
root: 43 72 47 68 7b da 81 17 c8 f5 10 dd fd 6b e9 57
all: c5 90 ed 22 24 3e 96 06 44 35 b6 63 7c 84 88 d5
checksum
global: a0 7f a7 ff ac 00 d5 b6 82 37 cc 13 3e 0b 9b 77
root: 43 72 47 68 7b da 81 17 c8 f5 10 dd fd 6b e9 57
all: c5 90 ed 22 24 3e 96 06 44 35 b6 63 7c 84 88 d5
Thefollowing commandoutput isforasubordinateunitofthesamecluster:
diagnose sys ha checksum show
is_manage_master()=0, is_root_master()=0
debugzone
global: a0 7f a7 ff ac 00 d5 b6 82 37 cc 13 3e 0b 9b 77
root: 43 72 47 68 7b da 81 17 c8 f5 10 dd fd 6b e9 57
all: c5 90 ed 22 24 3e 96 06 44 35 b6 63 7c 84 88 d5
checksum
global: a0 7f a7 ff ac 00 d5 b6 82 37 cc 13 3e 0b 9b 77
root: 43 72 47 68 7b da 81 17 c8 f5 10 dd fd 6b e9 57
all: c5 90 ed 22 24 3e 96 06 44 35 b6 63 7c 84 88 d5
Thefollowing exampleshowsusingthiscommandfortheprimaryunitofaclusterwithmultiple VDOMs.Two
VDOMshavebeenaddednamedtest andEng_vdm.
Fromtheprimaryunit:
config global
diagnose sys ha checksum show
is_manage_master()=1, is_root_master()=1
debugzone
global: 65 75 88 97 2d 58 1b bf 38 d3 3d 52 5b 0e 30 a9
test: a5 16 34 8c 7a 46 d6 a4 1e 1f c8 64 ec 1b 53 fe
root: 3c 12 45 98 69 f2 d8 08 24 cf 02 ea 71 57 a7 01
Eng_vdm: 64 51 7c 58 97 79 b1 b3 b3 ed 5c ec cd 07 74 09
all: 30 68 77 82 a1 5d 13 99 d1 42 a3 2f 9f b9 15 53
226 HighAvailability
Fortinet TechnologiesInc.

HAandfailoverprotection Synchronizingtheconfiguration
checksum
global: 65 75 88 97 2d 58 1b bf 38 d3 3d 52 5b 0e 30 a9
test: a5 16 34 8c 7a 46 d6 a4 1e 1f c8 64 ec 1b 53 fe
root: 3c 12 45 98 69 f2 d8 08 24 cf 02 ea 71 57 a7 01
Eng_vdm: 64 51 7c 58 97 79 b1 b3 b3 ed 5c ec cd 07 74 09
all: 30 68 77 82 a1 5d 13 99 d1 42 a3 2f 9f b9 15 53
Fromthesubordinateunit:
config global
diagnose sys ha checksum show
is_manage_master()=0, is_root_master()=0
debugzone
global: 65 75 88 97 2d 58 1b bf 38 d3 3d 52 5b 0e 30 a9
test: a5 16 34 8c 7a 46 d6 a4 1e 1f c8 64 ec 1b 53 fe
root: 3c 12 45 98 69 f2 d8 08 24 cf 02 ea 71 57 a7 01
Eng_vdm: 64 51 7c 58 97 79 b1 b3 b3 ed 5c ec cd 07 74 09
all: 30 68 77 82 a1 5d 13 99 d1 42 a3 2f 9f b9 15 53
checksum
global: 65 75 88 97 2d 58 1b bf 38 d3 3d 52 5b 0e 30 a9
test: a5 16 34 8c 7a 46 d6 a4 1e 1f c8 64 ec 1b 53 fe
root: 3c 12 45 98 69 f2 d8 08 24 cf 02 ea 71 57 a7 01
Eng_vdm: 64 51 7c 58 97 79 b1 b3 b3 ed 5c ec cd 07 74 09
all: 30 68 77 82 a1 5d 13 99 d1 42 a3 2f 9f b9 15 53
How to diagnose HA out of sync messages
Thissectiondescribeshowtousethediagnose sys ha checksum show anddiagnose debug
commandstodiagnosethecauseofHAoutofsyncmessages.
If HAsynchronizationisnotsuccessful,usethefollowing proceduresoneachclusterunittofindthecause.
To determine why HA synchronization does not occur
1. ConnecttoeachclusterunitCLIbyconnectedtotheconsoleport.
2. Enterthefollowing commandstoenabledebugginganddisplayHAoutofsyncmessages.
diagnose debug enable
diagnose debug console timestamp enable
diagnose debug application hatalk -1
diagnose debug application hasync -1
Collecttheconsoleoutput andcomparetheoutofsyncmessageswiththeinformation inthetable
HA outofsyncobjectmessagesandtheconfigurationobjectsthat theyreferenceonpage224.
3. Enterthefollowing commandstoturnoff debugging.
diagnose debug disable
diagnose debug reset
To determine what part of the configuration is causing the problem
If thepreviousproceduredisplaysmessagesthat includesyncobject0x30(forexample,HA_SYNC_SETTING_
CONFIGURATION = 0x03)thereisasynchronizationproblemwiththeconfiguration. Usethefollowing steps
todeterminethepartoftheconfigurationthat iscausingtheproblem.
HighAvailability
Fortinet TechnologiesInc.
227

