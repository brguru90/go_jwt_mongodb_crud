<meta name="description" content="go_gin_jwt_mongodb_crud  | Go/psql and Go/mongo and node js/psql">




# Testing Configuration [Go gin framework Benchmark]

```
  1. Mongo 5.0.6
  2. System memory 16GB
  3. AMD ryzen 7 5800H
  4. SSD 
  5. OS : ubuntu 21
```



## 1. Go run command [Will be used in comparison - <a href="https://github.com/brguru90/go_jwt_sql_crud">Go/psql</a> and  <a href="https://github.com/brguru90/go_jwt_mongodb_crud">Go/mongo</a> and <a href="https://github.com/brguru90/node-sql-and-jwt-demo">node js/psql]</a>

### ROW Count =  1220001  
### gin framework 

```
1. JWT auth + redis block list
2. Single record from DB at a time on 2nd benchmark  
3. 20 record from starting
4. 20 record from some where before end
```

```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/login_status/
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   3.723 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      24000000 bytes
HTML transferred:       13100000 bytes
Requests per second:    26860.10 [#/sec] (mean)
Time per request:       37.230 [ms] (mean)
Time per request:       0.037 [ms] (mean, across all concurrent requests)
Transfer rate:          6295.33 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   20  85.7     12    1052
Processing:     1   15  15.0     13     225
Waiting:        0   11  14.9      9     221
Total:          1   35  87.1     24    1073

Percentage of the requests served within a certain time (ms)
  50%     24
  66%     33
  75%     35
  80%     35
  90%     37
  95%     38
  98%     41
  99%    215
 100%   1073 (longest request)
```
```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/user/
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   4.770 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      78500000 bytes
HTML transferred:       67600000 bytes
Requests per second:    20966.12 [#/sec] (mean)
Time per request:       47.696 [ms] (mean)
Time per request:       0.048 [ms] (mean, across all concurrent requests)
Transfer rate:          16072.66 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   24 104.4     13    1042
Processing:     1   22  12.4     21     245
Waiting:        1   16  11.9     15     237
Total:          1   46 105.3     33    1072

Percentage of the requests served within a certain time (ms)
  50%     33
  66%     36
  75%     40
  80%     42
  90%     48
  95%     51
  98%     56
  99%   1031
 100%   1072 (longest request)
```
```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/user/?page=1&limit=20
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   7.344 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1319000000 bytes
HTML transferred:       1310200000 bytes
Requests per second:    13617.08 [#/sec] (mean)
Time per request:       73.437 [ms] (mean)
Time per request:       0.073 [ms] (mean, across all concurrent requests)
Transfer rate:          175399.72 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.9      0      16
Processing:     1   72  51.4     67     397
Waiting:        1   71  51.3     66     397
Total:          1   73  51.5     68     398

Percentage of the requests served within a certain time (ms)
  50%     68
  66%     91
  75%    104
  80%    113
  90%    143
  95%    171
  98%    196
  99%    216
 100%    398 (longest request)
```
```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/user/?page=1000&limit=20
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   97.956 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1319300000 bytes
HTML transferred:       1310500000 bytes
Requests per second:    1020.87 [#/sec] (mean)
Time per request:       979.560 [ms] (mean)
Time per request:       0.980 [ms] (mean, across all concurrent requests)
Transfer rate:          13152.63 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   1.2      0      14
Processing:    15  975  76.2    971    1263
Waiting:        7  975  76.2    971    1263
Total:         28  975  75.5    971    1263

Percentage of the requests served within a certain time (ms)
  50%    971
  66%    998
  75%   1017
  80%   1028
  90%   1057
  95%   1079
  98%   1103
  99%   1117
 100%   1263 (longest request)
```



## 2. Go run command

### ROW Count =  1220001  
### gin framework

### Change in source code: 
`1. added redis cache for /user/ API`

```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/login_status/
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   3.552 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      24000000 bytes
HTML transferred:       13100000 bytes
Requests per second:    28149.45 [#/sec] (mean)
Time per request:       35.525 [ms] (mean)
Time per request:       0.036 [ms] (mean, across all concurrent requests)
Transfer rate:          6597.53 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   17  70.3     12    1038
Processing:     3   16  19.0     13     232
Waiting:        0   12  18.8      9     228
Total:          5   32  72.8     25    1052

Percentage of the requests served within a certain time (ms)
  50%     25
  66%     27
  75%     28
  80%     29
  90%     35
  95%     36
  98%     62
  99%    222
 100%   1052 (longest request)
```
```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/user/
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   4.031 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      84288988 bytes
HTML transferred:       67600000 bytes
Requests per second:    24807.23 [#/sec] (mean)
Time per request:       40.311 [ms] (mean)
Time per request:       0.040 [ms] (mean, across all concurrent requests)
Transfer rate:          20419.69 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   18  65.4     14    1050
Processing:     6   22  13.2     21     240
Waiting:        1   17  13.0     15     231
Total:          8   40  66.4     37    1077

Percentage of the requests served within a certain time (ms)
  50%     37
  66%     40
  75%     41
  80%     42
  90%     44
  95%     45
  98%     48
  99%     55
 100%   1077 (longest request)
```
```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/user/?page=1&limit=20
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   5.250 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1324789063 bytes
HTML transferred:       1310200000 bytes
Requests per second:    19045.96 [#/sec] (mean)
Time per request:       52.505 [ms] (mean)
Time per request:       0.053 [ms] (mean, across all concurrent requests)
Transfer rate:          246405.04 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   23 117.3     10    1054
Processing:     2   28   8.6     27     114
Waiting:        1   18   9.1     15     102
Total:          2   51 118.3     36    1092

Percentage of the requests served within a certain time (ms)
  50%     36
  66%     39
  75%     41
  80%     44
  90%     51
  95%     55
  98%     65
  99%   1059
 100%   1092 (longest request)
```
```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8000

Document Path:          /api/user/?page=1000&limit=20
Document Length:        Variable

Concurrency Level:      1000
Time taken for tests:   5.277 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1325088840 bytes
HTML transferred:       1310500000 bytes
Requests per second:    18950.60 [#/sec] (mean)
Time per request:       52.769 [ms] (mean)
Time per request:       0.053 [ms] (mean, across all concurrent requests)
Transfer rate:          245226.88 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   20  95.7     12    1031
Processing:     4   30   9.3     29     243
Waiting:        1   18   9.0     16     225
Total:          4   51  96.4     40    1076

Percentage of the requests served within a certain time (ms)
  50%     40
  66%     44
  75%     49
  80%     51
  90%     56
  95%     59
  98%     65
  99%     76
 100%   1076 (longest request)
```