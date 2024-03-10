-- Create table for Job Listings
CREATE TABLE Jobs (
  id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  company_name VARCHAR(255) NOT NULL,
  location VARCHAR(255) NOT NULL,
  salary INT,
  posted_date DATE NOT NULL,
  category VARCHAR(255) NOT NULL,
  is_active BOOLEAN DEFAULT TRUE,
  user_id INT REFERENCES Users(id)
);

-- Create table for Users
CREATE TABLE Users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password_hash CHAR(60) NOT NULL,
  user_type ENUM('job_seeker', 'employer') NOT NULL,
  phone_number VARCHAR(20),
  rural_area VARCHAR(255) -- To capture details like village or town name
);

-- Additional tables can be created to store additional information as needed

-- Example: Table to store skills relevant to job seekers
CREATE TABLE Skills (
  id INT PRIMARY KEY AUTO_INCREMENT,
  skill_name VARCHAR(255) NOT NULL,
  user_id INT REFERENCES Users(id)
);

-- Example: Table to map skills to job listings
CREATE TABLE Job_Skills (
  job_id INT REFERENCES Jobs(id),
  skill_id INT REFERENCES Skills(id),
  PRIMARY KEY (job_id, skill_id)
);
