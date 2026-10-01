---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hc-en-us-articles-235247068-adding-ssh-keys-to-unifi-devices-7c2eb803
title: "hc-en-us-articles-235247068-adding-ssh-keys-to-unifi-devices-7c2eb803"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/hc-en-us-articles-235247068-adding-ssh-keys-to-unifi-devices-7c2eb803.md
source_anchor: ""
source_lines: [1, 37]
sha256: 10d0bb1ec8d26aea47d2e221d6380651e753d938ae0d7553c7dad49eacab1d78
---

# hc-en-us-articles-235247068-adding-ssh-keys-to-unifi-devices-7c2eb803

Adding SSH Keys to UniFi Devices
UniFi makes it simple for administrators to add SSH keys to UniFi Devices. If you are just looking to enable SSH or establish an SSH connection with UniFi devices, see our article here.
Instructions
Before you begin: Adding SSH keys to UniFi devices through the UniFi Network application should only be done by advanced system administrators familiar with SSH.
Instructions for each operating system are found below:
Windows
Linux
macOS
On Windows 10, it is also possible to use WSL to generate SSH keys. In this case, the commands will be identical to the macOS/Linux.
- Download and install the full PuTTY installer package or the standalone PuTTYgen binary.
- Open Putty Key Generator (PuTTYgen) to generate a new SSH key.
- Select RSA with a key size of 2048 and select Generate.
- Save the public key and private key in a secure location after they have been generated.
  - Note: Never share the private key file or its contents.
- 
Copy the complete contents of the SSH key file from the box labeled Public key for pasting into OpenSSH authorized_keys file. The SSH key starts with ssh-rsa and ends with== rsa-key-<date> , for example:ssh-rsa AbCdEfGh1234AbCdEfGh== rsa-key-20200911
- Add the key to UniFi Network by navigating to Settings > Devices > Device Updates and Settings > Device SSH Settings.
- Open the Linux Terminal.
- 
Use the ssh-keygen command to generate a new SSH Key: ssh-keygen -t rsa
- Either use the default file name or enter a file name of your choice and optionally add a passphrase.
  - Note: Never share the private key file or its contents.
- 
View the contents of the public SSH key file. Assuming the file name was not changed: cat .ssh/id_rsa.pub
- 
Copy the complete contents of the SSH key file from the Terminal window to your clipboard. The SSH key starts with ssh-rsa and ends withusername@machine , for example:ssh-rsa AbCdEfGh1234AbCdEfGh admin@computer
- Add the key to UniFi Network by navigating to Settings > System > Advanced > Device Authentication > SSH Keys.
- Open the macOS Terminal, located in Finder > Applications > Utilities.
- 
Use the ssh-keygen command to generate a new SSH Key: ssh-keygen -t rsa
- Either use the default file name or enter a file name of your choice and optionally add a passphrase.
  - Note: Never share the private key file or its contents.
- 
View the contents of the public SSH key file. Assuming the file name was not changed: cat .ssh/id_rsa.pub
- 
Copy the complete contents of the SSH key file from the Terminal window to your clipboard. The SSH key starts with ssh-rsa and ends withusername@machine , for example:ssh-rsa AbCdEfGh1234AbCdEfGh admin@computer
- Add the key to UniFi Network by navigating to Settings > System > Advanced > Device Authentication > SSH Keys.
