---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1184320-how-to-use-nginx-as-a-reverse-proxy-on-a-ubiquiti-router-516f3883-2
title: "questions-1184320-how-to-use-nginx-as-a-reverse-proxy-on-a-ubiquiti-router-516f3883"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1184320-how-to-use-nginx-as-a-reverse-proxy-on-a-ubiquiti-router-516f3883.md
source_anchor: ""
source_lines: [80, 191]
sha256: 044bded5e035ada5799e86a2eb61731b2355cf0c793670b13b4d97e558e09b48
---

# questions-1184320-how-to-use-nginx-as-a-reverse-proxy-on-a-ubiquiti-router-516f3883

There are plenty of examples of nginx configuration files online, but the example below may prove beneficial if you plan on putting your router interface (or some other web site that uses websockets) behind the nginx reverse proxy.  Or if you're looking specifically to put a Synology NAS behind it (the photo station in particular is a bit tricky).
upstream edgemax {
  server 192.168.1.1:553;
  keepalive 32;
}
upstream nas {
  server 192.168.1.7:5001;
  keepalive 32;
}
upstream nasphoto {
  server 192.168.1.5:443;
  keepalive 32;
}
upstream nasfile {
  server 192.168.1.7:7001;
  keepalive 32;
}
server {
  listen 443 ssl http2;
  server_name router.*;
  
  ssl_certificate /config/user-data/ssl_chain_key.pem;
  ssl_certificate_key /config/user-data/ssl_chain_key.pem;
  
  client_max_body_size 0;
  proxy_http_version 1.1;
  proxy_buffering off;
  proxy_set_header Upgrade $http_upgrade;
  proxy_set_header Connection "Upgrade";
  proxy_set_header Host $host;
  proxy_set_header X-Real-IP $remote_addr;
  proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
  
  location / {
    proxy_pass https:
  }
}
server {
  listen 443 ssl http2;
  server_name nas.*;
  
  ssl_certificate /config/user-data/ssl_chain_key.pem;
  ssl_certificate_key /config/user-data/ssl_chain_key.pem;
  client_max_body_size 0;
  proxy_http_version 1.1;
  proxy_buffering off;
  proxy_set_header Upgrade $http_upgrade;
  proxy_set_header Connection "Upgrade";
  proxy_set_header Host $host;
  proxy_set_header X-Real-IP $remote_addr;
  proxy_set_header X-Forward-For $proxy_add_x_forwarded_for;
  location / {
    proxy_pass https:
  }
  
  location /photo {
    proxy_pass https:
  }
}
server {
  listen 443 ssl http2;
  server_name files.*;
  ssl_certificate /config/user-data/ssl_chain_key.pem;
  ssl_certificate_key /config/user-data/ssl_chain_key.pem;
  client_max_body_size 0;
  proxy_http_version 1.1;
  proxy_buffering off;
  proxy_set_header Upgrade $http_upgrade;
  proxy_set_header Connection "Upgrade";
  proxy_set_header Host $host;
  proxy_set_header X-Real-IP $remote_addr;
  proxy_set_header X-Forward-For $proxy_add_x_forwarded_for;
  location / {
    proxy_pass https:
  }
}
server {
  listen 443 ssl http2;
  server_name photos.*;
 
  ssl_certificate /config/user-data/ssl_chain_key.pem;
  ssl_certificate_key /config/user-data/ssl_chain_key.pem;
  client_max_body_size 0;
  proxy_http_version 1.1;
  proxy_buffering off;
  proxy_set_header Upgrade $http_upgrade;
  proxy_set_header Connection "Upgrade";
  proxy_set_header Host $host;
  proxy_set_header X-Real-IP $remote_addr;
  proxy_set_header X-Forward-For $proxy_add_x_forwarded_for;
  
  rewrite ^/photo/(.*)$ /$1;
  
  location / {
    proxy_pass https:
  }
}
Once the reverse proxy config file is in place you can run sudo service nginx restart to make it effective.  Refer to /var/log/nginx/error.log if there are problems.
TIP #1: If you plan on using a non-standard port for your nginx reverse proxy (ie. the config for your reverse proxy says something other than listen 80 and/or listen 443) then you'd probably be well served to replace proxy_set_header Host $host; in the above config with proxy_set_header Host $http_host;.  This is especially important for websockets that seem to always want to make the web clients request traffic from the default ports otherwise - leading to interfaces that are missing components or live data.  (You can also try proxy_set_header Host $host:$server_port;, but note that this is distinctly different as it will always add a port to the header even when one wasn't used in the original request)
TIP #2: Also note the use of server_name router.* which simply means that using the URL router.<anything>.<anything> to browse the nginx reverse proxy will cause that section of the config to be applied.  (Most examples you'll find use a full/explicit domain name like router.domain.com)
This is useful not only for the purpose of potentially changing domain names/dns names (ie. works for router.home.com as well as router.home.net without changing the config), but also helps keep your server_name entries short.  This is important because having long names will often lead to errors saying something to the effect of "increase your bucket size" when you start nginx.  If you must have long names then you'll likely need to tweak the values of server_names_hash_max_size and/or server_names_hash_bucket_size within your config.
UPDATE:
I have successfully updated my router from EdgeOS version 1.9.1 to every version between and including 2.0.9-hotfix2 with essentially zero issues.  The router did not have nginx on it after the update (as expected), but I was able to run my script (the first code block in this answer) and it came right back.  Keeping your customizations under /config and using scripts under /config/scripts/post-config.d to maintain any changes that don't survive a reboot/upgrade is the key to smooth sailing here.
UPDATE #2:
This process works on smaller routers like the ER-X as well.  However, due to limited storage available on the device you'll be forced to remove the inactive firmware image to make room for nginx.  This needs to be done after each update as well.  To remove the inactive firmware image, run delete system image and answer yes to the prompt(s).
-- My ER-X update process --
- Backup my config!!!!
- Update to the latest firmware and reboot (and verify basic functionality other than nginx)
- Remove the old firmware (which is now the inactive image) using the delete system image command (show system storage reports 67% in use before removal and 33% in use afterwards)
- Verify external pings work (ping 8.8.8.8) and reboot the router if they do not.  This is a known issue in recent firmware versions and has nothing to do with this process or the scripts used.
- Use the scripts/steps above to re-install nginx (show system storage reports 67% in use after installation)
- Run sudo apt-get clean to free up space used by temp files and such (show system storage reports 52% in use after cleanup)
