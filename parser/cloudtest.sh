#!/bin/sh

docker build . -t europe-west1-docker.pkg.dev/telebot-350915/telebot/parser:test2
docker push europe-west1-docker.pkg.dev/telebot-350915/telebot/parser:test2
gcloud run services update parser --region europe-west1 --image europe-west1-docker.pkg.dev/telebot-350915/telebot/parser:test2