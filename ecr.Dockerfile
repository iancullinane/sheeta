FROM public.ecr.aws/lambda/go:1

ADD bin/main main

CMD [ "main" ]

