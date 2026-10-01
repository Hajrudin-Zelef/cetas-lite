---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-caddy-reverse-proxy-with-lets-encrypt-and-crowdsec-using-opnsense-fb0acb0a-2
title: "caddy.service"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-caddy-reverse-proxy-with-lets-encrypt-and-crowdsec-using-opnsense--fb0acb0a.md
source_anchor: ""
source_lines: [61, 154]
sha256: 1f72db17d0bd4233a2ada1cd4901d99a406081a1e9be2359f63a59f324d68ef2
---

# caddy.service

What I typically do is use a shorter, more generic hostname for the override alias and a more specific hostname for the actual app/service. For example, the Homepage dashboard app used in this example has the hostname homepage-server, but I plan to access the frontend which is behind the reverse proxy via the hostname homepage. If I need to SSH into my Homepage dashboard server, I would use the homepage-server hostname.
Click on the “+” button on the bottom portion of the page to create an alias for the Homepage dashboard app example that I am using. The frontend hostname will be homepage.
Set Up the Caddy Reverse Proxy
The main portion of this guide is setting up Caddy as a reverse proxy using Let’s Encrypt DNS challenges as well as setting up the CrowdSec agent and CrowdSec bouncer module. While the basic installation of Caddy is simple, the process I am demonstrating is going to require a manual installation so that additional modules can be built into the Caddy executable.
As mentioned earlier, I will be setting up Caddy in a LXC on Proxmox. Log in as the root user to complete the steps below. Note that I am including sudo in the commands even though it is not necessary for the root user in case you are using a different user with sudo privileges. You will have to ensure the files have the proper owner/group.
Install Go
Caddy was built with the Go programming language so you will need to install it in order to complete the steps in this guide.
On the Go download page, find the latest stable release and copy the URL for the download. You will want the Linux amd64 version of the file if you are using a LXC in Proxmox. Using the wget command, you can download it directly to the LXC.
The next command will extract the contents to the proper location on the system.
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz
You will want the go executable on the PATH.
export PATH=$PATH:/usr/local/go/bin
For the export to take effect immediate, issue the following source command.
source $HOME/.profile
You may check the Go version using the go version command. If you have a version number displayed instead of an error, you have successfully installed Go.
go version
You may remove the Go download once you have it set up properly.
rm go1.24.4.linux-amd64.tar.gz
Install xcaddy
You will need to install xcaddy for the next step since a custom Caddy executable file needs to be built. The xcaddy tool helps with this process. You should be able to copy/paste this entire block of commands and hit enter to install xcaddy.
sudo apt install -y curl debian-keyring debian-archive-keyring apt-transport-https
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/xcaddy/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-xcaddy-archive-keyring.gpg
curl -1sLf 'https://dl.cloudsmith.io/public/caddy/xcaddy/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-xcaddy.list
sudo apt update
sudo apt install xcaddy
Build Caddy with the Desired Modules
In order to use DNS challenges with Let’s Encrypt and to use a CrowdSec bouncer, you will need to a build a custom Caddy executable to extend the base functionality. Fortunately, the build process is easy with xcaddy since you can build a Caddy executable with a single command.
I will be using the Cloudflare module for Let’s Encrypt but you may use a different provider. Replace github.com/caddy-dns/cloudflare with a provider from the list found on GitHub.
xcaddy build \
    --with github.com/caddy-dns/cloudflare \
    --with github.com/hslatman/caddy-crowdsec-bouncer/http
Move the file to the /usr/bin/ folder:
sudo mv caddy /usr/bin/
You should be able to run the caddy executable to ensure it can be found on the path.
caddy version
Create User and Group for Caddy
Now is a good time to create a separate user and group for Caddy since the caddy executable has been built and moved to the proper location.
Create the caddy user group using the following command:
sudo groupadd --system caddy
Then create the caddy user and add it to the caddy group:
sudo useradd --system \
    --gid caddy \
    --create-home \
    --home-dir /var/lib/caddy \
    --shell /usr/sbin/nologin \
    --comment "Caddy web server" \
    caddy
Create Caddyfile
Caddy may be configured either by caddy.json or in a file called a Caddyfile, which is less verbose than using a JSON file. I will use the Caddyfile since for the purposes of this demonstration will suffice. If you need more advanced configuration, you can use the JSON configuration file instead.
mkdir /etc/caddy
touch /etc/caddy/Caddyfile
chown caddy:caddy /etc/caddy/Caddyfile
nano /etc/caddy/Caddyfile
In the configuration file, you will need to enter the Cloudflare API key used for editing DNS zones for the acme_dns cloudflare option.
In addition to the DNS API key, newer versions of Caddy (v2.8.0+) require you to enter an email address for ZeroSSL (Caddy uses both Let’s Encrypt and ZeroSSL for issuing certificates). I discovered I needed to add the email to the configuration when I tried upgrading Caddy.
Then in a crowdsec block, you will need to enter the API key what was generated from OPNsense earlier. The URL is for the CrowdSec LAPI on OPNsense which is 192.168.1.1:8080.
These first two settings should be contained in the global settings block as shown below.
Finally, you will need to include entries for your apps/services that will be behind your reverse proxy. I included one example of a Homepage dashboard app that I have running on my network. You simply include the URL next to the reverse_proxy option.
Please note that the Caddy server needs to have access to the apps/services if they live on a different network/VLAN so you will need to create the appropriate firewall rules.
Be sure to include the route section with crowdsec before the reverse_proxy line so that it will route the traffic through CrowdSec first. The log section is important because log files need to be generated for CrowdSec to parse for malicious activity.
{
        email youremail@example.com
        acme_dns cloudflare zNA25a5Wm8jxXesvShLP5U7kP2ENybNl7x8a9kxz
        crowdsec {
                api_key kLI6ljt3B/zWsBvRxu68vP8rL/cNj3O0hgjJ6B2s7Yk
                api_url http://192.168.1.1:8080/
        }
}
homepage.homenetworkguy.com {
        route {
                crowdsec
                reverse_proxy http://homepage-server.homenetworkguy.com:3000
        }
        log {
                output file /var/log/caddy/caddy.log
        }
}
There are 3 additional CrowdSec bouncer options you may include.
If you wish for the bouncer to check the LAPI each time instead of caching the decisions in memory and polling the LAPI every 10 seconds by default, you can disable streaming by adding the disable_streaming option to the crowdsec block. Streaming decision information is more efficient if you have a lot of requests, but it is possible there will be a slight increase in delay when decisions on the LAPI have changed.
...
        crowdsec {
            ...
            disable_streaming
...
If streaming is used (which is used by the default), then you may wan to tweak the ticker_interval option to control how many seconds to query the CrowdSec LAPI. The default is 10 seconds but you may check more or less frequently.
...
        crowdsec {
            ...
            ticker_interval 10s
...
Finally, there is an option to enable “hard fails” if there is an issue connecting to the CrowdSec LAPI. This feature might be nice if you wish to prevent your services from being accessed if something is wrong with the LAPI since they will be unprotected by CrowdSec. Of course, that would negatively affect uptime, but it would make it very apparent something bad has happened.
I noticed it takes about 30 seconds to fail after the CrowdSec service is stopped in OPNsense. I was starting to wonder if this option was working properly, but I simply was not patient enough during my testing.
...
        crowdsec {
