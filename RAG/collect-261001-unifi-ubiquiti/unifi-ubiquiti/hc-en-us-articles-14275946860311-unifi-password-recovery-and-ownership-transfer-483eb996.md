---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-14275946860311-unifi-password-recovery-and-ownership-transfer-483eb996
title: "hc-en-us-articles-14275946860311-unifi-password-recovery-and-ownership-transfer-483eb996"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-14275946860311-unifi-password-recovery-and-ownership-transfer-483eb996.md
source_anchor: ""
source_lines: [1, 28]
sha256: 85ea5507a92cabdee94d7eb0b858206f8e2988e6cbffeaf8185dfd459bd558ff
---

# hc-en-us-articles-14275946860311-unifi-password-recovery-and-ownership-transfer-483eb996

UniFi Password Recovery and Ownership Transfer
Use a UI Account to easily recover UniFi passwords and seamlessly transfer ownership of existing UniFi deployments.
Password Recovery
- Determine which of the three types of credentials you are using:
  - UI Account: A UI Account is used to sign in to the UniFi Site Manager to access all associated UniFi deployments when Remote Management is enabled. The email linked with a UI Account is also used receive password reset emails.
  - Local-Only: If Remote Management is not enabled, a local username and password are created during setup. These credentials are unique to a specific UniFi instance and cannot be recovered unless an SMTP server was manually created.
  - Standalone AP: If an AP is set up in Standalone Mode, a password is created during setup. This password is also found in the UniFi Mobile App under Configure > Device Credentials. This password will be unique to a specific AP and is not associated with any UI Account. Lost Standalone credentials cannot be recovered
- Users with a UI Account or local-only credentials should proceed to step (3). Users with Standalone APs should jump to step (6).
- Click the Forgot Password button on the login screen when attempting to connect to UniFi.
- You will receive a password reset email at the address associated with your UI Account, or at the SMTP server which was manually configured.
- If you did not receive an email, contact the original UniFi owner or administrator. This indicates that:
  - A UI Account was not created for that email address, or
  - The email address was not associated with a given instance of UniFi.
- If you are already the owner and still cannot access your credentials, there is no way of recovering the current UniFi deployment. You must factory reset all devices and set up UniFi as new. We do not have access to users' credentials.
  - If you are self-hosting the UniFi Network Server on Windows/macOS/Linux, you will also need to uninstall UniFi and set up your UniFi Network Server again.
Note: We highly recommend securely saving your device passwords to avoid a future lockout.
Transferring Ownership
Ownership transfer is only available to UniFi Consoles that support UniFi OS, such as the Dream Machine, CloudKey Gen2 Plus, or Network Video Recorder.
- Sign in to the Owner account. Click here to learn more about roles and permissions.
- Navigate to the Control Plane of your UniFi Console and click Transfer Ownership.
- Select the desired Owner from the dropdown. If the desired user is not in the list, they can be added in the Admins tab of your UniFi Console.
Transferring a Self-Hosted UniFi Network
UniFi Network Servers self-hosted on Windows, macOS, or Linux do not have Owners, only Administrators. A new Administrator can be invited in Settings > System > Administration.
I Lost Contact with the Original Owner
If you have inherited a UniFi deployment, or the original Owner is no longer available, then you will need to perform the following steps.
- Factory reset all UniFi devices. If you are self-hosting the UniFi Network Server on Windows/macOS/Linux, you will also need to uninstall UniFi.
- Set up UniFi.
- Re-adopt the UniFi devices.
