---
id: collect-261001-general-networking/general-networking/how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows-1
title: "how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "robotics", "gpu", "gpus", "latency", "memory", "parameters", "throughput", "training"]
source: docs/RAG/collect-261001-general-networking/how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows.md
source_anchor: ""
source_lines: [1, 106]
sha256: 32c73a854125febff69f28ca28f4eb9964bcf0d2f73b672a0a41eb0679ee489e
---

# how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows

MuJoCo Warp (MJWarp), built on NVIDIA Warp, takes compatible MuJoCo models into that GPU-scale regime. In this article, we will move an SO-101 follower arm from a familiar MuJoCo workflow to as many as 2,048 parallel MJWarp environments and examine the technology and validation steps that make the transition possible.

*Figure 1. How MJWarp connects Python to GPU simulation. MuJoCo loads and compiles the MJCF model; MJWarp implements the physics in NVIDIA Warp, which compiles CUDA kernels to advance simulation states on NVIDIA GPUs.*

*This is the second article in our State of Simulation for Physical AI series. The first article mapped the robot-simulation landscape. Here, we prepare and scale the simulation environment; we do not train a policy. The later Newton and Isaac Lab installments cover the next integration layers.*

| Layer | Role in the stack | 
|---|---|
| **NVIDIA Warp** | Python kernel language: single instruction, multiple threads (SIMT), autodiff, PyTorch/JAX interop | 
| **MJWarp** | MuJoCo physics on Warp: same MJCF, batched GPU throughput | 
| **Your scene (SO-101)** | Familiar Menagerie / Robot Studio assets + task geometry | 
| **Next (Newton / Isaac Lab)** | Multi-solver API, USD, sensors, managers, training loops | 

**Decision shortcut:**

| If you need… | Reach for… | 
|---|---|
| Single-robot MPC / teleop | MuJoCo CPU | 
| Max throughput on raw MuJoCo physics | **MJWarp** (or*mjlab* ) | 
| JAX training recipes | *MuJoCo Playground* / MJX (impl='warp') | 
| Multi-solver + Isaac Lab integration | **Newton** — next post in this series | 

***NVIDIA Warp*** is a Python framework for writing high-performance, GPU-accelerated kernels. Warp lets developers author statically typed kernels in Python and compiles them for CPU or CUDA execution. The first launch builds and caches a native module; later launches reuse it. The kernel language is a performance-oriented subset of Python, while ordinary Python remains responsible for configuration, allocation, and launch orchestration.

This small robotics-oriented kernel advances point positions under gravity. One logical thread handles one point, so the same code scales from two points to millions without introducing GPU terminology into the control flow.

The three value propositions of Warp are:

| Pillar | What you get | 
|---|---|
| **Performance** | Native-CUDA speed via JIT compilation, kernel fusion, and CUDA Graphs | 
| **Ease of use** | Pure Python authoring with built-in vectors, matrices, quaternions, BVHs, hash grids, sparse matrices, and tile primitives | 
| **Capability** | Differentiable kernels and DLPack-style interop so simulation can sit inside an ML training loop | 

```
import numpy as np
import warp as wp
@wp.kernel
def integrate(   positions: wp.array[wp.vec3],
   velocities: wp.array[wp.vec3],
   dt: float,
0.0, 0.0, -9.81) * dt
   positions[i] += velocities[i] * dt
wp.init()
device = "cuda:0" if wp.is_cuda_available() else "cpu"
start = np.array([[0.0, 0.0, 0.5], [0.2, 0.0, 0.5]], dtype=np.float32)
positions = wp.array(start, dtype=wp.vec3, device=device)
velocities = wp.zeros_like(positions)
wp.launch(
   integrate,
   dim=len(start),
   inputs=[positions, velocities, 0.01],
   device=device,
)
wp.synchronize_device(device)
print(positions.numpy())
```
- **Explicit parallel work.** wp.tid() identifies the point, contact, body, or world owned by the current logical thread.
- **Explicit device arrays.** An array lives on the selected device. Calling .numpy() on a CUDA array synchronizes and copies it to CPU memory; it is not a zero-copy path. For a device-resident PyTorch or JAX pipeline, use Warp’s framework adapters or DLPack-compatible sharing instead.
- **Composable kernel launches.** A program can launch a sequence of focused kernels and capture supported CUDA work into a graph to reduce repeated dispatch overhead. Graph capture replays launches against existing buffers; it does not fuse arbitrary kernels.

Two further Warp capabilities are worth knowing, even though neither is used in the SO-101 workflow in this article. Warp kernels are **differentiable**: a wp.Tape records the forward kernel launches made inside its context and replays their adjoints in reverse when backward() is called, which is why teams build differentiable geometry, CFD, and custom physics in Warp, including CAE workflows for simulation and design optimization. Warp also supports **deterministic execution**, introduced in *Warp 1.15*: GPU atomics are scheduler-dependent by default, so repeated launches of the same kernel can differ slightly, and the opt-in deterministic modes trade some performance for reproducible ordering in simulation, validation, and regression tests. These are Warp capabilities, not guarantees of differentiability or determinism for an entire MJWarp rollout. See the Warp documentation on differentiability and deterministic execution for the details.

**Try Warp:** pip install warp-lang (≥ 1.15 for GPU determinism), then python -m warp.examples.browse, or the *tutorial notebooks*.

A robot simulator repeatedly computes what happens next: given the current joint positions, velocities, controls, and contacts, it advances the scene by one small timestep. In this article, a **world** means one independent copy of that scene and its state. One world might contain the SO-101 arm reaching for a cube; another can contain the same arm starting from a slightly different pose.

MuJoCo and MJWarp can run the same compatible robot and task, but they organize the work differently. MuJoCo naturally suits developing and inspecting one or a few CPU worlds. **MJWarp** is a **NVIDIA Warp** implementation of **MuJoCo**’s physics pipeline that places the model and a batch of independent states on NVIDIA GPUs; one call to mjw.step advances the entire batch.

MJWarp’s value is not necessarily a faster step for one world. It is the ability to advance hundreds or thousands together, giving the GPU enough parallel work to improve **aggregate throughput**, the total world-steps completed per second. That favors reinforcement learning and large-scale sampling, where collecting experience matters more than minimizing one environment’s latency.

This blog covers the following:

1. validate one MuJoCo world,
2. move it to MJWarp, form a batch,
3. verify it, and measure it correctly.

Solver tuning, Jacobian representation, and specialized multi-GPU or determinism topics are not required for this migration and can be covered separately.

Then, the distinction is precise:

- **Latency** is wall-clock time for one simulation step.
- **Aggregate throughput** is the total number of world-steps completed per measured wall-clock second.

- The core API transition is small:

| MuJoCo host workflow | MJWarp workflow | 
|---|---|
| mujoco.MjModel | mjw.put_model(mjm) creates a device model | 
| mujoco.MjData | mjw.put_data(mjm, mjd, ...) preserves and batches an existing state | 
| mujoco.mj_step(mjm, mjd) | mjw.step(m, d) advances every world in d | 
| Host arrays such as mjd.ctrl | Batched device arrays such as d.ctrl with shape (nworld, nu) | 

Use mjw.make_data() when default/fresh state is intended. Use mjw.put_data() when the exact initialized MuJoCo state must cross the migration boundary.

Allocating batched resources requires defining the following parameters (refer to *Batch sizes*):

| Parameter | Meaning | 
|---|---|
| nworld | Total number of parallel environments | 
| nconmax | Expected contacts **per individual world** (overall capacity ≈ nconmax * nworld) | 
| naconmax | Alternative setting: global maximum contacts **across all environments combined** (takes precedence if both are defined) | 
| njmax | Hard upper limit on constraints **per world** | 

1. **CUDA graph capture**:mjw.step is many kernel launches; capture once, replay often:

