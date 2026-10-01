---
id: collect-261001-general-networking/general-networking/installationi
title: "Installationï"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/installationi.md
source_anchor: ""
source_lines: [1, 156]
sha256: fb7b716a48c6aa8de5159d9d485f293d364836ead3dcb0ddf4e05127c9405514
---

# Installationï

This topic describes how to install the **pyATS** within your system.

Warning

For all internal Cisco engineering users, skip this section, and refer to the engineering internal Wiki for detailed instructions on installing pyATS within Cisco.

| ï |  | 
|---|---|
| Type of user | Installation process | 
|---|---|
| Virtual Environment |  | 
| Docker |  | 

## Virtual Environmentï

pyATS development team recomments you to always develop and run Python scripts with a Python virtual environment. This section describes how to check your Python version, and creating a virtual environment using your available binary.

### Python Versionï

Note

Make sure your system has a supported version of Python installed:

- Python 3.7.x
- Python 3.8.x
- Python 3.9.x
- Python 3.10.x

To check your installed version:

*Result*: The system returns the installed version number:

Tip

`pyenv` a great utility for managing multiple Python versions within your
system.

### Create Virtual Environmentï

A Python Virtual Environment is simply a directory (folder). Within this virtual environment, you install the pyATS and pyATS Library packages, dependencies, and libraries, including everything else you need to run the system.

Note

In our examples, we use the directory `pyats`, but you can give your directory a different name.

1. Create a new directory:
2. Go to the new directory:
3. Initialize a virtual environment in this directory: *Result* : This creates a project âfolderâ (space) within the current directory. The folder keeps all dependencies, features, and components together in one place.
4. Activate the virtual environment:

*Result*: The system displays the directory in parentheses before the command prompt:

When you install the pyATS ecosystem within this virtual environment, the packages remain separate from those in other project spaces.

Hint

When youâre done with your pyATS session, you can close the terminal window or exit the environment:


### Pip Installï

1. If you havenât already done so, activate your virtual environment: *Result* : The system displays the directory in parentheses before the command prompt:
2. Upgrade pip with the latest setup tool packages:
3. Install pyATS and the pyATS Library, using the options described in the following table. ï Installation option Command Includes Full `$ pip install pyats[full]`
  - All pyATS and pyATS Library infrastructure
  - pyATS Library network automation packages
  - Optional extras (templates and the Robot Framework plug-in)
 Standard `$ pip install pyats[library]`
  - All pyATS and pyATS Library infrastructure
  - pyATS Library network automation packages
 Core pyATS `$ pip install pyats`
  - pyATS bare-bone infrastructure
 Robot Framework `$ pip install pyats[robot]`
  - Optional Robot Framework package
  - `pyats.robot` package (contains pyATS-specific keywords)
 Template command `$ pip install pyats[template]`
  - Enables use of the `template` command, which prompts you for input at runtime
 Individual packages (for three-part patch versions) `$ pip install <package_name>`
  - The specified package
 Hint Give the installer a few minutes to finish. *Result* : Youâre ready to start using pyATS and the pyATS Library!Note If you see warning messages, or the installation fails, first check your system requirements, especially your Linux and Python versions. If you need more help, contact us at 
pyats-support-ext@cisco.com.
4. To test the installation, from the current (pyATS) directory, clone the Git examples repository: git clone https://github.com/CiscoTestAutomation/examples
5. Run the following example: pyats run job examples/basic/basic_example_job.py

Or, for DevNet community users who want to receive an email summary:

pyats run job examples/basic/basic_example_job.py --mailto <address>
*Result*: pyATS runs three sample test cases, displays a summary of the results, and emails you the summary.


### Update Environmentï

On the last Tuesday of the month, the team releases a new version of **pyATS**.
This topic describes how to get the latest changes.

To upgrade the pyATS and pyATS Library infrastructure, and any or all of the feature libraries and components, run the relevant upgrade command **from your virtual environment**.

Tip

You can find the latest information about releases on Twitter at #pyATS.

For information about all things pyATS, see our discussion on Webex Teams.

You can check and upgrade your pyATS installation straigth from the command line:

```
# to check your current pyats version
(pyats)$ pyats version check
# to check if any packages are out-dated
(pyats)$ pyats version check --outdated
# to update version
(pyats)$ pyats version update
```
Otherwise, you can also update the packages manually using Pip:

| ï |  |  | 
|---|---|---|
| Upgrade option | Command | Includes | 
|---|---|---|
| Full | `$ pip install pyats[full] --upgrade` |  | 
| Standard | `$ pip install pyats[library] --upgrade` |  | 
| Core pyATS | `$ pip install pyats --upgrade` |  | 
| Robot Framework | `$ pip install pyats[robot] --upgrade` |  | 
| Template command | `$ pip install pyats[template] --upgrade` |  | 
| Individual packages (for three-part patch versions) | `$ pip install <package_name> --upgrade` |  | 

*Result*: The installer checks for and upgrades any dependencies, and gives you the latest version of the pyATS and pyATS Library core and library packages. To check the version:

*Result*: The system displays a list of the packages and the installed versions.

Attention

The major and minor versions must all match. Itâs okay if the patch version varies.

See alsoâ¦

## Using Dockerï

If you know how to use Docker, you can work with our pre-built docker image, which includes both pyATS and the pyATS Library. You can find the image and instructions at https://hub.docker.com/r/ciscotestautomation/pyats.

## Examples Repositoryï

Weâve provided some examples to help you start using the pyATS Library for some simple scenarios that demonstrate how the pyATS Library works.

- To clone the Git repository from your virtual environment:
- To download the Git repository from a browser: 
  - Go to https://github.com/CiscoTestAutomation/examples.
  - Select **Clone or download** .
  - Select **Open in Desktop** to download and use the GitHub Desktop app, or**Download Zip** to download and extract a zip file.

*Result*: You now have the example files stored in the `examples` directory.


See alsoâ¦
