---
id: collect-260926-mikrotik/mikrotik/questions-1053198-how-do-i-generate-an-sshfp-record-for-a-mikrotik-router-4822ae36
title: "questions-1053198-how-do-i-generate-an-sshfp-record-for-a-mikrotik-router-4822ae36"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1053198-how-do-i-generate-an-sshfp-record-for-a-mikrotik-router-4822ae36.md
source_anchor: ""
source_lines: [1, 21]
sha256: b29718a397d59605df9543a173c69478b09a47602afa2d3d9dedca327fea7304
---

# questions-1053198-how-do-i-generate-an-sshfp-record-for-a-mikrotik-router-4822ae36

I want to generate an SSHFP record for my Mikrotik CCR2004 running RouterOS 6.47.4, without getting the key over the network. How can I do this from the console?
1 Answer 1
Regenerate the host key if needed (e.g. if changing the modulus length):
ip ssh regenerate-host-key
The CLI will hang for a moment. When it's back:
ip ssh export-host-key
file print
You'll get something like:
[admin@MikroTik] > file print
 # NAME                   TYPE                        SIZE CREATION-TIME
 0 hostKey_rsa            file                        3272 jan/03/1970 18:35:53
 1 hostKey_rsa.pub        ssh key                      796 jan/03/1970 18:35:53
 2 hostKey_dsa            file                         668 jan/03/1970 18:35:53
 3 hostKey_dsa.pub        ssh key                      604 jan/03/1970 18:35:53
Print the one you want:
file print detail where name=hostKey_rsa.pub
Remove the extra spaces from the beginning of the printed key, and put it in a file e.g. router_pubkey.pem. Then on your box:
openssl rsa -in router_pubkey.pem -pubin -RSAPublicKey_out |
  ssh-keygen -f /dev/stdin -i -m PEM |
  ssh-keygen -f /dev/stdin -r hostname
The openssl command converts the key to PKCS#1 format (just an integer sequence, instead of an ASN.1 rsaEncryption object). The ssh-keygen invocations respectively import the key into OpenSSH format, and then generate the fingerprints and print SSHFP records.
