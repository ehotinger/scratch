FROM alpine
COPY ./addemup.sh foo/addemup.sh
WORKDIR /foo
ENTRYPOINT [ "/bin/sh" , "./addemup.sh"]
CMD ["1", "2"]