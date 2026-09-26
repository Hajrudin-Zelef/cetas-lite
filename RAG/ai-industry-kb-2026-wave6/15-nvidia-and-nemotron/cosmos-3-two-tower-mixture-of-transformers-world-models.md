---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/cosmos-3-two-tower-mixture-of-transformers-world-models
title: "Cosmos 3 — two-tower Mixture-of-Transformers world models"
domain: nvidia-and-nemotron
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Hugging Face", "Nvidia", "vLLM"]
dates: ["2026-05-31", "2026-06-01", "2026-07-15", "2026-07-20", "2026-08-13"]
keywords: ["attention", "attribution", "blackwell", "compute", "datacenter", "diffusion", "energy", "humanoid", "inference", "nvidia", "omni", "open weights"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7460, 7485]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: 1a41afae1d15ccd600b964370a4e94b58e49648b2c126268963853a0cbbcf2a7
---

# Cosmos 3 — two-tower Mixture-of-Transformers world models

### Cosmos 3 — two-tower Mixture-of-Transformers world models
- NVIDIA announced Cosmos 3 at its GTC Taipei event on June 1, 2026 (one research note dates the announcement ~May 31, 2026 — preserve the one-day discrepancy). [SECONDARY] [S11][S12][S13]
- Cosmos 3 is a family of omnimodal world foundation models for physical AI, unifying three capabilities that earlier Cosmos generations split across separate models: physical reasoning, world generation and action generation. [VENDOR] [S11][S13]
- Architecture: a two-tower Mixture-of-Transformers (MoT). The Reasoner tower is an autoregressive vision-language model (initialized from Qwen3-VL) that interprets text/image/video/audio and builds a world-state representation; the Generator tower is diffusion-based and produces physics-aware video, audio and action sequences conditioned on the reasoner's understanding. [VENDOR] [S11][S13]
- The towers have separate parameters per layer but interact via joint attention; information flows from reasoner to generator, and the reasoner can run independently while the generator always activates both towers. (single vendor coverage) [VENDOR] [S13]
- Two launch sizes: Cosmos 3 Nano = 16B total (8B reasoner + 8B generator), aimed at workstation inference (e.g., NVIDIA RTX PRO 6000); Cosmos 3 Super = 64B total (32B + 32B), aimed at datacenter deployment on Hopper/Blackwell for large-scale synthetic data and advanced physical reasoning. [VENDOR] [S11][S13]
- Cosmos 3 Edge = 4B on-device world model, launched July 15, 2026 with weights on Hugging Face July 20, 2026 (`nvidia/Cosmos3-Edge`); it targets Jetson Thor for real-time robot control at 15 Hz (640×360 resolution, 32 actions per inference). [SECONDARY] [S14][S15]
- NVIDIA claims Cosmos 3 Edge ranks #1 on VANTAGE-Bench for vision analytics among 4B-parameter models and state-of-the-art robot policy learning in its size class; specific scores were not disclosed. (single vendor coverage) [VENDOR] [S15]
- The Edge reasoner uses a 2B dense transformer and Qwen3-VL-compatible message conventions; reasoner context reaches up to 256K tokens. [SECONDARY] [S14][S16]
- Robotics operating modes: forward dynamics (action+image+text → future video, i.e., a learned simulator), inverse dynamics (text+video → action), and policy (image+text → video+action); the model emits numerical actions such as joint angles, gripper positions and trajectories. (single secondary coverage) [SECONDARY] [S16]
- Licensing: open weights under OpenMDW-1.1 (commercial use permitted, "Built on NVIDIA Cosmos" attribution); training scripts, deployment tools and datasets also open-sourced; code at github.com/nvidia/cosmos. [SECONDARY] [S13][S16]
- Access routes: Hugging Face (`nvidia/Cosmos3-Nano`, `nvidia/Cosmos3-Super`, `nvidia/Cosmos3-Edge`), build.nvidia.com (preview-gated), Cosmos NIM microservices (Reasoner NIM GA, Generator NIM v1.0.0), self-serve via vLLM-Omni, and third-party hosting (e.g., DeepInfra lists Nano at $0.0108/second). [SECONDARY] [S16][S53]
- A ready VLA-style policy fine-tune, Cosmos3-Nano-Policy-DROID (16B), was released alongside, trained on the DROID dataset. [SECONDARY] [S16][S54]
- CONTRADICTION (minor): announcement dated May 31, 2026 in one research note vs June 1, 2026 (GTC Taipei) in event records. Preserve both. [SECONDARY] [S12][S16]

### GR00T N1.7 and N2 — humanoid robot foundation models
- NVIDIA announced GR00T N1.7 in early access with commercial licensing, bringing generalized robot skills including advanced dexterous control to production-ready deployments. [VENDOR] [S17][S55]
- Industrial adopters named for the Isaac GR00T N line: AGIBOT, Humanoid, LG Electronics, NEURA Robotics and Noble Machines. [SECONDARY] [S17][S56]
- At the GTC keynote, Jensen Huang previewed GR00T N2, a next-generation robot foundation model based on DreamZero research with a new world-action-model architecture. [VENDOR] [S17][S56]
- NVIDIA's claim: N2 helps robots succeed at new tasks in new environments more than twice as often as leading vision-language-action models; it ranks No. 1 on MolmoSpaces and RoboArena for generalist robot policies. [VENDOR] [S17][S56]
- N2 was slated to be available by the end of the year (2026) at preview time — it remains a previewed, not generally available, model. [SECONDARY] [S17][S56]
- The GR00T stack sits on the Newton physics engine 1.0 plus the PhysX SDK (multiphysics simulation, dexterous manipulation) and the Jetson Thor robotic computing platform for sim-to-real deployment. (single secondary coverage) [SECONDARY] [S17]
- Separately, NVIDIA announced the Isaac GR00T Reference Humanoid Robot (June 1, 2026): a Unitree H2/H2 Plus chassis (~6 ft, 150 lb, 31 DOF) with Sharpa Wave tactile five-finger hands (22 DOF, 75 DOF total), head-mounted stereo camera (140° horizontal / 102° vertical FOV), wrist cameras, whole-body control with up to 120 Nm arm torque and 360 Nm leg torque, 7 kg rated / 15 kg peak arm payload, and Jetson Thor onboard compute. (single secondary coverage) [SECONDARY] [S18]
- LG Group and NVIDIA signed an MoU on August 13, 2026 (Santa Clara) to jointly develop a next-generation bipedal humanoid robot, targeting public unveiling in Q1 2027: Jetson Thor onboard compute, Isaac GR00T foundation model, and NVIDIA Halos for Robotics safety stack. LG Electronics supplies actuators, LG Innotek sensors, LG Energy Solution batteries — while LG also develops its own in-house robot foundation model (RFM) fed by manufacturing validation data. (single secondary coverage) [SECONDARY] [S19]
- The LG–NVIDIA scope extends beyond the robot: an AI factory pilot site powered by NVIDIA's Vera Rubin platform in H1 2027 (paving the way for an 80MW AI factory in Cheonan, South Korea, by H1 2028), CLOiD wheel-based robots on LG's Tennessee washing-machine line for real-world validation, a robot data platform with LG CNS, and an AI-defined vehicle platform on NVIDIA DRIVE Hyperion. (single secondary coverage) [SECONDARY] [S19]

