---
id: collect-240926-misc/misc/deepseek-v4-how-to-run-locally-unsloth-documentation-2
title: "deepseek-v4-how-to-run-locally-unsloth-documentation"
domain: unsloth
role: reference
task: reference
actors: ["DeepSeek", "Perplexity", "Unsloth"]
dates: []
keywords: ["deepseek", "agents", "benchmarks", "gguf", "gpu", "inference", "llama", "llama.cpp", "memory", "mxfp4", "parameters", "perplexity"]
source: docs/RAG/clean_en/misc/deepseek-v4-how-to-run-locally-unsloth-documentation.md
source_anchor: ""
source_lines: [106, 356]
sha256: d5b1e15ebcf63051ab29832d78240aa0bb860a56547b0de7153af7e5efcff4c8
---

# deepseek-v4-how-to-run-locally-unsloth-documentation

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
