FROM golang:1.17

WORKDIR /app
COPY entrypoint.sh entrypoint.sh
RUN ls -la .
RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
