# Project Management Service

## Introduction
The Project Management Service is designed to streamline the planning, execution, and monitoring of projects. It offers tools to manage tasks, timelines, and resources, ensuring that projects are completed efficiently and effectively.

## Features
- Create and update project
- Clone other project
- Resource management

## Installation Dependencies
Ensure you have Go installed. After cloning the project, install the necessary packages and libraries with the following command:

```bash
go mod tidy
```

## Environment Variables
Create a `.env` file in the root directory and configure it based on `.env.example`. This file should contain all necessary environment-specific configurations.

## Running the Project
To run the project, use the following command:

```bash
make restart
```

This command will build and start the service.

## Docker
### Building the Docker Image and Running the Container
To build the Docker image and run the container, execute:

```bash
make up
```
