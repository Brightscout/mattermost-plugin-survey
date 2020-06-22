import {PLUGIN_NAME} from './manifest';

export const ACTION_TYPES = {
    OPEN_SURVEY_MODAL: `${PLUGIN_NAME}_open_survey_modal`,
    CLOSE_SURVEY_MODAL: `${PLUGIN_NAME}_close_survey_modal`,
    SET_CURRENT_POST_ID: `${PLUGIN_NAME}_set_current_post_id`,
    SET_CURRENT_POST_PROPS: `${PLUGIN_NAME}_set_current_post_props`,
    RECEIVED_DASHBOARD_PATH: `${PLUGIN_NAME}_received_dashboard_path`,

    // From mattermost-plugin-riff-video-chat.
    // TODO: Update if the name of the other plugin changes
    OPEN_RIFF_DASHBOARD: 'com.rifflearning.video-chat_OPEN_RIFF_METRICS_MODAL',
};
