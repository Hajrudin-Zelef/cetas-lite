---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-70
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Cerebras", "Google", "Groq", "Microsoft", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "United States", "Z.ai"]
dates: ["2026-07-27", "2026-08-02", "2026-09-24", "2026-09-27", "2028-08-02"]
keywords: ["agent", "agents", "benchmark", "claude", "dram", "embedding", "embeddings", "glm", "gpu", "kimi", "llama", "luna"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [5888, 6000]
sha256: 9237854df6539acb31e0dd8585c6c2a73e853275da2159b256c0c29ac78ce82d
---

# IA — Le grand dossier

1. **Prix des frontier** : la guerre OpenAI/Google/Anthropic fait baisser les prix d'entrée (Luna à 0,20 $/1M) — revoir le seuil local/cloud chaque trimestre.
2. **Modèles ouverts >70B utilisables** : GPT-OSS-120B, Nemotron OML, Kimi K3 — la qualité locale monte, le hardware doit suivre.
3. **Contexte 1M+ standard** : Kimi K2.6, GLM-5.3 — le RAG « jette tout le corpus dans le prompt » devient envisageable (mais cher).
4. **Puces d'inférence** : Groq LPU (licence NVIDIA), Cerebras WSE, Trainium3, TPU v7 — la bataille perf/watt rebat les cartes du cloud GPU.
5. **Crise DRAM** : +24 % sur les 4090 d'occasion depuis mars 2026 — acheter du GPU au bon moment compte autant que le choisir.
6. **Licences** : glissement de certains labs (GLM-5.3 abandonne MIT, MiniMax/Kimi à seuils) — **relire la licence à chaque montée de version**.
7. **Régulation UE** : AI Act, exigences de souveraineté dans les appels d'offres — avantage structurel au local et aux hébergeurs UE.
8. **Routage agentique** : les gateways deviennent des « routeurs d'agents » (MCP, A2A) — LiteLLM suit, OpenRouter aussi.
9. **Délistages** : modèles retirés sans préavis (Llama chez Groq, Sora 2 API fermée le 24/09/2026) — d'où l'abstraction par noms logiques.
10. **Embeddings FR** : la qualité des embeddings ouverts multilingues (Qwen3-Embedding, Nomic) — tester contre text-embedding-3-small sur ton corpus avant de basculer.

---

*Document rédigé le 27/09/2026. Les prix, versions et disponibilités évoluent vite : toute décision budgétaire ou d'achat doit être précédée d'une vérification sur les pages officielles.*

---

---

### Quiz — Partie 3 : Débats, dangers, Claude, stacks, communautés (10 questions)

## 8. Quiz — 10 questions + corrigé

> Réponds d'abord, puis vérifie. L'objectif : ancrer les chiffres et les concepts qui comptent.

**Q1.** Selon l'étude MIT Project NANDA (juillet 2025), quelle proportion des pilotes d'IA
générative en entreprise atteint un impact financier (P&L) mesurable ?
- A. ~50 %
- B. ~25 %
- C. ~5 %
- D. ~80 %

**Q2.** Dario Amodei (CEO d'Anthropic) a alerté en 2025-2026 sur la disparition potentielle de
50 % des emplois de bureau débutants en 1 à 5 ans. Qui, *au sein même d'Anthropic*, a
publiquement contredit cette prédiction avec des données en juillet 2026 ?
- A. Daniela Amodei
- B. Peter McCrory, économiste en chef
- C. Jack Clark
- D. Personne

**Q3.** Dans l'affaire Anthropic vs auteurs (accord approuvé le 20 juillet 2026), quelle
distinction juridique le juge Alsup a-t-il établie en juin 2025 ?
- A. Aucune : tout entraînement est illégal
- B. L'entraînement sur livres légalement acquis = fair use ; le stockage d'une bibliothèque de livres piratés = pas de fair use
- C. Seuls les livres américains sont concernés
- D. L'accord ne couvre que les 3 plaignants nommés

**Q4.** Au 27/09/2026, quelle obligation de l'AI Act est déjà applicable (depuis le 2 août 2026) ?
- A. L'évaluation de conformité des systèmes à haut risque (annexe III)
- B. La transparence : les chatbots doivent déclarer qu'ils sont une IA, les contenus synthétiques doivent être marqués
- C. L'interdiction totale des modèles open source
- D. Aucune, tout a été reporté

**Q5.** Quelle est la date reportée (par l'AI Omnibus de juillet 2026) des obligations « haut
risque » pour les systèmes autonomes (annexe III) ?
- A. 2 août 2026
- B. 2 décembre 2026
- C. 2 décembre 2027
- D. 2 août 2028

**Q6.** Qu'est-ce que le « lethal trifecta » de Simon Willison ?
- A. Un benchmark de jailbreaks
- B. Les 3 conditions qui rendent un agent dangereux : accès à des données de valeur + exposition à du contenu externe non fiable + capacité d'exfiltrer des données
- C. Les 3 lois de la robotique d'Asimov
- D. Un type de quantification en 3 bits

**Q7.** Qui a créé le Model Context Protocol (MCP), devenu le standard de connexion outils/données
pour les agents ?
- A. OpenAI
- B. Google
- C. Anthropic (fin 2024)
- D. Microsoft

**Q8.** Dans un RAG, pourquoi faut-il utiliser le *même* modèle d'embedding à l'indexation et à
la requête ?
- A. Pour des raisons de licence
- B. Parce que les vecteurs produits par deux modèles différents ne sont pas comparables (espaces différents)
- C. C'est faux, on peut mélanger librement
- D. Pour réduire la latence uniquement

**Q9.** Selon les données Gartner citées en juin 2026, quelle part de l'électricité mondiale des
data centers consommaient les serveurs optimisés IA en 2026 ?
- A. 5 %
- B. 15 %
- C. 31 %
- D. 60 %

**Q10.** En Chine, depuis le 1er septembre 2025, que doivent comporter les contenus générés par IA
diffusés publiquement ?
- A. Rien de particulier
- B. Uniquement un filigrane invisible
- C. Un double marquage : label visible + métadonnées implicites (filigrane machine-readable)
- D. Une autorisation préalable du CAC par contenu

---

#### Corrigé — Partie 3

| Q | Réponse | Explication (avec renvoi) |
|---|---|---|
| Q1 | **C. ~5 %** | MIT Project NANDA, juillet 2025 : ~95 % des pilotes sans impact P&L mesurable ; les auteurs précisent que les chiffres sont directionnels. L'échec vient rarement de la techno. Voir 1.2. |
| Q2 | **B. Peter McCrory** | L'économiste en chef d'Anthropic a publié (juillet 2026) une analyse : chômage US à 4,2 %, aucune divergence d'emploi entre métiers exposés à l'IA et les autres. Voir 1.3. |
| Q3 | **B** | Juge William Alsup (juin 2025) : entraînement sur livres légalement acquis = fair use ; bibliothèque centrale de 7+ M de livres piratés (LibGen, PiLiMi) = pas de fair use. Accord final : 1,5 Md$, ~500 000 œuvres, ~3 000 $/titre. Voir 1.4. |
| Q4 | **B** | Article 50 (transparence) applicable depuis le 02/08/2026, amendes possibles. Le haut risque est reporté (voir Q5). Voir 2.4.1. |
| Q5 | **C. 2 décembre 2027** | AI Omnibus (règlement UE 2026/1744, en vigueur 27/07/2026) : +16 mois pour l'annexe III ; 02/08/2028 pour l'annexe I (produits). Voir 2.4.1. |
| Q6 | **B** | Le triptyque à cartographier avant de brancher un agent sur des données d'entreprise. Voir 2.1.1. |
| Q7 | **C. Anthropic** | MCP créé par Anthropic fin 2024, adopté ensuite par OpenAI, Google, Microsoft. Voir 3.3 et 4.4. |
| Q8 | **B** | Chaque modèle d'embedding définit son propre espace vectoriel ; mélanger deux modèles casse la recherche de similarité. Anti-pattern RAG n°2. Voir 4.3. |


---

*Fin du grand dossier IA.*
