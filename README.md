# features
- login/signup
- load crypto wallet
- load credit/debit card
- betting page
- add a poll so that we can bet on the item that maximum gamers want


# betting page
- item picture/video on the left
- betting being done on the right
- there should be a section for next-item in the page using a queue data structure

# main algorithm
- item change every 1 hour
- if someone matches current bet, it gets increased by 50 cents, if 2 people place the order, then the first one gets executed first and remaining get the message that the bet has already been placed

There should be a button 'Accept Bid', this should do the following:
- give an alert message if she wants to complete the bid
- upon clicking YES, deposit the money into the bank account
- change UI in the right-side of the page with the user's name
- change the bid price
- give 'bid already placed' message to other's

# communication between backend and frontend
-  create an API
- upon clicking button, it should hit an API path to perform a certain job in the backend