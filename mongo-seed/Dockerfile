FROM mongo:latest

ARG MONGO_INITDB_ROOT_USERNAME
ARG MONGO_INITDB_ROOT_PASSWORD

COPY init.json /init.json
CMD mongoimport --host mongodb --db sample_mflix --collection init --type json --file /init.json --jsonArray --username=${MONGO_INITDB_ROOT_USERNAME} --password=${MONGO_INITDB_ROOT_PASSWORD} --authenticationDatabase=admin