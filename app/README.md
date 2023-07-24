# Golang Backend template

This simple goland backend has a JWT generator that requires a specific client id and client secret. This is set in the example `.env` file. Utilize the `SAMPLE_REQUEST.sh` file to hit the endpoint. 

The application has CORs setup to allow only the front-end to communicate with the service. In addition, it's already rate limited.