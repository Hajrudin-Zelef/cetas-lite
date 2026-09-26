---
id: collect-240926-datacamp/datacamp/ipython-ou-jupyter-1
title: "Assign the result to `ls`"
domain: datacamp
role: reference
task: reference
actors: ["Google"]
dates: ["2010-10", "2011-12-21", "2013-03-23"]
keywords: ["funding", "research"]
source: docs/RAG/clean_en/datacamp/ipython-ou-jupyter.md
source_anchor: ""
source_lines: [1, 76]
sha256: 307db177baa8bf6dfae36977a1d1951b0d712ded42ba37bae5cfd17f8c896fa6
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

