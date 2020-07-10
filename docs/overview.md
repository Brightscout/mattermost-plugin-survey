# Overview

A mattermost plugin to send surveys to the user for Riff edu meetings.

- This plugin allows the meeting participants to fill a survey for the riff edu meeting they just participated in.
- This survey can be filled just after the meeting is finished or it can be sent as a direct message if they don't want to immediately fill out the survey.
- The current implementation does not check if the user was actually a part of the meeting as we don't interact with the riff meeting server. Such filtering should be done while analysing the responses and the metadata.
- Each user can update their responses indefinite number of times through an API request. The option to fill out the survey in the UI is only presented after leaving the meeting.
- If the survey is to be sent as a DM, it is sent through a bot user with Username: `riffbot`.
- For each survey, sent to the user, the reminder settings can be configured.
  - When the reminders feature is enabled, a reminder to fill the survey would be sent to the user as a reply of the `Submit Survey` post.
  - Each reminder would only be sent if the user has not responded to the survey yet.
  - The message for the reminder, interval between reminders and the maximum number of reminders sent to the user can be configured in the system console settings for the plugin.

- Each time the plugin is enabled, the plugin checks for the latest version of the survey with id: `f298903f8a80054ba09e342d0d9780635d3675a2` in the DB.
- If the survey does not exist in the DB, a new entry for survey is created with survey version set to one.
- Each time you update one or more fields for the survey in the system console settings, a new version would be created for the same SurveyID.
- Support for multiple surveys with different IDs is a feature for v2.
