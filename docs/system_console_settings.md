# System Console Settings

When this plugin is installed, plugin specific system console settings page is visible through the `/admin_console/plugins/plugin_survey` route of your mattermost instance.

Here you can configure and customise the functionality of the plugin.
These settings are validated by the plugin. If any one of these are invalid, the plugin would fail to start.

The information about which validation is failing can be seen in the mattermost server logs: `sudo journalctl -f -u mattermost`.

Here is a list of settings visible in the system console:

1. **Survey**:
    - This is the schema of the survey sent to the meeting participants.
    - It is a JSON object with the following fields:
        - title: (string) This is the title of the survey displayed on the header of the survey modal.
        - description: (string) This is a descriptive text visible on opening the survey modal.
        - questions (array) This is an array of questions. Each question is a JSON object with properties: type and text.
            - type: The type of the question. Currently supported types are:
                1. five-point-likert-scale: A five point likert type question with options ranging from strongly agree to strongly disagree.
                2. open: An open ended question with a textarea
            - text: The question text.

    - Example Schema:

        ```json
            {
                "title": "Survey Title",
                "description": "Survey Description",
                "questions": [
                    {
                        "type": "five-point-likert-scale",
                        "text": "What is your rating?"
                    },
                    {
                        "type": "open",
                        "text": "Please add any other comments."
                    }
                ]
            }
        ```

2. **Reminder Text**:
    - The message sent to the user as a reminder to fill the survey.
    - A reminder will only be sent to a user who has not yet submitted the survey yet.

3. **Max Reminder Count**:
    - The maximum number of times a reminder would be sent to the user.
    - This can be set to `0` to disable reminders.

4. **Reminder Interval**:
    - The time interval in minutes after the survey, or the last reminder for this survey, is sent to the user, after which the next reminder would be sent.
