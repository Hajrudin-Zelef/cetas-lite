---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-24
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Huawei", "LongCat", "MiniMax", "Moonshot", "Nvidia", "Z.ai"]
dates: ["2026-09-27"]
keywords: ["agent", "ascend", "asic", "attention", "attribution", "benchmarks", "cyber", "deepseek", "glm", "kimi", "mai", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [2397, 2470]
sha256: 5bc06772ba46cc3191ae6eb62b62186049b4116d92016f9cdb1002ebb6e55c22
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

| Mois | Sorties (modèles du volume) |
|---|---|
| Février 2026 | Qwen3.5-397B-A17B · Ling-2.5-1T / Ring-2.5-1T · Ming-Flash-Omni-2.0 · Qwen3-Coder-Next · GLM-5 (fondateur) |
| Mars 2026 | MiMo-V2-Pro/Omni/TTS (fermés) · MiniMax-M2.7 |
| Avril 2026 | Qwen3.6-Plus (fermé) · Qwen3.6-35B-A3B · Qwen3.6-Max-Preview (fermé) · Kimi K2.6 · MiMo-V2.5/V2.5-Pro · DeepSeek V4-Pro/Flash · Ling-2.6-flash |
| Mai 2026 | Qwen3.7-Max/Plus (fermés) |
| Juin 2026 | MiniMax-M3 · Kimi K2.7 Code · GLM-5.2 · LongCat-2.0 |
| Juillet 2026 | Tencent Hy3 · MiniMax-H3 · Ling-3.0-flash · Kimi K3 |
| Août 2026 | Qwen3.8-Max (+ snapshot 0902 le 02/09) · Qwen3.8-27B · GLM-5.3 · GLM-5.3-Flash |
| Septembre 2026 | Qwen3.8-Flash · humain-m3 · Ling-3.0-flash-VL/Fin · DeepSeek V4.1-Flash · MiMo-V2.6 (Pro/Flash/Distill) |

**Avril et août** sont les deux mois les plus denses (5–6 sorties chacun) : avril = le choc DeepSeek V4 +
la rupture « frontier-closed » d'Alibaba ; août = la contre-offensive ouverte (Qwen3.8-Max, GLM-5.3).
Septembre est le mois des **remplacements** (V4-Flash retiré, V4-Pro déprécié, snapshots). Le rythme
moyen : **~4 modèles significatifs par mois** — intenable pour une veille manuelle, d'où l'importance
des sources agrégées (Artificial Analysis, LMArena) et des changelogs officiels.

## 209. Les 10 faits à retenir en 30 secondes

1. **Kimi K3** (2,8T/104B, juillet 2026) est le plus gros open-weight de l'histoire — et il est chinois.
2. Le **1M de contexte** est devenu le standard (12 modèles du volume), grâce à l'attention linéaire/hybride.
3. Le token d'entrée est tombé à **$0,075/M** (Ling-3.0-flash) — divisé par ~19 en six semaines sur le segment efficace.
4. Un **dense 27B ouvert** (Qwen3.8-27B) bat un flagship fermé occidental sur le code.
5. L'**attention quadratique est morte** en Chine : GDN, KDA, DSA+IndexShare, hybride CSA/HCA/SWA — et les briques circulent (KDA Moonshot → Ant).
6. La **guerre des Flash** se joue à $0,01 près ($0,14–0,16/M input chez 4 labos).
7. **GLM-5.3** markete des capacités cyber offensives (53 CVE revendiqués) — première assumée.
8. La pile chinoise tourne **sans NVIDIA** : Ascend (GLM-5), 50K ASIC (LongCat-2.0), 100K puces (GLM-5.3-Flash).
9. Les **licences** sont le champ de mines : 5 modèles majeurs à licence non vérifiée, 1 restriction territoriale (MiniMax-H3), 1 clause d'attribution (Kimi).
10. Un **ID d'API est un pointeur mutable** : V4-Pro a failli changer de modèle sous les utilisateurs en 96 heures.

## 210. Index des modèles (renvoi aux sections)

- **DeepSeek** : V4 (§5), V4-Pro (§6), V4-Flash (§7), V4.1-Flash (§8), V4.1 Pro — annoncé (§9), prix (§11, §156), tableau (§12)
- **Qwen** : 3.5-397B (§15), 3.5-122B/35B (§16), 3.5-27B et petites tailles (§17), 3.6-35B (§19), 3.6-27B (§20), 3.6-Max-Preview (§21), 3.6-Plus (§22), 3.7-Max (§25), 3.7-Plus (§26), 3.8-Max (§27–28), 3.8-27B (§29), 3.8-Flash (§30), tiers API (§31, §199), Coder-Next (§32), Coder-480B (§33), déploiement local (§34, §175), tableaux (§24, §35), stratégie open/fermé (§14, §207)
- **Kimi** : K2.6 (§39), K2.7 Code (§40), K3 (§41), KDA+MLA (§42), Mooncake (§160), essaims (§161), licence Modified MIT (§38), tableau (§43)
- **GLM** : 5.2 (§45), 5.3 (§46, §202), 5.3-Flash / Ox Alpha (§47), MLA+DSA+IndexShare (§48), ZCode (§162), Coding Plan (§186), retenue sécurité (§185), tableau (§50)
- **MiMo** : 7B (§52), VL/Audio-7B (§53), V2-Flash (§54), V2-Pro/Omni/TTS (§55), V2.5 (§56), V2.6-Pro-RL (§57), V2.6-Flash-RL + UltraSpeed (§58), Distill-Qwen-9B (§59), RL en direct (§60), controverse Anthropic (§61), tableau (§62)
- **Ling** : séries Ling/Ring/Ming (§65), 2.0 (§66), 2.5-1T (§67), Ming-Omni (§68), 2.6-flash (§69), 3.0-flash (§70), 3.0-flash-VL (§71), 3.0-tiny (§72), 3.0-flash-Fin (§73, §165), KDA+MLA (§74), HiCache (§75), tableau (§76)
- **LongCat** : Flash-Chat (§78), Flash-Thinking (§79), Next (§80, §191), 2.0 (§81, §166), ScMoE (§82), tableau (§83)
- **MiniMax** : M2→M2.7 (§85–86, §203), M3 (§87, §167), H3 (§88–89), humain-m3 (§90), tableau (§91)
- **Tencent** : Hy3 Preview (§93), Hy3 (§94–95, §204), fenêtre gratuite (§168), tableau (§96)
- **Synthèses** : grand tableau (§97), timeline (§98, §208), que retenir (§99–102, §209), puces chinoises (§103), Qwen4 (§104), limites (§105), face-à-face (§169–172, §195–197), RAG/local (§174–175), check-list prod (§200), IDs API (§199)

---

*Fin du volume 1 — Encyclopédie des modèles IA : la Chine (février 2026 → 27 septembre 2026).*
*210 sections numérotées. Tous les faits sont issus de la recherche documentaire de septembre 2026 ; les points marqués « non vérifié au 27/09/2026 » le restent jusqu'à confirmation par source primaire.*

## 211. Ce que ce volume ne couvre pas (et pourquoi)

Transparence sur les frontières du volume — ce qui est exclu l'est par consigne ou par méthode, pas par
oubli. **Exclus par consigne éditoriale (→ volume 2, Occident)** : NVIDIA Nemotron (Nano 2 12B, Nemotron 3
Nano 30B-A3B, Super 120B-A12B, Cascade-2, Nano Omni, Ultra 550B-A55B, 3.5 Content Safety, 3.5 Lightning
30B-A3B) et NousResearch Hermes (Hermes 4, Hermes 4.3-36B, Hermes Agent, GEPA, Psyche, atropos) — tous
documentés dans les dossiers de recherche mais américains, donc hors périmètre Chine. **Couverts comme
fondations antérieures à février 2026** (fiches courtes, pas de détail 2026) : MiMo-7B (avril 2025),
LongCat-Flash-Chat/Thinking (août/sept. 2025), MiniMax-M2 originel (oct. 2025), GLM-5 (février 2026 —
fondateur de la lignée 5.x), Kimi K2.5 (janvier 2026, mentionné en contexte uniquement). **Non traités
faute de matière vérifiable** : MiMo-VL-7B et MiMo-Audio-7B (fiches prudentes, §53), Ming-Flash-Omni-2.0
(annonce sans specs, §68), les déclinaisons fines (ex. variantes Base/SFT/RL de chaque modèle — citées
quand connues, pas détaillées). **Non recherchés** (hors périmètre, à noter pour la veille) : d'éventuels
« DeepSeek V5 », « MiMo-V3 », « Ling-4.0 », « LongCat-3.0 », « Kimi K4 », « Qwen4 » — rien de tel n'est
apparu dans les résultats de septembre 2026. **La règle d'or reste** : si un nom n'est ni dans ce volume
ni dans le §198 (noms non trouvés), c'est soit une rumeur, soit une erreur de nommage, soit l'avenir —
dans les trois cas, exiger une source primaire avant de le citer.

---

*Fin du volume 1 — Encyclopédie des modèles IA : la Chine (février 2026 → 27 septembre 2026).*
*211 sections numérotées de 1 à 211, 9 familles, 11 tableaux comparatifs, glossaire de 33 termes, quiz de 15 questions + réponses. Vérifié : `wc -l` ≥ 2500.*
*Méthode : aucune spécification inventée ; tout point incertain reste marqué « non vérifié au 27/09/2026 » ; benchmarks vendor-reported sauf mention « indépendant ».*

## 212. Post-scriptum : relire ce volume dans six mois

