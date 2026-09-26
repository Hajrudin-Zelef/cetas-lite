---
id: collect-240926-storagereview/storagereview/fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f-3
title: "fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f"
domain: storagereview
role: reference
task: reference
actors: ["Nvidia", "OpenAI", "TensorRT-LLM"]
dates: []
keywords: ["chatgpt", "cost", "gpu", "nvidia", "parameters", "tensorrt"]
source: docs/RAG/clean_en/storagereview/fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f.md
source_anchor: ""
source_lines: [42, 79]
sha256: 4c9a627dbf9ae13f885124c730675490b4823acb0275a75fdccd2023e7615e57
---

# fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f

This turnkey application leverages the power of retrieval-augmented generation (RAG), combined with the efficiency of the TensorRT-optimized LLM and the high-speed capabilities of RTX acceleration. These provide contextual responses that are both fast and highly relevant. Running directly on your RTX Windows desktop or workstation, this setup ensures quick access to information and a high degree of privacy and security, as all processing is handled locally.
Implementing an LLM with RAG capabilities offers an excellent solution for professionals and advanced users who prioritize privacy, security, and personalized efficiency. Unlike public models such as ChatGPT, which process queries over the Internet, a local LLM operates entirely within the confines of your QNAP NAS.
This offline functionality ensures that all interactions remain private and secure. This allows users to customize the AI's knowledge base according to their specific needs, whether it's confidential business documents, specialized databases, or personal notes. This approach greatly improves the relevance and speed of AI responses, making it a valuable tool for those who need immediate, contextual information without compromising privacy or data security.
Also worth noting, and this may seem obvious, adding a GPU to the NAS directly simplifies the link between a company's data and the LLM. There is no need to move data to take advantage of this particular model, and the process is as simple and cost-effective as installing a mid-range GPU in the NAS. Moreover, at present, all this software is free, which greatly democratizes the potential of AI for small organizations.
Chat with RTX is still a beta program and at the time of writing, we were using version 0.2. But the ease of installing it and getting the web interface running was refreshing. Anyone who knows how to download and install an application can now get a local LLM with RAG in a few clicks.
Enabling remote access to chat with RTX via a universally accessible URL
We took our scenario to the next level and made it available to the entire office.
Step 1: Locate the configuration file
Start by navigating to the folder containing the configuration file:
- File path: C:\Users\{YourUserDir}\AppData\Local\NVIDIA\ChatWithRTX\RAG\trt-llm-rag-windows-main\ui\user_interface.py
Step 2: Update the launch code
Open the user_interface.py file and Ctrl-F for interface.launch. Locate the correct segment, which will appear by default as follows:
interface.launch(
    favicon_path=os.path.join(os.path.dirname(__file__), 'assets/nvidia_logo.png'),
    show_api=False,
    server_port=port
)
To enable network access, you need to add share=True as well:
interface.launch(
    favicon_path=os.path.join(os.path.dirname(__file__), 'assets/nvidia_logo.png'),
    show_api=False,
    share=True,
    server_port=port
)
Save the changes to the user_interface.py file. Then launch Chat with RTX via the Start menu, which will open a command prompt window and activate the interface.
Step 3: Find the public URL
The command prompt window will display both a local and public URL. To create a working public URL accessible from any device, merge the elements of both URLs. It would be best to take the public URL and add the local cookie information to the end:
- Public URL: https://62e1db9de99021560f.gradio.live
- Local URL with parameters: http://127.0.0.1:16852?cookie=4a56dd55-72a1-49c1-a6de-453fc5dba8f3&__theme=dark
Your combined URL should look like this, with the ?cookie added to the public URL:
https://62e1db9de99021560f.gradio.live?cookie=4a56dd55-72a1-49c1-a6de-453fc5dba8f3&__theme=dark
This URL allows access to Chat with RTX from any device on your network, extending its usability beyond local constraints.
Final thoughts
We have long been fans of QNAP's leadership in NAS hardware design, but QNAP customers have far more value at their disposal than they probably realize. To be honest, Virtualization Station is a great starting point, but why not take it to the next level and try GPU Passthrough? At the very least, organizations can provide a high-end GPU-powered virtual machine without having to set up a dedicated workstation. There are also the apparent benefits of a VM placed next to a huge internal storage pool with native performance levels. In this case, we had shared storage performance of nearly 10 GB/s, without worrying about a single 100 GbE connection or switch, all because the GPU-accelerated VM was inside the NAS itself.
Why not go even further to realize the benefits of AI for the organization? We have shown that adding a decent GPU to a QNAP NAS is relatively simple and inexpensive. We put an A4000 to work, and with a list price of around $1,050, that's not bad considering Virtualization Station is free and NVIDIA Chat with RTX is available for free. Being able to securely point this powerful LLM at a company's private data should provide actionable insights while making the business more agile.
Another goal to consider here is a file store for models that can be external to the QNAP system itself. This is ideal for small businesses that need a fast place to store their working data. With advanced networking capabilities, you could eventually use the NAS as a data storage location for RAG work on a larger GPU server, enabling an easily shareable data store from which to infer.
This is just one example of AI. The industry is evolving rapidly, so tools will continue to be made available. Smart businesses need to learn how to leverage AI, and this simple QNAP feature is a great way to start.
QNAP Virtualization Station
