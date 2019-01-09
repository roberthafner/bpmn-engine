create table BR_DEPLOYMENT (
  DEPLOYMENT_ID varchar(64),
  DEPLOYMENT_NAME varchar(255),
  DEPLOYMENT_TIME timestamp,
  primary key (DEPLOYMENT_ID)
);

create table BR_DEPLOYMENT_RESOURCE (
  RESOURCE_ID varchar(64),
  RESOURCE_NAME varchar(255),
  DEPLOYMENT_ID varchar(64),
  RESOURCE_BYTES bytea,
  primary key (RESOURCE_ID),
  foreign key (DEPLOYMENT_ID) REFERENCES BR_DEPLOYMENT
);