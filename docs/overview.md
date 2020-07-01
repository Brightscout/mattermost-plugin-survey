# Overview

A mattermost plugin to send surveys to the user for Riff edu meetings.

- A survey is sent to the meeting participants just after they leave a meeting by closing the browser tab or through the leave call button.
- The current implementation does not check if the user was actually a part of the meeting as we don't interact with the riff meeting server. Such filtering should be done while analysing the responses and the metadata.

- The survey is sent through a bot user configured in system console settings.

- For each survey, sent to the user, the reminder settings can be configured.
  - When the reminders feature is enabled, a reminder to fill the survey would be sent to the user as a reply of the 'submit survey' post.
  - Each reminder would only be sent if the user has not responded to the survey yet.
  - The message for the reminder, interval between reminders and the maximum number of reminders sent to the user can be configured in the system console settings for the plugin.

- Each time the plugin is enabled, the plugin checks for the latest version of the survey with id: `f298903f8a80054ba09e342d0d9780635d3675a2` in the DB.

- If the survey does not exist in the DB, a new entry is created with survey version set to one.

- Each time you update one or more fields for the survey in the system console settings, a new version would be created for the save SurveyID.

- Support for multiple surveys with different IDs is planned for v2.
