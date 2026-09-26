---
id: collect-240926-huggingface/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face-7
title: "Load tokenizer and model"
domain: huggingface
role: reference
task: reference
actors: ["Nvidia", "SGLang", "vLLM"]
dates: []
keywords: ["agentic", "distribution", "guardrails", "llama", "llama.cpp", "nvidia", "reasoning", "sglang", "training", "vllm"]
source: docs/RAG/clean_en/huggingface/nvidia-nvidia-nemotron-3-nano-30b-a3b-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [521, 611]
sha256: 80cbd8fff5ba89bb3f5c3321b7646e9958e878c680526374a878300a54f5fb6d
---

# Load tokenizer and model

| Language | Tokens | 
|---|---|
| Assembly | 750,628,764 | 
| C | 42,657,300,868 | 
| C# | 56,153,329,307 | 
| C++ | 67,773,701,658 | 
| CommonLisp | 263,234,672 | 
| CSS | 38,848,760,035 | 
| Cuda | 400,222,993 | 
| Dart | 3,816,960,470 | 
| Dockerfile | 474,958,084 | 
| Fortran | 1,105,049,387 | 
| Go | 8,332,419,480 | 
| Haskell | 1,294,613,669 | 
| HTML | 69,082,117,487 | 
| Java | 131,440,465,822 | 
| JavaScript | 75,573,420,861 | 
| JSON | 15,366,881,241 | 
| Julia | 621,046,949 | 
| JupyterNotebook | 2,241,893,197 | 
| Lua | 4,146,420,802 | 
| Makefile | 12,640,010,879 | 
| Markdown | 64,796,743,311 | 
| Mathematica | 320,504,225 | 
| OmniversePython | 26,946,093 | 
| Pascal | 1,625,013,876 | 
| Perl | 1,575,314,434 | 
| PHP | 61,575,339,005 | 
| Python | 126,916,727,384 | 
| R | 19,811,381,935 | 
| reStructuredText | 1,779,876,391 | 
| Ruby | 6,446,962,615 | 
| Rust | 4,438,640,533 | 
| Scala | 3,343,959,154 | 
| Shell | 18,758,779,250 | 
| SQL | 23,205,633,085 | 
| Swift | 5,976,714,881 | 
| SystemVerilog | 233,056,185 | 
| TeX | 7,347,157,527 | 
| TypeScript | 15,657,838,582 | 
| Verilog | 811,884,369 | 
| VHDL | 648,401,444 | 
| VisualBasic.NET | 1,005,680,881 | 
| XML | 12,616,779,741 | 
| YAML | 10,574,010,491 | 

For our post-training recipe, we focused on 5 main languages in addition to English: Spanish, French, Japanese, Italian, German.

Those languages were represented in the form of multilingual reasoning and translation task.

The following table depicts our sample distribution for the 6 languages and 5 translation pairs.

| Language | Size | 
|---|---|
| English | 16.2 M | 
| Italian | 0.252M | 
| German | 0.252M | 
| Spanish | 0.252M | 
| French | 0.252M | 
| Japanese | 0.252M | 
| English <-> Italian | 108k | 
| English <-> German | 108k | 
| English <-> Spanish | 108k | 
| English <-> French | 108k | 
| English <-> Japanese | 108k | 

- Data Collection Method by dataset: Hybrid: Human, Synthetic
- Labeling Method by dataset: Hybrid: Automated, Human, Synthetic

- Engines: HF, vLLM, TRT-LLM, SGLang, Llama.cpp
- Test Hardware: NVIDIA B200 192GB, RTX PRO 6000 96GB, Jetson Thor, DGX Spark

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. When downloaded or used in accordance with our Trustworthy AI terms of service, developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse.

We advise against circumvention of any provided safety guardrails contained in the Model without a substantially similar guardrail appropriate for your use case.For more details: Safety and Explainability Subcards.

For more detailed information on ethical considerations for this model, please see the Model Card++ Bias, and Privacy Subcards.

Please report security vulnerabilities or NVIDIA AI Concerns here.

```
@misc{nvidia_nemotron_nano_v3_2025,
  title  = {{Nemotron 3 Nano}: Open, Efficient Mixture-of-Experts Hybrid {Mamba}-{Transformer} Model for {Agentic} Reasoning},
  author = {{NVIDIA}},
  year   = {2025},
  url    = {https://arxiv.org/abs/2512.20848},
  note   = {Technical report}
}
```
- Downloads last month
- 1,077,751
