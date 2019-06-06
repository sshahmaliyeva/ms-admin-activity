FROM debian

WORKDIR /app
COPY ms-admin-activity ./

EXPOSE 80

ENTRYPOINT [ "./ms-admin-activity" ]
