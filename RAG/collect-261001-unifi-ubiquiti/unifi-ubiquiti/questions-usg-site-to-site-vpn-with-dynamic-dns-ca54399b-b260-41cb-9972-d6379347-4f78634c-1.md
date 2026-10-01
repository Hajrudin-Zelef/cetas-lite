---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-usg-site-to-site-vpn-with-dynamic-dns-ca54399b-b260-41cb-9972-d6379347-4f78634c-1
title: "questions-usg-site-to-site-vpn-with-dynamic-dns-ca54399b-b260-41cb-9972-d6379347-4f78634c"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-usg-site-to-site-vpn-with-dynamic-dns-ca54399b-b260-41cb-9972-d6379347-4f78634c.md
source_anchor: ""
source_lines: [1, 80]
sha256: 30abb045e09e8cd3eb54d3c2aa4782e6ec9d4f809cb02cd25a1aaa1fc9ec4ce0
---

# questions-usg-site-to-site-vpn-with-dynamic-dns-ca54399b-b260-41cb-9972-d6379347-4f78634c

This is very much desired. Especially for s2s with 2x USG Pro.
Comment by Stoogie/marting does not work. The current Unifi Controller version (5.11.50-12745-1) requires using IP addresses and not hostnames both on the IPsec-manual and Openvpn options. The only other is auto-ipsec, which DOES NOT fill all use cases as per:
See also the following, which mentions that the controller handles this magically, if both endpoint devices are on different sites of the same controller. I have not tried it yet:
support tells me that you can create a create a custom JSON file in order to do this, which begs the question "Why not just make it possible in the gui?" Has anyone done this and got it to work, they sent me instruction but it is not a simple thing to do.
@Mopar_Mudder wrote:
Please give us a update when finding a solution!
I haven't even tried, was hoping someone else had a copy of the file I could use.
maybe this link might help?
https://blog.azureinfra.com/2018/12/31/usg-vpns-and-dynamic-ips/
So I managed to get this working with a little bit of scripting and some custom config to run the script via config.gateway.json every 5 minutes.
To get this working, ssh to your usg device and create the file /config/scripts/reconfigvpn.sh with the following content and chmod 755 the file to make it executable:
#!/bin/vbash
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
declare -a MyStaticTunnelEndpoints=("1.2.3.4" "5.6.7.8" "9.10.11.12")
MyIP=$(curl -s https://ifconfig.me)
WR="/opt/vyatta/sbin/vyatta-cfg-cmd-wrapper"
for ip in "${MyStaticTunnelEndpoints[@]}"; do
    CurrentIPConfig=$($WR show vpn ipsec site-to-site peer $ip local-address | cut -d' ' -f2)
    if [[ "$MyIP" != "$CurrentIPConfig" ]]; then
        $WR begin
        $WR set vpn ipsec site-to-site peer $ip local-address $MyIP
        logger "updated peer $ip with $MyIP"
        $WR save
        $WR commit
        $WR end
        sleep 10
        logger "commited and saved peer $ip with $MyIP"
    else
        logger "No local IP address change detected for peer $ip"
    fi
done
All the script does is take an array (MyStaticTunnelEndpoints) of static IP addresses (which should already be in your configuration) and then it fetches your current public IP address via a webservice (ifconfig.me - choose your favorite service if you want...)
It then proceeds to loop over your static tunnel-endpoints, comparing the current public IP (MyIP) with the one currently written down in the configuration (CurrentIPConfig). If they match, nothing will be done (except a log-entry into /var/log/messages) - if they don't match the script will re-write the configuration for each endpoint, save and commit this (and then wait for 10 sec. so we don't get problems with - previous - unfinished save/commit operations) and finally write some log-lines into /var/log/messages indicating what was done.
As soon as the configuration is changed, the system will pick it up and re-start the IPSec tunnels with the correct current IP - thus allowing to establish tunnels even if one side should have a dynamic IP.
To schedule the script to run every five minutes, put this into your config.gateway.json (on your cloud-key controller - look in /usr/lib/unifi/data/sites/default if you are on ubuntu):
{
      "system": {
           "task-scheduler": {
                "task": {
                     "reconfigvpn": {
                          "executable": {
                               "path": "/config/scripts/reconfigvpn.sh"
                          },
                          "interval": "5m"
                     }
                }
           }
      }
}
Now if you want to trigger the script, issue a reconnect (if you are using pppoeX for example - change X to your interface numbering):
disconnect interface pppoe2 && sleep 10 && connect interface pppoe2
On a different terminal, watch what's happening on your machine via
tail -f /var/log/messages
Ideally you should see some output every 5 minutes in there relating to the above logger lines - depending which path of the code you took in your current state.
This solution *might* work without even using dyndns - but then the USG should be the one to initiate the connection for the tunnel (aka. if the other side is a pfsense-box, you can configure that to only respond to incoming vpn connections - also you might need to tweek phase1-identifiers to get this to work)
I currently use the USG dyndns service and I configured the other side of the tunnel (a pfsense box) to use the dynamic hostname to reconnect to the USG. Works like a charm so far - manually triggered fail-overs do re-establish the tunnel after 5 minutes (max).
Hope that information helps someone to achieve the same.
Can see where that would work but why have it running a script every 5 minutes and have the possibility of a down tunnel in the mean time? If you can some how just put in the host name instead of the actual IP address nothing ever has to run. I admit I am not real up on how all this works but it seems like it should be simple to fix.
Well, sometimes the simple things are the hardest, right?
And sometimes you have to work with the tools you are given - in your case, you could reduce the interval of the above script down to a minute if you wanted to.
Also: how big of a problem is this really for you? Let's say your ISP disconnects you every 24h - could you not schedule that to a time where your systems are not heavily frequented? Aka. during night hours?
I could see the problem in a 24/7 uptime required production environment (where you would have static IP's though, right?!?) - but wouldn't you then use different gear? something more... professional? and less semi-professional?
Running Sonicwall at office and in the past my ZyXel, Netgear devices have all been able to connect to office Sonicwall where remote sites didn't have option of static IP. Using xxxxxx.dyndns.org instead of an IP address has worked in all of them, now moving to UNIFI Dream Machine Pro and big roadblock not having this capability. Disappointed it's not an option yet. I've had this in these other products for 10 years now.
@mrmmdmaddad wrote:
Yep I never even though of researching if the Unifi had this capability because I thought it was a standard deal seeing as all my Sonic Walls and Ciscos had had it. Had I known I would have chosen something else.
If you want to connect 2 sites with USG, then site to site vpn with ddyns should be possible in this way:
set-inform Http://<IP of controller>:8080/inform
on the USG that should be connected to your remote site. Here you should be able to use ddyns as IP.
Now go to your remote location and do following:
Set a portforwarding on 8080 rule to your controller.
After that is done, create a new site which you will adopt your usg device in.
Adopt your USG in the new site.
After this you can simply choose a site to site vpn by 1 click by choosing the this site.
One problem with this might be that, if the remote controller goes down. So might your VPN.
Note! I haven't done this yet myself. But in theory it should work. Saw a video here that explains it a bit better: https://www.youtube.com/watch?v=ndgHd6TGRJs
BR
Henrik Lindgren @ piarelations
I ended up paying the extra $10 a month to have a static IP. Still can't believe they don't change it so you can enter a IP Name instead of number, can't be that difficult to do.
If you do not have a static ip, you can put 0.0.0.0 as local IP - it works for me with a dream machine (should work at the USG too). from home to my Office (a Sophos XG connects the office) The dm can not answer but initiate the connection..
The IP that you specify there is just the phase 1 identifer (Gateway ID for tunnel authentication), 0.0.0.0 is a good idea and works for me as well connecting to a WatchGuard.
