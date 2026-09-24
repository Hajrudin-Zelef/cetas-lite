---
id: collect-240926-misc/misc/glm-5-3-flash-how-to-run-locally-unsloth-documentation
title: "glm-5-3-flash-how-to-run-locally-unsloth-documentation"
domain: unsloth
role: reference
task: reference
actors: ["Anthropic", "Apple", "DeepSeek", "Google", "Nvidia", "OpenAI", "Unsloth", "Z.ai"]
dates: []
keywords: ["glm", "agentic", "agents", "attention", "benchmark", "benchmarks", "claude", "context window", "deepseek", "gemini", "gpt-5.6", "gpu"]
source: docs/RAG/clean_en/misc/glm-5-3-flash-how-to-run-locally-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 448]
sha256: 6b56b0ed0067eb34518eab549034f48603238eb4f017406143c7f5a84593f48c
---

# glm-5-3-flash-how-to-run-locally-unsloth-documentation

<!-- source: https://unsloth.ai/docs/models/glm-5.3-flash -->

GLM-5.3-Flash, also known as **ox-alpha****outperforms** GLM-5.2. GLM-5.3-Flash is the smaller version of GLM-5.3 and rivals **Claude Opus 4.8** on coding and agentic benchmarks. You can now run the 1-bit model locally on 102GB RAM/VRAM or 3-bit on 128GB setups via llama.cpp or Unsloth. Thank you Z.ai for day-zero access.

Unsloth dynamic **1-bit** (93GB) GGUFs retains **71% of top-1 accuracy** whilst being **85% smaller** vs BF16 (642GB). Dynamic 3-bit is 76% smaller and retains 87% accuracy.

GLM-5.3-Flash was trained on 30T tokens and is built on a newly trained base model. Its hybrid sparse and linear attention architecture lowers long-context serving costs without sacrificing accuracy.

You can now run the model directly in Unsloth Desktop.

The smallest 1-bit quant works on 100GB RAM while 3-bit works on 128GB devices like a Mac or NVIDIA DGX Spark.
**Table: Hardware requirements** (units = total memory: RAM + VRAM, or unified memory)

100 GB

115 GB

128-150 GB

162-210 GB

350 GB

650 GB

GLM-5.3-Flash has **3 thinking modes**: Low, High, and Max. Use Max Thinking for complicated tasks. In Unsloth, you can easily select Low, High, or Max Thinking with a toggle in the chat area.

Use these settings for most use cases:

`temperature` = 1.0

`temperature` = 0.95

`top_p` = 0.95

`top_p` = 1.0

- **Maximum context window:**`1,048,576` .

GLM-5.3-Flash uses Max reasoning by default. It also supports reasoning efforts where `reasoning_effort` can be "low", "high", or "max".

As of Sep 4, we’ve added several improvements and optimizations to our day-zero llama.cpp PR. We implemented faster decoding pat plus bonus MTP support, enabling up to **3.3× faster inference** at long context lengths!

Everything works out of the box in Unsloth Desktop, simply update to the latest version if needed. No additional modules or MTP files are required. Alternatively, you can follow our llama.cpp guide.

Using GLM-5.3-Flash UD-IQ1_S on 1xB200 and disregarding MTP first, we get:

pp512

1121.80

1122.0

tg32

62.79

63.10

tg32 @ 4096

53.52

59.50

tg32 @ 16384

41.02

57.99

tg32 @ 65536

20.66

48.99

Then once we add MTP, we see even larger payoffs especially for longer contexts. However we should stop at around n=2 as more draft tokens makes inference slower.

4096

58.6

86.5

80.2

63.7

16K

55.0

77.2

For shorter context lengths, we still see speedups, albeit up to 1.6x faster.

We quantized GLM-5.3-Flash down to UD-IQ1_S 1bit (93.09GB) and it retains 71% of top-1% accuracy whilst being 85% smaller vs BF16 (641.64GB)

Dynamic 2-bit UD-Q2_K_XL is 109GB, is 83% smaller and retains 78% of accuracy. Dynamic 3-bit UD-IQ3_XXS is 120GB, is 81% smaller and retains 82% of accuracy. Dynamic 4-bit UD-Q4_K_XL is 200GB, is 69% smaller and retains 93% of accuracy.

UD-IQ1_S

93.09

70.89%

0.669714

9.1658

UD-IQ1_M

97.58

73.06%

0.572413

8.5069

UD-IQ2_XXS

101.84

76.30%

0.450148

7.5764

UD-Q2_K_XL

108.72

78.34%

0.380134

6.8412

UD-IQ3_XXS

120.37

81.63%

0.283772

5.9611

UD-Q3_K_XL

147.54

86.25%

0.159697

4.0281

UD-IQ4_XS

156.82

88.18%

0.116652

3.1014

UD-Q4_K_XL

199.71

92.22%

0.049294

1.4894

UD-Q5_K_XL

240.31

94.35%

0.027052

0.8696

UD-Q6_K_XL

291.83

95.23%

0.019007

0.6267

You can now run GLM-5.3-Flash (Ox-Alpha) in Unsloth Desktop and llama.cpp with our specific PR. We are using 3-bit `UD-IQ3_XXS` in our demos as it fits on 128GB devices. Feel free to change quantization type.

GLM-5.3-Flash can now run in Unsloth Desktop, an open-source UI app for local AI. **Unsloth automatically offloads to RAM and detects multiGPU setups**. With Unsloth Desktop, you can run models locally on **MacOS, Windows**, Linux and:

- Fast CPU + GPU inference via MLX and llama.cpp

Inference parameters should be auto-set when using Unsloth, however you can still change it manually. You can also edit the context length, chat template and other settings.

For more information, you can view our Unsloth inference guide. 1-bit running below:

You can use `unsloth run` command and serve GLM-5.3-Flash via an API using `llama-server` runtime flags, including context sizing, GPU layers, threading, sampling, networking, and tool configuration. For more info see our API docs or unsloth start.

We need to use our specific llama.cpp PR here. You can follow the build instructions below as well. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.

To run the model, you can do:

Then to run it:

Replace `UD-IQ3_XXS` with your preferred quant, such as `IQ2_XXS` for 2-bit once uploaded.

Benchmark

GLM-5.3-Flash

GLM-5.2

DeepSeek-V4-Vision-Exp

Opus 4.8

GPT-5.6 Terra

Gemini 3.7 Flash

Coding

Terminal Bench 2.1

84.3

81.0

83.9

85.0

87.4

85.8

DeepSWE

v1.1

63.4

46.2

59.3

58.0

69.6

65.3

NL2Repo

56.3

48.9

57.7

69.7

-

-

Agentic

Toolathlon Verified

78.4

59.9

75.9

76.2

74.9

-

AutomationBench

v1.0.6

48.8

26.2

38.8

41.0

37.2

52.3

Agents' Last Exam

26.3

20.4

27.3

27.0

28.0

-

HLE w/ Tools

55.3

54.7

55.1

57.9

-

-

GDPval-AA v2

1773

1504

1675

1582

1571

1527

Vision

OfficeQA Pro

62.4

-

57.9

48.9

-

-

CharXiv Reasoning

w/ Tools

89.4

-

80.4

89.9

88.0

88.7

Chartography

w/ Tools

78.0

-

64.3

75.0

68.0

65.0

BabyVision

53.4

-

35.1

46.8

61.6

70.9

MVbench

77.8

-

69.4

67.1

75.0

82.2

MMVU

80.5

-

72.7

67.4

75.8

82.3

Last updated

Was this helpful?
