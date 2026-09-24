---
id: collect-240926-datacamp/datacamp/ipython-ou-jupyter
title: "Assign the result to `ls`"
domain: datacamp
role: reference
task: reference
actors: ["Google"]
dates: ["2010-10", "2011-12-21", "2013-03-23"]
keywords: ["apache", "attention", "distribution", "funding", "parameters", "research", "training"]
source: docs/RAG/clean_en/datacamp/ipython-ou-jupyter.md
source_anchor: ""
source_lines: [1, 331]
sha256: 889f3b4f61984a1c9b9871f1ba8d9b88fa34cedb068f8b518cfb7a75712e06cd
---

# Assign the result to `ls`

<!-- source: https://www.datacamp.com/fr/blog/ipython-or-jupyter -->

*A big thank you to Brian Granger, Fernando Pérez, and Robert Kern for their contributions to this article!*

For learners and seasoned data scientists alike, Jupyter Notebook is one of the flagship tools of data science: its interactive environment is ideal for teaching, learning, sharing work with peers, and ensuring research reproducibility. Yet, when discovering the notebook, you will very often encounter IPython.

In some cases, the two seem synonymous, and you will agree that it quickly becomes confusing as soon as you want to dig deeper: do "magics" belong to Jupyter or IPython? Are saving and loading notebooks features of IPython or Jupyter?

The list of questions can be long.

The goal of this article is to clearly lay out some fundamental differences between the two, starting from their origins to explain their connections, then reviewing features specific to one or the other, to help you better tell them apart.

Also consider checking out DataCamp's ultimate guide to Jupyter Notebook for tips, best practices, examples, and much more.

## The origins of IPython / Jupyter

To properly understand what Jupyter Notebook is and how it differs from IPython, it is useful to revisit the place of these two projects in the history (and future) of computational notebooks.

### The beginnings of computational notebooks: MATLAB, Mathematica, and Maple

In the mid-1980s, MATLAB was launched by The MathWorks, founded by Jack Little, Steve Bangert, and Cleve Moler.

In the late 1980s, 1987 to be precise, Theodore Gray began working on what would become Mathematica's notebook interface, published a year later. This graphical interface allows users to interactively create and edit notebook documents containing nicely formatted code, text, and many features such as typeset formulas, graphics, GUI components, tables, and sounds. It includes the classic functions of a word processor, including real-time multilingual spell checking. Documents can be played in slideshow mode for presentations.

The structure of these notebooks was based on a hierarchy of cells facilitating the layout and sections of documents — a principle that is found today in Jupyter notebooks.

In 1989, Maple also introduced its first notebook-style interface, included with version 4.3 for Macintosh. Versions for X11 and Windows followed in 1990. These precursor notebooks inspired and laid the foundations for what would later be called "data science notebooks."

### The rise of data science notebooks

Many computational notebooks emerged after Maple and Mathematica. This section, however, focuses on those that contributed to the rise of data science notebooks, some of which remain very popular among data scientists, beginners and experts alike.

#### Sage Notebook

The Sage notebook, as a browser-accessible system, appeared in the mid-2000s. In 2007, a new, more powerful version was released, with user account management and the ability to publish documents. Its interface resembles Google Docs, as the layout of the Sage notebook was inspired by Google Notebooks.

The creators of Sage confirmed that they were avid users of Mathematica notebooks and Maple worksheets. Other factors played a role in Sage's development: on the one hand, close exchanges with the IPython team, since Sage in terminal mode relied on IPython; on the other hand, "an abortive attempt" by two students to provide a graphical interface for IPython; and finally, the rise of "AJAX" (Asynchronous JavaScript And XML) web applications, which avoid reloading the entire page with each action.

For example, when you submit a form on a website: without AJAX, you are redirected to a new page returned by the server. With AJAX, JavaScript sends the request, retrieves the response, and updates the screen, without reloading or redirecting.

Among well-known AJAX applications: Gmail, Google Maps, Facebook, Twitter, etc. Almost everything now relies on AJAX.

#### IPython and Jupyter Notebook

In late 2001, about twenty years after Guido van Rossum began working on Python at CWI (Netherlands), Fernando Pérez started developing IPython. Like Sage and many other projects, IPython was strongly influenced by Mathematica notebooks and Maple worksheets.

In 2005, a first notebook attempt emerged with Wx, a widget toolkit for creating cross-platform graphical interfaces. Two students from Google Summer of Code worked on a prototype under the direction of Robert Kern and Fernando Pérez. Robert then pushed the work further. It was less a notebook prototype than a cleanup of IPython's core to facilitate writing a wxPython frontend for the IPython shell. These refactors helped make a proper notebook possible, and some were gradually integrated into IPython when the effort, this time successful, came to fruition.

The second prototype of an IPython notebook was made in the summer of 2006 by Min Ragan-Kelley, advised by Brian Granger. Based on the web with an SQL backend, it was eventually abandoned as too complex for the web technologies of the time.

The third prototype arrived in October 2010, developed by a third party during a hackathon of a few days. Finally, in the spring-summer of 2011, Brian Granger worked full-time on a web notebook prototype, building on his 2010 work, when he created with Fernando Pérez the IPython kernel architecture and message specification, as well as PyZMQ with Min Ragan Kelley.

PyZMQ is a library that provides Python bindings for ZeroMQ: necessary for parallel computing features, the Qt console, and the IPython notebook.

PyZMQ and websockets are the key technologies that make the notebook possible. In the fall, other contributors (including Matthias Bussonnier, Min, and Fernando) join the effort. On December 21, 2011, the first version of IPython Notebook (0.12) is released.

In the following years, the team receives awards (for example, the Advancement of Free Software for Fernando Pérez on March 23, 2013, the Jolt Productivity Award) and funding (Alfred P. Sloan Foundation, among others).

In 2014, Project Jupyter is finally born, stemming from IPython.

The last version of IPython before the split grouped into a single repository the interactive shell, the notebook server, the Qt console, etc. The project was becoming vast, with increasingly distinct components.

But size is not the only reason for the creation of Jupyter: from 2011 to 2014, the IPython Notebook begins to work with other languages. Julia arrives first, then R. The fact that the "IPython" Notebook — suggesting a notebook exclusively for Python — also works with kernels for Julia and R is confusing.

Why keep "IPython" if other languages are supported?

The name "Jupyter" is inspired by the main open languages for science (Julia, Python, and R). Even though it may seem like an acronym, it never meant that other languages were not welcome. Above all, it better represented the project and paid homage to its scientific roots.

After the creation of Jupyter, the language-agnostic elements of IPython — notebook format, message protocol, Qt console, notebook web application, etc. — are transferred to the Jupyter project. Jupyter's main GitHub organization is located here.

In the Jupyter and IPython communities, this is called "The Big Split."

IPython now has only two roles: serving as the Python backend for the Jupyter notebook (the kernel) and providing an interactive Python shell. But that's not all: the IPython ecosystem also includes a parallel computing framework, which we will discuss later.

Like IPython, Project Jupyter is an umbrella name for several projects: the three main applications are the Notebook, a Console, and a Qt console, but there are also sub-projects like JupyterHub for deploying notebooks, nbgrader for teaching, etc. An overview of the Jupyter architecture is available here.

This evolution explains the confusion of many Pythonistas regarding IPython and Jupyter: one stems from the other, and recently. Some still struggle to use the right terms. Even more confusing is the evolution itself: the common heritage leads to a notable overlap in the features of IPython and Jupyter Notebook, sometimes difficult to untangle. The following sections will clarify these distinctions. To learn more about the history of IPython, read the accounts by Fernando Pérez and William Stein about their notebooks.

#### R Notebooks

R Markdown and Jupyter Notebook share the goal of a reproducible workflow, weaving code, results, and text into a single document, with support for interactive widgets and multi-format export.

But they also differ: R Markdown favors reproducible batch execution, plain text representation, version control, production outputs, and the use of the same editor/tooling as for R scripts. Notebooks, on the other hand, display results alongside the code, cache outputs between sessions, and facilitate sharing code and results in a single file. They emphasize interactive execution. They do not use a plain text representation, but a structured data format, such as JSON.

This is what motivates RStudio's "notebook" application: it combines the strengths of R Markdown with those of computational notebooks.

To learn how to work with R notebooks and understand precisely the differences between Jupyter and R Markdown notebooks in terms of sharing, project management, version control, etc., read the DataCamp article Jupyter and R: notebooks with R.

#### Other data science notebooks

Other notebooks deserve your attention in data science. In recent years, many alternatives have emerged: Beaker Notebook, Apache Zeppelin, Spark Notebook, DataBricks Cloud, etc., but also tools like the Rodeo IDE or nteract, which make your analyses interactive and reproducible. Note that nteract stands out by relying on the Jupyter architecture (protocols and formats).

### The future of notebooks

Notebooks are here to stay. The new generation of Jupyter Notebook was recently introduced: JupyterLab. The Notebook application not only integrates notebook support, but also a file manager, a text editor, a terminal, a Jupyter process monitor, an IPython cluster manager, and a help pager.

Nothing new? In reality, JupyterLab allows you to leverage all these building blocks of interactive computing in novel ways.

Read more here.

The arsenal of Jupyter Notebook tools has grown organically, guided by the needs of users and developers. JupyterLab brings a next-generation architecture to unite them, with a flexible and responsive interface, offering a user-driven layout.

## IPython or Jupyter?

The evolution of the project and the "Big Split" that follows lay the foundation for understanding the real differences between the two. But since they are intimately linked, doubt sometimes persists about what belongs to what.

The following section reviews features belonging either to the IPython ecosystem or to the Jupyter project.

It’s up to you to identify the right answer and learn more about each function!

### Kernels?

Even though kernels feature prominently in the Jupyter Notebook application, the first tool to use a complete kernel and protocol was the Qt console, which predates the Notebook. Today, kernels are used flexibly by the Notebook, historically as well as architecturally, and belong to Jupyter rather than to the Notebook alone. In other words, kernels are much more than a feature: they are a central abstraction of the Jupyter architecture, used by non-notebook tools such as the text console, the Qt console, O’Reilly’s Thebe, Kernel Gateway, and nteract’s Hydrogen editor. Finally, in JupyterLab, a kernel can also connect to anything, from a notebook to a web console, or even a simple text file.

A kernel is a program that executes and introspects the user’s code: it handles computation and communication with front-end interfaces such as notebooks. The Jupyter Notebook application offers three main kernels: IPython, IRkernel, and IJulia.

Since “Jupyter” is inspired by Julia, Python, and R, that’s no surprise. The IPython kernel is maintained by the Jupyter team, a logical consequence of the project’s evolution.

However, you can run many other languages in the Jupyter Notebook application: Scala, JavaScript, Haskell, Ruby, etc. These are community-maintained kernels.

### Deploying notebooks?

Deploying notebooks is typically a topic you’ll encounter with Jupyter. Several packages in the Jupyter ecosystem can help you with that.

Here are a few:

- `docker-stacks`: stacks of Jupyter applications and kernels as Docker containers.
- `ipywidgets`: interactive HTML & JavaScript widgets (sliders, checkboxes, text fields, charts…) for the Jupyter architecture, connecting front-end controls and the kernel.
- `jupyter-drive`: allows IPython to use Google Drive for file management.
- `jupyter-sphinx-theme`: adds a Jupyter Sphinx theme to your notebooks to create polished and smart documentation.
- `kernel_gateway`: a web server that supports various mechanisms for launching and communicating with Jupyter kernels. Use cases here.
- `nbviewer`: for sharing your notebooks. Gallery here.
- `tmpnb`: creates temporary Jupyter Notebook servers via Docker. Try it here.
- `traitlets`: a framework that gives Python classes attributes with type checking, dynamic default values, and “on change” callbacks. Also useful for configuration (files, command-line arguments). traitlets powers the configuration system of IPython and Jupyter as well as the declarative API of IPython interactive widgets.

### Using the system shell?

IPython can be adapted for system-shell-like use thanks to the shell “escape”: lines beginning with `!` are passed directly to the shell. For example, !ls will run `ls` in the current directory. You can assign the result of a system command to a Python variable with the syntax `myfiles=!ls`. To explicitly display the result of ls as a list of strings without assigning it, use two exclamation marks (`!!ls`) or the `%sx` magic command without assignment.

```
# Assign the result to `ls`
ls = !ls
 
# Explicit `ls`
!!ls
 
# Or with magics
%sx
 
# Assign magics result
ls = %sx
```
Note that `!!` commands cannot be assigned to a variable, whereas the result of a magic (if it returns a value) can be.

IPython also lets you expand the value of Python variables in system calls: simply wrap variables or expressions in braces (`{}`). In a command with `!` or `!!`, any Python variable prefixed with `$` is expanded. In the snippet below, you display the argv attribute of the sys variable. You can also use the `$`/`$$` syntaxes to reference Python variables from system output, then use them in your scripts.

To pass a literal `$` to the shell, use a double `$$`. You’ll need this to access environment variables like $PATH:

```
# Import and initialize
import math
x = 4
 
# System call with variable
!echo {math.factorial(x)}
 
# Expand a variable
!echo $sys.argv
 
# Use $$ for a literal $
!echo "A system variable: $$HOME"
```
Learn more here.

Besides IPython, other kernels also offer special syntaxes for running lines in the shell! These aren’t necessarily “magics” in the strict sense, since the name and richness can vary depending on the implementation.

You can also define **aliases** for system commands: these are shortcuts to bash commands. An alias is a tuple: (“showTheDirectory”, “ls”). Run `%alias?` for more info!

**Tip**: use `%rehashx` to load your entire $PATH as IPython aliases.

### Magics?

If you’ve gone through DataCamp’s ultimate guide to Jupyter Notebook or if you’ve already used Jupyter, you’re probably familiar with “magic commands.” Magics rely on a syntax element that is invalid in the underlying language and a word that evokes a command. Under the hood, they are Python functions.

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

The first converts notebooks to other formats to present information in familiar forms, publish work, integrate notebooks into articles, collaborate, and share more widely.

The second defines the format of Jupyter notebooks: simple JSON documents containing metadata (kernel, language info), the format version (major and minor), and the cells where text, code, etc. are stored.

### Saving and loading notebooks?

Saving and loading notebooks are features of the Jupyter Notebook application. You can load notebooks (.ipynb files) created by others by downloading them and then opening them in the Jupyter application. Concretely, create a new notebook, then open the file via the "File" tab, "Open", and select your downloaded notebook.

Conversely, save your own notebooks via the same "File" tab by choosing "Download as" to retrieve the file, or "Save and Checkpoint" to set a milestone. Very useful for a lightweight form of version control and returning to a previous state if needed. Of course, your changes are saved automatically every few minutes, so the action is not always necessary.

Note that you can also preserve an original notebook by working on a copy and saving your changes there!

### Keyboard shortcuts and multi-cursor?

Selecting multiple cells, showing/hiding outputs, inserting new cells, etc.: all actions with keyboard shortcuts in Jupyter Notebook. The list is accessible via the top menu: "Help" tab, then "Keyboard Shortcuts".

Multi-cursor support is also a Jupyter Notebook feature!

### Distributed parallel computing?

The parallel computing network was part of IPython, but since version 4.0, it has been split into a standalone package, `ipyparallel`. This package groups CLI scripts to control Jupyter clusters.

Although separate, it remains a powerful component of the IPython ecosystem, often underestimated: instead of a single Python kernel, you can start many distributed kernels across multiple machines.

Typical use cases: running a model a very large number of times to estimate the distribution of its outputs or their sensitivity to input parameters. When runs are independent, you save time by parallelizing them across multiple machines in a cluster. Think of distributed model training or simulations.

### Terminal?

This feature belongs to the Jupyter ecosystem: there is Jupyter Console and a Jupyter terminal application. From the beginning, however, IPython referred to the initial interactive terminal for Python, offering an enriched REPL loop, particularly suited to scientific computing. It was the reference before 2011, when the Notebook introduced a modern and powerful web interface for Python.

There was also an IPython console, which launched two processes: the IPython terminal shell and the default profile/kernel (Python if unspecified). The IPython console is now deprecated and, to find similar usage, you must use Jupyter Console, a terminal frontend for Jupyter kernels. It takes up the interactive IPython terminal experience, but allows connecting to any Jupyter kernel, not just IPython. You can thus test any installed Jupyter kernel without opening a full Notebook. The Console allows terminal interactions with IJulia, IRKernel, etc.

Finally, the Jupyter Notebook application also includes a Terminal application: a simple bash shell that runs in your browser. You can easily find it by starting the application and then creating a new terminal via the dropdown menu.

### Qt Console?

The Qt console was part of IPython, but it joined the Jupyter project. It is a lightweight application that looks like a terminal while offering GUI-specific improvements: inline figures, true multiline editing with syntax highlighting, graphical contextual help, etc. The Qt console can use any Jupyter kernel.

## Conclusion

This article complements DataCamp's ultimate guide and traces in more detail the history of computational notebooks, while presenting key features of the IPython and Jupyter projects to better understand their evolution and differences. The boundary is not always clear if one forgets the historical perspective: some gray areas remain, difficult to classify.
