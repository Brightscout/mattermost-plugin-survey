import Constants from '../constants';

export const openRiffDashboard = (meetingID) => (dispatch) => {
    dispatch({
        type: Constants.ACTION_TYPES.OPEN_RIFF_DASHBOARD,
        meetingID,
    });
};
