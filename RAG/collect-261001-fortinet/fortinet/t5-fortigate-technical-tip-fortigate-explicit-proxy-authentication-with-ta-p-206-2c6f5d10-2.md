---
id: collect-261001-fortinet/fortinet/t5-fortigate-technical-tip-fortigate-explicit-proxy-authentication-with-ta-p-206-2c6f5d10-2
title: "t5-fortigate-technical-tip-fortigate-explicit-proxy-authentication-with-ta-p-206-2c6f5d10"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/t5-fortigate-technical-tip-fortigate-explicit-proxy-authentication-with-ta-p-206-2c6f5d10.md
source_anchor: ""
source_lines: [289, 457]
sha256: 95c7fb783c50f86ea504c31de116ed4863910dcb3cc6517166cc4e54d5809487
---

# t5-fortigate-technical-tip-fortigate-explicit-proxy-authentication-with-ta-p-206-2c6f5d10

                set group-name "CN=remoteAdmins,OU=France,OU=Europe,DC=amf,DC=meta,DC=local"

            next

            edit 2

                set server-name "2AD"

                set group-name "CN=Domain Users,CN=Users,DC=amf,DC=meta,DC=local"

           next

        end

    next

end


1. Create an **Authentication Scheme** and an**Authentication Rule** .

Configuration through the GUI.


Scheme configuration:




Rule configuration:




Configuration through CLI.


**config authentication scheme**

    edit "KRB2"

        set method negotiate

        set negotiate-ntlm disable

        set kerberos-keytab "service_fortigate2"

    next

end

**config authentication rule**

    edit "Rule_KRB2"

        set srcaddr "all"

        set ip-based disable

        set active-auth-method "KRB2"

    next

end


1. Create an explicit proxy-policy.


Configuration through the GUI.


Add the 'KRB' group created previously in the source.




Configuration through the CLI:


**config firewall proxy-policy**

    edit 1

        set name "KRB2_policy"

        set proxy explicit-web

        set dstintf "port1"

        set srcaddr "all"

        set dstaddr "all"

        set service "webproxy"

        set schedule "always"

        set logtraffic disable

        set groups "KRB"

    next

end


Windows proxy settings and monitoring. Testing the scenario with a non-domain-joined Windows PC.


Domain user: srogers.

Go to **Control Panel -> Internet Options -> Connections.**


Enter the FortiGate FQDN/IP as a proxy server in LAN settings and modify the port to 8080.




The captive portal prompt can be prevented, eliminating the need for the user to re-enter login credentials. To do this, open Internet Properties by typing 'inetcpl.cpl' in the Windows command prompt, then go to **Security -> Local Intranet -> Sites -> Advanced**. Add the FortiGate FQDN/IP address, including the HTTP/HTTPS prefix.


For example, http://fortigate2K.domain.com or https://fortigate2K.domain.com.


It is possible to verify user authentication in the FortiGate CLI. For this, run '**diagnose debug enable**' and then the command below:




In **Log & Report -> Events -> User events**, it is possible to monitor the user and authentication data.




On the Windows host, it is possible to check the granted tickets in CMD through the 'klist' command:




To capture the Request service packet and server Auth packet reply from the Server to the firewall, use the below WAD debug:


**diagnose wad filter clear**

**diagnose wad debug clear**

**diagnose wad debug display pid enable****diagnose wad filter src <IP address> / diagnose wad filter process-id-by-src <IP_address_of_client>**

**diagnose wad debug enable cate http****diagnose wad debug enable cate auth****diagnose wad debug enable cate config****diagnose debug cli 8**

**diagnose deb app fnbamd -1**

**diagnose wad debug enable level verb**

**diagnose debug console timestamp enable**

**diagnose debug enable**

**diagnose debug disable** -> To stop.


If the above WAD debug displays no output, it is recommended to use debugging without any source filter.**Note:**

The **wad filter process-id-by-src** option was introduced in **FortiProxy v7.0.13, v7.2.7, and v7.4.1**, and it was added in FortiGate starting version **7.6.3**. This option notably enhances the WAD filtering process by first filtering by the Source IP of a client connection and then further filtering based on the specific WAD worker process that is handling that traffic (which helps to reduce excess debug output).


**Related articles:**
