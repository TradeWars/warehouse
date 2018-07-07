FROM scratch
ADD warehouse warehouse
ENTRYPOINT [ "./warehouse" ]
