---
id: collect-250926-servers-hardware/servers-hardware/saving-to-gguf-unsloth-documentation
title: "https://github.com/ggml-org/llama.cpp/blob/master/examples/quantize/quantize.cpp#L19"
domain: servers-hardware
role: reference
task: reference
actors: ["Hugging Face", "Unsloth", "vLLM"]
dates: []
keywords: ["llama", "llama.cpp", "attention", "gguf", "gpu", "inference", "inference engine", "memory", "quantization", "training", "vllm"]
source: docs/RAG/clean4/saving-to-gguf-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 76]
sha256: 0af153bd21e2ef499660708d69ab806c1d25fbdbd87a408176860a0f084f0bfc
---

# https://github.com/ggml-org/llama.cpp/blob/master/examples/quantize/quantize.cpp#L19

Saving models to 16bit for GGUF so you can use it for Unsloth Studio, Ollama, llama.cpp and more!

To save to GGUF, use the below to save locally:

```
model.save_pretrained_gguf("directory", tokenizer, quantization_method = "q4_k_m")
model.save_pretrained_gguf("directory", tokenizer, quantization_method = "q8_0")
model.save_pretrained_gguf("directory", tokenizer, quantization_method = "f16")
```
To push to Hugging Face hub:

```
model.push_to_hub_gguf("hf_username/directory", tokenizer, quantization_method = "q4_k_m")
model.push_to_hub_gguf("hf_username/directory", tokenizer, quantization_method = "q8_0")
```
All supported quantization options for `quantization_method` are listed below:

```
# https://github.com/ggml-org/llama.cpp/blob/master/examples/quantize/quantize.cpp#L19
ALLOWED_QUANTS = \
{
    "not_quantized"  : "Recommended. Fast conversion. Slow inference, big files.",
    "fast_quantized" : "Recommended. Fast conversion. OK inference, OK file size.",
    "quantized"      : "Recommended. Slow conversion. Fast inference, small files.",
    "f32"     : "Not recommended. Retains 100% accuracy, but super slow and memory hungry.",
    "f16"     : "Fastest conversion + retains 100% accuracy. Slow and memory hungry.",
    "q8_0"    : "Fast conversion. High resource use, but generally acceptable.",
    "q4_k_m"  : "Recommended. Uses Q6_K for half of the attention.wv and feed_forward.w2 tensors, else Q4_K",
    "q5_k_m"  : "Recommended. Uses Q6_K for half of the attention.wv and feed_forward.w2 tensors, else Q5_K",
    "q2_k"    : "Uses Q4_K for the attention.wv and feed_forward.w2 tensors, Q2_K for the other tensors.",
    "q3_k_l"  : "Uses Q5_K for the attention.wv, attention.wo, and feed_forward.w2 tensors, else Q3_K",
    "q3_k_m"  : "Uses Q4_K for the attention.wv, attention.wo, and feed_forward.w2 tensors, else Q3_K",
    "q3_k_s"  : "Uses Q3_K for all tensors",
    "q4_0"    : "Original quant method, 4-bit.",
    "q4_1"    : "Higher accuracy than q4_0 but not as high as q5_0. However has quicker inference than q5 models.",
    "q4_k_s"  : "Uses Q4_K for all tensors",
    "q4_k"    : "alias for q4_k_m",
    "q5_k"    : "alias for q5_k_m",
    "q5_0"    : "Higher accuracy, higher resource usage and slower inference.",
    "q5_1"    : "Even higher accuracy, resource usage and slower inference.",
    "q5_k_s"  : "Uses Q5_K for all tensors",
    "q6_k"    : "Uses Q8_K for all tensors",
    "iq2_xxs" : "2.06 bpw quantization",
    "iq2_xs"  : "2.31 bpw quantization",
    "iq3_xxs" : "3.06 bpw quantization",
    "q3_k_xs" : "3-bit extra small quantization",
}
```
First save your model to 16bit:

Then use the terminal and do:

Or follow the steps at https://rentry.org/llama-cpp-conversions#merging-loras-into-a-model using the model name "merged_model" to merge to GGUF.

### Running in Unsloth works well, but after exporting & running on other platforms, the results are poor

You might sometimes encounter an issue where your model runs and produces good results on Unsloth, but when you use it on another platform like Ollama or vLLM, the results are poor or you might get gibberish, endless/infinite generations *or* repeated outputs**.**

- The most common cause of this error is using an **incorrect chat template****.** It’s essential to use the SAME chat template that was used when training the model in Unsloth and later when you run it in another framework, such as llama.cpp or Ollama. When inferencing from a saved model, it's crucial to apply the correct template.
- You must use the correct `eos token` . If not, you might get gibberish on longer generations.
- It might also be because your inference engine adds an unnecessary "start of sequence" token (or the lack of thereof on the contrary) so ensure you check both hypotheses!
- **Use our conversational notebooks to force the chat template - this will fix most issues.**

You can try reducing the maximum GPU usage during saving by changing `maximum_memory_usage`.

The default is `model.save_pretrained(..., maximum_memory_usage = 0.75)`. Reduce it to say 0.5 to use 50% of GPU peak memory or lower. This can reduce OOM crashes during saving.

First save your model to 16bit via:

Compile llama.cpp from source like below:

Then, save the model to F16:

Last updated

Was this helpful?
