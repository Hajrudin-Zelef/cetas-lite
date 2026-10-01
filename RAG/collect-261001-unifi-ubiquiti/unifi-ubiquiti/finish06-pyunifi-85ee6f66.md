---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/finish06-pyunifi-85ee6f66
title: "finish06-pyunifi-85ee6f66"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: ["2012-11-15"]
keywords: ["parameters"]
source: docs/RAG/collect-261001-unifi-ubiquiti/finish06-pyunifi-85ee6f66.md
source_anchor: ""
source_lines: [1, 130]
sha256: c6bf98e183ee48a1a701cc70f6e6e1a45ed51c0c21766e93b01b8a19686bf27d
---

# finish06-pyunifi-85ee6f66

A rewrite of https://github.com/unifi-hackers/unifi-lab in cleaner Python. Forked from https://github.com/calmh/unifi-api due to unmaintained status and rewritten to use the Requests module.
Perform all pull requests against the development branch. Pull requests against the master branch will not be merged, but closed.
sudo pip install -U pyunifi
from pyunifi.controller import Controller
c = Controller('192.168.1.99', 'admin', 'p4ssw0rd')
for ap in c.get_aps():
	print('AP named %s with MAC %s' % (ap.get('name'), ap['mac']))
See also the scripts unifi-ls-clients and unifi-low-rssi-reconnect for more
examples of how to use the API.
Interact with a UniFi controller.
Uses the JSON interface on port 8443 (HTTPS) to communicate with a UniFi controller. Operations will raise unifi.controller.APIError on obvious problems (such as login failure), but many errors (such as disconnecting a nonexistant client) will go unreported.
Create a Controller object.
- host -- the address of the controller host; IP or name
- username -- the username to log in with
- password -- the password to log in with
- port -- the port of the controller host
- version -- the base version of the controller API [v4|v5|unifiOS|UDMP-unifiOS]
- site_id -- the site ID to access
- ssl_verify -- Verify the controllers SSL certificate, default=True, can also be False or "path/to/custom_cert.pem"
Add a client to the block list.
- mac -- the MAC address of the client to block.
Disconnects a client, forcing them to reassociate. Useful when the connection is of bad quality to force a rescan.
- mac -- the MAC address of the client to disconnect.
Return a list of Alerts.
Return a list of unarchived Alerts.
Return a list of Events.
Return a list of all AP:s, with significant information about each.
Return a list of all active clients, with significant information about each.
Return statistical data of the last 24h
Return statistical data last 24h from endtime
- endtime -- the last time of statistics.
Return a list of all known clients, with significant information about each.
Return a list of user groups with its rate limiting settings.
Update user group bandwidth settings.
- group_id -- Group ID to modify.
- down_kbps -- New bandwidth in KBPS for download.
- up_kbps -- New bandwidth in KBPS for upload.
Return high level health information on status of the setup
Return a list of configured WLANs with their configuration parameters.
Restart an access point (by MAC).
- mac -- the MAC address of the AP to restart.
Restart an access point (by name).
- name -- the name address of the AP to restart.
Remove a client from the block list.
- mac -- the MAC address of the client to unblock.
Archive all alerts of site.
Tells the controller to create a backup archive that can be downloaded with download_backup() and then be used to restore a controller on another machine.
Remember that this puts significant load on a controller for some time (depending on the amount of users and managed APs).
Tells the controller to create a backup archive and downloads it to a file. It should have a .unf extension for later restore.
- targetfile -- the target file name, you can also use a full path. Default creates unifi-backup.unf in the current directoy.
authorize_guest(self, guest_mac, minutes, up_bandwidth=None, down_bandwidth=None, byte_quota=None, ap_mac=None)
Authorize a guest based on his MAC address.
- guest_mac -- the guest MAC address : aa:bb:cc:dd:ee:ff
- minutes -- duration of the authorization in minutes
- up_bandwith -- up speed allowed in kbps (optional)
- down_bandwith -- down speed allowed in kbps (optional)
- byte_quota -- quantity of bytes allowed in MB (optional)
- ap_mac -- access point MAC address (UniFi >= 3.x) (optional)
Unauthorize a guest based on his MAC address.
- guest_mac -- the guest MAC address : aa:bb:cc:dd:ee:ff
Set client alias. Use "" to reset to the default.
- mac: The target MAC: aa:bb:cc:dd:ee:ff
- alias: The alias to set
create_voucher(self, number, quota, expire, up_bandwidth=None, down_bandwidth=None, byte_quota=None, note=None)
Create voucher for guests. Return list of new vouchers.
- number -- number of vouchers
- quota -- maximal number of using; 0 = unlimited
- expire -- expiration of vouchers in minutes
- up_bandwidth -- up speed allowed in kbps (optional)
- down_bandwidth -- down speed allowed in kbps (optional)
- byte_quota -- quantity of bytes allowed in MB (optional)
- note -- description of vouchers (optional)
Get list of vouchers.
- filter --  Voucher filter  (create_time, code, quota, used, note, status_expires, status, ...)
  c.list_vouchers(code='12345-67890')
Delete / revoke voucher.
- id -- voucher id
Gets the current state & configuration of the given device based on its MAC Address.
- target_mac -- MAC address of the device
Returns a list of all RADIUS users, name, password, 24 digit user id, and 24 digit site id.
Add a new RADIUS user with this username and password.
- name -- the new user's username
- password -- the new user's password
Update a RADIUS user to this new username and password.
Requires the user's 24 digit user id, which can be gotten from get_radius_users(self).
- name -- the user's new username
- password -- the user's new password
- id -- the user's 24 digit user id.
Delete a RADIUS user.
Requires the user's 24 digit user id, which can be gotten from get_radius_users(self).
- id -- the user's 24 digit user id.
Gets a list of port overrides, in dictionary format, for the given target MAC address. The dictionary contains the port_idx, portconf_id, poe_mode, & name.
- target_mac -- MAC address of the device
Powers Off the given port on the Switch identified by the given MAC Address.
- target_mac -- MAC address of the device
- port_idx -- Port ID to power off
Powers On the given port on the Switch identified by the given MAC Address.
- target_mac -- MAC address of the device
- port_idx -- Port ID to power on
The following small utilities are bundled with the API:
Lists the currently active clients on the networks. Can take the following parameters:
| Parameters | Description | Default | 
|---|---|---|
| -c | controller address | unifi | 
| -u | controller username | admin | 
| -p | controller password |  | 
| -b | controller port | 8443 | 
| -v | controller base version | v5 | 
| -s | site ID, UniFi >=3.x only | default | 
| -V | ignore SSL certificates |  | 
| -C | verify with ssl certificate pem file |  | 
jb@unifi:~ % unifi-ls-clients -c localhost -u admin -p p4ssw0rd -v v3 -s default
NAME                             MAC  AP            CHAN  RSSI   RX   TX
client-kitchen     00:24:36:9a:0d:ab  Study          100    51  300  216
jborg-mbp          28:cf:da:d6:46:20  Study          100    45  300  300
jb-iphone          48:60:bc:44:36:a4  Living Room      1    45   65   65
jb-ipad            1c:ab:a7:af:05:65  Living Room      1    22   52   65
Periodically checks all clients for low SNR values, and disconnects those who
fall below the limit. The point being that these clients will then try to
reassociate, hopefully finding a closer AP. Take the same parameters as above,
plus settings for intervals and SNR threshold. Use unifi-low-snr-reconnect -h
for an option summary.
A good source of understanding for RSSI/SNR values is this article. According to that, an SNR of 15 dB seems like a good cutoff, and that's also the default value in the script. You can set a higher value for testing:
jb@unifi:~ % unifi-low-snr-reconnect -c localhost -u admin -p p4ssw0rd -v v3 -s default --minsnr 30
2012-11-15 11:23:01 INFO unifi-low-snr-reconnect: Disconnecting jb-ipad/1c:ab:a7:af:05:65@Study (SNR 22 dB < 30 dB)
2012-11-15 11:23:01 INFO unifi-low-snr-reconnect: Disconnecting Annas-Iphone/74:e2:f5:97:da:7e@Living Room (SNR 29 dB < 30 dB)
For production use, launching the script into the background is recommended...
Get a csv file with statistics
unifi-save-statistics -c localhost -u admin -p p4ssw0rd -v v3 -s default -f filename.csv
MIT
