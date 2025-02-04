FROM mongo:latest

ARG MONGO_INITDB_ROOT_USERNAME
ARG MONGO_INITDB_ROOT_PASSWORD

COPY comments.json /comments.json
COPY movies.json /movies.json
COPY sessions.json /sessions.json
COPY theaters.json /theaters.json
COPY users.json /users.json

CMD mongoimport --host mongodb --db sample_mflix --collection comments --type json --file /comments.json --jsonArray --username=${MONGO_INITDB_ROOT_USERNAME} --password=${MONGO_INITDB_ROOT_PASSWORD} --authenticationDatabase=admin && \
 mongoimport --host mongodb --db sample_mflix --collection movies --type json --file /movies.json --jsonArray --username=${MONGO_INITDB_ROOT_USERNAME} --password=${MONGO_INITDB_ROOT_PASSWORD} --authenticationDatabase=admin && \
 mongoimport --host mongodb --db sample_mflix --collection sessions --type json --file /sessions.json --jsonArray --username=${MONGO_INITDB_ROOT_USERNAME} --password=${MONGO_INITDB_ROOT_PASSWORD} --authenticationDatabase=admin && \
 mongoimport --host mongodb --db sample_mflix --collection theaters --type json --file /theaters.json --jsonArray --username=${MONGO_INITDB_ROOT_USERNAME} --password=${MONGO_INITDB_ROOT_PASSWORD} --authenticationDatabase=admin && \
 mongoimport --host mongodb --db sample_mflix --collection users --type json --file /users.json --jsonArray --username=${MONGO_INITDB_ROOT_USERNAME} --password=${MONGO_INITDB_ROOT_PASSWORD} --authenticationDatabase=admin