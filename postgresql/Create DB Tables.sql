-- public.community definition

-- Drop table

-- DROP TABLE community;

CREATE TABLE community (
	communityid int NOT NULL GENERATED ALWAYS AS IDENTITY,
	communityname varchar(50) NOT NULL,
	communityicon text NULL,
	communitymgrid int NULL,
	communitydesc varchar(250) NULL,
	isactive bool NOT NULL,
	CONSTRAINT community_pkey PRIMARY KEY (communityid)
);


-- public.communityadminandmanager definition

-- Drop table

-- DROP TABLE communityadminandmanager;

CREATE TABLE communityadminandmanager (
	communityadminandmanagerid int NOT NULL GENERATED ALWAYS AS IDENTITY,
	communityadminandmanagername varchar(100) NOT NULL,
	cognizantid varchar(10) NOT NULL,
	csvemail varchar(50) NOT NULL,
	password varchar(100) NOT NULL,
	roletype varchar(10) NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT communityadminandmanager_pkey PRIMARY KEY (communityadminandmanagerid)
);


-- public.joblevel definition

-- Drop table

-- DROP TABLE joblevel;

CREATE TABLE joblevel (
	joblevelid int NOT NULL GENERATED ALWAYS AS IDENTITY,
	jobleveldesc varchar(100) NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT joblevel_pkey PRIMARY KEY (joblevelid)
);


-- public.peopledetailsdesc definition

-- Drop table

-- DROP TABLE peopledetailsdesc;

CREATE TABLE peopledetailsdesc (
	peopledetailsdescid int NOT NULL GENERATED ALWAYS AS IDENTITY,
	peopledetailsdesc varchar(100) NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT peopleotherdetailsdesc_pkey PRIMARY KEY (peopledetailsdescid)
);


-- public.peopleskills definition

-- Drop table

-- DROP TABLE peopleskills;

CREATE TABLE peopleskills (
	peopleskillsid int NOT NULL GENERATED ALWAYS AS IDENTITY,
	peopleskillsdesc varchar(100) NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT peopleskills_pkey PRIMARY KEY (peopleskillsid)
);


-- public.project definition

-- Drop table

-- DROP TABLE project;

CREATE TABLE project (
	projectid int NOT NULL GENERATED ALWAYS AS IDENTITY,
	projectdesc varchar(100) NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT project_pkey PRIMARY KEY (projectid)
);


-- public.workstate definition

-- Drop table

-- DROP TABLE workstate;

CREATE TABLE workstate (
	workstateid int NOT NULL GENERATED ALWAYS AS IDENTITY,
	workstatedescription varchar(100) NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT workstate_pkey PRIMARY KEY (workstateid)
);


-- public.people definition

-- Drop table

-- DROP TABLE people;

CREATE TABLE people (
	peopleid int NOT NULL GENERATED ALWAYS AS IDENTITY,
	cognizantid int NOT NULL,
	lastname varchar(50) NOT NULL,
	firstname varchar(40) NOT NULL,
	middlename varchar(40) NOT NULL,
	fullname varchar(130) NOT NULL,
	csvemail varchar(50) NOT NULL,
	hireddate date NOT NULL,
	communityid int NULL,
	communityadminandmanagerid int NULL,
	workstateid int NULL,
	joblevelid int NULL,
	projectid int NULL,
	isprobationary bool NOT NULL DEFAULT false,
	isactive bool NOT NULL,
	CONSTRAINT people_pkey PRIMARY KEY (peopleid),
	CONSTRAINT community_fkey FOREIGN KEY (communityid) REFERENCES community(communityid),
	CONSTRAINT communityadminandmanager_fkey FOREIGN KEY (communityadminandmanagerid) REFERENCES communityadminandmanager(communityadminandmanagerid),
	CONSTRAINT joblevel_fkey FOREIGN KEY (joblevelid) REFERENCES joblevel(joblevelid),
	CONSTRAINT project_fkey FOREIGN KEY (projectid) REFERENCES project(projectid),
	CONSTRAINT workstate_fkey FOREIGN KEY (workstateid) REFERENCES workstate(workstateid)
);


-- public.peopledetails definition

-- Drop table

-- DROP TABLE peopledetails;

CREATE TABLE peopledetails (
	peopledetailsid int NOT NULL GENERATED ALWAYS AS IDENTITY,
	peopleid int NOT NULL,
	peopledetailsdescid int NOT NULL,
	activeflag bool NOT NULL,
	CONSTRAINT peopledetails_pkey PRIMARY KEY (peopledetailsid),
	CONSTRAINT peopledetailspeopleid_fkey FOREIGN KEY (peopleid) REFERENCES people(peopleid),
	CONSTRAINT peopleotherdetailsid_fkey FOREIGN KEY (peopledetailsdescid) REFERENCES peopledetailsdesc(peopledetailsdescid)
);


-- public.peopleprimaryskills definition

-- Drop table

-- DROP TABLE peopleprimaryskills;

CREATE TABLE peopleprimaryskills (
	peopleprimaryskillsid int NOT NULL GENERATED ALWAYS AS IDENTITY,
	peopleid int NOT NULL,
	peopleskillsid int NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT skillset_pkey PRIMARY KEY (peopleprimaryskillsid),
	CONSTRAINT peopleprimaryskillspeopleid_fkey FOREIGN KEY (peopleid) REFERENCES people(peopleid),
	CONSTRAINT peopleskills_fkey FOREIGN KEY (peopleskillsid) REFERENCES peopleskills(peopleskillsid)
);


-- public.peopletechnicalinterest definition

-- Drop table

-- DROP TABLE peopletechnicalinterest;

CREATE TABLE peopletechnicalinterest (
	peopletechnicalinterestid int NOT NULL GENERATED ALWAYS AS IDENTITY,
	peopleid int NOT NULL,
	peopleskillsid int NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT peopletechnicalinterest_pkey PRIMARY KEY (peopletechnicalinterestid),
	CONSTRAINT peopleskills_fkey FOREIGN KEY (peopleskillsid) REFERENCES peopleskills(peopleskillsid),
	CONSTRAINT peopletechnicalinterestpeopleid_fkey FOREIGN KEY (peopleid) REFERENCES people(peopleid)
);