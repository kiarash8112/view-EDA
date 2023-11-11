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
goka seperate to three section 

![goka](https://github.com/kiarash8112/view-EDA/assets/133909368/54ecd6f4-f815-4494-80ee-201f6a6f87ef)

 
