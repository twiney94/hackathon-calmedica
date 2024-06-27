# Instructions pour démarrer le projet

Pour démarrer le projet, veuillez suivre les étapes ci-dessous :

1. Lancez les services en utilisant Docker Compose :

```bash
docker compose up
```

2. Une fois les services démarrés, exécutez la commande suivante pour lancer ollama avec mistral :

```bash
docker exec -it ollama ollama run mistral
```