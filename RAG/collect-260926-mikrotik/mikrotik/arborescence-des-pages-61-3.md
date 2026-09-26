---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-61-3
title: "System Requirements"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "license", "memory"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-61.md
source_anchor: ""
source_lines: [291, 409]
sha256: c38e129913d897cd2f52c9f89b29145acf929ee3ce9e27985a7a3f37be876a78
---

# System Requirements

### Changing MTU

VMware ESXi supports MTU of up to 9000 bytes. To get the benefit of that, you have to adjust your ESXi installation to allow a higher MTU. Virtual Ethernet interface added **after** the MTU change will be properly allowed by the ESXi server to pass jumbo frames. Interfaces added prior to MTU change on the ESXi server will be barred by the ESXi server (it will still report the old MTU as the maximum possible size). If you have this, you have to re-add interfaces to the virtual guests.

**Example.** There are 2 interfaces added to the ESXi guest, auto-detected MTU on the interfaces show MTU size as it was at the time when the interface was added:

[admin@chr-vm] > interface ethernet print 
Flags: X - disabled, R - running, S - slave 
 #    NAME           MTU MAC-ADDRESS       ARP       
 0 R  ether1        9000 00:0C:29:35:37:5C enabled   
 1 R  ether2        1500 00:0C:29:35:37:66 enabled

### Using bridge on Linux

If Linux bridge supports IGMP snooping, and there are problems with IPv6 traffic it is required to disable that feature as it interacts with MLD packets (multicast) and is not passing them through.

echo -n 0 > /sys/class/net/vmbr0/bridge/multicast_snooping

### Packets not passing from guests

The problem: after configuring a software interface (VLAN, EoIP, bridge, etc.) on the guest CHR it stops passing data to the outside world beyond the router.

The solution: check your VMS (Virtualization Management System) security settings, if other MAC addresses are allowed to pass and if packets with VLAN tags are allowed to pass through. Adjust the security settings according to your needs like allowing MAC spoofing or a certain MAC address range. For VLAN interfaces, it is usually possible to define allowed VLAN tags or VLAN tag range.

### Using VLANs on CHR in various Hypervisors

In some hypervisors, before VLAN can be used on VMs, they need to first be configured on the hypervisor itself.

#### ESXI

Enable Promiscuous mode in a port group or virtual switch that you will use for a specific VM.

**ESX documentation:**

#### Hyper-V

**Hyper-V documentation:**

#### bhyve hypervisor

It won't be possible to run CHR on this hypervisor. CHR cannot be run as a para-virtualized platform.

#### Linode

When creating multiple Linodes with the same disk size, new Linodes will have the same systemID. This will cause issues to get a Trial/Paid license. To avoid this, run the command `/system license generate-new-id` after the first boot and before you request a trial or paid license. This will make sure the ID is unique.

**Some useful articles:**

Specific VLAN is untagged by NIC interface:

- https://blogs.msdn.microsoft.com/adamfazio/2008/11/14/understanding-hyper-v-vlans/
- https://www.aidanfinn.com/?p=10164

Allow passing other VLANs:

# Guest tools

## VMWare

### Time synchronization

Must be enabled from GUI ('Synchronize guest time with host'). Backward synchronization is disabled by default - if the guest is ahead of the host by more than ~5 seconds, synchronization is not performed

### Power operations

- *poweron* and*resume* scripts are executed (if present and enabled) after power on and resume operations respectively.
- *poweroff* and*suspend* scripts are executed before power off and suspend operations respectively.
- If scripts take longer than 30 seconds or contain errors, the operation fails
- In case of failure, retrying the same operation will ignore any errors and complete it successfully
- Failed script output is saved to a file (e. g. 'poweroff-script.log', 'resume-script.log' etc)
- Scripts can be enabled/disabled from hypervisor GUI ('run VMware Tools Scripts') or by enabling/disabling scripts from the console

### Quiescing/backup

Guest filesystem quiescing is performed only if requested.

- *freeze* script is executed before freezing the filesystem
- *freeze-fail* script is executed if the hypervisor failed to prepare for a snapshot or if*freeze* script failed
- *thaw* script is executed after the snapshot has been taken
- Script run time is limited to 60 seconds
- *freeze* script timeouts and errors result in the backup operation being aborted
- FAT32 disks are not quiesced
- Failed script output is saved to a file (e. g. 'freeze-script.log', 'freeze-fail-script.log', 'thaw-script.log')

### Guest info

Networking, disk, and OS info are reported to the hypervisor every 30 seconds (GuestStats (memory) are disabled by default, and can be enabled by setting 'guestinfo.disable-perfmon = "FALSE"' in VM config).

- The order, in which network interfaces are reported, can be controlled by setting 'guestinfo.exclude-nics', 'guestinfo.primary-nics' and 'guestinfo.low-priority-nics' options. Standard wildcard patterns can be used.

### Provisioning

You can use the ProcessManager from Vim API to execute scripts. Python bindings are available

- Main data structure: GuestProgramSpec
  - The *workingDirectory* and*envVariables* members are ignored
  - *programPath* must be set to either 'inline' or 'import'
  - If *programPath* is '**inline'** ,*arguments* are interpreted as script text
  - If *programPath* is '**import'** ,*arguments* are interpreted as file path
- The 

After using *GuestProgramSpec* together with an instance of GuestAuthentication as arguments to StartProgramInGuest unique *JobID* is obtained.

Script progress can be tracked by using the ListProcessesInGuest command. *ListProcessesInGuest* accepts an array of job id's; passing an empty array will report on all jobs started from the API

- *ListProcessesInGuest* returns an array of GuestProcessInfo instances:
  - *pid* field is set to*JobID*
  - *endTime* is only set after completion
  - *exitCode* is set to 0 on success and -1 on error
  - *name* is set to 'inline' or 'import' (same as*programPath* in*GuestProgramSpec* )

Information about completed jobs is kept around for ~1 minute, or until *ListProcessesInGuest* (with the corresponding *JobID*) is called. If the script fails, a file named 'vix_job_$JobID$ .txt' containing the script output is created. Script run time is limited to 120 seconds and script output is not saved on timeout,

- The vmrun command *runScriptInGuest* can also be used
- The PowerCLI cmdlet Invoke-VMScript is not supported
- Host/guest file transfer is not supported

##### Python example

