CREATE TABLE `todo_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` text,
  `status` enum('Doing', 'Done', 'Deleted') NULL DEFAULT 'Doing',
  `image` json NULL,
  `int_example` int(11) NULL,
  `double_example` double DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE dbtest.users (
	id INT auto_increment NOT NULL,
	email varchar(100) NOT NULL,
	salt varchar(100) NULL,
	password varchar(100) NOT NULL,
	first_name varchar(100) NULL,
	last_name varchar(100) NULL,
	phone varchar(100) NULL,
	`role` ENUM('user', 'admin', 'shipper', 'mod') DEFAULT 'user' NULL,
	status INT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP  NULL,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;



CREATE TABLE dbtest.user_like_items (
	user_id INT NOT NULL,
	item_id INT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP  NULL,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP NULL,
	CONSTRAINT user_like_items_users_FK FOREIGN KEY (user_id) REFERENCES dbtest.users(id),
	CONSTRAINT user_like_items_todo_items_FK FOREIGN KEY (item_id) REFERENCES dbtest.todo_items(id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
