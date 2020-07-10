# Survey Schema Design

## Survey

Stores the survey information. All surveys with different IDs, and all the different versions of surveys with the same ID, are maintained as different entries in the DB.

```
key: hash[survey_<survey_id>_<survey_version>]
value: {
    type: "survey",
    id: "string",
    version: <number>,
    created_at: <timestamp>,
    title: "string",
    description: "string",
    questions: [
        {
            id: "string",
            type: oneOf("five-point-likert-scale", "open"),
            text: "string"
        }
    ]
}
```

### Additional Notes

- Survey id would be a random alphanumeric uuid.
- Survey version would be a number, starting from 0, incremented on each update.
- For v1 plugin implementation:
  - Creation of only a single survey would be supported.
  - This survey would have a predetermined value of the “id” field: `f298903f8a80054ba09e342d0d9780635d3675a2` and the value of the “version’ field would start with 1.
  - Each time the plugin is started after updating the survey title, description and/or questions, the survey version would be incremented by 1.
  - The ID specific to each question would be a random alphanumeric uuid.

## Survey Response

Stores the user's response for the survey. A unique entry would be created for each unique survey response by the user for a meeting.

```
key: hash[survey_response_<user_id>_<meeting_id>_<survey_id>_<survey_version>]
value: {
    type: "survey_response",
    user_id: "string",
    meeting_id: "string",
    survey_id: "string",
    survey_version: <number>,
    created_at: <timestamp>,
    responses: {
        <question-id1>: "response1",
        ...
    }
}
```

### Additional Notes

- As of the current implementation, the user may submit a response through an API call even if they were not actually the part of the meeting.
- The user may submit the response as many times he wants. Each time, the same response entry in the DB is updated.
- The `created_at` field stores the time when the user submitted their response.
- The responses would be a map of questions ids from the survey to the users responses.
  - Open question response would be stored as string with the content user added to the textarea. It may be an empty string.
  - Likert question response would be a number string from "1" (Strongly Agree) - "5" (Strongly Disagree) or an empty string if the user did not add their response for that question.

## Latest Survey Info

Stores the information about the latest version for each survey. A unique entry is created for each survey with a different ID.

```
key: hash[“latest_survey_<survey_id>”]
value:  {
    type: "latest_survey_info",
    survey_id: "string",
    survey_version: <number>
}
```

## Meeting Metadata

Stores the information about survey sent for each meeting. A unique entry is created for each survey sent for a meeting with different meeting ID.

```
key: hash[meeting_metatata_<meeting_id>]
value:  {
    type: "meeting_metadata",
    meeting_id: "string",
    survey_id: "string",
    survey_version: <number>
}
```


### Additional Notes

- This entry makes sure that the same survey is sent to all the meeting participants of the same meeting.

## User Meeting Metadata

Stores the metadata information about surveys sent to a user. A unique entry is created for each user and meeting pair.

```
key: user_meeting_metadata_<user_id>_<meeting_id>
value:  {
    type: "meeting_metadata",
    user_id: "string",
    meeting_id: "string",
    survey_sent_at: <timestamp>,
    responded_at: <timestamp>
}
```

## Reminder Metadata

Stores the information about reminders. A unique entry is created for each survey post sent to the user.

```
key: reminder_metadata_<survey_post_id>
value:  {
    type: "meeting_metadata",
    meeting_id: "string",
    user_id: "string",
    post_id: "string",
    channel_id: "string",
    survey_sent_at: <timestamp>,
    previous_reminder_sent_at: <timestamp>,
    total_reminders_sent: <number>
}
```

## Notes

1. All these DB entries for various schemas will be in a single table in mattermost.

2. The table name would be:
    - ‘pluginkeyvaluestore’ for Postgres.
    - ‘PluginKeyValueStore’ for MySQL.

3. This table contains 3 columns:
    `pluginid | pkey | pvalue`

    - Where ‘pluginid’ is the ID of the plugin for this entry, ‘pkey’ is a string denoting the key and ‘pvalue’ is the value in bytes.
    - You can check the plugin id for this plugin in the `plugin.json` file in the root of this repo.

4. The data for KV Store table is stored as an array of bytes.
    - For Postgres DB, directly querying the table would result unintelligible data. A slightly modified query would be required:
        `select pkey, encode(pvalue, 'escape') from pluginkeyvaluestore;`
    - This would not be required for MySQL and a normal query for the value would work.

5. One of the entries in this table would just be an array of strings (reminders_list) and would not have a `type` field. This is being used to programmatically identify currently pending reminders.

6. All other entries in this table would have a “type” field. This field can be used to determine the type of the entry. It can have values:
    - "survey"
    - "survey_response"
    - "latest_survey_info"
    - "meeting_metadata"
    - "user_meeting_metadata"
    - "reminder_metadata"

7. Values stored in the DB for likert scale type questions would be:
    - {value: '1', text: 'Strongly Agree'},
    - {value: '2', text: 'Agree'},
    - {value: '3', text: 'Neutral'},
    - {value: '4', text: 'Disagree'},
    - {value: '5', text: 'Strongly Disagree'},
