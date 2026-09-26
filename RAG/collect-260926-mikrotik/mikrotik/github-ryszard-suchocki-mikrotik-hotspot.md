---
id: collect-260926-mikrotik/mikrotik/github-ryszard-suchocki-mikrotik-hotspot
title: "github-ryszard-suchocki-mikrotik-hotspot"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["apache"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/github-ryszard-suchocki-mikrotik-hotspot.md
source_anchor: ""
source_lines: [1, 33]
sha256: f82e6122fac35a0dedcdb221f73804f773d8076c4ab2d640de96bff668803c99
---

# github-ryszard-suchocki-mikrotik-hotspot

The following actions are required to use the code given in this repo:

Suppose your domain is `hotspot.example.com`. It can be setup like this:

```
cd /var/www
git clone https://github.com/splash-networks/mikrotik-yt-portal
mv mikrotik-yt-portal hotspot.example.com
cd /var/www/hotspot.example.com
```
Copy the `.env.example` file to `.env` and set the values of the given environment variables in it:

```
cp .env.example .env
nano .env
```
Navigate to public folder:

`cd /var/www/hotspot.example.com/public`

Use this link to install Composer. Then run `php composer.phar install` to install the packages given in `composer.json`.

Apache virtual host can be setup on the portal server using the instructions given here.

The portal files are in public folder in this repository. DocumentRoot will be:

`/var/www/hotspot.example.com/public`

It has been successfully tested with `RouterOS v7.4.1`

Add your domain to `login.html` in action field and upload it to Mikrotik router in Files => Hotspot folder.

To get rid of the "You are logged in" message during hotspot login, the file `alogin.html` may be replaced on the Mikrotik. The status page will not be shown after that.
