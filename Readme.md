<meta name="description" content="go_gin_jwt_mongodb_crud  | Go/psql and Go/mongo and node js/psql">



# go_jwt_mongodb_crud <br />

# Testing Configuration [Go gin framework Benchmark]

```
  1. Mongo 5.0.6
  2. System memory 16GB
  3. AMD ryzen 7 5800H
  4. SSD 
  5. OS : ubuntu 21
```



## 1. Go run command [Will be used in comparison - <a href="https://github.com/brguru90/go_jwt_sql_crud">Go/psql</a> and  <a href="https://github.com/brguru90/go_jwt_mongodb_crud">Go/mongo</a> and <a href="https://github.com/brguru90/node-sql-and-jwt-demo">node js/psql</a>]

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
Time taken for tests:   3.556 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      24000000 bytes
HTML transferred:       13100000 bytes
Requests per second:    28125.20 [#/sec] (mean)
Time per request:       35.555 [ms] (mean)
Time per request:       0.036 [ms] (mean, across all concurrent requests)
Transfer rate:          6591.84 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   19  81.8     12    1040
Processing:     0   16  19.9     14     228
Waiting:        0   12  20.0      9     221
Total:          0   35  83.9     26    1057

Percentage of the requests served within a certain time (ms)
  50%     26
  66%     28
  75%     31
  80%     33
  90%     35
  95%     36
  98%     92
  99%    227
 100%   1057 (longest request)
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
Time taken for tests:   4.034 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      84288962 bytes
HTML transferred:       67600000 bytes
Requests per second:    24789.38 [#/sec] (mean)
Time per request:       40.340 [ms] (mean)
Time per request:       0.040 [ms] (mean, across all concurrent requests)
Transfer rate:          20404.99 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   17  49.8     15    1047
Processing:     6   23  11.1     22     228
Waiting:        0   17  11.1     16     221
Total:         15   40  51.0     38    1076

Percentage of the requests served within a certain time (ms)
  50%     38
  66%     40
  75%     41
  80%     42
  90%     43
  95%     45
  98%     47
  99%     49
 100%   1076 (longest request)
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
Time taken for tests:   5.051 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1420089072 bytes
HTML transferred:       1405500000 bytes
Requests per second:    19797.26 [#/sec] (mean)
Time per request:       50.512 [ms] (mean)
Time per request:       0.051 [ms] (mean, across all concurrent requests)
Transfer rate:          274549.49 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   18  63.7     15    1043
Processing:     7   32   7.4     34      70
Waiting:        1   18   6.2     18      59
Total:         12   50  64.6     49    1088

Percentage of the requests served within a certain time (ms)
  50%     49
  66%     51
  75%     52
  80%     53
  90%     55
  95%     57
  98%     61
  99%     66
 100%   1088 (longest request)
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
Time taken for tests:   5.022 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      1427488932 bytes
HTML transferred:       1412900000 bytes
Requests per second:    19911.29 [#/sec] (mean)
Time per request:       50.223 [ms] (mean)
Time per request:       0.050 [ms] (mean, across all concurrent requests)
Transfer rate:          277569.80 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   19  75.5     13    1034
Processing:     9   31  11.5     31     246
Waiting:        1   18  11.4     16     229
Total:         11   50  76.5     45    1072

Percentage of the requests served within a certain time (ms)
  50%     45
  66%     48
  75%     50
  80%     51
  90%     53
  95%     55
  98%     59
  99%     64
 100%   1072 (longest request)
guruprasad@guruprasad-Nitro-AN515-45:~$ 
```
