FROM registry.access.redhat.com/ubi8/python-36:latest

USER root

RUN dnf install -y libvirt-devel

RUN pip3 install --upgrade pip \
 && pip3 install virtualbmc

ENTRYPOINT ["vbmcd"]
CMD ["--foreground"]