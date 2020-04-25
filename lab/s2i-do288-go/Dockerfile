# s2i-do288-go
FROM centos:7

# TODO: Put the maintainer name in the image metadata
LABEL maintainer="Marco"

# TODO: Rename the builder environment variable to inform users about application you provide them
ENV BUILDER_VERSION 1.0
ENV APP_ROOT=/opt/app-root
# TODO: Set labels used in OpenShift to describe the builder image
LABEL io.k8s.description="Platform for building xyz" \
      io.k8s.display-name="builder x.y.z" \
      io.openshift.expose-services="8080:http" \
      io.openshift.tags="builder,x.y.z,etc." \
      io.openshift.s2i.scripts-url="image:///usr/libexec/s2i"
# TODO: Install required packages here:
# RUN yum install -y ... && yum clean all -y
COPY go/ /usr/local/go

ENV GOROOT=/usr/local/go
ENV PATH $PATH:/usr/local/go/bin

#tar -zxvf go1.8.3.linux-amd64.tar.gz -C /usr/local
# TODO (optional): Copy the builder files into /opt/app-root
# COPY ./<builder_folder>/ /opt/app-root/

# TODO: Copy the S2I scripts to /usr/libexec/s2i, since openshift/base-centos7 image
# sets io.openshift.s2i.scripts-url label that way, or update that label
COPY ./.s2i/bin/ /usr/libexec/s2i

# TODO: Drop the root user and make the content of /opt/app-root owned by user 1001
# RUN chown -R 1001:1001 /opt/app-root

RUN mkdir ${APP_ROOT} && \
    chown -R 1001:1001 /opt/app-root

# This default user is created in the openshift/base-centos7 image
USER 1001



WORKDIR ${APP_ROOT}

# TODO: Set the default port for applications built using this image
EXPOSE 8080

# TODO: Set the default CMD for the image
CMD ["go", "version"]
