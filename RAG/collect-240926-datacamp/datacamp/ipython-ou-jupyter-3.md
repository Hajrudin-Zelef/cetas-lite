---
id: collect-240926-datacamp/datacamp/ipython-ou-jupyter-3
title: "Assign the result to `ls`"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean_en/datacamp/ipython-ou-jupyter.md
source_anchor: ""
source_lines: [185, 290]
sha256: 5f9a7c4b3e751646263a9ee7e3393566c2705b23631f74e8db58f2ac2d54cdcd
---

# Assign the result to `ls`

The IPython kernel uses the % symbol (which is not a valid unary operator in Python). Lines beginning with %% mark a “cell magic”: they take as arguments not only the rest of the line, but also all the following lines in the current block. Cell magics can freely modify the input they receive, which doesn’t even need to be valid Python. They receive the entire block as a single string.

Magics are specific to the kernels that provide them and are designed to make your work in Jupyter Notebook more interactive. Their availability depends on the kernel developers and varies from one kernel to another. As you may have gathered: magics are a kernel feature.

When you use the Python backend of Jupyter, IPython (the kernel), here are a few tips for accessing features that speed up, simplify, and enrich your programming. The list is not exhaustive. Consult the complete list of built-in magics.

#### Visualization

A major feature of the IPython kernel is the display of graphs produced by code cells. It is designed to work seamlessly with the matplotlib library. To enable it, use the `%matplotlib` magic.

By default, the graph opens in a separate window. You can also specify a backend (inline, qt, etc.) to display plots inline or via another GUI interface. More info here.

#### Navigating the file system

The kernel's magics also offer ways to navigate your file system. The `%cd` and `%bookmark` magics allow you to change directories or add bookmarks to access frequently used folders more quickly.

#### Accessing the debugger

You can invoke a Python debugger with `%pdb` on every uncaught exception. It guides you through the portion of code at the origin of the error to speed up analysis.

The `%run` magic with the -d option executes a script under the control of the Python debugger and automatically places initial breakpoints. The `%debug` magic offers even more direct access.

#### IPython extensions

The `%load_ext` magic loads an IPython extension by its module name. IPython extensions are Python modules that modify the behavior of the shell: they can register magics, define variables, and more generally enrich the user space to offer new functions in cells. Examples:

- `%load_ext oct2py.ipython`: transparently call M-files and Octave functions from Python,
- `%load_ext rpy2.ipython`: interface to R embedded in a Python process,
- `%load_ext Cython`: use the Python→C compiler,
- `sympy.init_printing()`: automatic pretty rendering of Sympy Basic objects,
- `%load_ext fortranmagic`: use Fortran in your interactive sessions.

… And many others! You can create and publish your own IPython extensions on PyPI: there are therefore many community extensions and magics. Example: `ipython_unittest`, and also consult this index of extensions.

Also keep an eye on sparkmagic, a set of tools for working interactively with remote Spark clusters via Livy (Spark REST server) in Jupyter notebooks. The `sparkmagic` library provides a `%%spark` magic to easily run code against a remote Spark cluster from a classic IPython notebook.

```
# Load in sparkmagic
%load_ext sparkmagic.magics
 
# Set the endpoint
%manage_spark
 
# Ask for help
%spark?
```
Go here for more examples of using these magics with a Spark cluster.

Note that IPython offers two other magics, `%reload_ext` and `%unload_ext`, to reload and unload extensions directly from Jupyter Notebook.

#### Different kernels, other magics

In other languages, the syntax element used by magics may already have a meaning. The R kernel, IRkernel, has no magic system. To run bash commands, for example, you will use R functions such as `system()` to invoke OS commands, for example `system("head -5 *.csv", intern=TRUE)`. By including the `intern` argument, you indicate that you want to capture the command's output as an R character vector. To display Markdown, use `display_markdown()`, to which you pass the Markdown code as a character vector. Similarly, the Julia kernel IJulia does not use "magics." Other syntaxes, more natural in Julia, work outside IJulia cells and are often more powerful. However, the IJulia developers have provided that when entering an IPython magic in an IJulia cell, a help message indicates how to achieve an equivalent effect in Julia, when possible.

For example, the equivalent of IPython's `%load` in IJulia is `IJulia.load()`.

Conversely, some kernels such as IScala (Scala) support magics, similarly to IPython, but with a set adapted to Scala and the JVM. A magic begins with `%` followed by an identifier and, optionally, input. Notable examples:

```
# Type Information
%type 1
 
# Library Management
%libraryDependencies
%update
```
As mentioned above, the `sparkmagic` library also provides Scala and Python kernels to automatically connect to a remote Spark cluster, run code and SQL queries, manage the configuration of your Livy server and your Spark jobs, and generate automatic visualizations — without writing additional code.

For example, you can launch SparkSQL queries with `%%sql` or access Spark application information and logs via the `%%info` magic.

If you use another kernel and wonder whether it supports magics, know that some kernels rely on the metakernel project and in many cases use the same magics as the IPython kernel. The list of metakernel magics is here. metakernel is a Jupyter/IPython kernel template including basic magics.

Examples:

- The MATLAB kernel `matlab_kernel`,
- The Octave kernel `octave_kernel`,
- The Java9 kernel `java9_kernel`,
- The Wolfram kernel `wolfram_kernel`,
- The SAS kernel. … And many others!

Thus, with the MATLAB kernel, you have for example the following magics:

```
Available line magics:
%cd  %connect_info  %download  %edit  %get  %help  %html  %install  %install_magic  %javascript  %kernel  %kx  %latex  %load  %ls  %lsmagic  %magic  %parallel  %plot  %pmap  %px  %python  %reload_magics  %restart  %run  %set  %shell  %spell
 
Available cell magics:
%%debug  %%file  %%help  %%html  %%javascript  %%kx  %%latex  %%processing  %%px  %%python  %%shell  %%show  %%spell
```
If you compare with the magics available by default in the IPython kernel, you will notice some overlaps:

```
Available line magics:
%alias  %alias_magic  %autocall  %automagic  %autosave  %bookmark  %cat  %cd  %clear  %colors  %config  %connect_info  %cp  %debug  %dhist  %dirs  %doctest_mode  %ed  %edit  %env  %gui  %hist  %history  %killbgscripts  %ldir  %less  %lf  %lk  %ll  %load  %load_ext  %loadpy  %logoff  %logon  %logstart  %logstate  %logstop  %ls  %lsmagic  %lx  %macro  %magic  %man  %matplotlib  %mkdir  %more  %mv  %notebook  %page  %pastebin  %pdb  %pdef  %pdoc  %pfile  %pinfo  %pinfo2  %popd  %pprint  %precision  %profile  %prun  %psearch  %psource  %pushd  %pwd  %pycat  %pylab  %qtconsole  %quickref  %recall  %rehashx  %reload_ext  %rep  %rerun  %reset  %reset_selective  %rm  %rmdir  %run  %save  %sc  %set_env  %store  %sx  %system  %tb  %time  %timeit  %unalias  %unload_ext  %who  %who_ls  %whos  %xdel  %xmode
 
Available cell magics:
%%!  %%HTML  %%SVG  %%bash  %%capture  %%debug  %%file  %%html  %%javascript  %%js  %%latex  %%perl  %%prun  %%pypy  %%python  %%python2  %%python3  %%ruby  %%script  %%sh  %%svg  %%sx  %%system  %%time  %%timeit  %%writefile
```
To distinguish magics specific to IPython from those usable elsewhere, ask yourself a simple question: is the functionality specific to Python or general to the language you are using?

For example, `%pdb` (Python debugger) or `%matplotlib` are specific to Python and make no sense with a JavaScript kernel. On the other hand, changing directory with `%cd` is fairly general and should work in any environment, provided the kernel handles magics.

### Converting and formatting notebooks?

Converting and formatting notebooks fall under the Jupyter ecosystem. Two typical tools: `nbconvert` and `nbformat`.

