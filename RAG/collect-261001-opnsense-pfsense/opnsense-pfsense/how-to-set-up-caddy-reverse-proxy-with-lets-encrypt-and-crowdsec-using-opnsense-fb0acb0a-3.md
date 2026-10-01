---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-caddy-reverse-proxy-with-lets-encrypt-and-crowdsec-using-opnsense-fb0acb0a-3
title: "caddy.service"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["acquisition", "agent"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-caddy-reverse-proxy-with-lets-encrypt-and-crowdsec-using-opnsense--fb0acb0a.md
source_anchor: ""
source_lines: [155, 257]
sha256: 284059941252fe9eace3ed8a6f37f7565a0777e1a7f696cd1f107ad64c219521
---

# caddy.service

            ...
            enable_hard_fails
...
Press “Ctrl + O”, “Enter”, and “Ctrl + X” to save and close the Caddyfile.
Create the Caddy Log File
Caddy will not automatically generate the initial file for logging so you will need to do that before starting the Caddy service.
In the Caddyfile example, the log file specified as /var/log/caddy/caddy.log.
Since the caddy folder does not exist under /var/log, it will need created.
sudo mkdir /var/log/caddy
Then you can create a blank file for Caddy to use by using the following command:
sudo touch /var/log/caddy/caddy.log
You will need to set the proper permissions so that the caddy user can access the log file instead of the root user.
sudo chown caddy:caddy /var/log/caddy/caddy.log
Update the Caddy Collection Acquisition Template
You will need to add the location of the Caddy log file(s) so that CrowdSec can parse the appropriate files to detect malicious activity.
In the example Caddyfile above, the log file /var/log/caddy/caddy.log was configured but if you put the log files in a different location, you will need to adjust the values used below.
Open a text editor for the acquis.yaml file.
sudo nano /etc/crowdsec/acquis.yaml
At the bottom of the file, add the following text. You find the example from CrowdSec’s caddy collection page.
filenames:
  - /var/log/caddy/*.log
labels:
  type: caddy
---
Caddy Systemd Service
Because Caddy was manually built, there is no systemd service installed so you will need to do it manually. Fortunately, it is not difficult to do since Caddy provides an example systemd service file. You only need to do this step once even if you build new versions of Caddy in the future.
Open a text editor for the caddy.service file.
sudo nano /etc/systemd/system/caddy.service
Enter the contents found on Caddy’s GitHub page:
# caddy.service
#
# For using Caddy with a config file.
#
# Make sure the ExecStart and ExecReload commands are correct
# for your installation.
#
# See https://caddyserver.com/docs/install for instructions.
#
# WARNING: This service does not use the --resume flag, so if you
# use the API to make changes, they will be overwritten by the
# Caddyfile next time the service is restarted. If you intend to
# use Caddy's API to configure it, add the --resume flag to the
# `caddy run` command or use the caddy-api.service file instead.
[Unit]
Description=Caddy
Documentation=https://caddyserver.com/docs/
After=network.target network-online.target
Requires=network-online.target
[Service]
Type=notify
User=caddy
Group=caddy
ExecStart=/usr/bin/caddy run --environ --config /etc/caddy/Caddyfile
ExecReload=/usr/bin/caddy reload --config /etc/caddy/Caddyfile --force
TimeoutStopSec=5s
LimitNOFILE=1048576
PrivateTmp=true
ProtectSystem=full
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
[Install]
WantedBy=multi-user.target
Press “Ctrl + O”, “Enter”, and “Ctrl + X” to save and close the file.
You will now need to reload the Systemd daemon, enable, and start the Caddy service.
sudo systemctl daemon-reload
sudo systemctl enable --now caddy
At this point, you should verify Caddy is running without any errors:
sudo systemctl status caddy
If you see that it is running and not stopped, you should have a good Caddy server configuration with Let’s Encrypt and a CrowdSec bouncer enabled. You may check in the OPNsense CrowdSec interface on the “Overview > Bouncers” page if there has been any recent API pulls from the CrowdSec bouncer on Caddy. That is a good sanity check to ensure communication is successfully occurring between the bouncer and the LAPI on OPNsense.
Install CrowdSec Agent on Caddy Server
The CrowdSec bouncer is what will block connections to services behind the reverse proxy, but the Caddy server will need a CrowdSec agent installed so it can run the parsers and scenarios on the server. The agent sends the information to the LAPI on OPNsense to make decisions on blocking content. For a single CrowdSec instance, this usually occurs on the same system, but in a multi-server configuration, the CrowdSec agent and bouncer communicates with the LAPI on a different server (in this case, OPNsense).
Install CrowdSec using the following commands. Basically the next step is following the CrowdSec installation guide. It is very simple to install.
curl -s https://install.crowdsec.net | sudo sh
sudo apt install crowdsec
The CrowdSec agent needs to be registered with OPNsense CrowdSec LAPI using the command below.
sudo cscli lapi register -u http://192.168.1.1:8080 --machine caddyDmz
Note in the command above that you may provide a friendly name such as caddyDmz using the --machine parameter. Otherwise, it will generate a long random string. This is useful when you have multiple systems you are registering and you wish to provide a more descriptive name that shows up in the OPNsense web UI. Also, you can easily enter that descriptive name when validating the machine in a later step.
Next, copy the default CrowdSec systemd file from /lib/systemd/system to /etc/systemd/system so customizations can be made to the service file.
sudo cp /lib/systemd/system/crowdsec.service /etc/systemd/system/crowdsec.service
Edit the /etc/systemd/system/crowdsec.service to add the -no-api flag to the end of the ExecStart command. This disables the LAPI on the Caddy server since it is not needed because the LAPI on OPNsense will be used instead (see CrowdSec’s multi-server configuration example).
[Unit]
Description=Crowdsec agent
After=syslog.target network.target remote-fs.target nss-lookup.target
[Service]
Type=notify
Environment=LC_ALL=C LANG=C
ExecStartPre=/usr/bin/crowdsec -c /etc/crowdsec/config.yaml -t -error
ExecStart=/usr/bin/crowdsec -c /etc/crowdsec/config.yaml -no-api
#ExecStartPost=/bin/sleep 0.1
ExecReload=/bin/kill -HUP $MAINPID
Restart=always
RestartSec=60
[Install]
WantedBy=multi-user.target
You will need to reload the systemd service since changes to the service was made.
sudo systemctl daemon-reload
Do not reload the CrowdSec service just yet because with the LAPI on the Caddy server disabled, CrowdSec will error on startup until you have validated the Caddy machine on OPNsense. CrowdSec will be unable to connect to the LAPI on OPNsense until validation occurs. When there is no reachable LAPI, the CrowdSec agent will fail to load.
Validate the Caddy Machine in OPNsense
While logged into OPNsense via SSH or the console, list the machines which have been registered or requesting to be registered:
sudo cscli machines list
You should see similar output to below. Notice the status of the machine caddyDmz (if you defined a name, otherwise it will be something random such as 02a3nfadce4ez4b19zh582e0f68f72a4CX4EzFJ2Th4PNkj1). There will be a “No” symbol under “Status”.
──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
 Name                                               IP Address       Last Update            Status   Version                   Auth Type   Last Heartbeat 
──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
