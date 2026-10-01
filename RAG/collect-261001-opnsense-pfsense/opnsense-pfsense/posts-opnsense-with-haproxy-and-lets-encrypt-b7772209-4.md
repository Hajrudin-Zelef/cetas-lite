---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/posts-opnsense-with-haproxy-and-lets-encrypt-b7772209-4
title: "local access only subdomains"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/posts-opnsense-with-haproxy-and-lets-encrypt-b7772209.md
source_anchor: ""
source_lines: [339, 448]
sha256: 4c0812c8a63213759ff78a373db3b49806e331d0d09599962300f69f322313fa
---

# local access only subdomains

I am using the LAN IP for OPNsense, 192.168.1.254 on which the 0_SNI_frontend is also listening on.
- Host set the name to*
- Domain should be the subdomain you registered the Let’s Encrypt certificate for, example: “test.house.lan”.
- Type is anA (IPv4 address)
- IP Address is the IP address of the reverse proxy server for local services, example: 192.168.1.254.
- Save.
Option B - NAT Reflection
This option uses firewall rules to check NAT first.
Please note that “NAT Reflection” is only applicable when port forwarding.
Open: Firewall --> NAT --> Port Forward
Create a new rule with the following:
- Interface is set toWAN
- Destination needs to beWAN address
- Destination port range should be set to the alias,HAProxy_Ports
- Redirect target IP should be one of the Virtual IPs we set earlier that the0_SNI_frontend is listening on.
- NAT reflection must changed to beEnable for this to work.
- Save.
Advanced Configuration: local-access-only subdomains
Imagine you have a service that you would like to access / protect using your brand new reverse proxy without making it available on the internet?
Well, HAProxy has got you covered!
Map file for local only subdomains
Going back to the map files section: Services --> HAProxy --> Settings --> Advanced --> Map Files
- Clone the current, PUBLIC_SUBDOMAINS_mapfile , rename it toLOCAL_SUBDOMAINS_mapfile .
- Add all of the subdomains you want to be local-access-only along with their corresponding backends.
Keep in mind that the contents of your “PUBLIC_SUBDOMAINS_mapfile” must also be in the “LOCAL_SUBDOMAINS_mapfile”.
HAProxy is processing the rules in the frontends based on the order they appear.
So if you place your PUBLIC_SUBDOMAINS_rule before your LOCAL_SUBDOMAINS_rule in the frontend configuration, you won’t get access to your local-access-only subdomains.
Vice versa this will also happen and you will no longer have access to your public subdomains.
To avoid this you have to also put the content of your PUBLIC_SUBDOMAINS_mapfile in the LOCAL_SUBDOMAINS_mapfile and place their rules in the correct order.
LOCAL_SUBDOMAINS_mapfile, example:
1
2
3
4
5
# local access only subdomains
localstuff MYSTUFF_backend
# public access subdomains
nextcloud NEXTCLOUD_backend
Create a local address condition
Conditions are in HAProxy: Services --> HAProxy --> Settings --> Rules & Checks --> Conditions
Now you need to create a new condition that detects if the source of the request is a local IP.
- Name is set set toLOCAL_SUBDOMAINS_SUBNETS_condition
- Condition type should beSource IP matches specified IP
- The Source IP should be as small and specific to your network as possible, or just use the entire RFC1918 IP range:1 10.0.0.0/8 172.16.0.0/12 192.168.0.0/16
You can use the predefined “Source IP is local” condition instead if needed.
Use DNS addresses for the condition
If you really need to, you can use DNS entries to match IPs that should be able to access local services.
But the resolving is only done once during the start / restart of HAProxy.
Adding the LOCAL_SUBDOMAINS_rule
Move over to rules: Services --> HAProxy --> Settings --> Rules & Checks --> Rules
- Clone the current, PUBLIC_SUBDOMAINS_rule , rename it toLOCAL_SUBDOMAINS_rule
- Name isLOCAL_SUBDOMAINS_rule
- Select Conditions and check/select theLOCAL_SUBDOMAINS_SUBNETS_condition you made earlier.
  - If you are using more than one set of SUBDOMAIN_conditions, you will need to change the Logical operator for conditions toOR .
- If you are using more than one set of SUBDOMAIN_conditions, you will need to change the 
- Execute function needs to be set toMap domains to backend pools using a map file
- Map file should be set to the name of our cloned and recreated mapfile,LOCAL_SUBDOMAINS_mapfile
- Save.
Ordering your HTTPS frontend rules
Final step: Services --> HAProxy --> Settings --> Virtual Services --> Public Services
- Edit your 1_HTTPS_frontend
- Scroll to the Select Rules section.
- Make sure you have your LOCAL_SUBDOMAINS_rule first.
  - This needs to be in the upper left corner, as in, first.
 Make sure your LOCAL_SUBDOMAINS_rule comes before your PUBLIC_SUBDOMAINS_rule entry in the rules for your HTTPS frontend public service.
Hide your certificate association with your WAN address
You might have noticed that if you can access your OPNsense using your public WAN IP (https://YOUR_PUBLIC_IP/) the connection will be secured, but upon further inspection you will see that your Let’s Encrypt certificate tied to your domain is beeing used.
While this is not a major security problem it still presents at least some privacy issues.
To fix this we can present a dummy certificate to everyone accessing your reverse proxy directly using your public WAN IP (https://YOUR_PUBLIC_IP/).
Create a placeholder certificate
Create a certificate authority
Create a new CA: System --> Trust --> Authorities
Here we have to create a placeholder certificate authority in order to create the placeholder certificate in the next step.
- Name will beInvalid_SNI
- Method is set toCreate an internal Certificate Authority
- Key Type should beElliptic Curve
- Curve needs to besecp521r1
- Digest Algorithm isSHA512
- Lifetime can be10950
- Fill in the rest of the fields with: Invalid_SNI
- Save.
Create a certificate
Now, the certificate:  System --> Trust --> Certificates
Create the placeholder certificate.
- Method isCreate an internal Certificate
- Descriptive Name isInvalid_SNI
- Certificate authority is alsoInvalid_SNI
- Type is aServer Certificate
- Key Type isElliptic Curve
- Curve issecp521r1
- Digest Algorithm isSHA512
- Lifetime can be10950
- Private key location should beSave on this firewall
- Fill in the rest of the fields with: Invalid_SNI .
- Save.
Connecting your frontend public service to a default placeholder certificate
Finally, get to: Services --> HAProxy --> Settings --> Virtual Services --> Public Services
The last thing left to do is to configure the placeholder certificate as “Default certificate” in your 1_HTTPS_frontend.
- Open your 1_HTTPS_frontend in yourPublic Services
- Scroll to the Default certificate section.
- Choose our newly minted placeholder certificate, Invalid_SNI
You should now no longer get presented with your trusted Let’s Encrypt certificate when accessing “https://YOUR_PUBLIC_IP”, but instead with the “Invalid_SNI” certificate. Thus masking your IP address connected to your domains.
Sources:
- TheMaw Tech for the basic pfsense style HAProxy Setup
- https://www.youtube.com/watch?v=uACQrhtsgFk
- TheHellSite forum post on OPNsense’s forums
- https://forum.opnsense.org/index.php?topic=23339.0
- OPNsense’s documentation
- https://docs.huihoo.com/m0n0wall/opnsense/manual/how-tos/haproxy.html
