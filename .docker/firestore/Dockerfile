FROM google/cloud-sdk:417.0.1-alpine

RUN apk --update add openjdk11-jre
RUN gcloud components update && \
    gcloud components install cloud-firestore-emulator --quiet

COPY ./.docker/firestore/entrypoint.sh .

ENV PORT 9000
EXPOSE "$PORT"

ENTRYPOINT ["sh", "entrypoint.sh"]
