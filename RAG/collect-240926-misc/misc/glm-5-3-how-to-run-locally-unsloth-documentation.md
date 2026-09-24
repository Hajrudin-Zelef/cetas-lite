---
id: collect-240926-misc/misc/glm-5-3-how-to-run-locally-unsloth-documentation
title: "glm-5-3-how-to-run-locally-unsloth-documentation"
domain: unsloth
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "DeepSeek", "ExploitGym", "Hugging Face", "Moonshot", "Nvidia", "OpenAI", "Unsloth", "Z.ai"]
dates: []
keywords: ["glm", "agentic", "agents", "benchmark", "context window", "cyber", "deepseek", "fable 5", "gguf", "gpt-5.6", "gpu", "inference"]
source: docs/RAG/clean_en/misc/glm-5-3-how-to-run-locally-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 537]
sha256: e86de6f05f56b3ccbfd95c4e1993f09098114ab8533c061f22c765e6fa36e586
---

# glm-5-3-how-to-run-locally-unsloth-documentation

<!-- source: https://unsloth.ai/docs/models/glm-5.3 -->

GLM-5.3 is Z.ai’s new 744B parameter (40B active) model. As of Aug 2026, GLM-5.3 is the **strongest open-model** to date, achieving SOTA on Terminal Bench 3.0 and Agents' Last Exam. GLM-5.3 uses the same base model as GLM-5.2, with every gain coming from post-training. The model has a **1M context** window and can now run locally via Unsloth Dynamic GGUFs with llama.cpp or Unsloth Desktop.

Dynamic 1-bit GGUF reaches **~76%** top-1 accuracy while being **85% smaller**. Dynamic 2-bit reaches **~81%** accuracy while being **83% smaller**. As GLM-5.3 has same size and arch as GLM-5.2, most requirements / settings are the same. Thanks Z.ai for Unsloth day-zero access. **GLM-5.3-GGUF**

The 2-bit dynamic quant `UD-IQ2_M` uses **239GB** of disk space works well on **256GB RAM** devices like a 2x NVIDIA DGX Sparks or a Mac Studio.

The **1-bit** quant will fit on 223GB RAM and 8-bit requires 810GB RAM.

**Table: Inference hardware requirements** (units = total memory: RAM + VRAM, or unified memory)

223GB

245GB

290-360GB

372-475GB

570GB

810GB

For best performance, make sure your total available memory, including VRAM and system RAM, exceeds the quantized model file size by a comfortable margin.

GLM-5.3 has **3 thinking modes**: **Low**, **High**, and **Max**. Use Max Thinking for complicated coding tasks. In Unsloth Desktop you can easily toggle Low, High, and Max Thinking with a UI.

Use these settings for most use cases:

`temperature` = 1.0

`temperature` = 1.0

`top_p` = 0.95

`top_p` = 1.0

The **maximum context window** is `1,048,576`.

GLM-5.3 uses max reasoning by default and thinking cannot be disabled. `reasoning_effort` can be `low`, `high`, or `max`.

`clear_thinking` is false by default and they recommend true.

For reasoning effort customization (change 'low' to 'high' or 'max'):

For multi-turn chat, use `clear_thinking=true` to remove reasoning from previous turns (recommended for this model):

We found GLM uses an interesting `.{id}.` notation for chat templates, but many engines don't support this. We edited it to all `[id]` so Python list indexing syntax. See this commit change for more details.

You can now run GLM-5.3 in llama.cpp and Unsloth Desktop. We will be utilizing the 239GB `UD-IQ2_M` quant for the best balance of accessibility and accuracy.

GLM-5.3 can now run in Unsloth Desktop, an open-source UI app for local AI. **Unsloth automatically offloads to RAM and detects multiGPU setups**. With Unsloth Desktop, you can run models locally on **MacOS, Windows**, Linux and:

- Fast CPU + GPU inference via MLX and llama.cpp

Inference parameters should be auto-set when using Unsloth, however you can still change it manually. You can also edit the context length, chat template and other settings.

For more information, you can view our Unsloth inference guide.

You can use `unsloth run` command and serve GLM-5.3 via an API using `llama-server` runtime flags, including context sizing, GPU layers, threading, sampling, networking, and tool configuration. For more info see our API docs or unsloth start.

For this guide we'll be running the `UD-IQ2_M` quant which will require at least 245GB RAM. Feel free to change quantization type. For these tutorials, we will using llama.cpp for fast local inference. GGUF: **GLM-5.3-GGUF** 

Obtain the latest `llama.cpp` **on** **GitHub here**. You can follow the build instructions below as well. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.

You can now use `llama.cpp` directly to load and download models, just like `ollama run`. First, select the quantization type you want like `UD-IQ2_M`. Also use `export LLAMA_CACHE="unsloth/GLM-5.3-GGUF"` to force `llama.cpp` to save to a specific location. **Note this download process might be very slow**, so it's probably best to use the manual download process in the next section.

If you want to download the model manually **(much faster!)**, we can download the model via the code below (after installing `pip install huggingface_hub`). If downloads get stuck, see: Hugging Face Hub, XET debugging

If you want to use the dynamic 1-bit, then do:

Then run the model in conversation mode. Use `unsloth/GLM-5.3-GGUF/UD-IQ2_M/GLM-5.3-UD-IQ2_M-00001-of-00006.gguf` for 2bit or `unsloth/GLM-5.3-GGUF/UD-IQ1_S/GLM-5.3-UD-IQ1_S-00001-of-00006.gguf`  for 1bit.

Similar to GLM-5.2, when you launch llama-cli, you will see:

Then after prompt, we made a cool small Snake game using `UD-IQ1_S` so 1-bit and it worked well with GLM-5.3!

To utilize long context in llama.cpp, use KV cache quantization to reduce memory usage.

Currently, these KV cache dtypes are supported: `f32`, `f16`, `bf16`, `q8_0`, `q4_0`, `q4_1`, `iq4_nl`, `q5_0`, and `q5_1`. By default `f16` is used. `q4_1` uses around 5 bits per weight, allowing around **3.2x longer context lengths**.

We ran KLD for the quants we uploaded as well and it shows Q4_K_XL and Q5_K_XL are very close to the baseline, so aim for those.

UD-IQ1_S

216.7

72.56

0.687991

9.104

4.6130

UD-IQ1_M

228.5

75.64

0.565455

8.595

4.1410

UD-IQ2_M

238.6

78.53

0.453992

7.717

3.7433

UD-Q2_K_XL

253.9

80.93

0.374219

7.076

3.5048

UD-IQ3_XXS

281.7

84.15

0.272796

6.252

3.2482

UD-Q3_K_XL

343.0

88.86

0.141406

4.127

2.9107

UD-IQ4_XS

365.3

90.59

0.101496

3.177

2.8460

UD-Q4_K_XL

467.3

94.29

0.036922

1.309

2.7006

UD-Q5_K_XL

562.5

95.82

0.019728

0.786

2.6842

UD-Q6_K_XL

684.4

96.59

0.013257

0.534

2.6771

You can view GLM-5.3's key benchmark improvements in table format further below:

Benchmark

GLM-5.3

GLM-5.2

Kimi K3

DeepSeek-V4

Pro-0813

Qwen3.8-Max

Opus 4.8

Fable 5

(w/ fallback)

GPT-5.6 Sol

Coding

Terminal Bench 2.1

88.2

81.0

88.3

87.9

86.6

85.0

88.0

88.8

Terminal Bench 3.0

28.3

4.6

17.4

-

-

21.1

33.7

34.6

DeepSWE

v1.1

66.9

46.2

67.5

62.7

56.6

58.0

69.7

72.7

NL2Repo

58.0

48.9

58.0

61.1

55.9

69.7

-

-

ProgramBench

Almost Solved

19.0

9.5

17.5

-

10.5

15.5

33.0

23.0

FrontierSWE

78.1

67.5

-

-

-

66.5

88.2

-

SWE-Marathon

v1.1

42.5

19.4

48.1

-

-

48.8

33.1

42.5

PostTrainBench

39.8

31.7

32.0

-

-

32.9

41.8

36.2

Cyber

CyberGym

84.5

77.2

80.0

83.3

78.5

78.1

83.8

83.6

ExploitGym

2h / 6h

105 / 130

29 / 39

36 / 70

-

14 / 26

80 / 120

181 / 247

216 / 293

ExploitBench

54.4

24.4

32.2

-

28.8

40.0

78.0

76.5

Agentic

Toolathlon Verified

73.0

59.9

76.5

74.1

72.5

76.2

74.7

74.9

AutomationBench

v1.0.6

48.2

26.2

46.7

43.2

39.8

41.0

46.2

45.8

Agents' Last Exam

ALE-CLI

28.5

23.8

27.6

25.7

27.0

25.7

23.8

28.6

HLE w/ Tools

62.5

54.7

59.8

60.0

56.2

57.9

63.9

64.5

GDPval-AA v2

1769

1508

1682

1590

1739

1588

1743

1730

Last updated

Was this helpful?
