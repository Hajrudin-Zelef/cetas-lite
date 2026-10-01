---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/posts-opnsense-with-haproxy-and-lets-encrypt-b7772209-2
title: "local access only subdomains"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/posts-opnsense-with-haproxy-and-lets-encrypt-b7772209.md
source_anchor: ""
source_lines: [110, 230]
sha256: 0dbb3574400a428d45ef48fd9422c52b9ed3331d55b5d6e55f06a7e863bab791
---

# local access only subdomains

Check what happened at: Services --> ACME Client --> Log Files --> ACME Log
There are two tabs, System and ACME Log.
- System will display what OPNsense decided to do.
- ACME Log is the script that runs. You can find most errors here.
- If everything went OK, you should have Installed a key to: /var/etc/acme-client/keys/99gbijk0z22.6749280/private.key
What if I used the Testing CA?
Go back to: Services --> ACME Client --> Accounts
Since you, presumably, successfully issued a staging certificate you can now change from the test environment to the default ACME CA Let's Encrypt and issue the production certificate.
Issue a production certificate
Back to: Services --> ACME Client --> Certificates
Now, again, you dont need to do this if you’ve already used the production server to generate a certificate.
Otherwise, forcefully issue your certificate with the circular arrow button
This time it you should receive a valid and trusted SSL certificate. Make sure to check the ACME log for any errors though!
HAProxy configuration
This is where the convoluted part begins. The instructions sound similar and become more difficult to understand. There are several sections, including using a map file, rules, conditions, backend server pools, services, dns overrides and more.
Please re-read if needed and be patient.
Please note, a lot of these pages are located as dropdowns in the menu headers on the settings page.
HAProxy inital configuration
HAProxy Service
First, we need to set the defaults we’re going to use across the HAProxy service in:
Services --> HAProxy --> Settings --> Service
On this page,
- Uncheck Show introduction pages
- Do checkbox Store OCSP responses
Enable HAProxy
Let’s turn it on: Services --> HAProxy --> Settings --> Service
On this page just checkbox Enable HAProxy and hit `Apply.
- Then, Apply.
HAProxy Global Paramaters
The next drop down we need is: Services --> HAProxy --> Settings --> Global Parameters
- In the upper left hand corner, check the advanced mode button.
- The number of HAProxy threads should not exceed the number of CPU threads of your OPNsense (example, 4).
- Maximum connections keeps from overloading the server, say800
- Increase the Maximum SSL DH Size to4096
- Apply.
HAProxy Default Paramaters
This area will helps us temper our reverse proxy to ensure there isnt an overloaded server.
Enter: Services --> HAProxy --> Settings --> Default Parameters
- Maximum Connections (Public Services) set to1200
- Maximum Connections (Servers) set to8000
- You can leave everything else the same.
- Apply.
HAProxy background services configuration
Setting conditions for non-HTTPS connections
Here we only need to create a NoSSL_condition, which is necessary in order to identify non-HTTPS traffic. This will be our only condition.
Next, go to: Services --> HAProxy --> Settings --> Rules & Checks --> Conditions
- Name should be set toNoSSL_condition .
- Condition type needs to be configured toTraffic is SSL (locally deciphered) .
- Negate condition has to be checked for this rule to make sense.
- Apply.
Map file for ease and convenience
Map files are found in: Services --> HAProxy --> Settings --> Advanced --> Map Files
Here we will create a new map file PUBLIC_SUBDOMAINS_map file for our public subdomains that we want to access from outside of our network.
This map file is telling HAProxy that any FQDN that starts with nextcloud then belongs to our NEXTCLOUD_backend (which finds our NEXTCLOUD_server”).
- Name isPUBLIC_SUBDOMAINS_mapfile
- Content for this example, should look like:1 2 # comment to tell you comments are allowed nextcloud NEXTCLOUD_backend
- Apply.
Rules to connect the moving parts
Finally, go to : Services --> HAProxy --> Settings --> Rules & Checks --> Rules
Here we add the rules that decide what to do with the traffic based on our map files (or conditions if necessary).
Forwarding HTTP to HTTPS
First, we must create a HTTPtoHTTPS_rule that will forward all HTTP traffic to HTTPS port, so it can go to our HTTPS_frontend.
- Name isHTTPtoHTTPS_rule .
- Select conditions this is set to the condition made earlier, theNoSSL_condition .
- Execute function are built into HAProxy, the one we want ishttp-request redirect .
- HTTP Redirect this is part of the paramateres for the function we chose earlier. Copy & Paste:1 scheme https code 301
- Save.
Mapping the map file to PUBLIC_SUBDOMAINS_rule
The PUBLIC_SUBDOMAINS_rule maps our subdomains to our backends using the map file we created in the previous step.
Continuing in Rules. Make a new rule for mapping domains to backends using a map file.
- Name isPUBLIC_SUBDOMAINS_rule
- Execute function must be set toMap domains to backend pools using a map file .
- Map file must also be set, the map file made above was namedPUBLIC_SUBDOMAINS_mapfile .
- Save.
- Apply.
Creating HAProxy Reverse Proxy
This is the meat and potatos of the tutorial. You will either get satiated or sick.
Adding the application service backend
First go to: Services --> HAProxy --> Settings --> Real Servers
This is where all of the servers on your network live. These are the “real servers” that exist on your network to which HAProxy can communicate with inorder to faciliate a reverse proxy. Note: unless you’re encrypting from the server to HAProxy with an internal CA, you dont need to verify SSL certificates.
The Virtual IP SSL_server
- Name addSSL_server for our virtual IP address we set earlier.
- IP will be the address of the virtual IP127.1.2.3
- Uncheck Verify SSL Certificate
- Save.
- Apply.
Adding application service servers you’re hosting
- Name add the name of your service for reference here, example:NEXTCLOUD_server
- IP enter the IP address of the machine you want to reverse proxy.
- Port enter the port number for the service(s) you want to proxy.
- Uncheck Verify SSL Certificate
- Save.
- Apply.
Loadbalance SSL_backend pool
Next go to: Services --> HAProxy --> Settings --> Virtual Services --> Backend Pools
These are servers that lie in the backend pool. The backend pool cares for health monitoring and load distribution. If you had multiple servers sending out the same content, they’d reside in the same pool group.
A Backend Pool must be configured, even if you only have one server.
Next, we will create the SSL_backend. This is the backend to which the SNI_frontend sends most of its traffic to.
Adding backend pools
SSL_backend
- In the upper left hand corner of the edit page, find advanced mode and turn it on (green).
- Name useSSL_backend for the first pool.
- Mode is set toTCP (Layer 4) for the SSL_backend.
- Proxy Protocol needs to be configured forVersion 2 .
- Servers should be set toSSL_server made earlier, the one you made with the virtual IP127.1.2.3 .
- Save.
Make sure that “SSL_backend” above is set to TCP mode, since the “SNI_frontend” is also running in TCP mode and you can’t mix HTTP mode and TCP mode with a frontend to backend.
Adding application service backends
Now we create the backend that belongs to an actual service.
- You will need one backend for each service.
- Make SURE the backend is named the same as the one in your mapfile!
- Name should reference the service you’re adding, example:NEXTCLOUD_backend
- Mode needs to beHTTP (Layer 7) [default] , this is different from the SSL_backend.
- Servers set to the correspondingReal Server you made earlier.
- Save.
- Apply.
Note: If you have multiple servers serving the exact same content than you will want to add all servers into a single backend so HAProxy can actually balance the load between the servers.
Life without a map file
Setting an application’s condition to a backend pool… without using a map file… if you so desire
This step is not required if you used the map file, this is an optional step!
If you didnt use a map file, or had a more particular configuration setup than “starts with”, feel free to set rules to point a condition to a backend pool that you made.
