FROM public.ecr.aws/lambda/provided:al2
ADD bin/sheeta /sheeta
ENTRYPOINT [ "/sheeta" ]   
