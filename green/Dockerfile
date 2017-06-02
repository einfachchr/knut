# von diesem Image starten wir (vanilla nginx)
FROM nginx:latest

# zum Debugging
RUN apt-get update
RUN apt-get install -y curl


# lokale Config-Daten in den Container kopieren
#   Wichtig: auf das richtige Verzeichnis achten!
COPY nginx.conf /etc/nginx/

# HTML-Seiten in den Container
COPY index.html /data/www/html/index.html
COPY welcome.html /data/www/html/welcome.html

# Port 80 nach außen veröffentlichen
EXPOSE 80


# führt das Kommando 'nginx -g "daemon off";' aus
#     -g      global configuration directives
#     daemon  Determines whether nginx should become a daemon.
#             Der Aufruf muss blocken, sonst weiss der Container
#             nicht, ob der Prozess noch läuft.
CMD nginx -g "daemon off;"
