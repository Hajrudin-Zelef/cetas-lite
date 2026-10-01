---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2-3
title: "how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2.md
source_anchor: ""
source_lines: [69, 151]
sha256: f28fd8330aeade58f435c28831c382f3b56d69b5f4cb2df8de0bb9fb1da5c6da
---

# how-to-deploy-nginx-proxy-manager-in-dmz-with-opnsense-11b386e2

Go to the “Services > Unbound DNS > Overrides” page and click on the “+” button to create a new DNS override in the “Host Overrides” section. Enter the “Host” and the “Domain” names in the appropriate fields. Then enter the IP address you are planning to use for the reverse proxy. In my example, I am using 192.168.2.50 for the reverse proxy. You should do this for every hostname for all services you are proxying through your reverse proxy.
Prepare Your System for Nginx Proxy Manager
If you plan to have a dedicated system, container, or virtual machine, you should install the Operating System, create the container from a template, or create the virtual machine with the desired OS installed. I am personally using a LXC container in Proxmox so my instructions are Linux focused, but the general process to install Nginx Proxy Manager is essentially the same for other platforms. You may need to consult the Docker or Docker Compose documentation for more specific instructions on how to install them on your platform of choice.
Set up a Static DHCP Mapping in OPNsense
Once your system has been installed (OS, container, virtual machine), you can create a static DHCP mapping in OPNsense to give the system a fixed IP address which will be used for DNS overrides, firewall rules, etc. You may add the static DHCP mapping by navigating to the “Services > DHCPv4 > Leases” page or by going to the “Services > DHCPv4 > [DMZ]” page. On the leases page, you can scroll down the list to find the system and click the “+” button to add the lease. The nice thing about adding a static mapping from the “Leases” page is the MAC address is already prefilled. Otherwise, if you go to the “[DMZ]” page, you will have to enter it manually (or copy/paste the value from console output).
Enter the “IP address” for the proxy manager which is 192.168.2.50 in my example. You may choose the “Hostname” you prefer to use to access the proxy manager service. An optional “Description” may be entered as well.
The lease will not take effect immediately, but if you disconnect/reconnect to the network or reboot the system or container, you will obtain a new lease. If you are using Linux, you can use the following commands to refresh the lease to obtain the static IP address (make sure you are not connecting remotely to the system or you will lose your connection!):
sudo dhclient -r
sudo dhclient
Install Docker
To install Docker on a Linux distro such as Ubuntu, use the following commands:
sudo apt-get update
sudo apt-get install apt-transport-https ca-certificates curl gnupg lsb-release
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
Note: If you are installing Docker into a LXC container on Proxmox as I am doing, you will need to enable the container options called nesting and keyctl.
If you wish to enable IPv6 in Docker, you will need to edit the /etc/docker/daemon.json file to add the following contents (update the IPv6 prefix for the “fixed-cidr-v6” to the same /64 prefix you use for your DMZ network):
{
  "ipv6": true,
  "fixed-cidr-v6": "2001:db8:1::/64"
}
Restart Docker with the following command:
sudo systemctl reload docker
You can test if IPv6 is functioning properly if you access the administration interface via the IPv6 address in the browser once you deploy the container (yes, you need to include the brackets “[ ]” around the IPv6 address): http://[your_ipv6_address]:81
Install Docker Compose
Use the following commands to install Docker Compose:
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
Verify installation:
docker-compose --version
It will show the version number:
docker-compose version 1.29.2, build 5becea4c
Create Docker Compose File
Create the docker-compose.yml file with the following contents. You will need to update the passwords in the file to your own unique passwords. Please note that your_password needs to be the same password in both places, but the your_root_password can be and should be a different password:
version: '3'
services:
  app:
    image: 'jc21/nginx-proxy-manager:latest'
    restart: unless-stopped
    ports:
      - '80:80'
      - '81:81'
      - '443:443'
    environment:
      DB_MYSQL_HOST: 'db'
      DB_MYSQL_PORT: 3306
      DB_MYSQL_USER: 'npm'
      DB_MYSQL_PASSWORD: 'your_password'
      DB_MYSQL_NAME: 'npm'
    volumes:
      - ./data:/data
      - ./letsencrypt:/etc/letsencrypt
  db:
    image: 'jc21/mariadb-aria:latest'
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: 'your_root_password'
      MYSQL_DATABASE: 'npm'
      MYSQL_USER: 'npm'
      MYSQL_PASSWORD: 'your_password'
    volumes:
      - ./data/mysql:/var/lib/mysql
Deploy the Docker Container
Now you can start up the Docker container using the following Docker Compose command:
docker-compose up -d
You can check on the status of the Docker container by entering:
docker ps
Log into the Nginx Proxy Manager Administration
The default credentials are:
username: admin@example.com
password: changeme
You will be prompted to change your user information.
Then you will prompted to change your password.
After changing your administrator credentials, the “Users” page will be displayed.
The proxy is now ready to be configured.
Determine Hostnames for the Proxy Host and Services
A proxy host is the entry point to each service. The network traffic is redirected to the host you have configured for each proxy host. When adding proxy hosts, you will need to choose a hostname for the proxy host. The proxy host can use the same hostname used by the server/container/virtual machine. However, it may be better to use a hostname for the proxy host that you are not using for the service itself. You will be able to refer to your internal servers directly by the hostname when performing operations such as ping which is helpful for troubleshooting network connectivity, for instance.
One possible naming convention could be to use something such as git.homenetworkguy.com as the proxy hostname and gitlab.homenetworkguy.com for your internal server (you can simply use the gitlab hostname within your network and do not have to use the fully qualified domain name). What I like about this naming convention is the fact the proxy hostname is more generic and the internal server is more specific. If you decide to use Gitea instead of GitLab, you can still use the same proxy hostname but simply update the proxy host to redirect to the gitea.homenetworkguy.com hostname. You could also ping either hostname separately if you wanted to test connectivity to the proxy or to the server itself (if you have ping allowed in your network).
Set up the Reverse Proxy Hosts
To set up a proxy host, click on the “Hosts > Proxy Hosts” menu to open the “Proxy Hosts” page. Click on the “Add Proxy Host” button.
For the “Domain Names”, enter the hostname(s) which will be the entry point to access your service. As discussed in the previous section, the hostname does not have to be the same as an existing host on your network. If you plan to access the service from outside your home network, you will need to create the hostname/subdomain with your DNS registrar such as Cloudflare.
