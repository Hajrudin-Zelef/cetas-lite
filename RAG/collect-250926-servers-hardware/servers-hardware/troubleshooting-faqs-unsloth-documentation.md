---
id: collect-250926-servers-hardware/servers-hardware/troubleshooting-faqs-unsloth-documentation
title: "troubleshooting-faqs-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Unsloth", "vLLM"]
dates: ["2025-06"]
keywords: ["deepseek", "fine-tuning", "gguf", "gpu", "inference", "inference engine", "llama", "llama.cpp", "lora", "memory", "throughput", "training"]
source: docs/RAG/clean4/troubleshooting-faqs-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 85]
sha256: d356f2785efbde8d3cba56f7b5264b254ebdfcc5243028fe6b729078c4cfc898
---

# troubleshooting-faqs-unsloth-documentation

If you're still encountering any issues with versions or dependencies, please use our Docker image which will have everything pre-installed.

**Try always to update Unsloth if you find any issues.**

`pip install --upgrade --force-reinstall --no-cache-dir --no-deps unsloth unsloth_zoo`

Unsloth works with any model supported by `transformers`. If a model isn’t in our uploads or doesn’t run out of the box, it’s usually still supported, some newer models may just need a small manual tweak due to our optimizations.

In most cases, you can enable compatibility by setting `trust_remote_code=True` in your fine-tuning script. Here’s an example using DeepSeek-OCR:

### Running in Unsloth works well, but after exporting & running on other platforms, the results are poor

You might sometimes encounter an issue where your model runs and produces good results on Unsloth, but when you use it on another platform like Ollama or vLLM, the results are poor or you might get gibberish, endless/infinite generations *or* repeated outputs**.**

- The most common cause of this error is using an **incorrect chat template****.** It’s essential to use the SAME chat template that was used when training the model in Unsloth and later when you run it in another framework, such as llama.cpp or Ollama. When inferencing from a saved model, it's crucial to apply the correct template.
- It might also be because your inference engine adds an unnecessary "start of sequence" token (or the lack of thereof on the contrary) so ensure you check both hypotheses!
- **Use our conversational notebooks to force the chat template - this will fix most issues.**

You can try reducing the maximum GPU usage during saving by changing `maximum_memory_usage`.

The default is `model.save_pretrained(..., maximum_memory_usage = 0.75)`. Reduce it to say 0.5 to use 50% of GPU peak memory or lower. This can reduce OOM crashes during saving.

First save your model to 16bit via:

Compile llama.cpp from source like below:

Then, save the model to F16:

On Mac devices, it seems like that BF16 might be slower than F16. Q8_K_XL upcasts some layers to BF16, so hence the slowdown, We are actively changing our conversion process to make F16 the default choice for Q8_K_XL to reduce performance hits.

To set up evaluation in your training run, you first have to split your dataset into a training and test split. You should **always shuffle the selection of the dataset**, otherwise your evaluation is wrong!

Then, we can set the training arguments to enable evaluation. Reminder evaluation can be very very slow especially if you set `eval_steps = 1` which means you are evaluating every single step. If you are, try reducing the eval_dataset size to say 100 rows or something.

A common issue when you OOM is because you set your batch size too high. Set it lower than 2 to use less VRAM. Also use `fp16_full_eval=True` to use float16 for evaluation which cuts memory by 1/2.

First split your training dataset into a train and test split. Set the trainer settings for evaluation to:

This will cause no OOMs and make it somewhat faster. You can also use `bf16_full_eval=True` for bf16 machines. By default Unsloth should have set these flags on by default as of June 2025.

If you want to stop the finetuning / training run since the evaluation loss is not decreasing, then you can use early stopping which stops the training process. Use `EarlyStoppingCallback`.

As usual, set up your trainer and your evaluation dataset. The below is used to stop the training run if the `eval_loss` (the evaluation loss) is not decreasing after 3 steps or so.

We then add the callback which can also be customized:

Then train the model as usual via `trainer.train() .`

If your model gets stuck at 90, 95% for a long time before you can disable some fast downloading processes to force downloads to be synchronous and to print out more error messages.

Simply use `UNSLOTH_STABLE_DOWNLOADS=1` before any Unsloth import.

Restart and run all, but place this at the start before any Unsloth import. Also please file a bug report asap thank you!

This means that your usage of `train_on_responses_only` is incorrect for that particular model. train_on_responses_only allows you to mask the user question, and train your model to output the assistant response with higher weighting. This is known to increase accuracy by 1% or more. See our **LoRA Hyperparameters Guide** for more details.

For Llama 3.1, 3.2, 3.3 type models, please use the below:

For Gemma 2, 3. 3n models, use the below:

If your speed seems slower at first, it’s likely because `torch.compile` typically takes ~5 minutes (or longer) to warm up and finish compiling. Make sure you measure throughput **after** it’s fully loaded as over longer runs, Unsloth should be much faster.

To disable use:

This is a critical error, since this means some weights are not parsed correctly, which will cause incorrect outputs. This can normally be fixed by upgrading Unsloth

`pip install --upgrade --force-reinstall --no-cache-dir --no-deps unsloth unsloth_zoo`

Then upgrade transformers and timm:

`pip install --upgrade --force-reinstall --no-cache-dir --no-deps transformers timm`

However if the issue still persists, please file a bug report asap!

See https://github.com/googlecolab/colabtools/issues/3409

In a new cell, run the below:

If you are citing the usage of our model uploads, use the below Bibtex. This is for Qwen3-30B-A3B-GGUF Q8_K_XL:

To cite the usage of our Github package or our work in general:

Last updated

Was this helpful?
