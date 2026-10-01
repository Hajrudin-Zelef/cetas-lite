---
id: collect-261001-general-networking/general-networking/how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows-3
title: "how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "benchmark", "benchmarks", "compute", "gpu", "humanoid", "latency", "memory", "throughput"]
source: docs/RAG/collect-261001-general-networking/how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows.md
source_anchor: ""
source_lines: [215, 314]
sha256: fd4a9860ef651d9ff9295bf369360397c5e53b842aa6c0e6542b41468cf01ff4
---

# how-to-use-nvidia-warp-and-mjwarp-to-accelerate-robotics-simulation-and-learning-workflows

```
def simulate_frame() -> None:
    ctrl = controller.step(mjm, mjd, frame_dt)
    for _ in range(sim_substeps):
        mjd.ctrl[: mjm.nu] = ctrl
        wp.copy(d.ctrl, wp.array(mjd.ctrl[None, :], dtype=wp.float32, device=device))
        mjw.step(m, d)
        mjd.qpos[:] = d.qpos.numpy()[0]
        mjd.qvel[:] = d.qvel.numpy()[0]
    mujoco.mj_forward(mjm, mjd)
```
The .numpy() reads synchronize and copy data to the host on every substep, so this is a task-validation path, not a throughput benchmark. It keeps inverse kinematics, viewing, and task checks on the host. After copying qpos and qvel, call mujoco.mj_forward(mjm, mjd) to refresh derived host quantities such as mjd.xpos before using them for control, viewing, or the stack check. Reading those fields after the loop does not refresh them automatically. Gate 4 removes these per-step host copies from the throughput path.

MJWarp allocates contact and constraint buffers before stepping. Exceeding those capacities invalidates the affected rollout for verification or benchmarking, even when execution continues with an overflow warning rather than an exception. Increase the relevant limit and rerun the task. Larger buffers use more GPU memory, so verify capacity over the full task before tightening the allocation.

Set contact and constraint limits for the robot and task being simulated. The SO-101 profile uses nconmax=128 and njmax=300 as starting capacities. Check that these limits are sufficient during the most contact-heavy part of the task:

```
d = mjw.make_data(mjm, nworld=nworld, nconmax=spec.nconmax, njmax=spec.njmax)
```
Size them against the most contact-heavy moment of the task, for pick-and-place, the instant both jaws and the table touch a cube, not the arm hovering in free space. An overflow is reported rather than raised: with Option.warn_overflow at its default, MJWarp prints the budget to increase (“narrowphase overflow - please increase nconmax to …”) to the terminal running your script or the viewer, and flags the affected worlds in Data.overflow for you to read back after a step. Only mjw.put_data raises an error outright, because it can compare the budgets against a MuJoCo state it already holds. mjwarp-testspeed --measure_alloc reports the contacts and constraints a scene actually consumed, and it aborts the rollout with the offending world IDs as soon as any world overflows. Treat those reports as failures: raise the limit and re-run before trusting either the trajectory or the benchmark, then tighten again whenever the model, collision geometry, or task changes.

Once one-world parity passes, reallocate at the target size and replicate the initialized state across the batch. Two things change relative to Gate 2: nworld, and the fact that nothing crosses the PCIe bus per step.

```
nworld = 2_048
d = mjw.make_data(mjm, nworld=nworld, nconmax=spec.nconmax, njmax=spec.njmax)
wp.copy(d.qpos, wp.array(np.tile(mjd.qpos, (nworld, 1)), dtype=wp.float32, device=device))
wp.copy(d.qvel, wp.array(np.tile(mjd.qvel, (nworld, 1)), dtype=wp.float32, device=device))
wp.copy(d.ctrl, wp.array(np.tile(mjd.ctrl, (nworld, 1)), dtype=wp.float32, device=device))
mjw.forward(m, d)
with wp.ScopedCapture() as capture:
    mjw.step(m, d)
step_graph = capture.graph
```
np.tile gives every world the same starting state, which is the right baseline for a throughput measurement; per-world randomization would instead write different rows of d.qpos on the device.

CUDA Graphs reuse the model and data buffers captured here. Update d.ctrl in place between replays, and capture a new graph after replacing buffers, changing nworld, or rebuilding the model. Graph capture requires CUDA.

*Figure 3. Scaling the SO-101 task from one CPU world to 2,048 independent GPU states using the same compatible model. A single MJWarp step advances the full batch. This conceptual illustration highlights aggregate throughput, measured as world-steps per wall-clock second.*

GPU launches are asynchronous, so a naive timer measures how fast Python queued work, not how fast the GPU finished it. Warm up first — the first launches pay kernel compilation and allocation — then synchronize immediately before and after the timed region:

```
import time
for _ in range(10): # warm-up: compilation, allocation, caches
    wp.capture_launch(step_graph)
wp.synchronize()
t0 = time.perf_counter()
for _ in range(200):
    wp.capture_launch(step_graph)
wp.synchronize() # without this you time the queue, not the work
elapsed = time.perf_counter() - t0
total = 200 * nworld
print(f"{total / elapsed:,.0f} world-steps/second")
```
Report both aggregate world-steps per second and milliseconds per batched step, together with the batch size. Use the measured curve to identify where additional worlds improve throughput and where memory or compute limits reduce the benefit. Results depend on the scene, simulation settings, and hardware; a one-world latency comparison does not establish batched throughput.

To see that curve on your own hardware, scaling_study.py sweeps the batch size and prints ms/step alongside throughput and speedup:

```
cd /tutorials/sim2real-blogs/notebooks/mujoco/part2
python solutions/so101_mjwarp_solution.py --headless-steps 600     # parity, needs CUDA
python scaling_study.py --worlds 1 64 1024 2048 8192 --steps 100
```
**Warp (kernel layer)**

 pip install warp-lang → python -m warp.examples.browse → *docs* · *GitHub*

**MJWarp (GPU MuJoCo)**

 pip install mujoco-warp → mjwarp-viewer benchmarks/humanoid/humanoid.xml → *docs* · *GitHub* · *Colab tutorial*

**SO-101 context**

 *SO-101 sim-to-real course* · *Physical AI learning paths*

**Train on top of MJWarp**

 *mjlab* · *MuJoCo Playground* · Isaac Lab + Newton (upcoming posts)

This post covered raw **Warp → MJWarp**: GPU kernels, batched stepping, and an SO-101 scene using mjw.step.

Next, we will port the same MJCF environment into **Newton**, using MuJoCo Warp as its rigid-body solver (newton.solvers.SolverMuJoCo). Newton will manage the model, state, controls, and contacts, while MJWarp runs underneath.

You will also see what Newton adds: multi-format assets, swappable solvers, sensors/IK helpers, and an Isaac Lab path.

The migration guide continues with the same SO-101 task and its optional reBot profile, explaining the changes required by Newton and the separate Isaac Lab integration.

If you build something with Warp or MJWarp, open an issue on the linked repositories or find us on Discord *NVIDIA Omniverse*.

- **Blog 1:** *The State of Simulation for Physical AI: An Overview* — Post 1 of this series.
- *NVIDIA Warp — GitHub* ·*Documentation* ·*v1.15.0 release (GPU determinism)* ·*Deterministic execution guide*
- *MuJoCo Warp — GitHub* ·*Official MJWarp docs*
- *Build Accelerated, Differentiable Computational Physics Code for AI with NVIDIA Warp*
- *Introducing Tile-Based Programming in Warp 1.5.0*
- *mjlab* ·*arXiv:2601.22074*
- *MuJoCo Playground*
- *NVIDIA SO-101 sim-to-real course*
- *Newton* next post: MJWarp as SolverMuJoCo and porting this environment
