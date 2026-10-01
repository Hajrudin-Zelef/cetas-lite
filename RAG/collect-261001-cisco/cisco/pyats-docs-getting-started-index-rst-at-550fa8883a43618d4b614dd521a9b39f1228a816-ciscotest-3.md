---
id: collect-261001-cisco/cisco/pyats-docs-getting-started-index-rst-at-550fa8883a43618d4b614dd521a9b39f1228a816-ciscotest-3
title: "Example"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2018-03-15"]
keywords: []
source: docs/RAG/collect-261001-cisco/pyats-docs-getting-started-index-rst-at-550fa8883a43618d4b614dd521a9b39f1228a816-ciscotestautomation.md
source_anchor: ""
source_lines: [285, 367]
sha256: 06f9f51e46aeb08975a661cbf927fb63e6a69111a3233a962f279b74558b40fb
---

# Example

`bash$ pyats run job ios_job.py --testbed-file ios_testbed.yaml --html-logs````
+------------------------------------------------------------------------------+
|                                Easypy Report                                 |
+------------------------------------------------------------------------------+
pyATS Instance   : /path/to/pyats
Tcl-ATS Tree     :
Python Version   : cpython-3.4.1 (32bit)
CLI Arguments    : pyats run job ios_job.py --testbed-file ios_testbed.yaml
User             : joe
Host Server      : automation
Host OS Version  : Red Hat Enterprise Linux Server 6.6 Santiago (x86_64)
Job Information
    Name         : ios_job
    Start time   : 2018-03-15 00:24:05.847263
    Stop time    : 2018-03-15 00:24:17.066042
    Elapsed time : 0:00:11.218779
    Archive      : archive/18-Mar/ios_job.2018Mar15_00:24:04.zip
Total Tasks    : 1
Overall Stats
    Passed     : 4
    Passx      : 0
    Failed     : 0
    Aborted    : 0
    Blocked    : 0
    Skipped    : 0
    Errored    : 0
    TOTAL      : 4
Success Rate   : 100.00 %
+------------------------------------------------------------------------------+
|                             Task Result Summary                              |
+------------------------------------------------------------------------------+
Task-1: connectivity_check.commonSetup                                    PASSED
Task-1: connectivity_check.PingTestcase[device=ios1]                      PASSED
Task-1: connectivity_check.PingTestcase[device=ios2]                      PASSED
Task-1: connectivity_check.commonCleanup                                  PASSED
+------------------------------------------------------------------------------+
|                             Task Result Details                              |
+------------------------------------------------------------------------------+
Task-1: connectivity_check
|-- common_setup                                                          PASSED
|   |-- check_topology                                                    PASSED
|   `-- establish_connections                                             PASSED
|       |-- Step 1: Connecting to ios-1                                   PASSED
|       `-- Step 2: Connecting to ios-2                                   PASSED
|-- PingTestcase[device=ios1]                                             PASSED
|   |-- ping[destination=10.10.10.1]                                      PASSED
|   `-- ping[destination=10.10.10.2]                                      PASSED
|-- PingTestcase[device=ios2]                                             PASSED
|   |-- ping[destination=10.10.10.1]                                      PASSED
|   `-- ping[destination=10.10.10.2]                                      PASSED
`-- common_cleanup                                                        PASSED
    `-- disconnect                                                        PASSED
        |-- Step 1: Disconnecting from ios-1                              PASSED
        `-- Step 2: Disconnecting from ios-2                              PASSED
```
By default, the results of a job file is an archive: a zipped folder containing files describing the runtime environment, what was run, result XML files, and log files - eg, everything that was generated in your job's :ref:`runinfo folder <easypy_runinfo>`.

Congratulation, you now understand the basic building blocks of pyATS: job, script, and testbed files. Aside from the above, there are tons more features in pyATS left for you to explore. Checkout the rest of the documentation for all the other awesome features that can help you with your day-to-day testing!

Various pyATS script examples can be found in GitHub:

- **Feature Usage** : https://github.com/CiscoTestAutomation/examples
- **Solutions and Scripts** : https://github.com/CiscoTestAutomation/solutions_examples

Feel free to clone them into your workspace and run.

```
# Example
# -------
#
#   launching pyats from the command line
# activating pyats instance
# (if you have not yet activated)
[tony@jarvis:~]$ cd /ws/tony-stark/pyats
[tony@jarvis:pyats]$ source env.sh
# clone example folder
(pyats) [tony@jarvis:pyats]$ git clone https://github.com/CiscoTestAutomation/examples
(pyats) [tony@jarvis:pyats]$ cd examples
# start with executing the basic examples jobfiles.
# this is a basic example demonstrating the usage of a jobfiles,
# and running through a single aetest testscript.
(pyats) [tony@jarvis:examples]$ pyats run job basic/basic_example_job.py
```
