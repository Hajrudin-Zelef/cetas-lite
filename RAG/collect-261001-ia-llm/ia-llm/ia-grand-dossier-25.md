---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-25
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Meta", "OpenAI", "vLLM"]
dates: []
keywords: ["agents", "agi", "arr", "capex", "compute", "datacenter", "gpu", "hyperscaler", "training", "valuation", "vllm"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1673, 1726]
sha256: d9d1e10119886f8126a0158e3451eaa247a497cb2f9e157b23dbf5c02f278e9b
---

# IA — Le grand dossier

**6.10.2. Pourquoi la charge GPU est un cauchemar pour l'électricien (et comment le dompter).**
1. **Appels de puissance brutaux** : un run d'entraînement fait passer les GPU de 10 % à 100 % en quelques secondes, de façon synchronisée sur des milliers de cartes — des **marches de plusieurs MW** sur le réseau du datacenter. L'inférence est plus douce (charge variable mais continue), d'où un argument de plus pour l'ère de l'inférence.
2. **Harmoniques** : les alimentations à découpage des serveurs GPU polluent le réseau (THDi élevé) → **transformateurs surdimensionnés**, filtres actifs dans les gros déploiements.
3. **Facteur de puissance et réactif** : à surveiller contractuellement (pénalités du fournisseur d'énergie au-delà d'un seuil de tan φ).
4. **Secours** : un onduleur dimensionné pour 15 minutes d'autonomie à 120 kW/rack = une **salle batteries** à part entière. Règle : l'UPS d'un cluster IA se dimensionne comme un **process industriel**, avec groupe électrogène à démarrage < 10 s et **bypass** maintenable (voir le guide onduleurs de Zelef : topologies VFI, N+1, test d'autonomie réel).
5. **Refroidissement = la moitié de la facture** : un PUE de 1,5 sur 1 MW IT = 500 kW de froid en plus. Le **free cooling** et le liquide (eau tiède 30-40 °C réutilisable pour chauffer des bureaux) sont les leviers — le « datacenter qui chauffe le quartier » n'est plus une blague mais un business model (plusieurs projets européens 2024-2026).

**6.10.3. Ce que ça change pour l'entreprise de Zelef (pas un hyperscaler).**
- **Ne pas mettre de GPU lourds dans la salle serveur existante** sans étude : bilan de puissance, bilan thermique, tenue des disjoncteurs et des câbles, plan de délestage.
- **Préférer l'inférence légère en local** (1-2 GPU prosumer, 1-3 kW) : ça rentre dans une baie standard, un onduleur 3-6 kVA existant l'encaisse, et ça couvre 80 % des besoins RAG.
- **Le cloud pour les pics** : l'élasticité du cloud, c'est aussi de l'**élasticité électrique** — on ne paie le MW que quand on s'en sert.
- **Supervision unifiée** : corréler les métriques IT (tokens/s, file d'attente vLLM) et les métriques énergie (kW, PUE, température) dans le même Grafana — c'est le FinOps IA (6.3) étendu au watt : **le coût du token, c'est aussi des kWh**.
- **Ordre de grandeur à retenir** : servir 1M tokens sur GPU local ≈ **quelques kWh** (selon modèle et quantification) — négligeable à l'échelle d'une équipe, significatif à l'échelle d'un datacenter.

**6.10.4. Checklist « accueil d'un serveur GPU » pour un responsable énergies.**
- [ ] Bilan de puissance disponible (transformateur, TGBT, départs) avec marge 20 %
- [ ] Bilan thermique de la salle (kW froid disponibles, redondance N+1)
- [ ] Protections : calibre disjoncteurs, sélectivité, parafoudre
- [ ] Onduleur : puissance, autonomie **mesurée** (pas théorique), bypass, plan de test annuel
- [ ] Groupe électrogène : puissance de reprise, contrat de maintenance, test en charge mensuel
- [ ] Qualité réseau : mesure THDi, facteur de puissance, corrective si besoin
- [ ] Supervision : puissance par départ, températures, alertes — intégrée à la supervision IT
- [ ] PRA : qu'est-ce qui s'arrête proprement en cas de coupure longue ? (scénario NUT du guide Proxmox)

*Cette section fait le pont entre le dossier IA et le guide onduleurs/UPS de Zelef : le même métier, la même rigueur, une nouvelle charge.*

---

## 4.9. Scénarios 2027-2030 : trois futurs possibles (opinions attribuées, pas des faits)

> Personne ne connaît l'avenir. Voici trois scénarios cohérents, chacun défendu par des voix réelles — pour décider en incertitude plutôt que parier sur une date.

**Scénario A — « Le plateau productif » (défendu par : les pragmatiques industrie, une partie des analystes 2026).**
Le training continue de ralentir, l'inférence et l'efficacité portent le progrès à un rythme **linéaire** : chaque année, -50 % de coût par unité d'intelligence, +30 % de capacités agents. Pas d'AGI avant 2035+. L'IA devient une **commodité industrielle** comme l'électricité : la valeur est dans l'intégration, pas dans le modèle. **Implication sysadmin :** investir dans le serving, le FinOps et les données — pas dans la course au modèle.

**Scénario B — « La percée par le raisonnement » (défendu par : Amodei, l'équipe o-series d'OpenAI, les tenants du RLVR).**
Le RL avec récompenses vérifiables (maths formelles, code, science simulée) produit des **sauts discontinus** : un « chercheur IA » crédible d'ici 2027-2028 accélère la R&D elle-même (boucle récursive). Les données humaines ne sont plus le plafond. **Implication sysadmin :** se préparer à des agents **beaucoup plus autonomes et consommateurs** (budgets, supervision, sécurité des agents deviennent critiques).

**Scénario C — « Le choc d'offre » (défendu par : les sceptiques du scaling, LeCun sur le fond).**
Ni le training ni l'inférence ne suffisent : les LLM plafonnent sur le raisonnement robuste, la planification long terme et le monde physique. Il faut une **rupture architecturale** (world models, JEPA, neuro-symbolique) qui n'arrive pas avant 5-10 ans. L'investissement se dégonfle partiellement (« hiver » sectoriel), les prix s'effondrent, seuls les cas d'usage rentables survivent. **Implication sysadmin :** ne pas sur-investir en CAPEX IA ; privilégier l'**OPEX réversible** (API, cloud) et les compétences transférables (data, éval, MLOps).

**Comment décider sans savoir : la stratégie « no-regret » (recommandation de l'auteur).**
1. **Abstraction** : interfaces standard (API OpenAI-compatible) → on change de modèle/fournisseur en un jour, quel que soit le scénario.
2. **Données propres** : utiles dans les trois scénarios (le corpus métier reste l'actif).
3. **OPEX > CAPEX** : louer le compute, pas l'acheter — sauf volume prouvé (14.1).
4. **Compétences transférables** : évaluation, data engineering, sécurité, énergétique — valables même si les modèles changent.
5. **Veille trimestrielle** : rejouer le gold set (14.3) et relire cette section — le scénario qui se réalise se voit dans les chiffres avant les discours.

---

## 4.10. Dix questions que le débat n'a pas tranchées (l'agenda 2026-2030)

> Ce que même les experts des deux camps admettent ne pas savoir. Une bonne question ouverte vaut mieux qu'une mauvaise certitude.

