# otc-rds-flavor

lookup available flavor for RDS in OTC


## usage


```bash
./rdsflavor -h
$ Usage of ./rdsflavor:
  -az1 string
    	AZ instance 1 (e.g. eu-de-01) (default "eu-de-01")
  -az2 string
    	AZ instance 1 (e.g. eu-de-02) (default "eu-de-02")
  -ds string
    	Datastore (e.g. mysql) (default "mysql")
  -dt string
    	Datatype (e.g. 8.0) (default "8.0")
  -ha
    	HA flavor
  -rr
    	RR flavor
  -v	version of the program

```

output from openstack cli

```bash
$ openstack rds datastore version list mysql
+--------------------------------------+------+
| ID                                   | Name |
+--------------------------------------+------+
| bf5a9a94-dbb1-3a0f-b37b-e257595441fa | 5.6  |
| b5863f8e-8081-3066-8221-7b3760218bc3 | 5.7  |
| c4f55bf1-0f4b-32ab-aa98-9becf6bdfef8 | 8.0  |
+--------------------------------------+------+
```

```bash
$ openstack rds flavor list mysql 8.0
+----------------------------+---------------+-------+-----+
| name                       | instance_mode | vcpus | ram |
+----------------------------+---------------+-------+-----+
| rds.mysql.c2.medium.rr     | replica       | 1     |   2 |
| rds.mysql.m1.large.rr      | replica       | 2     |  16 |
| rds.mysql.m1.xlarge.rr     | replica       | 4     |  32 |
| rds.mysql.m1.2xlarge.rr    | replica       | 8     |  64 |
| rds.mysql.s1.medium.ha     | ha            | 1     |   4 |
| rds.mysql.c2.medium.ha     | ha            | 1     |   2 |
| rds.mysql.m1.large.ha      | ha            | 2     |  16 |
| rds.mysql.m1.xlarge.ha     | ha            | 4     |  32 |
| rds.mysql.m1.2xlarge.ha    | ha            | 8     |  64 |
| rds.mysql.s1.medium        | single        | 1     |   4 |
| rds.mysql.c2.medium        | single        | 1     |   2 |
| rds.mysql.m1.large         | single        | 2     |  16 |
| rds.mysql.m1.xlarge        | single        | 4     |  32 |
| rds.mysql.m1.2xlarge       | single        | 8     |  64 |
| rds.mysql.s1.medium.rr     | replica       | 1     |   4 |
| rds.mysql.m3.15xlarge.8.rr | replica       | 60    | 512 |
| rds.mysql.m1.4xlarge       | single        | 16    | 128 |
| rds.mysql.m1.4xlarge.ha    | ha            | 16    | 128 |
| rds.mysql.m1.4xlarge.rr    | replica       | 16    | 128 |
| rds.mysql.m1.8xlarge.rr    | replica       | 32    | 256 |
| rds.mysql.m1.8xlarge       | single        | 32    | 256 |
| rds.mysql.m1.8xlarge.ha    | ha            | 32    | 256 |
| rds.mysql.m3.15xlarge.8    | single        | 60    | 512 |
| rds.mysql.m3.15xlarge.8.ha | ha            | 60    | 512 |
| rds.mysql.s1.large.rr      | replica       | 2     |   8 |
| rds.mysql.s1.xlarge.rr     | replica       | 4     |  16 |
| rds.mysql.s1.2xlarge.rr    | replica       | 8     |  32 |
| rds.mysql.c2.large.rr      | replica       | 2     |   4 |
| rds.mysql.c2.xlarge.rr     | replica       | 4     |   8 |
| rds.mysql.s1.large.ha      | ha            | 2     |   8 |
| rds.mysql.s1.xlarge.ha     | ha            | 4     |  16 |
| rds.mysql.s1.2xlarge.ha    | ha            | 8     |  32 |
| rds.mysql.c2.large.ha      | ha            | 2     |   4 |
| rds.mysql.c2.xlarge.ha     | ha            | 4     |   8 |
| rds.mysql.s1.large         | single        | 2     |   8 |
| rds.mysql.s1.xlarge        | single        | 4     |  16 |
| rds.mysql.s1.2xlarge       | single        | 8     |  32 |
| rds.mysql.c2.large         | single        | 2     |   4 |
| rds.mysql.c2.xlarge        | single        | 4     |   8 |
| rds.mysql.c3.15xlarge.4    | single        | 60    | 256 |
| rds.mysql.c3.15xlarge.4.ha | ha            | 60    | 256 |
| rds.mysql.c3.15xlarge.4.rr | replica       | 60    | 256 |
| rds.mysql.c2.2xlarge       | single        | 8     |  16 |
| rds.mysql.c2.2xlarge.ha    | ha            | 8     |  16 |
| rds.mysql.c2.2xlarge.rr    | replica       | 8     |  16 |
| rds.mysql.c2.4xlarge       | single        | 16    |  32 |
| rds.mysql.c2.4xlarge.ha    | ha            | 16    |  32 |
| rds.mysql.c2.4xlarge.rr    | replica       | 16    |  32 |
| rds.mysql.s1.4xlarge       | single        | 16    |  64 |
| rds.mysql.s1.4xlarge.ha    | ha            | 16    |  64 |
| rds.mysql.s1.4xlarge.rr    | replica       | 16    |  64 |
| rds.mysql.c2.8xlarge       | single        | 32    |  64 |
| rds.mysql.c2.8xlarge.ha    | ha            | 32    |  64 |
| rds.mysql.c2.8xlarge.rr    | replica       | 32    |  64 |
| rds.mysql.s1.8xlarge       | single        | 32    | 128 |
| rds.mysql.s1.8xlarge.ha    | ha            | 32    | 128 |
| rds.mysql.s1.8xlarge.rr    | replica       | 32    | 128 |
| rds.mysql.c3.15xlarge.2.ha | ha            | 60    | 128 |
| rds.mysql.c3.15xlarge.2.rr | replica       | 60    | 128 |
| rds.mysql.c3.15xlarge.2    | single        | 60    | 128 |
+----------------------------+---------------+-------+-----+
```

## Credits

Frank Kloeker f.kloeker@telekom.de

Life is for sharing. If you have an issue with the code or want to improve it, feel free to open an issue or an pull request.
