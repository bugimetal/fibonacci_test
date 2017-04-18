# Tapgerine test #

**Requirements:**
  * Go 1.7 or later is installed

**Packages:**
  * go get github.com/Sirupsen/logrus
  * go get goji.io

**Run:**
  * go build
  * ./fibonacci_test

**Tests:**
  *  2 * 18 With stdout
```
ab -c 10 -n 10000 http://localhost:8080/fib/number/262144
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)

Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /fib/number/262144
Document Length:        54787 bytes

Concurrency Level:      10

Time taken for tests:   23.823 seconds

Complete requests:      10000

Failed requests:        0

Total transferred:      548900000 bytes

HTML transferred:       547870000 bytes

Requests per second:    419.76 [#/sec] (mean)

Time per request:       23.823 [ms] (mean)

Time per request:       2.382 [ms] (mean, across all concurrent requests)

Transfer rate:          22500.47 [Kbytes/sec] received


Connection Times (ms)
              min  mean[+/-sd] median   max
              
Connect:        0    0   0.2      0       5

Processing:     9   24   6.5     23      72

Waiting:        9   24   6.5     23      72

Total:          9   24   6.5     23      73

Percentage of the requests served within a certain time (ms)
  50%     23
  66%     25
  75%     27
  80%     29
  90%     32
  95%     36
  98%     40
  99%     44
 100%     73 (longest request)
```
  * 2**18 without stdout
```
ab -c 10 -n 10000 http://localhost:8080/fib/number/262144
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /fib/number/262144
Document Length:        3 bytes

Concurrency Level:      10
Time taken for tests:   0.885 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1250000 bytes
HTML transferred:       30000 bytes
Requests per second:    11298.98 [#/sec] (mean)
Time per request:       0.885 [ms] (mean)
Time per request:       0.089 [ms] (mean, across all concurrent requests)
Transfer rate:          1379.27 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   9.3      0     387
Processing:     0    0   4.2      0     293
Waiting:        0    0   4.2      0     293
Total:          0    1  10.2      0     387

Percentage of the requests served within a certain time (ms)
  50%      0
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      1
  99%      1
 100%    387 (longest request)
```
  * 2**20 with stdout
```
ab -c 10 -n 10000 http://localhost:8080/fib/number/1048576
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /fib/number/1048576
Document Length:        219142 bytes

Concurrency Level:      10
Time taken for tests:   343.687 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      2192450000 bytes
HTML transferred:       2191420000 bytes
Requests per second:    29.10 [#/sec] (mean)
Time per request:       343.687 [ms] (mean)
Time per request:       34.369 [ms] (mean, across all concurrent requests)
Transfer rate:          6229.70 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   1.0      0      10
Processing:   144  343  89.0    332     950
Waiting:      144  343  88.9    332     949
Total:        144  344  89.0    332     950

Percentage of the requests served within a certain time (ms)
  50%    332
  66%    370
  75%    396
  80%    413
  90%    462
  95%    508
  98%    560
  99%    603
 100%    950 (longest request)
 ```
  * 2**24 without stdout
```
ab -c 10 -n 10000 http://localhost:8080/fib/number/16777216
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8080

Document Path:          /fib/number/16777216
Document Length:        3 bytes

Concurrency Level:      10
Time taken for tests:   15.331 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1250000 bytes
HTML transferred:       30000 bytes
Requests per second:    652.28 [#/sec] (mean)
Time per request:       15.331 [ms] (mean)
Time per request:       1.533 [ms] (mean, across all concurrent requests)
Transfer rate:          79.62 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0  10.7      0     661
Processing:     0   15 446.1      0   14604
Waiting:        0   15 446.1      0   14604
Total:          0   15 446.2      1   14604

Percentage of the requests served within a certain time (ms)
  50%      1
  66%      1
  75%      1
  80%      1
  90%      1
  95%      1
  98%      1
  99%      2
 100%  14604 (longest request)
```
