# Golang Authen with JWT
# Create table
```sh
CREATE DATABASE lak;
USE lak
```

# Create table
```sh
DROP TABLE IF EXISTS `products`;

CREATE TABLE `products` (
  `id` varchar(20) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `price` float DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` varchar(20) NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

# Change config file (config/myconfig.json)
```sh
{
  "mydb": "root:root@tcp(127.0.0.1:3306)/lak",
  "secret_key": "LeAnKhang"
}
```

# Handle concurrent requests
- According to my knowledge, we should use Load Balancer to forwards the user request to many backend server 
- With viewing Product we can configure some Slave DB for it
