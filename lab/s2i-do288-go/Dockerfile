FROM centos:7

LABEL Component="go" \
	Name="s2i-do288-go" \
	Version="1.0" \
	Release="1"
	
LABEL io.k8s.description="A Go programming language S2I builder image" \
	io.k8s.display-name="Go programming language S2I builder image for DO288" \
	io.openshift.expose-services="8080:http" \
	io.openshift.s2i.scripts-url="image:///usr/libexec/s2i"

# This label is used to categorize this image as a builder in the Openshift web console
LABEL io.openshift.tags="builder, go, golang"

# The application folder
ENV APP_ROOT /opt/app-root
RUN rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO && \
    curl -s https://mirror.go-repo.io/centos/go-repo.repo | tee /etc/yum.repos.d/go-repo.repo

# Install golang package
RUN yum install -y --setopt=tsflags=nodocs --noplugins golang && \
	yum clean all --noplugins -y
	
COPY ./s2i/bin/ /usr/libexec/s2i
RUN chmod +x /usr/libexec/s2i/assemble
# Permission to run on Openshift
RUN mkdir -p ${APP_ROOT} && \
	chown -R 1001:1001 ${APP_ROOT} && \
	chgrp -R 0 ${APP_ROOT} && \
	chmod -R g=u ${APP_ROOT}
	
# This default user is created in the rhel7 image
USER 1001

EXPOSE 8080

WORKDIR ${APP_ROOT}

CMD ["usage"]
