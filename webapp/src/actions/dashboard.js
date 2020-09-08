import Constants from '../constants';

export const openRiffDashboard = (meetingID) => (dispatch) => {
    dispatch({
        type: Constants.ACTION_TYPES.OPEN_RIFF_DASHBOARD,
        meetingID,

        // Always show meeting-specific metrics when showing from the survey post link
        selectedMetrics: Constants.METRICS_TYPE_MEETING_METRICS,
    });
};
