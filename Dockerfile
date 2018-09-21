FROM debian

# Non-root user `app`
RUN useradd app
WORKDIR /home/app

COPY bin/ipify-api ./

ENV LOGGING_LEVEL WARNING

COPY docker-entrypoint.sh ./

RUN chown -R app:app /home/app

# Change to user `app`
USER app

ENTRYPOINT ["./docker-entrypoint.sh"]

CMD ["./ipify-api"]
