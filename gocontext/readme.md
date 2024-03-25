### Context in Go:

* It is a inbuilt package provided by the go context library which provide more control to the developers while developing the applications.
* Software often kicks off long running, resource intnesive processes ofently in go routines.
* If the ation that caused this get cancelled or fails for some reason we need to stop the processes in a consistent way throught out the application.
* Context will be used to helps us managing the long running processes.
* It's important that you derive your contexts so that cancellations are propagated throughout the call stack for a given request.
* context has a method Done() which returns a channel which gets sent a signal when the context is "done" or "cancelled".
* Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propagate the Context
