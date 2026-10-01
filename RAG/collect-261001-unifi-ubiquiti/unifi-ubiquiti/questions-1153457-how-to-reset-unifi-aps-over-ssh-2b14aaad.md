---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1153457-how-to-reset-unifi-aps-over-ssh-2b14aaad
title: "questions-1153457-how-to-reset-unifi-aps-over-ssh-2b14aaad"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1153457-how-to-reset-unifi-aps-over-ssh-2b14aaad.md
source_anchor: ""
source_lines: [1, 22]
sha256: 8a114248bb85801d9986b9286b25c2246b63a0e7959408a1d30ff65264b006d0
---

# questions-1153457-how-to-reset-unifi-aps-over-ssh-2b14aaad

I want to Reset APs or switches over the Build-in SSH. I do no find a working anser since this is just turning the APs of.
2 Answers 2
After searching for some time i found this article which had an expanded command on doing it.
For those who don´t need the whole explanation, the Command is as follows:
syswrapper.sh restore-default & set-default &
Connect to your AP over ssh using the IP address of the AP:
ssh -oHostKeyAlgorithms=+ssh-dss ubnt@192.168.1.20
By default, both username and password should be ubnt:
ubnt@192.168.1.20's password:
When connected to the access point use help to show available commands:
U6-IW-BZ.6.6.77# help
UniFi Command Line Interface - Ubiquiti Networks
   info                      display device information
   set-default               restore to factory default
   set-inform <inform_url>   attempt inform URL (e.g. set-inform http://192.168.0.8:8080/inform)
   upgrade <firmware_url>    upgrade firmware (e.g. upgrade http://192.168.0.8/unifi_fw.bin)
   fwupdate --url <firmware_url|firmware_name> [--dl-only] [--md5sum <sum_of_fw>]
            [--keep-firmware] [--keep-running] [--reboot-sys] 
                                   new firmware update command
   reboot                    reboot the device
You'll see that the reboot command will reboot the device:
reboot
