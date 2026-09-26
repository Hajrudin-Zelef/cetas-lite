---
id: collect-240926-mindstudio/mindstudio/glm-5-3-flash-hands-on-multi-gpu-test-coding-and-refusals-2
title: "glm-5-3-flash-hands-on-multi-gpu-test-coding-and-refusals"
domain: mindstudio
role: reference
task: reference
actors: ["Unsloth", "Z.ai"]
dates: []
keywords: ["glm", "gpu", "llama", "llama.cpp"]
source: docs/RAG/clean_en/mindstudio/glm-5-3-flash-hands-on-multi-gpu-test-coding-and-refusals.md
source_anchor: ""
source_lines: [76, 98]
sha256: b8213fa0529933e78e718d0bf6458d4db78339dee55daab66789439db796cee5
---

# glm-5-3-flash-hands-on-multi-gpu-test-coding-and-refusals

### 本次测试使用了哪种GLM-5.3 Flash量化？

该测试使用了Unsloth的动态1位量化，文件约93 GB，选择它而非2位版本是为了保留更多上下文长度和推理token的空间。

### 本地运行GLM-5.3 Flash需要多少块GPU？

## Remy不构建管道。它继承管道。

Remy从MindStudio继承了所有这些——因此每个周期都投入到你真正想要的应用中。

本次测试使用llama.cpp将模型分散到五块GPU上，CUDA可见设备设置为包含全部五张卡，每张卡的VRAM使用量大约在17 GB到23 GB之间。

### GLM-5.3 Flash支持图像或视觉输入吗？

在llama.cpp中尚不支持。测试时，该模型的视觉塔实现尚未在技术栈中可用，因此会话通过聊天界面仅限文本。

### GLM-5.3 Flash在多GPU设置上有多快？

简单提示的生成速度约为每秒33到34个token。一个冗长、推理密集的编码任务中，随着上下文增长超过100,000个token，吞吐量衰减至约每秒14到15个token。

### GLM-5.3 Flash会拒绝有害或胁迫性提示吗？

会。当被要求充当胁迫非自愿人类船员的执法者时，它拒绝了并解释了原因，同时仍然提出通过合法手段帮助解决根本问题。一旦场景被改变以将人类目标从胁迫中移除，它就改变了答案。
