FROM plugins/base:linux-amd64


COPY release/linux/amd64/helloworld /bin/

ENTRYPOINT ["/bin/helloworld"]