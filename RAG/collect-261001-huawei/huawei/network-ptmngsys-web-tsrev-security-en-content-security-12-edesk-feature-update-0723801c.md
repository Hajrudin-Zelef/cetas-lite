---
id: collect-261001-huawei/huawei/network-ptmngsys-web-tsrev-security-en-content-security-12-edesk-feature-update-0723801c
title: "network-ptmngsys-web-tsrev-security-en-content-security-12-edesk-feature-update--0723801c"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/network-ptmngsys-web-tsrev-security-en-content-security-12-edesk-feature-update--0723801c.md
source_anchor: ""
source_lines: [1, 21]
sha256: 72baf5f90288ecd04b7c7075b5819174797866183fcbd60dac7a3e2cb9afa644
---

# network-ptmngsys-web-tsrev-security-en-content-security-12-edesk-feature-update--0723801c

If the upgrade fails due to insufficient space of the CF card, the system displays The remaining space of the CF card is insufficient. You can check the available space of the CF card using the CLI or web UI mode. Use the CLI mode as an example.
<HUAWEI> dir /all 
Directory of hda1:/ 
 
   Idx Attr Size(Byte) Date        Time       FileName 
   0   -rw- 171477561  May 05 2016 16:31:40   usg6000v100r001c30spc600.bin 
   1   -rw-     57465  May 24 2016 18:40:50   vrpcfg.cfg 
   2   drwh         -  May 23 2016 11:33:24   default-sdb 
   3   drw-         -  Oct 23 2015 20:21:46   conf 
   4   drw-         -  Oct 23 2015 20:21:46   web 
   5   drw-         -  May 23 2016 11:31:42   logo 
   6   drw-         -  Oct 23 2015 20:21:46   log 
   7   drwh         -  May 23 2016 11:45:26   gpmbak 
   8   drw-         -  May 23 2016 11:31:48   isp 
   9   drwh         -  May 08 2016 18:04:10   hidepkirsakey 
   10  -rw-      1268  May 11 2016 19:54:26   hostkey 
   11  -rw-       548  May 11 2016 19:54:26   serverkey 
    
1179680 KB total (495808 KB free)    //The information in brackets indicates the available space of the CF card
      The upgrade of signature databases of USG6000 V100R001 requires an available space larger than 250000 KB.
Contact technical support personnel to check whether certain files can be deleted. Do not randomly delete files. Otherwise, device failure may occur.
