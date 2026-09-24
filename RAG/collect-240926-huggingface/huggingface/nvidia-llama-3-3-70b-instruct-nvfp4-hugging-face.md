---
id: collect-240926-huggingface/huggingface/nvidia-llama-3-3-70b-instruct-nvfp4-hugging-face
title: "The entry point of the program need to be protected for spawning processes."
domain: huggingface
role: reference
task: reference
actors: ["Meta", "Nvidia", "TensorRT-LLM", "United States"]
dates: []
keywords: ["benchmark", "blackwell", "fp4", "gpu", "inference", "llama", "memory", "nvfp4", "nvidia", "parameters", "tensorrt", "tensorrt-llm"]
source: docs/RAG/clean_en/huggingface/nvidia-llama-3-3-70b-instruct-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [1, 91]
sha256: 0748438555042fa3857dbdf73bd083fcea18fd4726236cd76a0004db2fd6a8fd
---

# The entry point of the program need to be protected for spawning processes.

<!-- source: https://huggingface.co/nvidia/Llama-3.3-70B-Instruct-NVFP4 -->

The NVIDIA Llama 3.3 70B Instruct FP4 model is the quantized version of the Meta's Llama 3.3 70B Instruct model, which is an auto-regressive language model that uses an optimized transformer architecture. For more information, please check here. The NVIDIA Llama 3.3 70B Instruct FP4 model is quantized with TensorRT Model Optimizer.

This model is ready for commercial/non-commercial use.  

This model is not owned or developed by NVIDIA. This model has been developed and built to a third-party’s requirements for this application and use case; see link to Non-NVIDIA (Meta-Llama-3.3-70B-Instruct) Model Card.

**Architecture Type:** Transformers  

**Network Architecture:** Llama3.3 

**Input Type(s):** Text 

**Input Format(s):** String 

**Input Parameters:** 1D (One Dimensional): Sequences 

**Other Properties Related to Input:** Context length up to 128K 

**Output Type(s):** Text 

**Output Format:** String 

**Output Parameters:** 1D (One Dimensional): Sequences 

**Other Properties Related to Output:** N/A 

**Supported Runtime Engine(s):** 

- Tensor(RT)-LLM 

**Supported Hardware Microarchitecture Compatibility:** 

- NVIDIA Blackwell 

**Preferred Operating System(s):** 

- Linux 

The model is quantized with nvidia-modelopt **v0.23.0**  

- Calibration Dataset: cnn_dailymail 
 ** Data collection method: Automated.
** Labeling method: Unknown.

**Engine:** Tensor(RT)-LLM 

**Test Hardware:** B200 

This model was obtained by quantizing the weights and activations of Meta-Llama-3.3-70B-Instruct to FP4 data type, ready for inference with TensorRT-LLM. Only the weights and activations of the linear operators within transformers blocks are quantized. This optimization reduces the number of bits per parameter from 16 to 4, reducing the disk size and GPU memory requirements by approximately 3.3x.

To deploy the quantized checkpoint with TensorRT-LLM LLM API, follow the sample codes below:

- LLM API sample usage:

```
from tensorrt_llm import LLM, SamplingParams
def main():
    prompts = [
        "Hello, my name is",
        "The president of the United States is",
        "The capital of France is",
        "The future of AI is",
    ]
    sampling_params = SamplingParams(temperature=0.8, top_p=0.95)
    llm = LLM(model="nvidia/Llama-3.3-70B-Instruct-FP4")
    outputs = llm.generate(prompts, sampling_params)
    # Print the outputs.
    for output in outputs:
        prompt = output.prompt
        generated_text = output.outputs[0].text
        print(f"Prompt: {prompt!r}, Generated text: {generated_text!r}")
# The entry point of the program need to be protected for spawning processes.
if __name__ == '__main__':
    main()
```
Please refer to the TensorRT-LLM llm-api documentation for more details.

The accuracy benchmark results are presented in the table below:

| **Precision** | **MMLU** | **GSM8K_COT** | **ARC Challenge** | **IFEVAL** | 
| BF16 | 83.3 | 95.3 | 93.7 | 92.1 | 
| FP4 | 81.1 | 92.6 | 93.3 | 92.0 | 

NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. When downloaded or used in accordance with our terms of service, developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse.

Please report security vulnerabilities or NVIDIA AI Concerns here.

- Downloads last month
- 298,405
