---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-cfgm-0007-html-e2e477ee
title: "hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-cfgm-0007-html-e2e477ee"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-cfgm-0007-html-e2e477ee.md
source_anchor: ""
source_lines: [1, 4]
sha256: 5c8c1845cee110d00c6802c8c78ea5e1d43b0ad8480547e9fc8e0632146e9b9b
---

# hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-cfgm-0007-html-e2e477ee

When incorrect configurations are performed and functions are abnormal, you can use one of the following methods to restore the configuration file:
After restoring the configuration file, restart the device for the configuration file to take effect. Run the startup saved-configuration command to specify the configuration file for the next startup (if the configuration file name is not changed, skip this step), and then run the reboot command to restart the device. When the message "Warning: All the configuration will be saved to the next startup configuration. Continue? [y/n]:" is displayed, enter n to prevent the current configuration from being saved to the backup configuration file.
This method specifies the backup configuration file saved in the flash memory as the current configuration file. When the device is working properly, run the following commands.
The device supports configuration file restoration through FTP, SCP, TFTP, or SFTP. Configuration file restoration through FTP or TFTP is simple but may cause security risks. In scenarios requiring high security, SFTP and SCP are recommended for configuration file restoration. The following example describes the process of restoring the configuration file on a PC using FTP. For details about TFTP, SCP and SFTP, see "File Management" in NetEngine AR600, AR6100, AR6200, and AR6300 Configuration Guide - Basic Configuration.
