---
id: collect-240926-misc/misc/deepseek-v4-how-to-run-locally-unsloth-documentation
title: "deepseek-v4-how-to-run-locally-unsloth-documentation"
domain: unsloth
role: reference
task: reference
actors: ["Anthropic", "Apple", "DeepSeek", "Perplexity", "Unsloth"]
dates: []
keywords: ["deepseek", "agentic", "agents", "benchmark", "benchmarks", "claude", "context window", "fp8", "gguf", "gpu", "inference", "kv cache"]
source: docs/RAG/clean_en/misc/deepseek-v4-how-to-run-locally-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 356]
sha256: 72ffb93d4986ddeb167281a9b432f09fe8ffa97006540cec0622eb60e465a59a
---

# deepseek-v4-how-to-run-locally-unsloth-documentation

<!-- source: https://unsloth.ai/docs/models/deepseek-v4 -->

DeepSeek-V4, DeepSeek-V4-Flash-Vision-Exp, V4-**Pro-0813**, and V4-**Flash-0731** are new open-weight models - the Flash variant has 284B parameters (13B active), while V4-Pro has 1.6T (49B active). **V4-Pro-0813**, released on **Aug 13**, matches Claude-4.8-Opus performance, while **V4-Flash-0731**, released on **July 31**, delivers the best performance in its size class and **outperforms V4-Pro** (Preview). Built for coding, agentic, and chat workflows with a **1M context window**, this guide shows how to run DeepSeek-V4-Flash-0731 locally using Unsloth Dynamic GGUFs and Unsloth Desktop.

For **lossless** DeepSeek, use Q8 (`UD-Q8_K_XL`), which is only **7GB larger** than Q4 (`UD-Q4_K_XL`). The lossless 8-bit GGUF is **162 GB** and 3-bit is **103GB** which can run on a **110GB RAM** device**.** DeepSeek-V4-Flash-0731 scores 82.7% on Terminal Bench 2.1, 54.4% on DeepSWE, and 54.2% on NL2Repo. DSpark is also enabled for GGUFs, enabling up to **2x faster decoding speed**!

**Aug 6: DSpark is enabled for DeepSeek-V4-Flash-0731 - enabling 1.5x to 1.9x faster inference! DSpark is automatically enabled in** **Unsloth****.**

We also improved the DeepSeek-V4 chat jinja template, and tested over 4000 conversations to be equivalent with the official baseline.

Our `UD-Q8_K_XL` quant is fully lossless. DeepSeek-V4-Flash is quantization-aware-trained: the official checkpoint stores its routed experts (96% of the model) natively in MXFP4 and everything else in FP8 or BF16. GGUF's MXFP4 is exactly that format, so we repack the experts bit-for-bit, and FP8 dequantizes into BF16 with no rounding. We checked every tensor against the official DeepSeek weights: all 1,328 are bit-identical, and it stays lossless at inference (KL-divergence ~0, 100% top-token agreement).

**Non**-Unsloth DeepSeek-V4-Flash GGUFs were converted without these paths thus deviating from the official weights. `UD-Q4_K_XL` keeps the same bit-exact experts and only quantizes the non-expert tensors (4% of the model) to Q8_0, so it sits right next to Q8 in size and quality.

Measured against the official weights, both Unsloth quants are on the quality/size frontier. UD-Q8_K_XL is the only lossless point. UD-Q4_K_XL matches other community MXFP4 formats and is more accurate than the Q4_K-experts conversions, which are larger yet land at 0.029 KLD.

The error split by layer shows why. Keeping the native MXFP4 experts means 0% weight error at every layer. Conversions that re-quantize the experts to Q4_K or IQ2_XXS round almost every weight: 5% for Q4_K, over 30% for IQ2_XXS.

We also found using Q8_0 and F16 for some tensors is not lossless, and it gets worse since QAT was applied by DeepSeek to make MXFP4 / FP8 work well, so we had to leave them in BF16 directly. So use UD-Q8_K_XL for a true lossless quant, and UD-Q4_K_XL downcasts some of the BF16 items to Q8_0.

For full benchmark tables of GGUF Benchmarks, see here.

We also improved the DeepSeek-V4 chat jinja template, and tested over 4000 conversations to be equivalent with the golden baseline (official DS4)

We added `reasoning_effort` and you can select `max, high` just like official DeepSeek-V4. We prepend the correct system prompt as per DS4, and followed gpt-oss's style.

And for tool calls, `reasoning_content` was retained for DS4, but the jinja chat template would exclude them. We added it back.

DeepSeek-V4 uses reasoning by default. It also supports reasoning efforts where `reasoning_effort` can be "high", "max" or disabled.

To disable thinking, use `--chat-template-kwargs '{"enable_thinking":false}'`. If you're on **Windows** Powershell, use: `--chat-template-kwargs "{\"enable_thinking\":false}"`

You can also use `--reasoning on` or `--reasoning off` in llama.cpp as well now!

For reasoning effort customization and or to disable reasoning, use the below examples:

DeepSeek-V4-Flash is smaller and faster than DeepSeek-V4-Pro, with **284B** parameters (13B active), and a **1M context window**. The model has 3 modes, **Non-think**, **Think** **High** and **Think** **Max**. 

It's recommended to use `UD-IQ3_XXS` which is **103GB** for best results. Because the file size does not include KV cache, context allocation, try to have at least **110GB RAM** to run the model.

The `UD-Q8_K_XL` quant is DeepSeek-V4-Flash in full original precision. It is 162GB size and it's best to have at least 169GB of available RAM/VRAM available.

**Table: Inference hardware requirements** (units = total memory: RAM + VRAM, or unified memory)

Standard

92 GB

102 GB

110-135 GB

162 GB

169 GB

DSpark

102 GB

112 GB

120-145 GB

172 GB

179 GB

For best performance, make sure your total available memory, including VRAM and system RAM, exceeds the quantized model file size by a comfortable margin.

DeepSeek recommends these parameters for best performance: `temperature = 1.0`, `top-p = 1.0`. For **DeepSeek-V4-Flash-0731** and agentic scenarios, `top-p = 0.95` is suggested instead and `top-p = 1.0` for other tasks.

**Think High is on by default.** If disabled, you can enable it via: `--chat-template-kwargs '{"enable_thinking":true}'` or toggle it via the UI dropdown in Unsloth. Also see  DeepSeek V4 Chat template improvements

`temperature = 1.0`

`top-p = 1.0`

`top-p = 0.95` (agentic only)

`temperature = 1.0`

`top-p = 1.0`

- **Maximum context window:** `1,048,576`
- For Think Max, set context to at least **384K tokens** .

For this tutorial, we will use the 3-bit quant `UD-IQ3_XXS`, as it fits on a 128GB RAM device. Replace `UD-IQ3_XXS` with `UD-Q8_K_XL` (original quality) or another quant if your machine has enough memory. You can now run DeepSeek-V4-Flash-0731 in Unsloth Desktop . **DSpark is automatically enabled in** **Unsloth****.**

DeepSeek-V4-Flash-0731 can now be run and trained in Unsloth, our new open-source UI for local AI. Unsloth Desktop lets you run models locally on **MacOS**, **Windows**, Linux and:

- Fast CPU + GPU inference via llama.cpp

Inference parameters should be auto-set when using Unsloth, however you can still change it manually. Because **Think High is on by default**, you can go to the right dropdown to toggle it to Non-think or Think Max. You can also edit the context length, chat template and other settings. **DSpark is automatically enabled in** **Unsloth****.**

For more information, you can view our Unsloth inference guide.

You can use `unsloth run` command and serve DeepSeek via an API using `llama-server` runtime flags, including context sizing, GPU layers, threading, sampling, networking, and tool configuration. For more info see our API docs or unsloth start.

Obtain the latest `llama.cpp` **on** **GitHub here**. You can follow the build instructions below as well. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.

You can now use `llama.cpp` directly to load and download models, just like `ollama run`. First, select the quantization type you want like `IQ3_XXS`. Also use `export LLAMA_CACHE="folder"` to force `llama.cpp` to save to a specific location. Note this download process might be very slow, so it's probably best to use the manual download process in the next section.

You can edit `--threads 32` for the number of CPU threads, `--ctx-size 32768` for context length, `--n-gpu-layers 2` for GPU offloading on how many layers. Try adjusting it if your GPU goes out of memory. Also remove it if you have CPU only inference.

DeepSeek-V4-Flash-0731 has native DSpark, which allows for up to **2x faster decoding speed**! DSpark is a new algorithm by DeepSeek that is superior to naive MTP, and was introduced in this paper. DSpark enables DeepSeek-V4-Flash to reach **120 tokens/s** on a B200 GPU compared to the original 60 tokens/s baseline. **DSpark is automatically enabled in the** **Unsloth** **local UI.**

Llama.cpp integrated DSpark as part of PR 25784 with further improvements to multi GPU and more. We show using `--spec-draft-n-max 3` as a good default, allowing 1.9x faster inference speed. Larger values seem to be slower.

Download both the drafter and the GGUF - we made 2 Q8_0 and the lossless BF16 one. Note DSpark will need ~10GB more memory usage, so 128GB machines will need IQ3_XXS and Q8_0

Then load it via llama-cli or llama-server:

Benchmarks from the original DSpark paper as well showing how it fares against MTP:

See below for a table comparing benchmarks for quants from Unsloth and other providers. Reference = official weights. Perplexity and KL-divergence over wikitext-2 at ctx 512 on 4x B200.

Official (reference)

156.4

4.5319

0

0%

100%

100%

**Unsloth UD-Q8_K_XL**

161.9

4.5319

**~0 (lossless)**

0.000%

100.000%

**100.000%**

**Unsloth UD-Q4_K_XL**

155.1

4.5335

0.0102

3.40%

96.28%

97.46%

bartowski MXFP4

156.0

4.5351

0.0105

3.42%

96.18%

97.57%

antirez Q4KExperts-F16 (imatrix)

164.6

4.5743

0.0291

5.87%

93.95%

0.51%

antirez Q4KExperts-F16

164.6

4.5726

0.0290

5.89%

93.94%

0.93%

antirez mixed L37-42-Q4K (imatrix)

97.6

5.8169

0.3605

21.15%

79.74%

0.41%

antirez IQ2XXS (imatrix)

86.7

6.0808

0.4079

22.23%

78.15%

0.39%

antirez IQ2XXS

86.7

6.1518

0.4207

22.74%

77.92%

0.47%

DeepSeek-V4-Flash-0731 outperforms DeepSeek-V4-Pro (Preview) on the benchmarks below despite using far fewer activated parameters, and remains competitive with leading proprietary models.

Terminal Bench 2.1

82.7

61.8

72.1

81.0

85.0

NL2Repo

54.2

39.4

38.5

48.9

69.7

Cybergym

76.7

38.7

52.7

-

83.1

DeepSWE

54.4

7.3

12.8

46.2

58.0

Toolathlon-Verified

70.3

49.7

55.9

59.9

76.2

Agents' Last Exam

25.2

15.8

16.5

23.8

25.7

AutomationBench Public

25.1

10.8

12.8

12.9

27.2

DSBench-FullStack †

68.7

37.0

41.8

61.8

71.6

DSBench-Hard †

59.6

25.8

31.1

54.5

71.7

Last updated

Was this helpful?
