# GoTraining1
I tried to create an organized folder structure... maybe a bit overengineered. 

Currently presenting a counter at /counter
Serving at port 8090
Typical call http://localhost:8090/counter

Missing:
* Control the limit of calls
* Stress test
* Integration test (Calling the server end-end)

# Project Structure
The project follows the next structure

```
├───api
│   ├───handlers
│   └───serverfactory
├───application
│   └───services
└───infrastructure
    └───server
```
 * application
   * services: The business services. Right now the counter is placed here. I unit test all services.
 * api
   * handlers: The web controllers. We have one handler per service published. And also the aggregation of all handlers into the muxer. I added test for the muxer to check the routes are correct.
   * serverfactory: I uses the server at infrastucture, injects the handler and starts it.
* infrastructure
  * server: A basic http server. I expect to add inside the protections for DDOS.