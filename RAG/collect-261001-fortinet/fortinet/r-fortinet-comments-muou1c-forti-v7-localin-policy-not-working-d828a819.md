---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-muou1c-forti-v7-localin-policy-not-working-d828a819
title: "r-fortinet-comments-muou1c-forti-v7-localin-policy-not-working-d828a819"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: ["2015-12-01", "2018-04-09", "2021-04-20", "2021-10-29"]
keywords: ["license"]
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-muou1c-forti-v7-localin-policy-not-working-d828a819.md
source_anchor: ""
source_lines: [1, 210]
sha256: 8df32ba5938456d9931e1c97692313f87d95c601292530627604f88559acbb5d
---

# r-fortinet-comments-muou1c-forti-v7-localin-policy-not-working-d828a819

Forti v7 local-in policy not working 
        
        
        
    
    
    Just playing with new software FortiGate-60E v7.0.0,build0066,210330 and found that local-in-policy is not working anymore. Did anyone notice that already and know what to do?
Like to get it work again
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Can you be more specific?
Which local-in policy isn't working? (show the CLI config of it)
How is it not working? (completely ignored and allowing traffic? Some GUI bug? Some other behaviour?)
Started to get alarms as you see. Then i tested and yes, the fortigate was accessible from everywhere. Temporarily added trust host. Because this fw is for testing i am not worried, but curious, what the new version wants
config firewall local-in-policy edit 1 set intf "untrust" set srcaddr "all" set dstaddr "all" set action accept set service "PING" "HTTP" "HTTPS" "IKE" set schedule "always" next edit 2 set intf "any" set srcaddr "ADMIN_SUBNETS" set dstaddr "all" set action accept set service "ALL" set schedule "always" next edit 3 set intf "untrust" set srcaddr "all" set dstaddr "all" set service "ALL" set schedule "always" next end
Can you run debug flow on that, just to check what it's allowed by.
dia de res
dia de flow filter clear
dia de flow filter addr <destination-ip>
dia de flow filter addr <ssh-port used>
dia de en
dia de flow trace start 10
==> reproduce the issue (or wait for it to happen). Make sure you've captured the "unwanted" traffic, and not just your own allowed SSH. (you can change the address filter to some unwanted client IP that is trying to connect, if the source-IP repeats, or if you're reproducing this yourself, and therefore you know what the source-IP is)
My test results here seem to be effective：
FGVM04TM20007642 # config firewall local-in-policy
FGVM04TM20007642 (local-in-policy) # show
config firewall local-in-policy
edit 1
set intf "port2"
set srcaddr "all"
set dstaddr "all"
set action accept
set service "PING" "HTTP" "HTTPS" "IKE"
set schedule "always"
next
edit 2
set intf "any"
set srcaddr "ADMIN_SUBNETS"
set dstaddr "all"
set action accept
set service "ALL"
set schedule "always"
next
edit 3
set intf "port2"
set srcaddr "all"
set dstaddr "all"
set service "ALL"
set schedule "always"
next
end
FGVM04TM20007642 (local-in-policy) #
FGVM04TM20007642 # diagnose debug flow filter addr 192.168.100.2
FGVM04TM20007642 # diagnose debug flow trace start 100
FGVM04TM20007642 # diagnose debug enable
FGVM04TM20007642 # id=20085 trace_id=36 func=print_pkt_detail line=5723 msg="vd-root:0 received a packet(proto=6, 192.168.100.10:49167->192.168.100.2:22) from port2. flag [S], seq 3160216098, ack 0, win 8192"
id=20085 trace_id=36 func=init_ip_session_common line=5894 msg="allocate a new session-00003758"
id=20085 trace_id=36 func=vf_ip_route_input_common line=2621 msg="find a route: flag=84000000 gw-192.168.100.2 via root"
id=20085 trace_id=36 func=fw_local_in_handler line=455 msg="iprope_in_check() check failed on policy 3, drop"
id=20085 trace_id=37 func=print_pkt_detail line=5723 msg="vd-root:0 received a packet(proto=6, 192.168.100.10:49167->192.168.100.2:22) from port2. flag [S], seq 3160216098, ack 0, win 8192"
id=20085 trace_id=37 func=init_ip_session_common line=5894 msg="allocate a new session-00003759"
id=20085 trace_id=37 func=vf_ip_route_input_common line=2621 msg="find a route: flag=84000000 gw-192.168.100.2 via root"
id=20085 trace_id=37 func=fw_local_in_handler line=455 msg="iprope_in_check() check failed on policy 3, drop"
id=20085 trace_id=38 func=print_pkt_detail line=5723 msg="vd-root:0 received a packet(proto=6, 192.168.100.10:49167->192.168.100.2:22) from port2. flag [S], seq 3160216098, ack 0, win 8192"
id=20085 trace_id=38 func=init_ip_session_common line=5894 msg="allocate a new session-0000375a"
id=20085 trace_id=38 func=vf_ip_route_input_common line=2621 msg="find a route: flag=84000000 gw-192.168.100.2 via root"
id=20085 trace_id=38 func=fw_local_in_handler line=455 msg="iprope_in_check() check failed on policy 3, drop"
FGVM04TM20007642 # get system status
Version: FortiGate-VM64 v7.0.0,build0066,210330 (GA)
Virus-DB: 85.00602(2021-04-20 20:16)
Extended DB: 85.00602(2021-04-20 20:16)
Extreme DB: 1.00000(2018-04-09 18:07)
AV AI/ML Model: 2.00202(2021-04-20 19:45)
IPS-DB: 6.00741(2015-12-01 02:30)
IPS-ETDB: 6.00741(2015-12-01 02:30)
APP-DB: 6.00741(2015-12-01 02:30)
INDUSTRIAL-DB: 6.00741(2015-12-01 02:30)
IPS Malicious URL Database: 2.00984(2021-04-20 04:49)
Serial-Number: FGVM04TM20007642
License Status: Valid
License Expiration Date: 2021-10-29
VM Resources: 1 CPU/4 allowed, 2008 MB RAM
Log hard disk: Available
Hostname: FGVM04TM20007642
Private Encryption: Disable
Operation Mode: NAT
Current virtual domain: root
Max number of virtual domains: 10
Virtual domains status: 1 in NAT mode, 0 in TP mode
Virtual domain configuration: disable
FIPS-CC mode: disable
Current HA mode: standalone
Branch point: 0066
Release Version Information: GA
FortiOS x86-64: Yes
System time: Tue Apr 20 21:29:27 2021
Last reboot reason: warm reboot
FGVM04TM20007642 #
strange. But here it is not working, looks like not matching local-in policies at all.
what‘s the specific content of srcaddr "ADMIN_SUBNETS"？ and is wan1 in a zone untrust?
