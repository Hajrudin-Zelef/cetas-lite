---
id: collect-261001-cisco/cisco/enterprise-en-ne-router-troubleshooting-how-to-solve-the-vrrp-status-of-the-back-56e2c81d
title: "enterprise-en-ne-router-troubleshooting-how-to-solve-the-vrrp-status-of-the-back-56e2c81d"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-ne-router-troubleshooting-how-to-solve-the-vrrp-status-of-the-back-56e2c81d.md
source_anchor: ""
source_lines: [1, 28]
sha256: 4a5c5ecb05c10a953511e5874583d1963d1ec9171d502fc8aaf71578d0ad1e44
---

# enterprise-en-ne-router-troubleshooting-how-to-solve-the-vrrp-status-of-the-back-56e2c81d

When we use Huawei NE series routers, we sometimes encounter problems. How can these problems be solved? We provide the Troubleshooting Guide to help you.
If you encounter the VRRP Status of the Backup Device Flaps, perform the following steps:
1. Check whether the following log message VRRP/4/vrrpTrapNewMaster is displayed, which indicates that the backup device becomes the master device because the VRRP timer expires.
§ If this log message is displayed, go to Step 2.
§ If the preceding log message is not displayed and the fault persists, go to Step 7.
2. Run the display interface interface-type interface-number command several times to check whether the physical status of the interface transmitting the VRRP Advertisement packets flaps.
§ If the physical interface alternates between Up and Down, repair the network cable or rectify the fault on the interface. Then, go to Step 3.
§ If the physical interface status is stable and the fault persists, go to Step 3.
3. Run the display vrrp command several times to check whether the VRRP status of the backup device is always Backup.
§ If the backup device is in the Backup state, the VRRP backup group is working properly.
§ If the backup device is not in the Backup state, go to Step 4.
4. Run the display vrrp statistics command to check whether any VRRP Advertisement packets with the priority value of 0 have been received.
§ If such a packet has been received, go to Step 7.
§ If no such packet has been received, go to Step 5.
5. Run the display this command to check whether a policy for filtering packets is configured on the interface board.
§ If cpu-defend-policy is displayed in the command output, go to Step 6.
§ If cpu-defend-policy is not displayed in the command output, no policy for filtering packets is configured. Go to Step 7.
6. Run the display this command to check whether a policy (for example, an ACL rule) for filtering VRRP Advertisement packets is configured on the interface board.
§ If a policy for filtering VRRP Advertisement packets is configured, go to Step 7.
§ If no policy for filtering VRRP Advertisement packets is configured, go to Step 8.
7. Run the display vrrp command several times to check whether the backup device is always in the Backup state.
§ If the backup device remains in the Backup state, the VRRP backup group is working properly.
§ If the backup device status changes frequently, go to Step 8.
8. If the fault persists, collect the following information and contact technical support personnel.
§ Results of this troubleshooting procedure
§ Configuration files, log files, and alarm files of the devices
For more troubleshooting cases, see: NE40E Troubleshooting Guide V4.0 (VRPv8)
https://support.huawei.com/enterprise/en/doc/EDOC1000177634
