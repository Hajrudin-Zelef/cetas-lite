---
id: collect-261001-general-networking/general-networking/cyberchef-decoder-des-donnees-en-13-etapes-2026-3
title: "cyberchef-decoder-des-donnees-en-13-etapes-2026"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["decode", "attention", "incident"]
source: docs/RAG/collect-261001-general-networking/cyberchef-decoder-des-donnees-en-13-etapes-2026.md
source_anchor: ""
source_lines: [124, 200]
sha256: be5a488f10df0a6ddb06674fc09066aebdde8447f2ff845a8fc300862ccad5cb
---

# cyberchef-decoder-des-donnees-en-13-etapes-2026

- Collez le fichier de logs complet en Input (glisser-déposer le fichier fonctionne aussi, y compris pour des fichiers de plusieurs Mo).
- Ajoutez “Extract IP addresses” à la Recipe, avec l’option “Display total” activée pour compter les occurrences.
- Enchaînez avec “Unique” pour dédupliquer la liste, puis “Sort” pour la trier.
- Terminez par “Filter” si vous souhaitez exclure les plages d’IP internes (RFC 1918) avant de transmettre la liste à un outil de threat intelligence.

Cette combinaison “Extract → Unique → Sort → Filter” constitue l’une des recettes les plus réutilisées en triage d’incident, car elle transforme un fichier brut illisible en une liste d’IOC propre en quelques secondes, prête à être injectée dans un outil de corrélation comme Suricata pour générer des règles de détection.

## Étape 9 : Utiliser Magic Wand pour identifier un encodage inconnu

Que faire quand on ne sait absolument pas comment une donnée a été encodée ? C’est le cas d’usage de l’opération “Magic”, accessible via l’icône en forme de baguette au-dessus du panneau Output. Magic applique en parallèle des dizaines de décodages courants (Base64, Hex, ROT13, Gzip, différents charsets) et note chaque résultat selon un score de vraisemblance basé sur la fréquence des caractères imprimables et la présence de mots reconnaissables.

Dans les paramètres de Magic, augmentez la valeur “Depth” (profondeur) à 3 ou 4 pour permettre à l’opération de tester des encodages imbriqués (comme le double Base64 vu à l’étape 4) sans avoir à les empiler manuellement un par un. Attention cependant : sur un gros fichier binaire, une profondeur élevée peut ralentir sérieusement le navigateur, car chaque combinaison testée consomme de la mémoire côté client. Sur des captures réseau volumineuses ou des dumps mémoire, mieux vaut réduire la profondeur à 1 ou 2 et cibler manuellement l’opération suivante une fois une piste identifiée.

## Étape 10 : Créer, sauvegarder et partager une recette CyberChef

Une fois une recette au point, deux options s’offrent pour la conserver. La première, déjà mentionnée, consiste à copier l’URL de la page : CyberChef sérialise automatiquement la Recipe (mais pas l’Input, sauf activation explicite) dans les paramètres d’URL. La seconde, plus adaptée à un référentiel d’équipe, utilise le bouton “Save recipe” en haut du panneau Recipe, qui exporte la séquence d’opérations au format JSON.

```
[
  { "op": "From Base64", "args": ["A-Za-z0-9+/=", true] },
  { "op": "Decode text", "args": ["UTF-16LE (1200)"] }
]
```
Ce fichier JSON peut être versionné dans un dépôt Git interne aux côtés des runbooks d’investigation, puis réimporté via “Load recipe” par n’importe quel membre de l’équipe SOC. C’est la méthode recommandée pour construire une bibliothèque de recettes standardisées (une par famille de malware, une par type de log source) plutôt que de redécouvrir la même chaîne d’opérations à chaque garde.

## Étape 11 : Automatiser CyberChef avec CyberChef Server (API REST)

Pour intégrer CyberChef dans un pipeline automatisé (par exemple pour traiter en masse les pièces jointes d’une boîte mail de phishing signalé, ou enrichir automatiquement les alertes d’un SIEM), le GCHQ maintient un projet séparé nommé **CyberChef-server**, disponible sur gchq/CyberChef-server. Il expose une API REST permettant d’envoyer une recette et une donnée d’entrée, puis de récupérer le résultat “baked” (cuit, dans le vocabulaire du projet) sans passer par une interface graphique.

```
git clone https://github.com/gchq/CyberChef-server
cd CyberChef-server
docker build -t cyberchef-server .
docker run -d -p 3000:3000 cyberchef-server
```
Une fois le serveur lancé, un appel HTTP suffit pour exécuter une recette de façon programmatique, par exemple depuis un script Python de triage automatique ou une règle SOAR :

```
curl -X POST http://localhost:3000/bake \
  -H "Content-Type: application/json" \
  -d '{
    "input": "cG93ZXJzaGVsbCAtZW5j",
    "recipe": [{ "op": "From Base64", "args": [] }]
  }'
```
Cette approche permet de brancher CyberChef en amont d’un outil de scan de secrets comme GitLeaks dans une chaîne CI/CD, en décodant automatiquement les valeurs encodées avant de les soumettre à la détection de fuite.

## Étape 12 : Écrire une opération JavaScript personnalisée

Quand aucune des opérations natives ne correspond exactement au besoin (un format de log propriétaire, un algorithme d’obfuscation maison rencontré chez un client), CyberChef intègre une opération générique baptisée “JavaScript”, qui exécute du code arbitraire directement sur l’Input.

```
function decodeCustomFormat(input) {
  // Retire un préfixe propriétaire "ENC:" puis inverse la chaîne
  const cleaned = input.replace(/^ENC:/, "");
  return cleaned.split("").reverse().join("");
}
return decodeCustomFormat(input);
```
Glissez l’opération “JavaScript” dans la Recipe, collez ce type de code dans son éditeur intégré, et le résultat s’affiche dans l’Output comme n’importe quelle autre opération. Pour les besoins récurrents, il est également possible de contribuer une opération permanente au projet en soumettant une pull request sur le dépôt GitHub officiel, qui documente la structure exacte attendue (fichier dans `src/core/operations/`, tests unitaires associés) dans son wiki de contribution.

## Étape 13 : Projet complet — analyser un e-mail de phishing de bout en bout

Mettons tout bout à bout sur un cas réaliste, celui d’un e-mail de phishing signalé par un utilisateur, contenant une pièce jointe HTML avec un lien obfusqué. Voici le déroulé complet, étape par étape, tel qu’on le suivrait dans un vrai triage SOC.

1. Exportez l’e-mail au format .eml et ouvrez-le dans un éditeur de texte pour en extraire les en-têtes bruts (Received, Return-Path, Authentication-Results).
2. Collez les en-têtes dans CyberChef et appliquez “Extract email addresses” pour repérer rapidement tous les expéditeurs et domaines impliqués dans la chaîne de relais.
3. Isolez le corps HTML de la pièce jointe et appliquez “Strip HTML tags” puis “Extract URLs” pour lister tous les liens présents, y compris ceux cachés dans des attributs non visibles à l’affichage.
4. Pour chaque URL suspecte utilisant un raccourcisseur ou un encodage, appliquez “URL Decode” puis “From Base64” si un paramètre de la requête semble encodé (cas fréquent des kits de phishing qui embarquent l’adresse e-mail de la victime encodée dans le lien pour pré-remplir un faux formulaire).
5. Si la pièce jointe contient un script JavaScript obfusqué, réutilisez la recette “From Base64 → Decode text” mise au point à l’étape 7, en adaptant le charset si nécessaire.
6. Calculez le hash SHA-256 de la pièce jointe complète avec l’opération “SHA2” pour le documenter dans le ticket d’incident et le comparer à des bases de réputation.
7. Sauvegardez la recette complète au format JSON (étape 10) et archivez-la avec le rapport d’incident pour que la prochaine campagne similaire soit traitée deux fois plus vite.

Ce projet illustre bien la logique de CyberChef : chaque étape isolée est triviale, mais la capacité à les enchaîner sans changer d’outil ni recopier des données d’un terminal à un autre fait gagner un temps considérable sur un triage réel, où les premières minutes après un signalement comptent souvent le plus.

## Intégrer CyberChef dans un workflow SOC

