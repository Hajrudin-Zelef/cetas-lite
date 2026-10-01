---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1167958-unifi-cloud-key-gen-1-ssh-with-public-key-authentication-f21359cd
title: "questions-1167958-unifi-cloud-key-gen-1-ssh-with-public-key-authentication-f21359cd"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1167958-unifi-cloud-key-gen-1-ssh-with-public-key-authentication-f21359cd.md
source_anchor: ""
source_lines: [1, 42]
sha256: 8a4c5a5202b181ee0d3d497f4b9d842c01f220fc890cf8bd5662fe0b6f54a53e
---

# questions-1167958-unifi-cloud-key-gen-1-ssh-with-public-key-authentication-f21359cd

I'm trying to configure my Gen 1 Cloud Key (CK) to allow passwordless authentication so I can run a bash script to download autobackup files. I'm able to SSH into the CK "manually" and also with my bash script using username/password, but this would store the credentials in clear text which I (obviously) don't want to do.
I have:
- generated the keys on my local machine (MBP Sequoia 15.1.1) with ssh-keygen -t rsa ,
- copied the public key to the CK's ~/.ssh/authorized_keys file,
- modified /etc/ssh/sshd_config to include the following (uncommented) lines:RSAAuthentication yes ,PubkeyAuthentication yes , andAuthorizedKeysFile .ssh/authorized_keys ,
- restarted ssh on the CK with sudo systemctl restart sshd ,
- checked permissions with chmod 700 ~/.ssh andchmod 600 ~/.ssh/authorized_keys .
But still no joy. SSH log says,
...
debug1: Next authentication method: publickey
debug1: get_agent_identities: bound agent to hostkey
debug1: get_agent_identities: ssh_fetch_identitylist: agent contains no identities
debug1: Will attempt key: /Users/Hank/.ssh/id_rsa RSA SHA256:KEY_VALUE_REDACTED
debug1: Will attempt key: /Users/Hank/.ssh/id_ecdsa 
debug1: Will attempt key: /Users/Hank/.ssh/id_ecdsa_sk 
debug1: Will attempt key: /Users/Hank/.ssh/id_ed25519 
debug1: Will attempt key: /Users/Hank/.ssh/id_ed25519_sk 
debug1: Will attempt key: /Users/Hank/.ssh/id_xmss 
debug1: Will attempt key: /Users/Hank/.ssh/id_dsa 
debug1: Offering public key: /Users/Hank/.ssh/id_rsa RSA SHA256:KEY_VALUE_REDACTED
debug1: send_pubkey_test: no mutual signature algorithm
debug1: Trying private key: /Users/Hank/.ssh/id_ecdsa
debug1: Trying private key: /Users/Hank/.ssh/id_ecdsa_sk
debug1: Trying private key: /Users/Hank/.ssh/id_ed25519
debug1: Trying private key: /Users/Hank/.ssh/id_ed25519_sk
debug1: Trying private key: /Users/Hank/.ssh/id_xmss
debug1: Trying private key: /Users/Hank/.ssh/id_dsa
debug1: Next authentication method: password
Hank@10.0.1.123's password: 
...
and so I am still prompted for a password. I believe the clue to the problem lies with this line:
no mutual signature algorithm
Does this mean that the instances of SSH on the two devices are incompatible and thus using key-based auth is not possible? This seems like it should be straightforward, but something isn't quite right.
TLDR; The Unifi controller software supports SSH key-baed auth, but this only applies to attaching to network devices like APs, switches and routers. The CK is a client and it's apparently not possible to accomplish my goal via the controller software.
There's a fair bit of chatter surrounding this approach, but nothing seems to address this specific issue.
EDIT: I did find this, which is leading me to believe that my CK (which is admittedly ancient) may be the source of the problem. Upgrading ssh on the CK, while theoretically possible, is a suboptimal solution because I'd expect future firmware upgrades would overwrite any intermediate changes.
Is the answer to use something other than RSA?
EDIT2: Solution - FTF, here are the revised commands that worked in my case:
- Generate key pair: ssh-keygen -t ed25519
- Copy public key to server: ssh-copy-id -i ~/.ssh/id_ed25519.pub User@10.0.1.123
- Check permissions: chmod 700 ~/.ssh andchmod 600 ~/.ssh/authorized_keys
- Restart: sudo systemctl restart sshd (may not be required)
