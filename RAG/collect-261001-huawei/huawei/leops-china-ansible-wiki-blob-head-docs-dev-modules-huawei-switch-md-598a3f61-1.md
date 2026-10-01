---
id: collect-261001-huawei/huawei/leops-china-ansible-wiki-blob-head-docs-dev-modules-huawei-switch-md-598a3f61-1
title: "leops-china-ansible-wiki-blob-head-docs-dev-modules-huawei-switch-md-598a3f61"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-huawei/leops-china-ansible-wiki-blob-head-docs-dev-modules-huawei-switch-md-598a3f61.md
source_anchor: ""
source_lines: [1, 86]
sha256: 734171352ec7bb9a18b95433c984b510522ac2ef4cfb36e51d82763ab7cf2e46
---

# leops-china-ansible-wiki-blob-head-docs-dev-modules-huawei-switch-md-598a3f61

远程在switch上执行命令
使用ssh连接到switch上，远程输入命令到switch，获取返回结果。
华为的s5700交换机
s5700
<Huawei>system-view 
[Huawei]rsa local-key-pair create 
The key name will be: Huawei_Host
% RSA keys defined for Huawei_Host already exist.
Confirm to replace them? [y/n]:y
The range of public key size is (512 ~ 2048). 
NOTES: If the key modulus is greater than 512, 
       it will take a few minutes.
Input the bits in the modulus[default = 512]:2048
Generating keys...
..........+++
.....................................................................+++
..............++++++++
.............++++++++
[Huawei]user-interface vty 0 4
[Huawei-ui-vty0-4]authentication-mode aaa
[Huawei-ui-vty0-4]protocol inbound ssh
[Huawei-ui-vty0-4]quit
[Huawei]aaa
[Huawei-aaa]local-user user1 password cipher Huawei@123
[Huawei-aaa]local-user user1 service-type ssh
[Huawei-aaa]quit
[Huawei]ssh user user1 authentication-type password
Info: Succeeded in adding a new SSH user.
[Huawei]ssh user user1 service-type stelnet
[Huawei]stelnet server enable
Info: Succeeded in starting the Stelnet server.
就可以使用ssh方式连接交换机了。
- 先使用伪代码定义Hwcon类以及方法
class Hwcon(object): def __init__(self, address, username, password, port=22): # 初始化ssh连接 def close(self): # 关闭连接 def openShell(self): # 在ssh连接中打开一个terminal def send_command(self, command=''): # 向terminal发送命令 def get_command_result(self, cmd): # 获取命令的输出结果 def parse_result_data(self, data): # 解析命令的输出结果 def save_config(self): # 保存配置 def run(self, cmd): # 执行命令
使用方法：
python connection = Hwcon(b_host, b_user, b_password, module.params['sport']) connection.openShell() rc,stdout,stderr = connection.run(b_command) connection.save_config() connection.close() 
就可完成我们执行命令的功能，当然这只是伪代码，下面我们开始填充里面的功能代码。
- 
框架我们定义好了，下面定义命令的错误，正确输出信息的正则表达式，以便我们能确认命令输出结果是否正确。 terminal_stdout_re = [ re.compile(r'[\r\n]?<.+>(?:\s*)$'), re.compile(r'[\r\n]?\[.+\](?:\s*)$'), ] terminal_stderr_re = [ re.compile(r"% ?Error: "), re.compile(r"^% \w+", re.M), re.compile(r"% ?Bad secret"), re.compile(r"invalid input", re.I), re.compile(r"(?:incomplete|ambiguous) command", re.I), re.compile(r"connection timed out", re.I), re.compile(r"[^\r\n]+ not found", re.I), re.compile(r"'[^']' +returned error code: ?\d+"), re.compile(r"syntax error"), re.compile(r"unknown command"), re.compile(r"Error\[\d+\]: ", re.I), re.compile(r"Error:", re.I) ]
- 
导入我们需要的库 import time import re from ansible.module_utils.basic import AnsibleModule from ansible.module_utils._text import to_bytes from ansible.errors import AnsibleError, AnsibleConnectionFailure try: import paramiko except ImportError: raise AnsibleError("paramiko is not installed, please use pip install paramiko") try: from __main__ import display except ImportError: from ansible.utils.display import Display display = Display()
- 
填充我们的hwcon类 class Hwcon(object): shell = None client = None def __init__(self, address, username, password, port=22): display.vv("Connecting to network device on ip", str(address) + ".") self.client = paramiko.client.SSHClient() self.client.set_missing_host_key_policy(paramiko.client.AutoAddPolicy()) self.client.connect(address, port=port, username=username, password=password, look_for_keys=False, allow_agent=False) def close(self): if self.client is not None: self.client.close() def openShell(self): self.shell = self.client.invoke_shell() def send_command(self, command=''): if self.shell: if command not in ('?',): command += "\n" self.shell.send(command) while True: if self.shell.recv_ready() or self.shell.recv_stderr_ready(): break time.sleep(0.1) def get_command_result(self, cmd): buffersize = 4096 self.send_command() self.shell.recv(buffersize) self.send_command(cmd) stdout = self.shell.recv(buffersize) b_data = stdout.split('\r\n') result_tmp = '' while '- More -' in b_data[-1]: self.shell.send("\n") time.sleep(0.1) tmp = self.shell.recv(buffersize) b_data = tmp.split('\r\n') result_tmp += tmp if self.shell.recv_stderr_ready(): stderr = self.shell.recv_stderr(buffersize) else: stderr = '' stdout = '\r\n'.join(stdout.split('\r\n')[:-1]) + '\n' + result_tmp stdout = stdout.replace(' - More -', '').replace( '\x1b[42D \x1b[42D', '') return stdout, stderr def parse_result_data(self, data): b_data = data.split('\r\n') result = b_data[1:-1] return '\r\n'.join(result) def save_config(self): rc = 1 buffersize = 4096 self.send_command() stdout = self.shell.recv(buffersize) t1 = terminal_stdout_re[0].findall(stdout) while not t1: self.send_command('quit') stdout = self.shell.recv(buffersize) t1 = terminal_stdout_re[0].findall(stdout) self.send_command('save') time.sleep(0.1) self.send_command('y') time.sleep(3) stdout = self.shell.recv(buffersize) if stdout.find('successfully'): rc = 0 return rc def run(self, cmd): rc = 1 stdout, stderr = self.get_command_result(cmd) for regex in terminal_stderr_re: r1 = regex.findall(stdout) if not r1: stdout = self.parse_result_data(stdout) rc = 0 return rc, stdout, stderr
- 
把使用方法填入main方法中 def main(): module = AnsibleModule( argument_spec = dict( command=dict(required=True, type='str'), shost=dict(required=True, type='str'), sport=dict(required=False, type='int', default=22), suser=dict(required=True, type='str'), spass=dict(required=True, type='str', no_log=True), save=dict(required=False, type='bool') ), supports_check_mode=True ) b_command = to_bytes(module.params['command'], errors='surrogate_or_strict') b_host = to_bytes(module.params['shost'], errors='surrogate_or_strict') b_user = to_bytes(module.params['suser'], errors='surrogate_or_strict') b_password = to_bytes(module.params['spass'], errors='surrogate_or_strict') result = {'changed': False} try: connection = Hwcon(b_host, b_user, b_password, module.params['sport']) except Exception as e: raise AnsibleConnectionFailure(str(e)) try: connection.openShell() except Exception as e: msg = "Failed to open session" if len(str(e)) > 0: msg += ": %s" % str(e) raise AnsibleConnectionFailure(msg) display.vvv("EXEC %s" % b_command, host=b_host) try: rc,stdout,stderr = connection.run(b_command) except Exception as e: raise AnsibleError('Exec command error.\n' + str(e) ) if rc == 0 and b_command.find('display'): result['changed'] = True elif rc == 1: module.fail_json(msg=stdout+stderr) if module.params['save']: try: save_rc = connection.save_config() except Exception as e: raise AnsibleError('Save config error.\n' + str(e) ) if save_rc != 0: msg= "not save config!" module.fail_json(msg=msg) connection.close() result.update({ 'command': b_command, 'rc': rc, 'stdout': stdout, 'stderr': stderr}) module.exit_json(**result) if __name__ == '__main__': main()
添加完描述信息，完整的代码如下
#!/usr/bin/python
ANSIBLE_METADATA = {'status': ['preview'],
                    'supported_by': 'community',
                    'metadata_version': '1.0'}
DOCUMENTATION = ''' 
module: hwos_command
version_added: "2.3"
short_description: Run arbitrary command on HUAWEI network devices.
description:
     - Run arbitrary command on HUAWEI network devices.
options:
  command:
    description:
      - HUAWEI network devices command
    required: true
  shost:
    description:
      - remote network devices ip address
    required: true
  sport:
    description:
      - remote network devices ssh port
    required: false
    default: 22
  suser:
    description:
      - remote network devices ssh user
    required: true
  spass:
    description:
      - remote network devices ssh user password
    required: true
  save:
    description:
      - save current configuration
    required: false
    default: no
author:
    - "Lework"
'''
