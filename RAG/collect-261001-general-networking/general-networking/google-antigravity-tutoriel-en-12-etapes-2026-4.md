---
id: collect-261001-general-networking/general-networking/google-antigravity-tutoriel-en-12-etapes-2026-4
title: "macOS (Homebrew)"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Google"]
dates: []
keywords: ["agent", "agents", "attention", "claude", "gemini", "opus 4"]
source: docs/RAG/collect-261001-general-networking/google-antigravity-tutoriel-en-12-etapes-2026.md
source_anchor: ""
source_lines: [229, 326]
sha256: f1f089cfab367b7e793783b06012404a662b3ba4fa35146c78f5219e643ddb54
---

# macOS (Homebrew)

## Étape 9 : lire et valider les Artifacts

Les **Artifacts** sont la manière dont Antigravity rend le travail des agents vérifiable, au lieu de vous inonder de journaux techniques. Chaque mission en produit plusieurs types, que vous consultez dans un onglet dédié du Manager.

- **Liste de tâches** : la décomposition du travail, avec l’état d’avancement de chaque sous-tâche (à faire, en cours, terminé).
- **Plan d’implémentation** : les décisions de conception, l’arborescence des fichiers, l’ordre d’exécution. C’est le document que vous validez avant l’écriture.
- **Captures d’écran** : des images de l’interface au fur et à mesure, pratiques pour repérer un problème visuel sans lancer l’application.
- **Enregistrements du navigateur** : de courtes séquences d’actions montrant l’agent qui teste l’application dans Chrome.

L’atout majeur : vous laissez vos commentaires **directement sur l’Artifact**, et l’agent les intègre sans interrompre son exécution. Par exemple, sur le plan d’implémentation, vous pouvez annoter « Ajoute une validation : le titre ne doit pas être vide » et l’agent ajustera le code en conséquence. Cette boucle de rétroaction non bloquante est ce qui rend le travail multi-agents réellement gérable.

Prenez l’habitude de lire le plan d’implémentation *avant* de laisser l’agent coder pendant dix minutes. Corriger une hypothèse erronée au stade du plan coûte quelques secondes ; la corriger après 300 lignes générées coûte bien plus. C’est la discipline centrale du travail avec un IDE agentique.

## Étape 10 : vérifier l’application avec le sous-agent navigateur

La force distinctive d’Antigravity est sa capacité à **contrôler un navigateur** pour vérifier ce qu’il a construit. Le sous-agent navigateur ouvre Chrome, charge l’application, clique, remplit des champs et confirme que tout fonctionne – puis produit un enregistrement en Artifact. Pour l’activer, installez l’extension Chrome d’Antigravity depuis la boîte de dialogue proposée au premier usage du navigateur, et autorisez la connexion entre l’IDE et le navigateur.

Une fois le backend et le front générés, demandez à l’agent de lancer et de vérifier l’application :

```
Consigne à l'agent :
Lance l'application (uvicorn main:app --reload), ouvre http://127.0.0.1:8000
dans le navigateur, ajoute une tâche "Acheter du pain", coche-la comme faite,
puis supprime-la. Confirme par une capture que la liste est vide à la fin.
```
L’agent démarre le serveur, ouvre l’URL, exécute le scénario et enregistre la séquence. Voici l’extrait de front généré qui alimente la page (à titre d’illustration du résultat) :

```
<!-- index.html – extrait généré -->
<script>
async function chargerTaches() {
  const r = await fetch('/taches');
  const taches = await r.json();
  const liste = document.getElementById('liste');
  liste.innerHTML = '';
  for (const t of taches) {
    const li = document.createElement('li');
    li.textContent = t.titre + (t.faite ? ' ✓' : '');
    liste.appendChild(li);
  }
}
async function ajouter() {
  const titre = document.getElementById('titre').value;
  await fetch('/taches', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ titre, faite: false })
  });
  chargerTaches();
}
chargerTaches();
</script>
```
Attention à la sécurité : donner à un agent le contrôle du navigateur et du terminal est puissant mais sensible. Un agent qui navigue sur des pages tierces peut être exposé à des attaques par **injection de prompt** (des instructions cachées dans une page web tentant de détourner l’agent). Limitez le sous-agent navigateur à des URL de confiance (votre application locale), et ne le laissez pas parcourir librement le web avec des accès sensibles. Nous détaillons ces précautions dans la section astuces avancées.

## Étapes 11 et 12 : réviser, tester, committer et livrer

**Étape 11 – Réviser et itérer.** Repassez en vue Éditeur pour relire le code généré. Antigravity affiche les modifications sous forme de *diffs* que vous approuvez ou rejetez. Si un endpoint ne vous convient pas, sélectionnez-le et lancez une commande en ligne (« ajoute une gestion d’erreur 404 si l’id n’existe pas »). Les points de contrôle (checkpoints) automatiques vous permettent de revenir à un état antérieur en un clic si une itération casse quelque chose. C’est le moment d’affiner, pas de tout regénérer.

**Étape 12 – Tester et livrer.** Lancez la suite de tests générée pour valider le projet complet. Dans le terminal intégré :

```
pip install -r requirements.txt
pytest -q
```
Sortie attendue :

```
....                                                      [100%]
4 passed in 0.42s
```
Une fois les tests au vert, versionnez le projet. Antigravity peut rédiger le message de commit pour vous, mais vous pouvez aussi le faire manuellement :

```
git add .
git commit -m "feat: gestionnaire de taches (API FastAPI + front + tests)"
# Vérifier l'application une dernière fois
uvicorn main:app --reload
# Ouvrir http://127.0.0.1:8000 dans le navigateur
```
Félicitations : vous disposez d’un projet complet, testé et versionné, généré et vérifié par des agents. Vous pouvez désormais le déployer (par exemple sur un hébergeur compatible ASGI) ou l’enrichir avec de nouvelles missions. Le tout, sans avoir écrit une seule ligne à la main – mais en gardant le contrôle à chaque étape grâce aux plans, aux diffs et aux Artifacts.

## 5 pièges courants à éviter avec Google Antigravity

Les IDE agentiques transforment la productivité, mais ils introduisent aussi de nouveaux pièges. Voici les cinq erreurs les plus fréquentes chez les débutants – et comment les éviter.

1. **Lancer une mission sans AGENTS.md.** Sans règles, l’agent improvise la pile technique et les conventions. Résultat : du code incohérent, parfois des dépendances non désirées. Rédigez toujours`AGENTS.md` avant la première mission.
2. **Ne pas relire le plan d’implémentation.** Approuver un plan à la va-vite, c’est laisser l’agent bâtir sur une mauvaise fondation. Lisez le plan, corrigez-le, puis seulement laissez l’agent coder.
3. **Épuiser son quota sur des tâches triviales.** Utiliser Gemini 3 Pro ou Claude Opus 4.6 pour renommer une variable gaspille du quota. Réservez les gros modèles à la planification et à l’architecture ; passez sur Flash pour les petites itérations.
4. **Laisser le sous-agent navigateur en liberté.** Un agent qui parcourt le web peut subir une injection de prompt. Cantonnez-le à vos URL locales et de confiance.
5. **Ignorer Git et les checkpoints.** Sur une session longue, un agent peut casser du code qui marchait. Sans commits réguliers, vous perdez du temps à défaire les dégâts. Committez après chaque étape stable.

Un sixième piège mérite d’être cité : la **confiance aveugle**. Un agent produit du code plausible, pas forcément correct. Les tests générés automatiquement ne couvrent pas tout. Gardez un œil critique, surtout sur la logique métier et la sécurité. L’IDE agentique est un accélérateur, pas un pilote automatique infaillible.

## Dépannage : 8 problèmes fréquents et leurs solutions

Voici les incidents les plus courants rencontrés lors de l’installation et de l’utilisation de Google Antigravity, avec la cause probable et la solution.

