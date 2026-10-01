---
id: collect-261001-cisco/cisco/pyats-docs-getting-started-index-rst-at-550fa8883a43618d4b614dd521a9b39f1228a816-ciscotest-2
title: "Example"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/pyats-docs-getting-started-index-rst-at-550fa8883a43618d4b614dd521a9b39f1228a816-ciscotestautomation.md
source_anchor: ""
source_lines: [151, 284]
sha256: 896c468accd14c4ab05225827dc8c5bad26cce0f95b831dfdac372a6751aef2b
---

# Example

```
# Example
# -------
#
#   connectivity_check.py
from pyats import aetest
class CommonSetup(aetest.CommonSetup):
    @aetest.subsection
    def check_topology(self,
                       testbed,
                       ios1_name = 'ios-1',
                       ios2_name = 'ios-2'):
        ios1 = testbed.devices[ios1_name]
        ios2 = testbed.devices[ios2_name]
        # add them to testscript parameters
        self.parent.parameters.update(ios1 = ios1, ios2 = ios2)
        # get corresponding links
        links = ios1.find_links(ios2)
        assert len(links) >= 1, 'require one link between ios1 and ios2'
    @aetest.subsection
    def establish_connections(self, steps, ios1, ios2):
        with steps.start('Connecting to %s' % ios1.name):
            ios1.connect()
        with steps.start('Connecting to %s' % ios2.name):
            ios2.connect()
@aetest.loop(device=('ios1', 'ios2'))
class PingTestcase(aetest.Testcase):
    @aetest.test.loop(destination=('10.10.10.1', '10.10.10.2'))
    def ping(self, device, destination):
        try:
            result = self.parameters[device].ping(destination)
        except Exception as e:
            self.failed('Ping {} from device {} failed with error: {}'.format(
                                destination,
                                device,
                                str(e),
                            ),
                        goto = ['exit'])
        else:
            match = re.search(r'Success rate is (?P<rate>\d+) percent', result)
            success_rate = match.group('rate')
            logger.info('Ping {} with success rate of {}%'.format(
                                        destination,
                                        success_rate,
                                    )
                               )
class CommonCleanup(aetest.CommonCleanup):
    @aetest.subsection
    def disconnect(self, steps, ios1, ios2):
        with steps.start('Disconnecting from %s' % ios1.name):
            ios1.disconnect()
        with steps.start('Disconnecting from %s' % ios2.name):
            ios2.disconnect()
if __name__ == '__main__':
    import argparse
    from pyats.topology import loader
    parser = argparse.ArgumentParser()
    parser.add_argument('--testbed', dest = 'testbed',
                        type = loader.load)
    args, unknown = parser.parse_known_args()
    aetest.main(**vars(args))
```
This example uses Python argparse to parse command line arguments for a
testbed file input, and passes it to the script as the `testbed`
:ref:`script parameter <test_parameters>`. This is a good practice to do -
take arguments from command line makes your script more dynamic.

With your script written & saved, you can run it from the command line:

`bash$ python connectivity_check.py --testbed ios_testbed.yaml`
The `if __name__ == '__main__'` block in your testscript will invoke AEtest
to run the file's content when called from the command line, and when finished,
displays testcase results:

```
+------------------------------------------------------------------------------+
|                               Detailed Results                               |
+------------------------------------------------------------------------------+
 SECTIONS/TESTCASES                                                      RESULT
--------------------------------------------------------------------------------
.
|-- common_setup                                                         PASSED
|   |-- check_topology                                                   PASSED
|   `-- establish_connections                                            PASSED
|       |-- Step 1: Connecting to ios-1                                  PASSED
|       `-- Step 2: Connecting to ios-2                                  PASSED
|-- PingTestcase[device=ios1]                                            PASSED
|   |-- ping[destination=10.10.10.1]                                     PASSED
|   `-- ping[destination=10.10.10.2]                                     PASSED
|-- PingTestcase[device=ios2]                                            PASSED
|   |-- ping[destination=10.10.10.1]                                     PASSED
|   `-- ping[destination=10.10.10.2]                                     PASSED
`-- common_cleanup                                                       PASSED
    `-- disconnect                                                       PASSED
        |-- Step 1: Disconnecting from ios-1                             PASSED
        `-- Step 2: Disconnecting from ios-2                             PASSED
+------------------------------------------------------------------------------+
|                                   Summary                                    |
+------------------------------------------------------------------------------+
 Number of ABORTED                                                            0
 Number of BLOCKED                                                            0
 Number of ERRORED                                                            0
 Number of FAILED                                                             0
 Number of PASSED                                                             4
 Number of PASSX                                                              0
 Number of SKIPPED                                                            0
--------------------------------------------------------------------------------
```
This is the quickest way to see your testscript in action: everything is printed directly to screen, so you can edit, run, edit, and run again until your testing is tuned to perfection.

A :ref:`job <easypy_jobfile>` is a step above simply running testscripts as an executable and getting output in STDOUT. Job files enables the execution of testscripts as :ref:`tasks <easypy_tasks>` in standardized runtime environment, allowing testscripts to run in series or in parallel, and aggregates their logs and results together into a more manageable format.

In-effect, the engine around job files handle the typical boilerplate environment-setup, such as loading testbed files, through the use of :ref:`plugins <easypy_plugin>`.

A job must feature a `main()` method - this is entry point.

```
# Example: ios_job.py
# -------------------
#
#   a simple job file for the script above
from pyats.easypy import run
def main():
    # run api launches a testscript as an individual task.
    run('connectivity_check.py')
```
To launch a job, use `pyats`. The built-in testbed file handling plugin
accepts a `--testbed-file` argument, which automatically loads and parses the
provided testbed file into `testbed` parameter, and provide it to the
testscript. When launched, each testscript called by `run()` api inside the
job runs as a child process, and the contents inside its
```if __name__ == '__main__'``block is ignored. Add the```--html-logs`` argument
to enable generation of HTML log files - they are easier to read.

