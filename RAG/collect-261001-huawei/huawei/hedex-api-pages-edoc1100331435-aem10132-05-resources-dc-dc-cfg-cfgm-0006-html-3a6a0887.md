---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-cfgm-0006-html-3a6a0887
title: "hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-cfgm-0006-html-3a6a0887"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-cfgm-0006-html-3a6a0887.md
source_anchor: ""
source_lines: [1, 8]
sha256: e6214b3097896c09ba2a04adc3893830d1c2983e94fe788983c4b32721c7365f
---

# hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-cfgm-0006-html-3a6a0887

If the device is damaged unexpectedly, the configuration file cannot be restored. You can back up the configuration file in advance using one of the following methods:
Run the display current-configuration command and copy all the command output to a .txt file on a maintenance terminal. The configurations are then saved on the maintenance terminal.
If the configuration of a single command is too long, the configuration may be displayed in multiple lines on the terminal screen, depending on the terminal software. When copying a multi-line configuration from the screen to a .txt file, ensure that the configuration occupies one line in the .txt file. Otherwise, such a configuration may fail to be restored when the .txt file is used.
The current configuration file can be backed up immediately to the flash memory of the device. After the device starts, run the following commands to back up the configuration file to the flash memory of the device:
<Huawei> save config.cfg
<Huawei> copy config.cfg backup.cfg
To save the configuration in a directory other than the default storage device, specify an absolute path.
The device supports configuration file backup through FTP, SCP, TFTP, or SFTP. Configuration file backup through FTP or TFTP is simple but may cause security risks. In scenarios requiring high security, SFTP and SCP are recommended for configuration file backup. The following example describes the process of backing up the configuration file using FTP. For details about TFTP, SCP and SFTP, see "File Management" in NetEngine AR600, AR6100, AR6200, and AR6300 Configuration Guide - Basic Configuration.
