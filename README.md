# Travaux-memoire
Evaluation de l'impact de la connectivité internet sur l'accès aux services cloud depuis le Cameroun.

## Implémentation 
Nous allons utilisé une procédure basée sur des mesures de performance effectuées sur une infrastructure Cloud comme banc d’essai.
Pour le moment, les aspects des services Cloud à évaluer sont sur SaaS.

Les mesures de performance qui seront analysées sont :
- Temps de réponse du service (délai)
- Latence
- Bande passante
- Débit

## Etapes

1) Générer du trafic depuis le client sur les services cloud

2) Collecter les données depuis le client à l'aide du trafic généré

3) Analysée les données à l'aide d'outils : Prometheus et Grafana.


### 1. Générer du trafic depuis le client sur les services cloud

Nous allons  utiliser un script personnalisés à l'aide de l'outil curl pour générer du trafic depuis le client sur les services cloud:

1. Installez l'outil curl sur le système.
2. Écrire le script  en utilisant la syntaxe curl pour envoyer des requêtes HTTP aux services cloud.
3. Le script enverra  des requêtes HTTP aux services cloud en utilisant l'outil curl, ce qui générera du trafic depuis le client.
4. Dans notre cas spécifique :

	- Le script envoie 10 requêtes GET à une ou plusieurs URL spécifiée (dans l'exemple ci-dessus, `https://www.google.com` et `https://cloud.google.com/apis`), en attendant 5 secondes entre chaque requête.
	-  On peut modifier la valeur `NUM_REQUESTS` et `WAIT_TIME` pour en envoyer moins ou plus et ajuster l'intervalle de temps entre chaque requête pour répondre à nos besoins.

### 2.   Collecter les données depuis le client à l’aide du trafic généré

Pour collecter les données depuis le client à l'aide du trafic généré précédemment, on va  utiliser l'outil **collectd**.

**Collectd** est un démon de collecte de données open source qui fonctionne en collectant des données de performance à partir de différents services cloud et de systèmes d'exploitation.

- Le processus de collecte de données de collectd consiste à interroger régulièrement les sources de données pour récupérer des informations telles que  la consommation de bande passante, la latence, le temps de réponse, etc.

- Une fois que les données sont collectées, collectd les stocke dans un format standardisé ( exemple: csv)

- Les données peuvent ensuite être envoyées à un système de surveillance ou à un outil de visualisation pour analyse ultérieure (exemple: Prometheus, Grafana).
