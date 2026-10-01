---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f-3
title: "how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "intel", "memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f.md
source_anchor: ""
source_lines: [125, 182]
sha256: e49b5213357e6e49262629e63512c96c911c502bdfc843a692c7e081b1cb727a
---

# how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f

If you changed the “Protocol” to “HTTPS” when you are using “HTTP” or you changed the “TCP port” to a different value, you will need to change the URL you are using to access OPNsense accordingly (if you are not automatically redirected to the new URL).
Settings: Miscellaneous
The “System > Settings > Miscellaneous” page has a few options you may want to tweak such as the CPU type for the thermal sensors widget on the “Dashboard”. There are some periodic backup options, power savings options, and memory/swap options.
| Option | Value | 
|---|---|
| Thermal Sensors Hardware | Intel Core CPU (unless you have AMD hardware) | 
| Periodic RRD Backup | 24 hours (optional) | 
| Periodic DHCP Leases Backup | 24 hours (optional) | 
| Periodic NetFlow Backup | 24 hours (optional) | 
| Use PowerD | Checked (if you have power saving options enabled in the BIOS) | 
| Power Mode | Hiadaptive (to favor performance over power savings) | 
If you are using a SSD or a traditional hard disk, you should not need to adjust any of the disk/memory settings at the bottom of the page since those options are more designed for systems where you want to minimize wear on the disk or if disk space is very constrained. Modern SSDs can handle a lot of writes before the disks wear out.
Firmware: Status
If you have your new OPNsense box connected to your primary/existing network while you are setting it up, you may want to update to the latest OPNsense version before proceeding with the rest of the configuration described in this guide. Otherwise, you can perform this step once you complete this guide and you have swapped out your old router with your new OPNsense box.
Go to the “System > Firmware > Status” page and click the “Check for updates” button. It will jump over to the “Updates” tab while checking for updates. If any updates are available, you will be notified.
When you close the information box for an update, you can see at the bottom of the page if the update requires a reboot or not. This is very helpful in determining if you wish to proceed with an update. I will often wait to perform the update when nobody is using the network whenever I see that a reboot is required. Typically smaller updates will not require a reboot, but of course it depends on what part of the system is being updated.
Access: Users
Generally speaking, a security recommendation is to create a new user with administrator privileges rather than using the default root user. Once another user is created, you can actually disable the root user to minimize the likelihood of the root user account being used maliciously.
On the “System > Access > Users” page, click on the “+” button to add a new user. Enter the following information (use your own values where it makes sense):
| Option | Value | 
|---|---|
| Username | dustin | 
| Password | examplepassword | 
| Full Name | Dustin Casto | 
|  | example@homenetworkguy.com | 
| Login shell | /usr/local/bin/bash (to allow shell access) | 
| Group Memberships | admins (click admins and the right arrow button) | 
| OTP seed | Check “Generate new secret (160 bit)” (if using two factor authentication – requires server described in next section) | 
| Authorized keys | Paste a generated SSH key if you plan to use SSH for this account | 
If you wish to disable the root user, logout of the root user account and login with your new administrator account. Then go back to the “System > Access > Users” page. Click the pencil button next to the root user and click the “Disabled” check box. Save your changes.
Tip
If you only want to use the root user to log into OPNsense via console/SSH but not allow the root user access to the web interface, you can simply leave the root user enabled and remove the root user from the admins group. This only revokes the web interface access but not console/SSH access for the root user.
If you leave SSH password authentication disabled and only use SSH keys, your root account should still be more secure than using a password.
Note that you can only remove the root user from the admins group after creating another admin account.
After you save your account, you will be able to click the “Click to unhide” button for the “OTP QR code” option which is now visible. This allows you to use a QR code to add the TOTP to your favorite authenticator application on your phone or other device. If you are using a password manager, you can copy/paste the “OTP seed” to be able to generate codes from your password manager.
Access: Servers
The access servers in OPNsense allow you to add different authentication methods such as timebased one time password (TOTP) to enable two factor authentication. On the “System > Access > Servers” page, click on the “+” button to add a new authentication method. Enter the following values:
| Option | Value | 
|---|---|
| Descriptive name | TOTP server | 
| Type | Local + Timebased One Time Password | 
| Token length | 6 (the default value) | 
| Reverse token order | Checked | 
By default the token needs to be entered before the password, but if you check the “Reverse token order” option, the token should be entered after the password.
The reason I prefer to add the 6 digit token at the end of the password is that it works great for password managers such as BitWarden since the token is automatically copied to the clipboard after the username/password is filled it on the page. That allows you to click the password manager’s “auto fill” button and then immediately issue a “Ctrl + V” to paste the token to the end of the password. I find that to be convenient if you utilize password managers that also store the tokens.
Some users prefer not to store their second factor token in the same location as their passwords, which is certainly understandable from a security standpoint.
Note that at this point two factor authentication is not enabled yet. Before enabling it, you should test out the two factor authentication first, which leads us to the next section below.
Access: Tester
The “System > Access > Tester” page allows you to test out your other authentication methods before enabling them, which is very handy. You do not want to inadvertently lock yourself out the system.
To test out the TOTP server you just created, select the “TOTP server” from the dropdown menu. It will be named whatever you called it in the “Descriptive name” from the previous step. Simply enter the username and password + token of the admin user you just created. You will receive a pass/fail message after testing the user.
Access: Users (Again)
If the login for the new admin user passes the test from the previous step, log out and log back in as the new admin user before proceeding.
You can now remove the root user from the admins group to revoke root user access to the web interface. Go back to the “System > Access > Users” page, and click the pencil edit button for the root user.
For the “Group Memberships” option, click on “admins” in the “Member Of” box and click the left arrow to move it to “Not Member Of”. Click the “Save” button to persist the changes. You will no longer be able to log into OPNsense using your root user!
Keep in mind that you can still log into the server as the root user via the console or SSH if you have left it enabled earlier. If you allow root user access for SSH, I recommend using only SSH keys and disabling the password for SSH access to keep your root user account more secure.
Tip
If you are using your non-root user admin account when logging into SSH or the console, you can still access the same root user menu by entering the sudo su command. After entring your password, you will see the same menu options that is shown to the root user. For this command to work, you will need to ensure you have the “Sudo” option enabled on the “System > Settings > Administration” page.
Interface Configuration
