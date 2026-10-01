---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/guide-secure-unifi-cloud-controller-on-aws-lightsail-signed-with-lets-encrypt-ss-f5fdf716-2
title: "unifi_ssl_import.sh"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/guide-secure-unifi-cloud-controller-on-aws-lightsail-signed-with-lets-encrypt-ss-f5fdf716.md
source_anchor: ""
source_lines: [159, 203]
sha256: 0e350b982da56ed92968a81f2b2055dd0c6978aecd446b4daed9e99616108a35
---

# unifi_ssl_import.sh

This script will
- Backup the old keystore file (very handy, something i always forget to do)
- update the relevant keystore file with the LE cert
- restart the services to apply the new cert
7. Setup Automatic Certificate renewal
Lets-encrypt cert expeires every 3 months you can easily renew this by using
letsencrypt renew
This will use the existing config you used to generate the cert and renew it
then run the SSL-import script to update the controller cert
you can automate this using a cronjob
Copy the modified import Script you used in Step 6 to “/bin/certupdate/unifi_ssl_import.sh”
sudo mkdir /bin/certupdate/
cp /home/user/unifi_ssl_import.sh /bin/certupdate/unifi_ssl_import.sh
switch to sudo and edit your cron-tab for root and add the following lines
sudo su
crontab -e
0 1 31 1,3,5,7,9,11 * root certbot renew
15 1 31 1,3,5,7,9,11 * root /bin/certupdate/unifi_ssl_import.sh
Save and exit nano by doing CTRL+X followed by Y.
Check crontab for root and confirm
crontab -e
At 01:00 on day-of-month 31 in January, March, May, July, September, and November the command will attempt to renew the cert 
At 01:15 on day-of-month 31 in January, March, May, July, September, and November it will update the keystore with the new cert
Useful links –
https://kvz.io/blog/2007/07/29/schedule-tasks-on-linux-using-crontab/
8. Adopting UniFi devices to the new Controller with SSH or other L3 adoption methods 
1. Make sure the AP is running the same firmware as the controller. If it is not, see this guide: UniFi – Changing the Firmware of a UniFi Device.
2. Make sure the AP is in factory default state. If it’s not, do:
syswrapper.sh restore-default
3. SSH into the device and type the following and hit enter:
set-inform http://ip-of-controller:8080/inform
4. After issuing the set-inform, the UniFi device will show up for adoption. Once you click adopt, the device will appear to go offline.
5. Once the device goes offline, issue the  set-inform  command from step 3 again. This will permanently save the inform address, and the device will start provisioning.
Managing the Unify controller services
# to stop the controller
$ sudo service unifi stop
# to start the controller
$ sudo service unifi start
# to restart the controller
$ sudo service unifi restart
# to view the controller's current status
$ sudo service unifi status
Troubleshooting issues
cat /var/log/unifi/server.log
go through the system logs and google the issue, best part about ubiquity gear is the strong community support
