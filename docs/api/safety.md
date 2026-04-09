# User Safety

Content reporting.

## Report Content

Report a message, server, or user.

```
POST /safety/report
Auth: Session Token
Body: DataReportContent
Response 204: Success
```

### DataReportContent

```json
{
  "content": { /* ReportedContent */ },  // required
  "additional_context": "string"          // optional — additional details
}
```

### ReportedContent

Tagged union:

**Report a message:**
```json
{
  "type": "Message",
  "id": "message_id",
  "report_reason": "NoneSpecified"  // ContentReportReason
}
```

**Report a server:**
```json
{
  "type": "Server",
  "id": "server_id",
  "report_reason": "NoneSpecified"  // ContentReportReason
}
```

**Report a user:**
```json
{
  "type": "User",
  "id": "user_id",
  "report_reason": "NoneSpecified",  // UserReportReason
  "message_id": "string"              // optional — context message
}
```

### ContentReportReason

- `NoneSpecified`
- `Illegal`
- `IllegalGoods`
- `IllegalExtortion`
- `IllegalPornography`
- `IllegalHacking`
- `ExtremeViolence`
- `PromotesHarm`
- `UnsolicitedSpam`
- `Raid`
- `SpamAbuse`
- `ScamsFraud`
- `Malware`
- `Harassment`

### UserReportReason

- `NoneSpecified`
- `UnsolicitedSpam`
- `SpamAbuse`
- `InappropriateProfile`
- `Impersonation`
- `BanEvasion`
- `Underage`
