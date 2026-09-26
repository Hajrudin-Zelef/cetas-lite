---
id: collect-250926-servers-hardware/servers-hardware/fine-tuning-llms-on-intel-gpus-with-unsloth-2
title: "Fine-tuning LLMs on Intel GPUs with Unsloth"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Intel", "vLLM"]
dates: []
keywords: ["fine-tuning", "gpu", "intel", "agent", "agents", "inference", "lora", "memory", "training", "vllm"]
source: docs/RAG/clean4/Fine-tuning LLMs on Intel GPUs with Unsloth.md
source_anchor: ""
source_lines: [196, 326]
sha256: e7eb5f422385bf0e7a0214917641d529ed3e4c0f3435f5b9e7da775f488aa559
---

# Fine-tuning LLMs on Intel GPUs with Unsloth

        "-" * 20,
        f"Question:\n{q}",
        f"\nAnswer:\n{answer[0]}",
        f"\nResponse:\n{responses[0]}",
        f"\nExtracted:\n{extracted_responses[0]}",
    )
    return [2.0 if r == a else 0.0 for r, a in zip(extracted_responses, answer)]
def int_reward_func(completions, **kwargs) -> list[float]:
    responses = [completion[0]["content"] for completion in completions]
    extracted_responses = [extract_xml_answer(r) for r in responses]
    return [0.5 if r.isdigit() else 0.0 for r in extracted_responses]
def strict_format_reward_func(completions, **kwargs) -> list[float]:
    """Reward function that checks if the completion has a specific format."""
    pattern = r"^\n.*?\n\n\n.*?\n\n$"
    responses = [completion[0]["content"] for completion in completions]
    matches = [re.match(pattern, r) for r in responses]
    return [0.5 if match else 0.0 for match in matches]
def soft_format_reward_func(completions, **kwargs) -> list[float]:
    """Reward function that checks if the completion has a specific format."""
    pattern = r".*?\s*.*?"
    responses = [completion[0]["content"] for completion in completions]
    matches = [re.match(pattern, r) for r in responses]
    return [0.5 if match else 0.0 for match in matches]
def count_xml(text: str) -> float:
    count = 0.0
    if text.count("\n") == 1:
        count += 0.125
    if text.count("\n\n") == 1:
        count += 0.125
    if text.count("\n\n") == 1:
        count += 0.125
    count -= len(text.split("\n\n")[-1]) * 0.001
    if text.count("\n") == 1:
        count += 0.125
    count -= (len(text.split("\n")[-1]) - 1) * 0.001
    return count
def xmlcount_reward_func(completions, **kwargs) -> list[float]:
    contents = [completion[0]["content"] for completion in completions]
    return [count_xml(c) for c in contents]
if __name__ == "__main__":
    model, tokenizer = FastLanguageModel.from_pretrained(
        model_name="unsloth/Qwen3-0.6B",
        max_seq_length=max_seq_length,
        load_in_4bit=False,  # False for LoRA 16bit
        fast_inference=False,  # Enable vLLM fast inference
        max_lora_rank=lora_rank,
        gpu_memory_utilization=0.7,  # Reduce if out of memory
        device_map="xpu:0",
    )
    model = FastLanguageModel.get_peft_model(
        model,
        r=lora_rank,  # Choose any number > 0 ! Suggested 8, 16, 32, 64, 128
        target_modules=[
            "q_proj",
            "k_proj",
            "v_proj",
            "o_proj",
            "gate_proj",
            "up_proj",
            "down_proj",
        ],  # Remove QKVO if out of memory
        lora_alpha=lora_rank,
        use_gradient_checkpointing="unsloth",  # Enable long context finetuning
        random_state=3407,
    )
    dataset = get_gsm8k_questions()
    training_args = GRPOConfig(
        learning_rate=5e-6,
        adam_beta1=0.9,
        adam_beta2=0.99,
        weight_decay=0.1,
        warmup_ratio=0.1,
        lr_scheduler_type="cosine",
        optim="adamw_torch",
        logging_steps=1,
        per_device_train_batch_size=1,
        gradient_accumulation_steps=1,  # Increase to 4 for smoother training
        num_generations=4,  # Decrease if out of memory
        max_prompt_length=max_prompt_length,
        max_completion_length=max_seq_length - max_prompt_length,
        # num_train_epochs=1,  # Set to 1 for a full training run
        max_steps=20,
        save_steps=250,
        max_grad_norm=0.1,
        report_to="none",  # Can use Weights & Biases
        output_dir="outputs",
    )
    trainer = GRPOTrainer(
        model=model,
        processing_class=tokenizer,
        reward_funcs=[
            xmlcount_reward_func,
            soft_format_reward_func,
            strict_format_reward_func,
            int_reward_func,
            correctness_reward_func,
        ],
        args=training_args,
        train_dataset=dataset,
        dataset_num_proc=1,  # Recommended on Windows
    )
    trainer.train()
```
{% endcode %}
## Troubleshooting
### Out of Memory (OOM) Errors
If you run out of memory, try these solutions:
1. **Reduce batch size:** Lower `per_device_train_batch_size`.
2. **Use a smaller model:** Start with a smaller model to reduce memory requirements.
3. **Reduce sequence length:** Lower `max_seq_length`.
4. **Reduce LoRA rank:** Use `r=8` instead of `r=16` or `r=32`.
5. **For GRPO, reduce number of generations:** Lower `num_generations`.
### (Windows Only) Intel Ultra AIPC iGPU Shared Memory
For Intel Ultra AIPC with recent GPU drivers on Windows, the shared GPU memory for the integrated GPU typically defaults to **57%** of system memory. For larger models (e.g., **Qwen3-32B**), or when using longer max sequence length, larger batch size, LoRA adapters with larger LoRA rank, etc., during fine-tuning, you could increase available VRAM by raising the percentage of system memory allocated to the iGPU.
You can adjust this by modifying the registry:
* Path: `Computer\HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\GraphicsDrivers\MemoryManager`
* Key to change:\
  `SystemPartitionCommitLimitPercentage` (set to a larger percentage)
---
# Agent Instructions
This documentation is published with GitBook. GitBook is the documentation platform designed so that both humans and AI agents can read, navigate, and reason over technical content effectively. Learn more at gitbook.com.
## Querying This Documentation
If you need additional information that is not directly available in this page, you can query the documentation dynamically by asking a question.
Perform an HTTP GET request on the current page URL with the `ask` query parameter, and the optional `goal` query parameter:
```
GET https://unsloth.ai/docs/get-started/install/intel.md?ask=&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
