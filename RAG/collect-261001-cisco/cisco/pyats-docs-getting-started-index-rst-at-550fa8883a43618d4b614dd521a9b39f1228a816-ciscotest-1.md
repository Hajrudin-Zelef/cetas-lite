---
id: collect-261001-cisco/cisco/pyats-docs-getting-started-index-rst-at-550fa8883a43618d4b614dd521a9b39f1228a816-ciscotest-1
title: "Example"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/pyats-docs-getting-started-index-rst-at-550fa8883a43618d4b614dd521a9b39f1228a816-ciscotestautomation.md
source_anchor: ""
source_lines: [1, 150]
sha256: 039dee6b004f300cdcbbf3e433b2ff6b8281e1e4770a6f3a1268e9614c017b8f
---

# Example

pyATS is a foundation-layer test framework. It is designed to provide a reasonable end-to-end test environment for developers to write test cases in, featuring multiple packages and modules making writing networking-related tests a breeze.

The goal of this document is to provide an informal overview of how pyATS works, and just enough technical specifics to get you started with your first test project.

pyATS currently supports Python `v3.7.x`,``v3.8.x``,``v3.9.x``and``v3.10.x``, and is
tested to work on the following platforms:


- Linux (tested with CentOS, RHEL, Ubuntu, Alpine)
- Mac OSX (10.13+)

Note

**Cisco Internal Engineering**

For all Cisco engineering users (eg, if you are a Cisco employee), refer to installation section of the pyATS Wiki for guidelines on how to install pyATS correctly in your environment.

For all external users (Cisco DevNet users, customers, general public), you can install pyATS from public PyPI directly using pip install:

- `pip install pyats`
- **core framework only** : this installs just the core pyATS framework with
zero optional extras.
- `pip install pyats[library]`
- **pyATS + pyATS Library/Genie** : this installs the core pyATS framework,
and the standard pyATS network automation library,
Genie.*This is the recommended option*

- `pip install pyats[robot]`
- **pyATS framework with RobotFramework support** . This installs the optional
RobotFramework package, and`pyats.robot` package that features pyATS
specific keywords.
- `pip install pyats[template]`
- installs the cookiecutter dependency require for `pyats create project` command
- `pip install pyats[full]`
- installs pyATS, along with **all optional extras** .

Because we are in the business of network test automation, pyATS is designed
around the concept of :ref:`testbeds <topology_concept>`: where you describe
your *devices under testing* in YAML format.

This :ref:`testbed YAML file <topology_testbed_file>` provides many :ref:`sections <schema>` for you to describe your physical devices, and how they link together to form the topology.

```
# Example
# -------
#
#   an example testbed file - ios_testbed.yaml
testbed:
    name: IOS_Testbed
    credentials:
        default:
            username: admin
            password: cisco
        enable:
            password: cisco
devices:
    ios-1: # <----- must match to your device hostname in the prompt
        os: ios
        type: ios
        connections:
            a:
                protocol: telnet
                ip: 1.1.1.1
                port: 11023
    ios-2:
        os: ios
        type: ios
        connections:
            a:
                protocol: telnet
                ip: 1.1.1.2
                port: 11024
            vty:
                protocol: ssh
                ip: 5.5.5.5
topology:
    ios-1:
        interfaces:
            GigabitEthernet0/0:
                ipv4: 10.10.10.1/24
                ipv6: '10:10:10::1/64'
                link: link-1
                type: ethernet
            Loopback0:
                ipv4: 192.168.0.1/32
                ipv6: '192::1/128'
                link: ios1_Loopback0
                type: loopback
    ios-2:
        interfaces:
            GigabitEthernet0/0:
                ipv4: 10.10.10.2/24
                ipv6: '10:10:10::2/64'
                link: link-1
                type: ethernet
            Loopback0:
                ipv4: 192.168.0.2/32
                ipv6: '192::2/128'
                link: ios2_Loopback0
                type: loopback
```
Once a testbed yaml file is written, you can load it, query your topology, connect & issue commands to your devices using :ref:`the APIs <topology_usage>`.

This is the best way to validate whether your topology file is well formed, and your devices connectable.

```
# loader our newly minted testbed file
from pyats.topology import loader
testbed = loader.load('ios_testbed.yaml')
# access the devices
testbed.devices
# AttrDict({'ios-1': <Device ott-tb1-n7k4 at 0xf77190cc>,
#           'ios-2': <Device ott-tb1-n7k5 at 0xf744e16c>})
ios_1 = testbed.devices['ios-1']
ios_2 = testbed.devices['ios-2']
# find links from one device to another
for link in ios_1.find_links(ios_2):
    print(repr(link))
# <Link link-1 at 0xf744ef8c>
# establish basic connectivity
ios_1.connect()
# issue commands
print(ios_1.execute('show version'))
ios_1.configure('''
    interface GigabitEthernet0/0
        ip address 10.10.10.1 255.255.255.0
''')
# establish multiple, simultaneous connections
ios_2.connect(alias = 'console', via = 'a')
ios_2.connect(alias = 'vty_1', via = 'vty')
# issue commands through each connection separately
ios_2.vty_1.execute('show running')
ios_2.console.execute('reload')
# creating connection pools
ios_2.start_pool(alias = 'pool', size = 2)
# use connection pool in multiprocessing paradigms
# each process will be allocated a connection - whenever one is available
def sleep(seconds):
    ios_2.pool.execute('sleep %s' % seconds)
import multiprocessing
p1 = multiprocessing.Process(target=sleep, args = (10, ))
p2 = multiprocessing.Process(target=sleep, args = (10, ))
p3 = multiprocessing.Process(target=sleep, args = (10, ))
p1.start(); p2.start(); p3.start()
p1.join(); p2.join(); p3.join()
```
pyATS is all about testing; and the absolute cornerstone in testing is the actual testscript. In pyATS, test scripts are written and executed through :ref:`AEtest Package <aetest_index>`.

Testscripts are :ref:`structured <aetest_script_structure>` Python files that contains/describes the testing you want to do. A clean, elegant testscript is scalable, and generates easy-to-read test results and logs.

