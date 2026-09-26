---
id: collect-240926-datacamp/datacamp/ipython-ou-jupyter-2
title: "Assign the result to `ls`"
domain: datacamp
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["apache", "attention"]
source: docs/RAG/clean_en/datacamp/ipython-ou-jupyter.md
source_anchor: ""
source_lines: [77, 184]
sha256: 4467543ece42b4523c55c4fbdf5fdec020a610a13be32f8d6f4af1bcfc6118c2
---

# Assign the result to `ls`

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

