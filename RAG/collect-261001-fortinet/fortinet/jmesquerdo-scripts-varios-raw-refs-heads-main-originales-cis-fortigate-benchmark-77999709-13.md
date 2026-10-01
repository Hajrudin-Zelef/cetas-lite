---
id: collect-261001-fortinet/fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709-13
title: "jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["asic", "lean"]
source: docs/RAG/collect-261001-fortinet/jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709.md
source_anchor: ""
source_lines: [1983, 2211]
sha256: 1513212f3b6989b32e4e9d5eb91adb2ecc40f5cce513eb2a5b38204ebd52aed1
---

# jmesquerdo-scripts-varios-raw-refs-heads-main-originales-cis-fortigate-benchmark-77999709

Page 71 
FGT1 #config system ha 
FGT1 (ha) # show 
config system ha 
    set group-name "FGT-HA" 
    set mode a-p 
    set password ENC 
enrwD467hJmO6j6YW/l6FEOa1YNVYdo8Z5mCcTDEKUFpOVXcNYnPBmQDGX//ViXk6TkwNH0il5aJr
/fZY25lq+husndQHZVWp2LIlXmCv/n81U43nkZUWaIKvqkellGFbhv0/IHoOLzQPCsVcBbyrsgopr
YMvh6w7F06+nRriBtMNQxpiTE+12xAHz7lA3EoYZzf8A== 
    set ha-mgmt-status enable   
    config ha-mgmt-interfaces 
        edit 1 
            set interface "port6" 
            set gateway 10.10.10.1 
        next 
    end 
    set override disable 
end 
Validate that set ha-mgmt-status is enable 
and that config ha-mgmt-interfaces has at least one entry with an interface and gateway 
defined 
Remediation: 
Remediate through the GUI: 
go to System -> HA edit the "Master" device and enable "Management Interface 
Reservation" once this is enabled select an an interface, and configure the 
appropriate gateway. 
Remediate through the CLI:

Page 72 
FGT1 #config system ha 
FGT1 (ha) # set ha-mgmt-status enable  
FGT1 (ha) # config ha-mgmt-interfaces  
FGT1 (ha-mgmt-interfaces) # edit 1 
new entry '1' added 
FGT1 (1) # set interface port6 
FGT1 (1) # set gateway 10.10.10.1 
FGT1 (1) # end 
FGT1 (ha) # show 
config system ha 
    set group-name "FGT-HA" 
    set mode a-p 
    set password ENC 
enrwD467hJmO6j6YW/l6FEOa1YNVYdo8Z5mCcTDEKUFpOVXcNYnPBmQDGX//ViXk6TkwNH0il5aJr
/fZY25lq+husndQHZVWp2LIlXmCv/n81U43nkZUWaIKvqkellGFbhv0/IHoOLzQPCsVcBbyrsgopr
YMvh6w7F06+nRriBtMNQxpiTE+12xAHz7lA3EoYZzf8A== 
    set ha-mgmt-status enable   
    config ha-mgmt-interfaces 
        edit 1 
            set interface "port6" 
            set gateway 10.10.10.1 
        next 
    end 
    set override disable 
end 
FGT1 (ha) # end 
Default Value: 
N/A

Page 73 
3 Policy and Objects 
This section contains best practices related to configuring firewall policies, Objects and 
traffic shaping

Page 74 
3.1 Ensure that unused policies are reviewed regularly (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
All firewall policies should be reviewed regularly to verify the business purpose. Unused 
policies should be disabled and logged. 
Rationale: 
By reviewing policies regularly, we can determine if the policies are still needed by the 
business purpose. Thus, we can keep the firewall policies lean and efficient. It also 
prevents traffic being allowed or blocked accidently. 
Audit: 
In CLI, type "diag firewall iprope show 100004 <policy_id>". In this example, we'll verify 
policy with ID of 32. We'll also need to clear the counter after each review so that we 
can tell if the policy is still being used for the next review : 
FGT1 # diag firewall iprope show 100004 32 
idx=2 pkts/bytes=144967/135758174 asic_pkts/asic_bytes=0/0 flag=0x0 hit 
count:663 
FGT1 # diag firewall iprope clear 100004 32 
In the GUI, 
go to Policy & Objects -> IPv4 Policy. First make sure that either the 
columns "Bytes" or "Hit Count" are visible. To display either one of them, 
move the cursor to the top row where all the columns names are. Right click 
and select "Bytes" or "Hit Count" and click OK. To clear the counter, right 
click on the "Bytes" or "Hit Count" columns of that policy and click on 
"Clear Counters". 
Remediation: 
The remediation is to review and decide if you should delete unused policies. 
Default Value: 
By default, the hit count value is obviously 0 at the beginning. 
References: 
1. https://kb.fortinet.com/kb/documentLink.do?externalID=FD44631 
Additional Information: 
The CLI commands are only available after FortiOS 6.0. Before that, please use GUI.

Page 75 
3.2 Ensure that policies do not use "ALL" as Service (Automated) 
Profile Applicability: 
•  Level 1 
Description: 
We want to make sure that all security policies in effect clearly state which protocols / 
services they are allowing. 
Rationale: 
This is to make sure that the firewall do not allow traffic with unauthorized 
protocols/services by mistakes. 
Audit: 
In CLI: 
FGT1 # config firewall policy 
FGT1 (policy) # show 
TEST-FG-Third (policy) # show 
config firewall policy 
    edit 1 
        set uuid d0eed832-bb73-51e6-c3da-3cd2ec201608 
        set srcintf "internal" 
        set dstintf "wan" 
        set srcaddr "all" 
        set dstaddr "all" 
        set action accept 
        set schedule "always" 
        set service "HTTPS" "HTTP" 
        set ssl-ssh-profile "__tmp_no-inspection" 
        set nat enable 
    next 
end 
In the GUI, 
go to Policy & Objects -> IPv4 Policy. 
Make sure that none of the policies use "ALL" as its service 
Remediation: 
In this example, we will modify policy with ID of 2 to change the service from "ALL" to 
FTP and SNMP 
In CLI:

Page 76 
FGT1 # config firewall policy 
FGT1 (policy) # edit 2 
FGT1 (2) # set service "FTP" "SNMP" 
FGT1 (2) # end 
FGT1 # 
In the GUI, 
click on Policy & Objects -> IPv4 Policy. Select the policy, click "Edit". In 
the Service section, click on it and select FTP and SNMP. Click OK 
Default Value: 
By default, all new policy will have "ALL" in its service field. 
CIS Controls: 
Controls 
Version Control IG 1 IG 2 IG 3 
v7 
9.2 Ensure Only Approved Ports, Protocols and 
Services Are Running 
 Ensure that only network ports, protocols, and services listening on a 
system with validated business needs, are running on each system. 
 ● ●

Page 77 
3.3 Ensure Policies are Uniquely Named (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Ensure Policies are uniquely named 
Rationale: 
Uniquely named policies allow for better auditing and prevent multiple policies 
performing the same actions existing and possibly introducing misconfigurations 
Audit: 
Review firewall policies and validate that all policies are uniquely named 
Remediation: 
Provide all firewall policies with a unique name

Page 78 
3.4 Ensure there are no Unused Policies (Manual) 
Profile Applicability: 
•  Level 2 
Description: 
Ensure that there are no firewall policies that are unused 
Rationale: 
Unused policies may provide unintended or anticipated access to services or hosts 
Audit: 
Review all Firewall policies for use and validate the purpose of the policy 
Remediation: 
Disable and then delete any used firewall policies

Page 79 
3.5 Ensure firewall policy denying all traffic to/from Tor or 
malicious server IP addresses using ISDB (Manual) 
Profile Applicability: 
•  Level 1 
Description: 
Firewall policies should include a deny rule for traffic going to/from Tor or malicious 
server using ISDB (Internet Service Database). 
Rationale: 
FortiGate includes Tor or malicious server related IP address using ISDB. The idea is to 
filter out malicious traffics using firewall policies as first level filtering. This is done 
without involving more resource intensive process such as IPS inspection, hence 
optimizing FortiGate's performance. 
Audit: 
Go to "Policy & Objects". 
Validate that there is a firewall policy created to block inbound connections from 
sources named "Tor-Exit.Node", "Tor-Relay.Node", and "Malicious-Malicious.Server" on 
"All" services. 
Validate that there is a firewall policy created to block outbound connections to 
destination named "Tor-Relay.Node" and "Malicious-Malicious.Server". 
Remediation: 
Review firewall policies and ensure there are: 
1. A firewall policy created to block inbound connections with these settings: 
From: Any 
To: Any 
Source: "Tor-Exit.Node", "Tor-Relay.Node", and "Malicious-Malicious.Server" 
Destination: all 
Schedule: Always 
Services: All 
Action: Deny 
Log Violation Traffic: Enabled 
Enable this policy: Enabled 
 
2. A firewall policy created to block outbound connections with these settings:

Page 80 
From: Any 
To: Any 
Source: All 
Destination: "Tor-Relay.Node" and "Malicious-Malicious.Server" 
Schedule: Always 
Action: Deny 
Log Violation Traffic: Enabled 
Enable this policy: Enabled

