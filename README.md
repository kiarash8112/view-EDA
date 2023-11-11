# view-EDA
in this repo i implemented simple event-driven architecture with help of [CQRS pattern](https://docs.aws.amazon.com/prescriptive-guidance/latest/modernization-data-persistence/cqrs-pattern.html) for following problem :

# Problem & Architecture
user-request fulfilment is enabled by multiple distributed services, however the
system needs to present a view about the overall request. For example, in the hotel booking
flow, the user would like to see their booking as it progresses through various states, such
as INITIAL, PAYMENT_MADE, and BOOKING_CONFIRMED, along with ancillary
details such as the check-in time at the hotel. All this data is not available with one service,
thus to prepare the view, one naive approach might be to query all the services for the data
and compose the data. This is not always optimal since the service usually models data in
the format that it requires it, not this third-party requirement. The alternative is, in advance,
to pre-populate the data for the view in a format best suited to the view. Here is an
example:

![view](https://github.com/kiarash8112/view-EDA/assets/133909368/4f0d6132-d1d7-409f-a8c3-8f59ec33b8ae)

This pattern allows the decoupling promised by microservices, while still allowing the
composition of rich cross-cutting views.

# Implemention
i used [goka](https://github.com/lovoo/goka) package to implement this architecture 
goka seperate to three sections(details about each section is availabe [here](https://github.com/lovoo/goka/wiki/Introduction))

![goka](https://github.com/kiarash8112/view-EDA/assets/133909368/54ecd6f4-f815-4494-80ee-201f6a6f87ef)

 
i created request creator to simulate this scenario each request emit to emitter with two topics: payment,booking.

processor collect request through kafka message broker and update information in the database

after that view cache info and return it through http calls

# How to get it running 

```bash
# kafka and zookeeper must be running, 
# you can start them with docker-compose up

# run the example
go run .
```

This should output something like
```
View opened at http://localhost:9095/
2023/11/11 11:29:50 [Processor booking_group1] setup generation 1, claims=map[string][]int32{"booking1":[]int32{0, 1, 2, 3, 4, 5, 6, 7}}
2023/11/11 11:29:50 [Processor payment_group1] setup generation 1, claims=map[string][]int32{"payment1":[]int32{0, 1, 2, 3, 4, 5, 6, 7}}
```
Now open the browser and get the view for user-3: http://localhost:9095/user-3 (if you reload you get new requests results)
