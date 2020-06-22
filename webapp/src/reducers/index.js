import {combineReducers} from 'redux';

import {
    currentPostID, currentPostProps, dashboardPath, surveyModalVisible,
} from './survey';

export default combineReducers({
    currentPostID,
    currentPostProps,
    dashboardPath,
    surveyModalVisible,
});
