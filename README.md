## Steps to run:

1. Clone this repository
2. Install dependencies : `go mod tidy && go mod vendor`
3. Make sure Kafka is running, and change the value of `brokerAddress` to the address of you Kafka instance
4. Rename .env.default to .env, and the change all the values based on your own configurations
5. Run the code : `go run main.go`

## Email Service - Kafka Consumer

Check this repo to get Email Service (Kafka Consumer) https://github.com/klanjabrik/taskcrud-email-service