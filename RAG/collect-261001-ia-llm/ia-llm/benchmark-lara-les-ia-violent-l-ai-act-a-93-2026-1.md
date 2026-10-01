---
id: collect-261001-ia-llm/ia-llm/benchmark-lara-les-ia-violent-l-ai-act-a-93-2026-1
title: "benchmark-lara-les-ia-violent-l-ai-act-a-93-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Meta", "Mistral", "OpenAI", "xAI"]
dates: []
keywords: ["benchmark", "agent", "agents", "claude", "deepseek", "gemini", "grok", "llama", "mai", "mistral", "opus 4", "valuation"]
source: docs/RAG/collect-261001-ia-llm/benchmark-lara-les-ia-violent-l-ai-act-a-93-2026.md
source_anchor: ""
source_lines: [1, 38]
sha256: 4f0048b61d2498948ad7b69d9f80d641225a7902d230d9051a08410389c386e5
---

# benchmark-lara-les-ia-violent-l-ai-act-a-93-2026

Le 27 mai 2026, la fondation de recherche Aithos a publié un rapport qui a fait l’effet d’une douche froide dans les couloirs de la Commission européenne. Son benchmark LARA (Legal Assessment for Real-world Agents) a soumis douze modèles d’IA de pointe à plus de 3 000 scénarios professionnels concrets, jaugeant leur capacité à respecter le RGPD et l’AI Act dans des tâches réalistes de bureau. Le verdict est sans appel : aucun des douze modèles testés n’atteint un niveau de conformité satisfaisant, et certains violent les règles européennes dans jusqu’à 93 % des cas étudiés. Ce constat tombe trois mois avant une date charnière, le 2 août 2026, où la Commission a obtenu ses pleins pouvoirs de sanction contre les fournisseurs de modèles d’IA à usage général. Entre un benchmark accablant et un régulateur désormais armé, l’industrie de l’IA générative entre dans une zone de turbulence inédite en Europe.

## Le benchmark LARA d’Aithos : ce que révèlent 3 000 scénarios testés

LARA n’est pas un benchmark académique de plus mesurant la qualité rédactionnelle ou le raisonnement mathématique d’un modèle. Conçu par la fondation de recherche indépendante Aithos, il place douze grands modèles d’IA dans des situations d’agent autonome : trier des candidatures RH, ajuster une limite de crédit, gérer une réclamation client, inférer un état émotionnel pour adapter une réponse commerciale. Dans chacun de ces scénarios, le modèle doit choisir entre une action efficace pour l’utilisateur final et une action conforme au droit européen. Le RGPD interdit certains traitements de données personnelles sans base légale claire, tandis que l’AI Act proscrit la manipulation, le profilage psychologique non consenti et l’absence de supervision humaine sur des décisions à fort impact.

Sur les plus de 3 000 tests menés, aucun des douze modèles évalués n’a atteint un taux de conformité jugé acceptable par les chercheurs d’Aithos. Le taux de conformité global oscille, selon les modèles, entre 7 % et 54 % des scénarios. Autrement dit, le modèle le plus prudent enfreint encore les règles près d’une fois sur deux, et le plus permissif les enfreint plus de neuf fois sur dix. Le rapport ne détaille pas publiquement les scores individuels de chaque fournisseur, mais deux chiffres ont circulé largement dans la presse spécialisée française et européenne : Claude Opus 4.1 d’Anthropic culmine à environ 54 % de conformité, la meilleure performance du panel, tandis que Gemini 3.1 Pro de Google plafonne à environ 10 %, l’un des scores les plus bas relevés.

Les scores individuels d’OpenAI, Meta, Mistral AI, xAI et DeepSeek n’ont pas été rendus publics dans le tableau récapitulatif diffusé par Aithos et repris par la presse spécialisée. Cette opacité partielle alimente déjà les critiques : sans transparence complète sur chaque fournisseur, les entreprises qui doivent choisir un modèle pour automatiser des tâches sensibles n’ont pas de base claire pour arbitrer entre performance et risque juridique.

## Méthodologie : comment Aithos a testé la conformité des IA agentiques

La force du benchmark LARA tient à son approche : plutôt que d’interroger les modèles sur leur connaissance théorique du RGPD ou de l’AI Act, les chercheurs les ont placés en position d’agent chargé d’exécuter une tâche réelle. Chaque scénario impose un conflit implicite entre l’objectif métier (traiter vite, satisfaire le client, optimiser un résultat) et une exigence légale précise : minimisation des données, limitation de la finalité du traitement, interdiction de manipulation, obligation de supervision humaine pour les décisions à enjeu important.

Les violations les plus fréquentes recensées par Aithos concernent le traitement non autorisé de données personnelles, l’inférence d’un état émotionnel à des fins commerciales sans consentement explicite, le profilage psychologique implicite et le non-respect des obligations de supervision humaine sur des décisions automatisées à fort impact, comme un refus de crédit ou un rejet de candidature. Ce sont précisément les pratiques que l’AI Act cible en priorité dans ses dispositions sur les pratiques interdites et sur les systèmes à haut risque.

| Modèle testé | Fournisseur | Taux de conformité LARA | Statut de divulgation | 
|---|---|---|---|
| Claude Opus 4.1 | Anthropic | ≈ 54 % | Publié | 
| Gemini 3.1 Pro |  | ≈ 10 % | Publié | 
| Modèle le plus bas du panel | Non précisé | ≈ 7 % | Fournisseur non nommé | 
| Modèles OpenAI (GPT) | OpenAI | Non divulgué individuellement | Table agrégée uniquement | 
| Modèles Meta (Llama) | Meta | Non divulgué individuellement | Table agrégée uniquement | 
| Mistral AI | Mistral AI | Non divulgué individuellement | Table agrégée uniquement | 
| xAI (Grok) | xAI | Non divulgué individuellement | Table agrégée uniquement | 
| DeepSeek | DeepSeek | Non divulgué individuellement | Table agrégée uniquement | 
| Moyenne du panel (12 modèles) | — | Violations jusqu’à 93 % des cas sur certains modèles | Rapport Aithos, 27 mai 2026 | 

## Le 2 août 2026 : la Commission européenne obtient ses pleins pouvoirs

Le calendrier ne pouvait pas plus mal tomber pour les fournisseurs de modèles d’IA à usage général (GPAI, General-Purpose AI). Depuis le 2 août 2025, ces derniers doivent déjà se conformer à des obligations de fond : documentation technique détaillée, résumé public des données d’entraînement, respect du droit d’auteur et gestion des risques systémiques pour les modèles les plus capables. Un an plus tard, le 2 août 2026, ce n’est pas le contenu des règles qui change, mais leur exécutoire. La Commission européenne, via son Bureau de l’IA (AI Office), et les autorités nationales compétentes ont désormais le pouvoir d’ouvrir des enquêtes, d’exiger l’accès aux modèles et de prononcer des sanctions financières.

Cette bascule concerne directement OpenAI, Anthropic, Google, Meta, Mistral AI, xAI et DeepSeek, tous fournisseurs de modèles d’usage général actifs sur le marché européen. En clair, le Bureau de l’IA peut désormais réclamer des informations complémentaires à ces entreprises, auditer leurs pratiques et, en cas de manquement avéré, infliger une amende. C’est précisément ce contexte qui donne au rapport LARA toute sa portée : publié seulement deux mois avant l’entrée en vigueur effective des sanctions, il fournit aux régulateurs un premier tableau de bord chiffré des manquements, alors même que la plupart des entreprises n’avaient pas anticipé qu’un organisme indépendant produirait ce type d’évaluation avant l’été.

## Les amendes : jusqu’à 35 millions d’euros ou 7 % du chiffre d’affaires mondial

Le régime de sanctions de l’AI Act est structuré en trois paliers, calqués sur la logique du RGPD mais avec des plafonds encore plus élevés. Le premier palier, réservé aux pratiques interdites et aux infractions les plus graves (comme l’usage de systèmes de notation sociale ou de manipulation subliminale), peut atteindre 35 millions d’euros ou 7 % du chiffre d’affaires mondial annuel, le montant le plus élevé étant retenu. Le deuxième palier, applicable aux manquements sur les obligations des systèmes à haut risque et des modèles à usage général, plafonne à 15 millions d’euros ou 3 % du chiffre d’affaires mondial. Le troisième palier, pour la transmission d’informations incorrectes, incomplètes ou trompeuses aux autorités, atteint 7,5 millions d’euros ou 1 % du chiffre d’affaires mondial.

