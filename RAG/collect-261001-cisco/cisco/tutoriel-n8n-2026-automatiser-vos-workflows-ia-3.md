---
id: collect-261001-cisco/cisco/tutoriel-n8n-2026-automatiser-vos-workflows-ia-3
title: "Vérifier le statut du conteneur"
domain: cisco
role: reference
task: reference
actors: ["Anthropic", "Google", "OpenAI"]
dates: []
keywords: ["agent", "agents", "claude", "gemini", "mcp", "model context protocol", "open source", "opus 4"]
source: docs/RAG/collect-261001-cisco/tutoriel-n8n-2026-automatiser-vos-workflows-ia.md
source_anchor: ""
source_lines: [230, 333]
sha256: 562b1bd3abc155d6d3dd4b848718b2c6b72ee0b384ca06d5599dd56757d39d16
---

# Vérifier le statut du conteneur

La fonctionnalité phare de n8n 2.0 est sans conteste l'intégration native des agents IA. Avec le nœud « AI Agent », vous pouvez créer des workflows autonomes capables de raisonner, d'utiliser des outils et de prendre des décisions complexes — le tout sans écrire une seule ligne de code d'orchestration. Cette capacité positionne n8n comme une alternative sérieuse à LangChain et CrewAI pour les workflows d'automatisation IA en entreprise.

n8n supporte nativement les modèles de langage d'OpenAI (GPT-4o, GPT-4.5), d'Anthropic (Claude Opus 4, Claude Sonnet 4), de Google (Gemini 2.5 Pro) et les modèles open source via Ollama. Si vous souhaitez utiliser des modèles locaux, consultez notre tutoriel Ollama pour l'installation et la configuration.

Créons un agent IA pratique : un assistant qui analyse les tickets de support entrants, les catégorise automatiquement, suggère une réponse et les assigne au bon membre de l'équipe. Ce workflow combine un webhook (réception du ticket), un agent IA (analyse et classification) et des actions concrètes (mise à jour du ticket, notification Slack).

Commencez par ajouter un nœud « AI Agent » à votre workflow. Configurez-le avec le modèle de votre choix (Claude Sonnet 4 offre un excellent rapport qualité-prix pour ce cas d'usage). Le prompt système définit le comportement de l'agent :

```
// Configuration du nœud AI Agent dans n8n
// Prompt système de l'agent de tri de tickets
Tu es un agent de support technique senior spécialisé dans le tri des tickets.
Pour chaque ticket reçu, tu dois :
1. ANALYSER le contenu et identifier le problème principal
2. CATÉGORISER le ticket parmi : bug_critique, bug_mineur, 
   demande_fonctionnalite, question_technique, facturation
3. ÉVALUER la priorité : P1 (urgent), P2 (important), P3 (normal), P4 (faible)
4. IDENTIFIER l'équipe responsable : backend, frontend, infrastructure, 
   data, support_client
5. RÉDIGER une réponse préliminaire au client en français
Format de sortie JSON :
{
  "categorie": "...",
  "priorite": "P1|P2|P3|P4",
  "equipe": "...",
  "resume": "Résumé en une phrase",
  "reponse_client": "Réponse préliminaire...",
  "tags": ["tag1", "tag2"]
}
```
Connectez la sortie de l'agent IA à un nœud « Code » qui parse la réponse JSON, puis à un nœud « Switch » qui route vers les actions appropriées selon la priorité. Les tickets P1 déclenchent une notification Slack immédiate dans le canal `#urgences`, tandis que les tickets P3-P4 sont simplement mis à jour dans le système de ticketing.

L'un des atouts majeurs de n8n pour les workflows IA est la possibilité de donner des « outils » à l'agent. Par exemple, vous pouvez connecter un nœud de recherche base de données pour que l'agent consulte l'historique des tickets du client, ou un nœud HTTP pour vérifier le statut d'un service en temps réel. L'agent décide de manière autonome quand utiliser ces outils, exactement comme le ferait un agent LangChain — mais avec une interface visuelle et sans code Python.

n8n supporte également le protocole MCP (Model Context Protocol), ce qui permet à vos workflows d'être appelés par d'autres plateformes IA comme Claude ou des applications construites avec le SDK d'agents. Cette interopérabilité est un avantage compétitif majeur en 2026, où l'écosystème des agents IA se structure autour de standards ouverts. Pour approfondir le sujet des agents autonomes en entreprise, consultez notre analyse sur les agents IA qui transforment l'entreprise.

## Étape 7 : Manipuler les Données avec les Nœuds de Transformation

Dans la plupart des workflows réels, les données passent par plusieurs étapes de transformation entre la source et la destination. n8n offre un ensemble complet de nœuds de manipulation de données qui couvrent 95 % des cas d'usage sans écrire de code personnalisé. Ces nœuds sont essentiels pour construire des pipelines de données robustes.

Les nœuds de transformation les plus utilisés sont :

**Set** : Redéfinit la structure des données en sortie. Idéal pour ne garder que les champs nécessaires ou renommer des propriétés. Par exemple, transformer `{ "user_name": "Jean" }` en `{ "nom": "Jean" }`.

**Merge** : Combine les données de deux branches du workflow. Supporte les modes « Append » (concaténation), « Keep Key Matches » (jointure) et « Merge By Index ». C'est l'équivalent d'un JOIN SQL dans le monde visuel de n8n.

**Split In Batches** : Découpe un grand ensemble de données en lots plus petits pour respecter les rate limits des API. Indispensable quand vous devez traiter des centaines d'éléments via une API qui limite à 100 requêtes par minute.

**Aggregate** : Regroupe plusieurs éléments en un seul. Parfait pour créer un rapport récapitulatif à partir de données individuelles.

**Code** : Pour les transformations complexes, le nœud Code accepte du JavaScript ou du Python natif. Depuis n8n 2.0, le support Python est intégré nativement, ce qui ravira les data engineers habitués à pandas et numpy.

Voici un exemple concret de pipeline de transformation de données. Ce workflow récupère des données de ventes depuis une API, les transforme et les insère dans une base PostgreSQL :

```
// Nœud Code : Transformer les données de ventes
// Entrée : données brutes de l'API de ventes
// Sortie : données nettoyées et enrichies pour PostgreSQL
const items = $input.all();
const results = [];
for (const item of items) {
  const sale = item.json;
  
  // Convertir le montant en centimes (éviter les erreurs de virgule flottante)
  const montantCentimes = Math.round(parseFloat(sale.amount) * 100);
  
  // Normaliser la date au format ISO
  const dateVente = new Date(sale.date).toISOString();
  
  // Calculer la TVA française (20%)
  const montantHT = Math.round(montantCentimes / 1.20);
  const tva = montantCentimes - montantHT;
  
  // Catégoriser le montant
  let segment;
  if (montantCentimes > 100000) segment = 'enterprise';
  else if (montantCentimes > 10000) segment = 'business';
  else segment = 'starter';
  
  results.push({
    json: {
      client_id: sale.customer_id,
      montant_ttc: montantCentimes,
      montant_ht: montantHT,
      tva: tva,
      devise: 'EUR',
      segment: segment,
      date_vente: dateVente,
      source: sale.channel || 'web',
      pays: sale.country || 'FR',
      created_at: new Date().toISOString()
    }
  });
}
return results;
```
Ce nœud Code illustre plusieurs bonnes pratiques : manipulation des montants en centimes pour éviter les erreurs d'arrondi, normalisation des dates, calcul de la TVA française, et ajout de métadonnées. La sortie formatée est directement compatible avec un nœud PostgreSQL en mode INSERT.

## Étape 8 : Gérer les Erreurs et le Monitoring des Workflows

Un workflow de production doit être résilient. n8n offre plusieurs mécanismes de gestion des erreurs qui permettent de construire des automatisations robustes. Sans une stratégie d'erreur claire, un workflow en production peut échouer silencieusement pendant des jours sans que personne ne s'en aperçoive — un scénario malheureusement courant.

**Error Workflow** : n8n permet de définir un workflow dédié qui s'exécute automatiquement quand n'importe quel autre workflow échoue. Allez dans Settings → Error Workflow et sélectionnez votre workflow de notification d'erreurs. Ce workflow reçoit les détails de l'erreur (nom du workflow, nœud en échec, message d'erreur, timestamp) et peut envoyer des alertes sur Slack, par email, ou dans un outil de monitoring comme PagerDuty.

