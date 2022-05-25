-- public.community definition

-- Drop table

-- DROP TABLE public.community;

CREATE TABLE public.community (
	communityid int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	communityname varchar(50) NOT NULL,
	image bytea NULL,
--	colortheme varchar(30) NULL,
--	isactive bool NOT NULL,
	communitymgrid int4 NULL,
	CONSTRAINT community_pkey PRIMARY KEY (communityid)
);


-- public.joblevel definition

-- Drop table

-- DROP TABLE public.joblevel;

CREATE TABLE public.joblevel (
	joblevelid int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	jobleveldesc varchar(100) NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT joblevel_pkey PRIMARY KEY (joblevelid)
);


-- public.peopledetailsdesc definition

-- Drop table

-- DROP TABLE public.peopledetailsdesc;

CREATE TABLE public.peopledetailsdesc (
	peopledetailsdescid int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	peopledetailsdesc varchar(100) NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT peopleotherdetailsdesc_pkey PRIMARY KEY (peopledetailsdescid)
);


-- public.peopleskills definition

-- Drop table

-- DROP TABLE public.peopleskills;

CREATE TABLE public.peopleskills (
	peopleskillsid int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	peopleskillsdesc varchar(100) NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT peopleskills_pkey PRIMARY KEY (peopleskillsid)
);


-- public.project definition

-- Drop table

-- DROP TABLE public.project;

CREATE TABLE public.project (
	projectid int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	projectdesc varchar(100) NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT project_pkey PRIMARY KEY (projectid)
);


-- public.workstate definition

-- Drop table

-- DROP TABLE public.workstate;

CREATE TABLE public.workstate (
	workstateid int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	workstatedescription varchar(100) NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT workstate_pkey PRIMARY KEY (workstateid)
);


-- public.people definition

-- Drop table

-- DROP TABLE public.people;

CREATE TABLE public.people (
	peopleid int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	cognizantid int4 NOT NULL,
	lastname varchar(50) NOT NULL,
	firstname varchar(40) NOT NULL,
	middlename varchar(40) NOT NULL,
	fullname varchar(130) NOT NULL,
	csvemail varchar(50) NOT NULL,
	hireddate date NOT NULL,
	communityid int4 NULL,
	workstateid int4 NULL,
	joblevelid int4 NULL,
	projectid int4 NULL,
	isactive bool NOT NULL,
	isprobationary bool NOT NULL DEFAULT false,
	CONSTRAINT people_pkey PRIMARY KEY (peopleid),
	CONSTRAINT community_fkey FOREIGN KEY (communityid) REFERENCES public.community(communityid),
	CONSTRAINT joblevel_fkey FOREIGN KEY (joblevelid) REFERENCES public.joblevel(joblevelid),
	CONSTRAINT project_fkey FOREIGN KEY (projectid) REFERENCES public.project(projectid),
	CONSTRAINT workstate_fkey FOREIGN KEY (workstateid) REFERENCES public.workstate(workstateid)
);


-- public.peopledetails definition

-- Drop table

-- DROP TABLE public.peopledetails;

CREATE TABLE public.peopledetails (
	peopledetailsid int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	peopleid int4 NOT NULL,
	peopledetailsdescid int4 NOT NULL,
	activeflag bool NOT NULL,
	CONSTRAINT peopledetails_pkey PRIMARY KEY (peopledetailsid),
	CONSTRAINT peopleid_fkey FOREIGN KEY (peopleid) REFERENCES public.people(peopleid),
	CONSTRAINT peopleotherdetailsid_fkey FOREIGN KEY (peopledetailsdescid) REFERENCES public.peopledetailsdesc(peopledetailsdescid)
);


-- public.peopleprimaryskills definition

-- Drop table

-- DROP TABLE public.peopleprimaryskills;

CREATE TABLE public.peopleprimaryskills (
	peopleprimaryskillsid int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	peopleid int4 NOT NULL,
	peopleskillsid int4 NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT skillset_pkey PRIMARY KEY (peopleprimaryskillsid),
	CONSTRAINT peopleid_fkey FOREIGN KEY (peopleid) REFERENCES public.people(peopleid),
	CONSTRAINT peopleskills_fkey FOREIGN KEY (peopleskillsid) REFERENCES public.peopleskills(peopleskillsid)
);


-- public.peopletechnicalinterest definition

-- Drop table

-- DROP TABLE public.peopletechnicalinterest;

CREATE TABLE public.peopletechnicalinterest (
	peopletechnicalinterestid int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	peopleid int4 NOT NULL,
	peopleskillsid int4 NOT NULL,
	isactive bool NOT NULL,
	CONSTRAINT peopletechnicalinterest_pkey PRIMARY KEY (peopletechnicalinterestid),
	CONSTRAINT peopleid_fkey FOREIGN KEY (peopleid) REFERENCES public.people(peopleid),
	CONSTRAINT peopleskills_fkey FOREIGN KEY (peopleskillsid) REFERENCES public.peopleskills(peopleskillsid)
);