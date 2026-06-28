/*
Package inputport defines the driving ports (inbound interfaces) of the application core.

According to Hexagonal Architecture (Ports & Adapters) principles, this package contains the
abstract contracts and boundary specifications that define how external actors—such as HTTP APIs,
gRPC services, CLI commands, or event consumers—can trigger the application's use cases.

# Architectural Role

Input ports act as the formal entry points to the application's core business logic. They decouple
the execution of business capabilities from the transmission protocol. This ensures that the core
remains unaffected whether a request originates from a REST endpoint, an AMQP broker message,
or a cron job.

# Validation & Execution Boundaries

Primary entry points (driving adapters) are responsible for data parsing and basic protocol
validation, but they must pass execution control immediately to the input ports. In turn,
the application layer behind these ports guarantees semantic data consistency and returns
strongly typed, technology-agnostic errors back to the caller.
*/
package inputport
