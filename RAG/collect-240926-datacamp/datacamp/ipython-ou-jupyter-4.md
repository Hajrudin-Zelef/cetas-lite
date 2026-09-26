---
id: collect-240926-datacamp/datacamp/ipython-ou-jupyter-4
title: "Assign the result to `ls`"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "parameters", "training"]
source: docs/RAG/clean_en/datacamp/ipython-ou-jupyter.md
source_anchor: ""
source_lines: [291, 331]
sha256: b0c53026bee608fe7cb9e078e6eecb307f7fd0229a4ada39cecba7d6fe122866
---

# Assign the result to `ls`

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
