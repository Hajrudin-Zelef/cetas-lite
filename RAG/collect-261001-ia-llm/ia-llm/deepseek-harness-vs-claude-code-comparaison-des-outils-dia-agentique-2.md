---
id: collect-261001-ia-llm/ia-llm/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique-2
title: "deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "OpenAI"]
dates: []
keywords: ["claude", "deepseek", "agents", "benchmark", "benchmarks", "sonnet 5"]
source: docs/RAG/collect-261001-ia-llm/deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique.md
source_anchor: ""
source_lines: [99, 198]
sha256: 922dd768260d21676a3aa4590d4ae9eb917dd525249fb5ae190cef54b45bb6ca
---

# deepseek-harness-vs-claude-code-comparaison-des-outils-dia-agentique

- Les plugins Claude Code agrègent des fonctionnalités autour de la boucle.
- Les plugins DeepSeek Harness peuvent définir des pans du runtime lui-même.

### Prise en charge des modèles et choix du fournisseur

DeepSeek Harness prend en charge davantage de fournisseurs de modèles. Il exécute DeepSeek, Anthropic, OpenAI, des clouds et des endpoints locaux compatibles. Claude Code exécute… Claude. En d’autres termes, DeepSeek Harness peut héberger Claude, mais Claude Code ne peut pas héberger DeepSeek.

C’est ce qui m’a permis de garder le même modèle lors des tests. Comparer un modèle DeepSeek dans Harness à Claude dans Claude Code change deux variables à la fois.

Le choix du fournisseur implique la configuration des identifiants et de l’endpoint, y compris les IDs de modèle exacts. Claude Code contrôle la famille de modèles et la plupart des paramètres de requête.

Changer de fournisseur ne garantit pas un comportement identique. Les modèles peuvent différer par le format d’outillage, la taille de contexte ou les contrôles de raisonnement. Harness garde la configuration des plugins accessible, mais l’adaptateur fournisseur doit rester compatible avec l’endpoint choisi.

### Journaux de session et traçabilité

DeepSeek Harness enregistre invites, injections de contexte, appels d’outils et décisions de permission dans un flux d’événements append-only. Sa vue « trajectory » montre l’origine de chaque enregistrement et prend en charge la reprise, le fork, la recherche et la relecture.

Claude Code stocke des transcriptions JSONL, prend en charge reprise et fork, et émet des traces OpenTelemetry.

Les deux exposent l’historique de session, mais DeepSeek Harness met davantage le chemin d’exécution en avant dans l’UI.

Vue « trajectory » reconstruisant une exécution complète. Vidéo de l’auteur.

### Environnements d’exécution et permissions

DeepSeek Harness utilise bubblewrap ou Landlock sous Linux, Seatbelt sous macOS, et un jeton restreint par ACL sous Windows. Si le bac à sable ne peut pas démarrer, Harness bloque la commande. Claude Code propose des modes de permission ainsi que des règles allow, ask et deny. Dans ce run, le bac à sable Windows a influencé le chemin de vérification et le timing.

Les libellés d’accès diffèrent, rendant impossible un alignement strict des réglages.

- 
Les préréglages de permissions de DeepSeek Harness incluent `read-only` ,`workspace-write` et`danger-full-access` .
- 
Claude Code sépare les modifications automatiques, l’approbation manuelle, la planification et l’accès aux commandes par règles. Il prend aussi des snapshots des fichiers avant modification pour un rollback sans Git.

Sur les sessions Pro, Max et Team actuelles dans le terminal et VS Code, le mode Auto est le mode de démarrage intégré de Claude Code. Un classifieur révise les actions en arrière-plan. Le test a explicitement utilisé `acceptEdits` ; son comportement en matière de permissions reflète donc la configuration de test plutôt que le défaut actuel.

Cette nuance compte lorsqu’on compare le nombre d’approbations.

Le lieu d’exécution diffère aussi. Harness s’exécute principalement sur la machine qui héberge son interface web ou un processus headless. Claude Code peut tourner en local, dans des environnements cloud managés par Anthropic, ou sur une infrastructure autohébergée. Remote Control ajoute une interface navigateur tandis que l’exécution reste locale.

## Test pratique DeepSeek Harness vs Claude Code

Pour vérifier si ces différences de runtime affectaient l’exécution, j’ai confié aux deux outils le même bug TypeScript d’une ligne. Si vous voulez uniquement le résultat chronométré, c’est la section à lire.

### Configuration du test : même modèle, même dépôt, même prompt

J’ai écrit une bibliothèque TypeScript de « habit-streak » avec une ligne fautive. Elle mesurait les jours calendaires en arrondissant les millisecondes écoulées.

- Une validation à 23:50 suivie d’une autre à 00:10 semblait appartenir au même jour.
- Une validation à 08:00 suivie de 20:00 le lendemain semblait indiquer un jour manqué.

Dix tests figés révélaient trois échecs.

Chaque nouveau clone avait ses dépendances installées avant le chronométrage. Les deux agents ont reçu la même instruction et utilisé Claude Sonnet 5. Le dépôt et le prompt étaient fixes ; les outils, permissions et comportements de bac à sable restaient spécifiques à chaque produit.

Les configurations de permissions n’étaient pas équivalentes. DeepSeek Harness a démarré en `workspace-write`, mais son bac à sable Windows n’a pas pu se lancer. L’interface web a demandé une approbation, et j’ai accordé `danger-full-access` à trois reprises. Claude Code utilisait `accept-edits` et ne pouvait exécuter que la commande de test via bash.

Cette tâche ciblée n’utilisait que du code et des tests locaux. Les agents devaient trouver et corriger une ligne défectueuse, puis exécuter la suite. Les dépendances étaient préinstallées avant le chrono, bien que Harness ait choisi de relancer `npm install`.

### Résultats : patch identique, chemins d’approbation différents

Claude Code a indiqué 55,0 secondes. Il a exécuté la suite, trouvé le bug, modifié un fichier source et relancé les tests. Dix sur dix ont réussi.

Dans ce test Windows, DeepSeek Harness a pris 125,4 secondes. Il a trouvé le même helper `daysBetween` inutilisé, appliqué la même correction et exécuté avec succès la suite d’origine.

Les diffs étaient identiques à l’octet :

```
-import { DAY_MS } from './dates.js';
+import { daysBetween } from './dates.js';
 
 function gapInDays(earlier: number, later: number): number {
-  return Math.round((later - earlier) / DAY_MS);
+  return daysBetween(earlier, later);
 }
```
Le parcours d’approbation et les répétitions côté shell expliquent l’essentiel de l’écart de temps.

### Pourquoi DeepSeek Harness a été plus lent sous Windows

Le bac à sable `workspace-write` de DeepSeek Harness n’a pas pu démarrer sous Windows, si bien que l’interface web a demandé à trois reprises un mode de permission plus large. Après approbation, Harness a exécuté la suite d’origine, modifié `src/streak.ts`, puis relancé la suite. Dix tests sur dix ont réussi.

Ce résultat ne dit rien du comportement de Harness sous Linux ou macOS. Comme les deux exécutions utilisaient le même ID de modèle, l’écart de temps ne peut pas venir d’une différence de famille de modèles. Les permissions et chemins d’outils distincts ont toutefois façonné chaque run.

Les deux exécutions atteignent dix tests réussis. Image de l’auteur.

### Ce que le test avec le même modèle isole

Partager le même ID de modèle n’a pas rendu les requêtes identiques. Chaque produit a fourni son propre prompt système, ses descriptions d’outils, son contexte et ses règles de permission. Chacun de ces apports peut influencer la réponse suivante du modèle.

La configuration contrôle donc mieux la famille de modèles qu’une comparaison « modèle DeepSeek vs Claude », sans pour autant isoler toutes les variables.

Je n’ai réalisé qu’un essai chronométré par configuration. Des répétitions, dans un ordre aléatoire, seraient nécessaires avant de considérer les temps ou les comptes d’outils comme des mesures de performance stables.

## DeepSeek Harness vs Claude Code : coûts, benchmarks et limites des tests

Une étude de cas sous Windows ne peut pas trancher la comparaison générale de performance ou de prix. Elle montre cependant pourquoi les détails de benchmark comptent.

### Avertissement sur les benchmarks : DeepSeek a utilisé le mode Minimal

