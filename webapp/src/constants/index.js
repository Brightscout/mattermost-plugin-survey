import {PLUGIN_NAME} from './manifest';
import {ACTION_TYPES} from './action_types';

const OPEN_QUESTION_INITIAL_ROWS = 5;
const OPEN_QUESTION_MAX_LENGTH = 2000;

const QUESTION_TYPES = {
    OPEN: 'open',
    FIVE_POINT_LIKERT_SCALE: 'five-point-likert-scale',
};

const ERROR_MESSAGES = {
    GET_SURVEY: ' There was an error while retrieving the survey. Please try again later. If the problem persists, contact your System Administrator.',
    SUBMIT_SURVEY: ' There was an error while submitting your responses. Please try again later. If the problem persists, contact your System Administrator.',
    VALIDATE_SURVEY: 'Please provide your answers above.',
};

const FIVE_POINT_LIKERT_SCALE_RESPONSES = [
    {
        value: '1',
        text: 'Strongly Agree',
    },
    {
        value: '2',
        text: 'Agree',
    },
    {
        value: '3',
        text: 'Neutral',
    },
    {
        value: '4',
        text: 'Disagree',
    },
    {
        value: '5',
        text: 'Strongly Disagree',
    },
];

//
// Export the constants
//
export default {
    ACTION_TYPES,
    ERROR_MESSAGES,
    FIVE_POINT_LIKERT_SCALE_RESPONSES,
    MATTERMOST_CSRF_COOKIE: 'MMCSRF',
    OPEN_QUESTION_INITIAL_ROWS,
    OPEN_QUESTION_MAX_LENGTH,
    PLUGIN_NAME,
    QUESTION_TYPES,
    METRICS_TYPE_MEETING_METRICS: 'meetingMetrics',
};
