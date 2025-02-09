API Endpoints
-------------
1. Send an Email
   -------------
Endpoint: POST /v1/email/send

Description: This endpoint queues an email for sending. The email details such as recipient, sender, subject, and body are provided in the request payload.

Request Payload (JSON):
{
  "to": "sakshichalpe51@gmail.com",
  "from": "xyz@gmail.com",
  "subject": "Testing",
  "body": "I am Queen"
}
Response Payload (JSON):
{
  "message": "Email queued successfully"
}

Response Explanation:If the email is successfully added to the queue, a success message is returned.The actual sending process may be handled asynchronously, ensuring efficient email processing.

2. Get Email Statistics
   --------------------
Endpoint:GET /v1/email/statistics
Description:This endpoint returns email sending statistics, including the total number of emails sent and any failures encountered.
Response Payload (JSON):
{
  "TotalEmailsSent": 2,
  "TotalFailures": 0,
  "Mutex": {}
}

Response Explanation:
TotalEmailsSent: The number of emails successfully processed and sent.
TotalFailures: The number of emails that failed to send.
Mutex: A synchronization mechanism (likely used internally in the application) to manage concurrent access to email statistics, ensuring thread-safe operations.
