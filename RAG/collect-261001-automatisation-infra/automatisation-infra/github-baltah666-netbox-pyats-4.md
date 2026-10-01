---
id: collect-261001-automatisation-infra/automatisation-infra/github-baltah666-netbox-pyats-4
title: "Description: This script generates a testbed file based on the Netbox data"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/github-baltah666-netbox-pyats.md
source_anchor: ""
source_lines: [445, 503]
sha256: 737202fd8bff9a6db1155624134c1bc77978027af0d04884d9e02a85606ff20a
---

# Description: This script generates a testbed file based on the Netbox data

```
python ospf_neighbor_table.py
OSPF Neighbor Table for CSR1
+------------------+----------+-------------+---------+
|    Interface     | Neighbor |   Address   |  State  |
+------------------+----------+-------------+---------+
| GigabitEthernet2 | 2.2.2.2  | 192.168.1.2 | FULL/DR |
+------------------+----------+-------------+---------+
```
So far we have only scratched the surface of what you can do with pyATS, and test scripts will take your automated testing to the next level. Test scripts are written and executed with the AEtest Python package

I have included an example test script called `connectivity_check.py`, which is adapted from the example script in the official documentation. This script tests that both CSR1 and CSR2 can ping each other's loopback addresses. Test scripts are structured Python files that contain/describe the testing you want to do. A clean, elegant test script is scalable, and generates easy-to-read test results and logs.

To run the example test script, you must also specify the testbed file that is to be used, for example:

```
python connectivity_check.py --testbed testbed.yaml
```
Once the test script has been run then the results are displayed in the output:

```
2024-04-23T11:50:16: %AETEST-INFO: +------------------------------------------------------------------------------+
2024-04-23T11:50:16: %AETEST-INFO: |                               Detailed Results                               |
2024-04-23T11:50:16: %AETEST-INFO: +------------------------------------------------------------------------------+
2024-04-23T11:50:16: %AETEST-INFO:  SECTIONS/TESTCASES                                                      RESULT   
2024-04-23T11:50:16: %AETEST-INFO: --------------------------------------------------------------------------------
2024-04-23T11:50:16: %AETEST-INFO: .
2024-04-23T11:50:16: %AETEST-INFO: |-- common_setup                                                          PASSED
2024-04-23T11:50:16: %AETEST-INFO: |   |-- check_topology                                                    PASSED
2024-04-23T11:50:16: %AETEST-INFO: |   `-- establish_connections                                             PASSED
2024-04-23T11:50:16: %AETEST-INFO: |       |-- Step 1: Connecting to CSR1                                    PASSED
2024-04-23T11:50:16: %AETEST-INFO: |       `-- Step 2: Connecting to CSR2                                    PASSED
2024-04-23T11:50:16: %AETEST-INFO: |-- PingTestcase[device=device1]                                          PASSED
2024-04-23T11:50:16: %AETEST-INFO: |   |-- ping[destination=192.168.1.1]                                     PASSED
2024-04-23T11:50:16: %AETEST-INFO: |   `-- ping[destination=192.168.1.2]                                     PASSED
2024-04-23T11:50:16: %AETEST-INFO: |-- PingTestcase[device=device2]                                          PASSED
2024-04-23T11:50:16: %AETEST-INFO: |   |-- ping[destination=192.168.1.1]                                     PASSED
2024-04-23T11:50:16: %AETEST-INFO: |   `-- ping[destination=192.168.1.2]                                     PASSED
2024-04-23T11:50:16: %AETEST-INFO: `-- common_cleanup                                                        PASSED
2024-04-23T11:50:16: %AETEST-INFO:     `-- disconnect                                                        PASSED
2024-04-23T11:50:16: %AETEST-INFO:         |-- Step 1: Disconnecting from CSR1                               PASSED
2024-04-23T11:50:16: %AETEST-INFO:         `-- Step 2: Disconnecting from CSR2                               PASSED
2024-04-23T11:50:16: %AETEST-INFO: +------------------------------------------------------------------------------+
2024-04-23T11:50:16: %AETEST-INFO: |                                   Summary                                    |
2024-04-23T11:50:16: %AETEST-INFO: +------------------------------------------------------------------------------+
2024-04-23T11:50:16: %AETEST-INFO:  Number of ABORTED                                                            0 
2024-04-23T11:50:16: %AETEST-INFO:  Number of BLOCKED                                                            0 
2024-04-23T11:50:16: %AETEST-INFO:  Number of ERRORED                                                            0 
2024-04-23T11:50:16: %AETEST-INFO:  Number of FAILED                                                             0 
2024-04-23T11:50:16: %AETEST-INFO:  Number of PASSED                                                             4 
2024-04-23T11:50:16: %AETEST-INFO:  Number of PASSX                                                              0 
2024-04-23T11:50:16: %AETEST-INFO:  Number of SKIPPED                                                            0 
2024-04-23T11:50:16: %AETEST-INFO:  Total Number                                                                 4 
2024-04-23T11:50:16: %AETEST-INFO:  Success Rate                                                            100.0% 
2024-04-23T11:50:16: %AETEST-INFO: --------------------------------------------------------------------------------
```
pyATS Jobs take test scripts to another level and allow the execution of test scripts as tasks. This is beyond the scope of this webinar/repo but you can read all about pyATS jobs in the official docs

If you've got this far, first of all well done and thanks for reading! If you have any feedback or suggestions then please feel free raise a Github issue in this project repo.
