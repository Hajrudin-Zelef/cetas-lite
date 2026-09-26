---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-27
title: "Properties"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-27.md
source_anchor: ""
source_lines: [1, 62]
sha256: 150e362546088e00935920c80dd425dbc30b86d6c041808f6082d59f394241ff
---

# Properties

An E-mail tool is a utility that allows sending e-mails from the router. The tool can be used to send regular configuration backups and exports to a network administrator.

Email tool uses only plain authentication and TLS encryption. Other methods are not supported.

# Properties

This submenu allows setting SMTP server that will be used.

| Property | Description | 
|---|---|
| **address** (*IP/IPv6 address* ; Default:**0.0.0.0** ) | SMTP server's IP address. | 
| **certificate-verification**  (*yes* \|*yes-without-crl* \|*no* ; Default:**no** ) | Enables trust chain validation from local certificate store. *yes-without-crl,* validates a certificate, not performing CRL check (certificate revocation list). | 
| **from** (*string* ; Default:**<>** ) | Name or email address that will be shown as a receiver. | 
| **password** (*string* ; Default:**""** )*sensitive* | Password used for authenticating to an SMTP server. | 
| **port** (*integer[0..65535]* ; Default:**25** ) | SMTP server's port. | 
| **tls** (*no\|yes\|starttls* ; Default:**no** ) | Whether to use TLS encryption:  | 
| **user** (*string* ; Default:**""** ) | The username used for authenticating to an SMTP server. | 
| **vrf** (*VRF name* ; default value:**main** ) | Set VRF on which service is creating outgoing connections. | 

**Note:**All server's configurations (if specified) can be overridden by send command.

# Sending Email

Send command takes the following parameters:

| Property | Description | 
|---|---|
| **body** (*string* ; Default: ) | The actual body of the email message | 
| **cc** (*string* ; Default: ) | Send a copy to listed recipients. Multiple addresses allowed, use "," to separate entries | 
| **certificate-verification**  (*yes* \|*yes-without-crl* \|*no* ; Default:**no** ) | Enables trust chain validation from local certificate store. *yes-without-crl,* validates a certificate, not performing CRL check (certificate revocation list). | 
| **file** (*File[,File]* ; Default: ) | List of the file names that will be attached to the mail separated by a comma. | 
| **from** (*string* ; Default: ) | Name or email address which will appear as the sender. If a not specified value from the server's configuration is used. | 
| **password** (*string* ; Default: )*sensitive* | Password used to authenticate to an SMTP server. If a not specified value from the server's configuration is used. | 
| **port** (*integer[0..65535]* ; Default: ) | Port of SMTP server. If not specified, a value from the server's configuration is used. | 
| **server** (*IP/IPv6 address* ; Default: ) | Ip or IPv6 address of SMTP server. If not specified, a value from the server's configuration is used. | 
| **tls** (*yes\|no\|starttls* ; Default:**no** ) | Whether to use TLS encryption:  | 
| **subject** (*string* ; Default: ) | The subject of the message. | 
| **to** (*string* ; Default: ) | Destination email address. Single address allowed. | 
| **user** (*string* ; Default: ) | The username used to authenticate to an SMTP server. If not specified, a value from the server's configuration is used. | 

# Basic examples

**This example will show how to send an email with configuration export every 24hours.**

1. Configure SMTP server

2. Add a new script named "export-send":

3. Add scheduler to run our script:

**Send e-mail to a server using TLS/SSL encryption. For example, Google mail requires that.**

After the Google mail added **a new security policy** that **does not allow 3d-party devices to authenticate** **using** your standard **Gmail** **password** → you need to generate a 16-digit passcode ("App password") and use it instead of your Gmail password. To configure this, navigate to the "**Security>How you sign in to Google**" section settings and:

- **Enable 2-Step Verification;**
- **Generate an App password** .

Use the newly generated App password in the "set password=**mypassword**" setting shown below.

1. configure a client to connect to the correct server:

2. send e-mail using send command:
