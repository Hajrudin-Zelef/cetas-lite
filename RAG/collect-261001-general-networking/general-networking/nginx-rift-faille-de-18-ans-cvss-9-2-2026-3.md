---
id: collect-261001-general-networking/general-networking/nginx-rift-faille-de-18-ans-cvss-9-2-2026-3
title: "Vérifier la version de NGINX installée"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "attention", "cybersecurity", "distribution", "exploit", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/nginx-rift-faille-de-18-ans-cvss-9-2-2026.md
source_anchor: ""
source_lines: [107, 167]
sha256: 93776ce736ee31bf37afb98e0623dd01e64f4989dd5216951a8fe0ffaadd8d95
---

# Vérifier la version de NGINX installée

```
# Vérifier la version de NGINX installée
nginx -v
# Sur une distribution basée sur Debian/Ubuntu
apt list --installed | grep nginx
# Sur une distribution basée sur RHEL/AlmaLinux
rpm -qa | grep nginx
# Sur un cluster Kubernetes, vérifier l'image de l'Ingress Controller
kubectl get pods -A -o jsonpath="{.items[*].spec.containers[*].image}" | tr ' ' '\n' | grep -i nginx
```
En complément de la mise à jour, plusieurs mesures d’atténuation permettent de réduire le risque en attendant une fenêtre de maintenance :

- Auditer les fichiers de configuration à la recherche de directives `rewrite` suivies d’une directive`rewrite` ,`if` ou`set` utilisant des captures PCRE non nommées.
- Remplacer les captures non nommées (`$1` ,`$2` ) par des captures nommées lorsque c’est possible, en attendant la mise à jour.
- Vérifier que l’ASLR est bien activée sur l’ensemble des hôtes exposés, en particulier les conteneurs et machines virtuelles récemment provisionnés.
- Mettre en place une règle de pare-feu applicatif (WAF) en « virtual patching » le temps de qualifier la montée de version en environnement de test.
- Surveiller les journaux pour détecter des plantages anormaux et répétés des processus *worker* de NGINX, signe potentiel de tentatives d’exploitation.

## 5 prédictions pour la suite

Au vu de la chronologie observée depuis mai 2026, plusieurs évolutions semblent probables dans les prochains mois :

1. **D’autres CVE viseront des modules NGINX anciens.** L’attention portée à NGINX Rift va probablement pousser des chercheurs indépendants à auditer d’autres modules hérités des débuts du projet, avec de nouvelles découvertes possibles d’ici la fin 2026.
2. **L’analyse automatisée de versions va s’accélérer.** Les entreprises qui ont découvert leur exposition tardivement vont investir davantage dans des scanners de vulnérabilités intégrés au pipeline de déploiement, à l’image d’outils comme Trivy.
3. **La pression réglementaire européenne va s’intensifier.** Entre NIS2 et le Cybersecurity Act 2, les obligations de délai de correction pour les infrastructures numériques essentielles devraient se préciser, avec cette faille NGINX comme exemple récurrent dans les débats.
4. **Le recours au « virtual patching » par WAF va progresser.** Face à des cycles de correctifs rapprochés, de plus en plus d’organisations vont s’appuyer sur des règles de pare-feu applicatif comme filet de sécurité temporaire.
5. **De nouvelles vagues d’exploitation opportunistes sont à prévoir.** Comme pour la plupart des failles critiques largement médiatisées en 2026, les serveurs qui n’auront pas appliqué les correctifs d’ici la fin de l’été risquent de devenir des cibles privilégiées pour des campagnes automatisées.

## Foire aux questions

### Qu’est-ce que CVE-2026-42945, alias « NGINX Rift » ?

Il s’agit d’un débordement de tampon en zone de tas dans le module de réécriture d’URL de NGINX (`ngx_http_rewrite_module`), noté CVSS 9,2 sur l’échelle version 4.0, divulgué le 13 mai 2026 et activement exploité depuis.

### Mon serveur NGINX est-il vulnérable ?

Si votre version de NGINX Open Source se situe entre 0.6.27 et 1.30.0, ou votre NGINX Plus entre R32 et R36, et que votre configuration utilise des directives `rewrite` avec captures non nommées, vous êtes potentiellement exposé à cette faille NGINX. La commande `nginx -v` permet de vérifier rapidement la version installée.

### Comment savoir si mon serveur a déjà été compromis ?

Recherchez dans les journaux des redémarrages ou plantages anormaux et répétés des processus *worker* de NGINX, ainsi que des requêtes contenant des motifs de réécriture inhabituels ou des chaînes anormalement longues dans l’URI.

### Les failles du 18 juin sont-elles liées à NGINX Rift ?

Non, CVE-2026-42530 et CVE-2026-42055 touchent des modules différents (HTTP/3 et proxy/gRPC) de celui concerné par NGINX Rift. Elles ont cependant été révélées seulement cinq semaines plus tard, ce qui a renforcé la pression sur les équipes de correction.

### NGINX Ingress Controller sur Kubernetes est-il concerné ?

Oui, plusieurs versions de NGINX Ingress Controller (3.5.0 à 3.7.2, 4.0.0 à 4.0.1 et 5.0.0 à 5.5.0) sont concernées par les failles corrigées le 18 juin 2026. Une mise à jour de l’image du contrôleur est nécessaire.

### Existe-t-il un correctif temporaire si je ne peux pas mettre à jour immédiatement ?

Remplacer les captures PCRE non nommées par des captures nommées dans les directives de réécriture réduit le risque lié à NGINX Rift. Une règle de pare-feu applicatif en mode « virtual patching » peut aussi limiter l’exposition en attendant la montée de version définitive.

### Quelle est la différence entre NGINX Open Source et NGINX Plus face à ces failles ?

Les deux éditions sont concernées, mais avec des versions et des numéros de correctifs différents : NGINX Plus reçoit des correctifs dédiés identifiés par des lettres « P » (par exemple R36 P4), distribués séparément du cycle de NGINX Open Source.

### NGINX reste-t-il le serveur web le plus utilisé au monde malgré cette faille ?

Oui. Malgré un léger recul de 33,3 % à 31,8 % de parts de marché entre janvier et juillet 2026 selon W3Techs, NGINX conserve la première place devant Apache et le serveur de Cloudflare.
