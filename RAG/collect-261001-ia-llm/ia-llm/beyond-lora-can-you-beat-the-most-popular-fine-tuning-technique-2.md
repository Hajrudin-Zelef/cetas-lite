---
id: collect-261001-ia-llm/ia-llm/beyond-lora-can-you-beat-the-most-popular-fine-tuning-technique-2
title: "beyond-lora-can-you-beat-the-most-popular-fine-tuning-technique"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "vLLM"]
dates: []
keywords: ["fine-tuning", "lora", "benchmark", "benchmarks", "llama", "llama.cpp", "memory", "parameters", "quantization", "research", "vllm"]
source: docs/RAG/collect-261001-ia-llm/beyond-lora-can-you-beat-the-most-popular-fine-tuning-technique.md
source_anchor: ""
source_lines: [51, 111]
sha256: 3925cb0b07ef19c87d335f2f72751d724338311a90f27813a280b90d43f80435
---

# beyond-lora-can-you-beat-the-most-popular-fine-tuning-technique

Let's take a closer look at the results for the LLM Math dataset benchmark. When it comes to test accuracy vs memory, we find that LoRA is indeed on the Pareto frontier. It achieves 53.2% test accuracy and requires 22.6 GB of VRAM at the peak. There are, however, other PEFT techniques on the Pareto Frontier. For instance, BEFT achieves 32.9% test accuracy and requires only 20.2 GB of memory at max. On the other end, we have Lily, which achieves 54.9% test accuracy but requires 25.6 GB of memory. Depending on what's more important to you, you may conclude that LoRA does not present the best tradeoff for you.

| *Test accuracy vs memory usage tradeoff of fine-tuning `meta-llama/Llama-3.2-3B` and evaluating it on GSM8K. LoRA does well but so do other PEFT techniques.* | 

It is also worth noting that even though LoRA does well on this task, we're not talking about vanilla LoRA. On one side, we have LoRA with rank stabilized initialization, which is a technique to scale the LoRA contribution differently from the default initialization and provides very good test accuracy (53.2%). On the other end, we have LoRA-FA, which uses an optimizer specialized for LoRA that freezes part of the LoRA weights and is thus more memory efficient (20.2 GB). Normal LoRA only achieves an accuracy of 48.1% at 22.5 GB memory and should thus be avoided in favor of the alternatives.

Next let's take a look at the image generation benchmark. In the Hugging Face Space, choose “image-gen” in the “Select Task” dropdown to show the results. The goal of the task is to learn a new concept, namely a cat plushy, and generalize it to new prompts.

| *Cat plushy image created with LoRA fine-tuned on `FLUX.2-klein-base-4B`.* | 

For this task, the main metric is “dino similarity”, which measures how much a generated image resembles the picture from a holdout test dataset, with higher values being better. As always, we also want to keep an eye on memory usage. When plotting the Pareto Frontier of these two metrics, we find that LoRA is below that frontier. Let's get concrete numbers: LoRA achieves a similarity score of 0.697 whereas OFT achieves 0.708; in terms of memory, LoRA requires 9.97 GB, and OFT requires 9.01 GB. Therefore, OFT strictly dominates LoRA on these metrics.

| *Test accuracy vs memory usage tradeoff of fine-tuning `FLUX.2-klein-base-4B` and evaluating it on the test set. Other PEFT techniques like OFT beat LoRA in terms of test score and lower memory usage.* | 

Of course, you should also check the other PEFT methods that are close to the Pareto frontier, as metrics can be subject to small variations due to randomness. Also, you should explore other metrics: is runtime performance important to you or do you care about the size of the checkpoints? Choose the relevant metric from the dropdown and the picture can change considerably. For the image generation benchmark, do inspect the generated sample images to get a vibe of the fine-tuned model's capability.

Objection: But the benchmarks favor one method over another!


One criticism that could be leveled at the `PEFT` benchmarks is that the choice of hyper-parameters may favor one technique over another. This is true, doing an exhaustive and fair hyper-parameter sweep with this many techniques is difficult. It is, however, very easy for everyone to contribute their own experiments to `PEFT`: If you believe that a specific PEFT technique can be improved by choosing different hyper-parameters, create a PR! We added instructions on how to do that. In a similar vein, if you want to contribute a completely new benchmark, reach out to us to discuss your idea.

Another problem with the benchmarks is that they may not fully reflect the capabilities of a specific PEFT technique. We make it possible to compare the techniques along many different dimensions and discover the best ones according to these tradeoffs. But it's impossible to capture all facets this way. For instance, one PEFT technique called Cartridges was developed to compress long prompts, which is not measured in the benchmarks. Other factors can also influence the choice, for instance:

- Depending on the PEFT technique, only certain layer types can be modified.
- Not all PEFT techniques support quantized base models (but we actively expand the support in `PEFT` ).
- Some PEFT techniques allow merging of the adapter to reduce runtime overhead but others don't.

The benchmarks cannot fully lift the responsibility to do your research, but they can be reasonable pointers.

| *Click on the image to peruse the PEFT shop to find the best PEFT technique for you. It allows you to browse not only by benchmark metrics but also by capabilities, like quantization support.* | 

Objection: But llama.cpp/vLLM/... only supports LoRA


A limitation of using a PEFT technique other than LoRA is that they don't get the broad support in downstream packages that LoRA sees. For example, if you want to serve the model using vLLM, only LoRA checkpoints can be loaded. Thankfully, `PEFT` now supports converting other adapters into LoRA. That way, you can convert a non-LoRA checkpoint into LoRA and use it in vLLM or other downstream packages.

To test this, we converted an image adapter using the GraLoRA technique into a LoRA checkpoint. The test scores were virtually identical after conversion (similarity 0.702 → 0.694, 0.260 → 0.269). Below are test images for the prompt “sks cat at the beach”:

| *Left: Image generated by GraLoRA. Right: Image generated by the same GraLoRA checkpoint converted to a LoRA checkpoint. The images quality is comparable.* |  | 

At the moment, we haven't implemented conversion for all PEFT techniques, but if there is demand, we will expand the support.

While working on the `PEFT` package, we noticed that LoRA has a lot of momentum behind it, even though other PEFT techniques are potentially better. Therefore, we set out to add benchmarks to PEFT that could paint a more objective picture of how well different PEFT techniques perform on different metrics.

Given the results we found, we can confidently conclude that LoRA is not a bad choice at all, but there are potentially better choices. Especially when checking the image generation benchmark, LoRA is beaten by other techniques. We discussed that besides metrics, other considerations must be taken into account when choosing the right PEFT technique. However, even then, we are pushing `PEFT` further to achieve feature parity between LoRA and those other techniques.

Our journey is far from finished; we want to extend and improve the existing benchmarks, and we also plan to add more benchmarks in the future. We ensured that it is easy for the community to contribute, so if this is something you would like to do, please open an issue on the `PEFT` repository and let us know how you would like to contribute.

If you take away only one thing from this article, it is that LoRA should not be the automatic default when choosing a PEFT technique for your use case. Given the unified API provided by `PEFT`, changing from one PEFT technique to another is as easy as switching one config in your code. And even if you stick with LoRA, check out all the variants that are supported in `PEFT`: DoRA, rs-LoRA, LoRA-FA etc. Give these other techniques a try and you might be pleasantly surprised.

Example: Changing from LoRA to OFT using `PEFT`:

```
from transformers import AutoModelForCausalLM
-from peft import LoraConfig, get_peft_model
+from peft import OFTConfig, get_peft_model
base_model = AutoModelForCausalLM.from_pretrained("meta-llama/Llama-3.2-3B", dtype="bfloat16")
-config = LoraConfig(target_modules=["q_proj", "v_proj"])
+config = OFTConfig(target_modules=["q_proj", "v_proj"])
model = get_peft_model(base_model, config)
```
