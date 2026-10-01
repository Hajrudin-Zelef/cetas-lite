---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-ipsec-site-to-site-vpn-w-usg-3p-and-edgerouter-x-90eb13af-e730-41b7-94-b29516ef
title: "questions-ipsec-site-to-site-vpn-w-usg-3p-and-edgerouter-x-90eb13af-e730-41b7-94-b29516ef"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-ipsec-site-to-site-vpn-w-usg-3p-and-edgerouter-x-90eb13af-e730-41b7-94-b29516ef.md
source_anchor: ""
source_lines: [1, 126]
sha256: 87e6121ce05341b43919a8eeb9929c26702c220769f27465d64c27f8b4967754
---

# questions-ipsec-site-to-site-vpn-w-usg-3p-and-edgerouter-x-90eb13af-e730-41b7-94-b29516ef

@UI-Team
Hi all.
I've been tearing my hair out all night trying to get a site-to-site VPN set up between my flat (USG, 192.168.25.0/24) and my parents' house (ER-X, 10.0.0.0/24).
I've got it to a point wherein the two will connect happily every time but the USG seems to be having some kind of issue.
Packets will go out of the EdgeRouter but that's it - they're not picked up by the USG (as shown by show vpn ipsec sa).
show vpn ipsec sa ⬇️
From the USG:
peer-ERX-tunnel-vti: #2, ESTABLISHED, IKEv1, 13aab9023b0d8d7b:a80698b380e952df
  local  'USG' @ USG
  remote 'ER-X' @ ER-X
  AES_CBC-256/HMAC_SHA1_96/PRF_HMAC_SHA1/MODP_2048
  established 147s ago, reauth in 28085s
  peer-ERX-tunnel-vti: #2, INSTALLED, TUNNEL-in-UDP, ESP:AES_CBC-256/HMAC_SHA1_96/MODP_2048
    installed 146 ago, rekeying in 2525s, expires in 3455s
    in  ce0a6f8c,      0 bytes,     0 packets
    out c496b2cb,      0 bytes,     0 packets
    local  192.168.25.0/24
    remote 10.0.0.0/24
From the ER-X:
peer-USG-tunnel-1: #5, ESTABLISHED, IKEv1, 13aab9023b0d8d7b_i* a80698b380e952df_r
  local  'ER-X' @ ER-X[4500]
  remote 'USG' @ USG[4500]
  AES_CBC-256/HMAC_SHA1_96/PRF_HMAC_SHA1/MODP_2048
  established 218s ago, reauth in 27951s
  peer-USG-tunnel-1: #7, reqid 1, INSTALLED, TUNNEL-in-UDP, ESP:AES_CBC-256/HMAC_SHA1_96/MODP_2048
    installed 218s ago, rekeying in 2336s, expires in 3384s
    in  c496b2cb,      0 bytes,     0 packets
    out ce0a6f8c,   2856 bytes,    34 packets,   151s ago
    local  10.0.0.0/24
    remote 192.168.25.0/24
show vpn
 ipsec {
     auto-firewall-nat-exclude enable
     esp-group ESP_ER-X {
         compression disable
         lifetime 3600
         mode tunnel
         pfs enable
         proposal 1 {
             encryption aes256
             hash sha1
         }
     }
     ike-group IKE_ER-X {
         key-exchange ikev1
         lifetime 28800
         proposal 1 {
             dh-group 14
             encryption aes256
             hash sha1
         }
     }
     ipsec-interfaces {
         interface eth0
     }
     nat-networks {
         allowed-network 0.0.0.0/0 {
         }
     }
     nat-traversal enable
     site-to-site {
         peer ER-X {
             authentication {
                 mode pre-shared-secret
                 pre-shared-secret <secret>
             }
             connection-type initiate
             ike-group IKE_ER-X
             local-address USG
             vti {
                 bind vti64
                 esp-group ESP_ER-X
             }
         }
     }
 }
 ipsec {
     auto-firewall-nat-exclude enable
     esp-group FOO0 {
         proposal 1 {
             encryption aes256
             hash sha1
         }
     }
     ike-group FOO0 {
         proposal 1 {
             dh-group 14
             encryption aes256
             hash sha1
         }
     }
     site-to-site {
         peer USG {
             authentication {
                 mode pre-shared-secret
                 pre-shared-secret <secret>
             }
             connection-type initiate
             description "My Flat"
             force-encapsulation enable
             ike-group FOO0
             local-address ER-X
             tunnel 1 {
                 esp-group FOO0
                 local {
                     prefix 10.0.0.0/24
                 }
                 remote {
                     prefix 192.168.25.0/24
                 }
             }
         }
     }
 }
I can't help but think there is one heck of a difference between the two now I look...!
Is there something I'm missing? I tried disabling the remote-user VPN on the USG to no avail. I've tried forcing encapsulation - also nothing.
Please send help! 😄
You have a route-based VPN (VTI) configured on the USG side and a policy-based VPN configured on the ER-X side. I've never done it, but I think you should be able to make it work by configuring the ER-X to use VTI and mirror the USG config. All the parameters should match: pfs, lifetimes, compression, etc. Review this:
https://help.ui.com/hc/en-us/articles/115011377588-EdgeRouter-Route-Based-Site-to-Site-IPsec-VPN
@stshaw wrote:
If we weren’t in the middle of a pandemic, I would give you a massive kiss right now as that worked beautifully!!!
I cannot thank you enough!
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
