# TutorMe - A Tutoring App for Schools

TutorMe is an innovative app designed to streamline tutoring services in schools. By utilizing advanced multitenant architecture and cloud technologies, TutorMe ensures data isolation, scalability, and security for schools managing their tutoring programs.

## Features
- **Multitenant Database Model**: Each school is provisioned with a separate Turso database to ensure complete data isolation.
- **DNS Management with Cloudflare**: Automatically generates subdomains for schools, such as `schoolname.tutorme.tech`, enhancing security and accessibility.
- **Security and Performance**: Implements reverse proxy setups and Docker containers for efficient resource management and data safety.
- **Deployed on a VPS**: Ensures stability, scalability, and cost-effectiveness.

## Current Status
The app is currently under development, with an API accessible at [api.tutorme.tech](https://api.tutorme.tech). Here's what has been accomplished so far:
- **Server Setup**: Deployed on a Virtual Private Server with a reverse proxy and Docker containerization.
- **Database Configuration**: Initial setup of Turso databases for tenant isolation.
- **Subdomain Generation**: Automated using Cloudflare DNS management for seamless and secure onboarding.

## Upcoming Features
- User authentication and authorization tailored for students and tutors.
- A user-friendly front-end interface.
- Enhanced logging and monitoring for administrators.
- Integration with external educational tools and platforms.

## Deployment
The application is hosted on a VPS, leveraging:
- Docker swarm for remote containerized deployment.
- Traefik for reverse proxy setup and ssl management.
- Cloudflare for DNS and SSL management.

## Contributing
Contributions are welcome! Whether you're interested in frontend development, backend enhancements, or DevOps, there's plenty of room to collaborate.

To contribute:
1. Clone the repository.
2. Set up a local development environment (guidelines coming soon).
3. Open a pull request with detailed explanations for any changes.

## License
This project is licensed under the MIT License. See the `LICENSE` file for more details.
