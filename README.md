This instruction outline how to run [this project](https://github.com/aws/aws-lambda-runtime-interface-emulator) 

## Building aws-lambda-runtime-interface-emulator

1. Checkout source code
2. Run `make compile-lambda-linux-all`

## Build Go image

1. Change to `go` directory and run the command

```
 docker build  -f $PWD/Dockerfile  -t golambda:latest .
```

## Run the container

1. Execute the `run.sh` script. The output will look like the following

```
29 Jun 2024 13:45:20,074 [INFO] (rapid) exec '/main' (cwd=/var/task, handler=)
```

2. Use `curl` to test

```
curl -v -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}' | jq .
```

Output will look like the following

```
*   Trying 127.0.0.1:9000...
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* Connected to localhost (127.0.0.1) port 9000 (#0)
> POST /2015-03-31/functions/function/invocations HTTP/1.1
> Host: localhost:9000
> User-Agent: curl/7.81.0
> Accept: */*
> Content-Length: 2
> Content-Type: application/x-www-form-urlencoded
> 
} [2 bytes data]
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sat, 29 Jun 2024 13:46:15 GMT
< Content-Length: 151
< Content-Type: text/plain; charset=utf-8
< 
{ [151 bytes data]
100   153  100   151  100     2   6161     81 --:--:-- --:--:-- --:--:--  6375
* Connection #0 to host localhost left intact
{
  "principalId": "principalID",
  "policyDocument": {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": [
          "execute-api:Invoke"
        ],
        "Effect": "Deny",
        "Resource": [
          ""
        ]
      }
    ]
  }
}
```

3. The output of the emulator will look like the following when it successfully execute the lambda

```
29 Jun 2024 13:45:20,074 [INFO] (rapid) exec '/main' (cwd=/var/task, handler=)
START RequestId: fcc8869b-9472-4ae0-b3b6-a7463926a2ef Version: $LATEST
29 Jun 2024 13:46:15,635 [INFO] (rapid) INIT START(type: on-demand, phase: init)
29 Jun 2024 13:46:15,635 [INFO] (rapid) The extension's directory "/opt/extensions" does not exist, assuming no extensions to be loaded.
29 Jun 2024 13:46:15,635 [INFO] (rapid) Starting runtime without AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_SESSION_TOKEN , Expected?: false
29 Jun 2024 13:46:15,657 [INFO] (rapid) INIT RTDONE(status: success)
29 Jun 2024 13:46:15,657 [INFO] (rapid) INIT REPORT(durationMs: 21.947000)
29 Jun 2024 13:46:15,657 [INFO] (rapid) INVOKE START(requestId: 64c9d6b5-100b-46d0-9166-208b3ceb72dc)
29 Jun 2024 13:46:15,658 [INFO] (rapid) INVOKE RTDONE(status: success, produced bytes: 0, duration: 0.901000ms)
END RequestId: 64c9d6b5-100b-46d0-9166-208b3ceb72dc
REPORT RequestId: 64c9d6b5-100b-46d0-9166-208b3ceb72dc  Init Duration: 0.09 msDuration: 23.04 ms      Billed Duration: 24 ms  Memory Size: 3008 MB    Max Memory Used: 3008 MB
```