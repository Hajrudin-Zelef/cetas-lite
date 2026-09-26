---
id: collect-260926-mikrotik/mikrotik/router-upgrade-problem-routeros-powerpc-7-1-3-npk-missing-use-ignore-missing-or-disable-pa
title: "router-upgrade-problem-routeros-powerpc-7-1-3-npk-missing-use-ignore-missing-or-disable-package-s"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/router-upgrade-problem-routeros-powerpc-7-1-3-npk-missing-use-ignore-missing-or-disable-package-s.md
source_anchor: ""
source_lines: [1, 39]
sha256: 83f3ab2842a1d54070e37f3e4f83746be852c935c02650f25ebefdb0ba500ac8
---

# router-upgrade-problem-routeros-powerpc-7-1-3-npk-missing-use-ignore-missing-or-disable-package-s

Hi,

I try to upgrade my router os from 6.49.5 to 7.1.3 but getting error :

**routeros-powerpc-7.1.3.npk missing, use ignore-missing or disable package(s)**


screen shot is attached

My router is RB1100


             
            
           
          
            
            
              Hello

it is quite annoying but i believe that is happening the reason of the package name for PPC ( powerpc)  have changed.

Example:

For the 6.xx version the package name is routeros-powerpc-6.xx.xx.npk

https://download.mikrotik.com/routeros/6.48.6/routeros-powerpc-6.48.6.npk

For the 7.xx version the package name is routeros-7.2.2-ppc.npk

https://download.mikrotik.com/routeros/7.2.2/routeros-7.2.2-ppc.npk

The only solution is to run the upgrade for this architecture, for now,  by downloading the routeros-7.x.x-ppc.npk on your PC and the uploading it to the router.

After the file is uploaded reboot the router to proceed with the upgrade.

Best Regards,

Daniel
