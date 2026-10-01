---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-0-best-practices-555436-hardening-784ac110-2
title: "document-fortigate-7-6-0-best-practices-555436-hardening-784ac110"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["asic"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-0-best-practices-555436-hardening-784ac110.md
source_anchor: ""
source_lines: [81, 125]
sha256: fb8c3ab0a06571331d9cb4548f97a23a326484ddff5c2bedc6a7bd958efd3aa3
---

# document-fortigate-7-6-0-best-practices-555436-hardening-784ac110

- 
                                                    Where possible, enable ASIC DoS for offloading using network processor ASICs. The FortiOS Hardware Acceleration Guide contains more information about DoS-related NP ASIC features, such as configuring NP6 anomaly protection and using the host protection engine (HPE) to protect the FortiGate from DoS attacks.
Secure password storage
The passwords, and private keys used in certificates, that are stored on the FortiGate are encrypted using a predefined private key, and encoded when displayed in the CLI and configuration file. System admin passwords are hashed with SHA256 and encoded before being displayed. In FortiOS 7.6.1 and later, admin passwords are hashed with PBKDF2. See Enhanced administrator password security for more information.
Passwords cannot be decrypted without the private key and are not shown anywhere in clear text. The private key is required on other FortiGates to restore the system from a configuration file. In an HA cluster, the same key should be used on all of the units.
To enhance password security, specify a custom private key for the encryption process. This ensures that the key is only known by you.
FortiGate models with a Trusted Platform Module (TPM) can store the master encryption password, which is used to generate the master encryption key, on the TPM. For more information, see Trusted platform module support.
To configure your own private encryption key:
config system global
    set private-data-encryption enable
end
Please type your private data encryption key (32 hexadecimal numbers):
********************************
Please re-enter your private data encryption key (32 hexadecimal numbers) again:
********************************
Your private data encryption key is accepted.
                                            In FortiOS 7.6.1 and later, FortiGate no longer requires the user to input the key. Instead, FortiGate generates a random password. See Use per-FortiGate generated random password for private-data-encryption for more information.
Configuration backup
The FortiGate configuration file has important information that should always be kept secured, including details about your network, users, credentials, passwords, and keys. There are many reasons to back up your configuration, such as disaster recovery, preparing for migrating to another device, and troubleshooting. Evaluate the risk involved if your configurations were exposed, and manage your risk accordingly.
When backing up your configuration, consider the following steps to safeguard the file:
- 
                                                    Enable Encryption when backing up the configuration.
- 
                                                    Store the configuration file in a secure location.
- 
                                                    Delete old configuration files that are no longer needed.
If a configuration file must be shared with a third party for auditing, troubleshooting, or any other reasons, consider only providing a section of the file and not the entire file. Otherwise, consider the following steps:
- 
                                                    Enable Encryption when backing up the configuration and only share the password with the intended party.
- 
                                                    Manually replace the passwords in the backed up configuration file, or enable Password Masking when backing up the configuration.
- 
                                                    Request that the configuration file be deleted after the intended purpose has been satisfied.
If FortiGate has private-data-encryption enabled, you can only restore the configuration file on a FortiGate with the same encryption key configured.
Keep this in mind for FortiOS 7.6.1 and later where the encryption key is automatically generated. As such, a configuration that is backed up while private-data-encryption is enabled cannot be restored when private-data-encryption is disabled or when private-data-encryption is re-enabled because it generates a different random key.
RMA considerations
When a device has private-data-encryption enabled in FortiOS 7.6.1 and later, and the hardware malfunctions, you must disable private-data-encryption and back up the configuration. Then you can restore the configuration backup on a replacement unit with private-data-encryption disabled. After restoring the configuration backup, you can enable the private-data-encryption setting on the replacement unit.
Depending on the reason the hardware malfunctioned, you may be unable to complete this operation. Therefore, consider this risk when you enable private-data-encryption.
Non-standard admin ports and administrator usernames
FortiGate is configured with default administrative access ports under System > Settings. These ports are well known and likely to be targeted by malicious actors in the first pass. Similarly, the FortiGate has a default system administrator that is also well known. It is highly recommended to change the default ports and username to non-standard and non-guessable ports and names for an added layer of protection.
Blocking external access to administrative ports
It is generally not recommended to allow external (WAN) access to administrative ports on the FortiGate. A better solution is to configure administrative access on a trusted management interface where the management computer must be in the physical location, or accessible only through a trusted connection like a VPN for remote access. Ideally, this connection is out-of-band, meaning that it does not rely on the connection passing through the FortiGate. For information about configuring administrative access on interfaces, see Interface Settings > Configure administrative access to interfaces.
If access must be granted on an external and public interface, ensure that a local-in policy is defined to allow only trusted hosts to connect, or restrict administrator accounts logins to trusted hosts only. See Restricting logins to trusted hosts.
Local-in policies offer granularity in defining the hosts, or groups of hosts that are allowed or blocked. For example, using the ISDB or Geo-IP database, administrators can restrict a specific geo-location from accessing the administrative port and interface, or open up specific regions. By enabling logs on the local-in policy, you can also perform detailed forensic analysis on intrusion attempts.
For more information on Local-in policies, see Local-in policy.
