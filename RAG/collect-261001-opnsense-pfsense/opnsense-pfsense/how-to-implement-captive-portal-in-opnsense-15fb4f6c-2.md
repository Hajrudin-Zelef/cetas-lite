---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-implement-captive-portal-in-opnsense-15fb4f6c-2
title: "how-to-implement-captive-portal-in-opnsense-15fb4f6c"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-implement-captive-portal-in-opnsense-15fb4f6c.md
source_anchor: ""
source_lines: [56, 79]
sha256: f0f4759d54f88b3fc706476114574d974e6dbfd454175171b3d527d21fc03092
---

# how-to-implement-captive-portal-in-opnsense-15fb4f6c

In this scenario, you will need to create an additional rule to allow access to the appropriate ports on the GUEST interface. The port range shown as the example in the OPNsense documentation is from 8000-10000:
| Option | Value | 
|---|---|
| Action | Pass | 
| Interface | GUEST | 
| Protocol | TCP | 
| Source | GUEST net | 
| Source Port | any | 
| Destination | GUEST address | 
| Destination Port | 8000 to 10000 | 
| Description | Allow access to captive portal login | 
Test Captive Portal Login
All you need to do once you finish setting up the captive portal is to try to access the Internet from a device on your guest network. You should receive a popup with the captive portal login when you connect to the network or you will be redirected to the captive portal login in the browser.
Simply enter the username/password from the local database or a voucher which you have created.
Create Vouchers for Testing
If you are using vouchers, you need to create a few vouchers to test on the “Services > Captive Portal > Vouchers” page and clicking on “Create vouchers”.
You may choose a value for the “Validity” option to set how long a user has access to the network. The clock does not start ticking until the user has logged in. An expiration may be set as well for the vouchers. 1 or more vouchers can be created at the same time. A “Groupname” may be entered to identify the vouchers more easily.
Click “Generate” once you set the desired options.
A CSV file will be downloaded. It contains the usernames/passwords for all of the vouchers you just generated. You may use that file to import into some other tool or template which may be distributed to your guests.
Keep in mind that the plain text passwords are not stored on the firewall for security reasons so you will only have them in the downloaded file. If you lose or delete the file, you will need to recreate the vouchers.
An example of what you will see in the CSV file:
Conclusion
That is all there is to implementing a captive portal in OPNsense! The process is not very difficult especially if you use authentication servers such as the local database and vouchers.
With other authentication servers such as Radius and LDAP, additional configuration will be required to set up the users. However, the configuration of the captive portal will remain nearly the same regardless of the authentication servers that are used.
