---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-c7j6jg-recommended-ssl-cert-ff5e5fc1
title: "r-fortinet-comments-c7j6jg-recommended-ssl-cert-ff5e5fc1"
domain: fortinet
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-c7j6jg-recommended-ssl-cert-ff5e5fc1.md
source_anchor: ""
source_lines: [1, 19]
sha256: 2e49b59208bd96a69f8fad14c9b144d2dc8c79b7ccfa4b2d9f7dc424eceaf5ca
---

# r-fortinet-comments-c7j6jg-recommended-ssl-cert-ff5e5fc1

Recommended SSL cert 
        
    
      We are about to install HA 600E's and will be doing web filtering (including SSL inspection), and we will also use the 600E's for SSL VPN. We will buy certs rather than relying on the default self signed. Most internal devices will be domain joined and so have the cert deployed via GPO with the exception of BYOD devices. 
I'm wondering what level of validation is recommended -DV, OV, or EV?
Given past experience I'd like to avoid the need for intermediate certs so has anyone had experience with a CA/cert that doesn't have a requirement for intermediate certs?
    
Section des commentaires
What you're trying to accomplish is not possible. The certificate you would need to purchase from a vendor will not be usable for SSL inspection. No commercial 3rd party certificate authority is going to provide a certificate that allows you to perform "man-in-the-middle" because that has the potential to be misused by malicious users which in affect will ruin the Certificate Authority's trust chains.
The technical reason is, even though the intermediate/root certificates are "CA:TRUE" certs, the actual end entity certificate does not have the necessary extensions to perform the certificate resigning. The end entity "keyUsage" extension would require the "keyCertSign" attribute to allow this certificate to be used for SSL inspection.
If you need to use a custom certificate then you have to implement your own internal certificate authority. The most common one, to my knowledge, is Windows Certificate Services. You can refer to the link below for more information:
Setting up certificate services to sign the Fortigate SSL proxy cert: http://stuff.purdon.ca/?page_id=163 Fortigate HTTPS deep scanning and invalid certificates: http://stuff.purdon.ca/?page_id=155
EDIT: I must add that if you're hosting a webserver and have a CA cert for it. You can use this cert for inbound inspection to that specific server using " Protect SSL Server" option in your " SSL / SSH Inpsection" profile.
Thanks for the detailed explanation
Cheers -looks like a cert from our AD will be the way to go, previous webfilter had issues with it. And thinking about it, the VPN will be 90% internal devices so all will have the fortinet cert installed via GPO
short version: For accessing the gate, be it admin or VPN, anything you purchase will be just fine. For SSL inspection, you gotta use your own (enterprise) CA, commercial third parties won't/mustn't sell this to you.
As someone else said, the SSL VPN certs are different from the SSL MiTM inspection certs and you cannot used purchased certificates for the SSL MiTM
I use cheap Namecheap.com certs for the SSL VPN. They US$8 per year.
The SSL inspection cert, I distribute the one the FortiGate generates.
