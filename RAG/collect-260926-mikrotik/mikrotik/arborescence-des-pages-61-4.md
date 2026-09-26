---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-61-4
title: "System Requirements"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "datacenter"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-61.md
source_anchor: ""
source_lines: [410, 519]
sha256: e5ceef98706e85c5527f99c0a95000b899030ef291cf376f022f80b69f6c7624
---

# System Requirements

```
#!/usr/bin/env python
# -*- coding: utf-8 -*-
import sys,time
from pyVim import connect
from pyVmomi import vmodl,vim
def runInline(content,vm,creds,source):
    ''' Execute script source on vm '''
    if isinstance(source, list):
        source = '\n'.join(source)
    ps = vim.vm.guest.ProcessManager.ProgramSpec(
                programPath = 'console',
                arguments = source
        )
    return content.guestOperationsManager.processManager.StartProgramInGuest(vm,creds,ps)
def runFromFile(content,vm,creds,fileName):
    ''' Execute script file located on CHR '''
    ps = vim.vm.guest.ProcessManager.ProgramSpec(
                programPath = 'import',
                arguments = fileName
    )
    return content.guestOperationsManager.processManager.StartProgramInGuest(vm,creds,ps)
def findDatastore(content,name):
    sessionManager = content.sessionManager
    dcenterObjView = content.viewManager.CreateContainerView(content.rootFolder, [vim.Datacenter], True)
    datacenter = None
    datastore = None
    for dc in dcenterObjView.view:
        dstoreObjView = content.viewManager.CreateContainerView(dc, [vim.Datastore], True)
        for ds in dstoreObjView:
            if ds.info.name == name:
                datacenter = dc
                datastore = ds
                break
        dstoreObjView.Destroy()
    dcenterObjView.Destroy()
    return datacenter,datastore
def _FAILURE(s,*a):
    print(s.format(*a))
    sys.exit(-1)
#------------------------------------------------------------------------------#
if __name__ == '__main__':
    host = sys.argv[1] # ip or something
    user = 'root'
    pwd = 'MikroTik'
    vmName = 'chr-test'
    dataStoreName = 'datastore1'
    service = connect.SmartConnectNoSSL(host=host,user=user,pwd=pwd)
    if not service:
        _FAILURE("Could not connect to the specified host using specified username and password")
    content = service.RetrieveContent()
    #---------------------------------------------------------------------------
    # Find datacenter and datastore
    datacenter,datastore = findDatastore(content,dataStoreName)
    if not datacenter or not datastore:
        connect.Disconnect(service)
        _FAILURE('Could not find datastore \'{}\'',dataStorename)
    #---------------------------------------------------------------------------
    # Locate vm
    vmxPath = '[{0}] {1}/{1}.vmx'.format(dataStoreName, vmName)
    vm = content.searchIndex.FindByDatastorePath(datacenter, vmxPath)
    if not vm:
        connect.Disconnect(service)
        _FAILURE("Could not locate vm")
    #---------------------------------------------------------------------------
    # Setup credentials from user name and pasword
    creds = vim.vm.guest.NamePasswordAuthentication(username = 'admin', password = '')
    #---------------------------------------------------------------------------
    # Run script
    pm = content.guestOperationsManager.processManager
    try:
        # Run script
        src = [':ip address add address=192.168.0.1/24 interface=ether1;']
        jobID = runInline(content, vm, creds, src)
        # Or run file (from FTP root)
        # jobID = runFromFile(content,vm,creds, 'scripts/provision.rsc')
        #---------------------------------------------------------------------------
        # Wait for job to finish
        pm = content.guestOperationsManager.processManager
        jobInfo = pm.ListProcessesInGuest(vm, creds, [jobID])[0]
        while jobInfo.endTime is None:
            time.sleep(1.0)
            jobInfo = pm.ListProcessesInGuest(vm, creds, [jobID])[0]
        if jobInfo.exitCode != 0:
            _FAILURE('Script failed!')
    except:
        raise
    else:
        connect.Disconnect(service)
```
## KVM

QEMU guest agent is available. Supported agent commands can be retrieved by using the guest-info command. Host-guest file transfer can be performed by using guest-file-* commands. Guest networking information can be retrieved by using the guest-network-get-interfaces command.

- Scripts can be executed by using the guest-exec command together with the GuestExec data structure:
  - If the *path* member is provided, the corresponding file is executed
  - If the *path* member is not set and*input-data* member is provided,*input-data* value is used as script input
  - If *capture-output* is set, script output is reported back
  - *args* and*env* members are not used
- If the 

- Script job progress can be monitored with guest-exec-status command. The GuestExecStatus data structure is populated as follows:
  - On success, *exitcode* member is set to 0
  - If the script timed out *exitcode* is set to 1
  - If the script contained errors *exitcode* is set to -1
  - *signal* member is not set
  - The *err-data* member is not used
  - If *capture-output* was true, Base64 encoded script output is stored in*out-data*
- On success, 
- An additional agent channel ('chr.provision_channel') is also available
