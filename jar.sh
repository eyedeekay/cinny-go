#! /usr/bin/env
# You should set JAVA_HOME to point to a Java 17 classpath and set Debian to use Java 17
GO111MODULE=off gojava -o cinny.jar build .
mkdir jar && cd jar && jar -xvf ../cinny.jar && \
    javap -cp ../cinny.jar go.cinnygo.Cinnygo && \
    javap -cp ../cinny.jar go.cinnygo.Cinnygo.CinnyServer