FROM fuzzers/go-fuzz:1.2.0 as builder

ARG project='roaring'
ARG fuzzTargetPath='.'
ARG fuzzTarget='serialization-fuzz'
ARG fuzzFunction='FuzzSerializationBuffer'

ADD . ${project}
WORKDIR ${project}
RUN GO_ENABLED=0 go-fuzz-build -libfuzzer -func=${fuzzFunction} -o /${fuzzTarget}.a ./${fuzzTargetPath}
RUN clang -fsanitize=fuzzer /${fuzzTarget}.a -o /${fuzzTarget}.libfuzzer

FROM ubuntu:20.04
COPY --from=builder /serialization-fuzz.libfuzzer /
