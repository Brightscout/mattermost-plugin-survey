import Constants from '../constants';

const INITIAL_STATE = {
    visible: false,
    postID: '',
    meetingID: '',
    surveyID: '',
    surveyVersion: '',
};

export const survey = (state = INITIAL_STATE, action) => {
    switch (action.type) {
    case Constants.ACTION_TYPES.OPEN_SURVEY_MODAL:
        return {
            visible: true,
            postID: action.postID,
            meetingID: action.meetingID,
            surveyID: action.surveyID,
            surveyVersion: action.surveyVersion,
        };
    case Constants.ACTION_TYPES.CLOSE_SURVEY_MODAL:
        return INITIAL_STATE;
    default:
        return state;
    }
};
