---
id: collect-261001-fortinet/fortinet/document-fortigate-6-4-0-hardening-your-fortigate-344487-global-commands-for-str-a7c5b27c
title: "document-fortigate-6-4-0-hardening-your-fortigate-344487-global-commands-for-str-a7c5b27c"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-6-4-0-hardening-your-fortigate-344487-global-commands-for-str-a7c5b27c.md
source_anchor: ""
source_lines: [1, 23]
sha256: f093442aa23a2ac6864ddfc34df232ff76b6fee78f7ff8e5522ea69fdf4e11bd
---

# document-fortigate-6-4-0-hardening-your-fortigate-344487-global-commands-for-str-a7c5b27c

Global commands for stronger and more secure encryption
This section describes some best practices for employing stronger and more secure encryption.
Turn on global strong encryption
Enter the following command to configure FortiOS to use only strong encryption and allow only strong ciphers (AES, 3DES) and digest (SHA1) for HTTPS, SSH, TLS, and SSL functions.
config system global
set strong-crypto enable
end
Disable MD5 and CBC for SSH
In some cases, you may not be able to enable strong encryption. For example, your FortiGate may be communicating with a system that does not support strong encryption. With strong-crypto disabled you can use the following options to prevent SSH sessions with the FortiGate from using less secure MD5 and CBC algorithms:
config system global
set ssh-hmac-md5 disable
set ssh-cbc-cipher disable
end
Disable static keys for TLS
You can use the following command to prevent all TLS sessions that are terminated by FortiGate from using static keys (AES128-SHA, AES256-SHA, AES128-SHA256, AES256-SHA256):
config system global
set ssl-static-key-ciphers disable
end
Require larger values for Diffie-Hellman exchanges
Larger Diffie-Hellman values result in stronger encryption. Use the following command to force Diffie-Hellman exchanges to use 8192 bit values (the highest configurable DH value).
config system global
set dh-params 8192
end
