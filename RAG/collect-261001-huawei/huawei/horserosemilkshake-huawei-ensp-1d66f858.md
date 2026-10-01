---
id: collect-261001-huawei/huawei/horserosemilkshake-huawei-ensp-1d66f858
title: "horserosemilkshake-huawei-ensp-1d66f858"
domain: huawei
role: reference
task: reference
actors: ["Huawei", "Oracle"]
dates: []
keywords: ["memory", "training"]
source: docs/RAG/collect-261001-huawei/horserosemilkshake-huawei-ensp-1d66f858.md
source_anchor: ""
source_lines: [1, 20]
sha256: d5ae4408dc675347e9ec5ffa647eb7382f32e61193b5267a1d0c46f54c47ec31
---

# horserosemilkshake-huawei-ensp-1d66f858

huawei-ensp | 中文 |
Huawei eNSP (Enterprise Network Simulation Platform) is a network simulation tool designed for training and testing network configurations, particularly in enterprise environments. It allows users to create virtual network scenarios that replicate real-world configurations using Huawei's networking equipment. If you plan to take the HCIA/HCIP/HCIE certification exam, you can familiarize yourself with operating Huawei equipments (e.g., router, switch) on eNSP.
After 2019, Huawei has paused sharing eNSP for strategic reasons. A "professional version" is only available to training institutions or device dealers. However, I have a backup of eNSP for individual users.
Before running the eNSP installation executable, you should gather the following pre-requisites:
- Oracle VM VirtualBox 5.2.44 (https://download.virtualbox.org/virtualbox/5.2.44/VirtualBox-5.2.44-139111-Win.exe)
- WinPcap 4.1.3 (https://www.winpcap.org/install/bin/WinPcap_4_1_3.exe)
- Wireshark 4.4.6 (https://2.na.dl.wireshark.org/win64/Wireshark-4.4.6-x64.exe)
Having installed all the pre-requisites, download the eNSP installation executable (542.2 MB, the file may be too big for a virus scan, download at your own discretion). Check the release on the right hand side of this page, you should find a place to download the executable there. (https://github.com/horserosemilkshake/huawei-ensp/releases/tag/eNSP)
The installation processes should be simple, just click "Next" all the way and it will work.
The installed eNSP should launch normally at most cases, but starting a route device will (almost certainly) incur a 40 error.
You can test whether you have this problem by clicking one of the examples in the main page (e.g., 1-1RIPv1&v2), drag and select all machine, and start them by clicking the "Start Device" button in the toolbar:
To solve a 40 error, follow the below steps:
- 
Disable Hyper-V (Turn Windows feature on and off > Hyper-V):
- 
Disable memory integrity (Core isolation > Memory Integrity):
- 
Disable Hyper-V loading: Run cmd as an administrator. Runbcdedit . Ifhypervisorlaunchtype isOn orAuto , turn it off withbcdedit /set hypervisorlaunchtype off .
- 
Reboot the machine and run eNSP as an administrator. The 40 error should be resolved.
