---
id: vague2-datacamp/datacamp/ipython-or-jupyter
title: "IPython ou Jupyter ?"
domain: datacamp
role: reference
task: article
actors: ["Google"]
dates: ["2010-10", "2011-12", "2026-09-23"]
keywords: []
source: docs/RAG/Collect RAG Vague 2/02_datacamp/ipython-or-jupyter.md
source_anchor: ""
source_lines: [1, 63]
sha256: 7919d8724a4409049fb0d0921b777ed49cef9e5161896e4640731eb7e0ccce99
---

# IPython ou Jupyter ?

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/ipython-or-jupyter
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article clarifies the relationship and differences between IPython and Jupyter, starting from their origins. It traces computational notebooks back to MATLAB (mid-1980s), Mathematica's notebook interface (1987–1988, with cell hierarchies that Jupyter still uses), and Maple's notebook interface (1989). It then covers the rise of data science notebooks: Sage Notebook (mid-2000s, browser-accessible, inspired by Google Docs layout and AJAX apps like Gmail/Maps), and the IPython/Jupyter lineage.

Fernando Pérez started IPython in late 2001, influenced by Mathematica notebooks and Maple worksheets. A first notebook attempt came in 2005 (Wx toolkit, Google Summer of Code prototype under Robert Kern and Fernando Pérez). A second web-based prototype (2006, Min Ragan-Kelley advised by Brian Granger, SQL backend) was abandoned as too complex. A third prototype arrived October 2010, and in spring–summer 2011 Brian Granger worked full-time on a web notebook, building on the IPython kernel architecture and message spec (created with Pérez) and PyZMQ (with Min Ragan-Kelley). PyZMQ (ZeroMQ bindings) and websockets were key technologies. IPython Notebook 0.12 was released 21 December 2011.

In 2014, Project Jupyter was created out of IPython ("The Big Split"). The split was driven not only by project size but because from 2011–2014 the IPython Notebook began working with other languages (Julia first, then R), making the "IPython" name misleading. "Jupyter" is inspired by Julia, Python, and R. Language-agnostic parts (notebook format, message protocol, Qt console, notebook web app) moved to Jupyter. IPython retained two roles: the Python kernel backend for Jupyter and an interactive Python shell; its ecosystem also includes a parallel computing framework.

The article then classifies features by owner:
- **Kernels**: Jupyter (central architectural abstraction; IPython, IRkernel, IJulia; community kernels for Scala, JavaScript, Haskell, Ruby, etc.).
- **Notebook deployment**: Jupyter (docker-stacks, ipywidgets, jupyter-drive, jupyter-sphinx-theme, kernel_gateway, nbviewer, tmpnb, traitlets).
- **System shell use**: IPython (`!` escape, `!!`, `%sx`, `$`/`$$` variable expansion, aliases, `%rehashx`).
- **Magics**: kernel feature (IPython uses `%`/`%%`; cell magics take the whole block). IPython magics include `%matplotlib` (plotting), `%cd`/`%bookmark` (filesystem navigation), `%pdb`/`%debug` (debugger), and `%load_ext`/`%reload_ext`/`%unload_ext` (extensions like oct2py, rpy2, Cython, sympy, fortranmagic, sparkmagic). Other kernels differ: IRkernel has no magics (uses R's `system()`, `display_markdown()`); IJulia doesn't use magics but hints at Julia equivalents; IScala supports its own magics; metakernel-based kernels reuse many IPython magics.
- **Notebook conversion/formatting**: Jupyter (nbconvert, nbformat).
- **Saving/loading notebooks**: Jupyter Notebook application (.ipynb files, Save and Checkpoint).
- **Keyboard shortcuts and multi-cursor**: Jupyter Notebook.
- **Distributed parallel computing**: IPython ecosystem (ipyparallel, split off since v4.0).
- **Terminal**: Jupyter ecosystem (Jupyter Console; the old IPython console is deprecated).
- **Qt Console**: moved from IPython to Jupyter.

It concludes with a heuristic: is a feature Python-specific or general? (`%pdb` and `%matplotlib` are Python-specific; `%cd` is general.)

## Key points

- IPython (started 2001 by Fernando Pérez) is the ancestor; Project Jupyter was created from it in 2014 ("The Big Split").
- The split was driven by size and by non-Python language support (Julia, R); "Jupyter" = Julia, Python, R.
- IPython now serves as the Python kernel backend and an interactive Python shell; Jupyter owns language-agnostic components.
- Kernels, notebook deployment, conversion/formatting, saving/loading, shortcuts, and Qt console belong to Jupyter.
- System shell escapes, magics (in the IPython kernel), and distributed parallel computing (ipyparallel) belong to IPython.
- Magics are a kernel feature; different kernels (IRkernel, IJulia, IScala, metakernel) handle them differently.
- Heuristic: Python-specific features are IPython; general/language-agnostic features are Jupyter.

## Technical data / figures

| Feature | Owner |
| --- | --- |
| Kernels | Jupyter |
| Notebook deployment tools (nbviewer, tmpnb, kernel_gateway) | Jupyter |
| System shell escapes (`!`, `!!`, `%sx`) | IPython |
| Magics (`%`, `%%`, e.g. `%matplotlib`, `%pdb`) | IPython kernel (a kernel feature) |
| Notebook conversion/formatting (nbconvert, nbformat) | Jupyter |
| Saving/loading notebooks (.ipynb) | Jupyter Notebook app |
| Distributed parallel computing (ipyparallel) | IPython ecosystem |
| Terminal (Jupyter Console) | Jupyter |
| Qt Console | Jupyter (moved from IPython) |

- Timeline: MATLAB (1980s), Mathematica (1987), Maple (1989), Sage Notebook (mid-2000s), IPython (2001), IPython Notebook 0.12 (21 Dec 2011), Project Jupyter (2014).
- ipyparallel split from IPython at version 4.0.

## Why this source matters for the RAG

It provides authoritative historical and functional clarification of the IPython vs Jupyter distinction, a common point of confusion for Python/data-science learners. It also catalogues magics, extensions, and ecosystem tools useful for technical Q&A.
