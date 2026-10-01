---
id: collect-261001-general-networking/general-networking/platform-management-sm-endpoint-management-install-and-get-started-device-enroll-c8a95a61
title: "platform-management-sm-endpoint-management-install-and-get-started-device-enroll-c8a95a61"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/platform-management-sm-endpoint-management-install-and-get-started-device-enroll-c8a95a61.md
source_anchor: ""
source_lines: [1, 18]
sha256: 11e623b1b52b212e060d8c99c604367d36bbc52388e5c7470449b27c125c4149
---

# platform-management-sm-endpoint-management-install-and-get-started-device-enroll-c8a95a61

Windows Imaging and Systems Manager
Systems Manager uses a Universally Unique Identifier (UUID) to identify different Windows computers in Dashboard.
Duplicated UUID Clients
When a computer has the Systems Manager client installed on it and is then used as the base image and pushed out to multiple computers, this will result in all of them sharing the same UUID. When the computers report in to the Dashboard they will not appear as unique clients because of the matching UUID. Due to this fact, imaging with Systems Manager is not recommended. If a client is imaged with Systems Manager installed the following steps will allow you to remove the old UUID and generate a new one.
- Run the MerakiPCCAgent.msi and select Remove. This will uninstall the Systems Manger agent from the computer.
- In the Registry find and delete the entry under...
 For 32-bit: HKEY_LOCAL_MACHINE\SOFTWARE\Meraki\MerakiPCCAgent\UUID
For 64-bit: HKEY_LOCAL_MACHINE\SOFTWARE\Wow6432Node\Meraki\MerakiPCCAgent\WindowsCachedUUID
- Reinstall the agent on the computer.
To image and install Systems Manager unattended follow the steps below.
Using Imaging Software to install SM
To deploy a Systems Manager agent as part of a software image follow these steps.
- Add the Windows installer to the PC that will be the base image (do not run it). This can be found under Systems Manager > Manage > Add devices > Windows in a Systems Manager network.
- To minimize the times a Client needs to be touched, consider making a .bat file script that will run the image on the client upon startup.This .BAT example will silently run the installer, delete the installer executable, and delete itself. Exclude all '<' '>' characters. Refer to Windows Enrollment documentation for information on the msiexec arguments:
 MSIEXEC /i <path to installer .msi> ENROLLMENT_CODE=<network enrollment code / enrollment string> ENROLL_TOKEN=<bulk enrollment token> /qn
 del <path to installer .msi>
del <path to .bat script>
- Assign this script as a startup script on the base image and it will run the first time the client is booted and the client will be added to Systems Manager.
