---
id: collect-261001-ia-llm/ia-llm/opus-5-vs-gemini-3-5-pro-vs-deepseek-x4-5-d-ecart-2
title: "opus-5-vs-gemini-3-5-pro-vs-deepseek-x4-5-d-ecart"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Hugging Face"]
dates: []
keywords: ["deepseek", "gemini", "agents", "benchmarks", "claude", "opus 4", "opus 5", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/opus-5-vs-gemini-3-5-pro-vs-deepseek-x4-5-d-ecart.md
source_anchor: ""
source_lines: [31, 82]
sha256: 00d8093f83aa6dc28df64a6051dbf0a5fffcfd54c0bebed76fb4cf562c20e2fe
---

# opus-5-vs-gemini-3-5-pro-vs-deepseek-x4-5-d-ecart

DeepSeek a publié la variante V4-Pro-0813 le 13 août 2026, un correctif de mi-parcours de son modèle V4-Pro lancé plus tôt dans l’année. C’est le seul des trois modèles comparés ici à afficher des scores de benchmarks détaillés et publics. Le classement ModelGrep du 24 août 2026 lui attribue un score Intelligence de 53,2, un score de code de 68,8, un score agentique de 37,8, une vitesse de 92 tokens par seconde et un coût d’entrée de 1,12 dollar par million de tokens.

Le comparatif RohitAI, publié le 14 août 2026, donne des chiffres légèrement différents pour le même modèle : un score Intelligence de 53,20 (cohérent avec ModelGrep), mais un score agentique de 49,56 et une réussite Terminal-Bench 2.1 de 78,65 %, avec une vitesse de 78,44 tokens par seconde en configuration maximale. L’écart entre les deux scores agentiques, 37,8 contre 49,56, illustre bien à quel point les méthodologies de banc d’essai varient d’un évaluateur à l’autre. Pour rester prudent, il vaut mieux retenir la fourchette complète plutôt qu’un seul chiffre.

DeepSeek revendique par ailleurs, dans ses propres communications reprises par l’encyclopédie LLM de Stochastic Sandbox début août 2026, des résultats internes d’environ 95 % sur AIME 2025 et 78 % sur SWE-Bench Verified pour la lignée V4-Pro, avec une fenêtre de contexte confirmée d’un million de tokens. Ces chiffres proviennent de communications du fournisseur et méritent d’être traités comme tels, en attendant une validation par un banc d’essai indépendant.

Une variante ouverte, DeepSeek-V4-Pro-Max, publiée en avril 2026, reste disponible sous licence ouverte avec un contexte d’un million de tokens, un débit annoncé de 1600 tokens par seconde et un tarif d’environ 1,60 dollar en entrée et 3,20 dollars en sortie par million de tokens, selon LLM-Stats. C’est une option distincte de V4-Pro-0813, pensée pour l’auto-hébergement plutôt que pour l’accès API géré.

## Tableau comparatif complet des caractéristiques techniques

Voici la synthèse des caractéristiques publiques des trois modèles, avec leurs sources respectives. Les cases marquées « non communiqué » reflètent une absence réelle de donnée officielle plutôt qu’une estimation.

| Caractéristique | Claude Opus 5 | Gemini 3.5 Pro | DeepSeek V4-Pro-0813 | 
|---|---|---|---|
| Développeur | Anthropic | Google DeepMind | DeepSeek | 
| Date de sortie | Juillet 2026 | Juillet 2026 | 13 août 2026 | 
| Fenêtre de contexte | 1,0 million de tokens | 2,1 millions de tokens | 1,0 million de tokens | 
| Type de poids | Fermé (API) | Fermé (API) | Fermé (variante 0813) / ouvert pour V4-Pro-Max | 
| Prix entrée /1M tokens | ~5,00 $ | Non communiqué | 1,12 $ | 
| Prix sortie /1M tokens | ~25,00 $ | Non communiqué | Non détaillé séparément | 
| Vitesse de génération | 52,3 tokens/s | Non communiqué | 78,4 à 92 tokens/s selon la source | 
| Positionnement annoncé | Code, agents, raisonnement | Contexte long, raisonnement, vision | Code, raisonnement, rapport coût-performance | 
| Score Intelligence publié | Non communiqué | Non communiqué | 53,2 | 
| Accès | API Anthropic, Claude Cloud | Google AI Studio, Vertex AI | API DeepSeek, Hugging Face pour les variantes ouvertes | 
| Modèle précédent | Claude Opus 4.8 | Gemini 3.1 Pro | DeepSeek V4-Pro (version initiale, avant 0813) | 

## Benchmarks : intelligence, code et capacités agentiques

Le tableau ci-dessous compile les scores disponibles pour chaque modèle, en citant systématiquement la source et la date. Quand deux sources se contredisent, les deux chiffres sont indiqués plutôt qu’une moyenne arbitraire, conformément à une approche prudente sur des données encore mouvantes.

| Modèle | Score Intelligence | Score agentique | Terminal-Bench 2.1 | Vitesse | Source | 
|---|---|---|---|---|---|
| Claude Opus 5 | Non communiqué | Non communiqué | Non communiqué | 52,3 tok/s | AY Automate, août 2026 | 
| Gemini 3.5 Pro | Non communiqué | Non communiqué | Non communiqué | Non communiqué | AY Automate, août 2026 | 
| Gemini 3.7 Flash (variante rapide, référence) | 56,03 | 45,10 | 85,77 % | ~340 tok/s | RohitAI, 14 août 2026 | 
| DeepSeek V4-Pro-0813 | 53,2 | 37,8 (ModelGrep) / 49,56 (RohitAI) | 78,65 % (RohitAI) | 92 tok/s (ModelGrep) / 78,4 tok/s (RohitAI) | ModelGrep 24 août + RohitAI 14 août 2026 | 
| DeepSeek V4-Flash 0731 (référence) | 50 (Artificial Analysis Index) | Non communiqué | Non communiqué | Non communiqué | Stochastic Sandbox, 1er août 2026 | 

Trois enseignements ressortent de ce tableau. D’abord, DeepSeek reste le seul fournisseur des trois à publier des scores comparables d’une source à l’autre, ce qui facilite l’audit technique côté client mais complique aussi la vérification indépendante de ses propres claims marketing. Ensuite, l’écart entre les deux scores agentiques de DeepSeek montre qu’un même modèle peut afficher un score presque deux fois supérieur selon la méthode de test, un rappel utile avant de citer un seul chiffre dans une présentation interne. Enfin, l’absence de données chiffrées pour Claude Opus 5 et Gemini 3.5 Pro sur les bancs d’essai tiers consultés ne signifie pas que ces modèles sont moins performants. Cela reflète plutôt une stratégie de communication différente, où Anthropic et Google misent sur leur réputation de marque plutôt que sur des tableaux de scores publics.

## Vitesse de génération et latence en conditions réelles

La vitesse brute, mesurée en tokens par seconde, ne raconte qu’une partie de l’histoire. Un modèle rapide sur un banc d’essai contrôlé peut ralentir fortement sous forte charge, surtout en heure de pointe européenne quand les infrastructures américaines et asiatiques absorbent aussi le trafic domestique.

Sur le papier, DeepSeek V4-Pro-0813 domine largement Claude Opus 5 en débit brut, avec 78 à 92 tokens par seconde contre 52,3 pour Opus 5. Pour un usage de chat interactif ou de complétion de code en temps réel, cet écart se traduit par une sensation de réponse plus instantanée côté DeepSeek. Gemini 3.5 Pro n’a pas de chiffre de vitesse publié dans les sources consultées, ce qui empêche toute comparaison directe sur ce point précis. Sa variante Flash, en revanche, affiche un débit nettement supérieur, autour de 340 tokens par seconde, ce qui suggère que Google a choisi de sacrifier une part de vitesse sur son modèle Pro au profit de la profondeur de raisonnement et de la gestion du contexte long.

Pour les équipes qui développent des agents autonomes exécutant plusieurs appels d’outils à la chaîne, la latence cumulée compte davantage que le débit token par token. Un modèle qui répond vite mais se trompe souvent sur l’appel d’outil oblige à relancer la chaîne, ce qui annule tout le gain de vitesse initial. C’est un argument qui pèse en faveur de Claude Opus 5 malgré son débit plus modeste, du moins pour les scénarios agentiques complexes où la fiabilité de la première tentative compte plus que la rapidité pure.

## Tarification : combien coûte réellement chaque modèle

Le budget reste souvent le critère décisif pour les équipes qui traitent de gros volumes de requêtes. Voici la comparaison des tarifs API publics, avec les incertitudes clairement signalées là où elles existent.

