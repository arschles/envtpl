FROM alpine:3.3
COPY envtpl .
RUN mv envtpl /bin
CMD envtpl
