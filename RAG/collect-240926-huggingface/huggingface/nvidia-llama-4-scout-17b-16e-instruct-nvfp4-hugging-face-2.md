---
id: collect-240926-huggingface/huggingface/nvidia-llama-4-scout-17b-16e-instruct-nvfp4-hugging-face-2
title: "NVIDIA Open Model License Agreement"
domain: huggingface
role: reference
task: reference
actors: ["California", "EU", "Meta", "Microsoft", "Nvidia", "TensorRT-LLM", "United States"]
dates: []
keywords: ["license", "nvidia", "agent", "benchmark", "blackwell", "distribution", "fp4", "gpu", "inference", "liability", "llama", "memory"]
source: docs/RAG/clean_en/huggingface/nvidia-llama-4-scout-17b-16e-instruct-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [64, 200]
sha256: 06a5b4476270569c3e58d87e2946704c1e8dbe9297bdc8938abf342ebfae1a5e
---

# NVIDIA Open Model License Agreement

**Unless required by applicable law or agreed to in writing, NVIDIA provides the Model on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied, including, without limitation, any warranties or conditions of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE. You are solely responsible for determining the appropriateness of using or redistributing the Model, Derivative Models and outputs and assume any risks associated with Your exercise of permissions under this Agreement.**

## **7. Limitation of Liability**

**In no event and under no legal theory, whether in tort (including negligence), contract, or otherwise, unless required by applicable law (such as deliberate and grossly negligent acts) or agreed to in writing, will NVIDIA be liable to You for damages, including any direct, indirect, special, incidental, or consequential damages of any character arising as a result of this Agreement or out of the use or inability to use the Model, Derivative Models or outputs (including but not limited to damages for loss of goodwill, work stoppage, computer failure or malfunction, or any and all other commercial damages or losses), even if NVIDIA has been advised of the possibility of such damages.**

## 8. Indemnity

You will indemnify and hold harmless NVIDIA from and against any claim by any third party arising out of or related to your use or distribution of the Model, Model Derivatives or outputs.

## 9. Feedback

NVIDIA appreciates your feedback, and You agree that NVIDIA may use it without restriction or compensation to You.

## 10. Governing Law

This Agreement will be governed in all respects by the laws of the United States and the laws of the State of Delaware, without regard to conflict of laws principles or the United Nations Convention on Contracts for the International Sale of Goods. The state and federal courts residing in Santa Clara County, California will have exclusive jurisdiction over any dispute or claim arising out of or related to this Agreement, and the parties irrevocably consent to personal jurisdiction and venue in those courts; except that, either party may apply for injunctive remedies or an equivalent type of urgent legal relief in any jurisdiction.

## 11. Trade and Compliance

You agree to comply with all applicable export, import, trade and economic sanctions laws and regulations, as amended, including without limitation U.S. Export Administration Regulations and Office of Foreign Assets Control regulations. These laws include restrictions on destinations, end-users and end-use.

Log in or Sign Up to review the conditions and access this model content.

The NVIDIA Llama 4 Scout 17B 16E Instruct FP4 model is the quantized language model of the Meta's Llama 4 Scout 17B 16E model, which is an auto-regressive language model that uses a mixture-of-experts (MoE) architecture and incorporates early fusion for native multimodality. For more information, please check here.

This model is ready for commercial and non-commercial use.  

GOVERNING TERMS: Use of this model is governed by the  NVIDIA Open Model License.

ADDITIONAL INFORMATION: Llama4 Community License Agreement. Built with Llama.

Global, except in European Union 

Developers looking to take off the shelf pre-quantized models for deployment in AI Agent systems, chatbots, RAG systems, and other AI-powered applications.

Huggingface:  Jul 28th, 2025 via [https://huggingface.co/nvidia/Llama-4-Scout-17B-16E-Instruct-FP4] 

 

**Architecture Type:** Transformers  

**Network Architecture:** Llama4 

**Input Type(s):** Multilingual text, and up to 5 images 

**Input Format(s):** String, Images 

**Input Parameters:** One-Dimensional (1D), Two-Dimensional (2D) 

**Other Properties Related to Input:** Context length up to 1M 

**Output Type(s):** Multilingual text and code 

**Output Format:** String 

**Output Parameters:** One-Dimensional (1D) 

**Other Properties Related to Output:** Context length up to 1M

Our AI models are designed and/or optimized to run on NVIDIA GPU-accelerated systems. By leveraging NVIDIA’s hardware (e.g. GPU cores) and software frameworks (e.g., CUDA libraries), the model achieves faster training and inference times compared to CPU-only solutions. 

  

**Supported Runtime Engine(s):** 

- TensorRT-LLM 

**Supported Hardware Microarchitecture Compatibility:** 

- NVIDIA Blackwell 
- NVIDIA Hopper 
- NVIDIA Ampere 

**[Preferred/Supported] Operating System(s):** 

- Linux 

The model is quantized with nvidia-modelopt **v0.33.0**  

- Calibration Dataset: cnn_dailymail 
 ** Data collection method: Automated
** Labeling method: Automated.
- Evaluation Datasets: **Datasets: MMMU Pro, GPQA Diamond, HLE, LiveCodeBench, SciCode, HumanEval, AIME 2024, MATH-500 
 ** Data collection method: Hybrid: Automated, Human
** Labeling method: Hybrid: Human, Automated.

**Engine:** TensorRT-LLM 

**Test Hardware:** B200 

This model was obtained by quantizing the weights and activations of Llama 4 Scout 17B 16E Instruct to FP4 data type, ready for inference with TensorRT-LLM. Only the weights and activations of the linear operators within transformer blocks are quantized. This optimization reduces the number of bits per parameter from 16 to 4, reducing the disk size and GPU memory requirements by approximately 3.3x.

To serve the quantized checkpoint with TensorRT-LLM, follow the sample commands below with the TensorRT-LLM GitHub repo:

- LLM API sample usage:

```
import asyncio
from tensorrt_llm import LLM, SamplingParams
def main():
    llm = LLM(model="nvidia/Llama-4-Scout-17B-16E-Instruct-FP4", attn_backend="FLASHINFER", backend="pytorch", tensor_parallel_size=8)
    prompts = [
        "Hello, my name is",
        "The president of the United States is",
        "The capital of France is",
        "The future of AI is",
    ]
    sampling_params = SamplingParams(temperature=0.8, top_p=0.95)
    async def task(prompt: str):
        output = await llm.generate_async(prompt, sampling_params)
        print(
            f"Prompt: {output.prompt!r}, Generated text: {output.outputs[0].text!r}"
        )
    async def main():
        tasks = [task(prompt) for prompt in prompts]
        await asyncio.gather(*tasks)
    asyncio.run(main())
# The entry point of the program need to be protected for spawning processes.
if __name__ == '__main__':
    main()
```
The accuracy benchmark results are presented in the table below:

| **Precision** | **MMLU Pro** | **GPQA Diamond** | **HLE Challenge** | **SciCode** | **MATH-500** | **AIME 2024** | 
| Llama-4-Scout-17B-16E-Instruct <sup>1</sup> | 75 | 57 | 4 | 26 | 82 | 30 | 
| Llama-4-Scout-17B-16E-Instruct-FP4 | 74 | 56 | 4 | 24 | 81 | 31 | 

<sup>1</sup> Reference scores for Llama-4-Scout-17B-16E-Instruct sourced from artificialanalysis.


NVIDIA believes Trustworthy AI is a shared responsibility and we have established policies and practices to enable development for a wide array of AI applications. Developers should perform safety testing and tuning tailored to their specific applications of the model. When downloaded or used in accordance with our terms of service, developers should work with their internal model team to ensure this model meets requirements for the relevant industry and use case and addresses unforeseen product misuse.

Please report security vulnerabilities or NVIDIA AI Concerns here.

- Downloads last month
- 6,431
