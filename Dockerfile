FROM busybox:ubuntu-14.04

ENTRYPOINT ["/optimusPrime"]

COPY optimusPrime /optimusPrime
RUN chmod a+x /optimusPrime
