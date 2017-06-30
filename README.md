# stateful

An example of how to manage global resources in a microservices world.

# Tenets

- Lazy dependencies
	- Start instantly without regard for dependent services' status.
	 - i.e. All routes respond to requests as soon as the service is running.
- Never exit/fail under any circumstance under the service's control.
- Try to handle errors from dependent services graceful before propagating them.
